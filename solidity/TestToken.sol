pragma solidity 0.4.19;

import "./Math/SafeMath.sol";
import "./Modules/Administration.sol";

contract ERC20 {

    using SafeMath for uint256;

    string  public name;
    string  public symbol;
    uint256 public totalSupply;
    uint8   public decimals;

    mapping (address => uint256) public balances;
    mapping (address => mapping (address => uint256)) public allowed;

    event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount);
    event Approve(address indexed _owner, address indexed _spender, uint256 _amount);

    function ERC20() {
        name = "SafeOrDapp";
        symbol = "SOD";
        decimals = 18;
        totalSupply = 2000000000000000000000000000;
        balances[msg.sender] = totalSupply;
    }

    function transfer(
        address _recipient,
        uint256 _amount
    )
        public
        returns (bool transferred)
    {
        require(transferCheck(msg.sender, _recipient, _amount));
        balances[msg.sender] = balances[msg.sender].sub(_amount);
        balances[_recipient] = balances[_recipient].add(_amount);
        Transfer(msg.sender, _recipient, _amount);
        return true;
    }

    function transferFrom(
        address _owner,
        address _recipient,
        uint256 _amount
    )
        public
        returns (bool transferredFrom)
    {
        require(transferCheck(_owner, _recipient, _amount));
        require(allowed[_owner][msg.sender] >= _amount);
        require(allowed[_owner][msg.sender].sub(_amount) >= 0);
        allowed[_owner][msg.sender] = allowed[_owner][msg.sender].sub(_amount);
        balances[_owner] = balances[_owner].sub(_amount);
        balances[_recipient] = balances[_recipient].add(_amount);
        Transfer(_owner, _recipient, _amount);
        return true;
    }

    function approve(
        address _spender,
        uint256 _amount
    )
        public
        returns (bool approved)
    {
        require(_spender != address(0x0));
        require(_amount > 0);
        require(allowed[msg.sender][_spender].add(_amount) > allowed[msg.sender][_spender]);
        allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_amount);
        return true;
    }

    /**INTERNALS */

    function transferCheck(
        address _sender,
        address _recipient,
        uint256 _amount
    )
        internal
        view
        returns (bool valid)
    {
        require(_sender != address(0x0) && _recipient != address(0x0) && _amount > 0);
        require(balances[_sender].sub(_amount) >= 0);
        require(balances[_recipient].add(_amount) > 0);
        require(balances[_recipient].add(_amount) > balances[_recipient]);
        return true;
    }

    /**GETTERS */

    function totalSupply()
        public
        view
        returns (uint256)
    {
        return totalSupply;
    }

    function balanceOf(
        address _holder
    )
        public
        view
        returns (uint256)
    {
        return balances[_holder];
    }

    function allowance(
        address _owner,
        address _spender
    )
        public
        view
        returns (uint256)
    {
        return allowed[_owner][_spender];
    }
}