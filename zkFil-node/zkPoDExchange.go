// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// zkFilExchangeABI is the input ABI used to generate the binding from.
const zkFilExchangeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"publicVar_\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"bobDeposits_\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"unDepositAt\",\"type\":\"uint256\"},{\"name\":\"pendingCnt\",\"type\":\"uint256\"},{\"name\":\"stat\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"s_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t3\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bulletins_\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"size\",\"type\":\"uint64\"},{\"name\":\"s\",\"type\":\"uint64\"},{\"name\":\"n\",\"type\":\"uint64\"},{\"name\":\"sigma_mkl_root\",\"type\":\"uint256\"},{\"name\":\"vrf_meta_digest\",\"type\":\"uint256\"},{\"name\":\"pledge_value\",\"type\":\"uint256\"},{\"name\":\"unDepositAt\",\"type\":\"uint256\"},{\"name\":\"pendingCnt\",\"type\":\"uint256\"},{\"name\":\"blt_type\",\"type\":\"uint8\"},{\"name\":\"stat\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t1\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t4\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_publicVar\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_mode\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"OnDeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_seed0\",\"type\":\"bytes32\"}],\"name\":\"OnComplaintKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"OnComplaintClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"OnComplaintDeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_seed0\",\"type\":\"bytes32\"}],\"name\":\"OnAtomicSwapDeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_seed0\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_seed0_rand\",\"type\":\"uint256\"}],\"name\":\"OnAtomicSwapVCDeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_r\",\"type\":\"uint256\"}],\"name\":\"OnVRFDeal\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_size\",\"type\":\"uint64\"},{\"name\":\"_s\",\"type\":\"uint64\"},{\"name\":\"_n\",\"type\":\"uint64\"},{\"name\":\"_sigma_mkl_root\",\"type\":\"uint256\"},{\"name\":\"_vrf_meta_digest\",\"type\":\"uint256\"},{\"name\":\"_blt_type\",\"type\":\"uint256\"}],\"name\":\"publish\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bltKey\",\"type\":\"bytes32\"}],\"name\":\"unPublish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"bobDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"bobUnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bltKey\",\"type\":\"bytes32\"}],\"name\":\"withdrawA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"withdrawB\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_seed0\",\"type\":\"bytes32\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_addrs\",\"type\":\"address[2]\"},{\"name\":\"_bltKey\",\"type\":\"bytes32\"},{\"name\":\"_seed2\",\"type\":\"bytes32\"},{\"name\":\"_k_mkl_root\",\"type\":\"bytes32\"},{\"name\":\"_count\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"submitProofComplaint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_i\",\"type\":\"uint64\"},{\"name\":\"_j\",\"type\":\"uint64\"},{\"name\":\"_tx\",\"type\":\"uint256\"},{\"name\":\"_ty\",\"type\":\"uint256\"},{\"name\":\"_mkl_path\",\"type\":\"bytes32[]\"},{\"name\":\"_sCnt\",\"type\":\"uint64\"}],\"name\":\"claimComplaint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"settleComplaintDeal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_seed0\",\"type\":\"bytes32\"},{\"name\":\"_sCnt\",\"type\":\"uint64\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_addrs\",\"type\":\"address[2]\"},{\"name\":\"_seed2\",\"type\":\"bytes32\"},{\"name\":\"_sigma_vw\",\"type\":\"uint256\"},{\"name\":\"_count\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"submitProofAtomicSwap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_seed0\",\"type\":\"uint256\"},{\"name\":\"_seed0_rand\",\"type\":\"uint256\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_addrs\",\"type\":\"address[2]\"},{\"name\":\"_seed0_mimc3_digest\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"submitProofAtomicSwapVC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_s_r\",\"type\":\"uint256\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_addrs\",\"type\":\"address[2]\"},{\"name\":\"_g_exp_r\",\"type\":\"uint256[2]\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"submitProofVRF\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getSessionRecord\",\"outputs\":[{\"name\":\"submitAt\",\"type\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint8\"},{\"name\":\"stat\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getRecordComplaint\",\"outputs\":[{\"name\":\"seed0\",\"type\":\"bytes32\"},{\"name\":\"seed2\",\"type\":\"bytes32\"},{\"name\":\"k_mkl_root\",\"type\":\"bytes32\"},{\"name\":\"count\",\"type\":\"uint64\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"expireAt\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getRecordAtomicSwap\",\"outputs\":[{\"name\":\"seed0\",\"type\":\"bytes32\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getRecordAtomicSwapVC\",\"outputs\":[{\"name\":\"seed0\",\"type\":\"uint256\"},{\"name\":\"seed0_rand\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getRecordVRF\",\"outputs\":[{\"name\":\"r\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// zkFilExchange is an auto generated Go binding around an Ethereum contract.
type zkFilExchange struct {
	zkFilExchangeCaller     // Read-only binding to the contract
	zkFilExchangeTransactor // Write-only binding to the contract
	zkFilExchangeFilterer   // Log filterer for contract events
}

// zkFilExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type zkFilExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// zkFilExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type zkFilExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// zkFilExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type zkFilExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// zkFilExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type zkFilExchangeSession struct {
	Contract     *zkFilExchange    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// zkFilExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type zkFilExchangeCallerSession struct {
	Contract *zkFilExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// zkFilExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type zkFilExchangeTransactorSession struct {
	Contract     *zkFilExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// zkFilExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type zkFilExchangeRaw struct {
	Contract *zkFilExchange // Generic contract binding to access the raw methods on
}

// zkFilExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type zkFilExchangeCallerRaw struct {
	Contract *zkFilExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// zkFilExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type zkFilExchangeTransactorRaw struct {
	Contract *zkFilExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewzkFilExchange creates a new instance of zkFilExchange, bound to a specific deployed contract.
func NewzkFilExchange(address common.Address, backend bind.ContractBackend) (*zkFilExchange, error) {
	contract, err := bindzkFilExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &zkFilExchange{zkFilExchangeCaller: zkFilExchangeCaller{contract: contract}, zkFilExchangeTransactor: zkFilExchangeTransactor{contract: contract}, zkFilExchangeFilterer: zkFilExchangeFilterer{contract: contract}}, nil
}

// NewzkFilExchangeCaller creates a new read-only instance of zkFilExchange, bound to a specific deployed contract.
func NewzkFilExchangeCaller(address common.Address, caller bind.ContractCaller) (*zkFilExchangeCaller, error) {
	contract, err := bindzkFilExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeCaller{contract: contract}, nil
}

// NewzkFilExchangeTransactor creates a new write-only instance of zkFilExchange, bound to a specific deployed contract.
func NewzkFilExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*zkFilExchangeTransactor, error) {
	contract, err := bindzkFilExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeTransactor{contract: contract}, nil
}

// NewzkFilExchangeFilterer creates a new log filterer instance of zkFilExchange, bound to a specific deployed contract.
func NewzkFilExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*zkFilExchangeFilterer, error) {
	contract, err := bindzkFilExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeFilterer{contract: contract}, nil
}

// bindzkFilExchange binds a generic wrapper to an already deployed contract.
func bindzkFilExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(zkFilExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_zkFilExchange *zkFilExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _zkFilExchange.Contract.zkFilExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_zkFilExchange *zkFilExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _zkFilExchange.Contract.zkFilExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_zkFilExchange *zkFilExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _zkFilExchange.Contract.zkFilExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_zkFilExchange *zkFilExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _zkFilExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_zkFilExchange *zkFilExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _zkFilExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_zkFilExchange *zkFilExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _zkFilExchange.Contract.contract.Transact(opts, method, params...)
}

// BobDeposits is a free data retrieval call binding the contract method 0x98033565.
//
// Solidity: function bobDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint256 pendingCnt, uint8 stat)
func (_zkFilExchange *zkFilExchangeCaller) BobDeposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	PendingCnt  *big.Int
	Stat        uint8
}, error) {
	ret := new(struct {
		Value       *big.Int
		UnDepositAt *big.Int
		PendingCnt  *big.Int
		Stat        uint8
	})
	out := ret
	err := _zkFilExchange.contract.Call(opts, out, "bobDeposits_", arg0, arg1)
	return *ret, err
}

// BobDeposits is a free data retrieval call binding the contract method 0x98033565.
//
// Solidity: function bobDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint256 pendingCnt, uint8 stat)
func (_zkFilExchange *zkFilExchangeSession) BobDeposits(arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	PendingCnt  *big.Int
	Stat        uint8
}, error) {
	return _zkFilExchange.Contract.BobDeposits(&_zkFilExchange.CallOpts, arg0, arg1)
}

// BobDeposits is a free data retrieval call binding the contract method 0x98033565.
//
// Solidity: function bobDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint256 pendingCnt, uint8 stat)
func (_zkFilExchange *zkFilExchangeCallerSession) BobDeposits(arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	PendingCnt  *big.Int
	Stat        uint8
}, error) {
	return _zkFilExchange.Contract.BobDeposits(&_zkFilExchange.CallOpts, arg0, arg1)
}

// Bulletins is a free data retrieval call binding the contract method 0xd9417785.
//
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint256 pendingCnt, uint8 blt_type, uint8 stat)
func (_zkFilExchange *zkFilExchangeCaller) Bulletins(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner         common.Address
	Size          uint64
	S             uint64
	N             uint64
	SigmaMklRoot  *big.Int
	VrfMetaDigest *big.Int
	PledgeValue   *big.Int
	UnDepositAt   *big.Int
	PendingCnt    *big.Int
	BltType       uint8
	Stat          uint8
}, error) {
	ret := new(struct {
		Owner         common.Address
		Size          uint64
		S             uint64
		N             uint64
		SigmaMklRoot  *big.Int
		VrfMetaDigest *big.Int
		PledgeValue   *big.Int
		UnDepositAt   *big.Int
		PendingCnt    *big.Int
		BltType       uint8
		Stat          uint8
	})
	out := ret
	err := _zkFilExchange.contract.Call(opts, out, "bulletins_", arg0)
	return *ret, err
}

// Bulletins is a free data retrieval call binding the contract method 0xd9417785.
//
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint256 pendingCnt, uint8 blt_type, uint8 stat)
func (_zkFilExchange *zkFilExchangeSession) Bulletins(arg0 [32]byte) (struct {
	Owner         common.Address
	Size          uint64
	S             uint64
	N             uint64
	SigmaMklRoot  *big.Int
	VrfMetaDigest *big.Int
	PledgeValue   *big.Int
	UnDepositAt   *big.Int
	PendingCnt    *big.Int
	BltType       uint8
	Stat          uint8
}, error) {
	return _zkFilExchange.Contract.Bulletins(&_zkFilExchange.CallOpts, arg0)
}

// Bulletins is a free data retrieval call binding the contract method 0xd9417785.
//
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint256 pendingCnt, uint8 blt_type, uint8 stat)
func (_zkFilExchange *zkFilExchangeCallerSession) Bulletins(arg0 [32]byte) (struct {
	Owner         common.Address
	Size          uint64
	S             uint64
	N             uint64
	SigmaMklRoot  *big.Int
	VrfMetaDigest *big.Int
	PledgeValue   *big.Int
	UnDepositAt   *big.Int
	PendingCnt    *big.Int
	BltType       uint8
	Stat          uint8
}, error) {
	return _zkFilExchange.Contract.Bulletins(&_zkFilExchange.CallOpts, arg0)
}

// GetRecordAtomicSwap is a free data retrieval call binding the contract method 0x7c7c75a8.
//
// Solidity: function getRecordAtomicSwap(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeCaller) GetRecordAtomicSwap(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		Seed0    [32]byte
		SubmitAt *big.Int
	})
	out := ret
	err := _zkFilExchange.contract.Call(opts, out, "getRecordAtomicSwap", _a, _b, _sid)
	return *ret, err
}

// GetRecordAtomicSwap is a free data retrieval call binding the contract method 0x7c7c75a8.
//
// Solidity: function getRecordAtomicSwap(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeSession) GetRecordAtomicSwap(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	return _zkFilExchange.Contract.GetRecordAtomicSwap(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetRecordAtomicSwap is a free data retrieval call binding the contract method 0x7c7c75a8.
//
// Solidity: function getRecordAtomicSwap(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeCallerSession) GetRecordAtomicSwap(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	return _zkFilExchange.Contract.GetRecordAtomicSwap(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetRecordAtomicSwapVC is a free data retrieval call binding the contract method 0xa7e7b767.
//
// Solidity: function getRecordAtomicSwapVC(address _a, address _b, uint256 _sid) constant returns(uint256 seed0, uint256 seed0_rand, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeCaller) GetRecordAtomicSwapVC(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0     *big.Int
	Seed0Rand *big.Int
	SubmitAt  *big.Int
}, error) {
	ret := new(struct {
		Seed0     *big.Int
		Seed0Rand *big.Int
		SubmitAt  *big.Int
	})
	out := ret
	err := _zkFilExchange.contract.Call(opts, out, "getRecordAtomicSwapVC", _a, _b, _sid)
	return *ret, err
}

// GetRecordAtomicSwapVC is a free data retrieval call binding the contract method 0xa7e7b767.
//
// Solidity: function getRecordAtomicSwapVC(address _a, address _b, uint256 _sid) constant returns(uint256 seed0, uint256 seed0_rand, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeSession) GetRecordAtomicSwapVC(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0     *big.Int
	Seed0Rand *big.Int
	SubmitAt  *big.Int
}, error) {
	return _zkFilExchange.Contract.GetRecordAtomicSwapVC(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetRecordAtomicSwapVC is a free data retrieval call binding the contract method 0xa7e7b767.
//
// Solidity: function getRecordAtomicSwapVC(address _a, address _b, uint256 _sid) constant returns(uint256 seed0, uint256 seed0_rand, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeCallerSession) GetRecordAtomicSwapVC(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0     *big.Int
	Seed0Rand *big.Int
	SubmitAt  *big.Int
}, error) {
	return _zkFilExchange.Contract.GetRecordAtomicSwapVC(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetRecordComplaint is a free data retrieval call binding the contract method 0x1f31e95e.
//
// Solidity: function getRecordComplaint(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeCaller) GetRecordComplaint(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	Seed2    [32]byte
	KMklRoot [32]byte
	Count    uint64
	Price    *big.Int
	ExpireAt *big.Int
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		Seed0    [32]byte
		Seed2    [32]byte
		KMklRoot [32]byte
		Count    uint64
		Price    *big.Int
		ExpireAt *big.Int
		SubmitAt *big.Int
	})
	out := ret
	err := _zkFilExchange.contract.Call(opts, out, "getRecordComplaint", _a, _b, _sid)
	return *ret, err
}

// GetRecordComplaint is a free data retrieval call binding the contract method 0x1f31e95e.
//
// Solidity: function getRecordComplaint(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeSession) GetRecordComplaint(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	Seed2    [32]byte
	KMklRoot [32]byte
	Count    uint64
	Price    *big.Int
	ExpireAt *big.Int
	SubmitAt *big.Int
}, error) {
	return _zkFilExchange.Contract.GetRecordComplaint(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetRecordComplaint is a free data retrieval call binding the contract method 0x1f31e95e.
//
// Solidity: function getRecordComplaint(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeCallerSession) GetRecordComplaint(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	Seed2    [32]byte
	KMklRoot [32]byte
	Count    uint64
	Price    *big.Int
	ExpireAt *big.Int
	SubmitAt *big.Int
}, error) {
	return _zkFilExchange.Contract.GetRecordComplaint(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sid) constant returns(uint256 r, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeCaller) GetRecordVRF(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		R        *big.Int
		SubmitAt *big.Int
	})
	out := ret
	err := _zkFilExchange.contract.Call(opts, out, "getRecordVRF", _a, _b, _sid)
	return *ret, err
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sid) constant returns(uint256 r, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeSession) GetRecordVRF(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	return _zkFilExchange.Contract.GetRecordVRF(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sid) constant returns(uint256 r, uint256 submitAt)
func (_zkFilExchange *zkFilExchangeCallerSession) GetRecordVRF(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	return _zkFilExchange.Contract.GetRecordVRF(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetSessionRecord is a free data retrieval call binding the contract method 0xacff2f99.
//
// Solidity: function getSessionRecord(address _a, address _b, uint256 _sid) constant returns(uint256 submitAt, uint8 mode, uint8 stat)
func (_zkFilExchange *zkFilExchangeCaller) GetSessionRecord(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	SubmitAt *big.Int
	Mode     uint8
	Stat     uint8
}, error) {
	ret := new(struct {
		SubmitAt *big.Int
		Mode     uint8
		Stat     uint8
	})
	out := ret
	err := _zkFilExchange.contract.Call(opts, out, "getSessionRecord", _a, _b, _sid)
	return *ret, err
}

// GetSessionRecord is a free data retrieval call binding the contract method 0xacff2f99.
//
// Solidity: function getSessionRecord(address _a, address _b, uint256 _sid) constant returns(uint256 submitAt, uint8 mode, uint8 stat)
func (_zkFilExchange *zkFilExchangeSession) GetSessionRecord(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	SubmitAt *big.Int
	Mode     uint8
	Stat     uint8
}, error) {
	return _zkFilExchange.Contract.GetSessionRecord(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// GetSessionRecord is a free data retrieval call binding the contract method 0xacff2f99.
//
// Solidity: function getSessionRecord(address _a, address _b, uint256 _sid) constant returns(uint256 submitAt, uint8 mode, uint8 stat)
func (_zkFilExchange *zkFilExchangeCallerSession) GetSessionRecord(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	SubmitAt *big.Int
	Mode     uint8
	Stat     uint8
}, error) {
	return _zkFilExchange.Contract.GetSessionRecord(&_zkFilExchange.CallOpts, _a, _b, _sid)
}

// PublicVar is a free data retrieval call binding the contract method 0x0a6c93e0.
//
// Solidity: function publicVar_() constant returns(address)
func (_zkFilExchange *zkFilExchangeCaller) PublicVar(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _zkFilExchange.contract.Call(opts, out, "publicVar_")
	return *ret0, err
}

// PublicVar is a free data retrieval call binding the contract method 0x0a6c93e0.
//
// Solidity: function publicVar_() constant returns(address)
func (_zkFilExchange *zkFilExchangeSession) PublicVar() (common.Address, error) {
	return _zkFilExchange.Contract.PublicVar(&_zkFilExchange.CallOpts)
}

// PublicVar is a free data retrieval call binding the contract method 0x0a6c93e0.
//
// Solidity: function publicVar_() constant returns(address)
func (_zkFilExchange *zkFilExchangeCallerSession) PublicVar() (common.Address, error) {
	return _zkFilExchange.Contract.PublicVar(&_zkFilExchange.CallOpts)
}

// S is a free data retrieval call binding the contract method 0xc6b46963.
//
// Solidity: function s_() constant returns(uint64)
func (_zkFilExchange *zkFilExchangeCaller) S(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _zkFilExchange.contract.Call(opts, out, "s_")
	return *ret0, err
}

// S is a free data retrieval call binding the contract method 0xc6b46963.
//
// Solidity: function s_() constant returns(uint64)
func (_zkFilExchange *zkFilExchangeSession) S() (uint64, error) {
	return _zkFilExchange.Contract.S(&_zkFilExchange.CallOpts)
}

// S is a free data retrieval call binding the contract method 0xc6b46963.
//
// Solidity: function s_() constant returns(uint64)
func (_zkFilExchange *zkFilExchangeCallerSession) S() (uint64, error) {
	return _zkFilExchange.Contract.S(&_zkFilExchange.CallOpts)
}

// T1 is a free data retrieval call binding the contract method 0xfb5343f3.
//
// Solidity: function t1() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeCaller) T1(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _zkFilExchange.contract.Call(opts, out, "t1")
	return *ret0, err
}

// T1 is a free data retrieval call binding the contract method 0xfb5343f3.
//
// Solidity: function t1() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeSession) T1() (*big.Int, error) {
	return _zkFilExchange.Contract.T1(&_zkFilExchange.CallOpts)
}

// T1 is a free data retrieval call binding the contract method 0xfb5343f3.
//
// Solidity: function t1() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeCallerSession) T1() (*big.Int, error) {
	return _zkFilExchange.Contract.T1(&_zkFilExchange.CallOpts)
}

// T3 is a free data retrieval call binding the contract method 0xcfad78b1.
//
// Solidity: function t3() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeCaller) T3(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _zkFilExchange.contract.Call(opts, out, "t3")
	return *ret0, err
}

// T3 is a free data retrieval call binding the contract method 0xcfad78b1.
//
// Solidity: function t3() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeSession) T3() (*big.Int, error) {
	return _zkFilExchange.Contract.T3(&_zkFilExchange.CallOpts)
}

// T3 is a free data retrieval call binding the contract method 0xcfad78b1.
//
// Solidity: function t3() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeCallerSession) T3() (*big.Int, error) {
	return _zkFilExchange.Contract.T3(&_zkFilExchange.CallOpts)
}

// T4 is a free data retrieval call binding the contract method 0xfeae062d.
//
// Solidity: function t4() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeCaller) T4(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _zkFilExchange.contract.Call(opts, out, "t4")
	return *ret0, err
}

// T4 is a free data retrieval call binding the contract method 0xfeae062d.
//
// Solidity: function t4() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeSession) T4() (*big.Int, error) {
	return _zkFilExchange.Contract.T4(&_zkFilExchange.CallOpts)
}

// T4 is a free data retrieval call binding the contract method 0xfeae062d.
//
// Solidity: function t4() constant returns(uint256)
func (_zkFilExchange *zkFilExchangeCallerSession) T4() (*big.Int, error) {
	return _zkFilExchange.Contract.T4(&_zkFilExchange.CallOpts)
}

// BobDeposit is a paid mutator transaction binding the contract method 0x6bc76fdb.
//
// Solidity: function bobDeposit(address _to) returns()
func (_zkFilExchange *zkFilExchangeTransactor) BobDeposit(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "bobDeposit", _to)
}

// BobDeposit is a paid mutator transaction binding the contract method 0x6bc76fdb.
//
// Solidity: function bobDeposit(address _to) returns()
func (_zkFilExchange *zkFilExchangeSession) BobDeposit(_to common.Address) (*types.Transaction, error) {
	return _zkFilExchange.Contract.BobDeposit(&_zkFilExchange.TransactOpts, _to)
}

// BobDeposit is a paid mutator transaction binding the contract method 0x6bc76fdb.
//
// Solidity: function bobDeposit(address _to) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) BobDeposit(_to common.Address) (*types.Transaction, error) {
	return _zkFilExchange.Contract.BobDeposit(&_zkFilExchange.TransactOpts, _to)
}

// BobUnDeposit is a paid mutator transaction binding the contract method 0x30a84ef3.
//
// Solidity: function bobUnDeposit(address _to) returns()
func (_zkFilExchange *zkFilExchangeTransactor) BobUnDeposit(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "bobUnDeposit", _to)
}

// BobUnDeposit is a paid mutator transaction binding the contract method 0x30a84ef3.
//
// Solidity: function bobUnDeposit(address _to) returns()
func (_zkFilExchange *zkFilExchangeSession) BobUnDeposit(_to common.Address) (*types.Transaction, error) {
	return _zkFilExchange.Contract.BobUnDeposit(&_zkFilExchange.TransactOpts, _to)
}

// BobUnDeposit is a paid mutator transaction binding the contract method 0x30a84ef3.
//
// Solidity: function bobUnDeposit(address _to) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) BobUnDeposit(_to common.Address) (*types.Transaction, error) {
	return _zkFilExchange.Contract.BobUnDeposit(&_zkFilExchange.TransactOpts, _to)
}

// ClaimComplaint is a paid mutator transaction binding the contract method 0xf0e2e4e6.
//
// Solidity: function claimComplaint(address _a, uint256 _sid, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_zkFilExchange *zkFilExchangeTransactor) ClaimComplaint(opts *bind.TransactOpts, _a common.Address, _sid *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "claimComplaint", _a, _sid, _i, _j, _tx, _ty, _mkl_path, _sCnt)
}

// ClaimComplaint is a paid mutator transaction binding the contract method 0xf0e2e4e6.
//
// Solidity: function claimComplaint(address _a, uint256 _sid, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_zkFilExchange *zkFilExchangeSession) ClaimComplaint(_a common.Address, _sid *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _zkFilExchange.Contract.ClaimComplaint(&_zkFilExchange.TransactOpts, _a, _sid, _i, _j, _tx, _ty, _mkl_path, _sCnt)
}

// ClaimComplaint is a paid mutator transaction binding the contract method 0xf0e2e4e6.
//
// Solidity: function claimComplaint(address _a, uint256 _sid, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) ClaimComplaint(_a common.Address, _sid *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _zkFilExchange.Contract.ClaimComplaint(&_zkFilExchange.TransactOpts, _a, _sid, _i, _j, _tx, _ty, _mkl_path, _sCnt)
}

// Publish is a paid mutator transaction binding the contract method 0x56e25912.
//
// Solidity: function publish(uint64 _size, uint64 _s, uint64 _n, uint256 _sigma_mkl_root, uint256 _vrf_meta_digest, uint256 _blt_type) returns()
func (_zkFilExchange *zkFilExchangeTransactor) Publish(opts *bind.TransactOpts, _size uint64, _s uint64, _n uint64, _sigma_mkl_root *big.Int, _vrf_meta_digest *big.Int, _blt_type *big.Int) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "publish", _size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type)
}

// Publish is a paid mutator transaction binding the contract method 0x56e25912.
//
// Solidity: function publish(uint64 _size, uint64 _s, uint64 _n, uint256 _sigma_mkl_root, uint256 _vrf_meta_digest, uint256 _blt_type) returns()
func (_zkFilExchange *zkFilExchangeSession) Publish(_size uint64, _s uint64, _n uint64, _sigma_mkl_root *big.Int, _vrf_meta_digest *big.Int, _blt_type *big.Int) (*types.Transaction, error) {
	return _zkFilExchange.Contract.Publish(&_zkFilExchange.TransactOpts, _size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type)
}

// Publish is a paid mutator transaction binding the contract method 0x56e25912.
//
// Solidity: function publish(uint64 _size, uint64 _s, uint64 _n, uint256 _sigma_mkl_root, uint256 _vrf_meta_digest, uint256 _blt_type) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) Publish(_size uint64, _s uint64, _n uint64, _sigma_mkl_root *big.Int, _vrf_meta_digest *big.Int, _blt_type *big.Int) (*types.Transaction, error) {
	return _zkFilExchange.Contract.Publish(&_zkFilExchange.TransactOpts, _size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type)
}

// SettleComplaintDeal is a paid mutator transaction binding the contract method 0x629eee0a.
//
// Solidity: function settleComplaintDeal(address _a, address _b, uint256 _sid) returns()
func (_zkFilExchange *zkFilExchangeTransactor) SettleComplaintDeal(opts *bind.TransactOpts, _a common.Address, _b common.Address, _sid *big.Int) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "settleComplaintDeal", _a, _b, _sid)
}

// SettleComplaintDeal is a paid mutator transaction binding the contract method 0x629eee0a.
//
// Solidity: function settleComplaintDeal(address _a, address _b, uint256 _sid) returns()
func (_zkFilExchange *zkFilExchangeSession) SettleComplaintDeal(_a common.Address, _b common.Address, _sid *big.Int) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SettleComplaintDeal(&_zkFilExchange.TransactOpts, _a, _b, _sid)
}

// SettleComplaintDeal is a paid mutator transaction binding the contract method 0x629eee0a.
//
// Solidity: function settleComplaintDeal(address _a, address _b, uint256 _sid) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) SettleComplaintDeal(_a common.Address, _b common.Address, _sid *big.Int) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SettleComplaintDeal(&_zkFilExchange.TransactOpts, _a, _b, _sid)
}

// SubmitProofAtomicSwap is a paid mutator transaction binding the contract method 0x37c89b89.
//
// Solidity: function submitProofAtomicSwap(bytes32 _seed0, uint64 _sCnt, uint256 _sid, address[2] _addrs, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeTransactor) SubmitProofAtomicSwap(opts *bind.TransactOpts, _seed0 [32]byte, _sCnt uint64, _sid *big.Int, _addrs [2]common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "submitProofAtomicSwap", _seed0, _sCnt, _sid, _addrs, _seed2, _sigma_vw, _count, _price, _expireAt, _sig)
}

// SubmitProofAtomicSwap is a paid mutator transaction binding the contract method 0x37c89b89.
//
// Solidity: function submitProofAtomicSwap(bytes32 _seed0, uint64 _sCnt, uint256 _sid, address[2] _addrs, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeSession) SubmitProofAtomicSwap(_seed0 [32]byte, _sCnt uint64, _sid *big.Int, _addrs [2]common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SubmitProofAtomicSwap(&_zkFilExchange.TransactOpts, _seed0, _sCnt, _sid, _addrs, _seed2, _sigma_vw, _count, _price, _expireAt, _sig)
}

// SubmitProofAtomicSwap is a paid mutator transaction binding the contract method 0x37c89b89.
//
// Solidity: function submitProofAtomicSwap(bytes32 _seed0, uint64 _sCnt, uint256 _sid, address[2] _addrs, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) SubmitProofAtomicSwap(_seed0 [32]byte, _sCnt uint64, _sid *big.Int, _addrs [2]common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SubmitProofAtomicSwap(&_zkFilExchange.TransactOpts, _seed0, _sCnt, _sid, _addrs, _seed2, _sigma_vw, _count, _price, _expireAt, _sig)
}

// SubmitProofAtomicSwapVC is a paid mutator transaction binding the contract method 0x99ef8c5f.
//
// Solidity: function submitProofAtomicSwapVC(uint256 _seed0, uint256 _seed0_rand, uint256 _sid, address[2] _addrs, uint256 _seed0_mimc3_digest, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeTransactor) SubmitProofAtomicSwapVC(opts *bind.TransactOpts, _seed0 *big.Int, _seed0_rand *big.Int, _sid *big.Int, _addrs [2]common.Address, _seed0_mimc3_digest *big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "submitProofAtomicSwapVC", _seed0, _seed0_rand, _sid, _addrs, _seed0_mimc3_digest, _price, _expireAt, _sig)
}

// SubmitProofAtomicSwapVC is a paid mutator transaction binding the contract method 0x99ef8c5f.
//
// Solidity: function submitProofAtomicSwapVC(uint256 _seed0, uint256 _seed0_rand, uint256 _sid, address[2] _addrs, uint256 _seed0_mimc3_digest, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeSession) SubmitProofAtomicSwapVC(_seed0 *big.Int, _seed0_rand *big.Int, _sid *big.Int, _addrs [2]common.Address, _seed0_mimc3_digest *big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SubmitProofAtomicSwapVC(&_zkFilExchange.TransactOpts, _seed0, _seed0_rand, _sid, _addrs, _seed0_mimc3_digest, _price, _expireAt, _sig)
}

// SubmitProofAtomicSwapVC is a paid mutator transaction binding the contract method 0x99ef8c5f.
//
// Solidity: function submitProofAtomicSwapVC(uint256 _seed0, uint256 _seed0_rand, uint256 _sid, address[2] _addrs, uint256 _seed0_mimc3_digest, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) SubmitProofAtomicSwapVC(_seed0 *big.Int, _seed0_rand *big.Int, _sid *big.Int, _addrs [2]common.Address, _seed0_mimc3_digest *big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SubmitProofAtomicSwapVC(&_zkFilExchange.TransactOpts, _seed0, _seed0_rand, _sid, _addrs, _seed0_mimc3_digest, _price, _expireAt, _sig)
}

// SubmitProofComplaint is a paid mutator transaction binding the contract method 0x75843a12.
//
// Solidity: function submitProofComplaint(bytes32 _seed0, uint256 _sid, address[2] _addrs, bytes32 _bltKey, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeTransactor) SubmitProofComplaint(opts *bind.TransactOpts, _seed0 [32]byte, _sid *big.Int, _addrs [2]common.Address, _bltKey [32]byte, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "submitProofComplaint", _seed0, _sid, _addrs, _bltKey, _seed2, _k_mkl_root, _count, _price, _expireAt, _sig)
}

// SubmitProofComplaint is a paid mutator transaction binding the contract method 0x75843a12.
//
// Solidity: function submitProofComplaint(bytes32 _seed0, uint256 _sid, address[2] _addrs, bytes32 _bltKey, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeSession) SubmitProofComplaint(_seed0 [32]byte, _sid *big.Int, _addrs [2]common.Address, _bltKey [32]byte, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SubmitProofComplaint(&_zkFilExchange.TransactOpts, _seed0, _sid, _addrs, _bltKey, _seed2, _k_mkl_root, _count, _price, _expireAt, _sig)
}

// SubmitProofComplaint is a paid mutator transaction binding the contract method 0x75843a12.
//
// Solidity: function submitProofComplaint(bytes32 _seed0, uint256 _sid, address[2] _addrs, bytes32 _bltKey, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) SubmitProofComplaint(_seed0 [32]byte, _sid *big.Int, _addrs [2]common.Address, _bltKey [32]byte, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SubmitProofComplaint(&_zkFilExchange.TransactOpts, _seed0, _sid, _addrs, _bltKey, _seed2, _k_mkl_root, _count, _price, _expireAt, _sig)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0xe1752391.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sid, address[2] _addrs, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeTransactor) SubmitProofVRF(opts *bind.TransactOpts, _s_r *big.Int, _sid *big.Int, _addrs [2]common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "submitProofVRF", _s_r, _sid, _addrs, _g_exp_r, _price, _expireAt, _sig)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0xe1752391.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sid, address[2] _addrs, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeSession) SubmitProofVRF(_s_r *big.Int, _sid *big.Int, _addrs [2]common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SubmitProofVRF(&_zkFilExchange.TransactOpts, _s_r, _sid, _addrs, _g_exp_r, _price, _expireAt, _sig)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0xe1752391.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sid, address[2] _addrs, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) SubmitProofVRF(_s_r *big.Int, _sid *big.Int, _addrs [2]common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.SubmitProofVRF(&_zkFilExchange.TransactOpts, _s_r, _sid, _addrs, _g_exp_r, _price, _expireAt, _sig)
}

// UnPublish is a paid mutator transaction binding the contract method 0x52bb0780.
//
// Solidity: function unPublish(bytes32 _bltKey) returns()
func (_zkFilExchange *zkFilExchangeTransactor) UnPublish(opts *bind.TransactOpts, _bltKey [32]byte) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "unPublish", _bltKey)
}

// UnPublish is a paid mutator transaction binding the contract method 0x52bb0780.
//
// Solidity: function unPublish(bytes32 _bltKey) returns()
func (_zkFilExchange *zkFilExchangeSession) UnPublish(_bltKey [32]byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.UnPublish(&_zkFilExchange.TransactOpts, _bltKey)
}

// UnPublish is a paid mutator transaction binding the contract method 0x52bb0780.
//
// Solidity: function unPublish(bytes32 _bltKey) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) UnPublish(_bltKey [32]byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.UnPublish(&_zkFilExchange.TransactOpts, _bltKey)
}

// WithdrawA is a paid mutator transaction binding the contract method 0x9ecd2944.
//
// Solidity: function withdrawA(bytes32 _bltKey) returns()
func (_zkFilExchange *zkFilExchangeTransactor) WithdrawA(opts *bind.TransactOpts, _bltKey [32]byte) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "withdrawA", _bltKey)
}

// WithdrawA is a paid mutator transaction binding the contract method 0x9ecd2944.
//
// Solidity: function withdrawA(bytes32 _bltKey) returns()
func (_zkFilExchange *zkFilExchangeSession) WithdrawA(_bltKey [32]byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.WithdrawA(&_zkFilExchange.TransactOpts, _bltKey)
}

// WithdrawA is a paid mutator transaction binding the contract method 0x9ecd2944.
//
// Solidity: function withdrawA(bytes32 _bltKey) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) WithdrawA(_bltKey [32]byte) (*types.Transaction, error) {
	return _zkFilExchange.Contract.WithdrawA(&_zkFilExchange.TransactOpts, _bltKey)
}

// WithdrawB is a paid mutator transaction binding the contract method 0x2927251f.
//
// Solidity: function withdrawB(address _a) returns()
func (_zkFilExchange *zkFilExchangeTransactor) WithdrawB(opts *bind.TransactOpts, _a common.Address) (*types.Transaction, error) {
	return _zkFilExchange.contract.Transact(opts, "withdrawB", _a)
}

// WithdrawB is a paid mutator transaction binding the contract method 0x2927251f.
//
// Solidity: function withdrawB(address _a) returns()
func (_zkFilExchange *zkFilExchangeSession) WithdrawB(_a common.Address) (*types.Transaction, error) {
	return _zkFilExchange.Contract.WithdrawB(&_zkFilExchange.TransactOpts, _a)
}

// WithdrawB is a paid mutator transaction binding the contract method 0x2927251f.
//
// Solidity: function withdrawB(address _a) returns()
func (_zkFilExchange *zkFilExchangeTransactorSession) WithdrawB(_a common.Address) (*types.Transaction, error) {
	return _zkFilExchange.Contract.WithdrawB(&_zkFilExchange.TransactOpts, _a)
}

// zkFilExchangeOnAtomicSwapDealIterator is returned from FilterOnAtomicSwapDeal and is used to iterate over the raw logs and unpacked data for OnAtomicSwapDeal events raised by the zkFilExchange contract.
type zkFilExchangeOnAtomicSwapDealIterator struct {
	Event *zkFilExchangeOnAtomicSwapDeal // Event containing the contract specifics and raw log

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
func (it *zkFilExchangeOnAtomicSwapDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(zkFilExchangeOnAtomicSwapDeal)
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
		it.Event = new(zkFilExchangeOnAtomicSwapDeal)
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
func (it *zkFilExchangeOnAtomicSwapDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *zkFilExchangeOnAtomicSwapDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// zkFilExchangeOnAtomicSwapDeal represents a OnAtomicSwapDeal event raised by the zkFilExchange contract.
type zkFilExchangeOnAtomicSwapDeal struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	Seed0 [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnAtomicSwapDeal is a free log retrieval operation binding the contract event 0xff2c2f55ea8e89ebf103b88cb5f7806ec36789e2e8ffb192727fa47ccfec5d63.
//
// Solidity: event OnAtomicSwapDeal(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0)
func (_zkFilExchange *zkFilExchangeFilterer) FilterOnAtomicSwapDeal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*zkFilExchangeOnAtomicSwapDealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.FilterLogs(opts, "OnAtomicSwapDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeOnAtomicSwapDealIterator{contract: _zkFilExchange.contract, event: "OnAtomicSwapDeal", logs: logs, sub: sub}, nil
}

// WatchOnAtomicSwapDeal is a free log subscription operation binding the contract event 0xff2c2f55ea8e89ebf103b88cb5f7806ec36789e2e8ffb192727fa47ccfec5d63.
//
// Solidity: event OnAtomicSwapDeal(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0)
func (_zkFilExchange *zkFilExchangeFilterer) WatchOnAtomicSwapDeal(opts *bind.WatchOpts, sink chan<- *zkFilExchangeOnAtomicSwapDeal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.WatchLogs(opts, "OnAtomicSwapDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(zkFilExchangeOnAtomicSwapDeal)
				if err := _zkFilExchange.contract.UnpackLog(event, "OnAtomicSwapDeal", log); err != nil {
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

// zkFilExchangeOnAtomicSwapVCDealIterator is returned from FilterOnAtomicSwapVCDeal and is used to iterate over the raw logs and unpacked data for OnAtomicSwapVCDeal events raised by the zkFilExchange contract.
type zkFilExchangeOnAtomicSwapVCDealIterator struct {
	Event *zkFilExchangeOnAtomicSwapVCDeal // Event containing the contract specifics and raw log

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
func (it *zkFilExchangeOnAtomicSwapVCDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(zkFilExchangeOnAtomicSwapVCDeal)
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
		it.Event = new(zkFilExchangeOnAtomicSwapVCDeal)
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
func (it *zkFilExchangeOnAtomicSwapVCDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *zkFilExchangeOnAtomicSwapVCDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// zkFilExchangeOnAtomicSwapVCDeal represents a OnAtomicSwapVCDeal event raised by the zkFilExchange contract.
type zkFilExchangeOnAtomicSwapVCDeal struct {
	A         common.Address
	B         common.Address
	Sid       *big.Int
	Seed0     *big.Int
	Seed0Rand *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnAtomicSwapVCDeal is a free log retrieval operation binding the contract event 0x8bae2e4a75f8ccb83905dd727dbb4bf26457ca31ce29f776d48db5ed1794f028.
//
// Solidity: event OnAtomicSwapVCDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _seed0, uint256 _seed0_rand)
func (_zkFilExchange *zkFilExchangeFilterer) FilterOnAtomicSwapVCDeal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*zkFilExchangeOnAtomicSwapVCDealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.FilterLogs(opts, "OnAtomicSwapVCDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeOnAtomicSwapVCDealIterator{contract: _zkFilExchange.contract, event: "OnAtomicSwapVCDeal", logs: logs, sub: sub}, nil
}

// WatchOnAtomicSwapVCDeal is a free log subscription operation binding the contract event 0x8bae2e4a75f8ccb83905dd727dbb4bf26457ca31ce29f776d48db5ed1794f028.
//
// Solidity: event OnAtomicSwapVCDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _seed0, uint256 _seed0_rand)
func (_zkFilExchange *zkFilExchangeFilterer) WatchOnAtomicSwapVCDeal(opts *bind.WatchOpts, sink chan<- *zkFilExchangeOnAtomicSwapVCDeal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.WatchLogs(opts, "OnAtomicSwapVCDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(zkFilExchangeOnAtomicSwapVCDeal)
				if err := _zkFilExchange.contract.UnpackLog(event, "OnAtomicSwapVCDeal", log); err != nil {
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

// zkFilExchangeOnComplaintClaimIterator is returned from FilterOnComplaintClaim and is used to iterate over the raw logs and unpacked data for OnComplaintClaim events raised by the zkFilExchange contract.
type zkFilExchangeOnComplaintClaimIterator struct {
	Event *zkFilExchangeOnComplaintClaim // Event containing the contract specifics and raw log

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
func (it *zkFilExchangeOnComplaintClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(zkFilExchangeOnComplaintClaim)
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
		it.Event = new(zkFilExchangeOnComplaintClaim)
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
func (it *zkFilExchangeOnComplaintClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *zkFilExchangeOnComplaintClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// zkFilExchangeOnComplaintClaim represents a OnComplaintClaim event raised by the zkFilExchange contract.
type zkFilExchangeOnComplaintClaim struct {
	A   common.Address
	B   common.Address
	Sid *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnComplaintClaim is a free log retrieval operation binding the contract event 0xacba611c3ca0197a61bfbe86ec078a16c436808bdda2c1c248c534a595b1ad90.
//
// Solidity: event OnComplaintClaim(address indexed _a, address indexed _b, uint256 indexed _sid)
func (_zkFilExchange *zkFilExchangeFilterer) FilterOnComplaintClaim(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*zkFilExchangeOnComplaintClaimIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.FilterLogs(opts, "OnComplaintClaim", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeOnComplaintClaimIterator{contract: _zkFilExchange.contract, event: "OnComplaintClaim", logs: logs, sub: sub}, nil
}

// WatchOnComplaintClaim is a free log subscription operation binding the contract event 0xacba611c3ca0197a61bfbe86ec078a16c436808bdda2c1c248c534a595b1ad90.
//
// Solidity: event OnComplaintClaim(address indexed _a, address indexed _b, uint256 indexed _sid)
func (_zkFilExchange *zkFilExchangeFilterer) WatchOnComplaintClaim(opts *bind.WatchOpts, sink chan<- *zkFilExchangeOnComplaintClaim, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.WatchLogs(opts, "OnComplaintClaim", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(zkFilExchangeOnComplaintClaim)
				if err := _zkFilExchange.contract.UnpackLog(event, "OnComplaintClaim", log); err != nil {
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

// zkFilExchangeOnComplaintDealIterator is returned from FilterOnComplaintDeal and is used to iterate over the raw logs and unpacked data for OnComplaintDeal events raised by the zkFilExchange contract.
type zkFilExchangeOnComplaintDealIterator struct {
	Event *zkFilExchangeOnComplaintDeal // Event containing the contract specifics and raw log

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
func (it *zkFilExchangeOnComplaintDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(zkFilExchangeOnComplaintDeal)
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
		it.Event = new(zkFilExchangeOnComplaintDeal)
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
func (it *zkFilExchangeOnComplaintDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *zkFilExchangeOnComplaintDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// zkFilExchangeOnComplaintDeal represents a OnComplaintDeal event raised by the zkFilExchange contract.
type zkFilExchangeOnComplaintDeal struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnComplaintDeal is a free log retrieval operation binding the contract event 0xbb7539d148611cf3f9068b9e92a49185664f507b60523ae71e03216bcfc7e505.
//
// Solidity: event OnComplaintDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _price)
func (_zkFilExchange *zkFilExchangeFilterer) FilterOnComplaintDeal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*zkFilExchangeOnComplaintDealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.FilterLogs(opts, "OnComplaintDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeOnComplaintDealIterator{contract: _zkFilExchange.contract, event: "OnComplaintDeal", logs: logs, sub: sub}, nil
}

// WatchOnComplaintDeal is a free log subscription operation binding the contract event 0xbb7539d148611cf3f9068b9e92a49185664f507b60523ae71e03216bcfc7e505.
//
// Solidity: event OnComplaintDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _price)
func (_zkFilExchange *zkFilExchangeFilterer) WatchOnComplaintDeal(opts *bind.WatchOpts, sink chan<- *zkFilExchangeOnComplaintDeal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.WatchLogs(opts, "OnComplaintDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(zkFilExchangeOnComplaintDeal)
				if err := _zkFilExchange.contract.UnpackLog(event, "OnComplaintDeal", log); err != nil {
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

// zkFilExchangeOnComplaintKeyIterator is returned from FilterOnComplaintKey and is used to iterate over the raw logs and unpacked data for OnComplaintKey events raised by the zkFilExchange contract.
type zkFilExchangeOnComplaintKeyIterator struct {
	Event *zkFilExchangeOnComplaintKey // Event containing the contract specifics and raw log

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
func (it *zkFilExchangeOnComplaintKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(zkFilExchangeOnComplaintKey)
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
		it.Event = new(zkFilExchangeOnComplaintKey)
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
func (it *zkFilExchangeOnComplaintKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *zkFilExchangeOnComplaintKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// zkFilExchangeOnComplaintKey represents a OnComplaintKey event raised by the zkFilExchange contract.
type zkFilExchangeOnComplaintKey struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	Seed0 [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnComplaintKey is a free log retrieval operation binding the contract event 0x85218455b9d63667af97af61458f54558d6907d3e0b0eb772f6d4736b0c3eb52.
//
// Solidity: event OnComplaintKey(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0)
func (_zkFilExchange *zkFilExchangeFilterer) FilterOnComplaintKey(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*zkFilExchangeOnComplaintKeyIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.FilterLogs(opts, "OnComplaintKey", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeOnComplaintKeyIterator{contract: _zkFilExchange.contract, event: "OnComplaintKey", logs: logs, sub: sub}, nil
}

// WatchOnComplaintKey is a free log subscription operation binding the contract event 0x85218455b9d63667af97af61458f54558d6907d3e0b0eb772f6d4736b0c3eb52.
//
// Solidity: event OnComplaintKey(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0)
func (_zkFilExchange *zkFilExchangeFilterer) WatchOnComplaintKey(opts *bind.WatchOpts, sink chan<- *zkFilExchangeOnComplaintKey, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.WatchLogs(opts, "OnComplaintKey", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(zkFilExchangeOnComplaintKey)
				if err := _zkFilExchange.contract.UnpackLog(event, "OnComplaintKey", log); err != nil {
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

// zkFilExchangeOnDealIterator is returned from FilterOnDeal and is used to iterate over the raw logs and unpacked data for OnDeal events raised by the zkFilExchange contract.
type zkFilExchangeOnDealIterator struct {
	Event *zkFilExchangeOnDeal // Event containing the contract specifics and raw log

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
func (it *zkFilExchangeOnDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(zkFilExchangeOnDeal)
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
		it.Event = new(zkFilExchangeOnDeal)
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
func (it *zkFilExchangeOnDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *zkFilExchangeOnDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// zkFilExchangeOnDeal represents a OnDeal event raised by the zkFilExchange contract.
type zkFilExchangeOnDeal struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	Mode  uint8
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnDeal is a free log retrieval operation binding the contract event 0x37a5d6eac0a1d706c2eeea2c832256ebf3253a7b42d5a6cedd60e67f4195449b.
//
// Solidity: event OnDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint8 _mode, uint256 _price)
func (_zkFilExchange *zkFilExchangeFilterer) FilterOnDeal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*zkFilExchangeOnDealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.FilterLogs(opts, "OnDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeOnDealIterator{contract: _zkFilExchange.contract, event: "OnDeal", logs: logs, sub: sub}, nil
}

// WatchOnDeal is a free log subscription operation binding the contract event 0x37a5d6eac0a1d706c2eeea2c832256ebf3253a7b42d5a6cedd60e67f4195449b.
//
// Solidity: event OnDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint8 _mode, uint256 _price)
func (_zkFilExchange *zkFilExchangeFilterer) WatchOnDeal(opts *bind.WatchOpts, sink chan<- *zkFilExchangeOnDeal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.WatchLogs(opts, "OnDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(zkFilExchangeOnDeal)
				if err := _zkFilExchange.contract.UnpackLog(event, "OnDeal", log); err != nil {
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

// zkFilExchangeOnVRFDealIterator is returned from FilterOnVRFDeal and is used to iterate over the raw logs and unpacked data for OnVRFDeal events raised by the zkFilExchange contract.
type zkFilExchangeOnVRFDealIterator struct {
	Event *zkFilExchangeOnVRFDeal // Event containing the contract specifics and raw log

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
func (it *zkFilExchangeOnVRFDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(zkFilExchangeOnVRFDeal)
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
		it.Event = new(zkFilExchangeOnVRFDeal)
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
func (it *zkFilExchangeOnVRFDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *zkFilExchangeOnVRFDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// zkFilExchangeOnVRFDeal represents a OnVRFDeal event raised by the zkFilExchange contract.
type zkFilExchangeOnVRFDeal struct {
	A   common.Address
	B   common.Address
	Sid *big.Int
	R   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnVRFDeal is a free log retrieval operation binding the contract event 0x0cf813cbe764040d8f6a8bd61c65524fefa1b75e383c20d3673881e921173a52.
//
// Solidity: event OnVRFDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _r)
func (_zkFilExchange *zkFilExchangeFilterer) FilterOnVRFDeal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*zkFilExchangeOnVRFDealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.FilterLogs(opts, "OnVRFDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &zkFilExchangeOnVRFDealIterator{contract: _zkFilExchange.contract, event: "OnVRFDeal", logs: logs, sub: sub}, nil
}

// WatchOnVRFDeal is a free log subscription operation binding the contract event 0x0cf813cbe764040d8f6a8bd61c65524fefa1b75e383c20d3673881e921173a52.
//
// Solidity: event OnVRFDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _r)
func (_zkFilExchange *zkFilExchangeFilterer) WatchOnVRFDeal(opts *bind.WatchOpts, sink chan<- *zkFilExchangeOnVRFDeal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _zkFilExchange.contract.WatchLogs(opts, "OnVRFDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(zkFilExchangeOnVRFDeal)
				if err := _zkFilExchange.contract.UnpackLog(event, "OnVRFDeal", log); err != nil {
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
