// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package powerc20factory

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
	_ = abi.ConvertType
)

// Powerc20factoryMetaData contains all meta data concerning the Powerc20factory contract.
var Powerc20factoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"}],\"name\":\"PoWERC20Created\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_miningLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialLimitPerMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"name\":\"createPoWERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalCreatedContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Powerc20factoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Powerc20factoryMetaData.ABI instead.
var Powerc20factoryABI = Powerc20factoryMetaData.ABI

// Powerc20factory is an auto generated Go binding around an Ethereum contract.
type Powerc20factory struct {
	Powerc20factoryCaller     // Read-only binding to the contract
	Powerc20factoryTransactor // Write-only binding to the contract
	Powerc20factoryFilterer   // Log filterer for contract events
}

// Powerc20factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Powerc20factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Powerc20factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Powerc20factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Powerc20factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Powerc20factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Powerc20factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Powerc20factorySession struct {
	Contract     *Powerc20factory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Powerc20factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Powerc20factoryCallerSession struct {
	Contract *Powerc20factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Powerc20factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Powerc20factoryTransactorSession struct {
	Contract     *Powerc20factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Powerc20factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Powerc20factoryRaw struct {
	Contract *Powerc20factory // Generic contract binding to access the raw methods on
}

// Powerc20factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Powerc20factoryCallerRaw struct {
	Contract *Powerc20factoryCaller // Generic read-only contract binding to access the raw methods on
}

// Powerc20factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Powerc20factoryTransactorRaw struct {
	Contract *Powerc20factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPowerc20factory creates a new instance of Powerc20factory, bound to a specific deployed contract.
func NewPowerc20factory(address common.Address, backend bind.ContractBackend) (*Powerc20factory, error) {
	contract, err := bindPowerc20factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Powerc20factory{Powerc20factoryCaller: Powerc20factoryCaller{contract: contract}, Powerc20factoryTransactor: Powerc20factoryTransactor{contract: contract}, Powerc20factoryFilterer: Powerc20factoryFilterer{contract: contract}}, nil
}

// NewPowerc20factoryCaller creates a new read-only instance of Powerc20factory, bound to a specific deployed contract.
func NewPowerc20factoryCaller(address common.Address, caller bind.ContractCaller) (*Powerc20factoryCaller, error) {
	contract, err := bindPowerc20factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Powerc20factoryCaller{contract: contract}, nil
}

// NewPowerc20factoryTransactor creates a new write-only instance of Powerc20factory, bound to a specific deployed contract.
func NewPowerc20factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Powerc20factoryTransactor, error) {
	contract, err := bindPowerc20factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Powerc20factoryTransactor{contract: contract}, nil
}

// NewPowerc20factoryFilterer creates a new log filterer instance of Powerc20factory, bound to a specific deployed contract.
func NewPowerc20factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Powerc20factoryFilterer, error) {
	contract, err := bindPowerc20factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Powerc20factoryFilterer{contract: contract}, nil
}

// bindPowerc20factory binds a generic wrapper to an already deployed contract.
func bindPowerc20factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Powerc20factoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Powerc20factory *Powerc20factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Powerc20factory.Contract.Powerc20factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Powerc20factory *Powerc20factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powerc20factory.Contract.Powerc20factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Powerc20factory *Powerc20factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Powerc20factory.Contract.Powerc20factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Powerc20factory *Powerc20factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Powerc20factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Powerc20factory *Powerc20factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powerc20factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Powerc20factory *Powerc20factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Powerc20factory.Contract.contract.Transact(opts, method, params...)
}

// AllContracts is a free data retrieval call binding the contract method 0xe54e26b9.
//
// Solidity: function allContracts(uint256 ) view returns(address)
func (_Powerc20factory *Powerc20factoryCaller) AllContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Powerc20factory.contract.Call(opts, &out, "allContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllContracts is a free data retrieval call binding the contract method 0xe54e26b9.
//
// Solidity: function allContracts(uint256 ) view returns(address)
func (_Powerc20factory *Powerc20factorySession) AllContracts(arg0 *big.Int) (common.Address, error) {
	return _Powerc20factory.Contract.AllContracts(&_Powerc20factory.CallOpts, arg0)
}

// AllContracts is a free data retrieval call binding the contract method 0xe54e26b9.
//
// Solidity: function allContracts(uint256 ) view returns(address)
func (_Powerc20factory *Powerc20factoryCallerSession) AllContracts(arg0 *big.Int) (common.Address, error) {
	return _Powerc20factory.Contract.AllContracts(&_Powerc20factory.CallOpts, arg0)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns(address[])
func (_Powerc20factory *Powerc20factoryCaller) GetAllContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Powerc20factory.contract.Call(opts, &out, "getAllContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns(address[])
func (_Powerc20factory *Powerc20factorySession) GetAllContracts() ([]common.Address, error) {
	return _Powerc20factory.Contract.GetAllContracts(&_Powerc20factory.CallOpts)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns(address[])
func (_Powerc20factory *Powerc20factoryCallerSession) GetAllContracts() ([]common.Address, error) {
	return _Powerc20factory.Contract.GetAllContracts(&_Powerc20factory.CallOpts)
}

// GetTotalCreatedContracts is a free data retrieval call binding the contract method 0x9646ba64.
//
// Solidity: function getTotalCreatedContracts() view returns(uint256)
func (_Powerc20factory *Powerc20factoryCaller) GetTotalCreatedContracts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20factory.contract.Call(opts, &out, "getTotalCreatedContracts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalCreatedContracts is a free data retrieval call binding the contract method 0x9646ba64.
//
// Solidity: function getTotalCreatedContracts() view returns(uint256)
func (_Powerc20factory *Powerc20factorySession) GetTotalCreatedContracts() (*big.Int, error) {
	return _Powerc20factory.Contract.GetTotalCreatedContracts(&_Powerc20factory.CallOpts)
}

// GetTotalCreatedContracts is a free data retrieval call binding the contract method 0x9646ba64.
//
// Solidity: function getTotalCreatedContracts() view returns(uint256)
func (_Powerc20factory *Powerc20factoryCallerSession) GetTotalCreatedContracts() (*big.Int, error) {
	return _Powerc20factory.Contract.GetTotalCreatedContracts(&_Powerc20factory.CallOpts)
}

// CreatePoWERC20 is a paid mutator transaction binding the contract method 0x6b771618.
//
// Solidity: function createPoWERC20(string name, string symbol, uint256 _initialSupply, uint8 _decimals, uint256 _difficulty, uint256 _miningLimit, uint256 _initialLimitPerMint, address _verifier) returns(address)
func (_Powerc20factory *Powerc20factoryTransactor) CreatePoWERC20(opts *bind.TransactOpts, name string, symbol string, _initialSupply *big.Int, _decimals uint8, _difficulty *big.Int, _miningLimit *big.Int, _initialLimitPerMint *big.Int, _verifier common.Address) (*types.Transaction, error) {
	return _Powerc20factory.contract.Transact(opts, "createPoWERC20", name, symbol, _initialSupply, _decimals, _difficulty, _miningLimit, _initialLimitPerMint, _verifier)
}

// CreatePoWERC20 is a paid mutator transaction binding the contract method 0x6b771618.
//
// Solidity: function createPoWERC20(string name, string symbol, uint256 _initialSupply, uint8 _decimals, uint256 _difficulty, uint256 _miningLimit, uint256 _initialLimitPerMint, address _verifier) returns(address)
func (_Powerc20factory *Powerc20factorySession) CreatePoWERC20(name string, symbol string, _initialSupply *big.Int, _decimals uint8, _difficulty *big.Int, _miningLimit *big.Int, _initialLimitPerMint *big.Int, _verifier common.Address) (*types.Transaction, error) {
	return _Powerc20factory.Contract.CreatePoWERC20(&_Powerc20factory.TransactOpts, name, symbol, _initialSupply, _decimals, _difficulty, _miningLimit, _initialLimitPerMint, _verifier)
}

// CreatePoWERC20 is a paid mutator transaction binding the contract method 0x6b771618.
//
// Solidity: function createPoWERC20(string name, string symbol, uint256 _initialSupply, uint8 _decimals, uint256 _difficulty, uint256 _miningLimit, uint256 _initialLimitPerMint, address _verifier) returns(address)
func (_Powerc20factory *Powerc20factoryTransactorSession) CreatePoWERC20(name string, symbol string, _initialSupply *big.Int, _decimals uint8, _difficulty *big.Int, _miningLimit *big.Int, _initialLimitPerMint *big.Int, _verifier common.Address) (*types.Transaction, error) {
	return _Powerc20factory.Contract.CreatePoWERC20(&_Powerc20factory.TransactOpts, name, symbol, _initialSupply, _decimals, _difficulty, _miningLimit, _initialLimitPerMint, _verifier)
}

// Powerc20factoryPoWERC20CreatedIterator is returned from FilterPoWERC20Created and is used to iterate over the raw logs and unpacked data for PoWERC20Created events raised by the Powerc20factory contract.
type Powerc20factoryPoWERC20CreatedIterator struct {
	Event *Powerc20factoryPoWERC20Created // Event containing the contract specifics and raw log

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
func (it *Powerc20factoryPoWERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Powerc20factoryPoWERC20Created)
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
		it.Event = new(Powerc20factoryPoWERC20Created)
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
func (it *Powerc20factoryPoWERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Powerc20factoryPoWERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Powerc20factoryPoWERC20Created represents a PoWERC20Created event raised by the Powerc20factory contract.
type Powerc20factoryPoWERC20Created struct {
	NewContractAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoWERC20Created is a free log retrieval operation binding the contract event 0xe8d9455984bba86631742a4cfd8a6e4e82aedcc39a116c966eebebc47330e11b.
//
// Solidity: event PoWERC20Created(address newContractAddress)
func (_Powerc20factory *Powerc20factoryFilterer) FilterPoWERC20Created(opts *bind.FilterOpts) (*Powerc20factoryPoWERC20CreatedIterator, error) {

	logs, sub, err := _Powerc20factory.contract.FilterLogs(opts, "PoWERC20Created")
	if err != nil {
		return nil, err
	}
	return &Powerc20factoryPoWERC20CreatedIterator{contract: _Powerc20factory.contract, event: "PoWERC20Created", logs: logs, sub: sub}, nil
}

// WatchPoWERC20Created is a free log subscription operation binding the contract event 0xe8d9455984bba86631742a4cfd8a6e4e82aedcc39a116c966eebebc47330e11b.
//
// Solidity: event PoWERC20Created(address newContractAddress)
func (_Powerc20factory *Powerc20factoryFilterer) WatchPoWERC20Created(opts *bind.WatchOpts, sink chan<- *Powerc20factoryPoWERC20Created) (event.Subscription, error) {

	logs, sub, err := _Powerc20factory.contract.WatchLogs(opts, "PoWERC20Created")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Powerc20factoryPoWERC20Created)
				if err := _Powerc20factory.contract.UnpackLog(event, "PoWERC20Created", log); err != nil {
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

// ParsePoWERC20Created is a log parse operation binding the contract event 0xe8d9455984bba86631742a4cfd8a6e4e82aedcc39a116c966eebebc47330e11b.
//
// Solidity: event PoWERC20Created(address newContractAddress)
func (_Powerc20factory *Powerc20factoryFilterer) ParsePoWERC20Created(log types.Log) (*Powerc20factoryPoWERC20Created, error) {
	event := new(Powerc20factoryPoWERC20Created)
	if err := _Powerc20factory.contract.UnpackLog(event, "PoWERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
