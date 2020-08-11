// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package airdrop

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AirdropABI is the input ABI used to generate the binding from.
const AirdropABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_adminSet\",\"type\":\"bool\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"AirDropDispersed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"AirDropsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_ownershipTransferred\",\"type\":\"bool\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_proof\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prefixedProof\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SigDebug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignatureRecovered\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDrops\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"contractERC20Interface\",\"name\":\"intf\",\"type\":\"address\"},{\"internalType\":\"enumAirdropChannels.ChannelStates\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"closeChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"enableAirdrops\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"moderators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_channelValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_durationInDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dropAmount\",\"type\":\"uint256\"}],\"name\":\"openChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"retrieveAirdrop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// AirdropBin is the compiled bytecode used for deploying new contracts.
var AirdropBin = "0x60c0604052601c60808190527f19457468657265756d205369676e6564204d6573736167653a0a33320000000060a090815261003e9160039190610080565b506004805460ff1916600117905534801561005857600080fd5b5060008054336001600160a01b03199182168117909255600180549091169091179055610113565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c157805160ff19168380011785556100ee565b828001600101855582156100ee579182015b828111156100ee5782518255916020019190600101906100d3565b506100fa9291506100fe565b5090565b5b808211156100fa57600081556001016100ff565b6111f580620001236000396000f3fe6080604052600436106100ec5760003560e01c80638da5cb5b1161008a578063f2fde38b11610059578063f2fde38b1461039e578063f36bdb49146103d1578063f851a44014610409578063ffa1ad741461041e576100f3565b80638da5cb5b1461030457806391456d3d1461033557806391cca3db14610374578063a0ef91df14610389576100f3565b806356715148116100c6578063567151481461019c5780635c9bd273146101e7578063704b6c021461022c5780637a7ebd7b1461025f576100f3565b8063054f7d9c146100f557806314d0f1ba1461011e578063515982f614610151576100f3565b366100f357005b005b34801561010157600080fd5b5061010a6104a8565b604080519115158252519081900360200190f35b34801561012a57600080fd5b5061010a6004803603602081101561014157600080fd5b50356001600160a01b03166104b8565b34801561015d57600080fd5b5061010a600480360360c081101561017457600080fd5b5080359060ff6020820135169060408101359060608101359060808101359060a001356104cd565b3480156101a857600080fd5b5061010a600480360360c08110156101bf57600080fd5b5080359060ff6020820135169060408101359060608101359060808101359060a00135610822565b3480156101f357600080fd5b5061010a600480360360a081101561020a57600080fd5b5080359060ff6020820135169060408101359060608101359060800135610ad0565b34801561023857600080fd5b5061010a6004803603602081101561024f57600080fd5b50356001600160a01b0316610d37565b34801561026b57600080fd5b506102896004803603602081101561028257600080fd5b5035610db5565b604051808b6001600160a01b031681526020018a6001600160a01b03168152602001898152602001888152602001878152602001868152602001858152602001848152602001836001600160a01b031681526020018260028111156102ea57fe5b81526020019a505050505050505050505060405180910390f35b34801561031057600080fd5b50610319610e1b565b604080516001600160a01b039092168252519081900360200190f35b34801561034157600080fd5b506103196004803603608081101561035857600080fd5b5080359060ff6020820135169060408101359060600135610e2a565b34801561038057600080fd5b5061010a610e9a565b34801561039557600080fd5b5061010a610ea3565b3480156103aa57600080fd5b5061010a600480360360208110156103c157600080fd5b50356001600160a01b0316610f13565b61010a600480360360808110156103e757600080fd5b506001600160a01b038135169060208101359060408101359060600135610f94565b34801561041557600080fd5b5061031961115c565b34801561042a57600080fd5b5061043361116b565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561046d578181015183820152602001610455565b50505050905090810190601f16801561049a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600154600160a01b900460ff1681565b60026020526000908152604090205460ff1681565b60008281526005602052604081205460ff166104e857600080fd5b6001600084815260066020526040902060080154600160a01b900460ff16600281111561051157fe5b1461051b57600080fd5b60008381526006602052604090206005810154600290910154101561053f57600080fd5b600083815260076020908152604080832033845290915290205460ff161561056657600080fd5b600083833360405160200180848152602001838152602001826001600160a01b031660601b81526014019350505050604051602081830303815290604052805190602001209050600060038260405160200180838054600181600116156101000203166002900480156106105780601f106105ee576101008083540402835291820191610610565b820191906000526020600020905b8154815290600101906020018083116105fc575b505082815260200192505050604051602081830303815290604052805190602001209050600060018a8a8a8a60405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610690573d6000803e3d6000fd5b505060408051601f1901516000898152600660205291909120549092506001600160a01b0380841691161490506106c357fe5b8982146106cc57fe5b60008681526007602090815260408083203384528252808320805460ff1916600117905588835260069091529020600581015460029091015461070e91611191565b6000878152600660208190526040909120600281019290925501546107349060016111a6565b600087815260066020819052604080832090910192909255905187917fa467a409776bfc50e32dd8496edd970b5e7a84a38acda61529d8022bf757c2fd91a260008681526006602090815260408083206008810154600590910154825163a9059cbb60e01b8152336004820152602481019190915291516001600160a01b039091169363a9059cbb93604480850194919392918390030190829087803b1580156107dd57600080fd5b505af11580156107f1573d6000803e3d6000fd5b505050506040513d602081101561080757600080fd5b505161081257600080fd5b5060019998505050505050505050565b60008281526005602052604081205460ff1661083d57600080fd5b60008084815260066020526040902060080154600160a01b900460ff16600281111561086557fe5b1461086f57600080fd5b6000838152600660205260409020546001600160a01b0316331461089257600080fd5b600083836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209050600060038260405160200180838054600181600116156101000203166002900480156109285780601f10610906576101008083540402835291820191610928565b820191906000526020600020905b815481529060010190602001808311610914575b505082815260200192505050604051602081830303815290604052805190602001209050600060018a8a8a8a60405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156109a8573d6000803e3d6000fd5b505060408051601f1901516000898152600660205291909120549092506001600160a01b0380841691161490506109db57fe5b8982146109e457fe5b6000868152600660205260409020600801805460ff60a01b1916600160a01b17905560045460ff1615610a6157604080518b8152602081018590528082018490526001600160a01b038316606082015290517f6b7f08f909e42085ce2477fcf9f51be4ff86e5bcc7936c2dbe39e6757d33f9149181900360800190a15b60405186907ff857a98ed55303cc061ea3bc4c498fe49361c82cb72894280249b89848f9ce4690600090a26040516001600160a01b038216907f5c001d7523e0202fd001401f8393d128fbb1fafa8f637949f7d231d814f93b1390600090a25060019998505050505050505050565b60008181526005602052604081205460ff16610aeb57600080fd5b60008083815260066020526040902060080154600160a01b900460ff166002811115610b1357fe5b1480610b4557506001600083815260066020526040902060080154600160a01b900460ff166002811115610b4357fe5b145b610b4e57600080fd5b60045460ff16610b75576000828152600660205260409020600301544211610b7557600080fd5b6000828152600660205260409020546001600160a01b03163314610b9857600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610bf4573d6000803e3d6000fd5b505060408051601f1901516000868152600660205291909120549092506001600160a01b038084169116149050610c2757fe5b600083815260066020526040902060088101805460ff60a01b1916600160a11b1790554260038201556002015415610cff57600083815260066020908152604080832060028101805490859055600890910154825163a9059cbb60e01b815233600482015260248101839052925191946001600160a01b039091169363a9059cbb936044808201949293918390030190829087803b158015610cc857600080fd5b505af1158015610cdc573d6000803e3d6000fd5b505050506040513d6020811015610cf257600080fd5b5051610cfd57600080fd5b505b60405183907fceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba7290600090a25060019695505050505050565b600080546001600160a01b03163314610d4f57600080fd5b6001546001600160a01b0383811691161415610d6a57600080fd5b600180546001600160a01b0319166001600160a01b03841690811782556040517fe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e990600090a3919050565b60066020819052600091825260409091208054600182015460028301546003840154600485015460058601549686015460078701546008909701546001600160a01b0396871698958716979496939592949293919291811690600160a01b900460ff168a565b6000546001600160a01b031681565b600060018585858560405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610e86573d6000803e3d6000fd5b5050604051601f1901519695505050505050565b60045460ff1681565b600080546001600160a01b0316331480610ec757506001546001600160a01b031633145b610ed057600080fd5b60045460ff16610edf57600080fd5b60405133904780156108fc02916000818181858888f19350505050158015610f0b573d6000803e3d6000fd5b506001905090565b600080546001600160a01b03163314610f2b57600080fd5b6000546001600160a01b0383811691161415610f4657600080fd5b600080546001600160a01b0319166001600160a01b038416908117825560405160019233917f7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa589190a4919050565b6040805133606090811b6020808401919091529087901b6bffffffffffffffffffffffff191660348301524260488084018290528451808503909101815260689093018452825192820192909220600081815260059092529281205490929060ff161561100057600080fd5b600081815260066020908152604080832080546001600160a01b031990811633178255600180830180546001600160a01b038f16931683179055600283018c905542620151808c020160038401556004830188905560058084018b905560078401889055600890930180546001600160a81b0319169092179091559252808320805460ff19169092179091555182917f7ffc675d721b8768e035a323722842743bb523487b535906abc97e6b3e2095d091a260008181526006602090815260408083206008015481516323b872dd60e01b8152336004820152306024820152604481018b905291516001600160a01b03909116936323b872dd93606480850194919392918390030190829087803b15801561111a57600080fd5b505af115801561112e573d6000803e3d6000fd5b505050506040513d602081101561114457600080fd5b505161114f57600080fd5b5060019695505050505050565b6001546001600160a01b031681565b6040518060400160405280600a815260200169302e302e35616c70686160b01b81525081565b6000828211156111a057600080fd5b50900390565b6000828201838110156111b857600080fd5b939250505056fea2646970667358221220270639091140375d4a12c44e609a244a78f56046eaecbfb9b6ebfafd909425c764736f6c63430007000033"

// DeployAirdrop deploys a new Ethereum contract, binding an instance of Airdrop to it.
func DeployAirdrop(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Airdrop, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AirdropBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// Airdrop is an auto generated Go binding around an Ethereum contract.
type Airdrop struct {
	AirdropCaller     // Read-only binding to the contract
	AirdropTransactor // Write-only binding to the contract
	AirdropFilterer   // Log filterer for contract events
}

// AirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropSession struct {
	Contract     *Airdrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropCallerSession struct {
	Contract *AirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropTransactorSession struct {
	Contract     *AirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropRaw struct {
	Contract *Airdrop // Generic contract binding to access the raw methods on
}

// AirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropCallerRaw struct {
	Contract *AirdropCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropTransactorRaw struct {
	Contract *AirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdrop creates a new instance of Airdrop, bound to a specific deployed contract.
func NewAirdrop(address common.Address, backend bind.ContractBackend) (*Airdrop, error) {
	contract, err := bindAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// NewAirdropCaller creates a new read-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropCaller(address common.Address, caller bind.ContractCaller) (*AirdropCaller, error) {
	contract, err := bindAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropCaller{contract: contract}, nil
}

// NewAirdropTransactor creates a new write-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropTransactor, error) {
	contract, err := bindAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropTransactor{contract: contract}, nil
}

// NewAirdropFilterer creates a new log filterer instance of Airdrop, bound to a specific deployed contract.
func NewAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropFilterer, error) {
	contract, err := bindAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropFilterer{contract: contract}, nil
}

// bindAirdrop binds a generic wrapper to an already deployed contract.
func bindAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.AirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Airdrop *AirdropCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Airdrop.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Airdrop *AirdropSession) VERSION() (string, error) {
	return _Airdrop.Contract.VERSION(&_Airdrop.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Airdrop *AirdropCallerSession) VERSION() (string, error) {
	return _Airdrop.Contract.VERSION(&_Airdrop.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Airdrop *AirdropCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Airdrop.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Airdrop *AirdropSession) Admin() (common.Address, error) {
	return _Airdrop.Contract.Admin(&_Airdrop.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Airdrop *AirdropCallerSession) Admin() (common.Address, error) {
	return _Airdrop.Contract.Admin(&_Airdrop.CallOpts)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address source, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, uint256 dropAmount, uint256 totalDrops, bytes32 channelId, address intf, uint8 state)
func (_Airdrop *AirdropCaller) Channels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Source       common.Address
	TokenAddress common.Address
	Value        *big.Int
	ClosingDate  *big.Int
	OpenDate     *big.Int
	DropAmount   *big.Int
	TotalDrops   *big.Int
	ChannelId    [32]byte
	Intf         common.Address
	State        uint8
}, error) {
	ret := new(struct {
		Source       common.Address
		TokenAddress common.Address
		Value        *big.Int
		ClosingDate  *big.Int
		OpenDate     *big.Int
		DropAmount   *big.Int
		TotalDrops   *big.Int
		ChannelId    [32]byte
		Intf         common.Address
		State        uint8
	})
	out := ret
	err := _Airdrop.contract.Call(opts, out, "channels", arg0)
	return *ret, err
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address source, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, uint256 dropAmount, uint256 totalDrops, bytes32 channelId, address intf, uint8 state)
func (_Airdrop *AirdropSession) Channels(arg0 [32]byte) (struct {
	Source       common.Address
	TokenAddress common.Address
	Value        *big.Int
	ClosingDate  *big.Int
	OpenDate     *big.Int
	DropAmount   *big.Int
	TotalDrops   *big.Int
	ChannelId    [32]byte
	Intf         common.Address
	State        uint8
}, error) {
	return _Airdrop.Contract.Channels(&_Airdrop.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address source, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, uint256 dropAmount, uint256 totalDrops, bytes32 channelId, address intf, uint8 state)
func (_Airdrop *AirdropCallerSession) Channels(arg0 [32]byte) (struct {
	Source       common.Address
	TokenAddress common.Address
	Value        *big.Int
	ClosingDate  *big.Int
	OpenDate     *big.Int
	DropAmount   *big.Int
	TotalDrops   *big.Int
	ChannelId    [32]byte
	Intf         common.Address
	State        uint8
}, error) {
	return _Airdrop.Contract.Channels(&_Airdrop.CallOpts, arg0)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Airdrop *AirdropCaller) Dev(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Airdrop.contract.Call(opts, out, "dev")
	return *ret0, err
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Airdrop *AirdropSession) Dev() (bool, error) {
	return _Airdrop.Contract.Dev(&_Airdrop.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Airdrop *AirdropCallerSession) Dev() (bool, error) {
	return _Airdrop.Contract.Dev(&_Airdrop.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Airdrop *AirdropCaller) Frozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Airdrop.contract.Call(opts, out, "frozen")
	return *ret0, err
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Airdrop *AirdropSession) Frozen() (bool, error) {
	return _Airdrop.Contract.Frozen(&_Airdrop.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Airdrop *AirdropCallerSession) Frozen() (bool, error) {
	return _Airdrop.Contract.Frozen(&_Airdrop.CallOpts)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Airdrop *AirdropCaller) Moderators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Airdrop.contract.Call(opts, out, "moderators", arg0)
	return *ret0, err
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Airdrop *AirdropSession) Moderators(arg0 common.Address) (bool, error) {
	return _Airdrop.Contract.Moderators(&_Airdrop.CallOpts, arg0)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Airdrop *AirdropCallerSession) Moderators(arg0 common.Address) (bool, error) {
	return _Airdrop.Contract.Moderators(&_Airdrop.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Airdrop.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropCallerSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0x91456d3d.
//
// Solidity: function verifyProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_Airdrop *AirdropCaller) VerifyProof(opts *bind.CallOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Airdrop.contract.Call(opts, out, "verifyProof", _h, _v, _r, _s)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0x91456d3d.
//
// Solidity: function verifyProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_Airdrop *AirdropSession) VerifyProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _Airdrop.Contract.VerifyProof(&_Airdrop.CallOpts, _h, _v, _r, _s)
}

// VerifyProof is a free data retrieval call binding the contract method 0x91456d3d.
//
// Solidity: function verifyProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_Airdrop *AirdropCallerSession) VerifyProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _Airdrop.Contract.VerifyProof(&_Airdrop.CallOpts, _h, _v, _r, _s)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x5c9bd273.
//
// Solidity: function closeChannel(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Airdrop *AirdropTransactor) CloseChannel(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "closeChannel", _h, _v, _r, _s, _channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x5c9bd273.
//
// Solidity: function closeChannel(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Airdrop *AirdropSession) CloseChannel(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.CloseChannel(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x5c9bd273.
//
// Solidity: function closeChannel(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Airdrop *AirdropTransactorSession) CloseChannel(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.CloseChannel(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId)
}

// EnableAirdrops is a paid mutator transaction binding the contract method 0x56715148.
//
// Solidity: function enableAirdrops(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _id) returns(bool)
func (_Airdrop *AirdropTransactor) EnableAirdrops(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "enableAirdrops", _h, _v, _r, _s, _channelId, _id)
}

// EnableAirdrops is a paid mutator transaction binding the contract method 0x56715148.
//
// Solidity: function enableAirdrops(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _id) returns(bool)
func (_Airdrop *AirdropSession) EnableAirdrops(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.EnableAirdrops(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId, _id)
}

// EnableAirdrops is a paid mutator transaction binding the contract method 0x56715148.
//
// Solidity: function enableAirdrops(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _id) returns(bool)
func (_Airdrop *AirdropTransactorSession) EnableAirdrops(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.EnableAirdrops(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId, _id)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xf36bdb49.
//
// Solidity: function openChannel(address _tokenAddress, uint256 _channelValue, uint256 _durationInDays, uint256 _dropAmount) payable returns(bool)
func (_Airdrop *AirdropTransactor) OpenChannel(opts *bind.TransactOpts, _tokenAddress common.Address, _channelValue *big.Int, _durationInDays *big.Int, _dropAmount *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "openChannel", _tokenAddress, _channelValue, _durationInDays, _dropAmount)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xf36bdb49.
//
// Solidity: function openChannel(address _tokenAddress, uint256 _channelValue, uint256 _durationInDays, uint256 _dropAmount) payable returns(bool)
func (_Airdrop *AirdropSession) OpenChannel(_tokenAddress common.Address, _channelValue *big.Int, _durationInDays *big.Int, _dropAmount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.OpenChannel(&_Airdrop.TransactOpts, _tokenAddress, _channelValue, _durationInDays, _dropAmount)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xf36bdb49.
//
// Solidity: function openChannel(address _tokenAddress, uint256 _channelValue, uint256 _durationInDays, uint256 _dropAmount) payable returns(bool)
func (_Airdrop *AirdropTransactorSession) OpenChannel(_tokenAddress common.Address, _channelValue *big.Int, _durationInDays *big.Int, _dropAmount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.OpenChannel(&_Airdrop.TransactOpts, _tokenAddress, _channelValue, _durationInDays, _dropAmount)
}

// RetrieveAirdrop is a paid mutator transaction binding the contract method 0x515982f6.
//
// Solidity: function retrieveAirdrop(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _id) returns(bool)
func (_Airdrop *AirdropTransactor) RetrieveAirdrop(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "retrieveAirdrop", _h, _v, _r, _s, _channelId, _id)
}

// RetrieveAirdrop is a paid mutator transaction binding the contract method 0x515982f6.
//
// Solidity: function retrieveAirdrop(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _id) returns(bool)
func (_Airdrop *AirdropSession) RetrieveAirdrop(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.RetrieveAirdrop(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId, _id)
}

// RetrieveAirdrop is a paid mutator transaction binding the contract method 0x515982f6.
//
// Solidity: function retrieveAirdrop(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _id) returns(bool)
func (_Airdrop *AirdropTransactorSession) RetrieveAirdrop(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.RetrieveAirdrop(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId, _id)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Airdrop *AirdropTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Airdrop *AirdropSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.SetAdmin(&_Airdrop.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Airdrop *AirdropTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.SetAdmin(&_Airdrop.TransactOpts, _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Airdrop *AirdropTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Airdrop *AirdropSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Airdrop *AirdropTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, _newOwner)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Airdrop *AirdropTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Airdrop *AirdropSession) WithdrawEth() (*types.Transaction, error) {
	return _Airdrop.Contract.WithdrawEth(&_Airdrop.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Airdrop *AirdropTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _Airdrop.Contract.WithdrawEth(&_Airdrop.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Airdrop *AirdropTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Airdrop.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Airdrop *AirdropSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.Fallback(&_Airdrop.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Airdrop *AirdropTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.Fallback(&_Airdrop.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Airdrop *AirdropTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Airdrop *AirdropSession) Receive() (*types.Transaction, error) {
	return _Airdrop.Contract.Receive(&_Airdrop.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Airdrop *AirdropTransactorSession) Receive() (*types.Transaction, error) {
	return _Airdrop.Contract.Receive(&_Airdrop.TransactOpts)
}

// AirdropAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the Airdrop contract.
type AirdropAdminSetIterator struct {
	Event *AirdropAdminSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropAdminSet represents a AdminSet event raised by the Airdrop contract.
type AirdropAdminSet struct {
	Admin    common.Address
	AdminSet bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(address indexed _admin, bool indexed _adminSet)
func (_Airdrop *AirdropFilterer) FilterAdminSet(opts *bind.FilterOpts, _admin []common.Address, _adminSet []bool) (*AirdropAdminSetIterator, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return &AirdropAdminSetIterator{contract: _Airdrop.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(address indexed _admin, bool indexed _adminSet)
func (_Airdrop *AirdropFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *AirdropAdminSet, _admin []common.Address, _adminSet []bool) (event.Subscription, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropAdminSet)
				if err := _Airdrop.contract.UnpackLog(event, "AdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminSet is a log parse operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(address indexed _admin, bool indexed _adminSet)
func (_Airdrop *AirdropFilterer) ParseAdminSet(log types.Log) (*AirdropAdminSet, error) {
	event := new(AirdropAdminSet)
	if err := _Airdrop.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AirdropAirDropDispersedIterator is returned from FilterAirDropDispersed and is used to iterate over the raw logs and unpacked data for AirDropDispersed events raised by the Airdrop contract.
type AirdropAirDropDispersedIterator struct {
	Event *AirdropAirDropDispersed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropAirDropDispersedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropAirDropDispersed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropAirDropDispersed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropAirDropDispersedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropAirDropDispersedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropAirDropDispersed represents a AirDropDispersed event raised by the Airdrop contract.
type AirdropAirDropDispersed struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAirDropDispersed is a free log retrieval operation binding the contract event 0xa467a409776bfc50e32dd8496edd970b5e7a84a38acda61529d8022bf757c2fd.
//
// Solidity: event AirDropDispersed(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) FilterAirDropDispersed(opts *bind.FilterOpts, _channelId [][32]byte) (*AirdropAirDropDispersedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "AirDropDispersed", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AirdropAirDropDispersedIterator{contract: _Airdrop.contract, event: "AirDropDispersed", logs: logs, sub: sub}, nil
}

// WatchAirDropDispersed is a free log subscription operation binding the contract event 0xa467a409776bfc50e32dd8496edd970b5e7a84a38acda61529d8022bf757c2fd.
//
// Solidity: event AirDropDispersed(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) WatchAirDropDispersed(opts *bind.WatchOpts, sink chan<- *AirdropAirDropDispersed, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "AirDropDispersed", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropAirDropDispersed)
				if err := _Airdrop.contract.UnpackLog(event, "AirDropDispersed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAirDropDispersed is a log parse operation binding the contract event 0xa467a409776bfc50e32dd8496edd970b5e7a84a38acda61529d8022bf757c2fd.
//
// Solidity: event AirDropDispersed(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) ParseAirDropDispersed(log types.Log) (*AirdropAirDropDispersed, error) {
	event := new(AirdropAirDropDispersed)
	if err := _Airdrop.contract.UnpackLog(event, "AirDropDispersed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AirdropAirDropsEnabledIterator is returned from FilterAirDropsEnabled and is used to iterate over the raw logs and unpacked data for AirDropsEnabled events raised by the Airdrop contract.
type AirdropAirDropsEnabledIterator struct {
	Event *AirdropAirDropsEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropAirDropsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropAirDropsEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropAirDropsEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropAirDropsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropAirDropsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropAirDropsEnabled represents a AirDropsEnabled event raised by the Airdrop contract.
type AirdropAirDropsEnabled struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAirDropsEnabled is a free log retrieval operation binding the contract event 0xf857a98ed55303cc061ea3bc4c498fe49361c82cb72894280249b89848f9ce46.
//
// Solidity: event AirDropsEnabled(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) FilterAirDropsEnabled(opts *bind.FilterOpts, _channelId [][32]byte) (*AirdropAirDropsEnabledIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "AirDropsEnabled", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AirdropAirDropsEnabledIterator{contract: _Airdrop.contract, event: "AirDropsEnabled", logs: logs, sub: sub}, nil
}

// WatchAirDropsEnabled is a free log subscription operation binding the contract event 0xf857a98ed55303cc061ea3bc4c498fe49361c82cb72894280249b89848f9ce46.
//
// Solidity: event AirDropsEnabled(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) WatchAirDropsEnabled(opts *bind.WatchOpts, sink chan<- *AirdropAirDropsEnabled, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "AirDropsEnabled", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropAirDropsEnabled)
				if err := _Airdrop.contract.UnpackLog(event, "AirDropsEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAirDropsEnabled is a log parse operation binding the contract event 0xf857a98ed55303cc061ea3bc4c498fe49361c82cb72894280249b89848f9ce46.
//
// Solidity: event AirDropsEnabled(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) ParseAirDropsEnabled(log types.Log) (*AirdropAirDropsEnabled, error) {
	event := new(AirdropAirDropsEnabled)
	if err := _Airdrop.contract.UnpackLog(event, "AirDropsEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AirdropChannelClosedIterator is returned from FilterChannelClosed and is used to iterate over the raw logs and unpacked data for ChannelClosed events raised by the Airdrop contract.
type AirdropChannelClosedIterator struct {
	Event *AirdropChannelClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropChannelClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropChannelClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropChannelClosed represents a ChannelClosed event raised by the Airdrop contract.
type AirdropChannelClosed struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelClosed is a free log retrieval operation binding the contract event 0xceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba72.
//
// Solidity: event ChannelClosed(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) FilterChannelClosed(opts *bind.FilterOpts, _channelId [][32]byte) (*AirdropChannelClosedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ChannelClosed", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AirdropChannelClosedIterator{contract: _Airdrop.contract, event: "ChannelClosed", logs: logs, sub: sub}, nil
}

// WatchChannelClosed is a free log subscription operation binding the contract event 0xceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba72.
//
// Solidity: event ChannelClosed(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) WatchChannelClosed(opts *bind.WatchOpts, sink chan<- *AirdropChannelClosed, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ChannelClosed", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropChannelClosed)
				if err := _Airdrop.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelClosed is a log parse operation binding the contract event 0xceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba72.
//
// Solidity: event ChannelClosed(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) ParseChannelClosed(log types.Log) (*AirdropChannelClosed, error) {
	event := new(AirdropChannelClosed)
	if err := _Airdrop.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AirdropChannelOpenedIterator is returned from FilterChannelOpened and is used to iterate over the raw logs and unpacked data for ChannelOpened events raised by the Airdrop contract.
type AirdropChannelOpenedIterator struct {
	Event *AirdropChannelOpened // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropChannelOpened)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropChannelOpened)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropChannelOpened represents a ChannelOpened event raised by the Airdrop contract.
type AirdropChannelOpened struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelOpened is a free log retrieval operation binding the contract event 0x7ffc675d721b8768e035a323722842743bb523487b535906abc97e6b3e2095d0.
//
// Solidity: event ChannelOpened(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) FilterChannelOpened(opts *bind.FilterOpts, _channelId [][32]byte) (*AirdropChannelOpenedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &AirdropChannelOpenedIterator{contract: _Airdrop.contract, event: "ChannelOpened", logs: logs, sub: sub}, nil
}

// WatchChannelOpened is a free log subscription operation binding the contract event 0x7ffc675d721b8768e035a323722842743bb523487b535906abc97e6b3e2095d0.
//
// Solidity: event ChannelOpened(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) WatchChannelOpened(opts *bind.WatchOpts, sink chan<- *AirdropChannelOpened, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropChannelOpened)
				if err := _Airdrop.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelOpened is a log parse operation binding the contract event 0x7ffc675d721b8768e035a323722842743bb523487b535906abc97e6b3e2095d0.
//
// Solidity: event ChannelOpened(bytes32 indexed _channelId)
func (_Airdrop *AirdropFilterer) ParseChannelOpened(log types.Log) (*AirdropChannelOpened, error) {
	event := new(AirdropChannelOpened)
	if err := _Airdrop.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AirdropOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Airdrop contract.
type AirdropOwnershipTransferredIterator struct {
	Event *AirdropOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropOwnershipTransferred represents a OwnershipTransferred event raised by the Airdrop contract.
type AirdropOwnershipTransferred struct {
	PreviousOwner        common.Address
	NewOwner             common.Address
	OwnershipTransferred bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner, bool indexed _ownershipTransferred)
func (_Airdrop *AirdropFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (*AirdropOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}
	var _ownershipTransferredRule []interface{}
	for _, _ownershipTransferredItem := range _ownershipTransferred {
		_ownershipTransferredRule = append(_ownershipTransferredRule, _ownershipTransferredItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return &AirdropOwnershipTransferredIterator{contract: _Airdrop.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner, bool indexed _ownershipTransferred)
func (_Airdrop *AirdropFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AirdropOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}
	var _ownershipTransferredRule []interface{}
	for _, _ownershipTransferredItem := range _ownershipTransferred {
		_ownershipTransferredRule = append(_ownershipTransferredRule, _ownershipTransferredItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropOwnershipTransferred)
				if err := _Airdrop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner, bool indexed _ownershipTransferred)
func (_Airdrop *AirdropFilterer) ParseOwnershipTransferred(log types.Log) (*AirdropOwnershipTransferred, error) {
	event := new(AirdropOwnershipTransferred)
	if err := _Airdrop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AirdropSigDebugIterator is returned from FilterSigDebug and is used to iterate over the raw logs and unpacked data for SigDebug events raised by the Airdrop contract.
type AirdropSigDebugIterator struct {
	Event *AirdropSigDebug // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropSigDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropSigDebug)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropSigDebug)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropSigDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropSigDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropSigDebug represents a SigDebug event raised by the Airdrop contract.
type AirdropSigDebug struct {
	H             [32]byte
	Proof         [32]byte
	PrefixedProof [32]byte
	Signer        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSigDebug is a free log retrieval operation binding the contract event 0x6b7f08f909e42085ce2477fcf9f51be4ff86e5bcc7936c2dbe39e6757d33f914.
//
// Solidity: event SigDebug(bytes32 _h, bytes32 _proof, bytes32 prefixedProof, address signer)
func (_Airdrop *AirdropFilterer) FilterSigDebug(opts *bind.FilterOpts) (*AirdropSigDebugIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "SigDebug")
	if err != nil {
		return nil, err
	}
	return &AirdropSigDebugIterator{contract: _Airdrop.contract, event: "SigDebug", logs: logs, sub: sub}, nil
}

// WatchSigDebug is a free log subscription operation binding the contract event 0x6b7f08f909e42085ce2477fcf9f51be4ff86e5bcc7936c2dbe39e6757d33f914.
//
// Solidity: event SigDebug(bytes32 _h, bytes32 _proof, bytes32 prefixedProof, address signer)
func (_Airdrop *AirdropFilterer) WatchSigDebug(opts *bind.WatchOpts, sink chan<- *AirdropSigDebug) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "SigDebug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropSigDebug)
				if err := _Airdrop.contract.UnpackLog(event, "SigDebug", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSigDebug is a log parse operation binding the contract event 0x6b7f08f909e42085ce2477fcf9f51be4ff86e5bcc7936c2dbe39e6757d33f914.
//
// Solidity: event SigDebug(bytes32 _h, bytes32 _proof, bytes32 prefixedProof, address signer)
func (_Airdrop *AirdropFilterer) ParseSigDebug(log types.Log) (*AirdropSigDebug, error) {
	event := new(AirdropSigDebug)
	if err := _Airdrop.contract.UnpackLog(event, "SigDebug", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AirdropSignatureRecoveredIterator is returned from FilterSignatureRecovered and is used to iterate over the raw logs and unpacked data for SignatureRecovered events raised by the Airdrop contract.
type AirdropSignatureRecoveredIterator struct {
	Event *AirdropSignatureRecovered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropSignatureRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropSignatureRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropSignatureRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropSignatureRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropSignatureRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropSignatureRecovered represents a SignatureRecovered event raised by the Airdrop contract.
type AirdropSignatureRecovered struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignatureRecovered is a free log retrieval operation binding the contract event 0x5c001d7523e0202fd001401f8393d128fbb1fafa8f637949f7d231d814f93b13.
//
// Solidity: event SignatureRecovered(address indexed signer)
func (_Airdrop *AirdropFilterer) FilterSignatureRecovered(opts *bind.FilterOpts, signer []common.Address) (*AirdropSignatureRecoveredIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "SignatureRecovered", signerRule)
	if err != nil {
		return nil, err
	}
	return &AirdropSignatureRecoveredIterator{contract: _Airdrop.contract, event: "SignatureRecovered", logs: logs, sub: sub}, nil
}

// WatchSignatureRecovered is a free log subscription operation binding the contract event 0x5c001d7523e0202fd001401f8393d128fbb1fafa8f637949f7d231d814f93b13.
//
// Solidity: event SignatureRecovered(address indexed signer)
func (_Airdrop *AirdropFilterer) WatchSignatureRecovered(opts *bind.WatchOpts, sink chan<- *AirdropSignatureRecovered, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "SignatureRecovered", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropSignatureRecovered)
				if err := _Airdrop.contract.UnpackLog(event, "SignatureRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignatureRecovered is a log parse operation binding the contract event 0x5c001d7523e0202fd001401f8393d128fbb1fafa8f637949f7d231d814f93b13.
//
// Solidity: event SignatureRecovered(address indexed signer)
func (_Airdrop *AirdropFilterer) ParseSignatureRecovered(log types.Log) (*AirdropSignatureRecovered, error) {
	event := new(AirdropSignatureRecovered)
	if err := _Airdrop.contract.UnpackLog(event, "SignatureRecovered", log); err != nil {
		return nil, err
	}
	return event, nil
}
