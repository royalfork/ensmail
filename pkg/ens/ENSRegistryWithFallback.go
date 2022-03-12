// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ens

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

// ENSMetaData contains all meta data concerning the ENS contract.
var ENSMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e985e9c5": "isApprovedForAll(address,address)",
		"02571be3": "owner(bytes32)",
		"f79fe538": "recordExists(bytes32)",
		"0178b8bf": "resolver(bytes32)",
		"a22cb465": "setApprovalForAll(address,bool)",
		"5b0fc9c3": "setOwner(bytes32,address)",
		"cf408823": "setRecord(bytes32,address,address,uint64)",
		"1896f70a": "setResolver(bytes32,address)",
		"06ab5923": "setSubnodeOwner(bytes32,bytes32,address)",
		"5ef2c7f0": "setSubnodeRecord(bytes32,bytes32,address,address,uint64)",
		"14ab9038": "setTTL(bytes32,uint64)",
		"16a25cbd": "ttl(bytes32)",
	},
}

// ENSABI is the input ABI used to generate the binding from.
// Deprecated: Use ENSMetaData.ABI instead.
var ENSABI = ENSMetaData.ABI

// Deprecated: Use ENSMetaData.Sigs instead.
// ENSFuncSigs maps the 4-byte function signature to its string representation.
var ENSFuncSigs = ENSMetaData.Sigs

// ENS is an auto generated Go binding around an Ethereum contract.
type ENS struct {
	ENSCaller     // Read-only binding to the contract
	ENSTransactor // Write-only binding to the contract
	ENSFilterer   // Log filterer for contract events
}

// ENSCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSSession struct {
	Contract     *ENS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSCallerSession struct {
	Contract *ENSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ENSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSTransactorSession struct {
	Contract     *ENSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSRaw struct {
	Contract *ENS // Generic contract binding to access the raw methods on
}

// ENSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSCallerRaw struct {
	Contract *ENSCaller // Generic read-only contract binding to access the raw methods on
}

// ENSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSTransactorRaw struct {
	Contract *ENSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENS creates a new instance of ENS, bound to a specific deployed contract.
func NewENS(address common.Address, backend bind.ContractBackend) (*ENS, error) {
	contract, err := bindENS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENS{ENSCaller: ENSCaller{contract: contract}, ENSTransactor: ENSTransactor{contract: contract}, ENSFilterer: ENSFilterer{contract: contract}}, nil
}

// NewENSCaller creates a new read-only instance of ENS, bound to a specific deployed contract.
func NewENSCaller(address common.Address, caller bind.ContractCaller) (*ENSCaller, error) {
	contract, err := bindENS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSCaller{contract: contract}, nil
}

// NewENSTransactor creates a new write-only instance of ENS, bound to a specific deployed contract.
func NewENSTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSTransactor, error) {
	contract, err := bindENS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSTransactor{contract: contract}, nil
}

// NewENSFilterer creates a new log filterer instance of ENS, bound to a specific deployed contract.
func NewENSFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSFilterer, error) {
	contract, err := bindENS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSFilterer{contract: contract}, nil
}

// bindENS binds a generic wrapper to an already deployed contract.
func bindENS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENS *ENSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENS.Contract.ENSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENS *ENSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENS.Contract.ENSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENS *ENSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENS.Contract.ENSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENS *ENSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENS *ENSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENS *ENSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENS.Contract.contract.Transact(opts, method, params...)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENS *ENSCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENS *ENSSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ENS.Contract.IsApprovedForAll(&_ENS.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENS *ENSCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ENS.Contract.IsApprovedForAll(&_ENS.CallOpts, owner, operator)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "owner", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSSession) Owner(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Owner(&_ENS.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Owner(&_ENS.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENS *ENSCaller) RecordExists(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "recordExists", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENS *ENSSession) RecordExists(node [32]byte) (bool, error) {
	return _ENS.Contract.RecordExists(&_ENS.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENS *ENSCallerSession) RecordExists(node [32]byte) (bool, error) {
	return _ENS.Contract.RecordExists(&_ENS.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENS *ENSCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "resolver", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENS *ENSSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Resolver(&_ENS.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENS *ENSCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Resolver(&_ENS.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENS *ENSCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "ttl", node)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENS *ENSSession) Ttl(node [32]byte) (uint64, error) {
	return _ENS.Contract.Ttl(&_ENS.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENS *ENSCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _ENS.Contract.Ttl(&_ENS.CallOpts, node)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENS *ENSTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENS *ENSSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENS.Contract.SetApprovalForAll(&_ENS.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENS *ENSTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENS.Contract.SetApprovalForAll(&_ENS.TransactOpts, operator, approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENS *ENSTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENS *ENSSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetOwner(&_ENS.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENS *ENSTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetOwner(&_ENS.TransactOpts, node, owner)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSTransactor) SetRecord(opts *bind.TransactOpts, node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setRecord", node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetRecord(&_ENS.TransactOpts, node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSTransactorSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetRecord(&_ENS.TransactOpts, node, owner, resolver, ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENS *ENSTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENS *ENSSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetResolver(&_ENS.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENS *ENSTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetResolver(&_ENS.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENS *ENSTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENS *ENSSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeOwner(&_ENS.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENS *ENSTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeOwner(&_ENS.TransactOpts, node, label, owner)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSTransactor) SetSubnodeRecord(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setSubnodeRecord", node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeRecord(&_ENS.TransactOpts, node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENS *ENSTransactorSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeRecord(&_ENS.TransactOpts, node, label, owner, resolver, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENS *ENSTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENS *ENSSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetTTL(&_ENS.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENS *ENSTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENS.Contract.SetTTL(&_ENS.TransactOpts, node, ttl)
}

// ENSApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ENS contract.
type ENSApprovalForAllIterator struct {
	Event *ENSApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ENSApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSApprovalForAll)
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
		it.Event = new(ENSApprovalForAll)
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
func (it *ENSApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSApprovalForAll represents a ApprovalForAll event raised by the ENS contract.
type ENSApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENS *ENSFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ENSApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ENSApprovalForAllIterator{contract: _ENS.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENS *ENSFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ENSApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSApprovalForAll)
				if err := _ENS.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENS *ENSFilterer) ParseApprovalForAll(log types.Log) (*ENSApprovalForAll, error) {
	event := new(ENSApprovalForAll)
	if err := _ENS.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the ENS contract.
type ENSNewOwnerIterator struct {
	Event *ENSNewOwner // Event containing the contract specifics and raw log

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
func (it *ENSNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSNewOwner)
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
		it.Event = new(ENSNewOwner)
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
func (it *ENSNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSNewOwner represents a NewOwner event raised by the ENS contract.
type ENSNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENS *ENSFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*ENSNewOwnerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &ENSNewOwnerIterator{contract: _ENS.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENS *ENSFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ENSNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSNewOwner)
				if err := _ENS.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENS *ENSFilterer) ParseNewOwner(log types.Log) (*ENSNewOwner, error) {
	event := new(ENSNewOwner)
	if err := _ENS.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the ENS contract.
type ENSNewResolverIterator struct {
	Event *ENSNewResolver // Event containing the contract specifics and raw log

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
func (it *ENSNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSNewResolver)
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
		it.Event = new(ENSNewResolver)
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
func (it *ENSNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSNewResolver represents a NewResolver event raised by the ENS contract.
type ENSNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENS *ENSFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*ENSNewResolverIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSNewResolverIterator{contract: _ENS.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENS *ENSFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *ENSNewResolver, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSNewResolver)
				if err := _ENS.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ParseNewResolver is a log parse operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENS *ENSFilterer) ParseNewResolver(log types.Log) (*ENSNewResolver, error) {
	event := new(ENSNewResolver)
	if err := _ENS.contract.UnpackLog(event, "NewResolver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the ENS contract.
type ENSNewTTLIterator struct {
	Event *ENSNewTTL // Event containing the contract specifics and raw log

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
func (it *ENSNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSNewTTL)
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
		it.Event = new(ENSNewTTL)
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
func (it *ENSNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSNewTTL represents a NewTTL event raised by the ENS contract.
type ENSNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENS *ENSFilterer) FilterNewTTL(opts *bind.FilterOpts, node [][32]byte) (*ENSNewTTLIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSNewTTLIterator{contract: _ENS.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENS *ENSFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *ENSNewTTL, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSNewTTL)
				if err := _ENS.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ParseNewTTL is a log parse operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENS *ENSFilterer) ParseNewTTL(log types.Log) (*ENSNewTTL, error) {
	event := new(ENSNewTTL)
	if err := _ENS.contract.UnpackLog(event, "NewTTL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ENS contract.
type ENSTransferIterator struct {
	Event *ENSTransfer // Event containing the contract specifics and raw log

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
func (it *ENSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSTransfer)
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
		it.Event = new(ENSTransfer)
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
func (it *ENSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSTransfer represents a Transfer event raised by the ENS contract.
type ENSTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENS *ENSFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*ENSTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSTransferIterator{contract: _ENS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENS *ENSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ENSTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSTransfer)
				if err := _ENS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENS *ENSFilterer) ParseTransfer(log types.Log) (*ENSTransfer, error) {
	event := new(ENSTransfer)
	if err := _ENS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryMetaData contains all meta data concerning the ENSRegistry contract.
var ENSRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e985e9c5": "isApprovedForAll(address,address)",
		"02571be3": "owner(bytes32)",
		"f79fe538": "recordExists(bytes32)",
		"0178b8bf": "resolver(bytes32)",
		"a22cb465": "setApprovalForAll(address,bool)",
		"5b0fc9c3": "setOwner(bytes32,address)",
		"cf408823": "setRecord(bytes32,address,address,uint64)",
		"1896f70a": "setResolver(bytes32,address)",
		"06ab5923": "setSubnodeOwner(bytes32,bytes32,address)",
		"5ef2c7f0": "setSubnodeRecord(bytes32,bytes32,address,address,uint64)",
		"14ab9038": "setTTL(bytes32,uint64)",
		"16a25cbd": "ttl(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b03191633179055610939806100596000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80635b0fc9c3116100715780635b0fc9c3146101e85780635ef2c7f014610214578063a22cb46514610261578063cf4088231461028f578063e985e9c5146102d6578063f79fe53814610318576100b4565b80630178b8bf146100b957806302571be3146100f257806306ab59231461010f57806314ab90381461015357806316a25cbd146101825780631896f70a146101bc575b600080fd5b6100d6600480360360208110156100cf57600080fd5b5035610335565b604080516001600160a01b039092168252519081900360200190f35b6100d66004803603602081101561010857600080fd5b5035610356565b6101416004803603606081101561012557600080fd5b50803590602081013590604001356001600160a01b0316610386565b60408051918252519081900360200190f35b6101806004803603604081101561016957600080fd5b508035906020013567ffffffffffffffff16610456565b005b61019f6004803603602081101561019857600080fd5b503561052a565b6040805167ffffffffffffffff9092168252519081900360200190f35b610180600480360360408110156101d257600080fd5b50803590602001356001600160a01b0316610550565b610180600480360360408110156101fe57600080fd5b50803590602001356001600160a01b0316610616565b610180600480360360a081101561022a57600080fd5b5080359060208101359060408101356001600160a01b03908116916060810135909116906080013567ffffffffffffffff166106b9565b6101806004803603604081101561027757600080fd5b506001600160a01b03813516906020013515156106db565b610180600480360360808110156102a557600080fd5b5080359060208101356001600160a01b03908116916040810135909116906060013567ffffffffffffffff16610749565b610304600480360360408110156102ec57600080fd5b506001600160a01b0381358116916020013516610764565b604080519115158252519081900360200190f35b6103046004803603602081101561032e57600080fd5b5035610792565b6000818152602081905260409020600101546001600160a01b03165b919050565b6000818152602081905260408120546001600160a01b031630811415610380576000915050610351565b92915050565b60008381526020819052604081205484906001600160a01b0316338114806103d157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6103da57600080fd5b6040805160208082018990528183018890528251808303840181526060909201909252805191012061040c81866107af565b604080516001600160a01b03871681529051879189917fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e829181900360200190a39695505050505050565b60008281526020819052604090205482906001600160a01b0316338114806104a157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6104aa57600080fd5b6040805167ffffffffffffffff85168152905185917f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68919081900360200190a25050600091825260208290526040909120600101805467ffffffffffffffff909216600160a01b0267ffffffffffffffff60a01b19909216919091179055565b600090815260208190526040902060010154600160a01b900467ffffffffffffffff1690565b60008281526020819052604090205482906001600160a01b03163381148061059b57506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6105a457600080fd5b604080516001600160a01b0385168152905185917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a2505060009182526020829052604090912060010180546001600160a01b0319166001600160a01b03909216919091179055565b60008281526020819052604090205482906001600160a01b03163381148061066157506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b61066a57600080fd5b61067484846107af565b604080516001600160a01b0385168152905185917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a250505050565b60006106c6868686610386565b90506106d38184846107dd565b505050505050565b3360008181526001602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6107538484610616565b61075e8483836107dd565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6000908152602081905260409020546001600160a01b0316151590565b60009182526020829052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6000838152602081905260409020600101546001600160a01b03838116911614610863576000838152602081815260409182902060010180546001600160a01b0319166001600160a01b0386169081179091558251908152915185927f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a092908290030190a25b60008381526020819052604090206001015467ffffffffffffffff828116600160a01b90920416146108ff5760008381526020818152604091829020600101805467ffffffffffffffff60a01b1916600160a01b67ffffffffffffffff8616908102919091179091558251908152915185927f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa6892908290030190a25b50505056fea265627a7a72315820377341ef1aa4e60845c51e945d8fce1132a6b34b837c85a9e7c4737e2c748f3d64736f6c63430005110032",
}

// ENSRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ENSRegistryMetaData.ABI instead.
var ENSRegistryABI = ENSRegistryMetaData.ABI

// Deprecated: Use ENSRegistryMetaData.Sigs instead.
// ENSRegistryFuncSigs maps the 4-byte function signature to its string representation.
var ENSRegistryFuncSigs = ENSRegistryMetaData.Sigs

// ENSRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ENSRegistryMetaData.Bin instead.
var ENSRegistryBin = ENSRegistryMetaData.Bin

// DeployENSRegistry deploys a new Ethereum contract, binding an instance of ENSRegistry to it.
func DeployENSRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ENSRegistry, error) {
	parsed, err := ENSRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ENSRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENSRegistry{ENSRegistryCaller: ENSRegistryCaller{contract: contract}, ENSRegistryTransactor: ENSRegistryTransactor{contract: contract}, ENSRegistryFilterer: ENSRegistryFilterer{contract: contract}}, nil
}

// ENSRegistry is an auto generated Go binding around an Ethereum contract.
type ENSRegistry struct {
	ENSRegistryCaller     // Read-only binding to the contract
	ENSRegistryTransactor // Write-only binding to the contract
	ENSRegistryFilterer   // Log filterer for contract events
}

// ENSRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSRegistrySession struct {
	Contract     *ENSRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSRegistryCallerSession struct {
	Contract *ENSRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ENSRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSRegistryTransactorSession struct {
	Contract     *ENSRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ENSRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSRegistryRaw struct {
	Contract *ENSRegistry // Generic contract binding to access the raw methods on
}

// ENSRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSRegistryCallerRaw struct {
	Contract *ENSRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ENSRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSRegistryTransactorRaw struct {
	Contract *ENSRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENSRegistry creates a new instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistry(address common.Address, backend bind.ContractBackend) (*ENSRegistry, error) {
	contract, err := bindENSRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENSRegistry{ENSRegistryCaller: ENSRegistryCaller{contract: contract}, ENSRegistryTransactor: ENSRegistryTransactor{contract: contract}, ENSRegistryFilterer: ENSRegistryFilterer{contract: contract}}, nil
}

// NewENSRegistryCaller creates a new read-only instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryCaller(address common.Address, caller bind.ContractCaller) (*ENSRegistryCaller, error) {
	contract, err := bindENSRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryCaller{contract: contract}, nil
}

// NewENSRegistryTransactor creates a new write-only instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSRegistryTransactor, error) {
	contract, err := bindENSRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryTransactor{contract: contract}, nil
}

// NewENSRegistryFilterer creates a new log filterer instance of ENSRegistry, bound to a specific deployed contract.
func NewENSRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSRegistryFilterer, error) {
	contract, err := bindENSRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryFilterer{contract: contract}, nil
}

// bindENSRegistry binds a generic wrapper to an already deployed contract.
func bindENSRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSRegistry *ENSRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENSRegistry.Contract.ENSRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSRegistry *ENSRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSRegistry.Contract.ENSRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSRegistry *ENSRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSRegistry.Contract.ENSRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSRegistry *ENSRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENSRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSRegistry *ENSRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSRegistry *ENSRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENSRegistry *ENSRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ENSRegistry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENSRegistry *ENSRegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ENSRegistry.Contract.IsApprovedForAll(&_ENSRegistry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENSRegistry *ENSRegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ENSRegistry.Contract.IsApprovedForAll(&_ENSRegistry.CallOpts, owner, operator)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENSRegistry *ENSRegistryCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENSRegistry.contract.Call(opts, &out, "owner", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENSRegistry *ENSRegistrySession) Owner(node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Owner(&_ENSRegistry.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENSRegistry *ENSRegistryCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Owner(&_ENSRegistry.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENSRegistry *ENSRegistryCaller) RecordExists(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _ENSRegistry.contract.Call(opts, &out, "recordExists", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENSRegistry *ENSRegistrySession) RecordExists(node [32]byte) (bool, error) {
	return _ENSRegistry.Contract.RecordExists(&_ENSRegistry.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENSRegistry *ENSRegistryCallerSession) RecordExists(node [32]byte) (bool, error) {
	return _ENSRegistry.Contract.RecordExists(&_ENSRegistry.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENSRegistry *ENSRegistryCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENSRegistry.contract.Call(opts, &out, "resolver", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENSRegistry *ENSRegistrySession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Resolver(&_ENSRegistry.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENSRegistry *ENSRegistryCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSRegistry.Contract.Resolver(&_ENSRegistry.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENSRegistry *ENSRegistryCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var out []interface{}
	err := _ENSRegistry.contract.Call(opts, &out, "ttl", node)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENSRegistry *ENSRegistrySession) Ttl(node [32]byte) (uint64, error) {
	return _ENSRegistry.Contract.Ttl(&_ENSRegistry.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENSRegistry *ENSRegistryCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _ENSRegistry.Contract.Ttl(&_ENSRegistry.CallOpts, node)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENSRegistry *ENSRegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetApprovalForAll(&_ENSRegistry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetApprovalForAll(&_ENSRegistry.TransactOpts, operator, approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistry *ENSRegistrySession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetOwner(&_ENSRegistry.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetOwner(&_ENSRegistry.TransactOpts, node, owner)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetRecord(opts *bind.TransactOpts, node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setRecord", node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistrySession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetRecord(&_ENSRegistry.TransactOpts, node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetRecord(&_ENSRegistry.TransactOpts, node, owner, resolver, ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistry *ENSRegistrySession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetResolver(&_ENSRegistry.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetResolver(&_ENSRegistry.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENSRegistry *ENSRegistryTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENSRegistry *ENSRegistrySession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeOwner(&_ENSRegistry.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENSRegistry *ENSRegistryTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeOwner(&_ENSRegistry.TransactOpts, node, label, owner)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetSubnodeRecord(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setSubnodeRecord", node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistrySession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeRecord(&_ENSRegistry.TransactOpts, node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetSubnodeRecord(&_ENSRegistry.TransactOpts, node, label, owner, resolver, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistryTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistrySession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetTTL(&_ENSRegistry.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistry *ENSRegistryTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistry.Contract.SetTTL(&_ENSRegistry.TransactOpts, node, ttl)
}

// ENSRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ENSRegistry contract.
type ENSRegistryApprovalForAllIterator struct {
	Event *ENSRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ENSRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryApprovalForAll)
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
		it.Event = new(ENSRegistryApprovalForAll)
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
func (it *ENSRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryApprovalForAll represents a ApprovalForAll event raised by the ENSRegistry contract.
type ENSRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENSRegistry *ENSRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ENSRegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryApprovalForAllIterator{contract: _ENSRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENSRegistry *ENSRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ENSRegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryApprovalForAll)
				if err := _ENSRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENSRegistry *ENSRegistryFilterer) ParseApprovalForAll(log types.Log) (*ENSRegistryApprovalForAll, error) {
	event := new(ENSRegistryApprovalForAll)
	if err := _ENSRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the ENSRegistry contract.
type ENSRegistryNewOwnerIterator struct {
	Event *ENSRegistryNewOwner // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewOwner)
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
		it.Event = new(ENSRegistryNewOwner)
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
func (it *ENSRegistryNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewOwner represents a NewOwner event raised by the ENSRegistry contract.
type ENSRegistryNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*ENSRegistryNewOwnerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewOwnerIterator{contract: _ENSRegistry.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewOwner)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENSRegistry *ENSRegistryFilterer) ParseNewOwner(log types.Log) (*ENSRegistryNewOwner, error) {
	event := new(ENSRegistryNewOwner)
	if err := _ENSRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the ENSRegistry contract.
type ENSRegistryNewResolverIterator struct {
	Event *ENSRegistryNewResolver // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewResolver)
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
		it.Event = new(ENSRegistryNewResolver)
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
func (it *ENSRegistryNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewResolver represents a NewResolver event raised by the ENSRegistry contract.
type ENSRegistryNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryNewResolverIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewResolverIterator{contract: _ENSRegistry.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewResolver, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewResolver)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ParseNewResolver is a log parse operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENSRegistry *ENSRegistryFilterer) ParseNewResolver(log types.Log) (*ENSRegistryNewResolver, error) {
	event := new(ENSRegistryNewResolver)
	if err := _ENSRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the ENSRegistry contract.
type ENSRegistryNewTTLIterator struct {
	Event *ENSRegistryNewTTL // Event containing the contract specifics and raw log

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
func (it *ENSRegistryNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryNewTTL)
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
		it.Event = new(ENSRegistryNewTTL)
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
func (it *ENSRegistryNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryNewTTL represents a NewTTL event raised by the ENSRegistry contract.
type ENSRegistryNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENSRegistry *ENSRegistryFilterer) FilterNewTTL(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryNewTTLIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryNewTTLIterator{contract: _ENSRegistry.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENSRegistry *ENSRegistryFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *ENSRegistryNewTTL, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryNewTTL)
				if err := _ENSRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ParseNewTTL is a log parse operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENSRegistry *ENSRegistryFilterer) ParseNewTTL(log types.Log) (*ENSRegistryNewTTL, error) {
	event := new(ENSRegistryNewTTL)
	if err := _ENSRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ENSRegistry contract.
type ENSRegistryTransferIterator struct {
	Event *ENSRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *ENSRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryTransfer)
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
		it.Event = new(ENSRegistryTransfer)
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
func (it *ENSRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryTransfer represents a Transfer event raised by the ENSRegistry contract.
type ENSRegistryTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENSRegistry *ENSRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryTransferIterator{contract: _ENSRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENSRegistry *ENSRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ENSRegistryTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistry.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryTransfer)
				if err := _ENSRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENSRegistry *ENSRegistryFilterer) ParseTransfer(log types.Log) (*ENSRegistryTransfer, error) {
	event := new(ENSRegistryTransfer)
	if err := _ENSRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryWithFallbackMetaData contains all meta data concerning the ENSRegistryWithFallback contract.
var ENSRegistryWithFallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractENS\",\"name\":\"_old\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"old\",\"outputs\":[{\"internalType\":\"contractENS\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e985e9c5": "isApprovedForAll(address,address)",
		"b83f8663": "old()",
		"02571be3": "owner(bytes32)",
		"f79fe538": "recordExists(bytes32)",
		"0178b8bf": "resolver(bytes32)",
		"a22cb465": "setApprovalForAll(address,bool)",
		"5b0fc9c3": "setOwner(bytes32,address)",
		"cf408823": "setRecord(bytes32,address,address,uint64)",
		"1896f70a": "setResolver(bytes32,address)",
		"06ab5923": "setSubnodeOwner(bytes32,bytes32,address)",
		"5ef2c7f0": "setSubnodeRecord(bytes32,bytes32,address,address,uint64)",
		"14ab9038": "setTTL(bytes32,uint64)",
		"16a25cbd": "ttl(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610b82380380610b828339818101604052602081101561003357600080fd5b505160008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b03199081163317909155600280546001600160a01b0390931692909116919091179055610ae9806100996000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80635b0fc9c31161008c578063b83f866311610066578063b83f8663146102aa578063cf408823146102b2578063e985e9c5146102f9578063f79fe5381461033b576100cf565b80635b0fc9c3146102035780635ef2c7f01461022f578063a22cb4651461027c576100cf565b80630178b8bf146100d457806302571be31461010d57806306ab59231461012a57806314ab90381461016e57806316a25cbd1461019d5780631896f70a146101d7575b600080fd5b6100f1600480360360208110156100ea57600080fd5b5035610358565b604080516001600160a01b039092168252519081900360200190f35b6100f16004803603602081101561012357600080fd5b50356103f7565b61015c6004803603606081101561014057600080fd5b50803590602081013590604001356001600160a01b031661045b565b60408051918252519081900360200190f35b61019b6004803603604081101561018457600080fd5b508035906020013567ffffffffffffffff1661052b565b005b6101ba600480360360208110156101b357600080fd5b50356105ff565b6040805167ffffffffffffffff9092168252519081900360200190f35b61019b600480360360408110156101ed57600080fd5b50803590602001356001600160a01b0316610663565b61019b6004803603604081101561021957600080fd5b50803590602001356001600160a01b0316610729565b61019b600480360360a081101561024557600080fd5b5080359060208101359060408101356001600160a01b03908116916060810135909116906080013567ffffffffffffffff166107cc565b61019b6004803603604081101561029257600080fd5b506001600160a01b03813516906020013515156107ee565b6100f161085c565b61019b600480360360808110156102c857600080fd5b5080359060208101356001600160a01b03908116916040810135909116906060013567ffffffffffffffff1661086b565b6103276004803603604081101561030f57600080fd5b506001600160a01b0381358116916020013516610886565b604080519115158252519081900360200190f35b6103276004803603602081101561035157600080fd5b50356108b4565b6000610363826108b4565b6103e65760025460408051630178b8bf60e01b81526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156103b357600080fd5b505afa1580156103c7573d6000803e3d6000fd5b505050506040513d60208110156103dd57600080fd5b505190506103f2565b6103ef826108d1565b90505b919050565b6000610402826108b4565b61045257600254604080516302571be360e01b81526004810185905290516001600160a01b03909216916302571be391602480820192602092909190829003018186803b1580156103b357600080fd5b6103ef826108ef565b60008381526020819052604081205484906001600160a01b0316338114806104a657506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6104af57600080fd5b604080516020808201899052818301889052825180830384018152606090920190925280519101206104e18186610919565b604080516001600160a01b03871681529051879189917fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e829181900360200190a39695505050505050565b60008281526020819052604090205482906001600160a01b03163381148061057657506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b61057f57600080fd5b6040805167ffffffffffffffff85168152905185917f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68919081900360200190a25050600091825260208290526040909120600101805467ffffffffffffffff909216600160a01b0267ffffffffffffffff60a01b19909216919091179055565b600061060a826108b4565b61065a57600254604080516316a25cbd60e01b81526004810185905290516001600160a01b03909216916316a25cbd91602480820192602092909190829003018186803b1580156103b357600080fd5b6103ef8261093a565b60008281526020819052604090205482906001600160a01b0316338114806106ae57506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6106b757600080fd5b604080516001600160a01b0385168152905185917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a2505060009182526020829052604090912060010180546001600160a01b0319166001600160a01b03909216919091179055565b60008281526020819052604090205482906001600160a01b03163381148061077457506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b61077d57600080fd5b6107878484610919565b604080516001600160a01b0385168152905185917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a250505050565b60006107d986868661045b565b90506107e6818484610960565b505050505050565b3360008181526001602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6002546001600160a01b031681565b6108758484610729565b610880848383610960565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6000908152602081905260409020546001600160a01b0316151590565b6000908152602081905260409020600101546001600160a01b031690565b6000818152602081905260408120546001600160a01b0316308114156103ef5760009150506103f2565b806001600160a01b03811661092b5750305b6109358382610a86565b505050565b600090815260208190526040902060010154600160a01b900467ffffffffffffffff1690565b6000838152602081905260409020600101546001600160a01b038381169116146109e6576000838152602081815260409182902060010180546001600160a01b0319166001600160a01b0386169081179091558251908152915185927f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a092908290030190a25b60008381526020819052604090206001015467ffffffffffffffff828116600160a01b90920416146109355760008381526020818152604091829020600101805467ffffffffffffffff60a01b1916600160a01b67ffffffffffffffff8616908102919091179091558251908152915185927f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa6892908290030190a2505050565b60009182526020829052604090912080546001600160a01b0319166001600160a01b0390921691909117905556fea265627a7a723158204fb2c2a4e82d176edaabe1a2ad36a6687f4f32fa32ec60e2ce745063d423878f64736f6c63430005110032",
}

// ENSRegistryWithFallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use ENSRegistryWithFallbackMetaData.ABI instead.
var ENSRegistryWithFallbackABI = ENSRegistryWithFallbackMetaData.ABI

// Deprecated: Use ENSRegistryWithFallbackMetaData.Sigs instead.
// ENSRegistryWithFallbackFuncSigs maps the 4-byte function signature to its string representation.
var ENSRegistryWithFallbackFuncSigs = ENSRegistryWithFallbackMetaData.Sigs

// ENSRegistryWithFallbackBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ENSRegistryWithFallbackMetaData.Bin instead.
var ENSRegistryWithFallbackBin = ENSRegistryWithFallbackMetaData.Bin

// DeployENSRegistryWithFallback deploys a new Ethereum contract, binding an instance of ENSRegistryWithFallback to it.
func DeployENSRegistryWithFallback(auth *bind.TransactOpts, backend bind.ContractBackend, _old common.Address) (common.Address, *types.Transaction, *ENSRegistryWithFallback, error) {
	parsed, err := ENSRegistryWithFallbackMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ENSRegistryWithFallbackBin), backend, _old)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENSRegistryWithFallback{ENSRegistryWithFallbackCaller: ENSRegistryWithFallbackCaller{contract: contract}, ENSRegistryWithFallbackTransactor: ENSRegistryWithFallbackTransactor{contract: contract}, ENSRegistryWithFallbackFilterer: ENSRegistryWithFallbackFilterer{contract: contract}}, nil
}

// ENSRegistryWithFallback is an auto generated Go binding around an Ethereum contract.
type ENSRegistryWithFallback struct {
	ENSRegistryWithFallbackCaller     // Read-only binding to the contract
	ENSRegistryWithFallbackTransactor // Write-only binding to the contract
	ENSRegistryWithFallbackFilterer   // Log filterer for contract events
}

// ENSRegistryWithFallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSRegistryWithFallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryWithFallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSRegistryWithFallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryWithFallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSRegistryWithFallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSRegistryWithFallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSRegistryWithFallbackSession struct {
	Contract     *ENSRegistryWithFallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ENSRegistryWithFallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSRegistryWithFallbackCallerSession struct {
	Contract *ENSRegistryWithFallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ENSRegistryWithFallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSRegistryWithFallbackTransactorSession struct {
	Contract     *ENSRegistryWithFallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ENSRegistryWithFallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSRegistryWithFallbackRaw struct {
	Contract *ENSRegistryWithFallback // Generic contract binding to access the raw methods on
}

// ENSRegistryWithFallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSRegistryWithFallbackCallerRaw struct {
	Contract *ENSRegistryWithFallbackCaller // Generic read-only contract binding to access the raw methods on
}

// ENSRegistryWithFallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSRegistryWithFallbackTransactorRaw struct {
	Contract *ENSRegistryWithFallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENSRegistryWithFallback creates a new instance of ENSRegistryWithFallback, bound to a specific deployed contract.
func NewENSRegistryWithFallback(address common.Address, backend bind.ContractBackend) (*ENSRegistryWithFallback, error) {
	contract, err := bindENSRegistryWithFallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallback{ENSRegistryWithFallbackCaller: ENSRegistryWithFallbackCaller{contract: contract}, ENSRegistryWithFallbackTransactor: ENSRegistryWithFallbackTransactor{contract: contract}, ENSRegistryWithFallbackFilterer: ENSRegistryWithFallbackFilterer{contract: contract}}, nil
}

// NewENSRegistryWithFallbackCaller creates a new read-only instance of ENSRegistryWithFallback, bound to a specific deployed contract.
func NewENSRegistryWithFallbackCaller(address common.Address, caller bind.ContractCaller) (*ENSRegistryWithFallbackCaller, error) {
	contract, err := bindENSRegistryWithFallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallbackCaller{contract: contract}, nil
}

// NewENSRegistryWithFallbackTransactor creates a new write-only instance of ENSRegistryWithFallback, bound to a specific deployed contract.
func NewENSRegistryWithFallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSRegistryWithFallbackTransactor, error) {
	contract, err := bindENSRegistryWithFallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallbackTransactor{contract: contract}, nil
}

// NewENSRegistryWithFallbackFilterer creates a new log filterer instance of ENSRegistryWithFallback, bound to a specific deployed contract.
func NewENSRegistryWithFallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSRegistryWithFallbackFilterer, error) {
	contract, err := bindENSRegistryWithFallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallbackFilterer{contract: contract}, nil
}

// bindENSRegistryWithFallback binds a generic wrapper to an already deployed contract.
func bindENSRegistryWithFallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSRegistryWithFallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENSRegistryWithFallback.Contract.ENSRegistryWithFallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.ENSRegistryWithFallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.ENSRegistryWithFallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENSRegistryWithFallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.contract.Transact(opts, method, params...)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ENSRegistryWithFallback.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ENSRegistryWithFallback.Contract.IsApprovedForAll(&_ENSRegistryWithFallback.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ENSRegistryWithFallback.Contract.IsApprovedForAll(&_ENSRegistryWithFallback.CallOpts, owner, operator)
}

// Old is a free data retrieval call binding the contract method 0xb83f8663.
//
// Solidity: function old() view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCaller) Old(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ENSRegistryWithFallback.contract.Call(opts, &out, "old")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Old is a free data retrieval call binding the contract method 0xb83f8663.
//
// Solidity: function old() view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) Old() (common.Address, error) {
	return _ENSRegistryWithFallback.Contract.Old(&_ENSRegistryWithFallback.CallOpts)
}

// Old is a free data retrieval call binding the contract method 0xb83f8663.
//
// Solidity: function old() view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCallerSession) Old() (common.Address, error) {
	return _ENSRegistryWithFallback.Contract.Old(&_ENSRegistryWithFallback.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENSRegistryWithFallback.contract.Call(opts, &out, "owner", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) Owner(node [32]byte) (common.Address, error) {
	return _ENSRegistryWithFallback.Contract.Owner(&_ENSRegistryWithFallback.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENSRegistryWithFallback.Contract.Owner(&_ENSRegistryWithFallback.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCaller) RecordExists(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _ENSRegistryWithFallback.contract.Call(opts, &out, "recordExists", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) RecordExists(node [32]byte) (bool, error) {
	return _ENSRegistryWithFallback.Contract.RecordExists(&_ENSRegistryWithFallback.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCallerSession) RecordExists(node [32]byte) (bool, error) {
	return _ENSRegistryWithFallback.Contract.RecordExists(&_ENSRegistryWithFallback.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENSRegistryWithFallback.contract.Call(opts, &out, "resolver", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSRegistryWithFallback.Contract.Resolver(&_ENSRegistryWithFallback.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENSRegistryWithFallback.Contract.Resolver(&_ENSRegistryWithFallback.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var out []interface{}
	err := _ENSRegistryWithFallback.contract.Call(opts, &out, "ttl", node)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) Ttl(node [32]byte) (uint64, error) {
	return _ENSRegistryWithFallback.Contract.Ttl(&_ENSRegistryWithFallback.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _ENSRegistryWithFallback.Contract.Ttl(&_ENSRegistryWithFallback.CallOpts, node)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetApprovalForAll(&_ENSRegistryWithFallback.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetApprovalForAll(&_ENSRegistryWithFallback.TransactOpts, operator, approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetOwner(&_ENSRegistryWithFallback.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetOwner(&_ENSRegistryWithFallback.TransactOpts, node, owner)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactor) SetRecord(opts *bind.TransactOpts, node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.contract.Transact(opts, "setRecord", node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetRecord(&_ENSRegistryWithFallback.TransactOpts, node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetRecord(&_ENSRegistryWithFallback.TransactOpts, node, owner, resolver, ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetResolver(&_ENSRegistryWithFallback.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetResolver(&_ENSRegistryWithFallback.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetSubnodeOwner(&_ENSRegistryWithFallback.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetSubnodeOwner(&_ENSRegistryWithFallback.TransactOpts, node, label, owner)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactor) SetSubnodeRecord(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.contract.Transact(opts, "setSubnodeRecord", node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetSubnodeRecord(&_ENSRegistryWithFallback.TransactOpts, node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetSubnodeRecord(&_ENSRegistryWithFallback.TransactOpts, node, label, owner, resolver, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetTTL(&_ENSRegistryWithFallback.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _ENSRegistryWithFallback.Contract.SetTTL(&_ENSRegistryWithFallback.TransactOpts, node, ttl)
}

// ENSRegistryWithFallbackApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackApprovalForAllIterator struct {
	Event *ENSRegistryWithFallbackApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ENSRegistryWithFallbackApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryWithFallbackApprovalForAll)
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
		it.Event = new(ENSRegistryWithFallbackApprovalForAll)
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
func (it *ENSRegistryWithFallbackApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryWithFallbackApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryWithFallbackApprovalForAll represents a ApprovalForAll event raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ENSRegistryWithFallbackApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallbackApprovalForAllIterator{contract: _ENSRegistryWithFallback.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ENSRegistryWithFallbackApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryWithFallbackApprovalForAll)
				if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) ParseApprovalForAll(log types.Log) (*ENSRegistryWithFallbackApprovalForAll, error) {
	event := new(ENSRegistryWithFallbackApprovalForAll)
	if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryWithFallbackNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackNewOwnerIterator struct {
	Event *ENSRegistryWithFallbackNewOwner // Event containing the contract specifics and raw log

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
func (it *ENSRegistryWithFallbackNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryWithFallbackNewOwner)
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
		it.Event = new(ENSRegistryWithFallbackNewOwner)
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
func (it *ENSRegistryWithFallbackNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryWithFallbackNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryWithFallbackNewOwner represents a NewOwner event raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*ENSRegistryWithFallbackNewOwnerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallbackNewOwnerIterator{contract: _ENSRegistryWithFallback.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ENSRegistryWithFallbackNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryWithFallbackNewOwner)
				if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) ParseNewOwner(log types.Log) (*ENSRegistryWithFallbackNewOwner, error) {
	event := new(ENSRegistryWithFallbackNewOwner)
	if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryWithFallbackNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackNewResolverIterator struct {
	Event *ENSRegistryWithFallbackNewResolver // Event containing the contract specifics and raw log

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
func (it *ENSRegistryWithFallbackNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryWithFallbackNewResolver)
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
		it.Event = new(ENSRegistryWithFallbackNewResolver)
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
func (it *ENSRegistryWithFallbackNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryWithFallbackNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryWithFallbackNewResolver represents a NewResolver event raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryWithFallbackNewResolverIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallbackNewResolverIterator{contract: _ENSRegistryWithFallback.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *ENSRegistryWithFallbackNewResolver, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryWithFallbackNewResolver)
				if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ParseNewResolver is a log parse operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) ParseNewResolver(log types.Log) (*ENSRegistryWithFallbackNewResolver, error) {
	event := new(ENSRegistryWithFallbackNewResolver)
	if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "NewResolver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryWithFallbackNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackNewTTLIterator struct {
	Event *ENSRegistryWithFallbackNewTTL // Event containing the contract specifics and raw log

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
func (it *ENSRegistryWithFallbackNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryWithFallbackNewTTL)
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
		it.Event = new(ENSRegistryWithFallbackNewTTL)
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
func (it *ENSRegistryWithFallbackNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryWithFallbackNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryWithFallbackNewTTL represents a NewTTL event raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) FilterNewTTL(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryWithFallbackNewTTLIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.FilterLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallbackNewTTLIterator{contract: _ENSRegistryWithFallback.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *ENSRegistryWithFallbackNewTTL, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.WatchLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryWithFallbackNewTTL)
				if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ParseNewTTL is a log parse operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) ParseNewTTL(log types.Log) (*ENSRegistryWithFallbackNewTTL, error) {
	event := new(ENSRegistryWithFallbackNewTTL)
	if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "NewTTL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ENSRegistryWithFallbackTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackTransferIterator struct {
	Event *ENSRegistryWithFallbackTransfer // Event containing the contract specifics and raw log

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
func (it *ENSRegistryWithFallbackTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSRegistryWithFallbackTransfer)
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
		it.Event = new(ENSRegistryWithFallbackTransfer)
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
func (it *ENSRegistryWithFallbackTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSRegistryWithFallbackTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSRegistryWithFallbackTransfer represents a Transfer event raised by the ENSRegistryWithFallback contract.
type ENSRegistryWithFallbackTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*ENSRegistryWithFallbackTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSRegistryWithFallbackTransferIterator{contract: _ENSRegistryWithFallback.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ENSRegistryWithFallbackTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENSRegistryWithFallback.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSRegistryWithFallbackTransfer)
				if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_ENSRegistryWithFallback *ENSRegistryWithFallbackFilterer) ParseTransfer(log types.Log) (*ENSRegistryWithFallbackTransfer, error) {
	event := new(ENSRegistryWithFallbackTransfer)
	if err := _ENSRegistryWithFallback.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
