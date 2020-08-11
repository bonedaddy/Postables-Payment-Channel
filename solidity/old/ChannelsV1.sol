pragma solidity 0.4.20;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Libs/EcRecovery.sol";
/**
	Add message hash verificaton when submitting a second signature.
	first signature to submit commits the message hash, subsequent signatures mush also submit a signature
	which verifies off that message hash. This should help reduce exploitation attempts
*/
contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	
	uint256 private channelCount;
	bool	public	dev = true;

	struct ChannelStruct {
		address purchaser;
		address vendor;
		uint256 value;
		uint256 autoClosureDate;
		bytes32 channelId;
		bool	opened;
		bool	closed;
		bool	timedOut;
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
		uint256 _value)
		public
		payable
		returns (bool)
	{
		bytes32 channelId = keccak256(msg.sender, _vendor, _value);
		require(!channelIds[channelId]);
		require(!channels[channelId].opened);
		channelIds[channelId] = true;
		channels[channelId].purchaser = msg.sender;
		channels[channelId].vendor = _vendor;
		channels[channelId].value = _value;
		channels[channelId].channelId = channelId;
		channels[channelId].opened = true;
		channels[channelId].autoClosureDate = block.timestamp + 10 days;
		ChannelOpened(msg.sender, _vendor, channelId);
		return true;
	}


	/**
		Working
	*/
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
		PurchaserProofSubmitted(channelId, signer);
		return true;
	}

	/**
		For some reason, not with the purchaser proof but ONLY with the vendor proof (in testing at least)
		is not returning the correct value for `s`. In order to do so, we must convert `s` to bytes, then to hex.
		See readme for an explanation
	*/
	function submitVendorProof(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		address _vendor,
		uint256 _channelValue)
		public
		returns (bool)
	{
		bytes32 channelId = keccak256(msg.sender, _vendor, _channelValue);
		require(channels[channelId].purchaser == msg.sender);
		require(channels[channelId].closed == false && channels[channelId].timedOut == false);
		require(_vendor == channels[channelId].vendor);
		address signer = ecrecover(_h, _v, _r, _s);
		require(signer == _vendor);
		channels[channelId].proofSubmitted[signer] = true;
		VendorProofSubmitted(channelId, signer);
		return true;
	}

	function vendorWithdraw(
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(channels[_channelId].vendor == msg.sender);
		require(channels[_channelId].closed == false);
		require(channels[_channelId].timedOut == false);
		require(channels[_channelId].proofSubmitted[msg.sender] == true);
		address purchaser = channels[_channelId].purchaser;
		require(channels[_channelId].proofSubmitted[purchaser] = true);
		channels[_channelId].closed =true;
		ChannelClosed(_channelId);
		msg.sender.transfer(channels[_channelId].value);
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