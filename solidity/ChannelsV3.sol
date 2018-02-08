pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";

/*
	Will be used as the basis for V4 which will enable two parties to pay each other who payment have a channel with each other
		A --> B --> C
			A can pay C without having a direct path to C through its connection with B
*/
contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	string constant public VERSION = "0.3.1alpha";
	
	uint256 private channelCount;
	bytes private prefix = "\x19Ethereum Signed Message:\n32";
	// prevent any possible accidental triggering of developer only conditions
	bool	public	dev = true;

	enum ChannelStates { opened, expired, closed }

	ChannelStates public defaultState = ChannelStates.opened;

	struct ChannelStruct {
		address source;
		address destination;
		uint256 value;
		uint256 closingDate;
		uint256 openDate;
		bytes32 channelId;
		bool	sourceProofSubmitted;
		bool	destinationProofSubmitted;
		ChannelStates state;
	}

	mapping (bytes32 => ChannelStruct) public channels;
	mapping (uint256 => bytes32) private channelNumber;
	mapping (bytes32 => bool) private channelIds;
	// prevent resubmission of the same signed messages by a particular address within a channel
	mapping (bytes32 => mapping (bytes32 => mapping(address => bool))) private signedMessages;
	/**
		keeps track of micro payment proofs and prevent them from being reused.
		key 1 (bytes32) = channel ID
		key 2 (bytes32) = message hash
		val 1 (bool)    = boolean, true (message hash used), false (message hash not used)

	*/
	mapping (bytes32 => mapping (bytes32 => bool)) private microPaymentHashes;

	/** Micropayment proof
		hash: keccak256(channelId, paymentString, withdrawal amount)
		To generate on python:
			Web3.soliditySha3(['bytes32', 'string', 'uint256'], [Web3.toHex('foo'), 'bar', 200])
	*/

	event ChannelOpened(bytes32 indexed _channelId);
	event ChannelClosed(bytes32 indexed _channelId);
	event ChannelExpired(bytes32 indexed _channelId);
	event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue);

	modifier bothProofsSubmitted(bytes32 _channelId) {
		require(channels[_channelId].vendorProofSubmitted);
		require(channels[_channelId].purchaserProofSubmitted);
		_;
	}

	function () public payable {}



	function openChannel(
		address _destination,
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
		// make sure the channel id doens't already exist
		require(!channelIds[channelId]);
		channelIds[channelId] = true;
		channels[channelId].source = msg.sender;
		channels[channelId].destination = _destination;
		channels[channelId].value = _channelValueInWei;
		channels[channelId].closingDate = (now + (_durationInDays * 1 days));
		channels[channelId].openDate = currentDate;
		channels[channelId].channelId = channelId;
		channels[channelId].state = defaultState;
		ChannelOpened(channelId);
		return true;
	}

	function submitSourceProof(
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
		require(signer == channels[_channelId].source);
		channels[_channelId].sourceProofSubmitted = true;
		SourceProofSubmitted(_channelId, signer);
		return true;
	}

	function submitDestinationProof(
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
		require(signer == channels[_channelId].destination);
		channels[_channelId].vendorProofSubmitted = true;
		DestinationProofSubmitted(_channelId, signer);
		return true;
	}

	function closeChannel(
		bytes32 _channelId)
		public
		bothProofsSubmitted(_channelId)
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(channels[_channelId].state == ChannelStates.opened);
		require(channels[_channelId].value > 0);
		require(msg.sender == channels[_channelId].destination);
		uint256 deposit = channels[_channelId].value;
		channels[_channelId].value = 0;
		channels[_channelId].state = ChannelStates.closed;
		ChannelClosed(_channelId);
		msg.sender.transfer(deposit);
		return true;
	}

	function submitMicroPaymentProof(
		bytes32 _channelId,
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		uint256 _withdrawalAmount)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(channels[_channelId].state == ChannelStates.opened);
		require(msg.sender == channels[_channelId].destination);
		// prevent a micropayment from reducing the entire balance
		require(channels[_channelId].value > _withdrawalAmount && _withdrawalAmount > 0);
		bytes32  _proof = keccak256(_channelId, msg.sender, _withdrawalAmount);
		bytes32 proof = keccak256(prefix, _proof);
		assert(proof == _h);
		require(!microPaymentHashes[_channelId][_h]);
		microPaymentHashes[_channelId][_h] = true;
		uint256 remainingChannelValue = channels[_channelId].value.sub(_withdrawalAmount);
		channels[_channelId].value = remainingChannelValue;
		MicroPaymentWithdrawn(_channelId, _withdrawalAmount, remainingChannelValue);
		msg.sender.transfer(_withdrawalAmount);
		return true;
	}

	/**
		So long as both proofs have *NOT* been submitted, and it is past closing, terminate channel
	*/
	function expireChannel(
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(msg.sender == channels[_channelId].source);
		// require that the channel is open
		require(channels[_channelId].state == ChannelStates.opened);
		// if both proofs have been submitted, throw
		if (channels[_channelId].sourceProofSubmitted && channels[_channelId].destinationProofSubmitted) {
			revert();
		}
		if (!dev) {
			require(now >= channels[_channelId].closingDate);
		}
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