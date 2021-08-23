// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AccessControlUpgradeSafe

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

// AccessControlUpgradeSafeMetaData contains all meta data concerning the AccessControlUpgradeSafe contract.
var AccessControlUpgradeSafeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AccessControlUpgradeSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlUpgradeSafeMetaData.ABI instead.
var AccessControlUpgradeSafeABI = AccessControlUpgradeSafeMetaData.ABI

// AccessControlUpgradeSafe is an auto generated Go binding around an Ethereum contract.
type AccessControlUpgradeSafe struct {
	AccessControlUpgradeSafeCaller     // Read-only binding to the contract
	AccessControlUpgradeSafeTransactor // Write-only binding to the contract
	AccessControlUpgradeSafeFilterer   // Log filterer for contract events
}

// AccessControlUpgradeSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlUpgradeSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlUpgradeSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlUpgradeSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlUpgradeSafeSession struct {
	Contract     *AccessControlUpgradeSafe // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccessControlUpgradeSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlUpgradeSafeCallerSession struct {
	Contract *AccessControlUpgradeSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AccessControlUpgradeSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlUpgradeSafeTransactorSession struct {
	Contract     *AccessControlUpgradeSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AccessControlUpgradeSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlUpgradeSafeRaw struct {
	Contract *AccessControlUpgradeSafe // Generic contract binding to access the raw methods on
}

// AccessControlUpgradeSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlUpgradeSafeCallerRaw struct {
	Contract *AccessControlUpgradeSafeCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlUpgradeSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlUpgradeSafeTransactorRaw struct {
	Contract *AccessControlUpgradeSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlUpgradeSafe creates a new instance of AccessControlUpgradeSafe, bound to a specific deployed contract.
func NewAccessControlUpgradeSafe(address common.Address, backend bind.ContractBackend) (*AccessControlUpgradeSafe, error) {
	contract, err := bindAccessControlUpgradeSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeSafe{AccessControlUpgradeSafeCaller: AccessControlUpgradeSafeCaller{contract: contract}, AccessControlUpgradeSafeTransactor: AccessControlUpgradeSafeTransactor{contract: contract}, AccessControlUpgradeSafeFilterer: AccessControlUpgradeSafeFilterer{contract: contract}}, nil
}

// NewAccessControlUpgradeSafeCaller creates a new read-only instance of AccessControlUpgradeSafe, bound to a specific deployed contract.
func NewAccessControlUpgradeSafeCaller(address common.Address, caller bind.ContractCaller) (*AccessControlUpgradeSafeCaller, error) {
	contract, err := bindAccessControlUpgradeSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeSafeCaller{contract: contract}, nil
}

// NewAccessControlUpgradeSafeTransactor creates a new write-only instance of AccessControlUpgradeSafe, bound to a specific deployed contract.
func NewAccessControlUpgradeSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlUpgradeSafeTransactor, error) {
	contract, err := bindAccessControlUpgradeSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeSafeTransactor{contract: contract}, nil
}

// NewAccessControlUpgradeSafeFilterer creates a new log filterer instance of AccessControlUpgradeSafe, bound to a specific deployed contract.
func NewAccessControlUpgradeSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlUpgradeSafeFilterer, error) {
	contract, err := bindAccessControlUpgradeSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeSafeFilterer{contract: contract}, nil
}

// bindAccessControlUpgradeSafe binds a generic wrapper to an already deployed contract.
func bindAccessControlUpgradeSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlUpgradeSafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlUpgradeSafe.Contract.AccessControlUpgradeSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.AccessControlUpgradeSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.AccessControlUpgradeSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlUpgradeSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlUpgradeSafe.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlUpgradeSafe.Contract.DEFAULTADMINROLE(&_AccessControlUpgradeSafe.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlUpgradeSafe.Contract.DEFAULTADMINROLE(&_AccessControlUpgradeSafe.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlUpgradeSafe.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlUpgradeSafe.Contract.GetRoleAdmin(&_AccessControlUpgradeSafe.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlUpgradeSafe.Contract.GetRoleAdmin(&_AccessControlUpgradeSafe.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessControlUpgradeSafe.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControlUpgradeSafe.Contract.GetRoleMember(&_AccessControlUpgradeSafe.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControlUpgradeSafe.Contract.GetRoleMember(&_AccessControlUpgradeSafe.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlUpgradeSafe.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControlUpgradeSafe.Contract.GetRoleMemberCount(&_AccessControlUpgradeSafe.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControlUpgradeSafe.Contract.GetRoleMemberCount(&_AccessControlUpgradeSafe.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlUpgradeSafe.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlUpgradeSafe.Contract.HasRole(&_AccessControlUpgradeSafe.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlUpgradeSafe.Contract.HasRole(&_AccessControlUpgradeSafe.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.GrantRole(&_AccessControlUpgradeSafe.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.GrantRole(&_AccessControlUpgradeSafe.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.RenounceRole(&_AccessControlUpgradeSafe.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.RenounceRole(&_AccessControlUpgradeSafe.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.RevokeRole(&_AccessControlUpgradeSafe.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeSafe.Contract.RevokeRole(&_AccessControlUpgradeSafe.TransactOpts, role, account)
}

// AccessControlUpgradeSafeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControlUpgradeSafe contract.
type AccessControlUpgradeSafeRoleGrantedIterator struct {
	Event *AccessControlUpgradeSafeRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeSafeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeSafeRoleGranted)
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
		it.Event = new(AccessControlUpgradeSafeRoleGranted)
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
func (it *AccessControlUpgradeSafeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeSafeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeSafeRoleGranted represents a RoleGranted event raised by the AccessControlUpgradeSafe contract.
type AccessControlUpgradeSafeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlUpgradeSafeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlUpgradeSafe.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeSafeRoleGrantedIterator{contract: _AccessControlUpgradeSafe.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeSafeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlUpgradeSafe.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeSafeRoleGranted)
				if err := _AccessControlUpgradeSafe.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeFilterer) ParseRoleGranted(log types.Log) (*AccessControlUpgradeSafeRoleGranted, error) {
	event := new(AccessControlUpgradeSafeRoleGranted)
	if err := _AccessControlUpgradeSafe.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlUpgradeSafeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControlUpgradeSafe contract.
type AccessControlUpgradeSafeRoleRevokedIterator struct {
	Event *AccessControlUpgradeSafeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeSafeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeSafeRoleRevoked)
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
		it.Event = new(AccessControlUpgradeSafeRoleRevoked)
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
func (it *AccessControlUpgradeSafeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeSafeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeSafeRoleRevoked represents a RoleRevoked event raised by the AccessControlUpgradeSafe contract.
type AccessControlUpgradeSafeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlUpgradeSafeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlUpgradeSafe.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeSafeRoleRevokedIterator{contract: _AccessControlUpgradeSafe.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeSafeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlUpgradeSafe.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeSafeRoleRevoked)
				if err := _AccessControlUpgradeSafe.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeSafe *AccessControlUpgradeSafeFilterer) ParseRoleRevoked(log types.Log) (*AccessControlUpgradeSafeRoleRevoked, error) {
	event := new(AccessControlUpgradeSafeRoleRevoked)
	if err := _AccessControlUpgradeSafe.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
