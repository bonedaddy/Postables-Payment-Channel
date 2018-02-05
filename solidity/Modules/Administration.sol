pragma solidity 0.4.19;

contract Administration {

    address public owner;
    address public admin;
    bool    public frozen;

    mapping (address => bool) public moderators;

    event AdminSet(address indexed _admin, bool indexed _adminSet);
    event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner, bool indexed _ownershipTransferred);

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == owner || msg.sender == admin);
        _;
    }

    modifier onlyPrivileged() {
        require(msg.sender == owner || msg.sender == admin || moderators[msg.sender] == true);
        _;
    }

    function Administration() {
        owner = msg.sender;
        admin = msg.sender;
    }

    function setAdmin(
        address _newAdmin
    )
        public
        onlyOwner
        returns (bool)
    {
        require(_newAdmin != admin);
        admin = _newAdmin;
        AdminSet(_newAdmin, true);
    }

    function transferOwnership(
        address _newOwner
    )
        public
        onlyOwner
        returns (bool)
    {
        require(_newOwner != owner);
        owner = _newOwner;
        OwnershipTransferred(msg.sender, _newOwner, true);
    }
}
