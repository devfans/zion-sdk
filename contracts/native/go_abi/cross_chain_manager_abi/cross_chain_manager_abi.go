// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cross_chain_manager_abi

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

// CrossChainManagerMetaData contains all meta data concerning the CrossChainManager contract.
var CrossChainManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"txHashes\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"ReplenishEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"TxHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"MultiSign\",\"type\":\"bytes\"}],\"name\":\"btcTxMultiSignEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"FromChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"buf\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"FromTxHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"RedeemKey\",\"type\":\"string\"}],\"name\":\"btcTxToRelayEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rk\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"buf\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"amts\",\"type\":\"uint64[]\"}],\"name\":\"makeBtcTxEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"merkleValueHex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"BlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"makeProof\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"}],\"name\":\"BlackChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"RedeemKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"TxHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"Address\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"Signs\",\"type\":\"bytes[]\"}],\"name\":\"MultiSign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"}],\"name\":\"WhiteChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"crossChainID\",\"type\":\"bytes\"}],\"name\":\"checkDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"SourceChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Height\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"Proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Extra\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Signature\",\"type\":\"bytes\"}],\"name\":\"importOuterTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"txHashes\",\"type\":\"string[]\"}],\"name\":\"replenish\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8a449f03": "BlackChain(uint64)",
		"48c79d9d": "MultiSign(uint64,string,bytes,string,bytes[])",
		"99d0e87a": "WhiteChain(uint64)",
		"1245f8d5": "checkDone(uint64,bytes)",
		"bbc2a76a": "importOuterTransfer(uint64,uint32,bytes,bytes,bytes)",
		"06fdde03": "name()",
		"f8bac498": "replenish(uint64,string[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50610540806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638a449f031161005b5780638a449f03146100d357806399d0e87a146100d3578063bbc2a76a146100e6578063f8bac498146100f45761007d565b806306fdde03146100825780631245f8d5146100a057806348c79d9d146100c0575b600080fd5b61008a610107565b6040516100979190610477565b60405180910390f35b6100b36100ae366004610254565b61010c565b604051610097919061046c565b6100b36100ce3660046102a0565b610114565b6100b36100e13660046101b3565b61011f565b6100b36100ce3660046103bc565b6100b36101023660046101d4565b610127565b606090565b600092915050565b600095945050505050565b60005b919050565b60009392505050565b600082601f830112610140578081fd5b813567ffffffffffffffff81111561015a5761015a6104f4565b61016d601f8201601f19166020016104ca565b818152846020838601011115610181578283fd5b816020850160208301379081016020019190915292915050565b803567ffffffffffffffff8116811461012257600080fd5b6000602082840312156101c4578081fd5b6101cd8261019b565b9392505050565b6000806000604084860312156101e8578182fd5b6101f18461019b565b9250602084013567ffffffffffffffff8082111561020d578384fd5b818601915086601f830112610220578384fd5b81358181111561022e578485fd5b8760208083028501011115610241578485fd5b6020830194508093505050509250925092565b60008060408385031215610266578182fd5b61026f8361019b565b9150602083013567ffffffffffffffff81111561028a578182fd5b61029685828601610130565b9150509250929050565b600080600080600060a086880312156102b7578081fd5b6102c08661019b565b945060208087013567ffffffffffffffff808211156102dd578384fd5b6102e98a838b01610130565b965060408901359150808211156102fe578384fd5b61030a8a838b01610130565b9550606089013591508082111561031f578384fd5b61032b8a838b01610130565b94506080890135915080821115610340578384fd5b818901915089601f830112610353578384fd5b813581811115610365576103656104f4565b61037284858302016104ca565b8181528481019250838501865b838110156103a8576103968e888435890101610130565b8552938601939086019060010161037f565b505080955050505050509295509295909350565b600080600080600060a086880312156103d3578081fd5b6103dc8661019b565b9450602086013563ffffffff811681146103f4578182fd5b9350604086013567ffffffffffffffff80821115610410578283fd5b61041c89838a01610130565b94506060880135915080821115610431578283fd5b61043d89838a01610130565b93506080880135915080821115610452578283fd5b5061045f88828901610130565b9150509295509295909350565b901515815260200190565b6000602080835283518082850152825b818110156104a357858101830151858201604001528201610487565b818111156104b45783604083870101525b50601f01601f1916929092016040019392505050565b60405181810167ffffffffffffffff811182821017156104ec576104ec6104f4565b604052919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220d6be45f893ec188a13c5dfa7d370d6f46c369e4b406aaae431e8c21ef12c6a3064736f6c63430008000033",
}

// CrossChainManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossChainManagerMetaData.ABI instead.
var CrossChainManagerABI = CrossChainManagerMetaData.ABI

// Deprecated: Use CrossChainManagerMetaData.Sigs instead.
// CrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var CrossChainManagerFuncSigs = CrossChainManagerMetaData.Sigs

// CrossChainManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossChainManagerMetaData.Bin instead.
var CrossChainManagerBin = CrossChainManagerMetaData.Bin

// DeployCrossChainManager deploys a new Ethereum contract, binding an instance of CrossChainManager to it.
func DeployCrossChainManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossChainManager, error) {
	parsed, err := CrossChainManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossChainManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossChainManager{CrossChainManagerCaller: CrossChainManagerCaller{contract: contract}, CrossChainManagerTransactor: CrossChainManagerTransactor{contract: contract}, CrossChainManagerFilterer: CrossChainManagerFilterer{contract: contract}}, nil
}

// CrossChainManager is an auto generated Go binding around an Ethereum contract.
type CrossChainManager struct {
	CrossChainManagerCaller     // Read-only binding to the contract
	CrossChainManagerTransactor // Write-only binding to the contract
	CrossChainManagerFilterer   // Log filterer for contract events
}

// CrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainManagerSession struct {
	Contract     *CrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainManagerCallerSession struct {
	Contract *CrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainManagerTransactorSession struct {
	Contract     *CrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainManagerRaw struct {
	Contract *CrossChainManager // Generic contract binding to access the raw methods on
}

// CrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainManagerCallerRaw struct {
	Contract *CrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainManagerTransactorRaw struct {
	Contract *CrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChainManager creates a new instance of CrossChainManager, bound to a specific deployed contract.
func NewCrossChainManager(address common.Address, backend bind.ContractBackend) (*CrossChainManager, error) {
	contract, err := bindCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainManager{CrossChainManagerCaller: CrossChainManagerCaller{contract: contract}, CrossChainManagerTransactor: CrossChainManagerTransactor{contract: contract}, CrossChainManagerFilterer: CrossChainManagerFilterer{contract: contract}}, nil
}

// NewCrossChainManagerCaller creates a new read-only instance of CrossChainManager, bound to a specific deployed contract.
func NewCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*CrossChainManagerCaller, error) {
	contract, err := bindCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerCaller{contract: contract}, nil
}

// NewCrossChainManagerTransactor creates a new write-only instance of CrossChainManager, bound to a specific deployed contract.
func NewCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainManagerTransactor, error) {
	contract, err := bindCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerTransactor{contract: contract}, nil
}

// NewCrossChainManagerFilterer creates a new log filterer instance of CrossChainManager, bound to a specific deployed contract.
func NewCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainManagerFilterer, error) {
	contract, err := bindCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerFilterer{contract: contract}, nil
}

// bindCrossChainManager binds a generic wrapper to an already deployed contract.
func bindCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainManager *CrossChainManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainManager.Contract.CrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainManager *CrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainManager.Contract.CrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainManager *CrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainManager.Contract.CrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainManager *CrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainManager *CrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainManager *CrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_CrossChainManager *CrossChainManagerCaller) CheckDone(opts *bind.CallOpts, chainID uint64, crossChainID []byte) (bool, error) {
	var out []interface{}
	err := _CrossChainManager.contract.Call(opts, &out, "checkDone", chainID, crossChainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_CrossChainManager *CrossChainManagerSession) CheckDone(chainID uint64, crossChainID []byte) (bool, error) {
	return _CrossChainManager.Contract.CheckDone(&_CrossChainManager.CallOpts, chainID, crossChainID)
}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_CrossChainManager *CrossChainManagerCallerSession) CheckDone(chainID uint64, crossChainID []byte) (bool, error) {
	return _CrossChainManager.Contract.CheckDone(&_CrossChainManager.CallOpts, chainID, crossChainID)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactor) BlackChain(opts *bind.TransactOpts, ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManager.contract.Transact(opts, "BlackChain", ChainID)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_CrossChainManager *CrossChainManagerSession) BlackChain(ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManager.Contract.BlackChain(&_CrossChainManager.TransactOpts, ChainID)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactorSession) BlackChain(ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManager.Contract.BlackChain(&_CrossChainManager.TransactOpts, ChainID)
}

// MultiSign is a paid mutator transaction binding the contract method 0x48c79d9d.
//
// Solidity: function MultiSign(uint64 ChainID, string RedeemKey, bytes TxHash, string Address, bytes[] Signs) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactor) MultiSign(opts *bind.TransactOpts, ChainID uint64, RedeemKey string, TxHash []byte, Address string, Signs [][]byte) (*types.Transaction, error) {
	return _CrossChainManager.contract.Transact(opts, "MultiSign", ChainID, RedeemKey, TxHash, Address, Signs)
}

// MultiSign is a paid mutator transaction binding the contract method 0x48c79d9d.
//
// Solidity: function MultiSign(uint64 ChainID, string RedeemKey, bytes TxHash, string Address, bytes[] Signs) returns(bool success)
func (_CrossChainManager *CrossChainManagerSession) MultiSign(ChainID uint64, RedeemKey string, TxHash []byte, Address string, Signs [][]byte) (*types.Transaction, error) {
	return _CrossChainManager.Contract.MultiSign(&_CrossChainManager.TransactOpts, ChainID, RedeemKey, TxHash, Address, Signs)
}

// MultiSign is a paid mutator transaction binding the contract method 0x48c79d9d.
//
// Solidity: function MultiSign(uint64 ChainID, string RedeemKey, bytes TxHash, string Address, bytes[] Signs) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactorSession) MultiSign(ChainID uint64, RedeemKey string, TxHash []byte, Address string, Signs [][]byte) (*types.Transaction, error) {
	return _CrossChainManager.Contract.MultiSign(&_CrossChainManager.TransactOpts, ChainID, RedeemKey, TxHash, Address, Signs)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactor) WhiteChain(opts *bind.TransactOpts, ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManager.contract.Transact(opts, "WhiteChain", ChainID)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_CrossChainManager *CrossChainManagerSession) WhiteChain(ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManager.Contract.WhiteChain(&_CrossChainManager.TransactOpts, ChainID)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactorSession) WhiteChain(ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManager.Contract.WhiteChain(&_CrossChainManager.TransactOpts, ChainID)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactor) ImportOuterTransfer(opts *bind.TransactOpts, SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _CrossChainManager.contract.Transact(opts, "importOuterTransfer", SourceChainID, Height, Proof, Extra, Signature)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_CrossChainManager *CrossChainManagerSession) ImportOuterTransfer(SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _CrossChainManager.Contract.ImportOuterTransfer(&_CrossChainManager.TransactOpts, SourceChainID, Height, Proof, Extra, Signature)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactorSession) ImportOuterTransfer(SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _CrossChainManager.Contract.ImportOuterTransfer(&_CrossChainManager.TransactOpts, SourceChainID, Height, Proof, Extra, Signature)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string Name)
func (_CrossChainManager *CrossChainManagerTransactor) Name(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainManager.contract.Transact(opts, "name")
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string Name)
func (_CrossChainManager *CrossChainManagerSession) Name() (*types.Transaction, error) {
	return _CrossChainManager.Contract.Name(&_CrossChainManager.TransactOpts)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string Name)
func (_CrossChainManager *CrossChainManagerTransactorSession) Name() (*types.Transaction, error) {
	return _CrossChainManager.Contract.Name(&_CrossChainManager.TransactOpts)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactor) Replenish(opts *bind.TransactOpts, chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _CrossChainManager.contract.Transact(opts, "replenish", chainID, txHashes)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_CrossChainManager *CrossChainManagerSession) Replenish(chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _CrossChainManager.Contract.Replenish(&_CrossChainManager.TransactOpts, chainID, txHashes)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_CrossChainManager *CrossChainManagerTransactorSession) Replenish(chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _CrossChainManager.Contract.Replenish(&_CrossChainManager.TransactOpts, chainID, txHashes)
}

// CrossChainManagerReplenishEventIterator is returned from FilterReplenishEvent and is used to iterate over the raw logs and unpacked data for ReplenishEvent events raised by the CrossChainManager contract.
type CrossChainManagerReplenishEventIterator struct {
	Event *CrossChainManagerReplenishEvent // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerReplenishEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerReplenishEvent)
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
		it.Event = new(CrossChainManagerReplenishEvent)
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
func (it *CrossChainManagerReplenishEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerReplenishEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerReplenishEvent represents a ReplenishEvent event raised by the CrossChainManager contract.
type CrossChainManagerReplenishEvent struct {
	TxHashes []string
	ChainID  uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReplenishEvent is a free log retrieval operation binding the contract event 0xac3e52c0a7de47fbd0f9a52b8f205485cd725235d94d678f638e16d02404fb38.
//
// Solidity: event ReplenishEvent(string[] txHashes, uint64 chainID)
func (_CrossChainManager *CrossChainManagerFilterer) FilterReplenishEvent(opts *bind.FilterOpts) (*CrossChainManagerReplenishEventIterator, error) {

	logs, sub, err := _CrossChainManager.contract.FilterLogs(opts, "ReplenishEvent")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerReplenishEventIterator{contract: _CrossChainManager.contract, event: "ReplenishEvent", logs: logs, sub: sub}, nil
}

// WatchReplenishEvent is a free log subscription operation binding the contract event 0xac3e52c0a7de47fbd0f9a52b8f205485cd725235d94d678f638e16d02404fb38.
//
// Solidity: event ReplenishEvent(string[] txHashes, uint64 chainID)
func (_CrossChainManager *CrossChainManagerFilterer) WatchReplenishEvent(opts *bind.WatchOpts, sink chan<- *CrossChainManagerReplenishEvent) (event.Subscription, error) {

	logs, sub, err := _CrossChainManager.contract.WatchLogs(opts, "ReplenishEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerReplenishEvent)
				if err := _CrossChainManager.contract.UnpackLog(event, "ReplenishEvent", log); err != nil {
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

// ParseReplenishEvent is a log parse operation binding the contract event 0xac3e52c0a7de47fbd0f9a52b8f205485cd725235d94d678f638e16d02404fb38.
//
// Solidity: event ReplenishEvent(string[] txHashes, uint64 chainID)
func (_CrossChainManager *CrossChainManagerFilterer) ParseReplenishEvent(log types.Log) (*CrossChainManagerReplenishEvent, error) {
	event := new(CrossChainManagerReplenishEvent)
	if err := _CrossChainManager.contract.UnpackLog(event, "ReplenishEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainManagerBtcTxMultiSignEventIterator is returned from FilterBtcTxMultiSignEvent and is used to iterate over the raw logs and unpacked data for BtcTxMultiSignEvent events raised by the CrossChainManager contract.
type CrossChainManagerBtcTxMultiSignEventIterator struct {
	Event *CrossChainManagerBtcTxMultiSignEvent // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerBtcTxMultiSignEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerBtcTxMultiSignEvent)
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
		it.Event = new(CrossChainManagerBtcTxMultiSignEvent)
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
func (it *CrossChainManagerBtcTxMultiSignEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerBtcTxMultiSignEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerBtcTxMultiSignEvent represents a BtcTxMultiSignEvent event raised by the CrossChainManager contract.
type CrossChainManagerBtcTxMultiSignEvent struct {
	TxHash    []byte
	MultiSign []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBtcTxMultiSignEvent is a free log retrieval operation binding the contract event 0x62fb550ff7fa48f759b0e56ea24757e77b5612d11efbcbdbed9545982cfe1770.
//
// Solidity: event btcTxMultiSignEvent(bytes TxHash, bytes MultiSign)
func (_CrossChainManager *CrossChainManagerFilterer) FilterBtcTxMultiSignEvent(opts *bind.FilterOpts) (*CrossChainManagerBtcTxMultiSignEventIterator, error) {

	logs, sub, err := _CrossChainManager.contract.FilterLogs(opts, "btcTxMultiSignEvent")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerBtcTxMultiSignEventIterator{contract: _CrossChainManager.contract, event: "btcTxMultiSignEvent", logs: logs, sub: sub}, nil
}

// WatchBtcTxMultiSignEvent is a free log subscription operation binding the contract event 0x62fb550ff7fa48f759b0e56ea24757e77b5612d11efbcbdbed9545982cfe1770.
//
// Solidity: event btcTxMultiSignEvent(bytes TxHash, bytes MultiSign)
func (_CrossChainManager *CrossChainManagerFilterer) WatchBtcTxMultiSignEvent(opts *bind.WatchOpts, sink chan<- *CrossChainManagerBtcTxMultiSignEvent) (event.Subscription, error) {

	logs, sub, err := _CrossChainManager.contract.WatchLogs(opts, "btcTxMultiSignEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerBtcTxMultiSignEvent)
				if err := _CrossChainManager.contract.UnpackLog(event, "btcTxMultiSignEvent", log); err != nil {
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

// ParseBtcTxMultiSignEvent is a log parse operation binding the contract event 0x62fb550ff7fa48f759b0e56ea24757e77b5612d11efbcbdbed9545982cfe1770.
//
// Solidity: event btcTxMultiSignEvent(bytes TxHash, bytes MultiSign)
func (_CrossChainManager *CrossChainManagerFilterer) ParseBtcTxMultiSignEvent(log types.Log) (*CrossChainManagerBtcTxMultiSignEvent, error) {
	event := new(CrossChainManagerBtcTxMultiSignEvent)
	if err := _CrossChainManager.contract.UnpackLog(event, "btcTxMultiSignEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainManagerBtcTxToRelayEventIterator is returned from FilterBtcTxToRelayEvent and is used to iterate over the raw logs and unpacked data for BtcTxToRelayEvent events raised by the CrossChainManager contract.
type CrossChainManagerBtcTxToRelayEventIterator struct {
	Event *CrossChainManagerBtcTxToRelayEvent // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerBtcTxToRelayEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerBtcTxToRelayEvent)
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
		it.Event = new(CrossChainManagerBtcTxToRelayEvent)
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
func (it *CrossChainManagerBtcTxToRelayEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerBtcTxToRelayEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerBtcTxToRelayEvent represents a BtcTxToRelayEvent event raised by the CrossChainManager contract.
type CrossChainManagerBtcTxToRelayEvent struct {
	FromChainID uint64
	ChainID     uint64
	Buf         string
	FromTxHash  string
	RedeemKey   string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBtcTxToRelayEvent is a free log retrieval operation binding the contract event 0x59c070ab5215dda625f463061a0ad421505b2cca1066a8597411c72a5ecac51b.
//
// Solidity: event btcTxToRelayEvent(uint64 FromChainID, uint64 ChainID, string buf, string FromTxHash, string RedeemKey)
func (_CrossChainManager *CrossChainManagerFilterer) FilterBtcTxToRelayEvent(opts *bind.FilterOpts) (*CrossChainManagerBtcTxToRelayEventIterator, error) {

	logs, sub, err := _CrossChainManager.contract.FilterLogs(opts, "btcTxToRelayEvent")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerBtcTxToRelayEventIterator{contract: _CrossChainManager.contract, event: "btcTxToRelayEvent", logs: logs, sub: sub}, nil
}

// WatchBtcTxToRelayEvent is a free log subscription operation binding the contract event 0x59c070ab5215dda625f463061a0ad421505b2cca1066a8597411c72a5ecac51b.
//
// Solidity: event btcTxToRelayEvent(uint64 FromChainID, uint64 ChainID, string buf, string FromTxHash, string RedeemKey)
func (_CrossChainManager *CrossChainManagerFilterer) WatchBtcTxToRelayEvent(opts *bind.WatchOpts, sink chan<- *CrossChainManagerBtcTxToRelayEvent) (event.Subscription, error) {

	logs, sub, err := _CrossChainManager.contract.WatchLogs(opts, "btcTxToRelayEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerBtcTxToRelayEvent)
				if err := _CrossChainManager.contract.UnpackLog(event, "btcTxToRelayEvent", log); err != nil {
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

// ParseBtcTxToRelayEvent is a log parse operation binding the contract event 0x59c070ab5215dda625f463061a0ad421505b2cca1066a8597411c72a5ecac51b.
//
// Solidity: event btcTxToRelayEvent(uint64 FromChainID, uint64 ChainID, string buf, string FromTxHash, string RedeemKey)
func (_CrossChainManager *CrossChainManagerFilterer) ParseBtcTxToRelayEvent(log types.Log) (*CrossChainManagerBtcTxToRelayEvent, error) {
	event := new(CrossChainManagerBtcTxToRelayEvent)
	if err := _CrossChainManager.contract.UnpackLog(event, "btcTxToRelayEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainManagerMakeBtcTxEventIterator is returned from FilterMakeBtcTxEvent and is used to iterate over the raw logs and unpacked data for MakeBtcTxEvent events raised by the CrossChainManager contract.
type CrossChainManagerMakeBtcTxEventIterator struct {
	Event *CrossChainManagerMakeBtcTxEvent // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerMakeBtcTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerMakeBtcTxEvent)
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
		it.Event = new(CrossChainManagerMakeBtcTxEvent)
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
func (it *CrossChainManagerMakeBtcTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerMakeBtcTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerMakeBtcTxEvent represents a MakeBtcTxEvent event raised by the CrossChainManager contract.
type CrossChainManagerMakeBtcTxEvent struct {
	Rk   string
	Buf  string
	Amts []uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMakeBtcTxEvent is a free log retrieval operation binding the contract event 0xa00d721fd040a2b479d1adf886244de04bdbb3e3e310dc75f1036e2602726234.
//
// Solidity: event makeBtcTxEvent(string rk, string buf, uint64[] amts)
func (_CrossChainManager *CrossChainManagerFilterer) FilterMakeBtcTxEvent(opts *bind.FilterOpts) (*CrossChainManagerMakeBtcTxEventIterator, error) {

	logs, sub, err := _CrossChainManager.contract.FilterLogs(opts, "makeBtcTxEvent")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerMakeBtcTxEventIterator{contract: _CrossChainManager.contract, event: "makeBtcTxEvent", logs: logs, sub: sub}, nil
}

// WatchMakeBtcTxEvent is a free log subscription operation binding the contract event 0xa00d721fd040a2b479d1adf886244de04bdbb3e3e310dc75f1036e2602726234.
//
// Solidity: event makeBtcTxEvent(string rk, string buf, uint64[] amts)
func (_CrossChainManager *CrossChainManagerFilterer) WatchMakeBtcTxEvent(opts *bind.WatchOpts, sink chan<- *CrossChainManagerMakeBtcTxEvent) (event.Subscription, error) {

	logs, sub, err := _CrossChainManager.contract.WatchLogs(opts, "makeBtcTxEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerMakeBtcTxEvent)
				if err := _CrossChainManager.contract.UnpackLog(event, "makeBtcTxEvent", log); err != nil {
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

// ParseMakeBtcTxEvent is a log parse operation binding the contract event 0xa00d721fd040a2b479d1adf886244de04bdbb3e3e310dc75f1036e2602726234.
//
// Solidity: event makeBtcTxEvent(string rk, string buf, uint64[] amts)
func (_CrossChainManager *CrossChainManagerFilterer) ParseMakeBtcTxEvent(log types.Log) (*CrossChainManagerMakeBtcTxEvent, error) {
	event := new(CrossChainManagerMakeBtcTxEvent)
	if err := _CrossChainManager.contract.UnpackLog(event, "makeBtcTxEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainManagerMakeProofIterator is returned from FilterMakeProof and is used to iterate over the raw logs and unpacked data for MakeProof events raised by the CrossChainManager contract.
type CrossChainManagerMakeProofIterator struct {
	Event *CrossChainManagerMakeProof // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerMakeProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerMakeProof)
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
		it.Event = new(CrossChainManagerMakeProof)
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
func (it *CrossChainManagerMakeProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerMakeProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerMakeProof represents a MakeProof event raised by the CrossChainManager contract.
type CrossChainManagerMakeProof struct {
	MerkleValueHex string
	BlockHeight    uint64
	Key            string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMakeProof is a free log retrieval operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManager *CrossChainManagerFilterer) FilterMakeProof(opts *bind.FilterOpts) (*CrossChainManagerMakeProofIterator, error) {

	logs, sub, err := _CrossChainManager.contract.FilterLogs(opts, "makeProof")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerMakeProofIterator{contract: _CrossChainManager.contract, event: "makeProof", logs: logs, sub: sub}, nil
}

// WatchMakeProof is a free log subscription operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManager *CrossChainManagerFilterer) WatchMakeProof(opts *bind.WatchOpts, sink chan<- *CrossChainManagerMakeProof) (event.Subscription, error) {

	logs, sub, err := _CrossChainManager.contract.WatchLogs(opts, "makeProof")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerMakeProof)
				if err := _CrossChainManager.contract.UnpackLog(event, "makeProof", log); err != nil {
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

// ParseMakeProof is a log parse operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManager *CrossChainManagerFilterer) ParseMakeProof(log types.Log) (*CrossChainManagerMakeProof, error) {
	event := new(CrossChainManagerMakeProof)
	if err := _CrossChainManager.contract.UnpackLog(event, "makeProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

