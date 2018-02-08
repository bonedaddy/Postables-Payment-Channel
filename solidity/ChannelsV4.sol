pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Interfaces/ERC20Interface.sol";

contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	string constant public VERSION = "0.4.1alpha";
	
	uint256 private channelCount;
	bytes private prefix = "\x19Ethereum Signed Message:\n32";
	// prevent any possible accidental triggering of developer only conditions
	bool	public	dev = true;

	enum ChannelStates { opened, expired, closed }

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
		ercChannels[channelId].closingDate = (now + (_durationInDays * 1 days));
		ercChannels[channelId].openDate = now;
		ercChannels[channelId].channelId = channelId;
		ERC20Interface e = ERC20Interface(_tokenAddress);
		ErcChannelOpened(channelId);
		require(e.balanceOf(msg.sender) >= _channelValueInWei);
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