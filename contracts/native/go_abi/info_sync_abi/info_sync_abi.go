// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package info_sync_abi

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

// InfoSyncMetaData contains all meta data concerning the InfoSync contract.
var InfoSyncMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"}],\"name\":\"SyncRootInfoEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"}],\"name\":\"getInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"getLatestHeight\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"rootInfos\",\"type\":\"bytes[]\"}],\"name\":\"syncRootInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6a4a9f5e": "getInfo(uint64,uint32)",
		"3cb55d5e": "getLatestHeight(uint64)",
		"06fdde03": "name()",
		"48c8f119": "syncRootInfo(uint64,bytes[])",
	},
}

// InfoSyncABI is the input ABI used to generate the binding from.
// Deprecated: Use InfoSyncMetaData.ABI instead.
var InfoSyncABI = InfoSyncMetaData.ABI

// Deprecated: Use InfoSyncMetaData.Sigs instead.
// InfoSyncFuncSigs maps the 4-byte function signature to its string representation.
var InfoSyncFuncSigs = InfoSyncMetaData.Sigs

// InfoSync is an auto generated Go binding around an Ethereum contract.
type InfoSync struct {
	InfoSyncCaller     // Read-only binding to the contract
	InfoSyncTransactor // Write-only binding to the contract
	InfoSyncFilterer   // Log filterer for contract events
}

// InfoSyncCaller is an auto generated read-only Go binding around an Ethereum contract.
type InfoSyncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoSyncTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InfoSyncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoSyncFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfoSyncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoSyncSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfoSyncSession struct {
	Contract     *InfoSync         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InfoSyncCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfoSyncCallerSession struct {
	Contract *InfoSyncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// InfoSyncTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfoSyncTransactorSession struct {
	Contract     *InfoSyncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InfoSyncRaw is an auto generated low-level Go binding around an Ethereum contract.
type InfoSyncRaw struct {
	Contract *InfoSync // Generic contract binding to access the raw methods on
}

// InfoSyncCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfoSyncCallerRaw struct {
	Contract *InfoSyncCaller // Generic read-only contract binding to access the raw methods on
}

// InfoSyncTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfoSyncTransactorRaw struct {
	Contract *InfoSyncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInfoSync creates a new instance of InfoSync, bound to a specific deployed contract.
func NewInfoSync(address common.Address, backend bind.ContractBackend) (*InfoSync, error) {
	contract, err := bindInfoSync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfoSync{InfoSyncCaller: InfoSyncCaller{contract: contract}, InfoSyncTransactor: InfoSyncTransactor{contract: contract}, InfoSyncFilterer: InfoSyncFilterer{contract: contract}}, nil
}

// NewInfoSyncCaller creates a new read-only instance of InfoSync, bound to a specific deployed contract.
func NewInfoSyncCaller(address common.Address, caller bind.ContractCaller) (*InfoSyncCaller, error) {
	contract, err := bindInfoSync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfoSyncCaller{contract: contract}, nil
}

// NewInfoSyncTransactor creates a new write-only instance of InfoSync, bound to a specific deployed contract.
func NewInfoSyncTransactor(address common.Address, transactor bind.ContractTransactor) (*InfoSyncTransactor, error) {
	contract, err := bindInfoSync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfoSyncTransactor{contract: contract}, nil
}

// NewInfoSyncFilterer creates a new log filterer instance of InfoSync, bound to a specific deployed contract.
func NewInfoSyncFilterer(address common.Address, filterer bind.ContractFilterer) (*InfoSyncFilterer, error) {
	contract, err := bindInfoSync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfoSyncFilterer{contract: contract}, nil
}

// bindInfoSync binds a generic wrapper to an already deployed contract.
func bindInfoSync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InfoSyncABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfoSync *InfoSyncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfoSync.Contract.InfoSyncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfoSync *InfoSyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfoSync.Contract.InfoSyncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfoSync *InfoSyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfoSync.Contract.InfoSyncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfoSync *InfoSyncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfoSync.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfoSync *InfoSyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfoSync.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfoSync *InfoSyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfoSync.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InfoSync *InfoSyncCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InfoSync.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InfoSync *InfoSyncSession) Name() (string, error) {
	return _InfoSync.Contract.Name(&_InfoSync.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InfoSync *InfoSyncCallerSession) Name() (string, error) {
	return _InfoSync.Contract.Name(&_InfoSync.CallOpts)
}

// GetInfo is a paid mutator transaction binding the contract method 0x6a4a9f5e.
//
// Solidity: function getInfo(uint64 chainID, uint32 height) returns(bytes)
func (_InfoSync *InfoSyncTransactor) GetInfo(opts *bind.TransactOpts, chainID uint64, height uint32) (*types.Transaction, error) {
	return _InfoSync.contract.Transact(opts, "getInfo", chainID, height)
}

// GetInfo is a paid mutator transaction binding the contract method 0x6a4a9f5e.
//
// Solidity: function getInfo(uint64 chainID, uint32 height) returns(bytes)
func (_InfoSync *InfoSyncSession) GetInfo(chainID uint64, height uint32) (*types.Transaction, error) {
	return _InfoSync.Contract.GetInfo(&_InfoSync.TransactOpts, chainID, height)
}

// GetInfo is a paid mutator transaction binding the contract method 0x6a4a9f5e.
//
// Solidity: function getInfo(uint64 chainID, uint32 height) returns(bytes)
func (_InfoSync *InfoSyncTransactorSession) GetInfo(chainID uint64, height uint32) (*types.Transaction, error) {
	return _InfoSync.Contract.GetInfo(&_InfoSync.TransactOpts, chainID, height)
}

// GetLatestHeight is a paid mutator transaction binding the contract method 0x3cb55d5e.
//
// Solidity: function getLatestHeight(uint64 chainID) returns(uint32)
func (_InfoSync *InfoSyncTransactor) GetLatestHeight(opts *bind.TransactOpts, chainID uint64) (*types.Transaction, error) {
	return _InfoSync.contract.Transact(opts, "getLatestHeight", chainID)
}

// GetLatestHeight is a paid mutator transaction binding the contract method 0x3cb55d5e.
//
// Solidity: function getLatestHeight(uint64 chainID) returns(uint32)
func (_InfoSync *InfoSyncSession) GetLatestHeight(chainID uint64) (*types.Transaction, error) {
	return _InfoSync.Contract.GetLatestHeight(&_InfoSync.TransactOpts, chainID)
}

// GetLatestHeight is a paid mutator transaction binding the contract method 0x3cb55d5e.
//
// Solidity: function getLatestHeight(uint64 chainID) returns(uint32)
func (_InfoSync *InfoSyncTransactorSession) GetLatestHeight(chainID uint64) (*types.Transaction, error) {
	return _InfoSync.Contract.GetLatestHeight(&_InfoSync.TransactOpts, chainID)
}

// SyncRootInfo is a paid mutator transaction binding the contract method 0x48c8f119.
//
// Solidity: function syncRootInfo(uint64 chainID, bytes[] rootInfos) returns(bool)
func (_InfoSync *InfoSyncTransactor) SyncRootInfo(opts *bind.TransactOpts, chainID uint64, rootInfos [][]byte) (*types.Transaction, error) {
	return _InfoSync.contract.Transact(opts, "syncRootInfo", chainID, rootInfos)
}

// SyncRootInfo is a paid mutator transaction binding the contract method 0x48c8f119.
//
// Solidity: function syncRootInfo(uint64 chainID, bytes[] rootInfos) returns(bool)
func (_InfoSync *InfoSyncSession) SyncRootInfo(chainID uint64, rootInfos [][]byte) (*types.Transaction, error) {
	return _InfoSync.Contract.SyncRootInfo(&_InfoSync.TransactOpts, chainID, rootInfos)
}

// SyncRootInfo is a paid mutator transaction binding the contract method 0x48c8f119.
//
// Solidity: function syncRootInfo(uint64 chainID, bytes[] rootInfos) returns(bool)
func (_InfoSync *InfoSyncTransactorSession) SyncRootInfo(chainID uint64, rootInfos [][]byte) (*types.Transaction, error) {
	return _InfoSync.Contract.SyncRootInfo(&_InfoSync.TransactOpts, chainID, rootInfos)
}

// InfoSyncSyncRootInfoEventIterator is returned from FilterSyncRootInfoEvent and is used to iterate over the raw logs and unpacked data for SyncRootInfoEvent events raised by the InfoSync contract.
type InfoSyncSyncRootInfoEventIterator struct {
	Event *InfoSyncSyncRootInfoEvent // Event containing the contract specifics and raw log

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
func (it *InfoSyncSyncRootInfoEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfoSyncSyncRootInfoEvent)
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
		it.Event = new(InfoSyncSyncRootInfoEvent)
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
func (it *InfoSyncSyncRootInfoEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfoSyncSyncRootInfoEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfoSyncSyncRootInfoEvent represents a SyncRootInfoEvent event raised by the InfoSync contract.
type InfoSyncSyncRootInfoEvent struct {
	ChainID     uint64
	Height      uint32
	BlockHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSyncRootInfoEvent is a free log retrieval operation binding the contract event 0x24cd873c72270d09dcf64a7c56ff70129c04f398e38116d7ad773d43340509e2.
//
// Solidity: event SyncRootInfoEvent(uint64 chainID, uint32 height, uint256 BlockHeight)
func (_InfoSync *InfoSyncFilterer) FilterSyncRootInfoEvent(opts *bind.FilterOpts) (*InfoSyncSyncRootInfoEventIterator, error) {

	logs, sub, err := _InfoSync.contract.FilterLogs(opts, "SyncRootInfoEvent")
	if err != nil {
		return nil, err
	}
	return &InfoSyncSyncRootInfoEventIterator{contract: _InfoSync.contract, event: "SyncRootInfoEvent", logs: logs, sub: sub}, nil
}

// WatchSyncRootInfoEvent is a free log subscription operation binding the contract event 0x24cd873c72270d09dcf64a7c56ff70129c04f398e38116d7ad773d43340509e2.
//
// Solidity: event SyncRootInfoEvent(uint64 chainID, uint32 height, uint256 BlockHeight)
func (_InfoSync *InfoSyncFilterer) WatchSyncRootInfoEvent(opts *bind.WatchOpts, sink chan<- *InfoSyncSyncRootInfoEvent) (event.Subscription, error) {

	logs, sub, err := _InfoSync.contract.WatchLogs(opts, "SyncRootInfoEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfoSyncSyncRootInfoEvent)
				if err := _InfoSync.contract.UnpackLog(event, "SyncRootInfoEvent", log); err != nil {
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

// ParseSyncRootInfoEvent is a log parse operation binding the contract event 0x24cd873c72270d09dcf64a7c56ff70129c04f398e38116d7ad773d43340509e2.
//
// Solidity: event SyncRootInfoEvent(uint64 chainID, uint32 height, uint256 BlockHeight)
func (_InfoSync *InfoSyncFilterer) ParseSyncRootInfoEvent(log types.Log) (*InfoSyncSyncRootInfoEvent, error) {
	event := new(InfoSyncSyncRootInfoEvent)
	if err := _InfoSync.contract.UnpackLog(event, "SyncRootInfoEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

