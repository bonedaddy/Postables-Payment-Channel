pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Libs/EcRecovery.sol";

contract PaymentChannels is Administration {
	
	using SafeMath for uint256;
	using EcRecovery as ec;

	
	uint256 private channelCount;

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
	mapping (bytes32 => ChannelStruct) private channels;
	mapping (bytes32 => bool) private channelIds;

	function openChannel(
		address _vendor,
		uint256 _value)
		public
		returns (bool)
	{
		bytes32 channelId = keccak256(msg.sender, _vendor, _value);
		require(!channelIds[channelId]);
		require(!channels[channelId].opened);
		channelIds[channelId] = true;
		channels[channelId].purchaser = msg.sender;
		channels[channelId].vendor = _vendor;
		channels[channelId].channelId = channelId;
		channels[channelId].opened = true;
		channels[channelId].autoClosureDate = now + 10 days;
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
		require(channels[channelId].vendor == msg.sender)
		require(channels[channelId].closed == false && channels[channelId].timedOut == false);
		require(_purchaser == channels[channelId].purchaser);
		address signer = ecrecover(_h, _v, _r, _s);
		require(signer == _purchaser);
		channels[channelId].proofSubmitted[signer] = true;
		return true;
	}

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
		return true;
	}

	function vendorWithdrawFunds(
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(!channels[channelId].closed);
		address purchaser = channels[channelId].purchaser;
		require(msg.sender == channels[channelId].vendor);
		require(channels[channelId].proofSubmitted[purchaser]);
		require(channels[channelId].proofSubmitted[msg.sender]);
		channels[channelId].closed = true;
		// now that we have confirmed both proofs have been submitted vendor can withdraw funds
		require(msg.sender.transfer(channels[channelId].value));
		return true;
	}

	/**
		@dev Callable by channel opener if now proofs have been submitted wwhen the closure date is past
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
		require(now > channels[channelId].autoClosureDate);
		channels[channelId].closed = true;
		channels[channelId].timedOut = true;
		return true;
	}


	/**
		@dev Lets the channel opener withdraw their funds if the channel timesout
	*/
	function openerWithdrawFunds(
		address _vendor,
		uint256 _value)
		public
		returns (bool)
	{
		bytes32 channelId = keccak256(msg.sender, _vendor, _value);
		require(channels[channelId].closed == true && channels[channelId].timedOut == true);
		msg.sender.transfer(_value);
		return true;
	}

}