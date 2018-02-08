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
						bytes32,   uint256,   uint256
		hash: keccak256(channelId, paymentId, withdrawalAmount)
	*/

	event ChannelOpened(bytes32 indexed _channelId);
	event ChannelClosed(bytes32 indexed _channelId);
	event ChannelExpired(bytes32 indexed _channelId);
	event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue);

	modifier bothProofsSubmitted(bytes32 _channelId) {
		require(channels[_channelId].sourceProofSubmitted);
		require(channels[_channelId].destinationProofSubmitted);
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
		bytes32 channelId = keccak256(msg.sender, _destination, _channelValueInWei, currentDate);
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
		channels[_channelId].destinationProofSubmitted = true;
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
		uint256 _paymentId,
		uint256 _withdrawalAmount)
		public
		returns (bool)
	{
		// validate channel id
		require(channelIds[_channelId]);
		// validate channel state
		require(channels[_channelId].state == ChannelStates.opened);
		// make sure nobody else is trying to withdraw the funds
		require(msg.sender == channels[_channelId].destination);
		// prevent a micropayment from reducing the entire balance
		require(channels[_channelId].value > _withdrawalAmount && _withdrawalAmount > 0);
		// following two lines construct the proof, with prefix to validate _h
		bytes32  _proof = keccak256(_channelId, _paymentId, _withdrawalAmount);
		bytes32 proof = keccak256(prefix, _proof);
		// validate the proof, if it fails most likely malicious submitter, so lets waste their gas ;)
		assert(proof == _h);
		// make sure the proof hasn't already bee submitted
		require(!microPaymentHashes[_channelId][_h]);
		// mark proof as submitted
		microPaymentHashes[_channelId][_h] = true;
		// time to recover the signature
		address signer = ecrecover(_h, _v, _r, _s);
		// make sure its from the source, otherwise someones up to no good so lets waste their gas ;)
		assert(signer == channels[_channelId].source);
		// calculate remaining channel value
		uint256 remainingChannelValue = channels[_channelId].value.sub(_withdrawalAmount);
		// adjust channel value
		channels[_channelId].value = remainingChannelValue;
		// notify blockchain
		MicroPaymentWithdrawn(_channelId, _withdrawalAmount, remainingChannelValue);
		// withdraw funds
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