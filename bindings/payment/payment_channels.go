// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payment

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

// PaymentABI is the input ABI used to generate the binding from.
const PaymentABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_adminSet\",\"type\":\"bool\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ChannelExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recoveredAddress\",\"type\":\"address\"}],\"name\":\"DestinationProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"ErcChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"EthChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_remainingChannelValue\",\"type\":\"uint256\"}],\"name\":\"MicroPaymentWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_ownershipTransferred\",\"type\":\"bool\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recoveredAddress\",\"type\":\"address\"}],\"name\":\"SourceProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"closeErcChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultState\",\"outputs\":[{\"internalType\":\"enumPaymentChannels.ChannelStates\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ercChannels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openDate\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"sourceProofSubmitted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"destinationProofSubmitted\",\"type\":\"bool\"},{\"internalType\":\"enumPaymentChannels.ChannelStates\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ethChannels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openDate\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"sourceProofSubmitted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"destinationProofSubmitted\",\"type\":\"bool\"},{\"internalType\":\"enumPaymentChannels.ChannelStates\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"expireErcChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"moderators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_channelValueInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_durationInDays\",\"type\":\"uint256\"}],\"name\":\"openErcChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_channelValueInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_durationInDays\",\"type\":\"uint256\"}],\"name\":\"openEthChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"submitErcDestinationProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_paymentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submitErcMicroPaymentProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"submitErcSourceProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"submitEthDestinationProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_channelId\",\"type\":\"bytes32\"}],\"name\":\"submitEthSourceProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// PaymentBin is the compiled bytecode used for deploying new contracts.
var PaymentBin = "0x60c0604052601c60808190527f19457468657265756d205369676e6564204d6573736167653a0a33320000000060a090815262000040916003919062000089565b506004805460ff191660011761ff00191690553480156200006057600080fd5b5060008054336001600160a01b0319918216811790925560018054909116909117905562000125565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000cc57805160ff1916838001178555620000fc565b82800160010185558215620000fc579182015b82811115620000fc578251825591602001919060010190620000df565b506200010a9291506200010e565b5090565b5b808211156200010a57600081556001016200010f565b611c5c80620001356000396000f3fe6080604052600436106101395760003560e01c806391cca3db116100ab578063c21d92d81161006f578063c21d92d814610569578063f2fde38b146105ae578063f851a440146105e1578063f8be342e146105f6578063fcf0301d1461062c578063ffa1ad741461067157610140565b806391cca3db1461042f5780639ef3ae9d14610444578063a0ef91df14610489578063aa13a6121461049e578063b1653e67146104d057610140565b806349df728c116100fd57806349df728c146102c9578063704b6c02146102fc57806378bdd8391461032f5780637baccc6d146103745780638da5cb5b146103c55780638db53802146103f657610140565b8063054f7d9c1461014257806314d0f1ba1461016b5780632b2b4a751461019e5780633f73cfd7146101e75780633f93cc451461022057610140565b3661014057005b005b34801561014e57600080fd5b506101576106fb565b604080519115158252519081900360200190f35b34801561017757600080fd5b506101576004803603602081101561018e57600080fd5b50356001600160a01b031661070b565b3480156101aa57600080fd5b50610157600480360360808110156101c157600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135610720565b3480156101f357600080fd5b506101576004803603604081101561020a57600080fd5b50803590602001356001600160a01b03166108de565b34801561022c57600080fd5b5061024a6004803603602081101561024357600080fd5b5035610a86565b604051808b6001600160a01b031681526020018a6001600160a01b03168152602001896001600160a01b03168152602001888152602001878152602001868152602001858152602001841515815260200183151581526020018260038111156102af57fe5b81526020019a505050505050505050505060405180910390f35b3480156102d557600080fd5b50610157600480360360208110156102ec57600080fd5b50356001600160a01b0316610af0565b34801561030857600080fd5b506101576004803603602081101561031f57600080fd5b50356001600160a01b0316610c56565b34801561033b57600080fd5b50610157600480360360a081101561035257600080fd5b5080359060ff6020820135169060408101359060608101359060800135610cd4565b34801561038057600080fd5b50610157600480360360e081101561039757600080fd5b5080359060ff6020820135169060408101359060608101359060808101359060a08101359060c00135610e89565b3480156103d157600080fd5b506103da61115a565b604080516001600160a01b039092168252519081900360200190f35b34801561040257600080fd5b506101576004803603604081101561041957600080fd5b50803590602001356001600160a01b0316611169565b34801561043b57600080fd5b506101576112d2565b34801561045057600080fd5b50610157600480360360a081101561046757600080fd5b5080359060ff60208201351690604081013590606081013590608001356112db565b34801561049557600080fd5b506101576114a1565b610157600480360360608110156104b457600080fd5b506001600160a01b03813516906020810135906040013561153a565b3480156104dc57600080fd5b506104fa600480360360208110156104f357600080fd5b503561169b565b604051808a6001600160a01b03168152602001896001600160a01b031681526020018881526020018781526020018681526020018581526020018415158152602001831515815260200182600381111561055057fe5b8152602001995050505050505050505060405180910390f35b34801561057557600080fd5b50610157600480360360a081101561058c57600080fd5b5080359060ff60208201351690604081013590606081013590608001356116fc565b3480156105ba57600080fd5b50610157600480360360208110156105d157600080fd5b50356001600160a01b031661187f565b3480156105ed57600080fd5b506103da611900565b34801561060257600080fd5b5061060b61190f565b6040518082600381111561061b57fe5b815260200191505060405180910390f35b34801561063857600080fd5b50610157600480360360a081101561064f57600080fd5b5080359060ff602082013516906040810135906060810135906080013561191d565b34801561067d57600080fd5b50610686611ac1565b6040805160208082528351818301528351919283929083019185019080838360005b838110156106c05781810151838201526020016106a8565b50505050905090810190601f1680156106ed5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600154600160a01b900460ff1681565b60026020526000908152604090205460ff1681565b6040805133606090811b6020808401919091526bffffffffffffffffffffffff1987831b811660348501529188901b909116604883015242605c8084019190915283518084039091018152607c9092018352815191810191909120600081815260089092529181205490919060ff161561079957600080fd5b6000818152600860209081526040808320805460ff1916600190811790915560069283905281842080546001600160a01b03199081163317825591810180546001600160a01b038c811691851691909117909155600282018054918d169190931617909155600381018890554262015180880281016004830155600582015590910183905551879183917f41063c1edfbec39fdaea4c9b553c1fb2832454799895fcd189a36facbb7327629190a2604080516323b872dd60e01b81523360048201523060248201526044810187905290516001600160a01b038316916323b872dd9160648083019260209291908290030181600087803b15801561089c57600080fd5b505af11580156108b0573d6000803e3d6000fd5b505050506040513d60208110156108c657600080fd5b50516108d157600080fd5b5060019695505050505050565b60008281526008602052604081205460ff166108f957600080fd5b6000808481526006602052604090206007015462010000900460ff16600381111561092057fe5b1461092a57600080fd5b6000838152600660205260409020546001600160a01b0316331461094d57600080fd5b6000838152600660205260409020600201546001600160a01b0383811691161461097657600080fd5b600083815260066020526040812060038101805492905560070180546001919062ff0000191662010000830217905550604051839085907fb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe90600090a26040516001600160a01b03851690600080516020611c0783398151915290600090a26040805163a9059cbb60e01b81523360048201526024810184905290516001600160a01b0383169163a9059cbb9160448083019260209291908290030181600087803b158015610a4457600080fd5b505af1158015610a58573d6000803e3d6000fd5b505050506040513d6020811015610a6e57600080fd5b5051610a7957600080fd5b6001925050505b92915050565b6006602081905260009182526040909120805460018201546002830154600384015460048501546005860154968601546007909601546001600160a01b039586169794861696939095169491939092909160ff80821691610100810482169162010000909104168a565b600080546001600160a01b0316331480610b1457506001546001600160a01b031633145b610b1d57600080fd5b60045460ff16610b2c57600080fd5b604080516370a0823160e01b8152306004820152905183916000916001600160a01b038416916370a08231916024808301926020929190829003018186803b158015610b7757600080fd5b505afa158015610b8b573d6000803e3d6000fd5b505050506040513d6020811015610ba157600080fd5b50516040519091506001600160a01b03851690600080516020611c0783398151915290600090a26040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b158015610c1757600080fd5b505af1158015610c2b573d6000803e3d6000fd5b505050506040513d6020811015610c4157600080fd5b5051610c4c57600080fd5b5060019392505050565b600080546001600160a01b03163314610c6e57600080fd5b6001546001600160a01b0383811691161415610c8957600080fd5b600180546001600160a01b0319166001600160a01b03841690811782556040517fe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e990600090a3919050565b60008181526008602052604081205460ff16610cef57600080fd5b6000808381526006602052604090206007015462010000900460ff166003811115610d1657fe5b14610d2057600080fd5b60008281526006602052604090206007015460ff1615610d3f57600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610d9b573d6000803e3d6000fd5b505060408051601f1901516000868152600660205291909120549092506001600160a01b038084169116149050610dce57fe5b60008381526006602090815260408083206007018054600160ff199182168117909255600984528285208c86528452828520338652909352818420805490931617909155516001600160a01b0383169185917fef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a679190a3610e4f836001611ae7565b156108d157600083815260066020526040902060070180546002919062ff0000191662010000835b02179055505060019695505050505050565b60008381526008602052604081205460ff16610ea457600080fd5b6000808581526006602052604090206007015462010000900460ff166003811115610ecb57fe5b14610ed557600080fd5b600084815260066020526040902060030154610ef057600080fd5b6000848152600660205260409020600101546001600160a01b03163314610f1657600080fd5b600084848460405160200180848152602001838152602001828152602001935050505060405160208183030381529060405280519060200120905060006003826040516020018083805460018160011615610100020316600290048015610fb45780601f10610f92576101008083540402835291820191610fb4565b820191906000526020600020905b815481529060010190602001808311610fa0575b50509182525060408051808303815260209283019091528051910120915050898114610fdc57fe5b600060018b8b8b8b60405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611038573d6000803e3d6000fd5b505060408051601f19015160008a8152600660205291909120549092506001600160a01b03808416911614905061106b57fe5b6000878152600660205260408120600301546110879087611bf1565b600089815260066020526040808220600381018490556002015490519293506001600160a01b0316918291600080516020611c0783398151915291a26040805163a9059cbb60e01b81523360048201526024810189905290516001600160a01b0383169163a9059cbb9160448083019260209291908290030181600087803b15801561111257600080fd5b505af1158015611126573d6000803e3d6000fd5b505050506040513d602081101561113c57600080fd5b505161114757600080fd5b5060019c9b505050505050505050505050565b6000546001600160a01b031681565b60008281526008602052604081205460ff1661118457600080fd5b600260008481526006602052604090206007015462010000900460ff1660038111156111ac57fe5b146111b657600080fd5b6000838152600660205260409020600101546001600160a01b031633146111dc57600080fd5b6000838152600660205260409020600201546001600160a01b0383811691161461120557600080fd5b6000838152600660205260408120600380820180549390556007909101805462ff0000191662010000830217905550604051839085907fceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba7290600090a26040516001600160a01b03851690600080516020611c0783398151915290600090a26040805163a9059cbb60e01b81523360048201526024810184905290516001600160a01b0383169163a9059cbb9160448083019260209291908290030181600087803b158015610a4457600080fd5b60045460ff1681565b60008181526008602052604081205460ff166112f657600080fd5b6000808381526005602052604090206006015462010000900460ff16600381111561131d57fe5b1461132757600080fd5b60008281526005602052604090206006015460ff161561134657600080fd5b6000828152600960209081526040808320898452825280832033845290915290205460ff161561137557600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156113d1573d6000803e3d6000fd5b505060408051601f1901516000868152600560205291909120549092506001600160a01b03808416911614905061140457fe5b6000838152600560209081526040808320600601805460ff19166001179055600982528083208a84528252808320338452909152516001600160a01b0383169185917fef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a679190a3611475836000611ae7565b156108d157600083815260056020526040902060060180546002919062ff000019166201000083610e77565b600080546001600160a01b03163314806114c557506001546001600160a01b031633145b6114ce57600080fd5b60045460ff166114dd57600080fd5b6040517fe85db8c769971e615d7122e20c48e239d21cfd4b1a3561caf49010903e1ff5ac90600090a160405133904780156108fc02916000818181858888f19350505050158015611532573d6000803e3d6000fd5b506001905090565b600082341461154857600080fd5b6040805133606090811b6020808401919091529087901b6bffffffffffffffffffffffff19166034830152604882018690524260688084018290528451808503909101815260889093018452825192820192909220600081815260089092529290205490919060ff16156115bb57600080fd5b6000818152600860209081526040808320805460ff19166001908117909155600592839052922080546001600160a01b03199081163317825592810180546001600160a01b038b16941693909317909255600282018790554262015180870201600380840191909155600480840186905591830184905590546006909201805461010090930460ff1692909162ff000019909116906201000090849081111561166057fe5b021790555060405181907ec725122f067c5ca4dfb33fb90955e8c8ace1fbba4578cd42779dd95ed21faa90600090a250600195945050505050565b600560208190526000918252604090912080546001820154600283015460038401546004850154958501546006909501546001600160a01b0394851696939094169491939092919060ff808216916101008104821691620100009091041689565b60008181526008602052604081205460ff1661171757600080fd5b6000808381526006602052604090206007015462010000900460ff16600381111561173e57fe5b1461174857600080fd5b600082815260066020526040902060070154610100900460ff161561176c57600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156117c8573d6000803e3d6000fd5b505060408051601f1901516000868152600660205291909120600101549092506001600160a01b0380841691161490506117fe57fe5b6000838152600660209081526040808320600701805461ff001916610100179055600982528083208a84528252808320338452909152808220805460ff19166001179055516001600160a01b0383169185917f4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac9190a3610e4f836001611ae7565b600080546001600160a01b0316331461189757600080fd5b6000546001600160a01b03838116911614156118b257600080fd5b600080546001600160a01b0319166001600160a01b038416908117825560405160019233917f7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa589190a4919050565b6001546001600160a01b031681565b600454610100900460ff1681565b60008181526008602052604081205460ff1661193857600080fd5b6000808381526005602052604090206006015462010000900460ff16600381111561195f57fe5b1461196957600080fd5b600082815260056020526040902060060154610100900460ff161561198d57600080fd5b6000828152600960209081526040808320898452825280832033845290915290205460ff16156119bc57600080fd5b600060018787878760405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611a18573d6000803e3d6000fd5b505060408051601f1901516000868152600560205291909120600101549092506001600160a01b038084169116149050611a4e57fe5b6000838152600560209081526040808320600601805461ff001916610100179055600982528083208a84528252808320338452909152516001600160a01b0383169185917f4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac9190a3611475836000611ae7565b6040518060400160405280600a815260200169302e342e33616c70686160b01b81525081565b60008115611b6e576000808481526006602052604090206007015462010000900460ff166003811115611b1657fe5b14611b2057600080fd5b60008381526006602052604090206007015460ff1615156001148015611b5c5750600083815260066020526040902060070154610100900460ff165b15611b6957506001610a80565b611be8565b6000808481526005602052604090206006015462010000900460ff166003811115611b9557fe5b14611b9f57600080fd5b60008381526005602052604090206006015460ff1615156001148015611bdb5750600083815260056020526040902060060154610100900460ff165b15611be857506001610a80565b50600092915050565b600082821115611c0057600080fd5b5090039056feb47e37d80bb9227b793caed3e80256b8eac5e891b85d33f583d43ce03cf34273a264697066735822122028aef5485453317b6232bcd4a8222579f28cc077cc99c81aa482d95b4c7e574964736f6c63430007000033"

// DeployPayment deploys a new Ethereum contract, binding an instance of Payment to it.
func DeployPayment(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Payment, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaymentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Payment{PaymentCaller: PaymentCaller{contract: contract}, PaymentTransactor: PaymentTransactor{contract: contract}, PaymentFilterer: PaymentFilterer{contract: contract}}, nil
}

// Payment is an auto generated Go binding around an Ethereum contract.
type Payment struct {
	PaymentCaller     // Read-only binding to the contract
	PaymentTransactor // Write-only binding to the contract
	PaymentFilterer   // Log filterer for contract events
}

// PaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentSession struct {
	Contract     *Payment          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentCallerSession struct {
	Contract *PaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentTransactorSession struct {
	Contract     *PaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentRaw struct {
	Contract *Payment // Generic contract binding to access the raw methods on
}

// PaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentCallerRaw struct {
	Contract *PaymentCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentTransactorRaw struct {
	Contract *PaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayment creates a new instance of Payment, bound to a specific deployed contract.
func NewPayment(address common.Address, backend bind.ContractBackend) (*Payment, error) {
	contract, err := bindPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payment{PaymentCaller: PaymentCaller{contract: contract}, PaymentTransactor: PaymentTransactor{contract: contract}, PaymentFilterer: PaymentFilterer{contract: contract}}, nil
}

// NewPaymentCaller creates a new read-only instance of Payment, bound to a specific deployed contract.
func NewPaymentCaller(address common.Address, caller bind.ContractCaller) (*PaymentCaller, error) {
	contract, err := bindPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCaller{contract: contract}, nil
}

// NewPaymentTransactor creates a new write-only instance of Payment, bound to a specific deployed contract.
func NewPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentTransactor, error) {
	contract, err := bindPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentTransactor{contract: contract}, nil
}

// NewPaymentFilterer creates a new log filterer instance of Payment, bound to a specific deployed contract.
func NewPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentFilterer, error) {
	contract, err := bindPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentFilterer{contract: contract}, nil
}

// bindPayment binds a generic wrapper to an already deployed contract.
func bindPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.PaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Payment *PaymentCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Payment *PaymentSession) VERSION() (string, error) {
	return _Payment.Contract.VERSION(&_Payment.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Payment *PaymentCallerSession) VERSION() (string, error) {
	return _Payment.Contract.VERSION(&_Payment.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Payment *PaymentCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Payment *PaymentSession) Admin() (common.Address, error) {
	return _Payment.Contract.Admin(&_Payment.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Payment *PaymentCallerSession) Admin() (common.Address, error) {
	return _Payment.Contract.Admin(&_Payment.CallOpts)
}

// DefaultState is a free data retrieval call binding the contract method 0xf8be342e.
//
// Solidity: function defaultState() view returns(uint8)
func (_Payment *PaymentCaller) DefaultState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "defaultState")
	return *ret0, err
}

// DefaultState is a free data retrieval call binding the contract method 0xf8be342e.
//
// Solidity: function defaultState() view returns(uint8)
func (_Payment *PaymentSession) DefaultState() (uint8, error) {
	return _Payment.Contract.DefaultState(&_Payment.CallOpts)
}

// DefaultState is a free data retrieval call binding the contract method 0xf8be342e.
//
// Solidity: function defaultState() view returns(uint8)
func (_Payment *PaymentCallerSession) DefaultState() (uint8, error) {
	return _Payment.Contract.DefaultState(&_Payment.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Payment *PaymentCaller) Dev(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "dev")
	return *ret0, err
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Payment *PaymentSession) Dev() (bool, error) {
	return _Payment.Contract.Dev(&_Payment.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(bool)
func (_Payment *PaymentCallerSession) Dev() (bool, error) {
	return _Payment.Contract.Dev(&_Payment.CallOpts)
}

// ErcChannels is a free data retrieval call binding the contract method 0x3f93cc45.
//
// Solidity: function ercChannels(bytes32 ) view returns(address source, address destination, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Payment *PaymentCaller) ErcChannels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	TokenAddress              common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	ret := new(struct {
		Source                    common.Address
		Destination               common.Address
		TokenAddress              common.Address
		Value                     *big.Int
		ClosingDate               *big.Int
		OpenDate                  *big.Int
		ChannelId                 [32]byte
		SourceProofSubmitted      bool
		DestinationProofSubmitted bool
		State                     uint8
	})
	out := ret
	err := _Payment.contract.Call(opts, out, "ercChannels", arg0)
	return *ret, err
}

// ErcChannels is a free data retrieval call binding the contract method 0x3f93cc45.
//
// Solidity: function ercChannels(bytes32 ) view returns(address source, address destination, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Payment *PaymentSession) ErcChannels(arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	TokenAddress              common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	return _Payment.Contract.ErcChannels(&_Payment.CallOpts, arg0)
}

// ErcChannels is a free data retrieval call binding the contract method 0x3f93cc45.
//
// Solidity: function ercChannels(bytes32 ) view returns(address source, address destination, address tokenAddress, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Payment *PaymentCallerSession) ErcChannels(arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	TokenAddress              common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	return _Payment.Contract.ErcChannels(&_Payment.CallOpts, arg0)
}

// EthChannels is a free data retrieval call binding the contract method 0xb1653e67.
//
// Solidity: function ethChannels(bytes32 ) view returns(address source, address destination, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Payment *PaymentCaller) EthChannels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	ret := new(struct {
		Source                    common.Address
		Destination               common.Address
		Value                     *big.Int
		ClosingDate               *big.Int
		OpenDate                  *big.Int
		ChannelId                 [32]byte
		SourceProofSubmitted      bool
		DestinationProofSubmitted bool
		State                     uint8
	})
	out := ret
	err := _Payment.contract.Call(opts, out, "ethChannels", arg0)
	return *ret, err
}

// EthChannels is a free data retrieval call binding the contract method 0xb1653e67.
//
// Solidity: function ethChannels(bytes32 ) view returns(address source, address destination, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Payment *PaymentSession) EthChannels(arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	return _Payment.Contract.EthChannels(&_Payment.CallOpts, arg0)
}

// EthChannels is a free data retrieval call binding the contract method 0xb1653e67.
//
// Solidity: function ethChannels(bytes32 ) view returns(address source, address destination, uint256 value, uint256 closingDate, uint256 openDate, bytes32 channelId, bool sourceProofSubmitted, bool destinationProofSubmitted, uint8 state)
func (_Payment *PaymentCallerSession) EthChannels(arg0 [32]byte) (struct {
	Source                    common.Address
	Destination               common.Address
	Value                     *big.Int
	ClosingDate               *big.Int
	OpenDate                  *big.Int
	ChannelId                 [32]byte
	SourceProofSubmitted      bool
	DestinationProofSubmitted bool
	State                     uint8
}, error) {
	return _Payment.Contract.EthChannels(&_Payment.CallOpts, arg0)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Payment *PaymentCaller) Frozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "frozen")
	return *ret0, err
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Payment *PaymentSession) Frozen() (bool, error) {
	return _Payment.Contract.Frozen(&_Payment.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_Payment *PaymentCallerSession) Frozen() (bool, error) {
	return _Payment.Contract.Frozen(&_Payment.CallOpts)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Payment *PaymentCaller) Moderators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "moderators", arg0)
	return *ret0, err
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Payment *PaymentSession) Moderators(arg0 common.Address) (bool, error) {
	return _Payment.Contract.Moderators(&_Payment.CallOpts, arg0)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_Payment *PaymentCallerSession) Moderators(arg0 common.Address) (bool, error) {
	return _Payment.Contract.Moderators(&_Payment.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentSession) Owner() (common.Address, error) {
	return _Payment.Contract.Owner(&_Payment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentCallerSession) Owner() (common.Address, error) {
	return _Payment.Contract.Owner(&_Payment.CallOpts)
}

// CloseErcChannel is a paid mutator transaction binding the contract method 0x8db53802.
//
// Solidity: function closeErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Payment *PaymentTransactor) CloseErcChannel(opts *bind.TransactOpts, _channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "closeErcChannel", _channelId, _tokenAddress)
}

// CloseErcChannel is a paid mutator transaction binding the contract method 0x8db53802.
//
// Solidity: function closeErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Payment *PaymentSession) CloseErcChannel(_channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.Contract.CloseErcChannel(&_Payment.TransactOpts, _channelId, _tokenAddress)
}

// CloseErcChannel is a paid mutator transaction binding the contract method 0x8db53802.
//
// Solidity: function closeErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Payment *PaymentTransactorSession) CloseErcChannel(_channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.Contract.CloseErcChannel(&_Payment.TransactOpts, _channelId, _tokenAddress)
}

// ExpireErcChannel is a paid mutator transaction binding the contract method 0x3f73cfd7.
//
// Solidity: function expireErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Payment *PaymentTransactor) ExpireErcChannel(opts *bind.TransactOpts, _channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "expireErcChannel", _channelId, _tokenAddress)
}

// ExpireErcChannel is a paid mutator transaction binding the contract method 0x3f73cfd7.
//
// Solidity: function expireErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Payment *PaymentSession) ExpireErcChannel(_channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.Contract.ExpireErcChannel(&_Payment.TransactOpts, _channelId, _tokenAddress)
}

// ExpireErcChannel is a paid mutator transaction binding the contract method 0x3f73cfd7.
//
// Solidity: function expireErcChannel(bytes32 _channelId, address _tokenAddress) returns(bool)
func (_Payment *PaymentTransactorSession) ExpireErcChannel(_channelId [32]byte, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.Contract.ExpireErcChannel(&_Payment.TransactOpts, _channelId, _tokenAddress)
}

// OpenErcChannel is a paid mutator transaction binding the contract method 0x2b2b4a75.
//
// Solidity: function openErcChannel(address _tokenAddress, address _destination, uint256 _channelValueInWei, uint256 _durationInDays) returns(bool)
func (_Payment *PaymentTransactor) OpenErcChannel(opts *bind.TransactOpts, _tokenAddress common.Address, _destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "openErcChannel", _tokenAddress, _destination, _channelValueInWei, _durationInDays)
}

// OpenErcChannel is a paid mutator transaction binding the contract method 0x2b2b4a75.
//
// Solidity: function openErcChannel(address _tokenAddress, address _destination, uint256 _channelValueInWei, uint256 _durationInDays) returns(bool)
func (_Payment *PaymentSession) OpenErcChannel(_tokenAddress common.Address, _destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.OpenErcChannel(&_Payment.TransactOpts, _tokenAddress, _destination, _channelValueInWei, _durationInDays)
}

// OpenErcChannel is a paid mutator transaction binding the contract method 0x2b2b4a75.
//
// Solidity: function openErcChannel(address _tokenAddress, address _destination, uint256 _channelValueInWei, uint256 _durationInDays) returns(bool)
func (_Payment *PaymentTransactorSession) OpenErcChannel(_tokenAddress common.Address, _destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.OpenErcChannel(&_Payment.TransactOpts, _tokenAddress, _destination, _channelValueInWei, _durationInDays)
}

// OpenEthChannel is a paid mutator transaction binding the contract method 0xaa13a612.
//
// Solidity: function openEthChannel(address _destination, uint256 _channelValueInWei, uint256 _durationInDays) payable returns(bool)
func (_Payment *PaymentTransactor) OpenEthChannel(opts *bind.TransactOpts, _destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "openEthChannel", _destination, _channelValueInWei, _durationInDays)
}

// OpenEthChannel is a paid mutator transaction binding the contract method 0xaa13a612.
//
// Solidity: function openEthChannel(address _destination, uint256 _channelValueInWei, uint256 _durationInDays) payable returns(bool)
func (_Payment *PaymentSession) OpenEthChannel(_destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.OpenEthChannel(&_Payment.TransactOpts, _destination, _channelValueInWei, _durationInDays)
}

// OpenEthChannel is a paid mutator transaction binding the contract method 0xaa13a612.
//
// Solidity: function openEthChannel(address _destination, uint256 _channelValueInWei, uint256 _durationInDays) payable returns(bool)
func (_Payment *PaymentTransactorSession) OpenEthChannel(_destination common.Address, _channelValueInWei *big.Int, _durationInDays *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.OpenEthChannel(&_Payment.TransactOpts, _destination, _channelValueInWei, _durationInDays)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Payment *PaymentTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Payment *PaymentSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetAdmin(&_Payment.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Payment *PaymentTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetAdmin(&_Payment.TransactOpts, _newAdmin)
}

// SubmitErcDestinationProof is a paid mutator transaction binding the contract method 0xc21d92d8.
//
// Solidity: function submitErcDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentTransactor) SubmitErcDestinationProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "submitErcDestinationProof", _h, _v, _r, _s, _channelId)
}

// SubmitErcDestinationProof is a paid mutator transaction binding the contract method 0xc21d92d8.
//
// Solidity: function submitErcDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentSession) SubmitErcDestinationProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.Contract.SubmitErcDestinationProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitErcDestinationProof is a paid mutator transaction binding the contract method 0xc21d92d8.
//
// Solidity: function submitErcDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentTransactorSession) SubmitErcDestinationProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.Contract.SubmitErcDestinationProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitErcMicroPaymentProof is a paid mutator transaction binding the contract method 0x7baccc6d.
//
// Solidity: function submitErcMicroPaymentProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _paymentId, uint256 _amount) returns(bool)
func (_Payment *PaymentTransactor) SubmitErcMicroPaymentProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _paymentId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "submitErcMicroPaymentProof", _h, _v, _r, _s, _channelId, _paymentId, _amount)
}

// SubmitErcMicroPaymentProof is a paid mutator transaction binding the contract method 0x7baccc6d.
//
// Solidity: function submitErcMicroPaymentProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _paymentId, uint256 _amount) returns(bool)
func (_Payment *PaymentSession) SubmitErcMicroPaymentProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _paymentId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.SubmitErcMicroPaymentProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId, _paymentId, _amount)
}

// SubmitErcMicroPaymentProof is a paid mutator transaction binding the contract method 0x7baccc6d.
//
// Solidity: function submitErcMicroPaymentProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId, uint256 _paymentId, uint256 _amount) returns(bool)
func (_Payment *PaymentTransactorSession) SubmitErcMicroPaymentProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte, _paymentId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.SubmitErcMicroPaymentProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId, _paymentId, _amount)
}

// SubmitErcSourceProof is a paid mutator transaction binding the contract method 0x78bdd839.
//
// Solidity: function submitErcSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentTransactor) SubmitErcSourceProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "submitErcSourceProof", _h, _v, _r, _s, _channelId)
}

// SubmitErcSourceProof is a paid mutator transaction binding the contract method 0x78bdd839.
//
// Solidity: function submitErcSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentSession) SubmitErcSourceProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.Contract.SubmitErcSourceProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitErcSourceProof is a paid mutator transaction binding the contract method 0x78bdd839.
//
// Solidity: function submitErcSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentTransactorSession) SubmitErcSourceProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.Contract.SubmitErcSourceProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitEthDestinationProof is a paid mutator transaction binding the contract method 0xfcf0301d.
//
// Solidity: function submitEthDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentTransactor) SubmitEthDestinationProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "submitEthDestinationProof", _h, _v, _r, _s, _channelId)
}

// SubmitEthDestinationProof is a paid mutator transaction binding the contract method 0xfcf0301d.
//
// Solidity: function submitEthDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentSession) SubmitEthDestinationProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.Contract.SubmitEthDestinationProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitEthDestinationProof is a paid mutator transaction binding the contract method 0xfcf0301d.
//
// Solidity: function submitEthDestinationProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentTransactorSession) SubmitEthDestinationProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.Contract.SubmitEthDestinationProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitEthSourceProof is a paid mutator transaction binding the contract method 0x9ef3ae9d.
//
// Solidity: function submitEthSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentTransactor) SubmitEthSourceProof(opts *bind.TransactOpts, _h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "submitEthSourceProof", _h, _v, _r, _s, _channelId)
}

// SubmitEthSourceProof is a paid mutator transaction binding the contract method 0x9ef3ae9d.
//
// Solidity: function submitEthSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentSession) SubmitEthSourceProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.Contract.SubmitEthSourceProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId)
}

// SubmitEthSourceProof is a paid mutator transaction binding the contract method 0x9ef3ae9d.
//
// Solidity: function submitEthSourceProof(bytes32 _h, uint8 _v, bytes32 _r, bytes32 _s, bytes32 _channelId) returns(bool)
func (_Payment *PaymentTransactorSession) SubmitEthSourceProof(_h [32]byte, _v uint8, _r [32]byte, _s [32]byte, _channelId [32]byte) (*types.Transaction, error) {
	return _Payment.Contract.SubmitEthSourceProof(&_Payment.TransactOpts, _h, _v, _r, _s, _channelId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Payment *PaymentTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Payment *PaymentSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Payment.Contract.TransferOwnership(&_Payment.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Payment *PaymentTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Payment.Contract.TransferOwnership(&_Payment.TransactOpts, _newOwner)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Payment *PaymentTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Payment *PaymentSession) WithdrawEth() (*types.Transaction, error) {
	return _Payment.Contract.WithdrawEth(&_Payment.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns(bool)
func (_Payment *PaymentTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _Payment.Contract.WithdrawEth(&_Payment.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address _tokenAddress) returns(bool)
func (_Payment *PaymentTransactor) WithdrawTokens(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "withdrawTokens", _tokenAddress)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address _tokenAddress) returns(bool)
func (_Payment *PaymentSession) WithdrawTokens(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.Contract.WithdrawTokens(&_Payment.TransactOpts, _tokenAddress)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address _tokenAddress) returns(bool)
func (_Payment *PaymentTransactorSession) WithdrawTokens(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payment.Contract.WithdrawTokens(&_Payment.TransactOpts, _tokenAddress)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Payment *PaymentTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Payment.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Payment *PaymentSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Payment.Contract.Fallback(&_Payment.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Payment *PaymentTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Payment.Contract.Fallback(&_Payment.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Payment *PaymentTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Payment *PaymentSession) Receive() (*types.Transaction, error) {
	return _Payment.Contract.Receive(&_Payment.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Payment *PaymentTransactorSession) Receive() (*types.Transaction, error) {
	return _Payment.Contract.Receive(&_Payment.TransactOpts)
}

// PaymentAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the Payment contract.
type PaymentAdminSetIterator struct {
	Event *PaymentAdminSet // Event containing the contract specifics and raw log

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
func (it *PaymentAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentAdminSet)
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
		it.Event = new(PaymentAdminSet)
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
func (it *PaymentAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentAdminSet represents a AdminSet event raised by the Payment contract.
type PaymentAdminSet struct {
	Admin    common.Address
	AdminSet bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(address indexed _admin, bool indexed _adminSet)
func (_Payment *PaymentFilterer) FilterAdminSet(opts *bind.FilterOpts, _admin []common.Address, _adminSet []bool) (*PaymentAdminSetIterator, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return &PaymentAdminSetIterator{contract: _Payment.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xe68d2c359a771606c400cf8b87000cf5864010363d6a736e98f5047b7bbe18e9.
//
// Solidity: event AdminSet(address indexed _admin, bool indexed _adminSet)
func (_Payment *PaymentFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *PaymentAdminSet, _admin []common.Address, _adminSet []bool) (event.Subscription, error) {

	var _adminRule []interface{}
	for _, _adminItem := range _admin {
		_adminRule = append(_adminRule, _adminItem)
	}
	var _adminSetRule []interface{}
	for _, _adminSetItem := range _adminSet {
		_adminSetRule = append(_adminSetRule, _adminSetItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "AdminSet", _adminRule, _adminSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentAdminSet)
				if err := _Payment.contract.UnpackLog(event, "AdminSet", log); err != nil {
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
func (_Payment *PaymentFilterer) ParseAdminSet(log types.Log) (*PaymentAdminSet, error) {
	event := new(PaymentAdminSet)
	if err := _Payment.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentChannelClosedIterator is returned from FilterChannelClosed and is used to iterate over the raw logs and unpacked data for ChannelClosed events raised by the Payment contract.
type PaymentChannelClosedIterator struct {
	Event *PaymentChannelClosed // Event containing the contract specifics and raw log

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
func (it *PaymentChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentChannelClosed)
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
		it.Event = new(PaymentChannelClosed)
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
func (it *PaymentChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentChannelClosed represents a ChannelClosed event raised by the Payment contract.
type PaymentChannelClosed struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelClosed is a free log retrieval operation binding the contract event 0xceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba72.
//
// Solidity: event ChannelClosed(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) FilterChannelClosed(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentChannelClosedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "ChannelClosed", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentChannelClosedIterator{contract: _Payment.contract, event: "ChannelClosed", logs: logs, sub: sub}, nil
}

// WatchChannelClosed is a free log subscription operation binding the contract event 0xceeab2eef998c17fe96f30f83fbf3c55fc5047f6e40c55a0cf72d236e9d2ba72.
//
// Solidity: event ChannelClosed(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) WatchChannelClosed(opts *bind.WatchOpts, sink chan<- *PaymentChannelClosed, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "ChannelClosed", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentChannelClosed)
				if err := _Payment.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
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
func (_Payment *PaymentFilterer) ParseChannelClosed(log types.Log) (*PaymentChannelClosed, error) {
	event := new(PaymentChannelClosed)
	if err := _Payment.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentChannelExpiredIterator is returned from FilterChannelExpired and is used to iterate over the raw logs and unpacked data for ChannelExpired events raised by the Payment contract.
type PaymentChannelExpiredIterator struct {
	Event *PaymentChannelExpired // Event containing the contract specifics and raw log

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
func (it *PaymentChannelExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentChannelExpired)
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
		it.Event = new(PaymentChannelExpired)
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
func (it *PaymentChannelExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentChannelExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentChannelExpired represents a ChannelExpired event raised by the Payment contract.
type PaymentChannelExpired struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelExpired is a free log retrieval operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) FilterChannelExpired(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentChannelExpiredIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "ChannelExpired", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentChannelExpiredIterator{contract: _Payment.contract, event: "ChannelExpired", logs: logs, sub: sub}, nil
}

// WatchChannelExpired is a free log subscription operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) WatchChannelExpired(opts *bind.WatchOpts, sink chan<- *PaymentChannelExpired, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "ChannelExpired", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentChannelExpired)
				if err := _Payment.contract.UnpackLog(event, "ChannelExpired", log); err != nil {
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

// ParseChannelExpired is a log parse operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) ParseChannelExpired(log types.Log) (*PaymentChannelExpired, error) {
	event := new(PaymentChannelExpired)
	if err := _Payment.contract.UnpackLog(event, "ChannelExpired", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentDestinationProofSubmittedIterator is returned from FilterDestinationProofSubmitted and is used to iterate over the raw logs and unpacked data for DestinationProofSubmitted events raised by the Payment contract.
type PaymentDestinationProofSubmittedIterator struct {
	Event *PaymentDestinationProofSubmitted // Event containing the contract specifics and raw log

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
func (it *PaymentDestinationProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentDestinationProofSubmitted)
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
		it.Event = new(PaymentDestinationProofSubmitted)
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
func (it *PaymentDestinationProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentDestinationProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentDestinationProofSubmitted represents a DestinationProofSubmitted event raised by the Payment contract.
type PaymentDestinationProofSubmitted struct {
	ChannelId        [32]byte
	RecoveredAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDestinationProofSubmitted is a free log retrieval operation binding the contract event 0x4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac.
//
// Solidity: event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Payment *PaymentFilterer) FilterDestinationProofSubmitted(opts *bind.FilterOpts, _channelId [][32]byte, _recoveredAddress []common.Address) (*PaymentDestinationProofSubmittedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}
	var _recoveredAddressRule []interface{}
	for _, _recoveredAddressItem := range _recoveredAddress {
		_recoveredAddressRule = append(_recoveredAddressRule, _recoveredAddressItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "DestinationProofSubmitted", _channelIdRule, _recoveredAddressRule)
	if err != nil {
		return nil, err
	}
	return &PaymentDestinationProofSubmittedIterator{contract: _Payment.contract, event: "DestinationProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchDestinationProofSubmitted is a free log subscription operation binding the contract event 0x4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac.
//
// Solidity: event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Payment *PaymentFilterer) WatchDestinationProofSubmitted(opts *bind.WatchOpts, sink chan<- *PaymentDestinationProofSubmitted, _channelId [][32]byte, _recoveredAddress []common.Address) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}
	var _recoveredAddressRule []interface{}
	for _, _recoveredAddressItem := range _recoveredAddress {
		_recoveredAddressRule = append(_recoveredAddressRule, _recoveredAddressItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "DestinationProofSubmitted", _channelIdRule, _recoveredAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentDestinationProofSubmitted)
				if err := _Payment.contract.UnpackLog(event, "DestinationProofSubmitted", log); err != nil {
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

// ParseDestinationProofSubmitted is a log parse operation binding the contract event 0x4bfcb0e4fde38226ea7f6e76bc336e1aa934fec6abce5c6a98826c226c27bcac.
//
// Solidity: event DestinationProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Payment *PaymentFilterer) ParseDestinationProofSubmitted(log types.Log) (*PaymentDestinationProofSubmitted, error) {
	event := new(PaymentDestinationProofSubmitted)
	if err := _Payment.contract.UnpackLog(event, "DestinationProofSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentErcChannelOpenedIterator is returned from FilterErcChannelOpened and is used to iterate over the raw logs and unpacked data for ErcChannelOpened events raised by the Payment contract.
type PaymentErcChannelOpenedIterator struct {
	Event *PaymentErcChannelOpened // Event containing the contract specifics and raw log

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
func (it *PaymentErcChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentErcChannelOpened)
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
		it.Event = new(PaymentErcChannelOpened)
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
func (it *PaymentErcChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentErcChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentErcChannelOpened represents a ErcChannelOpened event raised by the Payment contract.
type PaymentErcChannelOpened struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterErcChannelOpened is a free log retrieval operation binding the contract event 0x41063c1edfbec39fdaea4c9b553c1fb2832454799895fcd189a36facbb732762.
//
// Solidity: event ErcChannelOpened(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) FilterErcChannelOpened(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentErcChannelOpenedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "ErcChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentErcChannelOpenedIterator{contract: _Payment.contract, event: "ErcChannelOpened", logs: logs, sub: sub}, nil
}

// WatchErcChannelOpened is a free log subscription operation binding the contract event 0x41063c1edfbec39fdaea4c9b553c1fb2832454799895fcd189a36facbb732762.
//
// Solidity: event ErcChannelOpened(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) WatchErcChannelOpened(opts *bind.WatchOpts, sink chan<- *PaymentErcChannelOpened, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "ErcChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentErcChannelOpened)
				if err := _Payment.contract.UnpackLog(event, "ErcChannelOpened", log); err != nil {
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

// ParseErcChannelOpened is a log parse operation binding the contract event 0x41063c1edfbec39fdaea4c9b553c1fb2832454799895fcd189a36facbb732762.
//
// Solidity: event ErcChannelOpened(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) ParseErcChannelOpened(log types.Log) (*PaymentErcChannelOpened, error) {
	event := new(PaymentErcChannelOpened)
	if err := _Payment.contract.UnpackLog(event, "ErcChannelOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentEthChannelOpenedIterator is returned from FilterEthChannelOpened and is used to iterate over the raw logs and unpacked data for EthChannelOpened events raised by the Payment contract.
type PaymentEthChannelOpenedIterator struct {
	Event *PaymentEthChannelOpened // Event containing the contract specifics and raw log

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
func (it *PaymentEthChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentEthChannelOpened)
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
		it.Event = new(PaymentEthChannelOpened)
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
func (it *PaymentEthChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentEthChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentEthChannelOpened represents a EthChannelOpened event raised by the Payment contract.
type PaymentEthChannelOpened struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthChannelOpened is a free log retrieval operation binding the contract event 0x00c725122f067c5ca4dfb33fb90955e8c8ace1fbba4578cd42779dd95ed21faa.
//
// Solidity: event EthChannelOpened(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) FilterEthChannelOpened(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentEthChannelOpenedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "EthChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentEthChannelOpenedIterator{contract: _Payment.contract, event: "EthChannelOpened", logs: logs, sub: sub}, nil
}

// WatchEthChannelOpened is a free log subscription operation binding the contract event 0x00c725122f067c5ca4dfb33fb90955e8c8ace1fbba4578cd42779dd95ed21faa.
//
// Solidity: event EthChannelOpened(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) WatchEthChannelOpened(opts *bind.WatchOpts, sink chan<- *PaymentEthChannelOpened, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "EthChannelOpened", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentEthChannelOpened)
				if err := _Payment.contract.UnpackLog(event, "EthChannelOpened", log); err != nil {
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

// ParseEthChannelOpened is a log parse operation binding the contract event 0x00c725122f067c5ca4dfb33fb90955e8c8ace1fbba4578cd42779dd95ed21faa.
//
// Solidity: event EthChannelOpened(bytes32 indexed _channelId)
func (_Payment *PaymentFilterer) ParseEthChannelOpened(log types.Log) (*PaymentEthChannelOpened, error) {
	event := new(PaymentEthChannelOpened)
	if err := _Payment.contract.UnpackLog(event, "EthChannelOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the Payment contract.
type PaymentEthWithdrawnIterator struct {
	Event *PaymentEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentEthWithdrawn)
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
		it.Event = new(PaymentEthWithdrawn)
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
func (it *PaymentEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentEthWithdrawn represents a EthWithdrawn event raised by the Payment contract.
type PaymentEthWithdrawn struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0xe85db8c769971e615d7122e20c48e239d21cfd4b1a3561caf49010903e1ff5ac.
//
// Solidity: event EthWithdrawn()
func (_Payment *PaymentFilterer) FilterEthWithdrawn(opts *bind.FilterOpts) (*PaymentEthWithdrawnIterator, error) {

	logs, sub, err := _Payment.contract.FilterLogs(opts, "EthWithdrawn")
	if err != nil {
		return nil, err
	}
	return &PaymentEthWithdrawnIterator{contract: _Payment.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0xe85db8c769971e615d7122e20c48e239d21cfd4b1a3561caf49010903e1ff5ac.
//
// Solidity: event EthWithdrawn()
func (_Payment *PaymentFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentEthWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Payment.contract.WatchLogs(opts, "EthWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentEthWithdrawn)
				if err := _Payment.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0xe85db8c769971e615d7122e20c48e239d21cfd4b1a3561caf49010903e1ff5ac.
//
// Solidity: event EthWithdrawn()
func (_Payment *PaymentFilterer) ParseEthWithdrawn(log types.Log) (*PaymentEthWithdrawn, error) {
	event := new(PaymentEthWithdrawn)
	if err := _Payment.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentMicroPaymentWithdrawnIterator is returned from FilterMicroPaymentWithdrawn and is used to iterate over the raw logs and unpacked data for MicroPaymentWithdrawn events raised by the Payment contract.
type PaymentMicroPaymentWithdrawnIterator struct {
	Event *PaymentMicroPaymentWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentMicroPaymentWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentMicroPaymentWithdrawn)
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
		it.Event = new(PaymentMicroPaymentWithdrawn)
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
func (it *PaymentMicroPaymentWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentMicroPaymentWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentMicroPaymentWithdrawn represents a MicroPaymentWithdrawn event raised by the Payment contract.
type PaymentMicroPaymentWithdrawn struct {
	ChannelId             [32]byte
	Amount                *big.Int
	RemainingChannelValue *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMicroPaymentWithdrawn is a free log retrieval operation binding the contract event 0x9b992a6dd822a48e6743018b4b1c20d242672ad05d1e485d2c43ed954fb0ff15.
//
// Solidity: event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue)
func (_Payment *PaymentFilterer) FilterMicroPaymentWithdrawn(opts *bind.FilterOpts, _channelId [][32]byte) (*PaymentMicroPaymentWithdrawnIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "MicroPaymentWithdrawn", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentMicroPaymentWithdrawnIterator{contract: _Payment.contract, event: "MicroPaymentWithdrawn", logs: logs, sub: sub}, nil
}

// WatchMicroPaymentWithdrawn is a free log subscription operation binding the contract event 0x9b992a6dd822a48e6743018b4b1c20d242672ad05d1e485d2c43ed954fb0ff15.
//
// Solidity: event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue)
func (_Payment *PaymentFilterer) WatchMicroPaymentWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentMicroPaymentWithdrawn, _channelId [][32]byte) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "MicroPaymentWithdrawn", _channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentMicroPaymentWithdrawn)
				if err := _Payment.contract.UnpackLog(event, "MicroPaymentWithdrawn", log); err != nil {
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

// ParseMicroPaymentWithdrawn is a log parse operation binding the contract event 0x9b992a6dd822a48e6743018b4b1c20d242672ad05d1e485d2c43ed954fb0ff15.
//
// Solidity: event MicroPaymentWithdrawn(bytes32 indexed _channelId, uint256 _amount, uint256 _remainingChannelValue)
func (_Payment *PaymentFilterer) ParseMicroPaymentWithdrawn(log types.Log) (*PaymentMicroPaymentWithdrawn, error) {
	event := new(PaymentMicroPaymentWithdrawn)
	if err := _Payment.contract.UnpackLog(event, "MicroPaymentWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Payment contract.
type PaymentOwnershipTransferredIterator struct {
	Event *PaymentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentOwnershipTransferred)
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
		it.Event = new(PaymentOwnershipTransferred)
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
func (it *PaymentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentOwnershipTransferred represents a OwnershipTransferred event raised by the Payment contract.
type PaymentOwnershipTransferred struct {
	PreviousOwner        common.Address
	NewOwner             common.Address
	OwnershipTransferred bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner, bool indexed _ownershipTransferred)
func (_Payment *PaymentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (*PaymentOwnershipTransferredIterator, error) {

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

	logs, sub, err := _Payment.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return &PaymentOwnershipTransferredIterator{contract: _Payment.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x7fdc2a4b6eb39ec3363d710d188620bd1e97b3c434161f187b4d0dc0544faa58.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner, bool indexed _ownershipTransferred)
func (_Payment *PaymentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address, _ownershipTransferred []bool) (event.Subscription, error) {

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

	logs, sub, err := _Payment.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule, _ownershipTransferredRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentOwnershipTransferred)
				if err := _Payment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Payment *PaymentFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentOwnershipTransferred, error) {
	event := new(PaymentOwnershipTransferred)
	if err := _Payment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentSourceProofSubmittedIterator is returned from FilterSourceProofSubmitted and is used to iterate over the raw logs and unpacked data for SourceProofSubmitted events raised by the Payment contract.
type PaymentSourceProofSubmittedIterator struct {
	Event *PaymentSourceProofSubmitted // Event containing the contract specifics and raw log

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
func (it *PaymentSourceProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentSourceProofSubmitted)
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
		it.Event = new(PaymentSourceProofSubmitted)
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
func (it *PaymentSourceProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentSourceProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentSourceProofSubmitted represents a SourceProofSubmitted event raised by the Payment contract.
type PaymentSourceProofSubmitted struct {
	ChannelId        [32]byte
	RecoveredAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSourceProofSubmitted is a free log retrieval operation binding the contract event 0xef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a67.
//
// Solidity: event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Payment *PaymentFilterer) FilterSourceProofSubmitted(opts *bind.FilterOpts, _channelId [][32]byte, _recoveredAddress []common.Address) (*PaymentSourceProofSubmittedIterator, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}
	var _recoveredAddressRule []interface{}
	for _, _recoveredAddressItem := range _recoveredAddress {
		_recoveredAddressRule = append(_recoveredAddressRule, _recoveredAddressItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "SourceProofSubmitted", _channelIdRule, _recoveredAddressRule)
	if err != nil {
		return nil, err
	}
	return &PaymentSourceProofSubmittedIterator{contract: _Payment.contract, event: "SourceProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchSourceProofSubmitted is a free log subscription operation binding the contract event 0xef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a67.
//
// Solidity: event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Payment *PaymentFilterer) WatchSourceProofSubmitted(opts *bind.WatchOpts, sink chan<- *PaymentSourceProofSubmitted, _channelId [][32]byte, _recoveredAddress []common.Address) (event.Subscription, error) {

	var _channelIdRule []interface{}
	for _, _channelIdItem := range _channelId {
		_channelIdRule = append(_channelIdRule, _channelIdItem)
	}
	var _recoveredAddressRule []interface{}
	for _, _recoveredAddressItem := range _recoveredAddress {
		_recoveredAddressRule = append(_recoveredAddressRule, _recoveredAddressItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "SourceProofSubmitted", _channelIdRule, _recoveredAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentSourceProofSubmitted)
				if err := _Payment.contract.UnpackLog(event, "SourceProofSubmitted", log); err != nil {
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

// ParseSourceProofSubmitted is a log parse operation binding the contract event 0xef7fbbdebba5df7bf59c00d02ad872d452a67a5a4a4170adc3c95f794b478a67.
//
// Solidity: event SourceProofSubmitted(bytes32 indexed _channelId, address indexed _recoveredAddress)
func (_Payment *PaymentFilterer) ParseSourceProofSubmitted(log types.Log) (*PaymentSourceProofSubmitted, error) {
	event := new(PaymentSourceProofSubmitted)
	if err := _Payment.contract.UnpackLog(event, "SourceProofSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the Payment contract.
type PaymentTokensWithdrawnIterator struct {
	Event *PaymentTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentTokensWithdrawn)
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
		it.Event = new(PaymentTokensWithdrawn)
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
func (it *PaymentTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentTokensWithdrawn represents a TokensWithdrawn event raised by the Payment contract.
type PaymentTokensWithdrawn struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0xb47e37d80bb9227b793caed3e80256b8eac5e891b85d33f583d43ce03cf34273.
//
// Solidity: event TokensWithdrawn(address indexed _tokenAddress)
func (_Payment *PaymentFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, _tokenAddress []common.Address) (*PaymentTokensWithdrawnIterator, error) {

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "TokensWithdrawn", _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &PaymentTokensWithdrawnIterator{contract: _Payment.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0xb47e37d80bb9227b793caed3e80256b8eac5e891b85d33f583d43ce03cf34273.
//
// Solidity: event TokensWithdrawn(address indexed _tokenAddress)
func (_Payment *PaymentFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentTokensWithdrawn, _tokenAddress []common.Address) (event.Subscription, error) {

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "TokensWithdrawn", _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentTokensWithdrawn)
				if err := _Payment.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0xb47e37d80bb9227b793caed3e80256b8eac5e891b85d33f583d43ce03cf34273.
//
// Solidity: event TokensWithdrawn(address indexed _tokenAddress)
func (_Payment *PaymentFilterer) ParseTokensWithdrawn(log types.Log) (*PaymentTokensWithdrawn, error) {
	event := new(PaymentTokensWithdrawn)
	if err := _Payment.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
