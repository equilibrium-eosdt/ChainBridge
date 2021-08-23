// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20PausableUpgradeSafe

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

// ERC20PausableUpgradeSafeMetaData contains all meta data concerning the ERC20PausableUpgradeSafe contract.
var ERC20PausableUpgradeSafeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20PausableUpgradeSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20PausableUpgradeSafeMetaData.ABI instead.
var ERC20PausableUpgradeSafeABI = ERC20PausableUpgradeSafeMetaData.ABI

// ERC20PausableUpgradeSafe is an auto generated Go binding around an Ethereum contract.
type ERC20PausableUpgradeSafe struct {
	ERC20PausableUpgradeSafeCaller     // Read-only binding to the contract
	ERC20PausableUpgradeSafeTransactor // Write-only binding to the contract
	ERC20PausableUpgradeSafeFilterer   // Log filterer for contract events
}

// ERC20PausableUpgradeSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PausableUpgradeSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PausableUpgradeSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PausableUpgradeSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PausableUpgradeSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PausableUpgradeSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PausableUpgradeSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PausableUpgradeSafeSession struct {
	Contract     *ERC20PausableUpgradeSafe // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20PausableUpgradeSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PausableUpgradeSafeCallerSession struct {
	Contract *ERC20PausableUpgradeSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ERC20PausableUpgradeSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PausableUpgradeSafeTransactorSession struct {
	Contract     *ERC20PausableUpgradeSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ERC20PausableUpgradeSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PausableUpgradeSafeRaw struct {
	Contract *ERC20PausableUpgradeSafe // Generic contract binding to access the raw methods on
}

// ERC20PausableUpgradeSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PausableUpgradeSafeCallerRaw struct {
	Contract *ERC20PausableUpgradeSafeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PausableUpgradeSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PausableUpgradeSafeTransactorRaw struct {
	Contract *ERC20PausableUpgradeSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20PausableUpgradeSafe creates a new instance of ERC20PausableUpgradeSafe, bound to a specific deployed contract.
func NewERC20PausableUpgradeSafe(address common.Address, backend bind.ContractBackend) (*ERC20PausableUpgradeSafe, error) {
	contract, err := bindERC20PausableUpgradeSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUpgradeSafe{ERC20PausableUpgradeSafeCaller: ERC20PausableUpgradeSafeCaller{contract: contract}, ERC20PausableUpgradeSafeTransactor: ERC20PausableUpgradeSafeTransactor{contract: contract}, ERC20PausableUpgradeSafeFilterer: ERC20PausableUpgradeSafeFilterer{contract: contract}}, nil
}

// NewERC20PausableUpgradeSafeCaller creates a new read-only instance of ERC20PausableUpgradeSafe, bound to a specific deployed contract.
func NewERC20PausableUpgradeSafeCaller(address common.Address, caller bind.ContractCaller) (*ERC20PausableUpgradeSafeCaller, error) {
	contract, err := bindERC20PausableUpgradeSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUpgradeSafeCaller{contract: contract}, nil
}

// NewERC20PausableUpgradeSafeTransactor creates a new write-only instance of ERC20PausableUpgradeSafe, bound to a specific deployed contract.
func NewERC20PausableUpgradeSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PausableUpgradeSafeTransactor, error) {
	contract, err := bindERC20PausableUpgradeSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUpgradeSafeTransactor{contract: contract}, nil
}

// NewERC20PausableUpgradeSafeFilterer creates a new log filterer instance of ERC20PausableUpgradeSafe, bound to a specific deployed contract.
func NewERC20PausableUpgradeSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PausableUpgradeSafeFilterer, error) {
	contract, err := bindERC20PausableUpgradeSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUpgradeSafeFilterer{contract: contract}, nil
}

// bindERC20PausableUpgradeSafe binds a generic wrapper to an already deployed contract.
func bindERC20PausableUpgradeSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PausableUpgradeSafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PausableUpgradeSafe.Contract.ERC20PausableUpgradeSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.ERC20PausableUpgradeSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.ERC20PausableUpgradeSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PausableUpgradeSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PausableUpgradeSafe.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PausableUpgradeSafe.Contract.Allowance(&_ERC20PausableUpgradeSafe.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PausableUpgradeSafe.Contract.Allowance(&_ERC20PausableUpgradeSafe.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PausableUpgradeSafe.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PausableUpgradeSafe.Contract.BalanceOf(&_ERC20PausableUpgradeSafe.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PausableUpgradeSafe.Contract.BalanceOf(&_ERC20PausableUpgradeSafe.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20PausableUpgradeSafe.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) Decimals() (uint8, error) {
	return _ERC20PausableUpgradeSafe.Contract.Decimals(&_ERC20PausableUpgradeSafe.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCallerSession) Decimals() (uint8, error) {
	return _ERC20PausableUpgradeSafe.Contract.Decimals(&_ERC20PausableUpgradeSafe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PausableUpgradeSafe.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) Name() (string, error) {
	return _ERC20PausableUpgradeSafe.Contract.Name(&_ERC20PausableUpgradeSafe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCallerSession) Name() (string, error) {
	return _ERC20PausableUpgradeSafe.Contract.Name(&_ERC20PausableUpgradeSafe.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20PausableUpgradeSafe.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) Paused() (bool, error) {
	return _ERC20PausableUpgradeSafe.Contract.Paused(&_ERC20PausableUpgradeSafe.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCallerSession) Paused() (bool, error) {
	return _ERC20PausableUpgradeSafe.Contract.Paused(&_ERC20PausableUpgradeSafe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PausableUpgradeSafe.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) Symbol() (string, error) {
	return _ERC20PausableUpgradeSafe.Contract.Symbol(&_ERC20PausableUpgradeSafe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCallerSession) Symbol() (string, error) {
	return _ERC20PausableUpgradeSafe.Contract.Symbol(&_ERC20PausableUpgradeSafe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PausableUpgradeSafe.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) TotalSupply() (*big.Int, error) {
	return _ERC20PausableUpgradeSafe.Contract.TotalSupply(&_ERC20PausableUpgradeSafe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20PausableUpgradeSafe.Contract.TotalSupply(&_ERC20PausableUpgradeSafe.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.Approve(&_ERC20PausableUpgradeSafe.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.Approve(&_ERC20PausableUpgradeSafe.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.DecreaseAllowance(&_ERC20PausableUpgradeSafe.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.DecreaseAllowance(&_ERC20PausableUpgradeSafe.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.IncreaseAllowance(&_ERC20PausableUpgradeSafe.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.IncreaseAllowance(&_ERC20PausableUpgradeSafe.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.Transfer(&_ERC20PausableUpgradeSafe.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.Transfer(&_ERC20PausableUpgradeSafe.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.TransferFrom(&_ERC20PausableUpgradeSafe.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PausableUpgradeSafe.Contract.TransferFrom(&_ERC20PausableUpgradeSafe.TransactOpts, sender, recipient, amount)
}

// ERC20PausableUpgradeSafeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20PausableUpgradeSafe contract.
type ERC20PausableUpgradeSafeApprovalIterator struct {
	Event *ERC20PausableUpgradeSafeApproval // Event containing the contract specifics and raw log

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
func (it *ERC20PausableUpgradeSafeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PausableUpgradeSafeApproval)
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
		it.Event = new(ERC20PausableUpgradeSafeApproval)
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
func (it *ERC20PausableUpgradeSafeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PausableUpgradeSafeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PausableUpgradeSafeApproval represents a Approval event raised by the ERC20PausableUpgradeSafe contract.
type ERC20PausableUpgradeSafeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20PausableUpgradeSafeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PausableUpgradeSafe.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUpgradeSafeApprovalIterator{contract: _ERC20PausableUpgradeSafe.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20PausableUpgradeSafeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PausableUpgradeSafe.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PausableUpgradeSafeApproval)
				if err := _ERC20PausableUpgradeSafe.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) ParseApproval(log types.Log) (*ERC20PausableUpgradeSafeApproval, error) {
	event := new(ERC20PausableUpgradeSafeApproval)
	if err := _ERC20PausableUpgradeSafe.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PausableUpgradeSafePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20PausableUpgradeSafe contract.
type ERC20PausableUpgradeSafePausedIterator struct {
	Event *ERC20PausableUpgradeSafePaused // Event containing the contract specifics and raw log

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
func (it *ERC20PausableUpgradeSafePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PausableUpgradeSafePaused)
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
		it.Event = new(ERC20PausableUpgradeSafePaused)
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
func (it *ERC20PausableUpgradeSafePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PausableUpgradeSafePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PausableUpgradeSafePaused represents a Paused event raised by the ERC20PausableUpgradeSafe contract.
type ERC20PausableUpgradeSafePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20PausableUpgradeSafePausedIterator, error) {

	logs, sub, err := _ERC20PausableUpgradeSafe.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUpgradeSafePausedIterator{contract: _ERC20PausableUpgradeSafe.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20PausableUpgradeSafePaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PausableUpgradeSafe.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PausableUpgradeSafePaused)
				if err := _ERC20PausableUpgradeSafe.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) ParsePaused(log types.Log) (*ERC20PausableUpgradeSafePaused, error) {
	event := new(ERC20PausableUpgradeSafePaused)
	if err := _ERC20PausableUpgradeSafe.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PausableUpgradeSafeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20PausableUpgradeSafe contract.
type ERC20PausableUpgradeSafeTransferIterator struct {
	Event *ERC20PausableUpgradeSafeTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20PausableUpgradeSafeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PausableUpgradeSafeTransfer)
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
		it.Event = new(ERC20PausableUpgradeSafeTransfer)
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
func (it *ERC20PausableUpgradeSafeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PausableUpgradeSafeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PausableUpgradeSafeTransfer represents a Transfer event raised by the ERC20PausableUpgradeSafe contract.
type ERC20PausableUpgradeSafeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20PausableUpgradeSafeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PausableUpgradeSafe.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUpgradeSafeTransferIterator{contract: _ERC20PausableUpgradeSafe.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20PausableUpgradeSafeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PausableUpgradeSafe.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PausableUpgradeSafeTransfer)
				if err := _ERC20PausableUpgradeSafe.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) ParseTransfer(log types.Log) (*ERC20PausableUpgradeSafeTransfer, error) {
	event := new(ERC20PausableUpgradeSafeTransfer)
	if err := _ERC20PausableUpgradeSafe.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PausableUpgradeSafeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20PausableUpgradeSafe contract.
type ERC20PausableUpgradeSafeUnpausedIterator struct {
	Event *ERC20PausableUpgradeSafeUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20PausableUpgradeSafeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PausableUpgradeSafeUnpaused)
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
		it.Event = new(ERC20PausableUpgradeSafeUnpaused)
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
func (it *ERC20PausableUpgradeSafeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PausableUpgradeSafeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PausableUpgradeSafeUnpaused represents a Unpaused event raised by the ERC20PausableUpgradeSafe contract.
type ERC20PausableUpgradeSafeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20PausableUpgradeSafeUnpausedIterator, error) {

	logs, sub, err := _ERC20PausableUpgradeSafe.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20PausableUpgradeSafeUnpausedIterator{contract: _ERC20PausableUpgradeSafe.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20PausableUpgradeSafeUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PausableUpgradeSafe.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PausableUpgradeSafeUnpaused)
				if err := _ERC20PausableUpgradeSafe.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ERC20PausableUpgradeSafe *ERC20PausableUpgradeSafeFilterer) ParseUnpaused(log types.Log) (*ERC20PausableUpgradeSafeUnpaused, error) {
	event := new(ERC20PausableUpgradeSafeUnpaused)
	if err := _ERC20PausableUpgradeSafe.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
