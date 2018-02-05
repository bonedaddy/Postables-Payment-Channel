pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Libs/EcRecovery.sol";
/**

	To Dos:
		> Implement enums for channel states

	What is a vendor proof?
		> A vendor proof is a message signed by a vendor, whose signature data is submitted by the 
			channel opener to verify the vendor
	What is a purchaser proof?
		> A purchaser proof is a message signed by the purchaser (channel opener), whose signature
			is submitted by the vendor in a channel
*/

contract PaymentChannels is Administration {
	
	using SafeMath for uint256;


	uint256 public channelFee;
	uint256 private channelCount;

	enum ChannelStates { Proposed, Accepted, Opened, Closed, Disconnected }
	ChannelStates public defaultState = ChannelStates.Proposed;

	struct ChannelStruct {
		address purchaser;
		address vendor;
		uint256 channelValue;
		uint256 autoClosureDate;
		bytes32 channelId;
		bool	opened;
		bool	closed;
		bool	timedOut;
		mapping (address => bool) proofSubmitted;
		ChannelSates state;
	}

	mapping (uint256 => bytes32) private channelNumber;
	mapping (bytes32 => bool) private channelIds;
	mapping (address => mapping (bytes32 => uint256)) private deposits;
	mapping (bytes32 => ChannelStruct) public channels;

	event ChannelOpened(bytes32 indexed _channelId);
	event ChannelClosed(bytes32 indexed _channelId);
	event ChannelTimedOut(bytes32 indexed _channeId);
	event PurchaserProofSubmitted(bytes32 indexed _channelId));
	event VendorProofSubmitted(bytes32 indexd _channelId);

	function () payable {}


	/**
		* Channel opener must also submit the vendor proof. This is so that a malicious channel opener
			is unable to withold the proof to prevent a vendor from being paid
		* Subsequently, this also allows the channel opener to with-hold submitting their signature to the
			vendor should the vendor prove to be malicious, they will be entitled to close their timeout their
			channel, pending a wait period 
	*/
	function openChannel(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		address _vendor,
		uint256 _channelValue,
		uint256 _durationInDays)
		public
		payable
		returns (bool)
	{
		bytes32 channelId = keccak256(msg.sender, _vendor, _channelValue);
		require(!channelIds[channelId]);
		require(msg.value > 0 && msg.value == _channelValue);
		channelIds[channelId] = true;
		channels[channelId].purchaser = msg.sender;
		channels[channelId].vendor = _vendor;
		channels[channelId].channelValue = _channelValue;
		channels[channelId].autoClosureDate = now + (_durationInDays * 1 days;
		channels[channelId].channelId = channelId;
		channels[channelId].opened = true;
		channels[channelId].state = defaultState;
		deposits[msg.sender][channelId] = msg.value;
		ChannelOpened(channelId);
		require(submitVendorProof(_h, _v, _r, _s, _vendor, _value));
		return true;
	}

	function submitPurchaserProof(
		bytes32 _h, // message hash signed by purchaser
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		address _purchaser, // address of the purchaser
		uint256 _channelValue) // value of the channel, used to generate channel ID
		public
		returns (bool)
	{
		bytes32 channelId = keccak256(_purchaser, msg.sender, _channelValue);
		require(channels[channelId].vendor == msg.sender);
		require(channels[channelId].closed == false && channels[channelId].timedOut == false);
		require(_purchaser == channels[channelId].purchaser);
		address signer = ecrecover(_h, _v, _r, _s);
		require(signer == _purchaser);
		channels[channelId].proofSubmitted[signer] = true;
		PurchaserProofSubmitted(channelId);
		return true;
	}

	/**
		@dev Must be supplied by channel opener when the channel is open
	*/
	function submitVendorProof(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		address _vendor,
		uint256 _channelValue)
		internal
		returns (bool)
	{
		bytes32 channelId = keccak256(msg.sender, _vendor, _channelValue);
		require(channels[channelId].purchaser == msg.sender);
		require(channels[channelId].closed == false && channels[channelId].timedOut == false);
		require(_vendor == channels[channelId].vendor);
		address signer = ecrecover(_h, _v, _r, _s);
		require(signer == _vendor);
		channels[channelId].proofSubmitted[signer] = true; 
		VendorProofSubmitted(channelId);
		return true;
	}

	/**
		* Used by the vendor "recipient" to withdraw funds from the channel so long as both proofs are valid
		* Also closes the channel (one of two ways)
	*/
	function vendorWithdrawFunds(
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(!channels[_channelId].closed);
		address purchaser = channels[_channelId].purchaser;
		require(msg.sender == channels[_channelId].vendor);
		require(channels[_channelId].proofSubmitted[purchaser]);
		require(channels[_channelId].proofSubmitted[msg.sender]);
		channels[_channelId].closed = true;
		ChannelClosed(_channelId);
		// now that we have confirmed both proofs have been submitted vendor can withdraw funds
		msg.sender.transfer(channels[_channelId].value);
		return true;
	}

	/**
		* Callable by channel opener if now proofs have been submitted wwhen the closure date is past
		* Also closes the channel (one of two ways to)
		* TO DO: Implement wait period
	*/
	function timeoutChannel(
		address _vendor,
		uint256 _value)
		public
		returns (bool)
	{
		bytes32 channelId = keccak256(msg.sender, _vendor, _value);
		require(channels[channelId].purchaser == msg.sender);
		require(!channels[channelId].closed);
		require(!channels[channelId].timedOut);
		require(deposits[msg.sender][channelId] > 0);
		require(now > channels[channelId].autoClosureDate);
		uint256 deposit = deposits[msg.sender][channelId];
		deposits[msg.sender][channelId] = 0;
		channels[channelId].closed = true;
		channels[channelId].timedOut = true;
		ChannelTimedOut(channelId);
		msg.semder.transfer(deposit);
		return true;
	}

}