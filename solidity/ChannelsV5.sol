pragma solidity 0.4.20;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Interfaces/ERC20Interface.sol";

contract PaymentChannels is Administration {
	
	using SafeMath for uint256;

	string constant public VERSION = "0.4.3alpha";
	
	bytes private prefix = "\x19Ethereum Signed Message:\n32";
	// prevent any possible accidental triggering of developer only conditions
	bool	public	dev = true;

	uint256 private numHosts;

	enum ChannelStates { opened, expired, finalized, closed }

	ChannelStates public defaultState = ChannelStates.opened;

	struct ChannelStruct { // a -- > b
		address source;
		address destination;
		uint256 value;
		bytes32 channelId;
	}


	struct HostStruct {
		uint256 id;
		uint256 peer1;
		uint256 peer2;
		address hostAddress;
	}

	mapping (uint256 => HostStruct) private hosts;
	mapping (address => bool)		private registeredHosts;

	function openChannel
}