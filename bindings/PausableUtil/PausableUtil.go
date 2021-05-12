// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PausableUtil

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

// PausableUtilABI is the input ABI used to generate the binding from.
const PausableUtilABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PausableUtil is an auto generated Go binding around an Ethereum contract.
type PausableUtil struct {
	PausableUtilCaller     // Read-only binding to the contract
	PausableUtilTransactor // Write-only binding to the contract
	PausableUtilFilterer   // Log filterer for contract events
}

// PausableUtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableUtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableUtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableUtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableUtilSession struct {
	Contract     *PausableUtil     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableUtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableUtilCallerSession struct {
	Contract *PausableUtilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PausableUtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableUtilTransactorSession struct {
	Contract     *PausableUtilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PausableUtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableUtilRaw struct {
	Contract *PausableUtil // Generic contract binding to access the raw methods on
}

// PausableUtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableUtilCallerRaw struct {
	Contract *PausableUtilCaller // Generic read-only contract binding to access the raw methods on
}

// PausableUtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableUtilTransactorRaw struct {
	Contract *PausableUtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableUtil creates a new instance of PausableUtil, bound to a specific deployed contract.
func NewPausableUtil(address common.Address, backend bind.ContractBackend) (*PausableUtil, error) {
	contract, err := bindPausableUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableUtil{PausableUtilCaller: PausableUtilCaller{contract: contract}, PausableUtilTransactor: PausableUtilTransactor{contract: contract}, PausableUtilFilterer: PausableUtilFilterer{contract: contract}}, nil
}

// NewPausableUtilCaller creates a new read-only instance of PausableUtil, bound to a specific deployed contract.
func NewPausableUtilCaller(address common.Address, caller bind.ContractCaller) (*PausableUtilCaller, error) {
	contract, err := bindPausableUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUtilCaller{contract: contract}, nil
}

// NewPausableUtilTransactor creates a new write-only instance of PausableUtil, bound to a specific deployed contract.
func NewPausableUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableUtilTransactor, error) {
	contract, err := bindPausableUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUtilTransactor{contract: contract}, nil
}

// NewPausableUtilFilterer creates a new log filterer instance of PausableUtil, bound to a specific deployed contract.
func NewPausableUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableUtilFilterer, error) {
	contract, err := bindPausableUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableUtilFilterer{contract: contract}, nil
}

// bindPausableUtil binds a generic wrapper to an already deployed contract.
func bindPausableUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableUtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUtil *PausableUtilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUtil.Contract.PausableUtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUtil *PausableUtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUtil.Contract.PausableUtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUtil *PausableUtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUtil.Contract.PausableUtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUtil *PausableUtilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUtil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUtil *PausableUtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUtil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUtil *PausableUtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUtil.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUtil *PausableUtilCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PausableUtil.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUtil *PausableUtilSession) Paused() (bool, error) {
	return _PausableUtil.Contract.Paused(&_PausableUtil.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUtil *PausableUtilCallerSession) Paused() (bool, error) {
	return _PausableUtil.Contract.Paused(&_PausableUtil.CallOpts)
}

// PausableUtilPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PausableUtil contract.
type PausableUtilPausedIterator struct {
	Event *PausableUtilPaused // Event containing the contract specifics and raw log

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
func (it *PausableUtilPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUtilPaused)
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
		it.Event = new(PausableUtilPaused)
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
func (it *PausableUtilPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUtilPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUtilPaused represents a Paused event raised by the PausableUtil contract.
type PausableUtilPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUtil *PausableUtilFilterer) FilterPaused(opts *bind.FilterOpts) (*PausableUtilPausedIterator, error) {

	logs, sub, err := _PausableUtil.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausableUtilPausedIterator{contract: _PausableUtil.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUtil *PausableUtilFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausableUtilPaused) (event.Subscription, error) {

	logs, sub, err := _PausableUtil.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUtilPaused)
				if err := _PausableUtil.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUtil *PausableUtilFilterer) ParsePaused(log types.Log) (*PausableUtilPaused, error) {
	event := new(PausableUtilPaused)
	if err := _PausableUtil.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUtilUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PausableUtil contract.
type PausableUtilUnpausedIterator struct {
	Event *PausableUtilUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUtilUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUtilUnpaused)
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
		it.Event = new(PausableUtilUnpaused)
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
func (it *PausableUtilUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUtilUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUtilUnpaused represents a Unpaused event raised by the PausableUtil contract.
type PausableUtilUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUtil *PausableUtilFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUtilUnpausedIterator, error) {

	logs, sub, err := _PausableUtil.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUtilUnpausedIterator{contract: _PausableUtil.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUtil *PausableUtilFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUtilUnpaused) (event.Subscription, error) {

	logs, sub, err := _PausableUtil.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUtilUnpaused)
				if err := _PausableUtil.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUtil *PausableUtilFilterer) ParseUnpaused(log types.Log) (*PausableUtilUnpaused, error) {
	event := new(PausableUtilUnpaused)
	if err := _PausableUtil.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
