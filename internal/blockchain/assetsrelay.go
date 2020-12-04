// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// AssetsRelayABI is the input ABI used to generate the binding from.
const AssetsRelayABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"int256\"}],\"name\":\"sign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"verifyAgainstPublishers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"getPublisherByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"verifyAgainstPublisherWithFallback\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKeys\",\"type\":\"address[]\"}],\"name\":\"disablePublisher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPublishers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"aContract\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAssetCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"getAssetCountForHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKeys\",\"type\":\"address[]\"},{\"name\":\"level\",\"type\":\"int256\"}],\"name\":\"setPublisherLevel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createdAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"verifyByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetsContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashIndex\",\"type\":\"uint256\"}],\"name\":\"getHashByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"aContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// AssetsRelayFuncSigs maps the 4-byte function signature to its string representation.
var AssetsRelayFuncSigs = map[string]string{
	"ddfe5b2d": "assetsContract()",
	"cf09e0d0": "createdAt()",
	"54df1eab": "disablePublisher(address[])",
	"a0aead4d": "getAssetCount()",
	"bf91e82f": "getAssetCountForHash(string)",
	"ec058186": "getHashByIndex(uint256)",
	"38052e27": "getPublisherByAddress(address)",
	"6c6071aa": "getPublishers()",
	"8da5cb5b": "owner()",
	"75f890ab": "setContract(address)",
	"c9e07f09": "setPublisherLevel(address[],int256)",
	"28f7f9b8": "sign(string,int256)",
	"bb9c6c3e": "verify(string)",
	"402cf5cc": "verifyAgainstPublisherWithFallback(string,address)",
	"32889c3d": "verifyAgainstPublishers(string,address[])",
	"d6ce25a9": "verifyByIndex(string,uint256)",
}

// AssetsRelayBin is the compiled bytecode used for deploying new contracts.
var AssetsRelayBin = "0x608060405234801561001057600080fd5b5060405160208061134e833981016040525160028054600160a060020a03909216600160a060020a03199283161790556000805490911633179055426001556112f08061005e6000396000f3006080604052600436106100cc5763ffffffff60e060020a60003504166328f7f9b881146100d157806332889c3d146100f757806338052e27146101b9578063402cf5cc1461020257806354df1eab146102665780636c6071aa1461028657806375f890ab146102eb5780638da5cb5b1461030c578063a0aead4d1461033d578063bb9c6c3e14610364578063bf91e82f14610384578063c9e07f09146103a4578063cf09e0d0146103c8578063d6ce25a9146103dd578063ddfe5b2d14610401578063ec05818614610416575b600080fd5b3480156100dd57600080fd5b506100f56024600480358281019291013590356104a3565b005b34801561010357600080fd5b506040805160206004803580820135601f8101849004840285018401909552848452610189943694929360249392840191908190840183828082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506105479650505050505050565b60408051600160a060020a0390951685526020850193909352838301919091526060830152519081900360800190f35b3480156101c557600080fd5b506101da600160a060020a03600435166107d1565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b34801561020e57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261018994369492936024939284019190819084018382808284375094975050509235600160a060020a0316935061088292505050565b34801561027257600080fd5b506100f56004803560248101910135610c06565b34801561029257600080fd5b5061029b610ca6565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102d75781810151838201526020016102bf565b505050509050019250505060405180910390f35b3480156102f757600080fd5b506100f5600160a060020a0360043516610d95565b34801561031857600080fd5b50610321610dd4565b60408051600160a060020a039092168252519081900360200190f35b34801561034957600080fd5b50610352610de3565b60408051918252519081900360200190f35b34801561037057600080fd5b506101896004803560248101910135610e73565b34801561039057600080fd5b506103526004803560248101910135610f46565b3480156103b057600080fd5b506100f5602460048035828101929101359035610ff9565b3480156103d457600080fd5b50610352611083565b3480156103e957600080fd5b50610189602460048035828101929101359035611089565b34801561040d57600080fd5b50610321611166565b34801561042257600080fd5b5061042e600435611175565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610468578181015183820152602001610450565b50505050905090810190601f1680156104955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600254604080517f28f7f9b8000000000000000000000000000000000000000000000000000000008152602481018490526004810191825260448101859052600160a060020a03909216916328f7f9b89186918691869181906064018585808284378201915050945050505050600060405180830381600087803b15801561052a57600080fd5b505af115801561053e573d6000803e3d6000fd5b50505050505050565b600080600080600080600080600080600260009054906101000a9004600160a060020a0316600160a060020a031663bf91e82f8d6040518263ffffffff1660e060020a0281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156105ca5781810151838201526020016105b2565b50505050905090810190601f1680156105f75780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561061657600080fd5b505af115801561062a573d6000803e3d6000fd5b505050506040513d602081101561064057600080fd5b505195506000861161065f5760009950899850600297508896506107c2565b5060009350839250829150819050845b60008111156107b357600260009054906101000a9004600160a060020a0316600160a060020a031663d6ce25a98d600184036040518363ffffffff1660e060020a0281526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b838110156106f65781810151838201526020016106de565b50505050905090810190601f1680156107235780820380516001836020036101000a031916815260200191505b509350505050608060405180830381600087803b15801561074357600080fd5b505af1158015610757573d6000803e3d6000fd5b505050506040513d608081101561076d57600080fd5b50805160208201516040830151606090930151919750955090935091506107948b86611269565b156107aa578484848499509950995099506107c2565b6000190161066f565b60009950899850600297508896505b50505050505092959194509250565b600254604080517f38052e27000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093849384939116916338052e279160248082019260609290919082900301818787803b15801561083e57600080fd5b505af1158015610852573d6000803e3d6000fd5b505050506040513d606081101561086857600080fd5b508051602082015160409092015190969195509350915050565b600080600080600080600080600080600260009054906101000a9004600160a060020a0316600160a060020a031663bf91e82f8d6040518263ffffffff1660e060020a0281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109055781810151838201526020016108ed565b50505050905090810190601f1680156109325780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561095157600080fd5b505af1158015610965573d6000803e3d6000fd5b505050506040513d602081101561097b57600080fd5b505195506000861161099a5760009950899850600297508896506107c2565b5060009350839250829150819050845b6000811115610af357600260009054906101000a9004600160a060020a0316600160a060020a031663d6ce25a98d600184036040518363ffffffff1660e060020a0281526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610a31578181015183820152602001610a19565b50505050905090810190601f168015610a5e5780820380516001836020036101000a031916815260200191505b509350505050608060405180830381600087803b158015610a7e57600080fd5b505af1158015610a92573d6000803e3d6000fd5b505050506040513d6080811015610aa857600080fd5b5080516020820151604083015160609093015191975095509093509150600160a060020a03808616908c161415610aea578484848499509950995099506107c2565b600019016109aa565b600260009054906101000a9004600160a060020a0316600160a060020a031663bb9c6c3e8d6040518263ffffffff1660e060020a0281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b67578181015183820152602001610b4f565b50505050905090810190601f168015610b945780820380516001836020036101000a031916815260200191505b5092505050608060405180830381600087803b158015610bb357600080fd5b505af1158015610bc7573d6000803e3d6000fd5b505050506040513d6080811015610bdd57600080fd5b50805160208201516040830151606090930151919f909e50919c509a5098505050505050505050565b6002546040517f54df1eab00000000000000000000000000000000000000000000000000000000815260206004820181815260248301859052600160a060020a03909316926354df1eab9286928692918291604490910190859085028082843782019150509350505050600060405180830381600087803b158015610c8a57600080fd5b505af1158015610c9e573d6000803e3d6000fd5b505050505050565b600254604080517f6c6071aa0000000000000000000000000000000000000000000000000000000081529051606092600160a060020a031691636c6071aa91600480830192600092919082900301818387803b158015610d0557600080fd5b505af1158015610d19573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610d4257600080fd5b810190808051640100000000811115610d5a57600080fd5b82016020810184811115610d6d57600080fd5b8151856020820283011164010000000082111715610d8a57600080fd5b509094505050505090565b600054600160a060020a0316321415610dd1576002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600054600160a060020a031681565b600254604080517fa0aead4d0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163a0aead4d91600480830192602092919082900301818787803b158015610e4257600080fd5b505af1158015610e56573d6000803e3d6000fd5b505050506040513d6020811015610e6c57600080fd5b5051905090565b6002546040517fbb9c6c3e00000000000000000000000000000000000000000000000000000000815260206004820190815260248201849052600092839283928392600160a060020a03169163bb9c6c3e918991899190819060440184848082843782019150509350505050608060405180830381600087803b158015610ef957600080fd5b505af1158015610f0d573d6000803e3d6000fd5b505050506040513d6080811015610f2357600080fd5b508051602082015160408301516060909301519199909850919650945092505050565b6002546040517fbf91e82f00000000000000000000000000000000000000000000000000000000815260206004820190815260248201849052600092600160a060020a03169163bf91e82f918691869190819060440184848082843782019150509350505050602060405180830381600087803b158015610fc657600080fd5b505af1158015610fda573d6000803e3d6000fd5b505050506040513d6020811015610ff057600080fd5b50519392505050565b600254604080517fc9e07f09000000000000000000000000000000000000000000000000000000008152602481018490526004810191825260448101859052600160a060020a039092169163c9e07f099186918691869181906064018560208602808284378201915050945050505050600060405180830381600087803b15801561052a57600080fd5b60015481565b600254604080517fd6ce25a9000000000000000000000000000000000000000000000000000000008152602481018490526004810191825260448101859052600092839283928392600160a060020a03169163d6ce25a9918a918a918a919081906064018585808284378201915050945050505050608060405180830381600087803b15801561111857600080fd5b505af115801561112c573d6000803e3d6000fd5b505050506040513d608081101561114257600080fd5b50805160208201516040830151606090930151919a90995091975095509350505050565b600254600160a060020a031681565b600254604080517fe9d1e616000000000000000000000000000000000000000000000000000000008152600481018490529051606092600160a060020a03169163e9d1e61691602480830192600092919082900301818387803b1580156111db57600080fd5b505af11580156111ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561121857600080fd5b81019080805164010000000081111561123057600080fd5b8201602081018481111561124357600080fd5b815164010000000081118282018710171561125d57600080fd5b50909695505050505050565b6000805b83518110156112b85782600160a060020a0316848281518110151561128e57fe5b90602001906020020151600160a060020a031614156112b057600191506112bd565b60010161126d565b600091505b50929150505600a165627a7a72305820e72b9311c3149c09f536f6b71ef3c7a7e3d8670cd2f9f0327a183e23da4f0ccc0029"

// DeployAssetsRelay deploys a new Ethereum contract, binding an instance of AssetsRelay to it.
func DeployAssetsRelay(auth *bind.TransactOpts, backend bind.ContractBackend, aContract common.Address) (common.Address, *types.Transaction, *AssetsRelay, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetsRelayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetsRelayBin), backend, aContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssetsRelay{AssetsRelayCaller: AssetsRelayCaller{contract: contract}, AssetsRelayTransactor: AssetsRelayTransactor{contract: contract}, AssetsRelayFilterer: AssetsRelayFilterer{contract: contract}}, nil
}

// AssetsRelay is an auto generated Go binding around an Ethereum contract.
type AssetsRelay struct {
	AssetsRelayCaller     // Read-only binding to the contract
	AssetsRelayTransactor // Write-only binding to the contract
	AssetsRelayFilterer   // Log filterer for contract events
}

// AssetsRelayCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetsRelayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsRelayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetsRelayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsRelayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetsRelayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsRelaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetsRelaySession struct {
	Contract     *AssetsRelay      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetsRelayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetsRelayCallerSession struct {
	Contract *AssetsRelayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AssetsRelayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetsRelayTransactorSession struct {
	Contract     *AssetsRelayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AssetsRelayRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetsRelayRaw struct {
	Contract *AssetsRelay // Generic contract binding to access the raw methods on
}

// AssetsRelayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetsRelayCallerRaw struct {
	Contract *AssetsRelayCaller // Generic read-only contract binding to access the raw methods on
}

// AssetsRelayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetsRelayTransactorRaw struct {
	Contract *AssetsRelayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetsRelay creates a new instance of AssetsRelay, bound to a specific deployed contract.
func NewAssetsRelay(address common.Address, backend bind.ContractBackend) (*AssetsRelay, error) {
	contract, err := bindAssetsRelay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetsRelay{AssetsRelayCaller: AssetsRelayCaller{contract: contract}, AssetsRelayTransactor: AssetsRelayTransactor{contract: contract}, AssetsRelayFilterer: AssetsRelayFilterer{contract: contract}}, nil
}

// NewAssetsRelayCaller creates a new read-only instance of AssetsRelay, bound to a specific deployed contract.
func NewAssetsRelayCaller(address common.Address, caller bind.ContractCaller) (*AssetsRelayCaller, error) {
	contract, err := bindAssetsRelay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetsRelayCaller{contract: contract}, nil
}

// NewAssetsRelayTransactor creates a new write-only instance of AssetsRelay, bound to a specific deployed contract.
func NewAssetsRelayTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetsRelayTransactor, error) {
	contract, err := bindAssetsRelay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetsRelayTransactor{contract: contract}, nil
}

// NewAssetsRelayFilterer creates a new log filterer instance of AssetsRelay, bound to a specific deployed contract.
func NewAssetsRelayFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetsRelayFilterer, error) {
	contract, err := bindAssetsRelay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetsRelayFilterer{contract: contract}, nil
}

// bindAssetsRelay binds a generic wrapper to an already deployed contract.
func bindAssetsRelay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetsRelayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetsRelay *AssetsRelayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetsRelay.Contract.AssetsRelayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetsRelay *AssetsRelayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetsRelay.Contract.AssetsRelayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetsRelay *AssetsRelayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetsRelay.Contract.AssetsRelayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetsRelay *AssetsRelayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetsRelay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetsRelay *AssetsRelayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetsRelay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetsRelay *AssetsRelayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetsRelay.Contract.contract.Transact(opts, method, params...)
}

// AssetsContract is a free data retrieval call binding the contract method 0xddfe5b2d.
//
// Solidity: function assetsContract() view returns(address)
func (_AssetsRelay *AssetsRelayCaller) AssetsContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "assetsContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetsContract is a free data retrieval call binding the contract method 0xddfe5b2d.
//
// Solidity: function assetsContract() view returns(address)
func (_AssetsRelay *AssetsRelaySession) AssetsContract() (common.Address, error) {
	return _AssetsRelay.Contract.AssetsContract(&_AssetsRelay.CallOpts)
}

// AssetsContract is a free data retrieval call binding the contract method 0xddfe5b2d.
//
// Solidity: function assetsContract() view returns(address)
func (_AssetsRelay *AssetsRelayCallerSession) AssetsContract() (common.Address, error) {
	return _AssetsRelay.Contract.AssetsContract(&_AssetsRelay.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_AssetsRelay *AssetsRelayCaller) CreatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_AssetsRelay *AssetsRelaySession) CreatedAt() (*big.Int, error) {
	return _AssetsRelay.Contract.CreatedAt(&_AssetsRelay.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_AssetsRelay *AssetsRelayCallerSession) CreatedAt() (*big.Int, error) {
	return _AssetsRelay.Contract.CreatedAt(&_AssetsRelay.CallOpts)
}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_AssetsRelay *AssetsRelayCaller) GetAssetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "getAssetCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_AssetsRelay *AssetsRelaySession) GetAssetCount() (*big.Int, error) {
	return _AssetsRelay.Contract.GetAssetCount(&_AssetsRelay.CallOpts)
}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_AssetsRelay *AssetsRelayCallerSession) GetAssetCount() (*big.Int, error) {
	return _AssetsRelay.Contract.GetAssetCount(&_AssetsRelay.CallOpts)
}

// GetAssetCountForHash is a free data retrieval call binding the contract method 0xbf91e82f.
//
// Solidity: function getAssetCountForHash(string hash) view returns(uint256)
func (_AssetsRelay *AssetsRelayCaller) GetAssetCountForHash(opts *bind.CallOpts, hash string) (*big.Int, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "getAssetCountForHash", hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetCountForHash is a free data retrieval call binding the contract method 0xbf91e82f.
//
// Solidity: function getAssetCountForHash(string hash) view returns(uint256)
func (_AssetsRelay *AssetsRelaySession) GetAssetCountForHash(hash string) (*big.Int, error) {
	return _AssetsRelay.Contract.GetAssetCountForHash(&_AssetsRelay.CallOpts, hash)
}

// GetAssetCountForHash is a free data retrieval call binding the contract method 0xbf91e82f.
//
// Solidity: function getAssetCountForHash(string hash) view returns(uint256)
func (_AssetsRelay *AssetsRelayCallerSession) GetAssetCountForHash(hash string) (*big.Int, error) {
	return _AssetsRelay.Contract.GetAssetCountForHash(&_AssetsRelay.CallOpts, hash)
}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 hashIndex) view returns(string)
func (_AssetsRelay *AssetsRelayCaller) GetHashByIndex(opts *bind.CallOpts, hashIndex *big.Int) (string, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "getHashByIndex", hashIndex)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 hashIndex) view returns(string)
func (_AssetsRelay *AssetsRelaySession) GetHashByIndex(hashIndex *big.Int) (string, error) {
	return _AssetsRelay.Contract.GetHashByIndex(&_AssetsRelay.CallOpts, hashIndex)
}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 hashIndex) view returns(string)
func (_AssetsRelay *AssetsRelayCallerSession) GetHashByIndex(hashIndex *big.Int) (string, error) {
	return _AssetsRelay.Contract.GetHashByIndex(&_AssetsRelay.CallOpts, hashIndex)
}

// GetPublisherByAddress is a free data retrieval call binding the contract method 0x38052e27.
//
// Solidity: function getPublisherByAddress(address publicKey) view returns(address, int256, uint256)
func (_AssetsRelay *AssetsRelayCaller) GetPublisherByAddress(opts *bind.CallOpts, publicKey common.Address) (common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "getPublisherByAddress", publicKey)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPublisherByAddress is a free data retrieval call binding the contract method 0x38052e27.
//
// Solidity: function getPublisherByAddress(address publicKey) view returns(address, int256, uint256)
func (_AssetsRelay *AssetsRelaySession) GetPublisherByAddress(publicKey common.Address) (common.Address, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.GetPublisherByAddress(&_AssetsRelay.CallOpts, publicKey)
}

// GetPublisherByAddress is a free data retrieval call binding the contract method 0x38052e27.
//
// Solidity: function getPublisherByAddress(address publicKey) view returns(address, int256, uint256)
func (_AssetsRelay *AssetsRelayCallerSession) GetPublisherByAddress(publicKey common.Address) (common.Address, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.GetPublisherByAddress(&_AssetsRelay.CallOpts, publicKey)
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(address[])
func (_AssetsRelay *AssetsRelayCaller) GetPublishers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "getPublishers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(address[])
func (_AssetsRelay *AssetsRelaySession) GetPublishers() ([]common.Address, error) {
	return _AssetsRelay.Contract.GetPublishers(&_AssetsRelay.CallOpts)
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(address[])
func (_AssetsRelay *AssetsRelayCallerSession) GetPublishers() ([]common.Address, error) {
	return _AssetsRelay.Contract.GetPublishers(&_AssetsRelay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetsRelay *AssetsRelayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetsRelay *AssetsRelaySession) Owner() (common.Address, error) {
	return _AssetsRelay.Contract.Owner(&_AssetsRelay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetsRelay *AssetsRelayCallerSession) Owner() (common.Address, error) {
	return _AssetsRelay.Contract.Owner(&_AssetsRelay.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string hash) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelayCaller) Verify(opts *bind.CallOpts, hash string) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "verify", hash)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string hash) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelaySession) Verify(hash string) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.Verify(&_AssetsRelay.CallOpts, hash)
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string hash) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelayCallerSession) Verify(hash string) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.Verify(&_AssetsRelay.CallOpts, hash)
}

// VerifyAgainstPublisherWithFallback is a free data retrieval call binding the contract method 0x402cf5cc.
//
// Solidity: function verifyAgainstPublisherWithFallback(string hash, address a) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelayCaller) VerifyAgainstPublisherWithFallback(opts *bind.CallOpts, hash string, a common.Address) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "verifyAgainstPublisherWithFallback", hash, a)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// VerifyAgainstPublisherWithFallback is a free data retrieval call binding the contract method 0x402cf5cc.
//
// Solidity: function verifyAgainstPublisherWithFallback(string hash, address a) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelaySession) VerifyAgainstPublisherWithFallback(hash string, a common.Address) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.VerifyAgainstPublisherWithFallback(&_AssetsRelay.CallOpts, hash, a)
}

// VerifyAgainstPublisherWithFallback is a free data retrieval call binding the contract method 0x402cf5cc.
//
// Solidity: function verifyAgainstPublisherWithFallback(string hash, address a) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelayCallerSession) VerifyAgainstPublisherWithFallback(hash string, a common.Address) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.VerifyAgainstPublisherWithFallback(&_AssetsRelay.CallOpts, hash, a)
}

// VerifyAgainstPublishers is a free data retrieval call binding the contract method 0x32889c3d.
//
// Solidity: function verifyAgainstPublishers(string hash, address[] addresses) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelayCaller) VerifyAgainstPublishers(opts *bind.CallOpts, hash string, addresses []common.Address) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "verifyAgainstPublishers", hash, addresses)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// VerifyAgainstPublishers is a free data retrieval call binding the contract method 0x32889c3d.
//
// Solidity: function verifyAgainstPublishers(string hash, address[] addresses) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelaySession) VerifyAgainstPublishers(hash string, addresses []common.Address) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.VerifyAgainstPublishers(&_AssetsRelay.CallOpts, hash, addresses)
}

// VerifyAgainstPublishers is a free data retrieval call binding the contract method 0x32889c3d.
//
// Solidity: function verifyAgainstPublishers(string hash, address[] addresses) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelayCallerSession) VerifyAgainstPublishers(hash string, addresses []common.Address) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.VerifyAgainstPublishers(&_AssetsRelay.CallOpts, hash, addresses)
}

// VerifyByIndex is a free data retrieval call binding the contract method 0xd6ce25a9.
//
// Solidity: function verifyByIndex(string hash, uint256 assetIndex) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelayCaller) VerifyByIndex(opts *bind.CallOpts, hash string, assetIndex *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AssetsRelay.contract.Call(opts, &out, "verifyByIndex", hash, assetIndex)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// VerifyByIndex is a free data retrieval call binding the contract method 0xd6ce25a9.
//
// Solidity: function verifyByIndex(string hash, uint256 assetIndex) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelaySession) VerifyByIndex(hash string, assetIndex *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.VerifyByIndex(&_AssetsRelay.CallOpts, hash, assetIndex)
}

// VerifyByIndex is a free data retrieval call binding the contract method 0xd6ce25a9.
//
// Solidity: function verifyByIndex(string hash, uint256 assetIndex) view returns(address, int256, int256, uint256)
func (_AssetsRelay *AssetsRelayCallerSession) VerifyByIndex(hash string, assetIndex *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsRelay.Contract.VerifyByIndex(&_AssetsRelay.CallOpts, hash, assetIndex)
}

// DisablePublisher is a paid mutator transaction binding the contract method 0x54df1eab.
//
// Solidity: function disablePublisher(address[] publicKeys) returns()
func (_AssetsRelay *AssetsRelayTransactor) DisablePublisher(opts *bind.TransactOpts, publicKeys []common.Address) (*types.Transaction, error) {
	return _AssetsRelay.contract.Transact(opts, "disablePublisher", publicKeys)
}

// DisablePublisher is a paid mutator transaction binding the contract method 0x54df1eab.
//
// Solidity: function disablePublisher(address[] publicKeys) returns()
func (_AssetsRelay *AssetsRelaySession) DisablePublisher(publicKeys []common.Address) (*types.Transaction, error) {
	return _AssetsRelay.Contract.DisablePublisher(&_AssetsRelay.TransactOpts, publicKeys)
}

// DisablePublisher is a paid mutator transaction binding the contract method 0x54df1eab.
//
// Solidity: function disablePublisher(address[] publicKeys) returns()
func (_AssetsRelay *AssetsRelayTransactorSession) DisablePublisher(publicKeys []common.Address) (*types.Transaction, error) {
	return _AssetsRelay.Contract.DisablePublisher(&_AssetsRelay.TransactOpts, publicKeys)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(address aContract) returns()
func (_AssetsRelay *AssetsRelayTransactor) SetContract(opts *bind.TransactOpts, aContract common.Address) (*types.Transaction, error) {
	return _AssetsRelay.contract.Transact(opts, "setContract", aContract)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(address aContract) returns()
func (_AssetsRelay *AssetsRelaySession) SetContract(aContract common.Address) (*types.Transaction, error) {
	return _AssetsRelay.Contract.SetContract(&_AssetsRelay.TransactOpts, aContract)
}

// SetContract is a paid mutator transaction binding the contract method 0x75f890ab.
//
// Solidity: function setContract(address aContract) returns()
func (_AssetsRelay *AssetsRelayTransactorSession) SetContract(aContract common.Address) (*types.Transaction, error) {
	return _AssetsRelay.Contract.SetContract(&_AssetsRelay.TransactOpts, aContract)
}

// SetPublisherLevel is a paid mutator transaction binding the contract method 0xc9e07f09.
//
// Solidity: function setPublisherLevel(address[] publicKeys, int256 level) returns()
func (_AssetsRelay *AssetsRelayTransactor) SetPublisherLevel(opts *bind.TransactOpts, publicKeys []common.Address, level *big.Int) (*types.Transaction, error) {
	return _AssetsRelay.contract.Transact(opts, "setPublisherLevel", publicKeys, level)
}

// SetPublisherLevel is a paid mutator transaction binding the contract method 0xc9e07f09.
//
// Solidity: function setPublisherLevel(address[] publicKeys, int256 level) returns()
func (_AssetsRelay *AssetsRelaySession) SetPublisherLevel(publicKeys []common.Address, level *big.Int) (*types.Transaction, error) {
	return _AssetsRelay.Contract.SetPublisherLevel(&_AssetsRelay.TransactOpts, publicKeys, level)
}

// SetPublisherLevel is a paid mutator transaction binding the contract method 0xc9e07f09.
//
// Solidity: function setPublisherLevel(address[] publicKeys, int256 level) returns()
func (_AssetsRelay *AssetsRelayTransactorSession) SetPublisherLevel(publicKeys []common.Address, level *big.Int) (*types.Transaction, error) {
	return _AssetsRelay.Contract.SetPublisherLevel(&_AssetsRelay.TransactOpts, publicKeys, level)
}

// Sign is a paid mutator transaction binding the contract method 0x28f7f9b8.
//
// Solidity: function sign(string hash, int256 status) returns()
func (_AssetsRelay *AssetsRelayTransactor) Sign(opts *bind.TransactOpts, hash string, status *big.Int) (*types.Transaction, error) {
	return _AssetsRelay.contract.Transact(opts, "sign", hash, status)
}

// Sign is a paid mutator transaction binding the contract method 0x28f7f9b8.
//
// Solidity: function sign(string hash, int256 status) returns()
func (_AssetsRelay *AssetsRelaySession) Sign(hash string, status *big.Int) (*types.Transaction, error) {
	return _AssetsRelay.Contract.Sign(&_AssetsRelay.TransactOpts, hash, status)
}

// Sign is a paid mutator transaction binding the contract method 0x28f7f9b8.
//
// Solidity: function sign(string hash, int256 status) returns()
func (_AssetsRelay *AssetsRelayTransactorSession) Sign(hash string, status *big.Int) (*types.Transaction, error) {
	return _AssetsRelay.Contract.Sign(&_AssetsRelay.TransactOpts, hash, status)
}

// AssetsV1ABI is the input ABI used to generate the binding from.
const AssetsV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"int256\"}],\"name\":\"sign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publisherPublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"getPublisherByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKeys\",\"type\":\"address[]\"}],\"name\":\"disablePublisher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPublishers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAssetCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"getAssetCountForHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKeys\",\"type\":\"address[]\"},{\"name\":\"level\",\"type\":\"int256\"}],\"name\":\"setPublisherLevel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createdAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"verifyByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"newPublisher\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"level\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"publisherLevelChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"newAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"level\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"assetSigned\",\"type\":\"event\"}]"

// AssetsV1FuncSigs maps the 4-byte function signature to its string representation.
var AssetsV1FuncSigs = map[string]string{
	"e9d1e616": "assetHashes(uint256)",
	"cf09e0d0": "createdAt()",
	"54df1eab": "disablePublisher(address[])",
	"a0aead4d": "getAssetCount()",
	"bf91e82f": "getAssetCountForHash(string)",
	"38052e27": "getPublisherByAddress(address)",
	"6c6071aa": "getPublishers()",
	"41c0e1b5": "kill()",
	"8da5cb5b": "owner()",
	"2c5f00a1": "publisherPublicKeys(uint256)",
	"c9e07f09": "setPublisherLevel(address[],int256)",
	"28f7f9b8": "sign(string,int256)",
	"bb9c6c3e": "verify(string)",
	"d6ce25a9": "verifyByIndex(string,uint256)",
}

// AssetsV1Bin is the compiled bytecode used for deploying new contracts.
var AssetsV1Bin = "0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905542600155611142806100366000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166328f7f9b881146100d45780632c5f00a1146100fa57806338052e271461012e57806341c0e1b51461017757806354df1eab1461018c5780636c6071aa146101ac5780638da5cb5b14610211578063a0aead4d14610226578063bb9c6c3e1461024d578063bf91e82f1461029d578063c9e07f09146102bd578063cf09e0d0146102e1578063d6ce25a9146102f6578063e9d1e6161461031a575b600080fd5b3480156100e057600080fd5b506100f86024600480358281019291013590356103a7565b005b34801561010657600080fd5b50610112600435610676565b60408051600160a060020a039092168252519081900360200190f35b34801561013a57600080fd5b5061014f600160a060020a036004351661069e565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b34801561018357600080fd5b506100f86106f7565b34801561019857600080fd5b506100f8600480356024810191013561071a565b3480156101b857600080fd5b506101c161090b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101fd5781810151838201526020016101e5565b505050509050019250505060405180910390f35b34801561021d57600080fd5b5061011261096e565b34801561023257600080fd5b5061023b61097d565b60408051918252519081900360200190f35b34801561025957600080fd5b5061026d6004803560248101910135610983565b60408051600160a060020a0390951685526020850193909352838301919091526060830152519081900360800190f35b3480156102a957600080fd5b5061023b6004803560248101910135610ba1565b3480156102c957600080fd5b506100f8602460048035828101929101359035610bd1565b3480156102ed57600080fd5b5061023b610db3565b34801561030257600080fd5b5061026d602460048035828101929101359035610db9565b34801561032657600080fd5b50610332600435610f05565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561036c578181015183820152602001610354565b50505050905090810190601f1680156103995780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103af610fac565b60006103b9610fd7565b32600090815260026020819052604082200154111561066e573260009081526002602081815260409283902083516060810185528154600160a060020a039081168252600183015482850152919093015483850152835160a081018552835190911681528351601f8a0183900483028101830190945288845291955042945090918282019189908990819084018382808284378201915050505050508152602001846020015181526020018581526020018381525090506000600387876040518083838082843782019150509250505090815260200160405180910390208054905011151561053b577f2bf03f94cf4fc3c57bcabb0cb88b62a4530d6a0de975f8f7a898bacde02e4d7a868684604051808060200183815260200182810382528585828181526020019250808284376040519201829003965090945050505050a16005805460018101808355600092909252610538907f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018888611010565b50505b600386866040518083838082843791909101948552505060405160209381900384019020805460018082018084556000938452928690208751600590930201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390931692909217825586860151805193968896509294506105c19391850192019061108e565b506040820151816002015560608201518160030155608082015181600401555050507f5246d0df7f13afa7adb921dfc9e2762d77304794bae8e4734ef030f543fe837d83600001518787866020015188876040518087600160a060020a0316600160a060020a031681526020018060200185815260200184815260200183815260200182810382528787828181526020019250808284376040519201829003995090975050505050505050a15b505050505050565b600480548290811061068457fe5b600091825260209091200154600160a060020a0316905081565b60008060006106ab610fac565b50505050600160a060020a039081166000908152600260208181526040928390208351606081018552815490951680865260018201549286018390529201549390920183905292909190565b600054600160a060020a031632141561071857600054600160a060020a0316ff5b565b600080600080610728610fac565b600054600160a060020a0316321415610902576000199450600093505b858410156109025786868581811061075957fe5b90506020020135600160a060020a0316925042915060606040519081016040528084600160a060020a0316815260200186815260200183815250905060006002600085600160a060020a0316600160a060020a03168152602001908152602001600020600201541115156108645760408051600160a060020a03851681526020810184905281517f8c810fbf05dab8bcb3508516cffe5f356d0e2450a2828f6910c4fdd34ad625c3929181900390910190a1600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161790555b600160a060020a038381166000818152600260208181526040928390208651815473ffffffffffffffffffffffffffffffffffffffff19169616959095178555858101516001860155858301519490910193909355805191825291810187905280820184905290517f86ec446c5d262f7d5d58e56a50600a6fa1a1db5678cb6d548e5383e65a3853f1916060908290030190a1600190930192610745565b50505050505050565b6060600480548060200260200160405190810160405280929190818152602001828054801561096357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610945575b505050505090505b90565b600054600160a060020a031681565b60055490565b6000806000806060610993610fd7565b600061099d610fd7565b60038a8a604051808383808284379091019485525050604080519384900360209081018520805480830287018301909352828652935090915060009084015b82821015610ad55760008481526020908190206040805160a081018252600586029092018054600160a060020a03168352600180820180548451600261010094831615949094026000190190911692909204601f810187900487028301870190945283825293949193858301939192909190830182828015610a9f5780601f10610a7457610100808354040283529160200191610a9f565b820191906000526020600020905b815481529060010190602001808311610a8257829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481525050815260200190600101906109dc565b50505050935060008451111515610af9576000975087965060029550869450610b94565b6040805160a0810182526000808252825160208181018552828252830152600019928201929092526002606082015260808101829052935091505b8351821015610b77578382815181101515610b4b57fe5b90602001906020020151905082604001518160400151121515610b6c578092505b600190910190610b34565b826000015183604001518460600151856080015197509750975097505b5050505092959194509250565b60006003838360405180838380828437909101948552505060405192839003602001909220549250505092915050565b6000806000610bde610fac565b600054600160a060020a031632141561090257600093505b8584101561090257868685818110610c0a57fe5b90506020020135600160a060020a0316925042915060606040519081016040528084600160a060020a0316815260200186815260200183815250905060006002600085600160a060020a0316600160a060020a0316815260200190815260200160002060020154111515610d155760408051600160a060020a03851681526020810184905281517f8c810fbf05dab8bcb3508516cffe5f356d0e2450a2828f6910c4fdd34ad625c3929181900390910190a1600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161790555b600160a060020a038381166000818152600260208181526040928390208651815473ffffffffffffffffffffffffffffffffffffffff19169616959095178555858101516001860155858301519490910193909355805191825291810187905280820184905290517f86ec446c5d262f7d5d58e56a50600a6fa1a1db5678cb6d548e5383e65a3853f1916060908290030190a1600190930192610bf6565b60015481565b600080600080610dc7610fd7565b6003888860405180838380828437820191505092505050908152602001604051809103902086815481101515610df957fe5b60009182526020918290206040805160a08101825260059093029091018054600160a060020a03168352600180820180548451601f60026000199584161561010002959095019092169390930490810187900487028301870190945283825293949193858301939192909190830182828015610eb65780601f10610e8b57610100808354040283529160200191610eb6565b820191906000526020600020905b815481529060010190602001808311610e9957829003601f168201915b5050505050815260200160028201548152602001600382015481526020016004820154815250509050806000015181604001518260600151836080015194509450945094505093509350935093565b6005805482908110610f1357fe5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815293509091830182828015610fa45780601f10610f7957610100808354040283529160200191610fa4565b820191906000526020600020905b815481529060010190602001808311610f8757829003601f168201915b505050505081565b6060604051908101604052806000600160a060020a0316815260200160008152602001600081525090565b60a0604051908101604052806000600160a060020a03168152602001606081526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110515782800160ff1982351617855561107e565b8280016001018555821561107e579182015b8281111561107e578235825591602001919060010190611063565b5061108a9291506110fc565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110cf57805160ff191683800117855561107e565b8280016001018555821561107e579182015b8281111561107e5782518255916020019190600101906110e1565b61096b91905b8082111561108a57600081556001016111025600a165627a7a723058206cb3af83a49539925a4f8ada9cb9aad7ae65ce26aad41b2410bf40af458c23e90029"

// DeployAssetsV1 deploys a new Ethereum contract, binding an instance of AssetsV1 to it.
func DeployAssetsV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AssetsV1, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetsV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetsV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssetsV1{AssetsV1Caller: AssetsV1Caller{contract: contract}, AssetsV1Transactor: AssetsV1Transactor{contract: contract}, AssetsV1Filterer: AssetsV1Filterer{contract: contract}}, nil
}

// AssetsV1 is an auto generated Go binding around an Ethereum contract.
type AssetsV1 struct {
	AssetsV1Caller     // Read-only binding to the contract
	AssetsV1Transactor // Write-only binding to the contract
	AssetsV1Filterer   // Log filterer for contract events
}

// AssetsV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type AssetsV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetsV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetsV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetsV1Session struct {
	Contract     *AssetsV1         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetsV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetsV1CallerSession struct {
	Contract *AssetsV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AssetsV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetsV1TransactorSession struct {
	Contract     *AssetsV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AssetsV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type AssetsV1Raw struct {
	Contract *AssetsV1 // Generic contract binding to access the raw methods on
}

// AssetsV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetsV1CallerRaw struct {
	Contract *AssetsV1Caller // Generic read-only contract binding to access the raw methods on
}

// AssetsV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetsV1TransactorRaw struct {
	Contract *AssetsV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetsV1 creates a new instance of AssetsV1, bound to a specific deployed contract.
func NewAssetsV1(address common.Address, backend bind.ContractBackend) (*AssetsV1, error) {
	contract, err := bindAssetsV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetsV1{AssetsV1Caller: AssetsV1Caller{contract: contract}, AssetsV1Transactor: AssetsV1Transactor{contract: contract}, AssetsV1Filterer: AssetsV1Filterer{contract: contract}}, nil
}

// NewAssetsV1Caller creates a new read-only instance of AssetsV1, bound to a specific deployed contract.
func NewAssetsV1Caller(address common.Address, caller bind.ContractCaller) (*AssetsV1Caller, error) {
	contract, err := bindAssetsV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetsV1Caller{contract: contract}, nil
}

// NewAssetsV1Transactor creates a new write-only instance of AssetsV1, bound to a specific deployed contract.
func NewAssetsV1Transactor(address common.Address, transactor bind.ContractTransactor) (*AssetsV1Transactor, error) {
	contract, err := bindAssetsV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetsV1Transactor{contract: contract}, nil
}

// NewAssetsV1Filterer creates a new log filterer instance of AssetsV1, bound to a specific deployed contract.
func NewAssetsV1Filterer(address common.Address, filterer bind.ContractFilterer) (*AssetsV1Filterer, error) {
	contract, err := bindAssetsV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetsV1Filterer{contract: contract}, nil
}

// bindAssetsV1 binds a generic wrapper to an already deployed contract.
func bindAssetsV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetsV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetsV1 *AssetsV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetsV1.Contract.AssetsV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetsV1 *AssetsV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetsV1.Contract.AssetsV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetsV1 *AssetsV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetsV1.Contract.AssetsV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetsV1 *AssetsV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetsV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetsV1 *AssetsV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetsV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetsV1 *AssetsV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetsV1.Contract.contract.Transact(opts, method, params...)
}

// AssetHashes is a free data retrieval call binding the contract method 0xe9d1e616.
//
// Solidity: function assetHashes(uint256 ) view returns(string)
func (_AssetsV1 *AssetsV1Caller) AssetHashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "assetHashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetHashes is a free data retrieval call binding the contract method 0xe9d1e616.
//
// Solidity: function assetHashes(uint256 ) view returns(string)
func (_AssetsV1 *AssetsV1Session) AssetHashes(arg0 *big.Int) (string, error) {
	return _AssetsV1.Contract.AssetHashes(&_AssetsV1.CallOpts, arg0)
}

// AssetHashes is a free data retrieval call binding the contract method 0xe9d1e616.
//
// Solidity: function assetHashes(uint256 ) view returns(string)
func (_AssetsV1 *AssetsV1CallerSession) AssetHashes(arg0 *big.Int) (string, error) {
	return _AssetsV1.Contract.AssetHashes(&_AssetsV1.CallOpts, arg0)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_AssetsV1 *AssetsV1Caller) CreatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_AssetsV1 *AssetsV1Session) CreatedAt() (*big.Int, error) {
	return _AssetsV1.Contract.CreatedAt(&_AssetsV1.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_AssetsV1 *AssetsV1CallerSession) CreatedAt() (*big.Int, error) {
	return _AssetsV1.Contract.CreatedAt(&_AssetsV1.CallOpts)
}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_AssetsV1 *AssetsV1Caller) GetAssetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "getAssetCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_AssetsV1 *AssetsV1Session) GetAssetCount() (*big.Int, error) {
	return _AssetsV1.Contract.GetAssetCount(&_AssetsV1.CallOpts)
}

// GetAssetCount is a free data retrieval call binding the contract method 0xa0aead4d.
//
// Solidity: function getAssetCount() view returns(uint256)
func (_AssetsV1 *AssetsV1CallerSession) GetAssetCount() (*big.Int, error) {
	return _AssetsV1.Contract.GetAssetCount(&_AssetsV1.CallOpts)
}

// GetAssetCountForHash is a free data retrieval call binding the contract method 0xbf91e82f.
//
// Solidity: function getAssetCountForHash(string hash) view returns(uint256)
func (_AssetsV1 *AssetsV1Caller) GetAssetCountForHash(opts *bind.CallOpts, hash string) (*big.Int, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "getAssetCountForHash", hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetCountForHash is a free data retrieval call binding the contract method 0xbf91e82f.
//
// Solidity: function getAssetCountForHash(string hash) view returns(uint256)
func (_AssetsV1 *AssetsV1Session) GetAssetCountForHash(hash string) (*big.Int, error) {
	return _AssetsV1.Contract.GetAssetCountForHash(&_AssetsV1.CallOpts, hash)
}

// GetAssetCountForHash is a free data retrieval call binding the contract method 0xbf91e82f.
//
// Solidity: function getAssetCountForHash(string hash) view returns(uint256)
func (_AssetsV1 *AssetsV1CallerSession) GetAssetCountForHash(hash string) (*big.Int, error) {
	return _AssetsV1.Contract.GetAssetCountForHash(&_AssetsV1.CallOpts, hash)
}

// GetPublisherByAddress is a free data retrieval call binding the contract method 0x38052e27.
//
// Solidity: function getPublisherByAddress(address publicKey) view returns(address, int256, uint256)
func (_AssetsV1 *AssetsV1Caller) GetPublisherByAddress(opts *bind.CallOpts, publicKey common.Address) (common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "getPublisherByAddress", publicKey)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPublisherByAddress is a free data retrieval call binding the contract method 0x38052e27.
//
// Solidity: function getPublisherByAddress(address publicKey) view returns(address, int256, uint256)
func (_AssetsV1 *AssetsV1Session) GetPublisherByAddress(publicKey common.Address) (common.Address, *big.Int, *big.Int, error) {
	return _AssetsV1.Contract.GetPublisherByAddress(&_AssetsV1.CallOpts, publicKey)
}

// GetPublisherByAddress is a free data retrieval call binding the contract method 0x38052e27.
//
// Solidity: function getPublisherByAddress(address publicKey) view returns(address, int256, uint256)
func (_AssetsV1 *AssetsV1CallerSession) GetPublisherByAddress(publicKey common.Address) (common.Address, *big.Int, *big.Int, error) {
	return _AssetsV1.Contract.GetPublisherByAddress(&_AssetsV1.CallOpts, publicKey)
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(address[])
func (_AssetsV1 *AssetsV1Caller) GetPublishers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "getPublishers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(address[])
func (_AssetsV1 *AssetsV1Session) GetPublishers() ([]common.Address, error) {
	return _AssetsV1.Contract.GetPublishers(&_AssetsV1.CallOpts)
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(address[])
func (_AssetsV1 *AssetsV1CallerSession) GetPublishers() ([]common.Address, error) {
	return _AssetsV1.Contract.GetPublishers(&_AssetsV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetsV1 *AssetsV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetsV1 *AssetsV1Session) Owner() (common.Address, error) {
	return _AssetsV1.Contract.Owner(&_AssetsV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssetsV1 *AssetsV1CallerSession) Owner() (common.Address, error) {
	return _AssetsV1.Contract.Owner(&_AssetsV1.CallOpts)
}

// PublisherPublicKeys is a free data retrieval call binding the contract method 0x2c5f00a1.
//
// Solidity: function publisherPublicKeys(uint256 ) view returns(address)
func (_AssetsV1 *AssetsV1Caller) PublisherPublicKeys(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "publisherPublicKeys", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PublisherPublicKeys is a free data retrieval call binding the contract method 0x2c5f00a1.
//
// Solidity: function publisherPublicKeys(uint256 ) view returns(address)
func (_AssetsV1 *AssetsV1Session) PublisherPublicKeys(arg0 *big.Int) (common.Address, error) {
	return _AssetsV1.Contract.PublisherPublicKeys(&_AssetsV1.CallOpts, arg0)
}

// PublisherPublicKeys is a free data retrieval call binding the contract method 0x2c5f00a1.
//
// Solidity: function publisherPublicKeys(uint256 ) view returns(address)
func (_AssetsV1 *AssetsV1CallerSession) PublisherPublicKeys(arg0 *big.Int) (common.Address, error) {
	return _AssetsV1.Contract.PublisherPublicKeys(&_AssetsV1.CallOpts, arg0)
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string hash) view returns(address, int256, int256, uint256)
func (_AssetsV1 *AssetsV1Caller) Verify(opts *bind.CallOpts, hash string) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "verify", hash)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string hash) view returns(address, int256, int256, uint256)
func (_AssetsV1 *AssetsV1Session) Verify(hash string) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsV1.Contract.Verify(&_AssetsV1.CallOpts, hash)
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string hash) view returns(address, int256, int256, uint256)
func (_AssetsV1 *AssetsV1CallerSession) Verify(hash string) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsV1.Contract.Verify(&_AssetsV1.CallOpts, hash)
}

// VerifyByIndex is a free data retrieval call binding the contract method 0xd6ce25a9.
//
// Solidity: function verifyByIndex(string hash, uint256 assetIndex) view returns(address, int256, int256, uint256)
func (_AssetsV1 *AssetsV1Caller) VerifyByIndex(opts *bind.CallOpts, hash string, assetIndex *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AssetsV1.contract.Call(opts, &out, "verifyByIndex", hash, assetIndex)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// VerifyByIndex is a free data retrieval call binding the contract method 0xd6ce25a9.
//
// Solidity: function verifyByIndex(string hash, uint256 assetIndex) view returns(address, int256, int256, uint256)
func (_AssetsV1 *AssetsV1Session) VerifyByIndex(hash string, assetIndex *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsV1.Contract.VerifyByIndex(&_AssetsV1.CallOpts, hash, assetIndex)
}

// VerifyByIndex is a free data retrieval call binding the contract method 0xd6ce25a9.
//
// Solidity: function verifyByIndex(string hash, uint256 assetIndex) view returns(address, int256, int256, uint256)
func (_AssetsV1 *AssetsV1CallerSession) VerifyByIndex(hash string, assetIndex *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AssetsV1.Contract.VerifyByIndex(&_AssetsV1.CallOpts, hash, assetIndex)
}

// DisablePublisher is a paid mutator transaction binding the contract method 0x54df1eab.
//
// Solidity: function disablePublisher(address[] publicKeys) returns()
func (_AssetsV1 *AssetsV1Transactor) DisablePublisher(opts *bind.TransactOpts, publicKeys []common.Address) (*types.Transaction, error) {
	return _AssetsV1.contract.Transact(opts, "disablePublisher", publicKeys)
}

// DisablePublisher is a paid mutator transaction binding the contract method 0x54df1eab.
//
// Solidity: function disablePublisher(address[] publicKeys) returns()
func (_AssetsV1 *AssetsV1Session) DisablePublisher(publicKeys []common.Address) (*types.Transaction, error) {
	return _AssetsV1.Contract.DisablePublisher(&_AssetsV1.TransactOpts, publicKeys)
}

// DisablePublisher is a paid mutator transaction binding the contract method 0x54df1eab.
//
// Solidity: function disablePublisher(address[] publicKeys) returns()
func (_AssetsV1 *AssetsV1TransactorSession) DisablePublisher(publicKeys []common.Address) (*types.Transaction, error) {
	return _AssetsV1.Contract.DisablePublisher(&_AssetsV1.TransactOpts, publicKeys)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_AssetsV1 *AssetsV1Transactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetsV1.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_AssetsV1 *AssetsV1Session) Kill() (*types.Transaction, error) {
	return _AssetsV1.Contract.Kill(&_AssetsV1.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_AssetsV1 *AssetsV1TransactorSession) Kill() (*types.Transaction, error) {
	return _AssetsV1.Contract.Kill(&_AssetsV1.TransactOpts)
}

// SetPublisherLevel is a paid mutator transaction binding the contract method 0xc9e07f09.
//
// Solidity: function setPublisherLevel(address[] publicKeys, int256 level) returns()
func (_AssetsV1 *AssetsV1Transactor) SetPublisherLevel(opts *bind.TransactOpts, publicKeys []common.Address, level *big.Int) (*types.Transaction, error) {
	return _AssetsV1.contract.Transact(opts, "setPublisherLevel", publicKeys, level)
}

// SetPublisherLevel is a paid mutator transaction binding the contract method 0xc9e07f09.
//
// Solidity: function setPublisherLevel(address[] publicKeys, int256 level) returns()
func (_AssetsV1 *AssetsV1Session) SetPublisherLevel(publicKeys []common.Address, level *big.Int) (*types.Transaction, error) {
	return _AssetsV1.Contract.SetPublisherLevel(&_AssetsV1.TransactOpts, publicKeys, level)
}

// SetPublisherLevel is a paid mutator transaction binding the contract method 0xc9e07f09.
//
// Solidity: function setPublisherLevel(address[] publicKeys, int256 level) returns()
func (_AssetsV1 *AssetsV1TransactorSession) SetPublisherLevel(publicKeys []common.Address, level *big.Int) (*types.Transaction, error) {
	return _AssetsV1.Contract.SetPublisherLevel(&_AssetsV1.TransactOpts, publicKeys, level)
}

// Sign is a paid mutator transaction binding the contract method 0x28f7f9b8.
//
// Solidity: function sign(string hash, int256 status) returns()
func (_AssetsV1 *AssetsV1Transactor) Sign(opts *bind.TransactOpts, hash string, status *big.Int) (*types.Transaction, error) {
	return _AssetsV1.contract.Transact(opts, "sign", hash, status)
}

// Sign is a paid mutator transaction binding the contract method 0x28f7f9b8.
//
// Solidity: function sign(string hash, int256 status) returns()
func (_AssetsV1 *AssetsV1Session) Sign(hash string, status *big.Int) (*types.Transaction, error) {
	return _AssetsV1.Contract.Sign(&_AssetsV1.TransactOpts, hash, status)
}

// Sign is a paid mutator transaction binding the contract method 0x28f7f9b8.
//
// Solidity: function sign(string hash, int256 status) returns()
func (_AssetsV1 *AssetsV1TransactorSession) Sign(hash string, status *big.Int) (*types.Transaction, error) {
	return _AssetsV1.Contract.Sign(&_AssetsV1.TransactOpts, hash, status)
}

// AssetsV1AssetSignedIterator is returned from FilterAssetSigned and is used to iterate over the raw logs and unpacked data for AssetSigned events raised by the AssetsV1 contract.
type AssetsV1AssetSignedIterator struct {
	Event *AssetsV1AssetSigned // Event containing the contract specifics and raw log

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
func (it *AssetsV1AssetSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetsV1AssetSigned)
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
		it.Event = new(AssetsV1AssetSigned)
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
func (it *AssetsV1AssetSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetsV1AssetSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetsV1AssetSigned represents a AssetSigned event raised by the AssetsV1 contract.
type AssetsV1AssetSigned struct {
	PublicKey common.Address
	Hash      string
	Level     *big.Int
	Status    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssetSigned is a free log retrieval operation binding the contract event 0x5246d0df7f13afa7adb921dfc9e2762d77304794bae8e4734ef030f543fe837d.
//
// Solidity: event assetSigned(address publicKey, string hash, int256 level, int256 status, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) FilterAssetSigned(opts *bind.FilterOpts) (*AssetsV1AssetSignedIterator, error) {

	logs, sub, err := _AssetsV1.contract.FilterLogs(opts, "assetSigned")
	if err != nil {
		return nil, err
	}
	return &AssetsV1AssetSignedIterator{contract: _AssetsV1.contract, event: "assetSigned", logs: logs, sub: sub}, nil
}

// WatchAssetSigned is a free log subscription operation binding the contract event 0x5246d0df7f13afa7adb921dfc9e2762d77304794bae8e4734ef030f543fe837d.
//
// Solidity: event assetSigned(address publicKey, string hash, int256 level, int256 status, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) WatchAssetSigned(opts *bind.WatchOpts, sink chan<- *AssetsV1AssetSigned) (event.Subscription, error) {

	logs, sub, err := _AssetsV1.contract.WatchLogs(opts, "assetSigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetsV1AssetSigned)
				if err := _AssetsV1.contract.UnpackLog(event, "assetSigned", log); err != nil {
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

// ParseAssetSigned is a log parse operation binding the contract event 0x5246d0df7f13afa7adb921dfc9e2762d77304794bae8e4734ef030f543fe837d.
//
// Solidity: event assetSigned(address publicKey, string hash, int256 level, int256 status, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) ParseAssetSigned(log types.Log) (*AssetsV1AssetSigned, error) {
	event := new(AssetsV1AssetSigned)
	if err := _AssetsV1.contract.UnpackLog(event, "assetSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetsV1NewAssetIterator is returned from FilterNewAsset and is used to iterate over the raw logs and unpacked data for NewAsset events raised by the AssetsV1 contract.
type AssetsV1NewAssetIterator struct {
	Event *AssetsV1NewAsset // Event containing the contract specifics and raw log

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
func (it *AssetsV1NewAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetsV1NewAsset)
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
		it.Event = new(AssetsV1NewAsset)
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
func (it *AssetsV1NewAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetsV1NewAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetsV1NewAsset represents a NewAsset event raised by the AssetsV1 contract.
type AssetsV1NewAsset struct {
	Hash      string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewAsset is a free log retrieval operation binding the contract event 0x2bf03f94cf4fc3c57bcabb0cb88b62a4530d6a0de975f8f7a898bacde02e4d7a.
//
// Solidity: event newAsset(string hash, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) FilterNewAsset(opts *bind.FilterOpts) (*AssetsV1NewAssetIterator, error) {

	logs, sub, err := _AssetsV1.contract.FilterLogs(opts, "newAsset")
	if err != nil {
		return nil, err
	}
	return &AssetsV1NewAssetIterator{contract: _AssetsV1.contract, event: "newAsset", logs: logs, sub: sub}, nil
}

// WatchNewAsset is a free log subscription operation binding the contract event 0x2bf03f94cf4fc3c57bcabb0cb88b62a4530d6a0de975f8f7a898bacde02e4d7a.
//
// Solidity: event newAsset(string hash, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) WatchNewAsset(opts *bind.WatchOpts, sink chan<- *AssetsV1NewAsset) (event.Subscription, error) {

	logs, sub, err := _AssetsV1.contract.WatchLogs(opts, "newAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetsV1NewAsset)
				if err := _AssetsV1.contract.UnpackLog(event, "newAsset", log); err != nil {
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

// ParseNewAsset is a log parse operation binding the contract event 0x2bf03f94cf4fc3c57bcabb0cb88b62a4530d6a0de975f8f7a898bacde02e4d7a.
//
// Solidity: event newAsset(string hash, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) ParseNewAsset(log types.Log) (*AssetsV1NewAsset, error) {
	event := new(AssetsV1NewAsset)
	if err := _AssetsV1.contract.UnpackLog(event, "newAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetsV1NewPublisherIterator is returned from FilterNewPublisher and is used to iterate over the raw logs and unpacked data for NewPublisher events raised by the AssetsV1 contract.
type AssetsV1NewPublisherIterator struct {
	Event *AssetsV1NewPublisher // Event containing the contract specifics and raw log

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
func (it *AssetsV1NewPublisherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetsV1NewPublisher)
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
		it.Event = new(AssetsV1NewPublisher)
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
func (it *AssetsV1NewPublisherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetsV1NewPublisherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetsV1NewPublisher represents a NewPublisher event raised by the AssetsV1 contract.
type AssetsV1NewPublisher struct {
	PublicKey common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewPublisher is a free log retrieval operation binding the contract event 0x8c810fbf05dab8bcb3508516cffe5f356d0e2450a2828f6910c4fdd34ad625c3.
//
// Solidity: event newPublisher(address publicKey, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) FilterNewPublisher(opts *bind.FilterOpts) (*AssetsV1NewPublisherIterator, error) {

	logs, sub, err := _AssetsV1.contract.FilterLogs(opts, "newPublisher")
	if err != nil {
		return nil, err
	}
	return &AssetsV1NewPublisherIterator{contract: _AssetsV1.contract, event: "newPublisher", logs: logs, sub: sub}, nil
}

// WatchNewPublisher is a free log subscription operation binding the contract event 0x8c810fbf05dab8bcb3508516cffe5f356d0e2450a2828f6910c4fdd34ad625c3.
//
// Solidity: event newPublisher(address publicKey, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) WatchNewPublisher(opts *bind.WatchOpts, sink chan<- *AssetsV1NewPublisher) (event.Subscription, error) {

	logs, sub, err := _AssetsV1.contract.WatchLogs(opts, "newPublisher")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetsV1NewPublisher)
				if err := _AssetsV1.contract.UnpackLog(event, "newPublisher", log); err != nil {
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

// ParseNewPublisher is a log parse operation binding the contract event 0x8c810fbf05dab8bcb3508516cffe5f356d0e2450a2828f6910c4fdd34ad625c3.
//
// Solidity: event newPublisher(address publicKey, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) ParseNewPublisher(log types.Log) (*AssetsV1NewPublisher, error) {
	event := new(AssetsV1NewPublisher)
	if err := _AssetsV1.contract.UnpackLog(event, "newPublisher", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetsV1PublisherLevelChangedIterator is returned from FilterPublisherLevelChanged and is used to iterate over the raw logs and unpacked data for PublisherLevelChanged events raised by the AssetsV1 contract.
type AssetsV1PublisherLevelChangedIterator struct {
	Event *AssetsV1PublisherLevelChanged // Event containing the contract specifics and raw log

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
func (it *AssetsV1PublisherLevelChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetsV1PublisherLevelChanged)
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
		it.Event = new(AssetsV1PublisherLevelChanged)
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
func (it *AssetsV1PublisherLevelChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetsV1PublisherLevelChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetsV1PublisherLevelChanged represents a PublisherLevelChanged event raised by the AssetsV1 contract.
type AssetsV1PublisherLevelChanged struct {
	PublicKey common.Address
	Level     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublisherLevelChanged is a free log retrieval operation binding the contract event 0x86ec446c5d262f7d5d58e56a50600a6fa1a1db5678cb6d548e5383e65a3853f1.
//
// Solidity: event publisherLevelChanged(address publicKey, int256 level, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) FilterPublisherLevelChanged(opts *bind.FilterOpts) (*AssetsV1PublisherLevelChangedIterator, error) {

	logs, sub, err := _AssetsV1.contract.FilterLogs(opts, "publisherLevelChanged")
	if err != nil {
		return nil, err
	}
	return &AssetsV1PublisherLevelChangedIterator{contract: _AssetsV1.contract, event: "publisherLevelChanged", logs: logs, sub: sub}, nil
}

// WatchPublisherLevelChanged is a free log subscription operation binding the contract event 0x86ec446c5d262f7d5d58e56a50600a6fa1a1db5678cb6d548e5383e65a3853f1.
//
// Solidity: event publisherLevelChanged(address publicKey, int256 level, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) WatchPublisherLevelChanged(opts *bind.WatchOpts, sink chan<- *AssetsV1PublisherLevelChanged) (event.Subscription, error) {

	logs, sub, err := _AssetsV1.contract.WatchLogs(opts, "publisherLevelChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetsV1PublisherLevelChanged)
				if err := _AssetsV1.contract.UnpackLog(event, "publisherLevelChanged", log); err != nil {
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

// ParsePublisherLevelChanged is a log parse operation binding the contract event 0x86ec446c5d262f7d5d58e56a50600a6fa1a1db5678cb6d548e5383e65a3853f1.
//
// Solidity: event publisherLevelChanged(address publicKey, int256 level, uint256 timestamp)
func (_AssetsV1 *AssetsV1Filterer) ParsePublisherLevelChanged(log types.Log) (*AssetsV1PublisherLevelChanged, error) {
	event := new(AssetsV1PublisherLevelChanged)
	if err := _AssetsV1.contract.UnpackLog(event, "publisherLevelChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
