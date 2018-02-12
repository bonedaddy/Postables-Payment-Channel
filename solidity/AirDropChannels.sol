pragma solidity 0.4.19;
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Interfaces/ERC20Interface.sol";

contract AirdropChannels is Administration {
	
	using SafeMath for uint256;

	string constant public VERSION = "0.0.5alpha";

	uint256 constant public AIRDROPAMOUNT = 300 * 1 ether;
	bytes private prefix = "\x19Ethereum Signed Message:\n32";
	// prevent any possible accidental triggering of developer only conditions
	bool	public	dev = true;

	ERC20Interface public ercI = ERC20Interface(0x0994F9595D28429584BfB5FCBFEA75b9c9Ea2c24);

	enum ChannelStates { opened, releasing, closed }

	struct ChannelStruct {
		address source;
		address tokenAddress;
		uint256 value;
		uint256 closingDate;
		uint256 openDate;
		uint256 totalDrops;
		bytes32 channelId;
		ChannelStates state;
	}

	mapping (bytes32 => bool) private channelIds;
	mapping (bytes32 => ChannelStruct) public channels;
	// k1 = channel id
	// k2 = address
	mapping (bytes32 => mapping (address => bool)) private receivedBonus;

	event ChannelOpened(bytes32 indexed _channelId);
	event ChannelClosed(bytes32 indexed _channelId);
	event AirDropsEnabled(bytes32 indexed _channelId);
	event AirDropDispersed(bytes32 indexed _channelId);
	event SignatureRecovered(address indexed signer);
	event SigDebug(bytes32 _h, bytes32 _proof, bytes32 proof, address signer);

	function () public payable {}

	function openChannel(
		address _tokenAddress,
		uint256 _channelValue,
		uint256 _durationInDays)
		public
		payable
		returns (bool)
	{
		uint256 currentDate = now;
		// channel hash = keccak256(purchaser, vendor, channel value, date of open)
		bytes32 channelId = keccak256(msg.sender, _tokenAddress, currentDate);
		// make sure the channel id doens't already exist
		require(!channelIds[channelId]);
		channels[channelId].source = msg.sender;
		channels[channelId].tokenAddress = _tokenAddress;
		channels[channelId].value = _channelValue;
		channels[channelId].closingDate = (now + (_durationInDays * 1 days));
		channels[channelId].openDate = currentDate;
		channels[channelId].channelId = channelId;
		channels[channelId].state = ChannelStates.opened;
		channelIds[channelId] = true;
		ChannelOpened(channelId);
		require(ercI.transferFrom(msg.sender, address(this), _channelValue));
		return true;
	}

	// airdrop enable proof = keccak256(prefix, keccak256(channelId, airdropID))
	// used to enable the air drop channel
	function enableAirdrops(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		bytes32 _channelId,
		uint256 _id)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(channels[_channelId].state == ChannelStates.opened);
		require(msg.sender == channels[_channelId].source);
		// we need to recreate the signed message hash, so first lets compute the raw hash using the preimages
		bytes32 _proof = keccak256(_channelId, _id);
		// now lets add the prefix, to get the signed message hash, or pefixed message hash
		bytes32 proof = keccak256(prefix, _proof);
		// retrieve the signer of the message
		address signer = ecrecover(_h, _v, _r, _s);
		// if someone fails this chances are it was malicious, so lets waste their gas
		assert(signer == channels[_channelId].source);
		// however we need to make sure they also submitted the right message
		// so we compare our dynamically generated proof, with the passed in signed message hash
		assert(proof == _h);
		// mark channel as releasing
		channels[_channelId].state = ChannelStates.releasing;
		if (dev) { SigDebug(_h, _proof, proof, signer); }
		AirDropsEnabled(_channelId);
		SignatureRecovered(signer);
		return true;
	}

	// air drop proof: keccak256(prefix, keccak256(_channelId, _id, msg.sender))
	function retrieveAirdrop(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		bytes32 _channelId,
		uint256 _id)
		public
		returns (bool)
	{
		// verify channel id
		require(channelIds[_channelId]);
		// verify we are in the correct state
		require(channels[_channelId].state == ChannelStates.releasing);
		// verify channel balance
		require(channels[_channelId].value >= AIRDROPAMOUNT);
		require(!receivedBonus[_channelId][msg.sender]);
		// this ensure only the intended recipient of a signed message can redeem
		bytes32 _proof = keccak256(_channelId, _id, msg.sender);
		bytes32 proof = keccak256(prefix, _proof);
		address signer = ecrecover(_h, _v, _r, _s);
		// validate the signer
		assert(signer == channels[_channelId].source);
		// signer checks out, but does the data?
		/**
			68 74 74 70 73 3a 2f 2f 70 69 63 73 2e 6d
			65 2e 6d 65 2f 69 74 73 2d 61 6e 2d 6f 6c
			64 65 72 2d 6d 65 6d 65 2d 73 69 72 2d 62
			75 74 2d 69 74 2d 63 68 65 63 6b 73 2d 6f
			75 74 2d 31 39 38 39 39 39 37 34 2e 70 6e
			67 
		*/
		assert(proof == _h);
		// mark them as having received a bonus so they can't double dip.
		receivedBonus[_channelId][msg.sender] = true;
		// reduce the channel value
		channels[_channelId].value = channels[_channelId].value.sub(AIRDROPAMOUNT);
		// increase number of air drops
		channels[_channelId].totalDrops = channels[_channelId].totalDrops.add(1);
		// notify blockchain
		AirDropDispersed(_channelId);
		// transfer tokens
		require(ercI.transfer(msg.sender, AIRDROPAMOUNT));
		return true;
	}

	function closeChannel(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		bytes32 _channelId)
		public
		returns (bool)
	{
		require(channelIds[_channelId]);
		require(channels[_channelId].state == ChannelStates.opened ||
			    channels[_channelId].state == ChannelStates.releasing);
		// if we aren't in dev mode, make sure the closing date has passed 
		if (!dev) {
			require(now > channels[_channelId].closingDate);
		}
		// make sure msg.sender is channel owner
		require(msg.sender == channels[_channelId].source);
		address signer = ecrecover(_h, _v, _r, _s);
		assert(signer == channels[_channelId].source);
		channels[_channelId].state = ChannelStates.closed;
		channels[_channelId].closingDate = now;
		// check to see if channe lvalue is greater than 0, if so withdraw remaining funds
		if (channels[_channelId].value > 0 ) {
			uint256 deposit = channels[_channelId].value;
			channels[_channelId].value = 0;
			require(ercI.transfer(msg.sender, deposit));
		}
		ChannelClosed(_channelId);
		return true;
	}

	function verifyProof(
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s)
		public
		pure 
		returns (address)
	{
		return ecrecover(_h, _v, _r, _s);
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

	function withdrawTokens()
		public
		onlyAdmin
		returns (bool)
	{
		require(dev);
		require(ercI.transfer(msg.sender, ercI.balanceOf(address(this))));
		return true;
	}

}