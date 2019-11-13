// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hello

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// HelloABI is the input ABI used to generate the binding from.
const HelloABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"speak\",\"outputs\":[{\"name\":\"itSays\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newSaying\",\"type\":\"string\"}],\"name\":\"saySomethingElse\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newSaying\",\"type\":\"string\"}],\"name\":\"onSaySomethingElse\",\"type\":\"event\"}]"

// HelloBin is the compiled bytecode used for deploying new contracts.
const HelloBin = `6060604052341561000c57fe5b5b604060405190810160405280600c81526020017f48656c6c6f20576f726c6421000000000000000000000000000000000000000081525060009080519060200190610059929190610060565b505b610105565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a157805160ff19168380011785556100cf565b828001600101855582156100cf579182015b828111156100ce5782518255916020019190600101906100b3565b5b5090506100dc91906100e0565b5090565b61010291905b808211156100fe5760008160009055506001016100e6565b5090565b90565b6103ac806101146000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806350d8531514610046578063b07d7e39146100df575bfe5b341561004e57fe5b610056610151565b60405180806020018281038252838181518152602001915080519060200190808383600083146100a5575b8051825260208311156100a557602082019150602081019050602083039250610081565b505050905090810190601f1680156100d15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156100e757fe5b610137600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506101fa565b604051808215151515815260200191505060405180910390f35b6101596102c7565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101ef5780601f106101c4576101008083540402835291602001916101ef565b820191906000526020600020905b8154815290600101906020018083116101d257829003601f168201915b505050505090505b90565b600081600090805190602001906102129291906102db565b507f87a4687b43c451ab438a8f4ac187d795809dd4e395a6a9ff46b9120ad40848cf826040518080602001828103825283818151815260200191508051906020019080838360008314610284575b80518252602083111561028457602082019150602081019050602083039250610260565b505050905090810190601f1680156102b05780820380516001836020036101000a031916815260200191505b509250505060405180910390a1600190505b919050565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061031c57805160ff191683800117855561034a565b8280016001018555821561034a579182015b8281111561034957825182559160200191906001019061032e565b5b509050610357919061035b565b5090565b61037d91905b80821115610379576000816000905550600101610361565b5090565b905600a165627a7a72305820df813e659d2476104cd89ce84cfadae31cd8cea723b0917533220d1a18eec44f0029`

// DeployHello deploys a new Ethereum contract, binding an instance of Hello to it.
func DeployHello(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hello, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelloBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}}, nil
}

// Hello is an auto generated Go binding around an Ethereum contract.
type Hello struct {
	HelloCaller     // Read-only binding to the contract
	HelloTransactor // Write-only binding to the contract
	HelloFilterer   // Log filterer for contract events
}

// HelloCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloSession struct {
	Contract     *Hello            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloCallerSession struct {
	Contract *HelloCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloTransactorSession struct {
	Contract     *HelloTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloRaw struct {
	Contract *Hello // Generic contract binding to access the raw methods on
}

// HelloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloCallerRaw struct {
	Contract *HelloCaller // Generic read-only contract binding to access the raw methods on
}

// HelloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloTransactorRaw struct {
	Contract *HelloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHello creates a new instance of Hello, bound to a specific deployed contract.
func NewHello(address common.Address, backend bind.ContractBackend) (*Hello, error) {
	contract, err := bindHello(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}}, nil
}

// NewHelloCaller creates a new read-only instance of Hello, bound to a specific deployed contract.
func NewHelloCaller(address common.Address, caller bind.ContractCaller) (*HelloCaller, error) {
	contract, err := bindHello(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloCaller{contract: contract}, nil
}

// NewHelloTransactor creates a new write-only instance of Hello, bound to a specific deployed contract.
func NewHelloTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloTransactor, error) {
	contract, err := bindHello(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloTransactor{contract: contract}, nil
}

// NewHelloFilterer creates a new log filterer instance of Hello, bound to a specific deployed contract.
func NewHelloFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloFilterer, error) {
	contract, err := bindHello(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloFilterer{contract: contract}, nil
}

// bindHello binds a generic wrapper to an already deployed contract.
func bindHello(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.HelloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transact(opts, method, params...)
}

// Speak is a free data retrieval call binding the contract method 0x50d85315.
//
// Solidity: function speak() constant returns(itSays string)
func (_Hello *HelloCaller) Speak(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hello.contract.Call(opts, out, "speak")
	return *ret0, err
}

// Speak is a free data retrieval call binding the contract method 0x50d85315.
//
// Solidity: function speak() constant returns(itSays string)
func (_Hello *HelloSession) Speak() (string, error) {
	return _Hello.Contract.Speak(&_Hello.CallOpts)
}

// Speak is a free data retrieval call binding the contract method 0x50d85315.
//
// Solidity: function speak() constant returns(itSays string)
func (_Hello *HelloCallerSession) Speak() (string, error) {
	return _Hello.Contract.Speak(&_Hello.CallOpts)
}

// SaySomethingElse is a paid mutator transaction binding the contract method 0xb07d7e39.
//
// Solidity: function saySomethingElse(newSaying string) returns(success bool)
func (_Hello *HelloTransactor) SaySomethingElse(opts *bind.TransactOpts, newSaying string) (*types.Transaction, error) {
	return _Hello.contract.Transact(opts, "saySomethingElse", newSaying)
}

// SaySomethingElse is a paid mutator transaction binding the contract method 0xb07d7e39.
//
// Solidity: function saySomethingElse(newSaying string) returns(success bool)
func (_Hello *HelloSession) SaySomethingElse(newSaying string) (*types.Transaction, error) {
	return _Hello.Contract.SaySomethingElse(&_Hello.TransactOpts, newSaying)
}

// SaySomethingElse is a paid mutator transaction binding the contract method 0xb07d7e39.
//
// Solidity: function saySomethingElse(newSaying string) returns(success bool)
func (_Hello *HelloTransactorSession) SaySomethingElse(newSaying string) (*types.Transaction, error) {
	return _Hello.Contract.SaySomethingElse(&_Hello.TransactOpts, newSaying)
}

// HelloOnSaySomethingElseIterator is returned from FilterOnSaySomethingElse and is used to iterate over the raw logs and unpacked data for OnSaySomethingElse events raised by the Hello contract.
type HelloOnSaySomethingElseIterator struct {
	Event *HelloOnSaySomethingElse // Event containing the contract specifics and raw log

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
func (it *HelloOnSaySomethingElseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloOnSaySomethingElse)
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
		it.Event = new(HelloOnSaySomethingElse)
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
func (it *HelloOnSaySomethingElseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloOnSaySomethingElseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloOnSaySomethingElse represents a OnSaySomethingElse event raised by the Hello contract.
type HelloOnSaySomethingElse struct {
	NewSaying string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnSaySomethingElse is a free log retrieval operation binding the contract event 0x87a4687b43c451ab438a8f4ac187d795809dd4e395a6a9ff46b9120ad40848cf.
//
// Solidity: e onSaySomethingElse(newSaying string)
func (_Hello *HelloFilterer) FilterOnSaySomethingElse(opts *bind.FilterOpts) (*HelloOnSaySomethingElseIterator, error) {

	logs, sub, err := _Hello.contract.FilterLogs(opts, "onSaySomethingElse")
	if err != nil {
		return nil, err
	}
	return &HelloOnSaySomethingElseIterator{contract: _Hello.contract, event: "onSaySomethingElse", logs: logs, sub: sub}, nil
}

// WatchOnSaySomethingElse is a free log subscription operation binding the contract event 0x87a4687b43c451ab438a8f4ac187d795809dd4e395a6a9ff46b9120ad40848cf.
//
// Solidity: e onSaySomethingElse(newSaying string)
func (_Hello *HelloFilterer) WatchOnSaySomethingElse(opts *bind.WatchOpts, sink chan<- *HelloOnSaySomethingElse) (event.Subscription, error) {

	logs, sub, err := _Hello.contract.WatchLogs(opts, "onSaySomethingElse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloOnSaySomethingElse)
				if err := _Hello.contract.UnpackLog(event, "onSaySomethingElse", log); err != nil {
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
