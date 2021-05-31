// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

const IMintBurnABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IMintBurnFuncSigs maps the 4-byte function signature to its string representation.
var IMintBurnFuncSigs = map[string]string{
	"9dc29fac": "burn(address,uint256)",
	"40c10f19": "mint(address,uint256)",
}

// IMintBurn is an auto generated Go binding around an Ethereum contract.
type IMintBurn struct {
	IMintBurnCaller     // Read-only binding to the contract
	IMintBurnTransactor // Write-only binding to the contract
	IMintBurnFilterer   // Log filterer for contract events
}

// IMintBurnCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMintBurnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintBurnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMintBurnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintBurnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMintBurnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintBurnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMintBurnSession struct {
	Contract     *IMintBurn        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMintBurnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMintBurnCallerSession struct {
	Contract *IMintBurnCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IMintBurnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMintBurnTransactorSession struct {
	Contract     *IMintBurnTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMintBurnRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMintBurnRaw struct {
	Contract *IMintBurn // Generic contract binding to access the raw methods on
}

// IMintBurnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMintBurnCallerRaw struct {
	Contract *IMintBurnCaller // Generic read-only contract binding to access the raw methods on
}

// IMintBurnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMintBurnTransactorRaw struct {
	Contract *IMintBurnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMintBurn creates a new instance of IMintBurn, bound to a specific deployed contract.
func NewIMintBurn(address common.Address, backend bind.ContractBackend) (*IMintBurn, error) {
	contract, err := bindIMintBurn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMintBurn{IMintBurnCaller: IMintBurnCaller{contract: contract}, IMintBurnTransactor: IMintBurnTransactor{contract: contract}, IMintBurnFilterer: IMintBurnFilterer{contract: contract}}, nil
}

// NewIMintBurnCaller creates a new read-only instance of IMintBurn, bound to a specific deployed contract.
func NewIMintBurnCaller(address common.Address, caller bind.ContractCaller) (*IMintBurnCaller, error) {
	contract, err := bindIMintBurn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMintBurnCaller{contract: contract}, nil
}

// NewIMintBurnTransactor creates a new write-only instance of IMintBurn, bound to a specific deployed contract.
func NewIMintBurnTransactor(address common.Address, transactor bind.ContractTransactor) (*IMintBurnTransactor, error) {
	contract, err := bindIMintBurn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMintBurnTransactor{contract: contract}, nil
}

// NewIMintBurnFilterer creates a new log filterer instance of IMintBurn, bound to a specific deployed contract.
func NewIMintBurnFilterer(address common.Address, filterer bind.ContractFilterer) (*IMintBurnFilterer, error) {
	contract, err := bindIMintBurn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMintBurnFilterer{contract: contract}, nil
}

// bindIMintBurn binds a generic wrapper to an already deployed contract.
func bindIMintBurn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMintBurnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintBurn *IMintBurnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintBurn.Contract.IMintBurnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintBurn *IMintBurnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintBurn.Contract.IMintBurnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintBurn *IMintBurnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintBurn.Contract.IMintBurnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintBurn *IMintBurnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintBurn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintBurn *IMintBurnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintBurn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintBurn *IMintBurnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintBurn.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_IMintBurn *IMintBurnTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurn.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_IMintBurn *IMintBurnSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurn.Contract.Burn(&_IMintBurn.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_IMintBurn *IMintBurnTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurn.Contract.Burn(&_IMintBurn.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IMintBurn *IMintBurnTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurn.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IMintBurn *IMintBurnSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurn.Contract.Mint(&_IMintBurn.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IMintBurn *IMintBurnTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurn.Contract.Mint(&_IMintBurn.TransactOpts, account, amount)
}

// InterchainSwapABI is the input ABI used to generate the binding from.
const InterchainSwapABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_relayers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayIndex\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appchainIndex\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayTokenAddr\",\"type\":\"address\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"ethTokenAddrs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"relayTokenAddrs\",\"type\":\"address[]\"}],\"name\":\"addSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bxh2ethToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"eth2bxhToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"index2Height\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_txid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_appchainIndex\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"}],\"name\":\"removeSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// InterchainSwapFuncSigs maps the 4-byte function signature to its string representation.
var InterchainSwapFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"fced6ada": "PIER_ROLE()",
	"926d7d7f": "RELAYER_ROLE()",
	"7010584c": "addSupportToken(address,address)",
	"ab1494de": "addSupportTokens(address[],address[])",
	"b8ce670d": "burn(address,uint256,address)",
	"520de93e": "bxh2ethToken(address)",
	"a1090028": "eth2bxhToken(address)",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"9010d07c": "getRoleMember(bytes32,uint256)",
	"ca15c873": "getRoleMemberCount(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"96f963b4": "index2Height(uint256)",
	"ca8efd75": "mint(address,address,address,address,uint256,string,uint256)",
	"3c0a25e2": "mintAmount(address,address)",
	"e2769cfa": "removeSupportToken(address)",
	"0daff621": "removeSupportTokens(address[])",
	"36568abe": "renounceRole(bytes32,address)",
	"d547741f": "revokeRole(bytes32,address)",
	"10c27402": "txMinted(string)",
}

// InterchainSwapBin is the compiled bytecode used for deploying new contracts.
var InterchainSwapBin = "0x6080604052600060065560006007553480156200001b57600080fd5b50604051620018d1380380620018d18339810160408190526200003e91620001ca565b6200004b6000336200009e565b60005b815181101562000096576200008d6b52454c415945525f524f4c4560a01b8383815181106200007957fe5b60200260200101516200009e60201b60201c565b6001016200004e565b5050620002a6565b620000aa8282620000ae565b5050565b600082815260208181526040909120620000d391839062000aa462000127821b17901c565b15620000aa57620000e362000147565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60006200013e836001600160a01b0384166200014b565b90505b92915050565b3390565b60006200015983836200019a565b620001915750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000141565b50600062000141565b60009081526001919091016020526040902054151590565b80516001600160a01b03811681146200014157600080fd5b60006020808385031215620001dd578182fd5b82516001600160401b0380821115620001f4578384fd5b818501915085601f83011262000208578384fd5b81518181111562000217578485fd5b8381029150620002298483016200027f565b8181528481019084860184860187018a101562000244578788fd5b8795505b8386101562000272576200025d8a82620001b2565b83526001959095019491860191860162000248565b5098975050505050505050565b6040518181016001600160401b03811182821017156200029e57600080fd5b604052919050565b61161b80620002b66000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c8063926d7d7f116100b8578063b8ce670d1161007c578063b8ce670d14610275578063ca15c87314610288578063ca8efd751461029b578063d547741f146102ae578063e2769cfa146102c1578063fced6ada146102d457610137565b8063926d7d7f1461022c57806396f963b414610234578063a109002814610247578063a217fddf1461025a578063ab1494de1461026257610137565b80633c0a25e2116100ff5780633c0a25e2146101c0578063520de93e146101d35780637010584c146101f35780639010d07c1461020657806391d148541461021957610137565b80630daff6211461013c57806310c2740214610151578063248a9ca31461017a5780632f2ff15d1461019a57806336568abe146101ad575b600080fd5b61014f61014a366004611001565b6102dc565b005b61016461015f366004611105565b610310565b604051610171919061123e565b60405180910390f35b61018d61018836600461109d565b610330565b6040516101719190611249565b61014f6101a83660046110b5565b610345565b61014f6101bb3660046110b5565b610392565b61018d6101ce366004610ef7565b6103d4565b6101e66101e1366004610edc565b6103f1565b6040516101719190611180565b61014f610201366004610ef7565b61040c565b6101e66102143660046110e4565b6104eb565b6101646102273660046110b5565b61050c565b61018d610524565b61018d61024236600461109d565b610537565b6101e6610255366004610edc565b610549565b61018d610564565b61014f61027036600461103c565b610569565b61014f610283366004610fc0565b6105d3565b61018d61029636600461109d565b610751565b61014f6102a9366004610f2b565b610768565b61014f6102bc3660046110b5565b6109b5565b61014f6102cf366004610edc565b6109ef565b61018d610a94565b60005b815181101561030c576103048282815181106102f757fe5b60200260200101516109ef565b6001016102df565b5050565b805160208183018101805160058252928201919093012091525460ff1681565b60009081526020819052604090206002015490565b60008281526020819052604090206002015461036390610227610ab9565b6103885760405162461bcd60e51b815260040161037f906112d4565b60405180910390fd5b61030c8282610abd565b61039a610ab9565b6001600160a01b0316816001600160a01b0316146103ca5760405162461bcd60e51b815260040161037f90611527565b61030c8282610b26565b600360209081526000928352604080842090915290825290205481565b6002602052600090815260409020546001600160a01b031681565b61041760003361050c565b6104335760405162461bcd60e51b815260040161037f90611265565b6001600160a01b03828116600090815260016020526040902054161561046b5760405162461bcd60e51b815260040161037f906114cd565b6001600160a01b0381811660009081526002602052604090205416156104a35760405162461bcd60e51b815260040161037f906114cd565b6001600160a01b0391821660008181526001602090815260408083208054969095166001600160a01b0319968716811790955593825260029052919091208054909216179055565b60008281526020819052604081206105039083610b8f565b90505b92915050565b60008281526020819052604081206105039083610b9b565b6b52454c415945525f524f4c4560a01b81565b60046020526000908152604090205481565b6001602052600090815260409020546001600160a01b031681565b600081565b805182511461058a5760405162461bcd60e51b815260040161037f90611323565b60005b82518110156105ce576105c68382815181106105a557fe5b60200260200101518383815181106105b957fe5b602002602001015161040c565b60010161058d565b505050565b6001600160a01b03808416600090815260026020908152604080832054841680845260019092529091205490911661061d5760405162461bcd60e51b815260040161037f906113c1565b6001600160a01b038416600090815260036020908152604080832033845290915290205461064b9084610bb0565b6001600160a01b038516600081815260036020908152604080832033845290915290819020929092559051632770a7eb60e21b8152639dc29fac906106969085908790600401611225565b600060405180830381600087803b1580156106b057600080fd5b505af11580156106c4573d6000803e3d6000fd5b50506007546106d7925090506001610bf2565b600781815560009182526004602090815260408084204390556001600160a01b03808916855260029092529283902054915492517f7a06eaf30bdc55024f045fe4b59735921f35f3c55dda426771972cd96afc77509361074393909216918891339188918a9190611194565b60405180910390a150505050565b600081815260208190526040812061050690610c17565b6001600160a01b038088166000908152600160205260409020548891166107a15760405162461bcd60e51b815260040161037f906113c1565b6107b768504945525f524f4c4560b81b3361050c565b6107d35760405162461bcd60e51b815260040161037f90611475565b826005816040516107e49190611164565b9081526040519081900360200190205460ff16156108145760405162461bcd60e51b815260040161037f90611504565b6001600160a01b038981166000908152600160205260409020548116908916146108505760405162461bcd60e51b815260040161037f90611353565b60018303600654146108745760405162461bcd60e51b815260040161037f906114a4565b60016005856040516108869190611164565b908152604051908190036020019020805491151560ff199092169190911790556006546108b4906001610bf2565b6006556001600160a01b038089166000908152600360209081526040808320938a16835292905220546108e79086610bf2565b6001600160a01b03808a166000818152600360209081526040808320948c16835293905282902092909255516340c10f1960e01b81526340c10f19906109339089908990600401611225565b600060405180830381600087803b15801561094d57600080fd5b505af1158015610961573d6000803e3d6000fd5b505050507f719d9ed5b4a376c7a386b65ea2d152ec6adc4e971b1283ebad4ec6d8395c343a8989898989896006546040516109a297969594939291906111ce565b60405180910390a1505050505050505050565b6000828152602081905260409020600201546109d390610227610ab9565b6103ca5760405162461bcd60e51b815260040161037f90611425565b6109fa60003361050c565b610a165760405162461bcd60e51b815260040161037f90611265565b6001600160a01b0381811660009081526001602052604090205416610a4d5760405162461bcd60e51b815260040161037f906113f8565b6001600160a01b0390811660008181526001602081815260408084208054909616845260028252832080546001600160a01b03199081169091559390925290528154169055565b68504945525f524f4c4560b81b81565b6000610503836001600160a01b038416610c22565b3390565b6000828152602081905260409020610ad59082610aa4565b1561030c57610ae2610ab9565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152602081905260409020610b3e9082610c6c565b1561030c57610b4b610ab9565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b60006105038383610c81565b6000610503836001600160a01b038416610cc6565b600061050383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610cde565b6000828201838110156105035760405162461bcd60e51b815260040161037f9061138a565b600061050682610d0a565b6000610c2e8383610cc6565b610c6457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610506565b506000610506565b6000610503836001600160a01b038416610d0e565b81546000908210610ca45760405162461bcd60e51b815260040161037f90611292565b826000018281548110610cb357fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b60008184841115610d025760405162461bcd60e51b815260040161037f9190611252565b505050900390565b5490565b60008181526001830160205260408120548015610dca5783546000198083019190810190600090879083908110610d4157fe5b9060005260206000200154905080876000018481548110610d5e57fe5b600091825260208083209091019290925582815260018981019092526040902090840190558654879080610d8e57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610506565b6000915050610506565b80356001600160a01b038116811461050657600080fd5b600082601f830112610dfb578081fd5b813567ffffffffffffffff811115610e11578182fd5b6020808202610e21828201611576565b83815293508184018583018287018401881015610e3d57600080fd5b600092505b84831015610e6857610e548882610dd4565b825260019290920191908301908301610e42565b505050505092915050565b600082601f830112610e83578081fd5b813567ffffffffffffffff811115610e99578182fd5b610eac601f8201601f1916602001611576565b9150808252836020828501011115610ec357600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610eed578081fd5b6105038383610dd4565b60008060408385031215610f09578081fd5b610f138484610dd4565b9150610f228460208501610dd4565b90509250929050565b600080600080600080600060e0888a031215610f45578283fd5b610f4f8989610dd4565b9650610f5e8960208a01610dd4565b9550610f6d8960408a01610dd4565b9450610f7c8960608a01610dd4565b93506080880135925060a088013567ffffffffffffffff811115610f9e578283fd5b610faa8a828b01610e73565b92505060c0880135905092959891949750929550565b600080600060608486031215610fd4578283fd5b8335610fdf816115cd565b9250602084013591506040840135610ff6816115cd565b809150509250925092565b600060208284031215611012578081fd5b813567ffffffffffffffff811115611028578182fd5b61103484828501610deb565b949350505050565b6000806040838503121561104e578182fd5b823567ffffffffffffffff80821115611065578384fd5b61107186838701610deb565b93506020850135915080821115611086578283fd5b5061109385828601610deb565b9150509250929050565b6000602082840312156110ae578081fd5b5035919050565b600080604083850312156110c7578182fd5b8235915060208301356110d9816115cd565b809150509250929050565b600080604083850312156110f6578182fd5b50508035926020909101359150565b600060208284031215611116578081fd5b813567ffffffffffffffff81111561112c578182fd5b61103484828501610e73565b6000815180845261115081602086016020860161159d565b601f01601f19169290920160200192915050565b6000825161117681846020870161159d565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b0396871681529486166020860152928516604085015293166060830152608082019290925260a081019190915260c00190565b6001600160a01b03888116825287811660208301528681166040830152851660608201526080810184905260e060a0820181905260009061121190830185611138565b90508260c083015298975050505050505050565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b90815260200190565b6000602082526105036020830184611138565b60208082526013908201527231b0b63632b91034b9903737ba1030b236b4b760691b604082015260600190565b60208082526022908201527f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e604082015261647360f01b606082015260800190565b6020808252602f908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201526e0818591b5a5b881d1bc819dc985b9d608a1b606082015260800190565b6020808252601690820152750a8ded6cadc40d8cadccee8d040dcdee840dac2e8c6d60531b604082015260600190565b60208082526017908201527f4275726e3a3a4e6f7420537570706f727420546f6b656e000000000000000000604082015260600190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252601f908201527f4d696e74206f72204275726e3a3a4e6f7420537570706f727420546f6b656e00604082015260600190565b602080825260139082015272151bdad95b881b9bdd0814dd5c1c1bdc9d1959606a1b604082015260600190565b60208082526030908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201526f2061646d696e20746f207265766f6b6560801b606082015260800190565b60208082526015908201527431b0b63632b91034b9903737ba1031b937b9b9b2b960591b604082015260600190565b6020808252600f908201526e0d2dcc8caf040dcdee840dac2e8c6d608b1b604082015260600190565b60208082526017908201527f546f6b656e20616c726561647920537570706f72746564000000000000000000604082015260600190565b6020808252600990820152681d1e081b5a5b9d195960ba1b604082015260600190565b6020808252602f908201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560408201526e103937b632b9903337b91039b2b63360891b606082015260800190565b60405181810167ffffffffffffffff8111828210171561159557600080fd5b604052919050565b60005b838110156115b85781810151838201526020016115a0565b838111156115c7576000848401525b50505050565b6001600160a01b03811681146115e257600080fd5b5056fea264697066735822122031fb171133808c07823e7d3100b6b2b9edaf28f2a9b632de63ce1bdb52f3040564736f6c634300060c0033"

// DeployInterchainSwap deploys a new Ethereum contract, binding an instance of InterchainSwap to it.
func DeployInterchainSwap(auth *bind.TransactOpts, backend bind.ContractBackend, _relayers []common.Address) (common.Address, *types.Transaction, *InterchainSwap, error) {
	parsed, err := abi.JSON(strings.NewReader(InterchainSwapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InterchainSwapBin), backend, _relayers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainSwap{InterchainSwapCaller: InterchainSwapCaller{contract: contract}, InterchainSwapTransactor: InterchainSwapTransactor{contract: contract}, InterchainSwapFilterer: InterchainSwapFilterer{contract: contract}}, nil
}

// InterchainSwap is an auto generated Go binding around an Ethereum contract.
type InterchainSwap struct {
	InterchainSwapCaller     // Read-only binding to the contract
	InterchainSwapTransactor // Write-only binding to the contract
	InterchainSwapFilterer   // Log filterer for contract events
}

// InterchainSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainSwapSession struct {
	Contract     *InterchainSwap   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterchainSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainSwapCallerSession struct {
	Contract *InterchainSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// InterchainSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainSwapTransactorSession struct {
	Contract     *InterchainSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InterchainSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainSwapRaw struct {
	Contract *InterchainSwap // Generic contract binding to access the raw methods on
}

// InterchainSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainSwapCallerRaw struct {
	Contract *InterchainSwapCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainSwapTransactorRaw struct {
	Contract *InterchainSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainSwap creates a new instance of InterchainSwap, bound to a specific deployed contract.
func NewInterchainSwap(address common.Address, backend bind.ContractBackend) (*InterchainSwap, error) {
	contract, err := bindInterchainSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainSwap{InterchainSwapCaller: InterchainSwapCaller{contract: contract}, InterchainSwapTransactor: InterchainSwapTransactor{contract: contract}, InterchainSwapFilterer: InterchainSwapFilterer{contract: contract}}, nil
}

// NewInterchainSwapCaller creates a new read-only instance of InterchainSwap, bound to a specific deployed contract.
func NewInterchainSwapCaller(address common.Address, caller bind.ContractCaller) (*InterchainSwapCaller, error) {
	contract, err := bindInterchainSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainSwapCaller{contract: contract}, nil
}

// NewInterchainSwapTransactor creates a new write-only instance of InterchainSwap, bound to a specific deployed contract.
func NewInterchainSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainSwapTransactor, error) {
	contract, err := bindInterchainSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainSwapTransactor{contract: contract}, nil
}

// NewInterchainSwapFilterer creates a new log filterer instance of InterchainSwap, bound to a specific deployed contract.
func NewInterchainSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainSwapFilterer, error) {
	contract, err := bindInterchainSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainSwapFilterer{contract: contract}, nil
}

// bindInterchainSwap binds a generic wrapper to an already deployed contract.
func bindInterchainSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterchainSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainSwap *InterchainSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainSwap.Contract.InterchainSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainSwap *InterchainSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainSwap.Contract.InterchainSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainSwap *InterchainSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainSwap.Contract.InterchainSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainSwap *InterchainSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainSwap *InterchainSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainSwap *InterchainSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainSwap.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _InterchainSwap.Contract.DEFAULTADMINROLE(&_InterchainSwap.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _InterchainSwap.Contract.DEFAULTADMINROLE(&_InterchainSwap.CallOpts)
}

// PIERROLE is a free data retrieval call binding the contract method 0xfced6ada.
//
// Solidity: function PIER_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapCaller) PIERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "PIER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PIERROLE is a free data retrieval call binding the contract method 0xfced6ada.
//
// Solidity: function PIER_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapSession) PIERROLE() ([32]byte, error) {
	return _InterchainSwap.Contract.PIERROLE(&_InterchainSwap.CallOpts)
}

// PIERROLE is a free data retrieval call binding the contract method 0xfced6ada.
//
// Solidity: function PIER_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapCallerSession) PIERROLE() ([32]byte, error) {
	return _InterchainSwap.Contract.PIERROLE(&_InterchainSwap.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapSession) RELAYERROLE() ([32]byte, error) {
	return _InterchainSwap.Contract.RELAYERROLE(&_InterchainSwap.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_InterchainSwap *InterchainSwapCallerSession) RELAYERROLE() ([32]byte, error) {
	return _InterchainSwap.Contract.RELAYERROLE(&_InterchainSwap.CallOpts)
}

// Bxh2ethToken is a free data retrieval call binding the contract method 0x520de93e.
//
// Solidity: function bxh2ethToken(address ) view returns(address)
func (_InterchainSwap *InterchainSwapCaller) Bxh2ethToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "bxh2ethToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bxh2ethToken is a free data retrieval call binding the contract method 0x520de93e.
//
// Solidity: function bxh2ethToken(address ) view returns(address)
func (_InterchainSwap *InterchainSwapSession) Bxh2ethToken(arg0 common.Address) (common.Address, error) {
	return _InterchainSwap.Contract.Bxh2ethToken(&_InterchainSwap.CallOpts, arg0)
}

// Bxh2ethToken is a free data retrieval call binding the contract method 0x520de93e.
//
// Solidity: function bxh2ethToken(address ) view returns(address)
func (_InterchainSwap *InterchainSwapCallerSession) Bxh2ethToken(arg0 common.Address) (common.Address, error) {
	return _InterchainSwap.Contract.Bxh2ethToken(&_InterchainSwap.CallOpts, arg0)
}

// Eth2bxhToken is a free data retrieval call binding the contract method 0xa1090028.
//
// Solidity: function eth2bxhToken(address ) view returns(address)
func (_InterchainSwap *InterchainSwapCaller) Eth2bxhToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "eth2bxhToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eth2bxhToken is a free data retrieval call binding the contract method 0xa1090028.
//
// Solidity: function eth2bxhToken(address ) view returns(address)
func (_InterchainSwap *InterchainSwapSession) Eth2bxhToken(arg0 common.Address) (common.Address, error) {
	return _InterchainSwap.Contract.Eth2bxhToken(&_InterchainSwap.CallOpts, arg0)
}

// Eth2bxhToken is a free data retrieval call binding the contract method 0xa1090028.
//
// Solidity: function eth2bxhToken(address ) view returns(address)
func (_InterchainSwap *InterchainSwapCallerSession) Eth2bxhToken(arg0 common.Address) (common.Address, error) {
	return _InterchainSwap.Contract.Eth2bxhToken(&_InterchainSwap.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_InterchainSwap *InterchainSwapCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_InterchainSwap *InterchainSwapSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _InterchainSwap.Contract.GetRoleAdmin(&_InterchainSwap.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_InterchainSwap *InterchainSwapCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _InterchainSwap.Contract.GetRoleAdmin(&_InterchainSwap.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_InterchainSwap *InterchainSwapCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_InterchainSwap *InterchainSwapSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _InterchainSwap.Contract.GetRoleMember(&_InterchainSwap.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_InterchainSwap *InterchainSwapCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _InterchainSwap.Contract.GetRoleMember(&_InterchainSwap.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_InterchainSwap *InterchainSwapCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_InterchainSwap *InterchainSwapSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _InterchainSwap.Contract.GetRoleMemberCount(&_InterchainSwap.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_InterchainSwap *InterchainSwapCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _InterchainSwap.Contract.GetRoleMemberCount(&_InterchainSwap.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_InterchainSwap *InterchainSwapCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_InterchainSwap *InterchainSwapSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _InterchainSwap.Contract.HasRole(&_InterchainSwap.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_InterchainSwap *InterchainSwapCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _InterchainSwap.Contract.HasRole(&_InterchainSwap.CallOpts, role, account)
}

// Index2Height is a free data retrieval call binding the contract method 0x96f963b4.
//
// Solidity: function index2Height(uint256 ) view returns(uint256)
func (_InterchainSwap *InterchainSwapCaller) Index2Height(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "index2Height", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index2Height is a free data retrieval call binding the contract method 0x96f963b4.
//
// Solidity: function index2Height(uint256 ) view returns(uint256)
func (_InterchainSwap *InterchainSwapSession) Index2Height(arg0 *big.Int) (*big.Int, error) {
	return _InterchainSwap.Contract.Index2Height(&_InterchainSwap.CallOpts, arg0)
}

// Index2Height is a free data retrieval call binding the contract method 0x96f963b4.
//
// Solidity: function index2Height(uint256 ) view returns(uint256)
func (_InterchainSwap *InterchainSwapCallerSession) Index2Height(arg0 *big.Int) (*big.Int, error) {
	return _InterchainSwap.Contract.Index2Height(&_InterchainSwap.CallOpts, arg0)
}

// MintAmount is a free data retrieval call binding the contract method 0x3c0a25e2.
//
// Solidity: function mintAmount(address , address ) view returns(uint256)
func (_InterchainSwap *InterchainSwapCaller) MintAmount(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "mintAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintAmount is a free data retrieval call binding the contract method 0x3c0a25e2.
//
// Solidity: function mintAmount(address , address ) view returns(uint256)
func (_InterchainSwap *InterchainSwapSession) MintAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InterchainSwap.Contract.MintAmount(&_InterchainSwap.CallOpts, arg0, arg1)
}

// MintAmount is a free data retrieval call binding the contract method 0x3c0a25e2.
//
// Solidity: function mintAmount(address , address ) view returns(uint256)
func (_InterchainSwap *InterchainSwapCallerSession) MintAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InterchainSwap.Contract.MintAmount(&_InterchainSwap.CallOpts, arg0, arg1)
}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_InterchainSwap *InterchainSwapCaller) TxMinted(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _InterchainSwap.contract.Call(opts, &out, "txMinted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_InterchainSwap *InterchainSwapSession) TxMinted(arg0 string) (bool, error) {
	return _InterchainSwap.Contract.TxMinted(&_InterchainSwap.CallOpts, arg0)
}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_InterchainSwap *InterchainSwapCallerSession) TxMinted(arg0 string) (bool, error) {
	return _InterchainSwap.Contract.TxMinted(&_InterchainSwap.CallOpts, arg0)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x7010584c.
//
// Solidity: function addSupportToken(address ethTokenAddr, address relayTokenAddr) returns()
func (_InterchainSwap *InterchainSwapTransactor) AddSupportToken(opts *bind.TransactOpts, ethTokenAddr common.Address, relayTokenAddr common.Address) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "addSupportToken", ethTokenAddr, relayTokenAddr)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x7010584c.
//
// Solidity: function addSupportToken(address ethTokenAddr, address relayTokenAddr) returns()
func (_InterchainSwap *InterchainSwapSession) AddSupportToken(ethTokenAddr common.Address, relayTokenAddr common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.AddSupportToken(&_InterchainSwap.TransactOpts, ethTokenAddr, relayTokenAddr)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x7010584c.
//
// Solidity: function addSupportToken(address ethTokenAddr, address relayTokenAddr) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) AddSupportToken(ethTokenAddr common.Address, relayTokenAddr common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.AddSupportToken(&_InterchainSwap.TransactOpts, ethTokenAddr, relayTokenAddr)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xab1494de.
//
// Solidity: function addSupportTokens(address[] ethTokenAddrs, address[] relayTokenAddrs) returns()
func (_InterchainSwap *InterchainSwapTransactor) AddSupportTokens(opts *bind.TransactOpts, ethTokenAddrs []common.Address, relayTokenAddrs []common.Address) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "addSupportTokens", ethTokenAddrs, relayTokenAddrs)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xab1494de.
//
// Solidity: function addSupportTokens(address[] ethTokenAddrs, address[] relayTokenAddrs) returns()
func (_InterchainSwap *InterchainSwapSession) AddSupportTokens(ethTokenAddrs []common.Address, relayTokenAddrs []common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.AddSupportTokens(&_InterchainSwap.TransactOpts, ethTokenAddrs, relayTokenAddrs)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xab1494de.
//
// Solidity: function addSupportTokens(address[] ethTokenAddrs, address[] relayTokenAddrs) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) AddSupportTokens(ethTokenAddrs []common.Address, relayTokenAddrs []common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.AddSupportTokens(&_InterchainSwap.TransactOpts, ethTokenAddrs, relayTokenAddrs)
}

// Burn is a paid mutator transaction binding the contract method 0xb8ce670d.
//
// Solidity: function burn(address relayToken, uint256 amount, address recipient) returns()
func (_InterchainSwap *InterchainSwapTransactor) Burn(opts *bind.TransactOpts, relayToken common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "burn", relayToken, amount, recipient)
}

// Burn is a paid mutator transaction binding the contract method 0xb8ce670d.
//
// Solidity: function burn(address relayToken, uint256 amount, address recipient) returns()
func (_InterchainSwap *InterchainSwapSession) Burn(relayToken common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.Burn(&_InterchainSwap.TransactOpts, relayToken, amount, recipient)
}

// Burn is a paid mutator transaction binding the contract method 0xb8ce670d.
//
// Solidity: function burn(address relayToken, uint256 amount, address recipient) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) Burn(relayToken common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.Burn(&_InterchainSwap.TransactOpts, relayToken, amount, recipient)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.GrantRole(&_InterchainSwap.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.GrantRole(&_InterchainSwap.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xca8efd75.
//
// Solidity: function mint(address ethToken, address relayToken, address from, address recipient, uint256 amount, string _txid, uint256 _appchainIndex) returns()
func (_InterchainSwap *InterchainSwapTransactor) Mint(opts *bind.TransactOpts, ethToken common.Address, relayToken common.Address, from common.Address, recipient common.Address, amount *big.Int, _txid string, _appchainIndex *big.Int) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "mint", ethToken, relayToken, from, recipient, amount, _txid, _appchainIndex)
}

// Mint is a paid mutator transaction binding the contract method 0xca8efd75.
//
// Solidity: function mint(address ethToken, address relayToken, address from, address recipient, uint256 amount, string _txid, uint256 _appchainIndex) returns()
func (_InterchainSwap *InterchainSwapSession) Mint(ethToken common.Address, relayToken common.Address, from common.Address, recipient common.Address, amount *big.Int, _txid string, _appchainIndex *big.Int) (*types.Transaction, error) {
	return _InterchainSwap.Contract.Mint(&_InterchainSwap.TransactOpts, ethToken, relayToken, from, recipient, amount, _txid, _appchainIndex)
}

// Mint is a paid mutator transaction binding the contract method 0xca8efd75.
//
// Solidity: function mint(address ethToken, address relayToken, address from, address recipient, uint256 amount, string _txid, uint256 _appchainIndex) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) Mint(ethToken common.Address, relayToken common.Address, from common.Address, recipient common.Address, amount *big.Int, _txid string, _appchainIndex *big.Int) (*types.Transaction, error) {
	return _InterchainSwap.Contract.Mint(&_InterchainSwap.TransactOpts, ethToken, relayToken, from, recipient, amount, _txid, _appchainIndex)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address ethTokenAddr) returns()
func (_InterchainSwap *InterchainSwapTransactor) RemoveSupportToken(opts *bind.TransactOpts, ethTokenAddr common.Address) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "removeSupportToken", ethTokenAddr)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address ethTokenAddr) returns()
func (_InterchainSwap *InterchainSwapSession) RemoveSupportToken(ethTokenAddr common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.RemoveSupportToken(&_InterchainSwap.TransactOpts, ethTokenAddr)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address ethTokenAddr) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) RemoveSupportToken(ethTokenAddr common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.RemoveSupportToken(&_InterchainSwap.TransactOpts, ethTokenAddr)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_InterchainSwap *InterchainSwapTransactor) RemoveSupportTokens(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "removeSupportTokens", addrs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_InterchainSwap *InterchainSwapSession) RemoveSupportTokens(addrs []common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.RemoveSupportTokens(&_InterchainSwap.TransactOpts, addrs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) RemoveSupportTokens(addrs []common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.RemoveSupportTokens(&_InterchainSwap.TransactOpts, addrs)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.RenounceRole(&_InterchainSwap.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.RenounceRole(&_InterchainSwap.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.RevokeRole(&_InterchainSwap.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_InterchainSwap *InterchainSwapTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _InterchainSwap.Contract.RevokeRole(&_InterchainSwap.TransactOpts, role, account)
}

// InterchainSwapBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the InterchainSwap contract.
type InterchainSwapBurnIterator struct {
	Event *InterchainSwapBurn // Event containing the contract specifics and raw log

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
func (it *InterchainSwapBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainSwapBurn)
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
		it.Event = new(InterchainSwapBurn)
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
func (it *InterchainSwapBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainSwapBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainSwapBurn represents a Burn event raised by the InterchainSwap contract.
type InterchainSwapBurn struct {
	EthToken   common.Address
	RelayToken common.Address
	Burner     common.Address
	Recipient  common.Address
	Amount     *big.Int
	RelayIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x7a06eaf30bdc55024f045fe4b59735921f35f3c55dda426771972cd96afc7750.
//
// Solidity: event Burn(address ethToken, address relayToken, address burner, address recipient, uint256 amount, uint256 relayIndex)
func (_InterchainSwap *InterchainSwapFilterer) FilterBurn(opts *bind.FilterOpts) (*InterchainSwapBurnIterator, error) {

	logs, sub, err := _InterchainSwap.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &InterchainSwapBurnIterator{contract: _InterchainSwap.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x7a06eaf30bdc55024f045fe4b59735921f35f3c55dda426771972cd96afc7750.
//
// Solidity: event Burn(address ethToken, address relayToken, address burner, address recipient, uint256 amount, uint256 relayIndex)
func (_InterchainSwap *InterchainSwapFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *InterchainSwapBurn) (event.Subscription, error) {

	logs, sub, err := _InterchainSwap.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainSwapBurn)
				if err := _InterchainSwap.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x7a06eaf30bdc55024f045fe4b59735921f35f3c55dda426771972cd96afc7750.
//
// Solidity: event Burn(address ethToken, address relayToken, address burner, address recipient, uint256 amount, uint256 relayIndex)
func (_InterchainSwap *InterchainSwapFilterer) ParseBurn(log types.Log) (*InterchainSwapBurn, error) {
	event := new(InterchainSwapBurn)
	if err := _InterchainSwap.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InterchainSwapMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the InterchainSwap contract.
type InterchainSwapMintIterator struct {
	Event *InterchainSwapMint // Event containing the contract specifics and raw log

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
func (it *InterchainSwapMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainSwapMint)
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
		it.Event = new(InterchainSwapMint)
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
func (it *InterchainSwapMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainSwapMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainSwapMint represents a Mint event raised by the InterchainSwap contract.
type InterchainSwapMint struct {
	EthToken      common.Address
	RelayToken    common.Address
	From          common.Address
	Recipient     common.Address
	Amount        *big.Int
	Txid          string
	AppchainIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x719d9ed5b4a376c7a386b65ea2d152ec6adc4e971b1283ebad4ec6d8395c343a.
//
// Solidity: event Mint(address ethToken, address relayToken, address from, address recipient, uint256 amount, string txid, uint256 appchainIndex)
func (_InterchainSwap *InterchainSwapFilterer) FilterMint(opts *bind.FilterOpts) (*InterchainSwapMintIterator, error) {

	logs, sub, err := _InterchainSwap.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &InterchainSwapMintIterator{contract: _InterchainSwap.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x719d9ed5b4a376c7a386b65ea2d152ec6adc4e971b1283ebad4ec6d8395c343a.
//
// Solidity: event Mint(address ethToken, address relayToken, address from, address recipient, uint256 amount, string txid, uint256 appchainIndex)
func (_InterchainSwap *InterchainSwapFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *InterchainSwapMint) (event.Subscription, error) {

	logs, sub, err := _InterchainSwap.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainSwapMint)
				if err := _InterchainSwap.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x719d9ed5b4a376c7a386b65ea2d152ec6adc4e971b1283ebad4ec6d8395c343a.
//
// Solidity: event Mint(address ethToken, address relayToken, address from, address recipient, uint256 amount, string txid, uint256 appchainIndex)
func (_InterchainSwap *InterchainSwapFilterer) ParseMint(log types.Log) (*InterchainSwapMint, error) {
	event := new(InterchainSwapMint)
	if err := _InterchainSwap.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InterchainSwapRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the InterchainSwap contract.
type InterchainSwapRoleAdminChangedIterator struct {
	Event *InterchainSwapRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *InterchainSwapRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainSwapRoleAdminChanged)
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
		it.Event = new(InterchainSwapRoleAdminChanged)
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
func (it *InterchainSwapRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainSwapRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainSwapRoleAdminChanged represents a RoleAdminChanged event raised by the InterchainSwap contract.
type InterchainSwapRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_InterchainSwap *InterchainSwapFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*InterchainSwapRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _InterchainSwap.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &InterchainSwapRoleAdminChangedIterator{contract: _InterchainSwap.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_InterchainSwap *InterchainSwapFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *InterchainSwapRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _InterchainSwap.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainSwapRoleAdminChanged)
				if err := _InterchainSwap.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_InterchainSwap *InterchainSwapFilterer) ParseRoleAdminChanged(log types.Log) (*InterchainSwapRoleAdminChanged, error) {
	event := new(InterchainSwapRoleAdminChanged)
	if err := _InterchainSwap.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InterchainSwapRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the InterchainSwap contract.
type InterchainSwapRoleGrantedIterator struct {
	Event *InterchainSwapRoleGranted // Event containing the contract specifics and raw log

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
func (it *InterchainSwapRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainSwapRoleGranted)
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
		it.Event = new(InterchainSwapRoleGranted)
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
func (it *InterchainSwapRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainSwapRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainSwapRoleGranted represents a RoleGranted event raised by the InterchainSwap contract.
type InterchainSwapRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_InterchainSwap *InterchainSwapFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InterchainSwapRoleGrantedIterator, error) {

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

	logs, sub, err := _InterchainSwap.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InterchainSwapRoleGrantedIterator{contract: _InterchainSwap.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_InterchainSwap *InterchainSwapFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *InterchainSwapRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _InterchainSwap.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainSwapRoleGranted)
				if err := _InterchainSwap.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_InterchainSwap *InterchainSwapFilterer) ParseRoleGranted(log types.Log) (*InterchainSwapRoleGranted, error) {
	event := new(InterchainSwapRoleGranted)
	if err := _InterchainSwap.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InterchainSwapRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the InterchainSwap contract.
type InterchainSwapRoleRevokedIterator struct {
	Event *InterchainSwapRoleRevoked // Event containing the contract specifics and raw log

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
func (it *InterchainSwapRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainSwapRoleRevoked)
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
		it.Event = new(InterchainSwapRoleRevoked)
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
func (it *InterchainSwapRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainSwapRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainSwapRoleRevoked represents a RoleRevoked event raised by the InterchainSwap contract.
type InterchainSwapRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_InterchainSwap *InterchainSwapFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InterchainSwapRoleRevokedIterator, error) {

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

	logs, sub, err := _InterchainSwap.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InterchainSwapRoleRevokedIterator{contract: _InterchainSwap.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_InterchainSwap *InterchainSwapFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *InterchainSwapRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _InterchainSwap.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainSwapRoleRevoked)
				if err := _InterchainSwap.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_InterchainSwap *InterchainSwapFilterer) ParseRoleRevoked(log types.Log) (*InterchainSwapRoleRevoked, error) {
	event := new(InterchainSwapRoleRevoked)
	if err := _InterchainSwap.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}
