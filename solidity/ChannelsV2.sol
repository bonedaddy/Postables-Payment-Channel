pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Libs/EcRecovery.sol";

contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	string constant public VERSION = "0.0.2alpha";
	
	uint256 private channelCount;
	bool	public	dev = true;

	/** State Description
		Initial channel state is "opened"
		It becomes "accepted" once the vendor proof is submitted
		It becomes "expired" if channel opener times out channel due to lack of vendor proof
		It becomes "finalized" once both proofs have been submitted
		It becomes "closed" if the channel is succesfully closed peacefully

	*/					/** 0       1          2           3*/
	enum ChannelStates { opened, expired, finalized, closed }

	ChannelStates public defaultState = ChannelStates.opened;

	struct ChannelStruct {
		address purchaser;
		address vendor;
		uint256 value;
		uint256 closingDate;
		uint256 openDate;
		bytes32 channelId;
		bool	purchaserProofSubmitted;
		bool	vendorProofSubmitted;
		ChannelStates state;
	}

	mapping (uint256 => bytes32) private channelNumber;
	mapping (bytes32 => ChannelStruct) public channels;
	mapping (bytes32 => bool) private channelIds;
	// prevent resubmission of the same signed messages by a particular address within a channel
	mapping (bytes32 => mapping (bytes32 => mapping(address => bool))) private signedMessages;
	event ChannelOpened(address indexed _purchaser, address indexed _vendor, bytes32 indexed _channelId);
	event ChannelClosed(bytes32 indexed _channelId);
	event ChannelExpired(bytes32 indexed _channelId);
	event VendorProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event PurchaserProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);

	modifier bothProofsSubmitted(bytes32 _channelId) {
		require(channels[_channelId].vendorProofSubmitted);
		require(channels[_channelId].purchaserProofSubmitted);
		_;
	}

	function () public payable {}


	/**
		@dev Used to open a channel with the vendor
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
		channels[channelId].purchaser = msg.sender;
		channels[channelId].vendor = _vendor;
		channels[channelId].value = _channelValueInWei;
		channels[channelId].closingDate = (now + (_durationInDays * 1 days));
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
		require(channels[_channelId].state == ChannelStates.opened);
		require(!signedMessages[_channelId][_h][msg.sender]);
		signedMessages[_channelId][_h][msg.sender] = true;
		address signer = ecrecover(_h, _v, _r, _s);
		require(signer == channels[_channelId].purchaser);
		channels[_channelId].purchaserProofSubmitted = true;
		PurchaserProofSubmitted(_channelId, signer);
		return true;
	}

	/**
		Used to submit vendor proof by channel purchaser
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
		require(channels[_channelId].state == ChannelStates.opened);
		require(!signedMessages[_channelId][_h][msg.sender]);
		signedMessages[_channelId][_h][msg.sender] = true;
		address signer = ecrecover(_h, _v, _r, _s);
		require(signer == channels[_channelId].vendor);
		channels[_channelId].vendorProofSubmitted = true;
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
		Used when the vendor has not sent their proof,  and its been 2 weeks since the closing date,
		kill the channel
	*/
	function expireChannel(
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(msg.sender == channels[_channelId].purchaser);
		if (!dev) {
			require(now > (channels[_channelId].closingDate + 2 weeks));
		}
		require(channels[_channelId].state != ChannelStates.closed &&
			    channels[_channelId].state != ChannelStates.finalized &&
			    channels[_channelId].state != ChannelStates.expired);
		channels[_channelId].state = ChannelStates.expired;
		uint256 deposit = channels[_channelId].value;
		channels[_channelId].value = 0;
		ChannelExpired(_channelId);
		msg.sender.transfer(deposit);
		return true;
	}

	/**tested
		withdraw any ether in the contract, only when in developer mode
	*/
	function withdrawEth()
		public
		onlyAdmin
		returns (bool)
	{
		require(dev);
		msg.sender.transfer(this.balance);
		return true;
	}

}