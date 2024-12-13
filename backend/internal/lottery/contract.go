// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lottery

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticketPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"WinnerSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipantsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTicketNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUnlockedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"restartEmpty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051610ca6380380610ca683398101604081905261002e916100f0565b60028311156100785760405162461bcd60e51b8152602060048201526011602482015270696c6c6567616c206f776e65722066656560781b60448201526064015b60405180910390fd5b620151808210156100ba5760405162461bcd60e51b815260206004820152600c60248201526b3130b210323ab930ba34b7b760a11b604482015260640161006f565b5f80546001600160a01b0319163317905560a0839052608081905260028290556100e4824261011b565b60015550610140915050565b5f5f5f60608486031215610102575f5ffd5b5050815160208301516040909301519094929350919050565b8082018082111561013a57634e487b7160e01b5f52601160045260245ffd5b92915050565b60805160a051610b226101845f395f81816102f0015261083501525f818161012c01528181610354015281816103d20152818161040901526104eb0152610b225ff3fe6080604052600436106100ef575f3560e01c806351cff8d911610087578063be9a655511610057578063be9a6555146102b7578063c7e284b8146102cb578063d5b2a01a146102df578063e9e8d83314610312575f5ffd5b806351cff8d9146102235780636bc1fb8a146102425780638404b377146102615780638a9e8671146102a2575f5ffd5b806327e235e3116100c257806327e235e3146101975780633197cbb6146101c257806335c1d349146101d7578063454a2ab31461020e575f5ffd5b80630fb5a6b4146100f35780631209b1f61461011b578063129de5bf1461014e578063146ca53114610182575b5f5ffd5b3480156100fe575f5ffd5b5061010860025481565b6040519081526020015b60405180910390f35b348015610126575f5ffd5b506101087f000000000000000000000000000000000000000000000000000000000000000081565b348015610159575f5ffd5b506101086101683660046109f6565b6001600160a01b03165f9081526006602052604090205490565b34801561018d575f5ffd5b5061010860035481565b3480156101a2575f5ffd5b506101086101b13660046109f6565b60066020525f908152604090205481565b3480156101cd575f5ffd5b5061010860015481565b3480156101e2575f5ffd5b506101f66101f1366004610a18565b610326565b6040516001600160a01b039091168152602001610112565b61022161021c366004610a18565b61034e565b005b34801561022e575f5ffd5b5061022161023d3660046109f6565b61056f565b34801561024d575f5ffd5b5061022161025c366004610a18565b61064e565b34801561026c575f5ffd5b5061010861027b3660046109f6565b6001600160a01b03165f908152600760209081526040808320600354845290915290205490565b3480156102ad575f5ffd5b5061010860045481565b3480156102c2575f5ffd5b506102216106fe565b3480156102d6575f5ffd5b5061010861094f565b3480156102ea575f5ffd5b506101087f000000000000000000000000000000000000000000000000000000000000000081565b34801561031d575f5ffd5b50600554610108565b60058181548110610335575f80fd5b5f918252602090912001546001600160a01b0316905081565b346103797f000000000000000000000000000000000000000000000000000000000000000083610a43565b11156103cc5760405162461bcd60e51b815260206004820152601f60248201527f6e6f7420656e6f75676820657468657220746f20627579207469636b6574730060448201526064015b60405180910390fd5b5f6103f77f000000000000000000000000000000000000000000000000000000000000000083610a43565b6104019034610a60565b90505f61042e7f000000000000000000000000000000000000000000000000000000000000000084610a43565b6104389034610a60565b111561046257335f908152600660205260408120805483929061045c908490610a73565b90915550505b5f5b828110156104b45760058054600181810183555f929092527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b0319163317905501610464565b50335f9081526007602090815260408083206003548452909152812080548492906104e0908490610a73565b9091555061051090507f000000000000000000000000000000000000000000000000000000000000000083610a43565b60045f8282546105209190610a73565b9091555050600354604080518481524260208201529081019190915233907f4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b3059060600160405180910390a25050565b335f908152600660205260409020546105c05760405162461bcd60e51b81526020600482015260136024820152726e6f7468696e6720746f20776974686472617760681b60448201526064016103c3565b335f90815260066020526040808220805490839055905190916001600160a01b0384169183156108fc0291849190818181858888f19350505050158015610609573d5f5f3e3d5ffd5b506040518181526001600160a01b0383169033907f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9060200160405180910390a35050565b5f546001600160a01b031633146106945760405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b60448201526064016103c3565b600554156106d85760405162461bcd60e51b815260206004820152601160248201527068617665207061727469636970616e747360781b60448201526064016103c3565b805f036106e7576002546106e9565b805b60028190556106f89042610a73565b60015550565b5f546001600160a01b031633146107445760405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b60448201526064016103c3565b6001544210156107965760405162461bcd60e51b815260206004820152601b60248201527f6c6f74746572792074696d656f7574206973206e6f74206f766572000000000060448201526064016103c3565b6005546107e55760405162461bcd60e51b815260206004820152601760248201527f6e6f7420656e6f756768207061727469636970616e747300000000000000000060448201526064016103c3565b6005545f906107f2610972565b6107fc9190610a9a565b90505f6005828154811061081257610812610aad565b5f9182526020822001546004546001600160a01b03909116925060649061085a907f000000000000000000000000000000000000000000000000000000000000000090610a43565b6108649190610ac1565b90505f816004546108759190610a60565b335f90815260066020526040812080549293508492909190610898908490610a73565b90915550506001600160a01b0383165f90815260066020526040812080548392906108c4908490610a73565b90915550505f6004556003546040516001600160a01b038516917f866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b939161091291858252602082015260400190565b60405180910390a260038054905f61092983610ad4565b909155505060025461093b9042610a73565b60015561094960055f6109ac565b50505050565b5f42600154101561095f57505f90565b4260015461096d9190610a60565b905090565b5f444260405160200161098f929190918252602082015260400190565b604051602081830303815290604052805190602001205f1c905090565b5080545f8255905f5260205f20908101906109c791906109ca565b50565b5b808211156109de575f81556001016109cb565b5090565b6001600160a01b03811681146109c7575f5ffd5b5f60208284031215610a06575f5ffd5b8135610a11816109e2565b9392505050565b5f60208284031215610a28575f5ffd5b5035919050565b634e487b7160e01b5f52601160045260245ffd5b8082028115828204841417610a5a57610a5a610a2f565b92915050565b81810381811115610a5a57610a5a610a2f565b80820180821115610a5a57610a5a610a2f565b634e487b7160e01b5f52601260045260245ffd5b5f82610aa857610aa8610a86565b500690565b634e487b7160e01b5f52603260045260245ffd5b5f82610acf57610acf610a86565b500490565b5f60018201610ae557610ae5610a2f565b506001019056fea2646970667358221220eab7f90637a5ddba06446b298c695fd06030e50b90a2067cc2a4c50513a930da64736f6c634300081c0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _ownerFee *big.Int, _duration *big.Int, _ticketPrice *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _ownerFee, _duration, _ticketPrice)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Contract *ContractCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Contract *ContractSession) Duration() (*big.Int, error) {
	return _Contract.Contract.Duration(&_Contract.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Contract *ContractCallerSession) Duration() (*big.Int, error) {
	return _Contract.Contract.Duration(&_Contract.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Contract *ContractCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Contract *ContractSession) EndTime() (*big.Int, error) {
	return _Contract.Contract.EndTime(&_Contract.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Contract *ContractCallerSession) EndTime() (*big.Int, error) {
	return _Contract.Contract.EndTime(&_Contract.CallOpts)
}

// GetParticipantsNum is a free data retrieval call binding the contract method 0xe9e8d833.
//
// Solidity: function getParticipantsNum() view returns(uint256)
func (_Contract *ContractCaller) GetParticipantsNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getParticipantsNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParticipantsNum is a free data retrieval call binding the contract method 0xe9e8d833.
//
// Solidity: function getParticipantsNum() view returns(uint256)
func (_Contract *ContractSession) GetParticipantsNum() (*big.Int, error) {
	return _Contract.Contract.GetParticipantsNum(&_Contract.CallOpts)
}

// GetParticipantsNum is a free data retrieval call binding the contract method 0xe9e8d833.
//
// Solidity: function getParticipantsNum() view returns(uint256)
func (_Contract *ContractCallerSession) GetParticipantsNum() (*big.Int, error) {
	return _Contract.Contract.GetParticipantsNum(&_Contract.CallOpts)
}

// GetTicketNum is a free data retrieval call binding the contract method 0x8404b377.
//
// Solidity: function getTicketNum(address account) view returns(uint256)
func (_Contract *ContractCaller) GetTicketNum(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTicketNum", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTicketNum is a free data retrieval call binding the contract method 0x8404b377.
//
// Solidity: function getTicketNum(address account) view returns(uint256)
func (_Contract *ContractSession) GetTicketNum(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTicketNum(&_Contract.CallOpts, account)
}

// GetTicketNum is a free data retrieval call binding the contract method 0x8404b377.
//
// Solidity: function getTicketNum(address account) view returns(uint256)
func (_Contract *ContractCallerSession) GetTicketNum(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTicketNum(&_Contract.CallOpts, account)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Contract *ContractCaller) GetTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Contract *ContractSession) GetTimeLeft() (*big.Int, error) {
	return _Contract.Contract.GetTimeLeft(&_Contract.CallOpts)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Contract *ContractCallerSession) GetTimeLeft() (*big.Int, error) {
	return _Contract.Contract.GetTimeLeft(&_Contract.CallOpts)
}

// GetUnlockedBalance is a free data retrieval call binding the contract method 0x129de5bf.
//
// Solidity: function getUnlockedBalance(address account) view returns(uint256)
func (_Contract *ContractCaller) GetUnlockedBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUnlockedBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnlockedBalance is a free data retrieval call binding the contract method 0x129de5bf.
//
// Solidity: function getUnlockedBalance(address account) view returns(uint256)
func (_Contract *ContractSession) GetUnlockedBalance(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetUnlockedBalance(&_Contract.CallOpts, account)
}

// GetUnlockedBalance is a free data retrieval call binding the contract method 0x129de5bf.
//
// Solidity: function getUnlockedBalance(address account) view returns(uint256)
func (_Contract *ContractCallerSession) GetUnlockedBalance(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetUnlockedBalance(&_Contract.CallOpts, account)
}

// OwnerFee is a free data retrieval call binding the contract method 0xd5b2a01a.
//
// Solidity: function ownerFee() view returns(uint256)
func (_Contract *ContractCaller) OwnerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerFee is a free data retrieval call binding the contract method 0xd5b2a01a.
//
// Solidity: function ownerFee() view returns(uint256)
func (_Contract *ContractSession) OwnerFee() (*big.Int, error) {
	return _Contract.Contract.OwnerFee(&_Contract.CallOpts)
}

// OwnerFee is a free data retrieval call binding the contract method 0xd5b2a01a.
//
// Solidity: function ownerFee() view returns(uint256)
func (_Contract *ContractCallerSession) OwnerFee() (*big.Int, error) {
	return _Contract.Contract.OwnerFee(&_Contract.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_Contract *ContractCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_Contract *ContractSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Participants(&_Contract.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Participants(&_Contract.CallOpts, arg0)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Contract *ContractCaller) Round(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "round")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Contract *ContractSession) Round() (*big.Int, error) {
	return _Contract.Contract.Round(&_Contract.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Contract *ContractCallerSession) Round() (*big.Int, error) {
	return _Contract.Contract.Round(&_Contract.CallOpts)
}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_Contract *ContractCaller) TicketPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ticketPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_Contract *ContractSession) TicketPrice() (*big.Int, error) {
	return _Contract.Contract.TicketPrice(&_Contract.CallOpts)
}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_Contract *ContractCallerSession) TicketPrice() (*big.Int, error) {
	return _Contract.Contract.TicketPrice(&_Contract.CallOpts)
}

// TotalBid is a free data retrieval call binding the contract method 0x8a9e8671.
//
// Solidity: function totalBid() view returns(uint256)
func (_Contract *ContractCaller) TotalBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBid is a free data retrieval call binding the contract method 0x8a9e8671.
//
// Solidity: function totalBid() view returns(uint256)
func (_Contract *ContractSession) TotalBid() (*big.Int, error) {
	return _Contract.Contract.TotalBid(&_Contract.CallOpts)
}

// TotalBid is a free data retrieval call binding the contract method 0x8a9e8671.
//
// Solidity: function totalBid() view returns(uint256)
func (_Contract *ContractCallerSession) TotalBid() (*big.Int, error) {
	return _Contract.Contract.TotalBid(&_Contract.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) payable returns()
func (_Contract *ContractTransactor) Bid(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "bid", amount)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) payable returns()
func (_Contract *ContractSession) Bid(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Bid(&_Contract.TransactOpts, amount)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) payable returns()
func (_Contract *ContractTransactorSession) Bid(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Bid(&_Contract.TransactOpts, amount)
}

// RestartEmpty is a paid mutator transaction binding the contract method 0x6bc1fb8a.
//
// Solidity: function restartEmpty(uint256 newDuration) returns()
func (_Contract *ContractTransactor) RestartEmpty(opts *bind.TransactOpts, newDuration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "restartEmpty", newDuration)
}

// RestartEmpty is a paid mutator transaction binding the contract method 0x6bc1fb8a.
//
// Solidity: function restartEmpty(uint256 newDuration) returns()
func (_Contract *ContractSession) RestartEmpty(newDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RestartEmpty(&_Contract.TransactOpts, newDuration)
}

// RestartEmpty is a paid mutator transaction binding the contract method 0x6bc1fb8a.
//
// Solidity: function restartEmpty(uint256 newDuration) returns()
func (_Contract *ContractTransactorSession) RestartEmpty(newDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RestartEmpty(&_Contract.TransactOpts, newDuration)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Contract *ContractTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Contract *ContractSession) Start() (*types.Transaction, error) {
	return _Contract.Contract.Start(&_Contract.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Contract *ContractTransactorSession) Start() (*types.Transaction, error) {
	return _Contract.Contract.Start(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Contract *ContractSession) Withdraw(_to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Contract *ContractTransactorSession) Withdraw(_to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, _to)
}

// ContractBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Contract contract.
type ContractBidIterator struct {
	Event *ContractBid // Event containing the contract specifics and raw log

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
func (it *ContractBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBid)
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
		it.Event = new(ContractBid)
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
func (it *ContractBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBid represents a Bid event raised by the Contract contract.
type ContractBid struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Round     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0x4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b305.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 timestamp, uint256 round)
func (_Contract *ContractFilterer) FilterBid(opts *bind.FilterOpts, account []common.Address) (*ContractBidIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Bid", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractBidIterator{contract: _Contract.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0x4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b305.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 timestamp, uint256 round)
func (_Contract *ContractFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *ContractBid, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Bid", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBid)
				if err := _Contract.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0x4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b305.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 timestamp, uint256 round)
func (_Contract *ContractFilterer) ParseBid(log types.Log) (*ContractBid, error) {
	event := new(ContractBid)
	if err := _Contract.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWinnerSelectedIterator is returned from FilterWinnerSelected and is used to iterate over the raw logs and unpacked data for WinnerSelected events raised by the Contract contract.
type ContractWinnerSelectedIterator struct {
	Event *ContractWinnerSelected // Event containing the contract specifics and raw log

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
func (it *ContractWinnerSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWinnerSelected)
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
		it.Event = new(ContractWinnerSelected)
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
func (it *ContractWinnerSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWinnerSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWinnerSelected represents a WinnerSelected event raised by the Contract contract.
type ContractWinnerSelected struct {
	Account common.Address
	Amount  *big.Int
	Round   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWinnerSelected is a free log retrieval operation binding the contract event 0x866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b93.
//
// Solidity: event WinnerSelected(address indexed account, uint256 amount, uint256 round)
func (_Contract *ContractFilterer) FilterWinnerSelected(opts *bind.FilterOpts, account []common.Address) (*ContractWinnerSelectedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "WinnerSelected", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractWinnerSelectedIterator{contract: _Contract.contract, event: "WinnerSelected", logs: logs, sub: sub}, nil
}

// WatchWinnerSelected is a free log subscription operation binding the contract event 0x866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b93.
//
// Solidity: event WinnerSelected(address indexed account, uint256 amount, uint256 round)
func (_Contract *ContractFilterer) WatchWinnerSelected(opts *bind.WatchOpts, sink chan<- *ContractWinnerSelected, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "WinnerSelected", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWinnerSelected)
				if err := _Contract.contract.UnpackLog(event, "WinnerSelected", log); err != nil {
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

// ParseWinnerSelected is a log parse operation binding the contract event 0x866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b93.
//
// Solidity: event WinnerSelected(address indexed account, uint256 amount, uint256 round)
func (_Contract *ContractFilterer) ParseWinnerSelected(log types.Log) (*ContractWinnerSelected, error) {
	event := new(ContractWinnerSelected)
	if err := _Contract.contract.UnpackLog(event, "WinnerSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Contract contract.
type ContractWithdrawIterator struct {
	Event *ContractWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdraw)
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
		it.Event = new(ContractWithdraw)
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
func (it *ContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdraw represents a Withdraw event raised by the Contract contract.
type ContractWithdraw struct {
	Account common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed account, address indexed to, uint256 amount)
func (_Contract *ContractFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address, to []common.Address) (*ContractWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdraw", accountRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawIterator{contract: _Contract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed account, address indexed to, uint256 amount)
func (_Contract *ContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractWithdraw, account []common.Address, to []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdraw", accountRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdraw)
				if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed account, address indexed to, uint256 amount)
func (_Contract *ContractFilterer) ParseWithdraw(log types.Log) (*ContractWithdraw, error) {
	event := new(ContractWithdraw)
	if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

