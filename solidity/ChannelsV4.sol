pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Interfaces/ERC20Interface.sol";

contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	string constant public VERSION = "0.4.1alpha";
	
	bytes private prefix = "\x19Ethereum Signed Message:\n32";
	// prevent any possible accidental triggering of developer only conditions
	bool	public	dev = true;

	enum ChannelStates { opened, expired, finalized, closed }

	ChannelStates public defaultState = ChannelStates.opened;

	struct EthChannelStruct {
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

	struct ErcChannelStruct {
		address source;
		address destination;
		address tokenAddress;
		uint256 value;
		uint256 closingDate;
		uint256 openDate;
		bytes32 channelId;
		bool 	sourceProofSubmitted;
		bool 	destinationProofSubmitted;
		ChannelStates state;
	}

	mapping (bytes32 => EthChannelStruct) public ethChannels;
	mapping (bytes32 => ErcChannelStruct) public ercChannels;
	mapping (uint256 => bytes32) private channelNumber;
	mapping (bytes32 => bool) private channelIds;
	/** prevent resubmission of the same signed messages by a particular address within a channel
		key 1 (bytes32) = channel id
		key 2 (bytes32) = messageHash
		key 3 (address) = address of the submitter
		val   (bool)    = whether or not the message has already been submitted
	*/
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

	event EthChannelOpened(bytes32 indexed _channelId);
	event ErcChannelOpened(bytes32 indexed _channelId);
	event ChannelClosed(bytes32 indexed _channelId);
	event ChannelExpired(bytes32 indexed _channelId);
	event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue);

	function () public payable {}

	function openEthChannel(
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
		ethChannels[channelId].source = msg.sender;
		ethChannels[channelId].destination = _destination;
		ethChannels[channelId].value = _channelValueInWei;
		ethChannels[channelId].closingDate = (now + (_durationInDays * 1 days));
		ethChannels[channelId].openDate = currentDate;
		ethChannels[channelId].channelId = channelId;
		ethChannels[channelId].state = defaultState;
		EthChannelOpened(channelId);
		return true;
	}

	function openErcChannel(
		address _tokenAddress,
		address _destination,
		uint256 _channelValueInWei,
		uint256 _durationInDays)
		public
		returns (bool)
	{
		bytes32 channelId = keccak256(msg.sender, _destination, _tokenAddress, now);
		// make sure the chanel ID doesn't already exist
		require(!channelIds[channelId]);
		channelIds[channelId] = true;
		ercChannels[channelId].source = msg.sender;
		ercChannels[channelId].destination = _destination;
		ercChannels[channelId].tokenAddress = _tokenAddress;
		ercChannels[channelId].value = _channelValueInWei;
		ercChannels[channelId].closingDate = (now + (_durationInDays * 1 days));
		ercChannels[channelId].openDate = now;
		ercChannels[channelId].channelId = channelId;
		ERC20Interface e = ERC20Interface(_tokenAddress);
		ErcChannelOpened(channelId);
		require(e.transferFrom(msg.sender, address(this), _channelValueInWei));
		return true;
	}	

	function submitErcSourceProof(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		bytes32 _channelId)
		public
		returns (bool)
	{
		// verify correct channel id
		require(channelIds[_channelId]);
		// ensure channel state is opened
		require(ercChannels[_channelId].state == ChannelStates.opened);
		// make sure source proof hasn't already been submitted
		require(!ercChannels[_channelId].sourceProofSubmitted);
		// recover address of signer
		address signer = ecrecover(_h, _v, _r, _s);
		// ensure that the signer is the source of the channel
		assert (signer == ercChannels[_channelId].source);
		// mark source proof as submitted
		ercChannels[_channelId].sourceProofSubmitted = true;
		// ensure this proof can't be reused by the same sender
		signedMessages[_channelId][_h][msg.sender] = true;
		// notify blockchain
		SourceProofSubmitted(_channelId, signer);
		return true;
	}

	function submitErcDestinationProof(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s, 
		bytes32 _channelId)
		public
		returns (bool)
	{
		// verify correct channel id
		require(channelIds[_channelId]);
		// ensure channel state is opened
		require(ercChannels[_channelId].state == ChannelStates.opened);
		// make sure destination proof hasn't already been submitted
		require(!ercChannels[_channelId].destinationProofSubmitted);
		// recover address of signer
		address signer = ecrecover(_h, _v, _r, _s);
		// ensure that the signer is the destionation of the channel
		assert (signer == ercChannels[_channelId].destination);
		// mark destination proof as submitted
		ercChannels[_channelId].destinationProofSubmitted = true;
		// ensure this proof can't be reused by the same sender
		signedMessages[_channelId][_h][msg.sender] = true;
		// notify blockchain
		SourceProofSubmitted(_channelId, signer);
		return true;
	}

	function verifyDoubleProof(
		bytes32 _channelId,
		bool    _ercChannel)
		internal
		returns (bool)
	{
		if (_ercChannel) {
			require(ercChannels[_channelId].state == ChannelStates.opened);
			if (ercChannels[_channelId].sourceProofSubmitted == true && ercChannels[_channelId].destinationProofSubmitted) {
				// both proofs have been submited, lets mark as finalized
				ercChannels[_channelId].state = ChannelStates.finalized;
			}
		} else {
			require(ethChannels[_channelId].state == ChannelStates.opened);
			if (ethChannels[_channelId].sourceProofSubmitted == true && ethChannels[_channelId].destinationProofSubmitted) {
				// both proofs have been submitted, lets mark as finalized
				ethChannels[_channelId].state = ChannelStates.finalized;
			}		
		}
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

	function withdrawTokens(
		address _tokenAddress)
		public
		onlyAdmin
		returns (bool)
	{
		require(dev);
		ERC20Interface e = ERC20Interface(_tokenAddress);
		uint256 balance = e.balanceOf(address(this));
		require(e.transfer(msg.sender, balance));
		return true;
	}

}