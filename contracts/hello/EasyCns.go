// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hello

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EasyCnsABI is the input ABI used to generate the binding from.
const EasyCnsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// EasyCnsBin is the compiled bytecode used for deploying new contracts.
const EasyCnsBin = `0x6060604052341561000c57fe5b5b60008054600160a060020a03191633600160a060020a03161790555b5b610165806100396000396000f300606060405263ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663693ec85e811461003a575bfe5b341561004257fe5b610090600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437509496506100b995505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60006001826040518082805190602001908083835b602083106100ed5780518252601f1990920191602091820191016100ce565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205473ffffffffffffffffffffffffffffffffffffffff16925050505b9190505600a165627a7a72305820f395cb225ed00c70277bfd934e389d348f6548d52f8c498fc337805a686c7a560029`

// DeployEasyCns deploys a new Ethereum contract, binding an instance of EasyCns to it.
func DeployEasyCns(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EasyCns, error) {
	parsed, err := abi.JSON(strings.NewReader(EasyCnsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EasyCnsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EasyCns{EasyCnsCaller: EasyCnsCaller{contract: contract}, EasyCnsTransactor: EasyCnsTransactor{contract: contract}, EasyCnsFilterer: EasyCnsFilterer{contract: contract}}, nil
}

// EasyCns is an auto generated Go binding around an Ethereum contract.
type EasyCns struct {
	EasyCnsCaller     // Read-only binding to the contract
	EasyCnsTransactor // Write-only binding to the contract
	EasyCnsFilterer   // Log filterer for contract events
}

// EasyCnsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EasyCnsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EasyCnsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EasyCnsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EasyCnsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EasyCnsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EasyCnsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EasyCnsSession struct {
	Contract     *EasyCns          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EasyCnsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EasyCnsCallerSession struct {
	Contract *EasyCnsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EasyCnsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EasyCnsTransactorSession struct {
	Contract     *EasyCnsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EasyCnsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EasyCnsRaw struct {
	Contract *EasyCns // Generic contract binding to access the raw methods on
}

// EasyCnsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EasyCnsCallerRaw struct {
	Contract *EasyCnsCaller // Generic read-only contract binding to access the raw methods on
}

// EasyCnsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EasyCnsTransactorRaw struct {
	Contract *EasyCnsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEasyCns creates a new instance of EasyCns, bound to a specific deployed contract.
func NewEasyCns(address common.Address, backend bind.ContractBackend) (*EasyCns, error) {
	contract, err := bindEasyCns(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EasyCns{EasyCnsCaller: EasyCnsCaller{contract: contract}, EasyCnsTransactor: EasyCnsTransactor{contract: contract}, EasyCnsFilterer: EasyCnsFilterer{contract: contract}}, nil
}

// NewEasyCnsCaller creates a new read-only instance of EasyCns, bound to a specific deployed contract.
func NewEasyCnsCaller(address common.Address, caller bind.ContractCaller) (*EasyCnsCaller, error) {
	contract, err := bindEasyCns(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EasyCnsCaller{contract: contract}, nil
}

// NewEasyCnsTransactor creates a new write-only instance of EasyCns, bound to a specific deployed contract.
func NewEasyCnsTransactor(address common.Address, transactor bind.ContractTransactor) (*EasyCnsTransactor, error) {
	contract, err := bindEasyCns(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EasyCnsTransactor{contract: contract}, nil
}

// NewEasyCnsFilterer creates a new log filterer instance of EasyCns, bound to a specific deployed contract.
func NewEasyCnsFilterer(address common.Address, filterer bind.ContractFilterer) (*EasyCnsFilterer, error) {
	contract, err := bindEasyCns(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EasyCnsFilterer{contract: contract}, nil
}

// bindEasyCns binds a generic wrapper to an already deployed contract.
func bindEasyCns(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EasyCnsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EasyCns *EasyCnsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EasyCns.Contract.EasyCnsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EasyCns *EasyCnsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EasyCns.Contract.EasyCnsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EasyCns *EasyCnsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EasyCns.Contract.EasyCnsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EasyCns *EasyCnsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EasyCns.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EasyCns *EasyCnsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EasyCns.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EasyCns *EasyCnsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EasyCns.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(address)
func (_EasyCns *EasyCnsCaller) Get(opts *bind.CallOpts, name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EasyCns.contract.Call(opts, out, "get", name)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(address)
func (_EasyCns *EasyCnsSession) Get(name string) (common.Address, error) {
	return _EasyCns.Contract.Get(&_EasyCns.CallOpts, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(address)
func (_EasyCns *EasyCnsCallerSession) Get(name string) (common.Address, error) {
	return _EasyCns.Contract.Get(&_EasyCns.CallOpts, name)
}
