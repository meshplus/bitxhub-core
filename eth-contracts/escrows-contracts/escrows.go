// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package escrows_contracts

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

// EscrowsABI is the input ABI used to generate the binding from.
const EscrowsABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_relayers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appchainIndex\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dstChainId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dstContract\",\"type\":\"string\"}],\"name\":\"QuickSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayTokenAddr\",\"type\":\"address\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"ethTokenAddrs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"relayTokenAddrs\",\"type\":\"address[]\"}],\"name\":\"addSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appchainIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"index2Height\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dstChainId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dstContract\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"name\":\"quickSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"}],\"name\":\"removeSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_txid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_relayIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EscrowsFuncSigs maps the 4-byte function signature to its string representation.
var EscrowsFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"fced6ada": "PIER_ROLE()",
	"926d7d7f": "RELAYER_ROLE()",
	"7010584c": "addSupportToken(address,address)",
	"ab1494de": "addSupportTokens(address[],address[])",
	"82815874": "appchainIndex()",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"9010d07c": "getRoleMember(bytes32,uint256)",
	"ca15c873": "getRoleMemberCount(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"96f963b4": "index2Height(uint256)",
	"e1c7392a": "init()",
	"c267ce5f": "lock(address,uint256,string)",
	"95bc3bd0": "lockAmount(address)",
	"eedf5b7b": "quickSwap(string,string,address,uint256,string)",
	"a26c16a4": "relayIndex()",
	"e2769cfa": "removeSupportToken(address)",
	"0daff621": "removeSupportTokens(address[])",
	"36568abe": "renounceRole(bytes32,address)",
	"d547741f": "revokeRole(bytes32,address)",
	"2a4f1621": "supportToken(address)",
	"967145af": "txUnlocked(string)",
	"ef2abc1f": "unlock(address,address,address,uint256,string,uint256,bytes[])",
}

// EscrowsBin is the compiled bytecode used for deploying new contracts.
var EscrowsBin = "0x6080604052600060065560006007553480156200001b57600080fd5b5060405162001f6738038062001f678339810160408190526200003e91620001ca565b6200004b6000336200009e565b60005b815181101562000096576200008d6b52454c415945525f524f4c4560a01b8383815181106200007957fe5b60200260200101516200009e60201b60201c565b6001016200004e565b5050620002a6565b620000aa8282620000ae565b5050565b600082815260208181526040909120620000d391839062000b4d62000127821b17901c565b15620000aa57620000e362000147565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60006200013e836001600160a01b0384166200014b565b90505b92915050565b3390565b60006200015983836200019a565b620001915750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000141565b50600062000141565b60009081526001919091016020526040902054151590565b80516001600160a01b03811681146200014157600080fd5b60006020808385031215620001dd578182fd5b82516001600160401b0380821115620001f4578384fd5b818501915085601f83011262000208578384fd5b81518181111562000217578485fd5b8381029150620002298483016200027f565b8181528481019084860184860187018a101562000244578788fd5b8795505b8386101562000272576200025d8a82620001b2565b83526001959095019491860191860162000248565b5098975050505050505050565b6040518181016001600160401b03811182821017156200029e57600080fd5b604052919050565b611cb180620002b66000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806396f963b4116100c3578063d547741f1161007c578063d547741f146102b9578063e1c7392a146102cc578063e2769cfa146102d4578063eedf5b7b146102e7578063ef2abc1f146102fa578063fced6ada1461030d57610158565b806396f963b41461025d578063a217fddf14610270578063a26c16a414610278578063ab1494de14610280578063c267ce5f14610293578063ca15c873146102a657610158565b8063828158741161011557806382815874146101f45780639010d07c146101fc57806391d148541461020f578063926d7d7f1461022f57806395bc3bd014610237578063967145af1461024a57610158565b80630daff6211461015d578063248a9ca3146101725780632a4f16211461019b5780632f2ff15d146101bb57806336568abe146101ce5780637010584c146101e1575b600080fd5b61017061016b36600461144a565b610315565b005b6101856101803660046114fe565b610349565b60405161019291906117ff565b60405180910390f35b6101ae6101a93660046112a5565b61035e565b604051610192919061170b565b6101706101c9366004611516565b610379565b6101706101dc366004611516565b6103c6565b6101706101ef3660046112c0565b610408565b610185610495565b6101ae61020a366004611545565b61049b565b61022261021d366004611516565b6104bc565b60405161019291906117f4565b6101856104d4565b6101856102453660046112a5565b6104e7565b610222610258366004611566565b6104f9565b61018561026b3660046114fe565b610519565b61018561052b565b610185610530565b61017061028e36600461147d565b610536565b6101706102a13660046113f4565b6105a0565b6101856102b43660046114fe565b6106aa565b6101706102c7366004611516565b6106c1565b6101706106fb565b6101706102e23660046112a5565b61072e565b6101706102f5366004611599565b6107b3565b6101706103083660046112f4565b610838565b610185610b3d565b60005b81518110156103455761033d82828151811061033057fe5b602002602001015161072e565b600101610318565b5050565b60009081526020819052604090206002015490565b6001602052600090815260409020546001600160a01b031681565b6000828152602081905260409020600201546103979061021d610b62565b6103bc5760405162461bcd60e51b81526004016103b3906118d6565b60405180910390fd5b6103458282610b66565b6103ce610b62565b6001600160a01b0316816001600160a01b0316146103fe5760405162461bcd60e51b81526004016103b390611ba1565b6103458282610bcf565b6104136000336104bc565b61042f5760405162461bcd60e51b81526004016103b390611867565b6001600160a01b0382811660009081526001602052604090205416156104675760405162461bcd60e51b81526004016103b390611ab2565b6001600160a01b03918216600090815260016020526040902080546001600160a01b03191691909216179055565b60065481565b60008281526020819052604081206104b39083610c38565b90505b92915050565b60008281526020819052604081206104b39083610c44565b6b52454c415945525f524f4c4560a01b81565b60026020526000908152604090205481565b805160208183018101805160048252928201919093012091525460ff1681565b60036020526000908152604090205481565b600081565b60075481565b80518251146105575760405162461bcd60e51b81526004016103b390611925565b60005b825181101561059b5761059383828151811061057257fe5b602002602001015183838151811061058657fe5b6020026020010151610408565b60010161055a565b505050565b6001600160a01b038084166000908152600160205260409020548491166105d95760405162461bcd60e51b81526004016103b390611ae9565b6001600160a01b0384166000908152600260205260409020546105fc9084610c59565b6001600160a01b03851660008181526002602052604090209190915561062490333086610c7e565b600654610632906001610c59565b600681815560009182526003602090815260408084204390556001600160a01b03808916855260019092529283902054915492517f67741de31257ee580484c64e9f8b91449aa7df22ae38fcb86e50bdadfca0ad239361069c9389931691339188918a919061171f565b60405180910390a150505050565b60008181526020819052604081206104b690610cdc565b6000828152602081905260409020600201546106df9061021d610b62565b6103fe5760405162461bcd60e51b81526004016103b390611a0a565b6107066000336104bc565b6107225760405162461bcd60e51b81526004016103b390611867565b60006006819055600755565b6107396000336104bc565b6107555760405162461bcd60e51b81526004016103b390611867565b6001600160a01b038181166000908152600160205260409020541661078c5760405162461bcd60e51b81526004016103b3906119dd565b6001600160a01b0316600090815260016020526040902080546001600160a01b0319169055565b6001600160a01b038084166000908152600160205260409020548491166107ec5760405162461bcd60e51b81526004016103b390611ae9565b6107f78484846105a0565b7fbf42a9b8a78d1a7612fe5abd3e8bb6a6d68d4a31cc3a0ead24cdad19450fab838686604051610828929190611839565b60405180910390a1505050505050565b6001600160a01b038088166000908152600160205260409020548891166108715760405162461bcd60e51b81526004016103b390611ae9565b61088768504945525f524f4c4560b81b336104bc565b6108a35760405162461bcd60e51b81526004016103b390611a5a565b836004816040516108b491906116be565b9081526040519081900360200190205460ff16156108e45760405162461bcd60e51b81526004016103b3906119b8565b60018403600754146109085760405162461bcd60e51b81526004016103b390611a89565b60006109226b52454c415945525f524f4c4560a01b6106aa565b9050600060026003600019840104830181010490508085511015610947575050610b32565b60008b8b8b8b8b604051602001610962959493929190611665565b60405160208183030381529060405280519060200120905060005b86518110156109f05760006109ad61099484610ce7565b8984815181106109a057fe5b6020026020010151610d17565b90506109c86b52454c415945525f524f4c4560a01b826104bc565b156109e75760008381526005602052604090206109e59082610b4d565b505b5060010161097d565b5060008181526005602052604090208290610a0a90610cdc565b1015610a285760405162461bcd60e51b81526004016103b390611955565b6001600489604051610a3a91906116be565b908152604051908190036020019020805491151560ff19909216919091179055600754610a68906001610c59565b6007556001600160a01b038c16600090815260026020526040902054610a8e908a610df9565b6001600160a01b038d16600081815260026020526040902091909155610ab5908b8b610e3b565b7f5800596ed55e41c52e6e763640e29ff0d4c09c47408a512c4fb0974a452c1f0c8c600160008f6001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03168d8d8d8d604051610b2696959493929190611768565b60405180910390a15050505b505050505050505050565b68504945525f524f4c4560b81b81565b60006104b3836001600160a01b038416610e5a565b3390565b6000828152602081905260409020610b7e9082610b4d565b1561034557610b8b610b62565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152602081905260409020610be79082610ea4565b1561034557610bf4610b62565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b60006104b38383610eb9565b60006104b3836001600160a01b038416610efe565b6000828201838110156104b35760405162461bcd60e51b81526004016103b390611981565b610cd6846323b872dd60e01b858585604051602401610c9f939291906117b7565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610f16565b50505050565b60006104b682610fa5565b600081604051602001610cfa91906116da565b604051602081830303815290604052805190602001209050919050565b60008151604114610d2a575060006104b6565b60208201516040830151606084015160001a601b017f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610d7357600093505050506104b6565b8060ff16601b14158015610d8b57508060ff16601c14155b15610d9c57600093505050506104b6565b600060018783868660405160008152602001604052604051610dc19493929190611808565b6020604051602081039080840390855afa158015610de3573d6000803e3d6000fd5b5050604051601f19015198975050505050505050565b60006104b383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610fa9565b61059b8363a9059cbb60e01b8484604051602401610c9f9291906117db565b6000610e668383610efe565b610e9c575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556104b6565b5060006104b6565b60006104b3836001600160a01b038416610fd5565b81546000908210610edc5760405162461bcd60e51b81526004016103b390611894565b826000018281548110610eeb57fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b6060610f6b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661109b9092919063ffffffff16565b80519091501561059b5780806020019051810190610f8991906114de565b61059b5760405162461bcd60e51b81526004016103b390611b57565b5490565b60008184841115610fcd5760405162461bcd60e51b81526004016103b39190611826565b505050900390565b60008181526001830160205260408120548015611091578354600019808301919081019060009087908390811061100857fe5b906000526020600020015490508087600001848154811061102557fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061105557fe5b600190038181906000526020600020016000905590558660010160008781526020019081526020016000206000905560019450505050506104b6565b60009150506104b6565b60606110aa84846000856110b2565b949350505050565b60606110bd85611176565b6110d95760405162461bcd60e51b81526004016103b390611b20565b60006060866001600160a01b031685876040516110f691906116be565b60006040518083038185875af1925050503d8060008114611133576040519150601f19603f3d011682016040523d82523d6000602084013e611138565b606091505b5091509150811561114c5791506110aa9050565b80511561115c5780518082602001fd5b8360405162461bcd60e51b81526004016103b39190611826565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906110aa575050151592915050565b80356001600160a01b03811681146104b657600080fd5b600082601f8301126111d6578081fd5b81356111e96111e482611c17565b611bf0565b81815291506020808301908481018184028601820187101561120a57600080fd5b60005b848110156112315761121f88836111af565b8452928201929082019060010161120d565b505050505092915050565b600082601f83011261124c578081fd5b813567ffffffffffffffff811115611262578182fd5b611275601f8201601f1916602001611bf0565b915080825283602082850101111561128c57600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156112b6578081fd5b6104b383836111af565b600080604083850312156112d2578081fd5b6112dc84846111af565b91506112eb84602085016111af565b90509250929050565b600080600080600080600060e0888a03121561130e578283fd5b6113188835611c63565b87359650602088013561132a81611c63565b9550604088013561133a81611c63565b945060608801359350608088013567ffffffffffffffff8082111561135d578485fd5b6113698b838c0161123c565b945060a08a0135935060c08a0135915080821115611385578283fd5b508801601f81018a13611396578182fd5b80356113a46111e482611c17565b818152602080820191908401855b848110156113df576113ca8f6020843589010161123c565b845260209384019391909101906001016113b2565b50508094505050505092959891949750929550565b600080600060608486031215611408578283fd5b61141285856111af565b925060208401359150604084013567ffffffffffffffff811115611434578182fd5b6114408682870161123c565b9150509250925092565b60006020828403121561145b578081fd5b813567ffffffffffffffff811115611471578182fd5b6110aa848285016111c6565b6000806040838503121561148f578182fd5b823567ffffffffffffffff808211156114a6578384fd5b6114b2868387016111c6565b935060208501359150808211156114c7578283fd5b506114d4858286016111c6565b9150509250929050565b6000602082840312156114ef578081fd5b815180151581146104b3578182fd5b60006020828403121561150f578081fd5b5035919050565b60008060408385031215611528578182fd5b82359150602083013561153a81611c63565b809150509250929050565b60008060408385031215611557578182fd5b50508035926020909101359150565b600060208284031215611577578081fd5b813567ffffffffffffffff81111561158d578182fd5b6110aa8482850161123c565b600080600080600060a086880312156115b0578081fd5b853567ffffffffffffffff808211156115c7578283fd5b6115d389838a0161123c565b965060208801359150808211156115e8578283fd5b6115f489838a0161123c565b95506116038960408a016111af565b945060608801359350608088013591508082111561161f578283fd5b5061162c8882890161123c565b9150509295509295909350565b60008151808452611651816020860160208601611c37565b601f01601f19169290920160200192915050565b60006bffffffffffffffffffffffff19808860601b168352808760601b166014840152808660601b1660288401525083603c83015282516116ad81605c850160208701611c37565b91909101605c019695505050505050565b600082516116d0818460208701611c37565b9190910192915050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b6001600160a01b0391909116815260200190565b6001600160a01b03878116825286811660208301528516604082015260c06060820181905260009061175390830186611639565b60808301949094525060a00152949350505050565b6001600160a01b03878116825286811660208301528581166040830152841660608201526080810183905260c060a082018190526000906117ab90830184611639565b98975050505050505050565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b90815260200190565b93845260ff9290921660208401526040830152606082015260800190565b6000602082526104b36020830184611639565b60006040825261184c6040830185611639565b828103602084015261185e8185611639565b95945050505050565b60208082526013908201527231b0b63632b91034b9903737ba1030b236b4b760691b604082015260600190565b60208082526022908201527f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e604082015261647360f01b606082015260800190565b6020808252602f908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201526e0818591b5a5b881d1bc819dc985b9d608a1b606082015260800190565b6020808252601690820152750a8ded6cadc40d8cadccee8d040dcdee840dac2e8c6d60531b604082015260600190565b6020808252601290820152711cda59db985d1d5c995cc81a5b9d985b1a5960721b604082015260600190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252600b908201526a1d1e081d5b9b1bd8dad95960aa1b604082015260600190565b602080825260139082015272151bdad95b881b9bdd0814dd5c1c1bdc9d1959606a1b604082015260600190565b60208082526030908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201526f2061646d696e20746f207265766f6b6560801b606082015260800190565b60208082526015908201527431b0b63632b91034b9903737ba1031b937b9b9b2b960591b604082015260600190565b6020808252600f908201526e0d2dcc8caf040dcdee840dac2e8c6d608b1b604082015260600190565b60208082526017908201527f546f6b656e20616c726561647920537570706f72746564000000000000000000604082015260600190565b60208082526017908201527f4c6f636b3a3a4e6f7420537570706f727420546f6b656e000000000000000000604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b6020808252602f908201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560408201526e103937b632b9903337b91039b2b63360891b606082015260800190565b60405181810167ffffffffffffffff81118282101715611c0f57600080fd5b604052919050565b600067ffffffffffffffff821115611c2d578081fd5b5060209081020190565b60005b83811015611c52578181015183820152602001611c3a565b83811115610cd65750506000910152565b6001600160a01b0381168114611c7857600080fd5b5056fea2646970667358221220a6c5c4a3e80ca683a02c09f9febdf3a5caa2adfda89b00f8fb853c56f22e165464736f6c634300060c0033"

// DeployEscrows deploys a new Ethereum contract, binding an instance of Escrows to it.
func DeployEscrows(auth *bind.TransactOpts, backend bind.ContractBackend, _relayers []common.Address) (common.Address, *types.Transaction, *Escrows, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EscrowsBin), backend, _relayers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Escrows{EscrowsCaller: EscrowsCaller{contract: contract}, EscrowsTransactor: EscrowsTransactor{contract: contract}, EscrowsFilterer: EscrowsFilterer{contract: contract}}, nil
}

// Escrows is an auto generated Go binding around an Ethereum contract.
type Escrows struct {
	EscrowsCaller     // Read-only binding to the contract
	EscrowsTransactor // Write-only binding to the contract
	EscrowsFilterer   // Log filterer for contract events
}

// EscrowsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowsSession struct {
	Contract     *Escrows          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowsCallerSession struct {
	Contract *EscrowsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EscrowsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowsTransactorSession struct {
	Contract     *EscrowsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EscrowsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowsRaw struct {
	Contract *Escrows // Generic contract binding to access the raw methods on
}

// EscrowsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowsCallerRaw struct {
	Contract *EscrowsCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowsTransactorRaw struct {
	Contract *EscrowsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrows creates a new instance of Escrows, bound to a specific deployed contract.
func NewEscrows(address common.Address, backend bind.ContractBackend) (*Escrows, error) {
	contract, err := bindEscrows(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrows{EscrowsCaller: EscrowsCaller{contract: contract}, EscrowsTransactor: EscrowsTransactor{contract: contract}, EscrowsFilterer: EscrowsFilterer{contract: contract}}, nil
}

// NewEscrowsCaller creates a new read-only instance of Escrows, bound to a specific deployed contract.
func NewEscrowsCaller(address common.Address, caller bind.ContractCaller) (*EscrowsCaller, error) {
	contract, err := bindEscrows(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowsCaller{contract: contract}, nil
}

// NewEscrowsTransactor creates a new write-only instance of Escrows, bound to a specific deployed contract.
func NewEscrowsTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowsTransactor, error) {
	contract, err := bindEscrows(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowsTransactor{contract: contract}, nil
}

// NewEscrowsFilterer creates a new log filterer instance of Escrows, bound to a specific deployed contract.
func NewEscrowsFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowsFilterer, error) {
	contract, err := bindEscrows(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowsFilterer{contract: contract}, nil
}

// bindEscrows binds a generic wrapper to an already deployed contract.
func bindEscrows(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrows *EscrowsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrows.Contract.EscrowsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrows *EscrowsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrows.Contract.EscrowsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrows *EscrowsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrows.Contract.EscrowsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrows *EscrowsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrows.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrows *EscrowsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrows.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrows *EscrowsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrows.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Escrows *EscrowsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Escrows *EscrowsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Escrows.Contract.DEFAULTADMINROLE(&_Escrows.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Escrows *EscrowsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Escrows.Contract.DEFAULTADMINROLE(&_Escrows.CallOpts)
}

// PIERROLE is a free data retrieval call binding the contract method 0xfced6ada.
//
// Solidity: function PIER_ROLE() view returns(bytes32)
func (_Escrows *EscrowsCaller) PIERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "PIER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PIERROLE is a free data retrieval call binding the contract method 0xfced6ada.
//
// Solidity: function PIER_ROLE() view returns(bytes32)
func (_Escrows *EscrowsSession) PIERROLE() ([32]byte, error) {
	return _Escrows.Contract.PIERROLE(&_Escrows.CallOpts)
}

// PIERROLE is a free data retrieval call binding the contract method 0xfced6ada.
//
// Solidity: function PIER_ROLE() view returns(bytes32)
func (_Escrows *EscrowsCallerSession) PIERROLE() ([32]byte, error) {
	return _Escrows.Contract.PIERROLE(&_Escrows.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Escrows *EscrowsCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Escrows *EscrowsSession) RELAYERROLE() ([32]byte, error) {
	return _Escrows.Contract.RELAYERROLE(&_Escrows.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Escrows *EscrowsCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Escrows.Contract.RELAYERROLE(&_Escrows.CallOpts)
}

// AppchainIndex is a free data retrieval call binding the contract method 0x82815874.
//
// Solidity: function appchainIndex() view returns(uint256)
func (_Escrows *EscrowsCaller) AppchainIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "appchainIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AppchainIndex is a free data retrieval call binding the contract method 0x82815874.
//
// Solidity: function appchainIndex() view returns(uint256)
func (_Escrows *EscrowsSession) AppchainIndex() (*big.Int, error) {
	return _Escrows.Contract.AppchainIndex(&_Escrows.CallOpts)
}

// AppchainIndex is a free data retrieval call binding the contract method 0x82815874.
//
// Solidity: function appchainIndex() view returns(uint256)
func (_Escrows *EscrowsCallerSession) AppchainIndex() (*big.Int, error) {
	return _Escrows.Contract.AppchainIndex(&_Escrows.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Escrows *EscrowsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Escrows *EscrowsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Escrows.Contract.GetRoleAdmin(&_Escrows.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Escrows *EscrowsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Escrows.Contract.GetRoleAdmin(&_Escrows.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Escrows *EscrowsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Escrows *EscrowsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Escrows.Contract.GetRoleMember(&_Escrows.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Escrows *EscrowsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Escrows.Contract.GetRoleMember(&_Escrows.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Escrows *EscrowsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Escrows *EscrowsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Escrows.Contract.GetRoleMemberCount(&_Escrows.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Escrows *EscrowsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Escrows.Contract.GetRoleMemberCount(&_Escrows.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Escrows *EscrowsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Escrows *EscrowsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Escrows.Contract.HasRole(&_Escrows.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Escrows *EscrowsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Escrows.Contract.HasRole(&_Escrows.CallOpts, role, account)
}

// Index2Height is a free data retrieval call binding the contract method 0x96f963b4.
//
// Solidity: function index2Height(uint256 ) view returns(uint256)
func (_Escrows *EscrowsCaller) Index2Height(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "index2Height", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index2Height is a free data retrieval call binding the contract method 0x96f963b4.
//
// Solidity: function index2Height(uint256 ) view returns(uint256)
func (_Escrows *EscrowsSession) Index2Height(arg0 *big.Int) (*big.Int, error) {
	return _Escrows.Contract.Index2Height(&_Escrows.CallOpts, arg0)
}

// Index2Height is a free data retrieval call binding the contract method 0x96f963b4.
//
// Solidity: function index2Height(uint256 ) view returns(uint256)
func (_Escrows *EscrowsCallerSession) Index2Height(arg0 *big.Int) (*big.Int, error) {
	return _Escrows.Contract.Index2Height(&_Escrows.CallOpts, arg0)
}

// LockAmount is a free data retrieval call binding the contract method 0x95bc3bd0.
//
// Solidity: function lockAmount(address ) view returns(uint256)
func (_Escrows *EscrowsCaller) LockAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "lockAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockAmount is a free data retrieval call binding the contract method 0x95bc3bd0.
//
// Solidity: function lockAmount(address ) view returns(uint256)
func (_Escrows *EscrowsSession) LockAmount(arg0 common.Address) (*big.Int, error) {
	return _Escrows.Contract.LockAmount(&_Escrows.CallOpts, arg0)
}

// LockAmount is a free data retrieval call binding the contract method 0x95bc3bd0.
//
// Solidity: function lockAmount(address ) view returns(uint256)
func (_Escrows *EscrowsCallerSession) LockAmount(arg0 common.Address) (*big.Int, error) {
	return _Escrows.Contract.LockAmount(&_Escrows.CallOpts, arg0)
}

// RelayIndex is a free data retrieval call binding the contract method 0xa26c16a4.
//
// Solidity: function relayIndex() view returns(uint256)
func (_Escrows *EscrowsCaller) RelayIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "relayIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayIndex is a free data retrieval call binding the contract method 0xa26c16a4.
//
// Solidity: function relayIndex() view returns(uint256)
func (_Escrows *EscrowsSession) RelayIndex() (*big.Int, error) {
	return _Escrows.Contract.RelayIndex(&_Escrows.CallOpts)
}

// RelayIndex is a free data retrieval call binding the contract method 0xa26c16a4.
//
// Solidity: function relayIndex() view returns(uint256)
func (_Escrows *EscrowsCallerSession) RelayIndex() (*big.Int, error) {
	return _Escrows.Contract.RelayIndex(&_Escrows.CallOpts)
}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(address)
func (_Escrows *EscrowsCaller) SupportToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "supportToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(address)
func (_Escrows *EscrowsSession) SupportToken(arg0 common.Address) (common.Address, error) {
	return _Escrows.Contract.SupportToken(&_Escrows.CallOpts, arg0)
}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(address)
func (_Escrows *EscrowsCallerSession) SupportToken(arg0 common.Address) (common.Address, error) {
	return _Escrows.Contract.SupportToken(&_Escrows.CallOpts, arg0)
}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_Escrows *EscrowsCaller) TxUnlocked(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "txUnlocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_Escrows *EscrowsSession) TxUnlocked(arg0 string) (bool, error) {
	return _Escrows.Contract.TxUnlocked(&_Escrows.CallOpts, arg0)
}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_Escrows *EscrowsCallerSession) TxUnlocked(arg0 string) (bool, error) {
	return _Escrows.Contract.TxUnlocked(&_Escrows.CallOpts, arg0)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x7010584c.
//
// Solidity: function addSupportToken(address ethTokenAddr, address relayTokenAddr) returns()
func (_Escrows *EscrowsTransactor) AddSupportToken(opts *bind.TransactOpts, ethTokenAddr common.Address, relayTokenAddr common.Address) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "addSupportToken", ethTokenAddr, relayTokenAddr)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x7010584c.
//
// Solidity: function addSupportToken(address ethTokenAddr, address relayTokenAddr) returns()
func (_Escrows *EscrowsSession) AddSupportToken(ethTokenAddr common.Address, relayTokenAddr common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.AddSupportToken(&_Escrows.TransactOpts, ethTokenAddr, relayTokenAddr)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x7010584c.
//
// Solidity: function addSupportToken(address ethTokenAddr, address relayTokenAddr) returns()
func (_Escrows *EscrowsTransactorSession) AddSupportToken(ethTokenAddr common.Address, relayTokenAddr common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.AddSupportToken(&_Escrows.TransactOpts, ethTokenAddr, relayTokenAddr)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xab1494de.
//
// Solidity: function addSupportTokens(address[] ethTokenAddrs, address[] relayTokenAddrs) returns()
func (_Escrows *EscrowsTransactor) AddSupportTokens(opts *bind.TransactOpts, ethTokenAddrs []common.Address, relayTokenAddrs []common.Address) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "addSupportTokens", ethTokenAddrs, relayTokenAddrs)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xab1494de.
//
// Solidity: function addSupportTokens(address[] ethTokenAddrs, address[] relayTokenAddrs) returns()
func (_Escrows *EscrowsSession) AddSupportTokens(ethTokenAddrs []common.Address, relayTokenAddrs []common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.AddSupportTokens(&_Escrows.TransactOpts, ethTokenAddrs, relayTokenAddrs)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xab1494de.
//
// Solidity: function addSupportTokens(address[] ethTokenAddrs, address[] relayTokenAddrs) returns()
func (_Escrows *EscrowsTransactorSession) AddSupportTokens(ethTokenAddrs []common.Address, relayTokenAddrs []common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.AddSupportTokens(&_Escrows.TransactOpts, ethTokenAddrs, relayTokenAddrs)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.GrantRole(&_Escrows.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.GrantRole(&_Escrows.TransactOpts, role, account)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Escrows *EscrowsTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Escrows *EscrowsSession) Init() (*types.Transaction, error) {
	return _Escrows.Contract.Init(&_Escrows.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Escrows *EscrowsTransactorSession) Init() (*types.Transaction, error) {
	return _Escrows.Contract.Init(&_Escrows.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address token, uint256 amount, string recipient) returns()
func (_Escrows *EscrowsTransactor) Lock(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "lock", token, amount, recipient)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address token, uint256 amount, string recipient) returns()
func (_Escrows *EscrowsSession) Lock(token common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Escrows.Contract.Lock(&_Escrows.TransactOpts, token, amount, recipient)
}

// Lock is a paid mutator transaction binding the contract method 0xc267ce5f.
//
// Solidity: function lock(address token, uint256 amount, string recipient) returns()
func (_Escrows *EscrowsTransactorSession) Lock(token common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Escrows.Contract.Lock(&_Escrows.TransactOpts, token, amount, recipient)
}

// QuickSwap is a paid mutator transaction binding the contract method 0xeedf5b7b.
//
// Solidity: function quickSwap(string dstChainId, string dstContract, address token, uint256 amount, string recipient) returns()
func (_Escrows *EscrowsTransactor) QuickSwap(opts *bind.TransactOpts, dstChainId string, dstContract string, token common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "quickSwap", dstChainId, dstContract, token, amount, recipient)
}

// QuickSwap is a paid mutator transaction binding the contract method 0xeedf5b7b.
//
// Solidity: function quickSwap(string dstChainId, string dstContract, address token, uint256 amount, string recipient) returns()
func (_Escrows *EscrowsSession) QuickSwap(dstChainId string, dstContract string, token common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Escrows.Contract.QuickSwap(&_Escrows.TransactOpts, dstChainId, dstContract, token, amount, recipient)
}

// QuickSwap is a paid mutator transaction binding the contract method 0xeedf5b7b.
//
// Solidity: function quickSwap(string dstChainId, string dstContract, address token, uint256 amount, string recipient) returns()
func (_Escrows *EscrowsTransactorSession) QuickSwap(dstChainId string, dstContract string, token common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Escrows.Contract.QuickSwap(&_Escrows.TransactOpts, dstChainId, dstContract, token, amount, recipient)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address ethTokenAddr) returns()
func (_Escrows *EscrowsTransactor) RemoveSupportToken(opts *bind.TransactOpts, ethTokenAddr common.Address) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "removeSupportToken", ethTokenAddr)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address ethTokenAddr) returns()
func (_Escrows *EscrowsSession) RemoveSupportToken(ethTokenAddr common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.RemoveSupportToken(&_Escrows.TransactOpts, ethTokenAddr)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address ethTokenAddr) returns()
func (_Escrows *EscrowsTransactorSession) RemoveSupportToken(ethTokenAddr common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.RemoveSupportToken(&_Escrows.TransactOpts, ethTokenAddr)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_Escrows *EscrowsTransactor) RemoveSupportTokens(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "removeSupportTokens", addrs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_Escrows *EscrowsSession) RemoveSupportTokens(addrs []common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.RemoveSupportTokens(&_Escrows.TransactOpts, addrs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_Escrows *EscrowsTransactorSession) RemoveSupportTokens(addrs []common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.RemoveSupportTokens(&_Escrows.TransactOpts, addrs)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.RenounceRole(&_Escrows.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.RenounceRole(&_Escrows.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.RevokeRole(&_Escrows.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Escrows *EscrowsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.RevokeRole(&_Escrows.TransactOpts, role, account)
}

// Unlock is a paid mutator transaction binding the contract method 0xef2abc1f.
//
// Solidity: function unlock(address token, address from, address recipient, uint256 amount, string _txid, uint256 _relayIndex, bytes[] signatures) returns()
func (_Escrows *EscrowsTransactor) Unlock(opts *bind.TransactOpts, token common.Address, from common.Address, recipient common.Address, amount *big.Int, _txid string, _relayIndex *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "unlock", token, from, recipient, amount, _txid, _relayIndex, signatures)
}

// Unlock is a paid mutator transaction binding the contract method 0xef2abc1f.
//
// Solidity: function unlock(address token, address from, address recipient, uint256 amount, string _txid, uint256 _relayIndex, bytes[] signatures) returns()
func (_Escrows *EscrowsSession) Unlock(token common.Address, from common.Address, recipient common.Address, amount *big.Int, _txid string, _relayIndex *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Escrows.Contract.Unlock(&_Escrows.TransactOpts, token, from, recipient, amount, _txid, _relayIndex, signatures)
}

// Unlock is a paid mutator transaction binding the contract method 0xef2abc1f.
//
// Solidity: function unlock(address token, address from, address recipient, uint256 amount, string _txid, uint256 _relayIndex, bytes[] signatures) returns()
func (_Escrows *EscrowsTransactorSession) Unlock(token common.Address, from common.Address, recipient common.Address, amount *big.Int, _txid string, _relayIndex *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Escrows.Contract.Unlock(&_Escrows.TransactOpts, token, from, recipient, amount, _txid, _relayIndex, signatures)
}

// EscrowsLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the Escrows contract.
type EscrowsLockIterator struct {
	Event *EscrowsLock // Event containing the contract specifics and raw log

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
func (it *EscrowsLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowsLock)
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
		it.Event = new(EscrowsLock)
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
func (it *EscrowsLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowsLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowsLock represents a Lock event raised by the Escrows contract.
type EscrowsLock struct {
	EthToken      common.Address
	RelayToken    common.Address
	Locker        common.Address
	Recipient     string
	Amount        *big.Int
	AppchainIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x67741de31257ee580484c64e9f8b91449aa7df22ae38fcb86e50bdadfca0ad23.
//
// Solidity: event Lock(address ethToken, address relayToken, address locker, string recipient, uint256 amount, uint256 appchainIndex)
func (_Escrows *EscrowsFilterer) FilterLock(opts *bind.FilterOpts) (*EscrowsLockIterator, error) {

	logs, sub, err := _Escrows.contract.FilterLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return &EscrowsLockIterator{contract: _Escrows.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x67741de31257ee580484c64e9f8b91449aa7df22ae38fcb86e50bdadfca0ad23.
//
// Solidity: event Lock(address ethToken, address relayToken, address locker, string recipient, uint256 amount, uint256 appchainIndex)
func (_Escrows *EscrowsFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *EscrowsLock) (event.Subscription, error) {

	logs, sub, err := _Escrows.contract.WatchLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowsLock)
				if err := _Escrows.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0x67741de31257ee580484c64e9f8b91449aa7df22ae38fcb86e50bdadfca0ad23.
//
// Solidity: event Lock(address ethToken, address relayToken, address locker, string recipient, uint256 amount, uint256 appchainIndex)
func (_Escrows *EscrowsFilterer) ParseLock(log types.Log) (*EscrowsLock, error) {
	event := new(EscrowsLock)
	if err := _Escrows.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowsQuickSwapIterator is returned from FilterQuickSwap and is used to iterate over the raw logs and unpacked data for QuickSwap events raised by the Escrows contract.
type EscrowsQuickSwapIterator struct {
	Event *EscrowsQuickSwap // Event containing the contract specifics and raw log

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
func (it *EscrowsQuickSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowsQuickSwap)
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
		it.Event = new(EscrowsQuickSwap)
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
func (it *EscrowsQuickSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowsQuickSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowsQuickSwap represents a QuickSwap event raised by the Escrows contract.
type EscrowsQuickSwap struct {
	DstChainId  string
	DstContract string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterQuickSwap is a free log retrieval operation binding the contract event 0xbf42a9b8a78d1a7612fe5abd3e8bb6a6d68d4a31cc3a0ead24cdad19450fab83.
//
// Solidity: event QuickSwap(string dstChainId, string dstContract)
func (_Escrows *EscrowsFilterer) FilterQuickSwap(opts *bind.FilterOpts) (*EscrowsQuickSwapIterator, error) {

	logs, sub, err := _Escrows.contract.FilterLogs(opts, "QuickSwap")
	if err != nil {
		return nil, err
	}
	return &EscrowsQuickSwapIterator{contract: _Escrows.contract, event: "QuickSwap", logs: logs, sub: sub}, nil
}

// WatchQuickSwap is a free log subscription operation binding the contract event 0xbf42a9b8a78d1a7612fe5abd3e8bb6a6d68d4a31cc3a0ead24cdad19450fab83.
//
// Solidity: event QuickSwap(string dstChainId, string dstContract)
func (_Escrows *EscrowsFilterer) WatchQuickSwap(opts *bind.WatchOpts, sink chan<- *EscrowsQuickSwap) (event.Subscription, error) {

	logs, sub, err := _Escrows.contract.WatchLogs(opts, "QuickSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowsQuickSwap)
				if err := _Escrows.contract.UnpackLog(event, "QuickSwap", log); err != nil {
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

// ParseQuickSwap is a log parse operation binding the contract event 0xbf42a9b8a78d1a7612fe5abd3e8bb6a6d68d4a31cc3a0ead24cdad19450fab83.
//
// Solidity: event QuickSwap(string dstChainId, string dstContract)
func (_Escrows *EscrowsFilterer) ParseQuickSwap(log types.Log) (*EscrowsQuickSwap, error) {
	event := new(EscrowsQuickSwap)
	if err := _Escrows.contract.UnpackLog(event, "QuickSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Escrows contract.
type EscrowsRoleAdminChangedIterator struct {
	Event *EscrowsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EscrowsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowsRoleAdminChanged)
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
		it.Event = new(EscrowsRoleAdminChanged)
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
func (it *EscrowsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowsRoleAdminChanged represents a RoleAdminChanged event raised by the Escrows contract.
type EscrowsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Escrows *EscrowsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EscrowsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Escrows.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EscrowsRoleAdminChangedIterator{contract: _Escrows.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Escrows *EscrowsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EscrowsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Escrows.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowsRoleAdminChanged)
				if err := _Escrows.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Escrows *EscrowsFilterer) ParseRoleAdminChanged(log types.Log) (*EscrowsRoleAdminChanged, error) {
	event := new(EscrowsRoleAdminChanged)
	if err := _Escrows.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Escrows contract.
type EscrowsRoleGrantedIterator struct {
	Event *EscrowsRoleGranted // Event containing the contract specifics and raw log

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
func (it *EscrowsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowsRoleGranted)
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
		it.Event = new(EscrowsRoleGranted)
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
func (it *EscrowsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowsRoleGranted represents a RoleGranted event raised by the Escrows contract.
type EscrowsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Escrows *EscrowsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EscrowsRoleGrantedIterator, error) {

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

	logs, sub, err := _Escrows.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EscrowsRoleGrantedIterator{contract: _Escrows.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Escrows *EscrowsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EscrowsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Escrows.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowsRoleGranted)
				if err := _Escrows.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Escrows *EscrowsFilterer) ParseRoleGranted(log types.Log) (*EscrowsRoleGranted, error) {
	event := new(EscrowsRoleGranted)
	if err := _Escrows.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Escrows contract.
type EscrowsRoleRevokedIterator struct {
	Event *EscrowsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EscrowsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowsRoleRevoked)
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
		it.Event = new(EscrowsRoleRevoked)
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
func (it *EscrowsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowsRoleRevoked represents a RoleRevoked event raised by the Escrows contract.
type EscrowsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Escrows *EscrowsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EscrowsRoleRevokedIterator, error) {

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

	logs, sub, err := _Escrows.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EscrowsRoleRevokedIterator{contract: _Escrows.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Escrows *EscrowsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EscrowsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Escrows.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowsRoleRevoked)
				if err := _Escrows.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Escrows *EscrowsFilterer) ParseRoleRevoked(log types.Log) (*EscrowsRoleRevoked, error) {
	event := new(EscrowsRoleRevoked)
	if err := _Escrows.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowsUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the Escrows contract.
type EscrowsUnlockIterator struct {
	Event *EscrowsUnlock // Event containing the contract specifics and raw log

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
func (it *EscrowsUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowsUnlock)
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
		it.Event = new(EscrowsUnlock)
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
func (it *EscrowsUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowsUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowsUnlock represents a Unlock event raised by the Escrows contract.
type EscrowsUnlock struct {
	EthToken   common.Address
	RelayToken common.Address
	From       common.Address
	Recipient  common.Address
	Amount     *big.Int
	Txid       string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x5800596ed55e41c52e6e763640e29ff0d4c09c47408a512c4fb0974a452c1f0c.
//
// Solidity: event Unlock(address ethToken, address relayToken, address from, address recipient, uint256 amount, string txid)
func (_Escrows *EscrowsFilterer) FilterUnlock(opts *bind.FilterOpts) (*EscrowsUnlockIterator, error) {

	logs, sub, err := _Escrows.contract.FilterLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return &EscrowsUnlockIterator{contract: _Escrows.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x5800596ed55e41c52e6e763640e29ff0d4c09c47408a512c4fb0974a452c1f0c.
//
// Solidity: event Unlock(address ethToken, address relayToken, address from, address recipient, uint256 amount, string txid)
func (_Escrows *EscrowsFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *EscrowsUnlock) (event.Subscription, error) {

	logs, sub, err := _Escrows.contract.WatchLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowsUnlock)
				if err := _Escrows.contract.UnpackLog(event, "Unlock", log); err != nil {
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

// ParseUnlock is a log parse operation binding the contract event 0x5800596ed55e41c52e6e763640e29ff0d4c09c47408a512c4fb0974a452c1f0c.
//
// Solidity: event Unlock(address ethToken, address relayToken, address from, address recipient, uint256 amount, string txid)
func (_Escrows *EscrowsFilterer) ParseUnlock(log types.Log) (*EscrowsUnlock, error) {
	event := new(EscrowsUnlock)
	if err := _Escrows.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
