// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// AirdropABI is the input ABI used to generate the binding from.
const AirdropABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"moderators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_h\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"retrieveAirdrop\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_h\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"enableAirdrops\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_h\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"closeChannel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"name\":\"source\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"closingDate\",\"type\":\"uint256\"},{\"name\":\"openDate\",\"type\":\"uint256\"},{\"name\":\"dropAmount\",\"type\":\"uint256\"},{\"name\":\"totalDrops\",\"type\":\"uint256\"},{\"name\":\"channelId\",\"type\":\"bytes32\"},{\"name\":\"intf\",\"type\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_h\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_channelValue\",\"type\":\"uint256\"},{\"name\":\"_durationInDays\",\"type\":\"uint256\"},{\"name\":\"_dropAmount\",\"type\":\"uint256\"}],\"name\":\"openChannel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"AirDropsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"AirDropDispersed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignatureRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_h\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_proof\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"proof\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SigDebug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_admin\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_adminSet\",\"type\":\"bool\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_ownershipTransferred\",\"type\":\"bool\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AirdropBin is the compiled bytecode used for deploying new contracts.
const AirdropBin = `60c0604052601c60808190527f19457468657265756d205369676e6564204d6573736167653a0a33320000000060a090815262000040916003919062000074565b506004805460ff1916600190811790915560008054600160a060020a03199081163390811790925582541617905562000119565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000b757805160ff1916838001178555620000e7565b82800160010185558215620000e7579182015b82811115620000e7578251825591602001919060010190620000ca565b50620000f5929150620000f9565b5090565b6200011691905b80821115620000f5576000815560010162000100565b90565b6113b680620001296000396000f3006080604052600436106100da5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663054f7d9c81146100dc57806314d0f1ba14610105578063515982f61461012657806356715148146101505780635c9bd2731461017a578063704b6c02146101a15780637a7ebd7b146101c25780638da5cb5b1461025357806391456d3d1461028457806391cca3db146102a8578063a0ef91df146102bd578063f2fde38b146102d2578063f36bdb49146102f3578063f851a44014610310578063ffa1ad7414610325575b005b3480156100e857600080fd5b506100f16103af565b604080519115158252519081900360200190f35b34801561011157600080fd5b506100f1600160a060020a03600435166103bf565b34801561013257600080fd5b506100f160043560ff6024351660443560643560843560a4356103d4565b34801561015c57600080fd5b506100f160043560ff6024351660443560643560843560a435610813565b34801561018657600080fd5b506100f160043560ff60243516604435606435608435610b93565b3480156101ad57600080fd5b506100f1600160a060020a0360043516610e40565b3480156101ce57600080fd5b506101da600435610ecb565b60408051600160a060020a03808d1682528b811660208301529181018a9052606081018990526080810188905260a0810187905260c0810186905260e08101859052908316610100820152610120810182600281111561023657fe5b60ff1681526020019a505050505050505050505060405180910390f35b34801561025f57600080fd5b50610268610f31565b60408051600160a060020a039092168252519081900360200190f35b34801561029057600080fd5b5061026860043560ff60243516604435606435610f40565b3480156102b457600080fd5b506100f1610fad565b3480156102c957600080fd5b506100f1610fb6565b3480156102de57600080fd5b506100f1600160a060020a036004351661102b565b6100f1600160a060020a03600435166024356044356064356110b9565b34801561031c57600080fd5b50610268611316565b34801561033157600080fd5b5061033a611325565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561037457818101518382015260200161035c565b50505050905090810190601f1680156103a15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60015460a060020a900460ff1681565b60026020526000908152604090205460ff1681565b60008281526005602052604081205481908190819060ff1615156103f757600080fd5b600160008781526006602052604090206008015460a060020a900460ff16600281111561042057fe5b1461042a57600080fd5b60008681526006602052604090206005810154600290910154101561044e57600080fd5b600086815260076020908152604080832033845290915290205460ff161561047557600080fd5b6040805160208082018990528183018890526c010000000000000000000000003302606083015282516054818403018152607490920192839052815191929182918401908083835b602083106104dc5780518252601f1990920191602091820191016104bd565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250600383604051602001808380546001816001161561010002031660029004801561056c5780601f1061054a57610100808354040283529182019161056c565b820191906000526020600020905b815481529060010190602001808311610558575b505091825250604080518083038152602092830191829052805190935090918291908401908083835b602083106105b45780518252601f199092019160209182019101610595565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915060018a8a8a8a604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af115801561065a573d6000803e3d6000fd5b505060408051601f190151600089815260066020529190912054909250600160a060020a03808416911614905061068d57fe5b818a1461069657fe5b60008681526007602090815260408083203384528252808320805460ff191660011790558883526006909152902060058101546002909101546106de9163ffffffff61135c16565b60008781526006602081905260409091206002810192909255015461070a90600163ffffffff61137116565b600087815260066020819052604080832090910192909255905187917fa467a409776bfc50e32dd8496edd970b5e7a84a38acda61529d8022bf757c2fd91a26000868152600660209081526040808320600881015460059091015482517fa9059cbb00000000000000000000000000000000000000000000000000000000815233600482015260248101919091529151600160a060020a039091169363a9059cbb93604480850194919392918390030190829087803b1580156107cc57600080fd5b505af11580156107e0573d6000803e3d6000fd5b505050506040513d60208110156107f657600080fd5b5051151561080357600080fd5b5060019998505050505050505050565b60008281526005602052604081205481908190819060ff16151561083657600080fd5b6000808781526006602052604090206008015460a060020a900460ff16600281111561085e57fe5b1461086857600080fd5b600086815260066020526040902054600160a060020a0316331461088b57600080fd5b604080516020808201899052818301889052825180830384018152606090920192839052815191929182918401908083835b602083106108dc5780518252601f1990920191602091820191016108bd565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250600383604051602001808380546001816001161561010002031660029004801561096c5780601f1061094a57610100808354040283529182019161096c565b820191906000526020600020905b815481529060010190602001808311610958575b505091825250604080518083038152602092830191829052805190935090918291908401908083835b602083106109b45780518252601f199092019160209182019101610995565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915060018a8a8a8a604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015610a5a573d6000803e3d6000fd5b505060408051601f190151600089815260066020529190912054909250600160a060020a038084169116149050610a8d57fe5b818a14610a9657fe5b6000868152600660205260409020600801805474ff0000000000000000000000000000000000000000191660a060020a17905560045460ff1615610b2457604080518b815260208101859052808201849052600160a060020a038316606082015290517f6b7f08f909e42085ce2477fcf9f51be4ff86e5bcc7936c2dbe39e6757d33f9149181900360800190a15b60405186907ff857a98ed55303cc061ea3bc4c498fe49361c82cb72894280249b89848f9ce4690600090a2604051600160a060020a038216907f5c001d7523e0202fd001401f8393d128fbb1fafa8f637949f7d231d814f93b1390600090a25060019998505050505050505050565b6000818152600560205260408120548190819060ff161515610bb457600080fd5b6000808581526006602052604090206008015460a060020a900460ff166002811115610bdc57fe5b1480610c0e5750600160008581526006602052604090206008015460a060020a900460ff166002811115610c0c57fe5b145b1515610c1957600080fd5b60045460ff161515610c42576000848152600660205260409020600301544211610c4257600080fd5b600084815260066020526040902054600160a060020a03163314610c6557600080fd5b60408051600080825260208083018085528c905260ff8b1683850152606083018a905260808301899052925160019360a0808501949193601f19840193928390039091019190865af1158015610cbf573d6000803e3d6000fd5b505060408051601f190151600087815260066020529190912054909350600160a060020a038085169116149050610cf257fe5b600084815260066020526040812060088101805474ff0000000000000000000000000000000000000000191674020000000000000000000000000000000000000000179055426003820155600201541115610e07575060008381526006602090815260408083206002810180549085905560089091015482517fa9059cbb0000000000000000000000000000000000000000000000000000000081523360048201526024810183905292519194600160a060020a039091169363a9059cbb936044808201949293918390030190829087803b158015610dd057600080fd5b505af1158015610de4573d6000803e3d6000fd5b505050506040513d6020811015610dfa57600080fd5b50511515610e0757600080fd5b60405184907fceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba7290600090a2506001979650505050505050565b60008054600160a060020a03163314610e5857600080fd5b600154600160a060020a0383811691161415610e7357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03841690811782556040517fe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e990600090a3919050565b6006602081905260009182526040909120805460018201546002830154600384015460048501546005860154968601546007870154600890970154600160a060020a039687169895871697949693959294929391929181169060a060020a900460ff168a565b600054600160a060020a031681565b604080516000808252602080830180855288905260ff87168385015260608301869052608083018590529251909260019260a080820193601f198101928190039091019086865af1158015610f99573d6000803e3d6000fd5b5050604051601f1901519695505050505050565b60045460ff1681565b60008054600160a060020a0316331480610fda5750600154600160a060020a031633145b1515610fe557600080fd5b60045460ff161515610ff657600080fd5b6040513390303180156108fc02916000818181858888f19350505050158015611023573d6000803e3d6000fd5b506001905090565b60008054600160a060020a0316331461104357600080fd5b600054600160a060020a038381169116141561105e57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038416908117825560405160019233917f7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa589190a4919050565b604080516c01000000000000000000000000338102602080840191909152600160a060020a03881690910260348301524260488084018290528451808503909101815260689093019384905282516000949193859390929182918401908083835b602083106111395780518252601f19909201916020918201910161111a565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008181526005909252929020549194505060ff16159150611184905057600080fd5b6000818152600660209081526040808320805473ffffffffffffffffffffffffffffffffffffffff199081163317825560018083018054600160a060020a038f16931683179055600283018c905542620151808c020160038401556004830188905560058084018b9055600784018890556008909301805474ffffffffffffffffffffffffffffffffffffffffff19169092179091559252808320805460ff19169092179091555182917f7ffc675d721b8768e035a323722842743bb523487b535906abc97e6b3e2095d091a260008181526006602090815260408083206008015481517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018b90529151600160a060020a03909116936323b872dd93606480850194919392918390030190829087803b1580156112d257600080fd5b505af11580156112e6573d6000803e3d6000fd5b505050506040513d60208110156112fc57600080fd5b5051151561130957600080fd5b5060019695505050505050565b600154600160a060020a031681565b60408051808201909152600a81527f302e302e35616c70686100000000000000000000000000000000000000000000602082015281565b60008282111561136b57600080fd5b50900390565b60008282018381101561138357600080fd5b93925050505600a165627a7a72305820dcf23ce18912d1f7161bae0c3b73119901a432b213df7c9ba72923f7a2ae6ae10029`

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
// Solidity: function VERSION() constant returns(string)
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
// Solidity: function VERSION() constant returns(string)
func (_Airdrop *AirdropSession) VERSION() (string, error) {
	return _Airdrop.Contract.VERSION(&_Airdrop.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Airdrop *AirdropCallerSession) VERSION() (string, error) {
	return _Airdrop.Contract.VERSION(&_Airdrop.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
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
// Solidity: function admin() constant returns(address)
func (_Airdrop *AirdropSession) Admin() (common.Address, error) {
	return _Airdrop.Contract.Admin(&_Airdrop.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Airdrop *AirdropCallerSession) Admin() (common.Address, error) {
	return _Airdrop.Contract.Admin(&_Airdrop.CallOpts)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels( bytes32) constant returns(source address, tokenAddress address, value uint256, closingDate uint256, openDate uint256, dropAmount uint256, totalDrops uint256, channelId bytes32, intf address, state uint8)
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
// Solidity: function channels( bytes32) constant returns(source address, tokenAddress address, value uint256, closingDate uint256, openDate uint256, dropAmount uint256, totalDrops uint256, channelId bytes32, intf address, state uint8)
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
// Solidity: function channels( bytes32) constant returns(source address, tokenAddress address, value uint256, closingDate uint256, openDate uint256, dropAmount uint256, totalDrops uint256, channelId bytes32, intf address, state uint8)
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
// Solidity: function dev() constant returns(bool)
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
// Solidity: function dev() constant returns(bool)
func (_Airdrop *AirdropSession) Dev() (bool, error) {
	return _Airdrop.Contract.Dev(&_Airdrop.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() constant returns(bool)
func (_Airdrop *AirdropCallerSession) Dev() (bool, error) {
	return _Airdrop.Contract.Dev(&_Airdrop.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
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
// Solidity: function frozen() constant returns(bool)
func (_Airdrop *AirdropSession) Frozen() (bool, error) {
	return _Airdrop.Contract.Frozen(&_Airdrop.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() constant returns(bool)
func (_Airdrop *AirdropCallerSession) Frozen() (bool, error) {
	return _Airdrop.Contract.Frozen(&_Airdrop.CallOpts)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
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
// Solidity: function moderators( address) constant returns(bool)
func (_Airdrop *AirdropSession) Moderators(arg0 common.Address) (bool, error) {
	return _Airdrop.Contract.Moderators(&_Airdrop.CallOpts, arg0)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators( address) constant returns(bool)
func (_Airdrop *AirdropCallerSession) Moderators(arg0 common.Address) (bool, error) {
	return _Airdrop.Contract.Moderators(&_Airdrop.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
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
// Solidity: function owner() constant returns(address)
func (_Airdrop *AirdropSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Airdrop *AirdropCallerSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0x91456d3d.
//
// Solidity: function verifyProof(_h bytes32, _v uint8, _r bytes32, _s bytes32) constant returns(address)
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
// Solidity: function verifyProof(_h bytes32, _v uint8, _r bytes32, _s bytes32) constant returns(address)
func (_Airdrop *AirdropSession) VerifyProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _Airdrop.Contract.VerifyProof(&_Airdrop.CallOpts, _h, _v, _r, _s)
}

// VerifyProof is a free data retrieval call binding the contract method 0x91456d3d.
//
// Solidity: function verifyProof(_h bytes32, _v uint8, _r bytes32, _s bytes32) constant returns(address)
func (_Airdrop *AirdropCallerSession) VerifyProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _Airdrop.Contract.VerifyProof(&_Airdrop.CallOpts, _h, _v, _r, _s)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x5c9bd273.
//
// Solidity: function closeChannel(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32) returns(bool)
func (_Airdrop *AirdropTransactor) CloseChannel(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "closeChannel", _h, _v, _r, _s, _channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x5c9bd273.
//
// Solidity: function closeChannel(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32) returns(bool)
func (_Airdrop *AirdropSession) CloseChannel(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.CloseChannel(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x5c9bd273.
//
// Solidity: function closeChannel(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32) returns(bool)
func (_Airdrop *AirdropTransactorSession) CloseChannel(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.CloseChannel(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId)
}

// EnableAirdrops is a paid mutator transaction binding the contract method 0x56715148.
//
// Solidity: function enableAirdrops(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32, _id uint256) returns(bool)
func (_Airdrop *AirdropTransactor) EnableAirdrops(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "enableAirdrops", _h, _v, _r, _s, _channelId, _id)
}

// EnableAirdrops is a paid mutator transaction binding the contract method 0x56715148.
//
// Solidity: function enableAirdrops(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32, _id uint256) returns(bool)
func (_Airdrop *AirdropSession) EnableAirdrops(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.EnableAirdrops(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId, _id)
}

// EnableAirdrops is a paid mutator transaction binding the contract method 0x56715148.
//
// Solidity: function enableAirdrops(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32, _id uint256) returns(bool)
func (_Airdrop *AirdropTransactorSession) EnableAirdrops(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.EnableAirdrops(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId, _id)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xf36bdb49.
//
// Solidity: function openChannel(_tokenAddress address, _channelValue uint256, _durationInDays uint256, _dropAmount uint256) returns(bool)
func (_Airdrop *AirdropTransactor) OpenChannel(opts *bind.TransactOpts, _tokenAddress common.Address, _channelValue *big.Int, _durationInDays *big.Int, _dropAmount *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "openChannel", _tokenAddress, _channelValue, _durationInDays, _dropAmount)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xf36bdb49.
//
// Solidity: function openChannel(_tokenAddress address, _channelValue uint256, _durationInDays uint256, _dropAmount uint256) returns(bool)
func (_Airdrop *AirdropSession) OpenChannel(_tokenAddress common.Address, _channelValue *big.Int, _durationInDays *big.Int, _dropAmount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.OpenChannel(&_Airdrop.TransactOpts, _tokenAddress, _channelValue, _durationInDays, _dropAmount)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xf36bdb49.
//
// Solidity: function openChannel(_tokenAddress address, _channelValue uint256, _durationInDays uint256, _dropAmount uint256) returns(bool)
func (_Airdrop *AirdropTransactorSession) OpenChannel(_tokenAddress common.Address, _channelValue *big.Int, _durationInDays *big.Int, _dropAmount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.OpenChannel(&_Airdrop.TransactOpts, _tokenAddress, _channelValue, _durationInDays, _dropAmount)
}

// RetrieveAirdrop is a paid mutator transaction binding the contract method 0x515982f6.
//
// Solidity: function retrieveAirdrop(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32, _id uint256) returns(bool)
func (_Airdrop *AirdropTransactor) RetrieveAirdrop(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "retrieveAirdrop", _h, _v, _r, _s, _channelId, _id)
}

// RetrieveAirdrop is a paid mutator transaction binding the contract method 0x515982f6.
//
// Solidity: function retrieveAirdrop(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32, _id uint256) returns(bool)
func (_Airdrop *AirdropSession) RetrieveAirdrop(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.RetrieveAirdrop(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId, _id)
}

// RetrieveAirdrop is a paid mutator transaction binding the contract method 0x515982f6.
//
// Solidity: function retrieveAirdrop(_h bytes32, _v uint8, _r bytes32, _s bytes32, _channelId bytes32, _id uint256) returns(bool)
func (_Airdrop *AirdropTransactorSession) RetrieveAirdrop(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.RetrieveAirdrop(&_Airdrop.TransactOpts, _h, _v, _r, _s, _channelId, _id)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_Airdrop *AirdropTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_Airdrop *AirdropSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.SetAdmin(&_Airdrop.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_newAdmin address) returns(bool)
func (_Airdrop *AirdropTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.SetAdmin(&_Airdrop.TransactOpts, _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Airdrop *AirdropTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
func (_Airdrop *AirdropSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns(bool)
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
// Solidity: e AdminSet(_admin indexed address, _adminSet indexed bool)
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
// Solidity: e AdminSet(_admin indexed address, _adminSet indexed bool)
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
// Solidity: e AirDropDispersed(_channelId indexed bytes32)
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
// Solidity: e AirDropDispersed(_channelId indexed bytes32)
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
// Solidity: e AirDropsEnabled(_channelId indexed bytes32)
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
// Solidity: e AirDropsEnabled(_channelId indexed bytes32)
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
// Solidity: e ChannelClosed(_channelId indexed bytes32)
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
// Solidity: e ChannelClosed(_channelId indexed bytes32)
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
// Solidity: e ChannelOpened(_channelId indexed bytes32)
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
// Solidity: e ChannelOpened(_channelId indexed bytes32)
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
// Solidity: e OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address, _ownershipTransferred indexed bool)
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
// Solidity: e OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address, _ownershipTransferred indexed bool)
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
	H      [32]byte
	Proof  [32]byte
	Proof  [32]byte
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSigDebug is a free log retrieval operation binding the contract event 0x6b7f08f909e42085ce2477fcf9f51be4ff86e5bcc7936c2dbe39e6757d33f914.
//
// Solidity: e SigDebug(_h bytes32, _proof bytes32, proof bytes32, signer address)
func (_Airdrop *AirdropFilterer) FilterSigDebug(opts *bind.FilterOpts) (*AirdropSigDebugIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "SigDebug")
	if err != nil {
		return nil, err
	}
	return &AirdropSigDebugIterator{contract: _Airdrop.contract, event: "SigDebug", logs: logs, sub: sub}, nil
}

// WatchSigDebug is a free log subscription operation binding the contract event 0x6b7f08f909e42085ce2477fcf9f51be4ff86e5bcc7936c2dbe39e6757d33f914.
//
// Solidity: e SigDebug(_h bytes32, _proof bytes32, proof bytes32, signer address)
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
// Solidity: e SignatureRecovered(signer indexed address)
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
// Solidity: e SignatureRecovered(signer indexed address)
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
