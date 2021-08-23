// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PausableUpgradeSafe

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
)

// PausableUpgradeSafeMetaData contains all meta data concerning the PausableUpgradeSafe contract.
var PausableUpgradeSafeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50609a8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80635c975abb14602d575b600080fd5b6033604d565b604051808215151515815260200191505060405180910390f35b6000606560009054906101000a900460ff1690509056fea26469706673582212207ac17eb3e39a1780071cfc794380fbbf4c5f9520b90e0248c965a6381216934164736f6c63430006040033",
}

// PausableUpgradeSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableUpgradeSafeMetaData.ABI instead.
var PausableUpgradeSafeABI = PausableUpgradeSafeMetaData.ABI

// PausableUpgradeSafeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PausableUpgradeSafeMetaData.Bin instead.
var PausableUpgradeSafeBin = PausableUpgradeSafeMetaData.Bin

// DeployPausableUpgradeSafe deploys a new Ethereum contract, binding an instance of PausableUpgradeSafe to it.
func DeployPausableUpgradeSafe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PausableUpgradeSafe, error) {
	parsed, err := PausableUpgradeSafeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PausableUpgradeSafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PausableUpgradeSafe{PausableUpgradeSafeCaller: PausableUpgradeSafeCaller{contract: contract}, PausableUpgradeSafeTransactor: PausableUpgradeSafeTransactor{contract: contract}, PausableUpgradeSafeFilterer: PausableUpgradeSafeFilterer{contract: contract}}, nil
}

// PausableUpgradeSafe is an auto generated Go binding around an Ethereum contract.
type PausableUpgradeSafe struct {
	PausableUpgradeSafeCaller     // Read-only binding to the contract
	PausableUpgradeSafeTransactor // Write-only binding to the contract
	PausableUpgradeSafeFilterer   // Log filterer for contract events
}

// PausableUpgradeSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableUpgradeSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableUpgradeSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableUpgradeSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableUpgradeSafeSession struct {
	Contract     *PausableUpgradeSafe // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PausableUpgradeSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableUpgradeSafeCallerSession struct {
	Contract *PausableUpgradeSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PausableUpgradeSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableUpgradeSafeTransactorSession struct {
	Contract     *PausableUpgradeSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PausableUpgradeSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableUpgradeSafeRaw struct {
	Contract *PausableUpgradeSafe // Generic contract binding to access the raw methods on
}

// PausableUpgradeSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableUpgradeSafeCallerRaw struct {
	Contract *PausableUpgradeSafeCaller // Generic read-only contract binding to access the raw methods on
}

// PausableUpgradeSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableUpgradeSafeTransactorRaw struct {
	Contract *PausableUpgradeSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableUpgradeSafe creates a new instance of PausableUpgradeSafe, bound to a specific deployed contract.
func NewPausableUpgradeSafe(address common.Address, backend bind.ContractBackend) (*PausableUpgradeSafe, error) {
	contract, err := bindPausableUpgradeSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeSafe{PausableUpgradeSafeCaller: PausableUpgradeSafeCaller{contract: contract}, PausableUpgradeSafeTransactor: PausableUpgradeSafeTransactor{contract: contract}, PausableUpgradeSafeFilterer: PausableUpgradeSafeFilterer{contract: contract}}, nil
}

// NewPausableUpgradeSafeCaller creates a new read-only instance of PausableUpgradeSafe, bound to a specific deployed contract.
func NewPausableUpgradeSafeCaller(address common.Address, caller bind.ContractCaller) (*PausableUpgradeSafeCaller, error) {
	contract, err := bindPausableUpgradeSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeSafeCaller{contract: contract}, nil
}

// NewPausableUpgradeSafeTransactor creates a new write-only instance of PausableUpgradeSafe, bound to a specific deployed contract.
func NewPausableUpgradeSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableUpgradeSafeTransactor, error) {
	contract, err := bindPausableUpgradeSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeSafeTransactor{contract: contract}, nil
}

// NewPausableUpgradeSafeFilterer creates a new log filterer instance of PausableUpgradeSafe, bound to a specific deployed contract.
func NewPausableUpgradeSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableUpgradeSafeFilterer, error) {
	contract, err := bindPausableUpgradeSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeSafeFilterer{contract: contract}, nil
}

// bindPausableUpgradeSafe binds a generic wrapper to an already deployed contract.
func bindPausableUpgradeSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableUpgradeSafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeSafe *PausableUpgradeSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeSafe.Contract.PausableUpgradeSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeSafe *PausableUpgradeSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeSafe.Contract.PausableUpgradeSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeSafe *PausableUpgradeSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeSafe.Contract.PausableUpgradeSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeSafe *PausableUpgradeSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeSafe *PausableUpgradeSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeSafe *PausableUpgradeSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeSafe.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeSafe *PausableUpgradeSafeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PausableUpgradeSafe.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeSafe *PausableUpgradeSafeSession) Paused() (bool, error) {
	return _PausableUpgradeSafe.Contract.Paused(&_PausableUpgradeSafe.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeSafe *PausableUpgradeSafeCallerSession) Paused() (bool, error) {
	return _PausableUpgradeSafe.Contract.Paused(&_PausableUpgradeSafe.CallOpts)
}

// PausableUpgradeSafePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PausableUpgradeSafe contract.
type PausableUpgradeSafePausedIterator struct {
	Event *PausableUpgradeSafePaused // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeSafePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeSafePaused)
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
		it.Event = new(PausableUpgradeSafePaused)
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
func (it *PausableUpgradeSafePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeSafePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeSafePaused represents a Paused event raised by the PausableUpgradeSafe contract.
type PausableUpgradeSafePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeSafe *PausableUpgradeSafeFilterer) FilterPaused(opts *bind.FilterOpts) (*PausableUpgradeSafePausedIterator, error) {

	logs, sub, err := _PausableUpgradeSafe.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeSafePausedIterator{contract: _PausableUpgradeSafe.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeSafe *PausableUpgradeSafeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeSafePaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeSafe.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeSafePaused)
				if err := _PausableUpgradeSafe.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PausableUpgradeSafe *PausableUpgradeSafeFilterer) ParsePaused(log types.Log) (*PausableUpgradeSafePaused, error) {
	event := new(PausableUpgradeSafePaused)
	if err := _PausableUpgradeSafe.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUpgradeSafeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PausableUpgradeSafe contract.
type PausableUpgradeSafeUnpausedIterator struct {
	Event *PausableUpgradeSafeUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeSafeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeSafeUnpaused)
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
		it.Event = new(PausableUpgradeSafeUnpaused)
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
func (it *PausableUpgradeSafeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeSafeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeSafeUnpaused represents a Unpaused event raised by the PausableUpgradeSafe contract.
type PausableUpgradeSafeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeSafe *PausableUpgradeSafeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUpgradeSafeUnpausedIterator, error) {

	logs, sub, err := _PausableUpgradeSafe.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeSafeUnpausedIterator{contract: _PausableUpgradeSafe.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeSafe *PausableUpgradeSafeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeSafeUnpaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeSafe.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeSafeUnpaused)
				if err := _PausableUpgradeSafe.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PausableUpgradeSafe *PausableUpgradeSafeFilterer) ParseUnpaused(log types.Log) (*PausableUpgradeSafeUnpaused, error) {
	event := new(PausableUpgradeSafeUnpaused)
	if err := _PausableUpgradeSafe.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
