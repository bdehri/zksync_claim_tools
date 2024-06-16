// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ZkMerkleDistributorClaimSignatureInfo is an auto generated low-level Go binding around an user-defined struct.
type ZkMerkleDistributorClaimSignatureInfo struct {
	Claimant  common.Address
	Expiry    *big.Int
	Signature []byte
}

// ZkMerkleDistributorDelegateInfo is an auto generated low-level Go binding around an user-defined struct.
type ZkMerkleDistributorDelegateInfo struct {
	Delegatee common.Address
	Expiry    *big.Int
	Signature []byte
}

// ZkSyncMerkleDistributorMetaData contains all meta data concerning the ZkSyncMerkleDistributor contract.
var ZkSyncMerkleDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"contractIMintableAndDelegatable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maximumTotalClaimable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_windowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_windowEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkMerkleDistributor__AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkMerkleDistributor__ClaimAmountExceedsMaximum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkMerkleDistributor__ClaimWindowNotOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkMerkleDistributor__ClaimWindowNotYetClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkMerkleDistributor__ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkMerkleDistributor__InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkMerkleDistributor__InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkMerkleDistributor__SweepAlreadyDone\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ZkMerkleDistributor__Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_TOTAL_CLAIMABLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MERKLE_ROOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"contractIMintableAndDelegatable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WINDOW_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WINDOW_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZK_CLAIM_AND_DELEGATE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZK_CLAIM_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structZkMerkleDistributor.DelegateInfo\",\"name\":\"_delegateInfo\",\"type\":\"tuple\"}],\"name\":\"claimAndDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structZkMerkleDistributor.ClaimSignatureInfo\",\"name\":\"_claimSignatureInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structZkMerkleDistributor.DelegateInfo\",\"name\":\"_delegateInfo\",\"type\":\"tuple\"}],\"name\":\"claimAndDelegateOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structZkMerkleDistributor.ClaimSignatureInfo\",\"name\":\"_claimSignatureInfo\",\"type\":\"tuple\"}],\"name\":\"claimOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invalidateNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_unclaimedReceiver\",\"type\":\"address\"}],\"name\":\"sweepUnclaimed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZkSyncMerkleDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkSyncMerkleDistributorMetaData.ABI instead.
var ZkSyncMerkleDistributorABI = ZkSyncMerkleDistributorMetaData.ABI

// ZkSyncMerkleDistributor is an auto generated Go binding around an Ethereum contract.
type ZkSyncMerkleDistributor struct {
	ZkSyncMerkleDistributorCaller     // Read-only binding to the contract
	ZkSyncMerkleDistributorTransactor // Write-only binding to the contract
	ZkSyncMerkleDistributorFilterer   // Log filterer for contract events
}

// ZkSyncMerkleDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkSyncMerkleDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkSyncMerkleDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkSyncMerkleDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkSyncMerkleDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkSyncMerkleDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkSyncMerkleDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkSyncMerkleDistributorSession struct {
	Contract     *ZkSyncMerkleDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZkSyncMerkleDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkSyncMerkleDistributorCallerSession struct {
	Contract *ZkSyncMerkleDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ZkSyncMerkleDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkSyncMerkleDistributorTransactorSession struct {
	Contract     *ZkSyncMerkleDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ZkSyncMerkleDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkSyncMerkleDistributorRaw struct {
	Contract *ZkSyncMerkleDistributor // Generic contract binding to access the raw methods on
}

// ZkSyncMerkleDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkSyncMerkleDistributorCallerRaw struct {
	Contract *ZkSyncMerkleDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// ZkSyncMerkleDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkSyncMerkleDistributorTransactorRaw struct {
	Contract *ZkSyncMerkleDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkSyncMerkleDistributor creates a new instance of ZkSyncMerkleDistributor, bound to a specific deployed contract.
func NewZkSyncMerkleDistributor(address common.Address, backend bind.ContractBackend) (*ZkSyncMerkleDistributor, error) {
	contract, err := bindZkSyncMerkleDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkSyncMerkleDistributor{ZkSyncMerkleDistributorCaller: ZkSyncMerkleDistributorCaller{contract: contract}, ZkSyncMerkleDistributorTransactor: ZkSyncMerkleDistributorTransactor{contract: contract}, ZkSyncMerkleDistributorFilterer: ZkSyncMerkleDistributorFilterer{contract: contract}}, nil
}

// NewZkSyncMerkleDistributorCaller creates a new read-only instance of ZkSyncMerkleDistributor, bound to a specific deployed contract.
func NewZkSyncMerkleDistributorCaller(address common.Address, caller bind.ContractCaller) (*ZkSyncMerkleDistributorCaller, error) {
	contract, err := bindZkSyncMerkleDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkSyncMerkleDistributorCaller{contract: contract}, nil
}

// NewZkSyncMerkleDistributorTransactor creates a new write-only instance of ZkSyncMerkleDistributor, bound to a specific deployed contract.
func NewZkSyncMerkleDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkSyncMerkleDistributorTransactor, error) {
	contract, err := bindZkSyncMerkleDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkSyncMerkleDistributorTransactor{contract: contract}, nil
}

// NewZkSyncMerkleDistributorFilterer creates a new log filterer instance of ZkSyncMerkleDistributor, bound to a specific deployed contract.
func NewZkSyncMerkleDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkSyncMerkleDistributorFilterer, error) {
	contract, err := bindZkSyncMerkleDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkSyncMerkleDistributorFilterer{contract: contract}, nil
}

// bindZkSyncMerkleDistributor binds a generic wrapper to an already deployed contract.
func bindZkSyncMerkleDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkSyncMerkleDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkSyncMerkleDistributor.Contract.ZkSyncMerkleDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.ZkSyncMerkleDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.ZkSyncMerkleDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkSyncMerkleDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.contract.Transact(opts, method, params...)
}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(address)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) ADMIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "ADMIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(address)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) ADMIN() (common.Address, error) {
	return _ZkSyncMerkleDistributor.Contract.ADMIN(&_ZkSyncMerkleDistributor.CallOpts)
}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(address)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) ADMIN() (common.Address, error) {
	return _ZkSyncMerkleDistributor.Contract.ADMIN(&_ZkSyncMerkleDistributor.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ZkSyncMerkleDistributor.Contract.DOMAINSEPARATOR(&_ZkSyncMerkleDistributor.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ZkSyncMerkleDistributor.Contract.DOMAINSEPARATOR(&_ZkSyncMerkleDistributor.CallOpts)
}

// MAXIMUMTOTALCLAIMABLE is a free data retrieval call binding the contract method 0x392e2303.
//
// Solidity: function MAXIMUM_TOTAL_CLAIMABLE() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) MAXIMUMTOTALCLAIMABLE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "MAXIMUM_TOTAL_CLAIMABLE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMTOTALCLAIMABLE is a free data retrieval call binding the contract method 0x392e2303.
//
// Solidity: function MAXIMUM_TOTAL_CLAIMABLE() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) MAXIMUMTOTALCLAIMABLE() (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.MAXIMUMTOTALCLAIMABLE(&_ZkSyncMerkleDistributor.CallOpts)
}

// MAXIMUMTOTALCLAIMABLE is a free data retrieval call binding the contract method 0x392e2303.
//
// Solidity: function MAXIMUM_TOTAL_CLAIMABLE() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) MAXIMUMTOTALCLAIMABLE() (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.MAXIMUMTOTALCLAIMABLE(&_ZkSyncMerkleDistributor.CallOpts)
}

// MERKLEROOT is a free data retrieval call binding the contract method 0x51e75e8b.
//
// Solidity: function MERKLE_ROOT() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) MERKLEROOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "MERKLE_ROOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MERKLEROOT is a free data retrieval call binding the contract method 0x51e75e8b.
//
// Solidity: function MERKLE_ROOT() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) MERKLEROOT() ([32]byte, error) {
	return _ZkSyncMerkleDistributor.Contract.MERKLEROOT(&_ZkSyncMerkleDistributor.CallOpts)
}

// MERKLEROOT is a free data retrieval call binding the contract method 0x51e75e8b.
//
// Solidity: function MERKLE_ROOT() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) MERKLEROOT() ([32]byte, error) {
	return _ZkSyncMerkleDistributor.Contract.MERKLEROOT(&_ZkSyncMerkleDistributor.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) TOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) TOKEN() (common.Address, error) {
	return _ZkSyncMerkleDistributor.Contract.TOKEN(&_ZkSyncMerkleDistributor.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) TOKEN() (common.Address, error) {
	return _ZkSyncMerkleDistributor.Contract.TOKEN(&_ZkSyncMerkleDistributor.CallOpts)
}

// WINDOWEND is a free data retrieval call binding the contract method 0x1f23c68a.
//
// Solidity: function WINDOW_END() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) WINDOWEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "WINDOW_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WINDOWEND is a free data retrieval call binding the contract method 0x1f23c68a.
//
// Solidity: function WINDOW_END() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) WINDOWEND() (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.WINDOWEND(&_ZkSyncMerkleDistributor.CallOpts)
}

// WINDOWEND is a free data retrieval call binding the contract method 0x1f23c68a.
//
// Solidity: function WINDOW_END() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) WINDOWEND() (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.WINDOWEND(&_ZkSyncMerkleDistributor.CallOpts)
}

// WINDOWSTART is a free data retrieval call binding the contract method 0x37b82556.
//
// Solidity: function WINDOW_START() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) WINDOWSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "WINDOW_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WINDOWSTART is a free data retrieval call binding the contract method 0x37b82556.
//
// Solidity: function WINDOW_START() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) WINDOWSTART() (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.WINDOWSTART(&_ZkSyncMerkleDistributor.CallOpts)
}

// WINDOWSTART is a free data retrieval call binding the contract method 0x37b82556.
//
// Solidity: function WINDOW_START() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) WINDOWSTART() (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.WINDOWSTART(&_ZkSyncMerkleDistributor.CallOpts)
}

// ZKCLAIMANDDELEGATETYPEHASH is a free data retrieval call binding the contract method 0xb217293c.
//
// Solidity: function ZK_CLAIM_AND_DELEGATE_TYPEHASH() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) ZKCLAIMANDDELEGATETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "ZK_CLAIM_AND_DELEGATE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZKCLAIMANDDELEGATETYPEHASH is a free data retrieval call binding the contract method 0xb217293c.
//
// Solidity: function ZK_CLAIM_AND_DELEGATE_TYPEHASH() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) ZKCLAIMANDDELEGATETYPEHASH() ([32]byte, error) {
	return _ZkSyncMerkleDistributor.Contract.ZKCLAIMANDDELEGATETYPEHASH(&_ZkSyncMerkleDistributor.CallOpts)
}

// ZKCLAIMANDDELEGATETYPEHASH is a free data retrieval call binding the contract method 0xb217293c.
//
// Solidity: function ZK_CLAIM_AND_DELEGATE_TYPEHASH() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) ZKCLAIMANDDELEGATETYPEHASH() ([32]byte, error) {
	return _ZkSyncMerkleDistributor.Contract.ZKCLAIMANDDELEGATETYPEHASH(&_ZkSyncMerkleDistributor.CallOpts)
}

// ZKCLAIMTYPEHASH is a free data retrieval call binding the contract method 0x92a233c4.
//
// Solidity: function ZK_CLAIM_TYPEHASH() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) ZKCLAIMTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "ZK_CLAIM_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZKCLAIMTYPEHASH is a free data retrieval call binding the contract method 0x92a233c4.
//
// Solidity: function ZK_CLAIM_TYPEHASH() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) ZKCLAIMTYPEHASH() ([32]byte, error) {
	return _ZkSyncMerkleDistributor.Contract.ZKCLAIMTYPEHASH(&_ZkSyncMerkleDistributor.CallOpts)
}

// ZKCLAIMTYPEHASH is a free data retrieval call binding the contract method 0x92a233c4.
//
// Solidity: function ZK_CLAIM_TYPEHASH() view returns(bytes32)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) ZKCLAIMTYPEHASH() ([32]byte, error) {
	return _ZkSyncMerkleDistributor.Contract.ZKCLAIMTYPEHASH(&_ZkSyncMerkleDistributor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ZkSyncMerkleDistributor.Contract.Eip712Domain(&_ZkSyncMerkleDistributor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ZkSyncMerkleDistributor.Contract.Eip712Domain(&_ZkSyncMerkleDistributor.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 _index) view returns(bool)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) IsClaimed(opts *bind.CallOpts, _index *big.Int) (bool, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "isClaimed", _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 _index) view returns(bool)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) IsClaimed(_index *big.Int) (bool, error) {
	return _ZkSyncMerkleDistributor.Contract.IsClaimed(&_ZkSyncMerkleDistributor.CallOpts, _index)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 _index) view returns(bool)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) IsClaimed(_index *big.Int) (bool, error) {
	return _ZkSyncMerkleDistributor.Contract.IsClaimed(&_ZkSyncMerkleDistributor.CallOpts, _index)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.Nonces(&_ZkSyncMerkleDistributor.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.Nonces(&_ZkSyncMerkleDistributor.CallOpts, owner)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCaller) TotalClaimed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZkSyncMerkleDistributor.contract.Call(opts, &out, "totalClaimed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) TotalClaimed() (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.TotalClaimed(&_ZkSyncMerkleDistributor.CallOpts)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorCallerSession) TotalClaimed() (*big.Int, error) {
	return _ZkSyncMerkleDistributor.Contract.TotalClaimed(&_ZkSyncMerkleDistributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 _index, uint256 _amount, bytes32[] _merkleProof) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactor) Claim(opts *bind.TransactOpts, _index *big.Int, _amount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.contract.Transact(opts, "claim", _index, _amount, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 _index, uint256 _amount, bytes32[] _merkleProof) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) Claim(_index *big.Int, _amount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.Claim(&_ZkSyncMerkleDistributor.TransactOpts, _index, _amount, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 _index, uint256 _amount, bytes32[] _merkleProof) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactorSession) Claim(_index *big.Int, _amount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.Claim(&_ZkSyncMerkleDistributor.TransactOpts, _index, _amount, _merkleProof)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0xc41ac887.
//
// Solidity: function claimAndDelegate(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _delegateInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactor) ClaimAndDelegate(opts *bind.TransactOpts, _index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _delegateInfo ZkMerkleDistributorDelegateInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.contract.Transact(opts, "claimAndDelegate", _index, _amount, _merkleProof, _delegateInfo)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0xc41ac887.
//
// Solidity: function claimAndDelegate(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _delegateInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) ClaimAndDelegate(_index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _delegateInfo ZkMerkleDistributorDelegateInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.ClaimAndDelegate(&_ZkSyncMerkleDistributor.TransactOpts, _index, _amount, _merkleProof, _delegateInfo)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0xc41ac887.
//
// Solidity: function claimAndDelegate(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _delegateInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactorSession) ClaimAndDelegate(_index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _delegateInfo ZkMerkleDistributorDelegateInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.ClaimAndDelegate(&_ZkSyncMerkleDistributor.TransactOpts, _index, _amount, _merkleProof, _delegateInfo)
}

// ClaimAndDelegateOnBehalf is a paid mutator transaction binding the contract method 0x00d8e391.
//
// Solidity: function claimAndDelegateOnBehalf(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _claimSignatureInfo, (address,uint256,bytes) _delegateInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactor) ClaimAndDelegateOnBehalf(opts *bind.TransactOpts, _index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _claimSignatureInfo ZkMerkleDistributorClaimSignatureInfo, _delegateInfo ZkMerkleDistributorDelegateInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.contract.Transact(opts, "claimAndDelegateOnBehalf", _index, _amount, _merkleProof, _claimSignatureInfo, _delegateInfo)
}

// ClaimAndDelegateOnBehalf is a paid mutator transaction binding the contract method 0x00d8e391.
//
// Solidity: function claimAndDelegateOnBehalf(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _claimSignatureInfo, (address,uint256,bytes) _delegateInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) ClaimAndDelegateOnBehalf(_index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _claimSignatureInfo ZkMerkleDistributorClaimSignatureInfo, _delegateInfo ZkMerkleDistributorDelegateInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.ClaimAndDelegateOnBehalf(&_ZkSyncMerkleDistributor.TransactOpts, _index, _amount, _merkleProof, _claimSignatureInfo, _delegateInfo)
}

// ClaimAndDelegateOnBehalf is a paid mutator transaction binding the contract method 0x00d8e391.
//
// Solidity: function claimAndDelegateOnBehalf(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _claimSignatureInfo, (address,uint256,bytes) _delegateInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactorSession) ClaimAndDelegateOnBehalf(_index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _claimSignatureInfo ZkMerkleDistributorClaimSignatureInfo, _delegateInfo ZkMerkleDistributorDelegateInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.ClaimAndDelegateOnBehalf(&_ZkSyncMerkleDistributor.TransactOpts, _index, _amount, _merkleProof, _claimSignatureInfo, _delegateInfo)
}

// ClaimOnBehalf is a paid mutator transaction binding the contract method 0x875a00c3.
//
// Solidity: function claimOnBehalf(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _claimSignatureInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactor) ClaimOnBehalf(opts *bind.TransactOpts, _index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _claimSignatureInfo ZkMerkleDistributorClaimSignatureInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.contract.Transact(opts, "claimOnBehalf", _index, _amount, _merkleProof, _claimSignatureInfo)
}

// ClaimOnBehalf is a paid mutator transaction binding the contract method 0x875a00c3.
//
// Solidity: function claimOnBehalf(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _claimSignatureInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) ClaimOnBehalf(_index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _claimSignatureInfo ZkMerkleDistributorClaimSignatureInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.ClaimOnBehalf(&_ZkSyncMerkleDistributor.TransactOpts, _index, _amount, _merkleProof, _claimSignatureInfo)
}

// ClaimOnBehalf is a paid mutator transaction binding the contract method 0x875a00c3.
//
// Solidity: function claimOnBehalf(uint256 _index, uint256 _amount, bytes32[] _merkleProof, (address,uint256,bytes) _claimSignatureInfo) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactorSession) ClaimOnBehalf(_index *big.Int, _amount *big.Int, _merkleProof [][32]byte, _claimSignatureInfo ZkMerkleDistributorClaimSignatureInfo) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.ClaimOnBehalf(&_ZkSyncMerkleDistributor.TransactOpts, _index, _amount, _merkleProof, _claimSignatureInfo)
}

// InvalidateNonce is a paid mutator transaction binding the contract method 0x5a57b46f.
//
// Solidity: function invalidateNonce() returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactor) InvalidateNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.contract.Transact(opts, "invalidateNonce")
}

// InvalidateNonce is a paid mutator transaction binding the contract method 0x5a57b46f.
//
// Solidity: function invalidateNonce() returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) InvalidateNonce() (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.InvalidateNonce(&_ZkSyncMerkleDistributor.TransactOpts)
}

// InvalidateNonce is a paid mutator transaction binding the contract method 0x5a57b46f.
//
// Solidity: function invalidateNonce() returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactorSession) InvalidateNonce() (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.InvalidateNonce(&_ZkSyncMerkleDistributor.TransactOpts)
}

// SweepUnclaimed is a paid mutator transaction binding the contract method 0x404330bc.
//
// Solidity: function sweepUnclaimed(address _unclaimedReceiver) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactor) SweepUnclaimed(opts *bind.TransactOpts, _unclaimedReceiver common.Address) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.contract.Transact(opts, "sweepUnclaimed", _unclaimedReceiver)
}

// SweepUnclaimed is a paid mutator transaction binding the contract method 0x404330bc.
//
// Solidity: function sweepUnclaimed(address _unclaimedReceiver) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorSession) SweepUnclaimed(_unclaimedReceiver common.Address) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.SweepUnclaimed(&_ZkSyncMerkleDistributor.TransactOpts, _unclaimedReceiver)
}

// SweepUnclaimed is a paid mutator transaction binding the contract method 0x404330bc.
//
// Solidity: function sweepUnclaimed(address _unclaimedReceiver) returns()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorTransactorSession) SweepUnclaimed(_unclaimedReceiver common.Address) (*types.Transaction, error) {
	return _ZkSyncMerkleDistributor.Contract.SweepUnclaimed(&_ZkSyncMerkleDistributor.TransactOpts, _unclaimedReceiver)
}

// ZkSyncMerkleDistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the ZkSyncMerkleDistributor contract.
type ZkSyncMerkleDistributorClaimedIterator struct {
	Event *ZkSyncMerkleDistributorClaimed // Event containing the contract specifics and raw log

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
func (it *ZkSyncMerkleDistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncMerkleDistributorClaimed)
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
		it.Event = new(ZkSyncMerkleDistributorClaimed)
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
func (it *ZkSyncMerkleDistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncMerkleDistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncMerkleDistributorClaimed represents a Claimed event raised by the ZkSyncMerkleDistributor contract.
type ZkSyncMerkleDistributorClaimed struct {
	Index   *big.Int
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 index, address account, uint256 amount)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorFilterer) FilterClaimed(opts *bind.FilterOpts) (*ZkSyncMerkleDistributorClaimedIterator, error) {

	logs, sub, err := _ZkSyncMerkleDistributor.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &ZkSyncMerkleDistributorClaimedIterator{contract: _ZkSyncMerkleDistributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 index, address account, uint256 amount)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ZkSyncMerkleDistributorClaimed) (event.Subscription, error) {

	logs, sub, err := _ZkSyncMerkleDistributor.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncMerkleDistributorClaimed)
				if err := _ZkSyncMerkleDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 index, address account, uint256 amount)
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorFilterer) ParseClaimed(log types.Log) (*ZkSyncMerkleDistributorClaimed, error) {
	event := new(ZkSyncMerkleDistributorClaimed)
	if err := _ZkSyncMerkleDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkSyncMerkleDistributorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the ZkSyncMerkleDistributor contract.
type ZkSyncMerkleDistributorEIP712DomainChangedIterator struct {
	Event *ZkSyncMerkleDistributorEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ZkSyncMerkleDistributorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkSyncMerkleDistributorEIP712DomainChanged)
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
		it.Event = new(ZkSyncMerkleDistributorEIP712DomainChanged)
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
func (it *ZkSyncMerkleDistributorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkSyncMerkleDistributorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkSyncMerkleDistributorEIP712DomainChanged represents a EIP712DomainChanged event raised by the ZkSyncMerkleDistributor contract.
type ZkSyncMerkleDistributorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ZkSyncMerkleDistributorEIP712DomainChangedIterator, error) {

	logs, sub, err := _ZkSyncMerkleDistributor.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ZkSyncMerkleDistributorEIP712DomainChangedIterator{contract: _ZkSyncMerkleDistributor.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ZkSyncMerkleDistributorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _ZkSyncMerkleDistributor.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkSyncMerkleDistributorEIP712DomainChanged)
				if err := _ZkSyncMerkleDistributor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ZkSyncMerkleDistributor *ZkSyncMerkleDistributorFilterer) ParseEIP712DomainChanged(log types.Log) (*ZkSyncMerkleDistributorEIP712DomainChanged, error) {
	event := new(ZkSyncMerkleDistributorEIP712DomainChanged)
	if err := _ZkSyncMerkleDistributor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
