pragma solidity 0.4.19;

/*
    ERC20 Standard Token interface
*/
interface ERC20Interface {
    function owner() public view returns (address);
    function decimals() public view returns (uint8);
    function transfer(address _to, uint256 _value) public returns (bool);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool);
    function approve(address _spender, uint256 _amount) public returns (bool);
    function totalSupply() public view returns (uint256);
    function balanceOf(address _owner) public view returns (uint256);
    function allowance(address _owner, address _spender) public view returns (uint256);
}

