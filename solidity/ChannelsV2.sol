pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Libs/EcRecovery.sol";
/**
	Add message hash verificaton when submitting a second signature.
	first signature to submit commits the message hash, subsequent signatures mush also submit a signature
	which verifies off that message hash. This should help reduce exploitation attempts

	Implementing enums for channel states

	Implement option for two step channel opening or combine both steps:
		> open channel with vendor message
		> open channel without vendor message, manually submit message

*/
contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	
	uint256 private channelCount;
	bool	public	dev = true;

	/** State Description
		Initial channel state is "proposed"
		It becomes "accepted" once the vendor proof is submitted
		It becomes "expired" if channel opener times out channel due to lack of vendor proof
		It becomes "closed" if the channel is succesfully closed peacefully

	*/
	enum ChannelStates { proposed, accepted, expired, closed }

	ChannelStates public defaultState = ChannelStates.proposed;

	struct ChannelStruct {
		address purchaser;
		address vendor;
		uint256 value;
		uint256 expirationDate;
		uint256 openDate;
		bytes32 channelId;
		ChannelStates state;
		mapping (address => bool) proofSubmitted;
	}

	mapping (uint256 => bytes32) private channelNumber;
	mapping (bytes32 => ChannelStruct) public channels;
	mapping (bytes32 => bool) private channelIds;

	event ChannelOpened(address indexed _purchaser, address indexed _vendor, bytes32 indexed _channelId);
	event ChannelClosed(bytes32 indexed _channelId);
	event VendorProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);
	event PurchaserProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress);

	function () public payable {}

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
		uint256 autoClosureDate = (now + ((_durationInDays * 1 days) * 2));
		channels[channelId].purchaser = msg.sender;
		channels[channelId].vendor = _vendor;
		channels[channelId].value = _channelValueInWei;
		channels[channelId].expirationDate = autoClosureDate;
		channels[channelId].openDate = currentDate;
		channels[channelId].channelId = channelId;
		channels[channelId].state = defaultState;
		ChannelOpened(msg.sender, _vendor, channelId);
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

}