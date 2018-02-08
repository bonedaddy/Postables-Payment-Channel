pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Libs/EcRecovery.sol";

/*
	Will be used as the basis for V3 which will enable trustless micro payments, ontop of full channel payments
*/
contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	string constant public VERSION = "0.0.3alpha";
	
	uint256 private channelCount;
	// prevent any possible accidental triggering of developer only conditions
	bool	public	dev = true;

	enum ChannelStates { opened, expired, closed }

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

	mapping (bytes32 => ChannelStruct) public channels;
	mapping (uint256 => bytes32) private channelNumber;
	mapping (bytes32 => bool) private channelIds;
	// prevent resubmission of the same signed messages by a particular address within a channel
	mapping (bytes32 => mapping (bytes32 => mapping(address => bool))) private signedMessages;

	event ChannelOpened(bytes32 indexed _channelId);
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
		// make sure the channel id doens't already exist
		require(!channelIds[channelId]);
		channelIds[channelId] = true;
		channels[channelId].purchaser = msg.sender;
		channels[channelId].vendor = _vendor;
		channels[channelId].value = _channelValueInWei;
		channels[channelId].closingDate = (now + (_durationInDays * 1 days));
		channels[channelId].openDate = currentDate;
		channels[channelId].channelId = channelId;
		channels[channelId].state = defaultState;
		ChannelOpened(channelId);
		return true;
	}

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

	function closeChannel(
		bytes32 _channelId)
		public
		bothProofsSubmitted(_channelId)
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(channels[_channelId].state == ChannelStates.opened);
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
		So long as both proofs have *NOT* been submitted, and it is past closing, terminate channel
	*/
	function expireChannel(
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(msg.sender == channels[_channelId].purchaser);
		// require that the channel is open
		require(channels[_channelId].state == ChannelStates.opened);
		// if both proofs have been submitted, throw
		if (channels[_channelId].purchaserProofSubmitted && channels[_channelId].vendorProofSubmitted) {
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