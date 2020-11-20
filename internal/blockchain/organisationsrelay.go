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

// OrganisationsRelayABI is the input ABI used to generate the binding from.
const OrganisationsRelayABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"organisationsContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getOrganisationByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getOrganisationByHash\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"removeOrganisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getOrganisation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrganisationCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createdAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"members\",\"type\":\"address[]\"}],\"name\":\"setMembers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"organisationOwner\",\"type\":\"address\"}],\"name\":\"changeOrganisationOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oldId\",\"type\":\"string\"},{\"name\":\"newId\",\"type\":\"string\"}],\"name\":\"renameOrganisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"organisationOwner\",\"type\":\"address\"}],\"name\":\"addOrganisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"oContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OrganisationsRelayFuncSigs maps the 4-byte function signature to its string representation.
var OrganisationsRelayFuncSigs = map[string]string{
	"fe794b00": "addOrganisation(string,address)",
	"f26937b5": "changeOrganisationOwner(string,address)",
	"cf09e0d0": "createdAt()",
	"50a38744": "getOrganisation(string)",
	"06e80cd9": "getOrganisationByHash(bytes32)",
	"04fc1902": "getOrganisationByIndex(uint256)",
	"b082b9c7": "getOrganisationCount()",
	"03a4692a": "organisationsContract()",
	"8da5cb5b": "owner()",
	"13ed4cb6": "removeOrganisation(string)",
	"f4e1bb87": "renameOrganisation(string,string)",
	"d317f35f": "setMembers(string,address[])",
}

// OrganisationsRelayBin is the compiled bytecode used for deploying new contracts.
var OrganisationsRelayBin = "0x608060405234801561001057600080fd5b50604051602080610b50833981016040525160028054600160a060020a03909216600160a060020a0319928316179055600080549091163217905542600155610af28061005e6000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166303a4692a81146100be57806304fc1902146100ef57806306e80cd9146101e457806313ed4cb6146101fc57806350a387441461021e5780638da5cb5b1461023e578063b082b9c714610253578063cf09e0d01461027a578063d317f35f1461028f578063f26937b5146102bb578063f4e1bb87146102e9578063fe794b0014610315575b600080fd5b3480156100ca57600080fd5b506100d3610343565b60408051600160a060020a039092168252519081900360200190f35b3480156100fb57600080fd5b50610107600435610352565b6040518085600160a060020a0316600160a060020a031681526020018060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b83811015610169578181015183820152602001610151565b50505050905001838103825285818151815260200191508051906020019080838360005b838110156101a557818101518382015260200161018d565b50505050905090810190601f1680156101d25780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156101f057600080fd5b506101076004356104b0565b34801561020857600080fd5b5061021c600480356024810191013561051c565b005b34801561022a57600080fd5b5061010760048035602481019101356105ca565b34801561024a57600080fd5b506100d3610744565b34801561025f57600080fd5b50610268610753565b60408051918252519081900360200190f35b34801561028657600080fd5b506102686107e3565b34801561029b57600080fd5b5061021c60246004803582810192908201359181359182019101356107e9565b3480156102c757600080fd5b5061021c6024600480358281019291013590600160a060020a039035166108bf565b3480156102f557600080fd5b5061021c6024600480358281019290820135918135918201910135610978565b34801561032157600080fd5b5061021c6024600480358281019291013590600160a060020a03903516610a2b565b600254600160a060020a031681565b600254604080517f04fc190200000000000000000000000000000000000000000000000000000000815260048101849052905160009260609283928592600160a060020a0316916304fc1902916024808301928692919082900301818387803b1580156103be57600080fd5b505af11580156103d2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260808110156103fb57600080fd5b81516020830180519193928301929164010000000081111561041c57600080fd5b8201602081018481111561042f57600080fd5b815185602082028301116401000000008211171561044c57600080fd5b5050929190602001805164010000000081111561046857600080fd5b8201602081018481111561047b57600080fd5b815164010000000081118282018710171561049557600080fd5b505060209190910151949a9399509750929550909350505050565b600254604080517f06e80cd900000000000000000000000000000000000000000000000000000000815260048101849052905160009260609283928592600160a060020a0316916306e80cd9916024808301928692919082900301818387803b1580156103be57600080fd5b600054600160a060020a03163214156105c6576002546040517f13ed4cb600000000000000000000000000000000000000000000000000000000815260206004820190815260248201849052600160a060020a03909216916313ed4cb69185918591819060440184848082843782019150509350505050600060405180830381600087803b1580156105ad57600080fd5b505af11580156105c1573d6000803e3d6000fd5b505050505b5050565b6002546040517f50a387440000000000000000000000000000000000000000000000000000000081526020600482019081526024820184905260009260609283928592600160a060020a0316916350a38744918991899190819060440184848082843782019150509350505050600060405180830381600087803b15801561065157600080fd5b505af1158015610665573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052608081101561068e57600080fd5b8151602083018051919392830192916401000000008111156106af57600080fd5b820160208101848111156106c257600080fd5b81518560208202830111640100000000821117156106df57600080fd5b505092919060200180516401000000008111156106fb57600080fd5b8201602081018481111561070e57600080fd5b815164010000000081118282018710171561072857600080fd5b505060209190910151949b939a50985092965090945050505050565b600054600160a060020a031681565b600254604080517fb082b9c70000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163b082b9c791600480830192602092919082900301818787803b1580156107b257600080fd5b505af11580156107c6573d6000803e3d6000fd5b505050506040513d60208110156107dc57600080fd5b5051905090565b60015481565b600054600160a060020a03163214156108b957600254604080517fd317f35f0000000000000000000000000000000000000000000000000000000081526004810191825260448101869052600160a060020a039092169163d317f35f9187918791879187919081906024810190606401878780828437909101848103835285815260209081019150869086028082843782019150509650505050505050600060405180830381600087803b1580156108a057600080fd5b505af11580156108b4573d6000803e3d6000fd5b505050505b50505050565b600054600160a060020a031632141561097357600254604080517ff26937b5000000000000000000000000000000000000000000000000000000008152600160a060020a03848116602483015260048201928352604482018690529092169163f26937b59186918691869181906064018585808284378201915050945050505050600060405180830381600087803b15801561095a57600080fd5b505af115801561096e573d6000803e3d6000fd5b505050505b505050565b600054600160a060020a03163214156108b957600254604080517ff4e1bb870000000000000000000000000000000000000000000000000000000081526004810191825260448101869052600160a060020a039092169163f4e1bb8791879187918791879190819060248101906064018787808284379091018481038352858152602001905085858082843782019150509650505050505050600060405180830381600087803b1580156108a057600080fd5b600054600160a060020a031632141561097357600254604080517ffe794b00000000000000000000000000000000000000000000000000000000008152600160a060020a03848116602483015260048201928352604482018690529092169163fe794b009186918691869181906064018585808284378201915050945050505050600060405180830381600087803b15801561095a57600080fd00a165627a7a72305820048edd1af7148282552c7b52f6505fb319ba652a5dfb27c9ccc1b8d8680800330029"

// DeployOrganisationsRelay deploys a new Ethereum contract, binding an instance of OrganisationsRelay to it.
func DeployOrganisationsRelay(auth *bind.TransactOpts, backend bind.ContractBackend, oContract common.Address) (common.Address, *types.Transaction, *OrganisationsRelay, error) {
	parsed, err := abi.JSON(strings.NewReader(OrganisationsRelayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrganisationsRelayBin), backend, oContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrganisationsRelay{OrganisationsRelayCaller: OrganisationsRelayCaller{contract: contract}, OrganisationsRelayTransactor: OrganisationsRelayTransactor{contract: contract}, OrganisationsRelayFilterer: OrganisationsRelayFilterer{contract: contract}}, nil
}

// OrganisationsRelay is an auto generated Go binding around an Ethereum contract.
type OrganisationsRelay struct {
	OrganisationsRelayCaller     // Read-only binding to the contract
	OrganisationsRelayTransactor // Write-only binding to the contract
	OrganisationsRelayFilterer   // Log filterer for contract events
}

// OrganisationsRelayCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrganisationsRelayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsRelayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrganisationsRelayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsRelayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrganisationsRelayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsRelaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrganisationsRelaySession struct {
	Contract     *OrganisationsRelay // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OrganisationsRelayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrganisationsRelayCallerSession struct {
	Contract *OrganisationsRelayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OrganisationsRelayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrganisationsRelayTransactorSession struct {
	Contract     *OrganisationsRelayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OrganisationsRelayRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrganisationsRelayRaw struct {
	Contract *OrganisationsRelay // Generic contract binding to access the raw methods on
}

// OrganisationsRelayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrganisationsRelayCallerRaw struct {
	Contract *OrganisationsRelayCaller // Generic read-only contract binding to access the raw methods on
}

// OrganisationsRelayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrganisationsRelayTransactorRaw struct {
	Contract *OrganisationsRelayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrganisationsRelay creates a new instance of OrganisationsRelay, bound to a specific deployed contract.
func NewOrganisationsRelay(address common.Address, backend bind.ContractBackend) (*OrganisationsRelay, error) {
	contract, err := bindOrganisationsRelay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrganisationsRelay{OrganisationsRelayCaller: OrganisationsRelayCaller{contract: contract}, OrganisationsRelayTransactor: OrganisationsRelayTransactor{contract: contract}, OrganisationsRelayFilterer: OrganisationsRelayFilterer{contract: contract}}, nil
}

// NewOrganisationsRelayCaller creates a new read-only instance of OrganisationsRelay, bound to a specific deployed contract.
func NewOrganisationsRelayCaller(address common.Address, caller bind.ContractCaller) (*OrganisationsRelayCaller, error) {
	contract, err := bindOrganisationsRelay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationsRelayCaller{contract: contract}, nil
}

// NewOrganisationsRelayTransactor creates a new write-only instance of OrganisationsRelay, bound to a specific deployed contract.
func NewOrganisationsRelayTransactor(address common.Address, transactor bind.ContractTransactor) (*OrganisationsRelayTransactor, error) {
	contract, err := bindOrganisationsRelay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationsRelayTransactor{contract: contract}, nil
}

// NewOrganisationsRelayFilterer creates a new log filterer instance of OrganisationsRelay, bound to a specific deployed contract.
func NewOrganisationsRelayFilterer(address common.Address, filterer bind.ContractFilterer) (*OrganisationsRelayFilterer, error) {
	contract, err := bindOrganisationsRelay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrganisationsRelayFilterer{contract: contract}, nil
}

// bindOrganisationsRelay binds a generic wrapper to an already deployed contract.
func bindOrganisationsRelay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrganisationsRelayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganisationsRelay *OrganisationsRelayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrganisationsRelay.Contract.OrganisationsRelayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganisationsRelay *OrganisationsRelayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.OrganisationsRelayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganisationsRelay *OrganisationsRelayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.OrganisationsRelayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganisationsRelay *OrganisationsRelayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrganisationsRelay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganisationsRelay *OrganisationsRelayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganisationsRelay *OrganisationsRelayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.contract.Transact(opts, method, params...)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_OrganisationsRelay *OrganisationsRelayCaller) CreatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrganisationsRelay.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_OrganisationsRelay *OrganisationsRelaySession) CreatedAt() (*big.Int, error) {
	return _OrganisationsRelay.Contract.CreatedAt(&_OrganisationsRelay.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_OrganisationsRelay *OrganisationsRelayCallerSession) CreatedAt() (*big.Int, error) {
	return _OrganisationsRelay.Contract.CreatedAt(&_OrganisationsRelay.CallOpts)
}

// GetOrganisation is a free data retrieval call binding the contract method 0x50a38744.
//
// Solidity: function getOrganisation(string id) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelayCaller) GetOrganisation(opts *bind.CallOpts, id string) (common.Address, []common.Address, string, *big.Int, error) {
	var out []interface{}
	err := _OrganisationsRelay.contract.Call(opts, &out, "getOrganisation", id)

	if err != nil {
		return *new(common.Address), *new([]common.Address), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetOrganisation is a free data retrieval call binding the contract method 0x50a38744.
//
// Solidity: function getOrganisation(string id) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelaySession) GetOrganisation(id string) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsRelay.Contract.GetOrganisation(&_OrganisationsRelay.CallOpts, id)
}

// GetOrganisation is a free data retrieval call binding the contract method 0x50a38744.
//
// Solidity: function getOrganisation(string id) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelayCallerSession) GetOrganisation(id string) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsRelay.Contract.GetOrganisation(&_OrganisationsRelay.CallOpts, id)
}

// GetOrganisationByHash is a free data retrieval call binding the contract method 0x06e80cd9.
//
// Solidity: function getOrganisationByHash(bytes32 hash) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelayCaller) GetOrganisationByHash(opts *bind.CallOpts, hash [32]byte) (common.Address, []common.Address, string, *big.Int, error) {
	var out []interface{}
	err := _OrganisationsRelay.contract.Call(opts, &out, "getOrganisationByHash", hash)

	if err != nil {
		return *new(common.Address), *new([]common.Address), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetOrganisationByHash is a free data retrieval call binding the contract method 0x06e80cd9.
//
// Solidity: function getOrganisationByHash(bytes32 hash) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelaySession) GetOrganisationByHash(hash [32]byte) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsRelay.Contract.GetOrganisationByHash(&_OrganisationsRelay.CallOpts, hash)
}

// GetOrganisationByHash is a free data retrieval call binding the contract method 0x06e80cd9.
//
// Solidity: function getOrganisationByHash(bytes32 hash) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelayCallerSession) GetOrganisationByHash(hash [32]byte) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsRelay.Contract.GetOrganisationByHash(&_OrganisationsRelay.CallOpts, hash)
}

// GetOrganisationByIndex is a free data retrieval call binding the contract method 0x04fc1902.
//
// Solidity: function getOrganisationByIndex(uint256 index) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelayCaller) GetOrganisationByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, []common.Address, string, *big.Int, error) {
	var out []interface{}
	err := _OrganisationsRelay.contract.Call(opts, &out, "getOrganisationByIndex", index)

	if err != nil {
		return *new(common.Address), *new([]common.Address), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetOrganisationByIndex is a free data retrieval call binding the contract method 0x04fc1902.
//
// Solidity: function getOrganisationByIndex(uint256 index) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelaySession) GetOrganisationByIndex(index *big.Int) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsRelay.Contract.GetOrganisationByIndex(&_OrganisationsRelay.CallOpts, index)
}

// GetOrganisationByIndex is a free data retrieval call binding the contract method 0x04fc1902.
//
// Solidity: function getOrganisationByIndex(uint256 index) view returns(address, address[], string, uint256)
func (_OrganisationsRelay *OrganisationsRelayCallerSession) GetOrganisationByIndex(index *big.Int) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsRelay.Contract.GetOrganisationByIndex(&_OrganisationsRelay.CallOpts, index)
}

// GetOrganisationCount is a free data retrieval call binding the contract method 0xb082b9c7.
//
// Solidity: function getOrganisationCount() view returns(uint256)
func (_OrganisationsRelay *OrganisationsRelayCaller) GetOrganisationCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrganisationsRelay.contract.Call(opts, &out, "getOrganisationCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrganisationCount is a free data retrieval call binding the contract method 0xb082b9c7.
//
// Solidity: function getOrganisationCount() view returns(uint256)
func (_OrganisationsRelay *OrganisationsRelaySession) GetOrganisationCount() (*big.Int, error) {
	return _OrganisationsRelay.Contract.GetOrganisationCount(&_OrganisationsRelay.CallOpts)
}

// GetOrganisationCount is a free data retrieval call binding the contract method 0xb082b9c7.
//
// Solidity: function getOrganisationCount() view returns(uint256)
func (_OrganisationsRelay *OrganisationsRelayCallerSession) GetOrganisationCount() (*big.Int, error) {
	return _OrganisationsRelay.Contract.GetOrganisationCount(&_OrganisationsRelay.CallOpts)
}

// OrganisationsContract is a free data retrieval call binding the contract method 0x03a4692a.
//
// Solidity: function organisationsContract() view returns(address)
func (_OrganisationsRelay *OrganisationsRelayCaller) OrganisationsContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrganisationsRelay.contract.Call(opts, &out, "organisationsContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrganisationsContract is a free data retrieval call binding the contract method 0x03a4692a.
//
// Solidity: function organisationsContract() view returns(address)
func (_OrganisationsRelay *OrganisationsRelaySession) OrganisationsContract() (common.Address, error) {
	return _OrganisationsRelay.Contract.OrganisationsContract(&_OrganisationsRelay.CallOpts)
}

// OrganisationsContract is a free data retrieval call binding the contract method 0x03a4692a.
//
// Solidity: function organisationsContract() view returns(address)
func (_OrganisationsRelay *OrganisationsRelayCallerSession) OrganisationsContract() (common.Address, error) {
	return _OrganisationsRelay.Contract.OrganisationsContract(&_OrganisationsRelay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganisationsRelay *OrganisationsRelayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrganisationsRelay.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganisationsRelay *OrganisationsRelaySession) Owner() (common.Address, error) {
	return _OrganisationsRelay.Contract.Owner(&_OrganisationsRelay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganisationsRelay *OrganisationsRelayCallerSession) Owner() (common.Address, error) {
	return _OrganisationsRelay.Contract.Owner(&_OrganisationsRelay.CallOpts)
}

// AddOrganisation is a paid mutator transaction binding the contract method 0xfe794b00.
//
// Solidity: function addOrganisation(string id, address organisationOwner) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactor) AddOrganisation(opts *bind.TransactOpts, id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.contract.Transact(opts, "addOrganisation", id, organisationOwner)
}

// AddOrganisation is a paid mutator transaction binding the contract method 0xfe794b00.
//
// Solidity: function addOrganisation(string id, address organisationOwner) returns()
func (_OrganisationsRelay *OrganisationsRelaySession) AddOrganisation(id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.AddOrganisation(&_OrganisationsRelay.TransactOpts, id, organisationOwner)
}

// AddOrganisation is a paid mutator transaction binding the contract method 0xfe794b00.
//
// Solidity: function addOrganisation(string id, address organisationOwner) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactorSession) AddOrganisation(id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.AddOrganisation(&_OrganisationsRelay.TransactOpts, id, organisationOwner)
}

// ChangeOrganisationOwner is a paid mutator transaction binding the contract method 0xf26937b5.
//
// Solidity: function changeOrganisationOwner(string id, address organisationOwner) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactor) ChangeOrganisationOwner(opts *bind.TransactOpts, id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.contract.Transact(opts, "changeOrganisationOwner", id, organisationOwner)
}

// ChangeOrganisationOwner is a paid mutator transaction binding the contract method 0xf26937b5.
//
// Solidity: function changeOrganisationOwner(string id, address organisationOwner) returns()
func (_OrganisationsRelay *OrganisationsRelaySession) ChangeOrganisationOwner(id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.ChangeOrganisationOwner(&_OrganisationsRelay.TransactOpts, id, organisationOwner)
}

// ChangeOrganisationOwner is a paid mutator transaction binding the contract method 0xf26937b5.
//
// Solidity: function changeOrganisationOwner(string id, address organisationOwner) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactorSession) ChangeOrganisationOwner(id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.ChangeOrganisationOwner(&_OrganisationsRelay.TransactOpts, id, organisationOwner)
}

// RemoveOrganisation is a paid mutator transaction binding the contract method 0x13ed4cb6.
//
// Solidity: function removeOrganisation(string id) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactor) RemoveOrganisation(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _OrganisationsRelay.contract.Transact(opts, "removeOrganisation", id)
}

// RemoveOrganisation is a paid mutator transaction binding the contract method 0x13ed4cb6.
//
// Solidity: function removeOrganisation(string id) returns()
func (_OrganisationsRelay *OrganisationsRelaySession) RemoveOrganisation(id string) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.RemoveOrganisation(&_OrganisationsRelay.TransactOpts, id)
}

// RemoveOrganisation is a paid mutator transaction binding the contract method 0x13ed4cb6.
//
// Solidity: function removeOrganisation(string id) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactorSession) RemoveOrganisation(id string) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.RemoveOrganisation(&_OrganisationsRelay.TransactOpts, id)
}

// RenameOrganisation is a paid mutator transaction binding the contract method 0xf4e1bb87.
//
// Solidity: function renameOrganisation(string oldId, string newId) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactor) RenameOrganisation(opts *bind.TransactOpts, oldId string, newId string) (*types.Transaction, error) {
	return _OrganisationsRelay.contract.Transact(opts, "renameOrganisation", oldId, newId)
}

// RenameOrganisation is a paid mutator transaction binding the contract method 0xf4e1bb87.
//
// Solidity: function renameOrganisation(string oldId, string newId) returns()
func (_OrganisationsRelay *OrganisationsRelaySession) RenameOrganisation(oldId string, newId string) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.RenameOrganisation(&_OrganisationsRelay.TransactOpts, oldId, newId)
}

// RenameOrganisation is a paid mutator transaction binding the contract method 0xf4e1bb87.
//
// Solidity: function renameOrganisation(string oldId, string newId) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactorSession) RenameOrganisation(oldId string, newId string) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.RenameOrganisation(&_OrganisationsRelay.TransactOpts, oldId, newId)
}

// SetMembers is a paid mutator transaction binding the contract method 0xd317f35f.
//
// Solidity: function setMembers(string id, address[] members) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactor) SetMembers(opts *bind.TransactOpts, id string, members []common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.contract.Transact(opts, "setMembers", id, members)
}

// SetMembers is a paid mutator transaction binding the contract method 0xd317f35f.
//
// Solidity: function setMembers(string id, address[] members) returns()
func (_OrganisationsRelay *OrganisationsRelaySession) SetMembers(id string, members []common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.SetMembers(&_OrganisationsRelay.TransactOpts, id, members)
}

// SetMembers is a paid mutator transaction binding the contract method 0xd317f35f.
//
// Solidity: function setMembers(string id, address[] members) returns()
func (_OrganisationsRelay *OrganisationsRelayTransactorSession) SetMembers(id string, members []common.Address) (*types.Transaction, error) {
	return _OrganisationsRelay.Contract.SetMembers(&_OrganisationsRelay.TransactOpts, id, members)
}

// OrganisationsV1ABI is the input ABI used to generate the binding from.
const OrganisationsV1ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getOrganisationByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getOrganisationByHash\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"removeOrganisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"organisationHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getOrganisation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrganisationCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createdAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"members\",\"type\":\"address[]\"}],\"name\":\"setMembers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"organisationOwner\",\"type\":\"address\"}],\"name\":\"changeOrganisationOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oldId\",\"type\":\"string\"},{\"name\":\"newId\",\"type\":\"string\"}],\"name\":\"renameOrganisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"organisationOwner\",\"type\":\"address\"}],\"name\":\"addOrganisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"organisationOwner\",\"type\":\"address\"}],\"name\":\"organisationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"organisationRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"newHash\",\"type\":\"bytes32\"}],\"name\":\"organisationRenamed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"organisationOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"oldMembers\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"newMembers\",\"type\":\"address[]\"}],\"name\":\"organisationMembersChanged\",\"type\":\"event\"}]"

// OrganisationsV1FuncSigs maps the 4-byte function signature to its string representation.
var OrganisationsV1FuncSigs = map[string]string{
	"fe794b00": "addOrganisation(string,address)",
	"f26937b5": "changeOrganisationOwner(string,address)",
	"cf09e0d0": "createdAt()",
	"50a38744": "getOrganisation(string)",
	"06e80cd9": "getOrganisationByHash(bytes32)",
	"04fc1902": "getOrganisationByIndex(uint256)",
	"b082b9c7": "getOrganisationCount()",
	"41c0e1b5": "kill()",
	"1afffdcb": "organisationHashes(uint256)",
	"8da5cb5b": "owner()",
	"13ed4cb6": "removeOrganisation(string)",
	"f4e1bb87": "renameOrganisation(string,string)",
	"d317f35f": "setMembers(string,address[])",
}

// OrganisationsV1Bin is the compiled bytecode used for deploying new contracts.
var OrganisationsV1Bin = "0x608060405234801561001057600080fd5b5060008054600160a060020a03191632179055426001556117e9806100366000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304fc190281146100c957806306e80cd9146101be57806313ed4cb6146101d65780631afffdcb146101f857806341c0e1b51461022257806350a38744146102375780638da5cb5b14610257578063b082b9c714610288578063cf09e0d01461029d578063d317f35f146102b2578063f26937b5146102de578063f4e1bb871461030c578063fe794b0014610338575b600080fd5b3480156100d557600080fd5b506100e1600435610366565b6040518085600160a060020a0316600160a060020a031681526020018060200180602001848152602001838103835286818151815260200191508051906020019060200280838360005b8381101561014357818101518382015260200161012b565b50505050905001838103825285818151815260200191508051906020019080838360005b8381101561017f578181015183820152602001610167565b50505050905090810190601f1680156101ac5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156101ca57600080fd5b506100e1600435610478565b3480156101e257600080fd5b506101f66004803560248101910135610577565b005b34801561020457600080fd5b5061021060043561073a565b60408051918252519081900360200190f35b34801561022e57600080fd5b506101f6610759565b34801561024357600080fd5b506100e160048035602481019101356107b7565b34801561026357600080fd5b5061026c6108f6565b60408051600160a060020a039092168252519081900360200190f35b34801561029457600080fd5b50610210610905565b3480156102a957600080fd5b5061021061090c565b3480156102be57600080fd5b506101f66024600480358281019290820135918135918201910135610912565b3480156102ea57600080fd5b506101f66024600480358281019291013590600160a060020a03903516610b95565b34801561031857600080fd5b506101f66024600480358281019290820135918135918201910135610d5e565b34801561034457600080fd5b506101f66024600480358281019291013590600160a060020a03903516611109565b6000606080600080610376611644565b60025460609088106103a757604080516000808252818301909252602081018281529198509650945086935061046e565b60028054899081106103b557fe5b6000918252602080832090910154808352600382526040808420815180830183528154600160a060020a031681526001909101548185015282855260048452938190208054825181860281018601909352808352929750939550929183018282801561044a57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161042c575b505050505090508160000151816104608561130f565b846020015196509650965096505b5050509193509193565b60006060806000610487611644565b5060008581526003602090815260409182902082518084019093528054600160a060020a03168084526001909101549183019190915260609015156104ec5781516020808401516040805192830190526000825291975091955090935091508361056e565b6000878152600460209081526040918290208054835181840281018401909452808452909183018282801561054a57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161052c575b505050505090508160000151816105608961130f565b846020015195509550955095505b50509193509193565b60008054600160a060020a031632146105c8576040805160e560020a62461bcd028152602060048201526013602482015260008051602061179e833981519152604482015290519081900360640190fd5b61060183838080601f01602080910402602001604051908101604052809392919081815260200183838082843750611446945050505050565b905060008211610656576040805160e560020a62461bcd028152602060048201526021602482015260008051602061177e833981519152604482015260f860020a607902606482015290519081900360840190fd5b600081815260036020526040902054600160a060020a031615156106c4576040805160e560020a62461bcd02815260206004820152601b60248201527f4f7267616e69736174696f6e20646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b60008181526003602090815260408083208054600160a060020a0319168155600101839055600490915281206106f99161165b565b61070281611513565b6040805182815290517f9f11f05fe2810e17e680075c9e99bca8f9acd92f043d2cad824ab94decc9e9789181900360200190a1505050565b600280548290811061074857fe5b600091825260209091200154905081565b600054600160a060020a031632146107a9576040805160e560020a62461bcd028152602060048201526013602482015260008051602061179e833981519152604482015290519081900360640190fd5b600054600160a060020a0316ff5b60006060806000806107c7611644565b606061080289898080601f01602080910402602001604051908101604052809392919081815260200183838082843750611446945050505050565b60008181526003602090815260409182902082518084019093528054600160a060020a03168084526001909101549183019190915291945092501515610868578151602080840151604080519283019052600082529198509196509094509250846108ea565b600083815260046020908152604091829020805483518184028101840190945280845290918301828280156108c657602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108a8575b505050505090508160000151816108dc8561130f565b846020015196509650965096505b50505092959194509250565b600054600160a060020a031681565b6002545b90565b60015481565b60008054606090600160a060020a03163214610966576040805160e560020a62461bcd028152602060048201526013602482015260008051602061179e833981519152604482015290519081900360640190fd5b61099f86868080601f01602080910402602001604051908101604052809392919081815260200183838082843750611446945050505050565b9150600085116109f4576040805160e560020a62461bcd028152602060048201526021602482015260008051602061177e833981519152604482015260f860020a607902606482015290519081900360840190fd5b600082815260036020526040902054600160a060020a03161515610a62576040805160e560020a62461bcd02815260206004820152601b60248201527f4f7267616e69736174696f6e20646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b60008281526004602090815260409182902080548351818402810184019094528084529091830182828015610ac057602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610aa2575b5050506000858152600460205260409020929350610ae39291508690508561167c565b507fd89a1da0f5f472e29dce6d5509d4b3c637e5f58daa235cf7a0b650e46a3051a8828286866040518085600019166000191681526020018060200180602001838103835286818151815260200191508051906020019060200280838360005b83811015610b5b578181015183820152602001610b43565b50505050905001838103825285858281815260200192506020028082843760405192018290039850909650505050505050a1505050505050565b600080548190600160a060020a03163214610be8576040805160e560020a62461bcd028152602060048201526013602482015260008051602061179e833981519152604482015290519081900360640190fd5b610c2185858080601f01602080910402602001604051908101604052809392919081815260200183838082843750611446945050505050565b915060008411610c76576040805160e560020a62461bcd028152602060048201526021602482015260008051602061177e833981519152604482015260f860020a607902606482015290519081900360840190fd5b600082815260036020526040902054600160a060020a03161515610ce4576040805160e560020a62461bcd02815260206004820152601b60248201527f4f7267616e69736174696f6e20646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b506000818152600360209081526040918290208054600160a060020a03868116600160a060020a0319831681179093558451868152911692810183905280840191909152915190917fe6ea8e5ace912ff2fa4c02f7bafe620691508973d49774daee6a08862f9614e8919081900360600190a15050505050565b600080548190600160a060020a03163214610db1576040805160e560020a62461bcd028152602060048201526013602482015260008051602061179e833981519152604482015290519081900360640190fd5b610dea86868080601f01602080910402602001604051908101604052809392919081815260200183838082843750611446945050505050565b9150610e2584848080601f01602080910402602001604051908101604052809392919081815260200183838082843750611446945050505050565b905060008511610ea5576040805160e560020a62461bcd02815260206004820152602560248201527f4f6c64206f7267616e69736174696f6e206964206d757374206e6f742062652060448201527f656d707479000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008311610f23576040805160e560020a62461bcd02815260206004820152602560248201527f4e6577206f7267616e69736174696f6e206964206d757374206e6f742062652060448201527f656d707479000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600082815260036020526040902054600160a060020a03161515610f91576040805160e560020a62461bcd02815260206004820152601f60248201527f4f6c64206f7267616e69736174696f6e20646f6573206e6f7420657869737400604482015290519081900360640190fd5b600081815260036020526040902054600160a060020a031615610ffe576040805160e560020a62461bcd02815260206004820152601f60248201527f4e6577206f7267616e69736174696f6e20616c72656164792065786973747300604482015290519081900360640190fd5b600082815260036020908152604080832084845281842081548154600160a060020a031916600160a060020a0390911617815560019182015491015584835260049091528082208383529120815461105692906116df565b506002805460018181019092557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0182905560008381526003602090815260408083208054600160a060020a0319168155909301829055600490529081206110bd9161165b565b6110c682611513565b604080518381526020810183905281517f83edd531f7a6ab02a34fcaea2a9f5922b2eda44c477002ad4eba9945a50ef55e929181900390910190a1505050505050565b60008054600160a060020a0316321461115a576040805160e560020a62461bcd028152602060048201526013602482015260008051602061179e833981519152604482015290519081900360640190fd5b61119384848080601f01602080910402602001604051908101604052809392919081815260200183838082843750611446945050505050565b9050600083116111e8576040805160e560020a62461bcd028152602060048201526021602482015260008051602061177e833981519152604482015260f860020a607902606482015290519081900360840190fd5b600081815260036020526040902054600160a060020a031615611255576040805160e560020a62461bcd02815260206004820152601b60248201527f4f7267616e69736174696f6e20616c7265616479206578697374730000000000604482015290519081900360640190fd5b6002805460018082019092557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01829055604080518082018252600160a060020a0385811680835242602080850191825260008881526003825286902094518554600160a060020a031916941693909317845551929094019190915581518481529081019290925280517f6101ecf9ce19dbf2bf40dd4b2c74280d8a76b327a5ea53c42857f9d14e5b83589281900390910190a150505050565b6040805181815260608082018352918291600091829182918291906020820161080080388339019050509450600093505b602084101561143b5786846020811061135557fe5b1a60f860020a02925060108360f860020a900460ff1681151561137457fe5b0460f860020a0291508160f860020a90046010028360f860020a90040360f860020a0290506113a2826115c8565b85856002028151811015156113b357fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506113ec816115c8565b858560020260010181518110151561140057fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600190930192611340565b509295945050505050565b6000816040516020018082805190602001908083835b6020831061147b5780518252601f19909201916020918201910161145c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106114de5780518252601f1990920191602091820191016114bf565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090505b919050565b6000805b60025482101561155757600280548491908490811061153257fe5b600091825260209091200154141561154c57506001611557565b600190910190611517565b80156115c3575b600254600019018210156115ae57600280546001840190811061157d57fe5b906000526020600020015460028381548110151561159757fe5b60009182526020909120015560019091019061155e565b60028054906115c190600019830161171f565b505b505050565b60007f0a000000000000000000000000000000000000000000000000000000000000007fff000000000000000000000000000000000000000000000000000000000000008316101561162c578160f860020a900460300160f860020a02905061150e565b8160f860020a900460570160f860020a02905061150e565b604080518082019091526000808252602082015290565b5080546000825590600052602060002090810190611679919061173f565b50565b8280548282559060005260206000209081019282156116cf579160200282015b828111156116cf578154600160a060020a031916600160a060020a0384351617825560209092019160019091019061169c565b506116db929150611759565b5090565b8280548282559060005260206000209081019282156116cf5760005260206000209182015b828111156116cf578254825591600101919060010190611704565b8154818355818111156115c3576000838152602090206115c39181019083015b61090991905b808211156116db5760008155600101611745565b61090991905b808211156116db578054600160a060020a031916815560010161175f56004f7267616e69736174696f6e206964206d757374206e6f7420626520656d70745265737472696374656420746f206f776e657200000000000000000000000000a165627a7a72305820763357f93cd5e789bc4332439a1445cf7ad5444a517596419285f8a4b23f172c0029"

// DeployOrganisationsV1 deploys a new Ethereum contract, binding an instance of OrganisationsV1 to it.
func DeployOrganisationsV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrganisationsV1, error) {
	parsed, err := abi.JSON(strings.NewReader(OrganisationsV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrganisationsV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrganisationsV1{OrganisationsV1Caller: OrganisationsV1Caller{contract: contract}, OrganisationsV1Transactor: OrganisationsV1Transactor{contract: contract}, OrganisationsV1Filterer: OrganisationsV1Filterer{contract: contract}}, nil
}

// OrganisationsV1 is an auto generated Go binding around an Ethereum contract.
type OrganisationsV1 struct {
	OrganisationsV1Caller     // Read-only binding to the contract
	OrganisationsV1Transactor // Write-only binding to the contract
	OrganisationsV1Filterer   // Log filterer for contract events
}

// OrganisationsV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type OrganisationsV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OrganisationsV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrganisationsV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrganisationsV1Session struct {
	Contract     *OrganisationsV1  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrganisationsV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrganisationsV1CallerSession struct {
	Contract *OrganisationsV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// OrganisationsV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrganisationsV1TransactorSession struct {
	Contract     *OrganisationsV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OrganisationsV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type OrganisationsV1Raw struct {
	Contract *OrganisationsV1 // Generic contract binding to access the raw methods on
}

// OrganisationsV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrganisationsV1CallerRaw struct {
	Contract *OrganisationsV1Caller // Generic read-only contract binding to access the raw methods on
}

// OrganisationsV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrganisationsV1TransactorRaw struct {
	Contract *OrganisationsV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOrganisationsV1 creates a new instance of OrganisationsV1, bound to a specific deployed contract.
func NewOrganisationsV1(address common.Address, backend bind.ContractBackend) (*OrganisationsV1, error) {
	contract, err := bindOrganisationsV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1{OrganisationsV1Caller: OrganisationsV1Caller{contract: contract}, OrganisationsV1Transactor: OrganisationsV1Transactor{contract: contract}, OrganisationsV1Filterer: OrganisationsV1Filterer{contract: contract}}, nil
}

// NewOrganisationsV1Caller creates a new read-only instance of OrganisationsV1, bound to a specific deployed contract.
func NewOrganisationsV1Caller(address common.Address, caller bind.ContractCaller) (*OrganisationsV1Caller, error) {
	contract, err := bindOrganisationsV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1Caller{contract: contract}, nil
}

// NewOrganisationsV1Transactor creates a new write-only instance of OrganisationsV1, bound to a specific deployed contract.
func NewOrganisationsV1Transactor(address common.Address, transactor bind.ContractTransactor) (*OrganisationsV1Transactor, error) {
	contract, err := bindOrganisationsV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1Transactor{contract: contract}, nil
}

// NewOrganisationsV1Filterer creates a new log filterer instance of OrganisationsV1, bound to a specific deployed contract.
func NewOrganisationsV1Filterer(address common.Address, filterer bind.ContractFilterer) (*OrganisationsV1Filterer, error) {
	contract, err := bindOrganisationsV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1Filterer{contract: contract}, nil
}

// bindOrganisationsV1 binds a generic wrapper to an already deployed contract.
func bindOrganisationsV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrganisationsV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganisationsV1 *OrganisationsV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrganisationsV1.Contract.OrganisationsV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganisationsV1 *OrganisationsV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.OrganisationsV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganisationsV1 *OrganisationsV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.OrganisationsV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganisationsV1 *OrganisationsV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrganisationsV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganisationsV1 *OrganisationsV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganisationsV1 *OrganisationsV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.contract.Transact(opts, method, params...)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_OrganisationsV1 *OrganisationsV1Caller) CreatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrganisationsV1.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_OrganisationsV1 *OrganisationsV1Session) CreatedAt() (*big.Int, error) {
	return _OrganisationsV1.Contract.CreatedAt(&_OrganisationsV1.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_OrganisationsV1 *OrganisationsV1CallerSession) CreatedAt() (*big.Int, error) {
	return _OrganisationsV1.Contract.CreatedAt(&_OrganisationsV1.CallOpts)
}

// GetOrganisation is a free data retrieval call binding the contract method 0x50a38744.
//
// Solidity: function getOrganisation(string id) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1Caller) GetOrganisation(opts *bind.CallOpts, id string) (common.Address, []common.Address, string, *big.Int, error) {
	var out []interface{}
	err := _OrganisationsV1.contract.Call(opts, &out, "getOrganisation", id)

	if err != nil {
		return *new(common.Address), *new([]common.Address), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetOrganisation is a free data retrieval call binding the contract method 0x50a38744.
//
// Solidity: function getOrganisation(string id) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1Session) GetOrganisation(id string) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsV1.Contract.GetOrganisation(&_OrganisationsV1.CallOpts, id)
}

// GetOrganisation is a free data retrieval call binding the contract method 0x50a38744.
//
// Solidity: function getOrganisation(string id) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1CallerSession) GetOrganisation(id string) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsV1.Contract.GetOrganisation(&_OrganisationsV1.CallOpts, id)
}

// GetOrganisationByHash is a free data retrieval call binding the contract method 0x06e80cd9.
//
// Solidity: function getOrganisationByHash(bytes32 hash) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1Caller) GetOrganisationByHash(opts *bind.CallOpts, hash [32]byte) (common.Address, []common.Address, string, *big.Int, error) {
	var out []interface{}
	err := _OrganisationsV1.contract.Call(opts, &out, "getOrganisationByHash", hash)

	if err != nil {
		return *new(common.Address), *new([]common.Address), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetOrganisationByHash is a free data retrieval call binding the contract method 0x06e80cd9.
//
// Solidity: function getOrganisationByHash(bytes32 hash) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1Session) GetOrganisationByHash(hash [32]byte) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsV1.Contract.GetOrganisationByHash(&_OrganisationsV1.CallOpts, hash)
}

// GetOrganisationByHash is a free data retrieval call binding the contract method 0x06e80cd9.
//
// Solidity: function getOrganisationByHash(bytes32 hash) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1CallerSession) GetOrganisationByHash(hash [32]byte) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsV1.Contract.GetOrganisationByHash(&_OrganisationsV1.CallOpts, hash)
}

// GetOrganisationByIndex is a free data retrieval call binding the contract method 0x04fc1902.
//
// Solidity: function getOrganisationByIndex(uint256 index) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1Caller) GetOrganisationByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, []common.Address, string, *big.Int, error) {
	var out []interface{}
	err := _OrganisationsV1.contract.Call(opts, &out, "getOrganisationByIndex", index)

	if err != nil {
		return *new(common.Address), *new([]common.Address), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetOrganisationByIndex is a free data retrieval call binding the contract method 0x04fc1902.
//
// Solidity: function getOrganisationByIndex(uint256 index) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1Session) GetOrganisationByIndex(index *big.Int) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsV1.Contract.GetOrganisationByIndex(&_OrganisationsV1.CallOpts, index)
}

// GetOrganisationByIndex is a free data retrieval call binding the contract method 0x04fc1902.
//
// Solidity: function getOrganisationByIndex(uint256 index) view returns(address, address[], string, uint256)
func (_OrganisationsV1 *OrganisationsV1CallerSession) GetOrganisationByIndex(index *big.Int) (common.Address, []common.Address, string, *big.Int, error) {
	return _OrganisationsV1.Contract.GetOrganisationByIndex(&_OrganisationsV1.CallOpts, index)
}

// GetOrganisationCount is a free data retrieval call binding the contract method 0xb082b9c7.
//
// Solidity: function getOrganisationCount() view returns(uint256)
func (_OrganisationsV1 *OrganisationsV1Caller) GetOrganisationCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrganisationsV1.contract.Call(opts, &out, "getOrganisationCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrganisationCount is a free data retrieval call binding the contract method 0xb082b9c7.
//
// Solidity: function getOrganisationCount() view returns(uint256)
func (_OrganisationsV1 *OrganisationsV1Session) GetOrganisationCount() (*big.Int, error) {
	return _OrganisationsV1.Contract.GetOrganisationCount(&_OrganisationsV1.CallOpts)
}

// GetOrganisationCount is a free data retrieval call binding the contract method 0xb082b9c7.
//
// Solidity: function getOrganisationCount() view returns(uint256)
func (_OrganisationsV1 *OrganisationsV1CallerSession) GetOrganisationCount() (*big.Int, error) {
	return _OrganisationsV1.Contract.GetOrganisationCount(&_OrganisationsV1.CallOpts)
}

// OrganisationHashes is a free data retrieval call binding the contract method 0x1afffdcb.
//
// Solidity: function organisationHashes(uint256 ) view returns(bytes32)
func (_OrganisationsV1 *OrganisationsV1Caller) OrganisationHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _OrganisationsV1.contract.Call(opts, &out, "organisationHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrganisationHashes is a free data retrieval call binding the contract method 0x1afffdcb.
//
// Solidity: function organisationHashes(uint256 ) view returns(bytes32)
func (_OrganisationsV1 *OrganisationsV1Session) OrganisationHashes(arg0 *big.Int) ([32]byte, error) {
	return _OrganisationsV1.Contract.OrganisationHashes(&_OrganisationsV1.CallOpts, arg0)
}

// OrganisationHashes is a free data retrieval call binding the contract method 0x1afffdcb.
//
// Solidity: function organisationHashes(uint256 ) view returns(bytes32)
func (_OrganisationsV1 *OrganisationsV1CallerSession) OrganisationHashes(arg0 *big.Int) ([32]byte, error) {
	return _OrganisationsV1.Contract.OrganisationHashes(&_OrganisationsV1.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganisationsV1 *OrganisationsV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrganisationsV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganisationsV1 *OrganisationsV1Session) Owner() (common.Address, error) {
	return _OrganisationsV1.Contract.Owner(&_OrganisationsV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrganisationsV1 *OrganisationsV1CallerSession) Owner() (common.Address, error) {
	return _OrganisationsV1.Contract.Owner(&_OrganisationsV1.CallOpts)
}

// AddOrganisation is a paid mutator transaction binding the contract method 0xfe794b00.
//
// Solidity: function addOrganisation(string id, address organisationOwner) returns()
func (_OrganisationsV1 *OrganisationsV1Transactor) AddOrganisation(opts *bind.TransactOpts, id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.contract.Transact(opts, "addOrganisation", id, organisationOwner)
}

// AddOrganisation is a paid mutator transaction binding the contract method 0xfe794b00.
//
// Solidity: function addOrganisation(string id, address organisationOwner) returns()
func (_OrganisationsV1 *OrganisationsV1Session) AddOrganisation(id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.AddOrganisation(&_OrganisationsV1.TransactOpts, id, organisationOwner)
}

// AddOrganisation is a paid mutator transaction binding the contract method 0xfe794b00.
//
// Solidity: function addOrganisation(string id, address organisationOwner) returns()
func (_OrganisationsV1 *OrganisationsV1TransactorSession) AddOrganisation(id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.AddOrganisation(&_OrganisationsV1.TransactOpts, id, organisationOwner)
}

// ChangeOrganisationOwner is a paid mutator transaction binding the contract method 0xf26937b5.
//
// Solidity: function changeOrganisationOwner(string id, address organisationOwner) returns()
func (_OrganisationsV1 *OrganisationsV1Transactor) ChangeOrganisationOwner(opts *bind.TransactOpts, id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.contract.Transact(opts, "changeOrganisationOwner", id, organisationOwner)
}

// ChangeOrganisationOwner is a paid mutator transaction binding the contract method 0xf26937b5.
//
// Solidity: function changeOrganisationOwner(string id, address organisationOwner) returns()
func (_OrganisationsV1 *OrganisationsV1Session) ChangeOrganisationOwner(id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.ChangeOrganisationOwner(&_OrganisationsV1.TransactOpts, id, organisationOwner)
}

// ChangeOrganisationOwner is a paid mutator transaction binding the contract method 0xf26937b5.
//
// Solidity: function changeOrganisationOwner(string id, address organisationOwner) returns()
func (_OrganisationsV1 *OrganisationsV1TransactorSession) ChangeOrganisationOwner(id string, organisationOwner common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.ChangeOrganisationOwner(&_OrganisationsV1.TransactOpts, id, organisationOwner)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_OrganisationsV1 *OrganisationsV1Transactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationsV1.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_OrganisationsV1 *OrganisationsV1Session) Kill() (*types.Transaction, error) {
	return _OrganisationsV1.Contract.Kill(&_OrganisationsV1.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_OrganisationsV1 *OrganisationsV1TransactorSession) Kill() (*types.Transaction, error) {
	return _OrganisationsV1.Contract.Kill(&_OrganisationsV1.TransactOpts)
}

// RemoveOrganisation is a paid mutator transaction binding the contract method 0x13ed4cb6.
//
// Solidity: function removeOrganisation(string id) returns()
func (_OrganisationsV1 *OrganisationsV1Transactor) RemoveOrganisation(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _OrganisationsV1.contract.Transact(opts, "removeOrganisation", id)
}

// RemoveOrganisation is a paid mutator transaction binding the contract method 0x13ed4cb6.
//
// Solidity: function removeOrganisation(string id) returns()
func (_OrganisationsV1 *OrganisationsV1Session) RemoveOrganisation(id string) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.RemoveOrganisation(&_OrganisationsV1.TransactOpts, id)
}

// RemoveOrganisation is a paid mutator transaction binding the contract method 0x13ed4cb6.
//
// Solidity: function removeOrganisation(string id) returns()
func (_OrganisationsV1 *OrganisationsV1TransactorSession) RemoveOrganisation(id string) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.RemoveOrganisation(&_OrganisationsV1.TransactOpts, id)
}

// RenameOrganisation is a paid mutator transaction binding the contract method 0xf4e1bb87.
//
// Solidity: function renameOrganisation(string oldId, string newId) returns()
func (_OrganisationsV1 *OrganisationsV1Transactor) RenameOrganisation(opts *bind.TransactOpts, oldId string, newId string) (*types.Transaction, error) {
	return _OrganisationsV1.contract.Transact(opts, "renameOrganisation", oldId, newId)
}

// RenameOrganisation is a paid mutator transaction binding the contract method 0xf4e1bb87.
//
// Solidity: function renameOrganisation(string oldId, string newId) returns()
func (_OrganisationsV1 *OrganisationsV1Session) RenameOrganisation(oldId string, newId string) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.RenameOrganisation(&_OrganisationsV1.TransactOpts, oldId, newId)
}

// RenameOrganisation is a paid mutator transaction binding the contract method 0xf4e1bb87.
//
// Solidity: function renameOrganisation(string oldId, string newId) returns()
func (_OrganisationsV1 *OrganisationsV1TransactorSession) RenameOrganisation(oldId string, newId string) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.RenameOrganisation(&_OrganisationsV1.TransactOpts, oldId, newId)
}

// SetMembers is a paid mutator transaction binding the contract method 0xd317f35f.
//
// Solidity: function setMembers(string id, address[] members) returns()
func (_OrganisationsV1 *OrganisationsV1Transactor) SetMembers(opts *bind.TransactOpts, id string, members []common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.contract.Transact(opts, "setMembers", id, members)
}

// SetMembers is a paid mutator transaction binding the contract method 0xd317f35f.
//
// Solidity: function setMembers(string id, address[] members) returns()
func (_OrganisationsV1 *OrganisationsV1Session) SetMembers(id string, members []common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.SetMembers(&_OrganisationsV1.TransactOpts, id, members)
}

// SetMembers is a paid mutator transaction binding the contract method 0xd317f35f.
//
// Solidity: function setMembers(string id, address[] members) returns()
func (_OrganisationsV1 *OrganisationsV1TransactorSession) SetMembers(id string, members []common.Address) (*types.Transaction, error) {
	return _OrganisationsV1.Contract.SetMembers(&_OrganisationsV1.TransactOpts, id, members)
}

// OrganisationsV1OrganisationAddedIterator is returned from FilterOrganisationAdded and is used to iterate over the raw logs and unpacked data for OrganisationAdded events raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationAddedIterator struct {
	Event *OrganisationsV1OrganisationAdded // Event containing the contract specifics and raw log

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
func (it *OrganisationsV1OrganisationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganisationsV1OrganisationAdded)
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
		it.Event = new(OrganisationsV1OrganisationAdded)
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
func (it *OrganisationsV1OrganisationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganisationsV1OrganisationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganisationsV1OrganisationAdded represents a OrganisationAdded event raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationAdded struct {
	Hash              [32]byte
	OrganisationOwner common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrganisationAdded is a free log retrieval operation binding the contract event 0x6101ecf9ce19dbf2bf40dd4b2c74280d8a76b327a5ea53c42857f9d14e5b8358.
//
// Solidity: event organisationAdded(bytes32 hash, address organisationOwner)
func (_OrganisationsV1 *OrganisationsV1Filterer) FilterOrganisationAdded(opts *bind.FilterOpts) (*OrganisationsV1OrganisationAddedIterator, error) {

	logs, sub, err := _OrganisationsV1.contract.FilterLogs(opts, "organisationAdded")
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1OrganisationAddedIterator{contract: _OrganisationsV1.contract, event: "organisationAdded", logs: logs, sub: sub}, nil
}

// WatchOrganisationAdded is a free log subscription operation binding the contract event 0x6101ecf9ce19dbf2bf40dd4b2c74280d8a76b327a5ea53c42857f9d14e5b8358.
//
// Solidity: event organisationAdded(bytes32 hash, address organisationOwner)
func (_OrganisationsV1 *OrganisationsV1Filterer) WatchOrganisationAdded(opts *bind.WatchOpts, sink chan<- *OrganisationsV1OrganisationAdded) (event.Subscription, error) {

	logs, sub, err := _OrganisationsV1.contract.WatchLogs(opts, "organisationAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganisationsV1OrganisationAdded)
				if err := _OrganisationsV1.contract.UnpackLog(event, "organisationAdded", log); err != nil {
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

// ParseOrganisationAdded is a log parse operation binding the contract event 0x6101ecf9ce19dbf2bf40dd4b2c74280d8a76b327a5ea53c42857f9d14e5b8358.
//
// Solidity: event organisationAdded(bytes32 hash, address organisationOwner)
func (_OrganisationsV1 *OrganisationsV1Filterer) ParseOrganisationAdded(log types.Log) (*OrganisationsV1OrganisationAdded, error) {
	event := new(OrganisationsV1OrganisationAdded)
	if err := _OrganisationsV1.contract.UnpackLog(event, "organisationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganisationsV1OrganisationMembersChangedIterator is returned from FilterOrganisationMembersChanged and is used to iterate over the raw logs and unpacked data for OrganisationMembersChanged events raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationMembersChangedIterator struct {
	Event *OrganisationsV1OrganisationMembersChanged // Event containing the contract specifics and raw log

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
func (it *OrganisationsV1OrganisationMembersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganisationsV1OrganisationMembersChanged)
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
		it.Event = new(OrganisationsV1OrganisationMembersChanged)
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
func (it *OrganisationsV1OrganisationMembersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganisationsV1OrganisationMembersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganisationsV1OrganisationMembersChanged represents a OrganisationMembersChanged event raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationMembersChanged struct {
	Hash       [32]byte
	OldMembers []common.Address
	NewMembers []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrganisationMembersChanged is a free log retrieval operation binding the contract event 0xd89a1da0f5f472e29dce6d5509d4b3c637e5f58daa235cf7a0b650e46a3051a8.
//
// Solidity: event organisationMembersChanged(bytes32 hash, address[] oldMembers, address[] newMembers)
func (_OrganisationsV1 *OrganisationsV1Filterer) FilterOrganisationMembersChanged(opts *bind.FilterOpts) (*OrganisationsV1OrganisationMembersChangedIterator, error) {

	logs, sub, err := _OrganisationsV1.contract.FilterLogs(opts, "organisationMembersChanged")
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1OrganisationMembersChangedIterator{contract: _OrganisationsV1.contract, event: "organisationMembersChanged", logs: logs, sub: sub}, nil
}

// WatchOrganisationMembersChanged is a free log subscription operation binding the contract event 0xd89a1da0f5f472e29dce6d5509d4b3c637e5f58daa235cf7a0b650e46a3051a8.
//
// Solidity: event organisationMembersChanged(bytes32 hash, address[] oldMembers, address[] newMembers)
func (_OrganisationsV1 *OrganisationsV1Filterer) WatchOrganisationMembersChanged(opts *bind.WatchOpts, sink chan<- *OrganisationsV1OrganisationMembersChanged) (event.Subscription, error) {

	logs, sub, err := _OrganisationsV1.contract.WatchLogs(opts, "organisationMembersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganisationsV1OrganisationMembersChanged)
				if err := _OrganisationsV1.contract.UnpackLog(event, "organisationMembersChanged", log); err != nil {
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

// ParseOrganisationMembersChanged is a log parse operation binding the contract event 0xd89a1da0f5f472e29dce6d5509d4b3c637e5f58daa235cf7a0b650e46a3051a8.
//
// Solidity: event organisationMembersChanged(bytes32 hash, address[] oldMembers, address[] newMembers)
func (_OrganisationsV1 *OrganisationsV1Filterer) ParseOrganisationMembersChanged(log types.Log) (*OrganisationsV1OrganisationMembersChanged, error) {
	event := new(OrganisationsV1OrganisationMembersChanged)
	if err := _OrganisationsV1.contract.UnpackLog(event, "organisationMembersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganisationsV1OrganisationOwnerChangedIterator is returned from FilterOrganisationOwnerChanged and is used to iterate over the raw logs and unpacked data for OrganisationOwnerChanged events raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationOwnerChangedIterator struct {
	Event *OrganisationsV1OrganisationOwnerChanged // Event containing the contract specifics and raw log

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
func (it *OrganisationsV1OrganisationOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganisationsV1OrganisationOwnerChanged)
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
		it.Event = new(OrganisationsV1OrganisationOwnerChanged)
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
func (it *OrganisationsV1OrganisationOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganisationsV1OrganisationOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganisationsV1OrganisationOwnerChanged represents a OrganisationOwnerChanged event raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationOwnerChanged struct {
	Hash     [32]byte
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrganisationOwnerChanged is a free log retrieval operation binding the contract event 0xe6ea8e5ace912ff2fa4c02f7bafe620691508973d49774daee6a08862f9614e8.
//
// Solidity: event organisationOwnerChanged(bytes32 hash, address oldOwner, address newOwner)
func (_OrganisationsV1 *OrganisationsV1Filterer) FilterOrganisationOwnerChanged(opts *bind.FilterOpts) (*OrganisationsV1OrganisationOwnerChangedIterator, error) {

	logs, sub, err := _OrganisationsV1.contract.FilterLogs(opts, "organisationOwnerChanged")
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1OrganisationOwnerChangedIterator{contract: _OrganisationsV1.contract, event: "organisationOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOrganisationOwnerChanged is a free log subscription operation binding the contract event 0xe6ea8e5ace912ff2fa4c02f7bafe620691508973d49774daee6a08862f9614e8.
//
// Solidity: event organisationOwnerChanged(bytes32 hash, address oldOwner, address newOwner)
func (_OrganisationsV1 *OrganisationsV1Filterer) WatchOrganisationOwnerChanged(opts *bind.WatchOpts, sink chan<- *OrganisationsV1OrganisationOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _OrganisationsV1.contract.WatchLogs(opts, "organisationOwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganisationsV1OrganisationOwnerChanged)
				if err := _OrganisationsV1.contract.UnpackLog(event, "organisationOwnerChanged", log); err != nil {
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

// ParseOrganisationOwnerChanged is a log parse operation binding the contract event 0xe6ea8e5ace912ff2fa4c02f7bafe620691508973d49774daee6a08862f9614e8.
//
// Solidity: event organisationOwnerChanged(bytes32 hash, address oldOwner, address newOwner)
func (_OrganisationsV1 *OrganisationsV1Filterer) ParseOrganisationOwnerChanged(log types.Log) (*OrganisationsV1OrganisationOwnerChanged, error) {
	event := new(OrganisationsV1OrganisationOwnerChanged)
	if err := _OrganisationsV1.contract.UnpackLog(event, "organisationOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganisationsV1OrganisationRemovedIterator is returned from FilterOrganisationRemoved and is used to iterate over the raw logs and unpacked data for OrganisationRemoved events raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationRemovedIterator struct {
	Event *OrganisationsV1OrganisationRemoved // Event containing the contract specifics and raw log

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
func (it *OrganisationsV1OrganisationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganisationsV1OrganisationRemoved)
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
		it.Event = new(OrganisationsV1OrganisationRemoved)
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
func (it *OrganisationsV1OrganisationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganisationsV1OrganisationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganisationsV1OrganisationRemoved represents a OrganisationRemoved event raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationRemoved struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrganisationRemoved is a free log retrieval operation binding the contract event 0x9f11f05fe2810e17e680075c9e99bca8f9acd92f043d2cad824ab94decc9e978.
//
// Solidity: event organisationRemoved(bytes32 hash)
func (_OrganisationsV1 *OrganisationsV1Filterer) FilterOrganisationRemoved(opts *bind.FilterOpts) (*OrganisationsV1OrganisationRemovedIterator, error) {

	logs, sub, err := _OrganisationsV1.contract.FilterLogs(opts, "organisationRemoved")
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1OrganisationRemovedIterator{contract: _OrganisationsV1.contract, event: "organisationRemoved", logs: logs, sub: sub}, nil
}

// WatchOrganisationRemoved is a free log subscription operation binding the contract event 0x9f11f05fe2810e17e680075c9e99bca8f9acd92f043d2cad824ab94decc9e978.
//
// Solidity: event organisationRemoved(bytes32 hash)
func (_OrganisationsV1 *OrganisationsV1Filterer) WatchOrganisationRemoved(opts *bind.WatchOpts, sink chan<- *OrganisationsV1OrganisationRemoved) (event.Subscription, error) {

	logs, sub, err := _OrganisationsV1.contract.WatchLogs(opts, "organisationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganisationsV1OrganisationRemoved)
				if err := _OrganisationsV1.contract.UnpackLog(event, "organisationRemoved", log); err != nil {
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

// ParseOrganisationRemoved is a log parse operation binding the contract event 0x9f11f05fe2810e17e680075c9e99bca8f9acd92f043d2cad824ab94decc9e978.
//
// Solidity: event organisationRemoved(bytes32 hash)
func (_OrganisationsV1 *OrganisationsV1Filterer) ParseOrganisationRemoved(log types.Log) (*OrganisationsV1OrganisationRemoved, error) {
	event := new(OrganisationsV1OrganisationRemoved)
	if err := _OrganisationsV1.contract.UnpackLog(event, "organisationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrganisationsV1OrganisationRenamedIterator is returned from FilterOrganisationRenamed and is used to iterate over the raw logs and unpacked data for OrganisationRenamed events raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationRenamedIterator struct {
	Event *OrganisationsV1OrganisationRenamed // Event containing the contract specifics and raw log

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
func (it *OrganisationsV1OrganisationRenamedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganisationsV1OrganisationRenamed)
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
		it.Event = new(OrganisationsV1OrganisationRenamed)
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
func (it *OrganisationsV1OrganisationRenamedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganisationsV1OrganisationRenamedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganisationsV1OrganisationRenamed represents a OrganisationRenamed event raised by the OrganisationsV1 contract.
type OrganisationsV1OrganisationRenamed struct {
	OldHash [32]byte
	NewHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrganisationRenamed is a free log retrieval operation binding the contract event 0x83edd531f7a6ab02a34fcaea2a9f5922b2eda44c477002ad4eba9945a50ef55e.
//
// Solidity: event organisationRenamed(bytes32 oldHash, bytes32 newHash)
func (_OrganisationsV1 *OrganisationsV1Filterer) FilterOrganisationRenamed(opts *bind.FilterOpts) (*OrganisationsV1OrganisationRenamedIterator, error) {

	logs, sub, err := _OrganisationsV1.contract.FilterLogs(opts, "organisationRenamed")
	if err != nil {
		return nil, err
	}
	return &OrganisationsV1OrganisationRenamedIterator{contract: _OrganisationsV1.contract, event: "organisationRenamed", logs: logs, sub: sub}, nil
}

// WatchOrganisationRenamed is a free log subscription operation binding the contract event 0x83edd531f7a6ab02a34fcaea2a9f5922b2eda44c477002ad4eba9945a50ef55e.
//
// Solidity: event organisationRenamed(bytes32 oldHash, bytes32 newHash)
func (_OrganisationsV1 *OrganisationsV1Filterer) WatchOrganisationRenamed(opts *bind.WatchOpts, sink chan<- *OrganisationsV1OrganisationRenamed) (event.Subscription, error) {

	logs, sub, err := _OrganisationsV1.contract.WatchLogs(opts, "organisationRenamed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganisationsV1OrganisationRenamed)
				if err := _OrganisationsV1.contract.UnpackLog(event, "organisationRenamed", log); err != nil {
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

// ParseOrganisationRenamed is a log parse operation binding the contract event 0x83edd531f7a6ab02a34fcaea2a9f5922b2eda44c477002ad4eba9945a50ef55e.
//
// Solidity: event organisationRenamed(bytes32 oldHash, bytes32 newHash)
func (_OrganisationsV1 *OrganisationsV1Filterer) ParseOrganisationRenamed(log types.Log) (*OrganisationsV1OrganisationRenamed, error) {
	event := new(OrganisationsV1OrganisationRenamed)
	if err := _OrganisationsV1.contract.UnpackLog(event, "organisationRenamed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
