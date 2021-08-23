// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20BurnableUpgradeSafe

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

// ERC20BurnableUpgradeSafeMetaData contains all meta data concerning the ERC20BurnableUpgradeSafe contract.
var ERC20BurnableUpgradeSafeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20BurnableUpgradeSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BurnableUpgradeSafeMetaData.ABI instead.
var ERC20BurnableUpgradeSafeABI = ERC20BurnableUpgradeSafeMetaData.ABI

// ERC20BurnableUpgradeSafe is an auto generated Go binding around an Ethereum contract.
type ERC20BurnableUpgradeSafe struct {
	ERC20BurnableUpgradeSafeCaller     // Read-only binding to the contract
	ERC20BurnableUpgradeSafeTransactor // Write-only binding to the contract
	ERC20BurnableUpgradeSafeFilterer   // Log filterer for contract events
}

// ERC20BurnableUpgradeSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableUpgradeSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableUpgradeSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableUpgradeSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableUpgradeSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableUpgradeSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableUpgradeSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableUpgradeSafeSession struct {
	Contract     *ERC20BurnableUpgradeSafe // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20BurnableUpgradeSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableUpgradeSafeCallerSession struct {
	Contract *ERC20BurnableUpgradeSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ERC20BurnableUpgradeSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableUpgradeSafeTransactorSession struct {
	Contract     *ERC20BurnableUpgradeSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ERC20BurnableUpgradeSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableUpgradeSafeRaw struct {
	Contract *ERC20BurnableUpgradeSafe // Generic contract binding to access the raw methods on
}

// ERC20BurnableUpgradeSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableUpgradeSafeCallerRaw struct {
	Contract *ERC20BurnableUpgradeSafeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableUpgradeSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableUpgradeSafeTransactorRaw struct {
	Contract *ERC20BurnableUpgradeSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20BurnableUpgradeSafe creates a new instance of ERC20BurnableUpgradeSafe, bound to a specific deployed contract.
func NewERC20BurnableUpgradeSafe(address common.Address, backend bind.ContractBackend) (*ERC20BurnableUpgradeSafe, error) {
	contract, err := bindERC20BurnableUpgradeSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeSafe{ERC20BurnableUpgradeSafeCaller: ERC20BurnableUpgradeSafeCaller{contract: contract}, ERC20BurnableUpgradeSafeTransactor: ERC20BurnableUpgradeSafeTransactor{contract: contract}, ERC20BurnableUpgradeSafeFilterer: ERC20BurnableUpgradeSafeFilterer{contract: contract}}, nil
}

// NewERC20BurnableUpgradeSafeCaller creates a new read-only instance of ERC20BurnableUpgradeSafe, bound to a specific deployed contract.
func NewERC20BurnableUpgradeSafeCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableUpgradeSafeCaller, error) {
	contract, err := bindERC20BurnableUpgradeSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeSafeCaller{contract: contract}, nil
}

// NewERC20BurnableUpgradeSafeTransactor creates a new write-only instance of ERC20BurnableUpgradeSafe, bound to a specific deployed contract.
func NewERC20BurnableUpgradeSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableUpgradeSafeTransactor, error) {
	contract, err := bindERC20BurnableUpgradeSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeSafeTransactor{contract: contract}, nil
}

// NewERC20BurnableUpgradeSafeFilterer creates a new log filterer instance of ERC20BurnableUpgradeSafe, bound to a specific deployed contract.
func NewERC20BurnableUpgradeSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableUpgradeSafeFilterer, error) {
	contract, err := bindERC20BurnableUpgradeSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeSafeFilterer{contract: contract}, nil
}

// bindERC20BurnableUpgradeSafe binds a generic wrapper to an already deployed contract.
func bindERC20BurnableUpgradeSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BurnableUpgradeSafeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20BurnableUpgradeSafe.Contract.ERC20BurnableUpgradeSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.ERC20BurnableUpgradeSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.ERC20BurnableUpgradeSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20BurnableUpgradeSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeSafe.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Allowance(&_ERC20BurnableUpgradeSafe.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Allowance(&_ERC20BurnableUpgradeSafe.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeSafe.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20BurnableUpgradeSafe.Contract.BalanceOf(&_ERC20BurnableUpgradeSafe.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20BurnableUpgradeSafe.Contract.BalanceOf(&_ERC20BurnableUpgradeSafe.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeSafe.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) Decimals() (uint8, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Decimals(&_ERC20BurnableUpgradeSafe.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCallerSession) Decimals() (uint8, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Decimals(&_ERC20BurnableUpgradeSafe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeSafe.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) Name() (string, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Name(&_ERC20BurnableUpgradeSafe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCallerSession) Name() (string, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Name(&_ERC20BurnableUpgradeSafe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeSafe.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) Symbol() (string, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Symbol(&_ERC20BurnableUpgradeSafe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCallerSession) Symbol() (string, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Symbol(&_ERC20BurnableUpgradeSafe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeSafe.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) TotalSupply() (*big.Int, error) {
	return _ERC20BurnableUpgradeSafe.Contract.TotalSupply(&_ERC20BurnableUpgradeSafe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20BurnableUpgradeSafe.Contract.TotalSupply(&_ERC20BurnableUpgradeSafe.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Approve(&_ERC20BurnableUpgradeSafe.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Approve(&_ERC20BurnableUpgradeSafe.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Burn(&_ERC20BurnableUpgradeSafe.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Burn(&_ERC20BurnableUpgradeSafe.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.BurnFrom(&_ERC20BurnableUpgradeSafe.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.BurnFrom(&_ERC20BurnableUpgradeSafe.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.DecreaseAllowance(&_ERC20BurnableUpgradeSafe.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.DecreaseAllowance(&_ERC20BurnableUpgradeSafe.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.IncreaseAllowance(&_ERC20BurnableUpgradeSafe.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.IncreaseAllowance(&_ERC20BurnableUpgradeSafe.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Transfer(&_ERC20BurnableUpgradeSafe.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.Transfer(&_ERC20BurnableUpgradeSafe.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.TransferFrom(&_ERC20BurnableUpgradeSafe.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeSafe.Contract.TransferFrom(&_ERC20BurnableUpgradeSafe.TransactOpts, sender, recipient, amount)
}

// ERC20BurnableUpgradeSafeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20BurnableUpgradeSafe contract.
type ERC20BurnableUpgradeSafeApprovalIterator struct {
	Event *ERC20BurnableUpgradeSafeApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableUpgradeSafeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableUpgradeSafeApproval)
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
		it.Event = new(ERC20BurnableUpgradeSafeApproval)
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
func (it *ERC20BurnableUpgradeSafeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableUpgradeSafeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableUpgradeSafeApproval represents a Approval event raised by the ERC20BurnableUpgradeSafe contract.
type ERC20BurnableUpgradeSafeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableUpgradeSafeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20BurnableUpgradeSafe.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeSafeApprovalIterator{contract: _ERC20BurnableUpgradeSafe.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableUpgradeSafeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20BurnableUpgradeSafe.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableUpgradeSafeApproval)
				if err := _ERC20BurnableUpgradeSafe.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeFilterer) ParseApproval(log types.Log) (*ERC20BurnableUpgradeSafeApproval, error) {
	event := new(ERC20BurnableUpgradeSafeApproval)
	if err := _ERC20BurnableUpgradeSafe.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableUpgradeSafeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20BurnableUpgradeSafe contract.
type ERC20BurnableUpgradeSafeTransferIterator struct {
	Event *ERC20BurnableUpgradeSafeTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableUpgradeSafeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableUpgradeSafeTransfer)
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
		it.Event = new(ERC20BurnableUpgradeSafeTransfer)
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
func (it *ERC20BurnableUpgradeSafeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableUpgradeSafeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableUpgradeSafeTransfer represents a Transfer event raised by the ERC20BurnableUpgradeSafe contract.
type ERC20BurnableUpgradeSafeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableUpgradeSafeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20BurnableUpgradeSafe.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeSafeTransferIterator{contract: _ERC20BurnableUpgradeSafe.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableUpgradeSafeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20BurnableUpgradeSafe.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableUpgradeSafeTransfer)
				if err := _ERC20BurnableUpgradeSafe.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20BurnableUpgradeSafe *ERC20BurnableUpgradeSafeFilterer) ParseTransfer(log types.Log) (*ERC20BurnableUpgradeSafeTransfer, error) {
	event := new(ERC20BurnableUpgradeSafeTransfer)
	if err := _ERC20BurnableUpgradeSafe.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
