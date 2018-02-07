pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Libs/EcRecovery.sol";
/**
	Add message hash verificaton when submitting a second signature.
	first signature to submit commits the message hash, subsequent signatures mush also submit a signature
	which verifies off that message hash. This should help reduce exploitation attempts

	To Do:
		> Implement channel timeout
		> Implement two in one proof submission for the vendor and channel open process

*/
contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	string constant public VERSION = "0.0.2alpha";
	
	uint256 private channelCount;
	bool	public	dev = true;

	/** State Description
		Initial channel state is "proposed"
		It becomes "accepted" once the vendor proof is submitted
		It becomes "expired" if channel opener times out channel due to lack of vendor proof
		It becomes "finalized" once both proofs have been submitted
		It becomes "closed" if the channel is succesfully closed peacefully

	*/
	enum ChannelStates { proposed, accepted, expired, finalized, closed }

	ChannelStates public defaultState = ChannelStates.proposed;

	struct ChannelStruct {
		address purchaser;
		address vendor;
		uint256 value;
		uint256 expirationDate;
		uint256 closingDate;
		uint256 openDate;
		bytes32 channelId;
		ChannelStates state;
		bool	bothProofsSubmitted;
		mapping (address => bool) proofSubmitted;
	}

	mapping (uint256 => bytes32) private channelNumber;
	mapping (bytes32 => ChannelStruct) public channels;
	mapping (bytes32 => bool) private channelIds;

	event ChannelOpened(address indexed _purchaser, address indexed _vendor, bytes32 indexed _channelId);
	event ChannelClosed(bytes32 indexed _channelId);
	event ChannelExpired(bytes32 indexed _channelId);
	event VendorProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event PurchaserProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);

	function () public payable {}


	/**
		Used to open a channeltwo-in-one vendor verification
		Useful in situations where the vendor and purchaser don't know each other.
	*/
	function openChannel(
		address _vendor,
		uint256 _channelValueInWei,
		uint256 _durationInDays)
		public
		payable
		returns (bool)
	{
		require(msg.value == _channelValueInWei);
		uint256 currentDate = now;
		// channel hash = keccak256(purchaser, vendor, channel value, date of open)
		bytes32 channelId = keccak256(msg.sender, _vendor, _channelValueInWei, currentDate);
		// make sure the channel id doens't already exit
		require(!channelIds[channelId]);
		channelIds[channelId] = true;
		uint256 autoClosureDate = (now + ((_durationInDays * 1 days) * 2));
		channels[channelId].purchaser = msg.sender;
		channels[channelId].vendor = _vendor;
		channels[channelId].value = _channelValueInWei;
		channels[channelId].expirationDate = autoClosureDate;
		channels[channelId].closingDate = _durationInDays * 1 days;
		channels[channelId].openDate = currentDate;
		channels[channelId].channelId = channelId;
		channels[channelId].state = defaultState;
		ChannelOpened(msg.sender, _vendor, channelId);
		return true;
	}

	/**
		Used by the vendor to accept the channel request
	*/
	function submitPurchaserProof(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(msg.sender == channels[_channelId].vendor);
		require(channels[_channelId].state == ChannelStates.proposed);
		channels[_channelId].state = ChannelStates.accepted;
		address signer = ecrecover(_h, _v, _r, _s);
		require(signer == channels[_channelId].purchaser);
		channels[_channelId].proofSubmitted[signer] = true;
		PurchaserProofSubmitted(_channelId, signer);
		return true;
	}

	/**
		Used to submit vendor proof
	*/
	function submitVendorProof(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(msg.sender == channels[_channelId].purchaser);
		require(channels[_channelId].state == ChannelStates.accepted);
		channels[_channelId].state = ChannelStates.finalized;
		address signer = ecrecover(_h, _v, _r, _s);
		require(signer == channels[_channelId].vendor);
		channels[_channelId].proofSubmitted[signer] = true;
		channels[_channelId].bothProofsSubmitted = true;
		VendorProofSubmitted(_channelId, signer);
		return true;
	}

	/**
		This is used by a vendor to close the channel and withdraw their funds
	*/
	function closeChannel(
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(channels[_channelId].state == ChannelStates.finalized);
		require(channels[_channelId].bothProofsSubmitted);
		require(channels[_channelId].value > 0);
		require(msg.sender == channels[_channelId].vendor);
		uint256 deposit = channels[_channelId].value;
		channels[_channelId].value = 0;
		channels[_channelId].state = ChannelStates.closed;
		ChannelClosed(_channelId);
		msg.sender.transfer(deposit);
		return true;
	}

	/**
		Used when the vendor has not sent their proof, and allows channel opener to recover their funds
	*/
	function expireChannel(
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelId[_channelId]);
		require(msg.sender == channels[_channelId].purchaser);
		require(channels[_channelId].state != ChannelStates.closed &&
			    channels[_channelId].state != ChannelStates.finalized &&
			    channels[_channelId].state != ChannelStates.expired);
		channels[_channelId].state = ChannelStates.expired;
		uint256 deposit = channels[_channelId].value;
		channels[_channelId].value = 0.
		ChannelExpired(_channelId);
		msg.sender.transfer(deposit);
		return true;
	}

	function withdrawEth()
		public
		onlyAdmin
		returns (bool)
	{
		require(dev);
		msg.sender.transfer(this.balance);
		return true;
	}

	/**GETTERS*/

	function getChannelPurchaser(
		bytes32 _channelId)
		public
		view
		returns (address)
	{
		return channels[_channelId].purchaser;
	}

	function getChannelVendor(
		bytes32 _channelId)
		public
		view
		returns (address)
	{
		return channels[_channelId].vendor;
	}

	function getChannelValue(
		bytes32 _channelId)
		public
		view
		returns (uint256)
	{
		return channels[_channelId].value;
	}

	function getExpirationDate(
		bytes32 _channelId)
		public
		view
		returns (uint256)
	{
		return channels[_channelId].expirationDate;
	}

	function getOpenDate(
		bytes32 _channelId)
		public
		view
		returns (uint256)
	{
		return channels[_channelId].openDate;
	}

	function getChannelState(
		bytes32 _channelId)
		public
		view
		returns (ChannelStates)
	{
		return channels[_channelId].state;
	}

	function checkIfProofSubmitted(
		bytes32 _channelId,
		address _addr)
		public
		view
		returns (bool)
	{
		return channels[_channelId].proofSubmitted[_addr];
	}

	function calculateChannelId(
		address _purchaser,
		address _vendor,
		uint256 _channelValueInWei,
		uint256 _date)
		public
		pure 
		returns (bytes32)
	{
		return keccak256(_purchaser, _vendor, _channelValueInWei, _date);
	}

}