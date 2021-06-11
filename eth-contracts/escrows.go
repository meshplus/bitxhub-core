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


// EscrowsABI is the input ABI used to generate the binding from.
const EscrowsABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_relayers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appchainIndex\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayTokenAddr\",\"type\":\"address\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"ethTokenAddrs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"relayTokenAddrs\",\"type\":\"address[]\"}],\"name\":\"addSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appchainIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"index2Height\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"}],\"name\":\"removeSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_txid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_relayIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
	"4bbc170a": "lock(address,uint256,address)",
	"e2095ab4": "lockAmount(address,address)",
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
var EscrowsBin = "0x6080604052600060065560006007553480156200001b57600080fd5b5060405162001e4238038062001e428339810160408190526200003e91620001ca565b6200004b6000336200009e565b60005b815181101562000096576200008d6b52454c415945525f524f4c4560a01b8383815181106200007957fe5b60200260200101516200009e60201b60201c565b6001016200004e565b5050620002a6565b620000aa8282620000ae565b5050565b600082815260208181526040909120620000d391839062000b1b62000127821b17901c565b15620000aa57620000e362000147565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60006200013e836001600160a01b0384166200014b565b90505b92915050565b3390565b60006200015983836200019a565b620001915750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000141565b50600062000141565b60009081526001919091016020526040902054151590565b80516001600160a01b03811681146200014157600080fd5b60006020808385031215620001dd578182fd5b82516001600160401b0380821115620001f4578384fd5b818501915085601f83011262000208578384fd5b81518181111562000217578485fd5b8381029150620002298483016200027f565b8181528481019084860184860187018a101562000244578788fd5b8795505b8386101562000272576200025d8a82620001b2565b83526001959095019491860191860162000248565b5098975050505050505050565b6040518181016001600160401b03811182821017156200029e57600080fd5b604052919050565b611b8c80620002b66000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c8063967145af116100c3578063d547741f1161007c578063d547741f1461029b578063e1c7392a146102ae578063e2095ab4146102b6578063e2769cfa146102c9578063ef2abc1f146102dc578063fced6ada146102ef5761014d565b8063967145af1461023f57806396f963b414610252578063a217fddf14610265578063a26c16a41461026d578063ab1494de14610275578063ca15c873146102885761014d565b80634bbc170a116101155780634bbc170a146101d65780637010584c146101e957806382815874146101fc5780639010d07c1461020457806391d1485414610217578063926d7d7f146102375761014d565b80630daff62114610152578063248a9ca3146101675780632a4f1621146101905780632f2ff15d146101b057806336568abe146101c3575b600080fd5b610165610160366004611402565b6102f7565b005b61017a6101753660046114b6565b61032b565b6040516101879190611708565b60405180910390f35b6101a361019e366004611273565b610340565b6040516101879190611623565b6101656101be3660046114ce565b61035b565b6101656101d13660046114ce565b6103a8565b6101656101e43660046113c1565b6103ea565b6101656101f736600461128e565b61050a565b61017a610597565b6101a36102123660046114fd565b61059d565b61022a6102253660046114ce565b6105be565b60405161018791906116fd565b61017a6105d6565b61022a61024d36600461151e565b6105e9565b61017a6102603660046114b6565b610609565b61017a61061b565b61017a610620565b610165610283366004611435565b610626565b61017a6102963660046114b6565b610690565b6101656102a93660046114ce565b6106a7565b6101656106e1565b61017a6102c436600461128e565b610714565b6101656102d7366004611273565b610731565b6101656102ea3660046112c2565b6107b6565b61017a610b0b565b60005b81518110156103275761031f82828151811061031257fe5b6020026020010151610731565b6001016102fa565b5050565b60009081526020819052604090206002015490565b6001602052600090815260409020546001600160a01b031681565b60008281526020819052604090206002015461037990610225610b30565b61039e5760405162461bcd60e51b8152600401610395906117b1565b60405180910390fd5b6103278282610b34565b6103b0610b30565b6001600160a01b0316816001600160a01b0316146103e05760405162461bcd60e51b815260040161039590611a7c565b6103278282610b9d565b6001600160a01b038084166000908152600160205260409020548491166104235760405162461bcd60e51b8152600401610395906119c4565b6001600160a01b03841660009081526002602090815260408083203384529091529020546104519084610c06565b6001600160a01b038516600081815260026020908152604080832033808552925290912092909255610484913086610c2b565b600654610492906001610c06565b600681815560009182526003602090815260408084204390556001600160a01b03808916855260019092529283902054915492517f5470ff67bafe441e81f97ab58498aee8999bf604f37158490e83c01753cd6e21936104fc9389931691339188918a9190611637565b60405180910390a150505050565b6105156000336105be565b6105315760405162461bcd60e51b815260040161039590611742565b6001600160a01b0382811660009081526001602052604090205416156105695760405162461bcd60e51b81526004016103959061198d565b6001600160a01b03918216600090815260016020526040902080546001600160a01b03191691909216179055565b60065481565b60008281526020819052604081206105b59083610c89565b90505b92915050565b60008281526020819052604081206105b59083610c95565b6b52454c415945525f524f4c4560a01b81565b805160208183018101805160048252928201919093012091525460ff1681565b60036020526000908152604090205481565b600081565b60075481565b80518251146106475760405162461bcd60e51b815260040161039590611800565b60005b825181101561068b5761068383828151811061066257fe5b602002602001015183838151811061067657fe5b602002602001015161050a565b60010161064a565b505050565b60008181526020819052604081206105b890610caa565b6000828152602081905260409020600201546106c590610225610b30565b6103e05760405162461bcd60e51b8152600401610395906118e5565b6106ec6000336105be565b6107085760405162461bcd60e51b815260040161039590611742565b60006006819055600755565b600260209081526000928352604080842090915290825290205481565b61073c6000336105be565b6107585760405162461bcd60e51b815260040161039590611742565b6001600160a01b038181166000908152600160205260409020541661078f5760405162461bcd60e51b8152600401610395906118b8565b6001600160a01b0316600090815260016020526040902080546001600160a01b0319169055565b6001600160a01b038088166000908152600160205260409020548891166107ef5760405162461bcd60e51b8152600401610395906119c4565b61080568504945525f524f4c4560b81b336105be565b6108215760405162461bcd60e51b815260040161039590611935565b8360048160405161083291906115d6565b9081526040519081900360200190205460ff16156108625760405162461bcd60e51b815260040161039590611893565b60018403600754146108865760405162461bcd60e51b815260040161039590611964565b60006108a06b52454c415945525f524f4c4560a01b610690565b90506000600260036000198401048301810104905080855110156108c5575050610b00565b60008b8b8b8b8b6040516020016108e095949392919061157d565b60405160208183030381529060405280519060200120905060005b865181101561096e57600061092b61091284610cb5565b89848151811061091e57fe5b6020026020010151610ce5565b90506109466b52454c415945525f524f4c4560a01b826105be565b156109655760008381526005602052604090206109639082610b1b565b505b506001016108fb565b506000818152600560205260409020829061098890610caa565b10156109a65760405162461bcd60e51b815260040161039590611830565b60016004896040516109b891906115d6565b908152604051908190036020019020805491151560ff199092169190911790556007546109e6906001610c06565b6007556001600160a01b03808d166000908152600260209081526040808320938e1683529290522054610a19908a610dc7565b600260008e6001600160a01b03166001600160a01b0316815260200190815260200160002060008c6001600160a01b03166001600160a01b0316815260200190815260200160002081905550610a838a8a8e6001600160a01b0316610e099092919063ffffffff16565b7f5800596ed55e41c52e6e763640e29ff0d4c09c47408a512c4fb0974a452c1f0c8c600160008f6001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03168d8d8d8d604051610af496959493929190611671565b60405180910390a15050505b505050505050505050565b68504945525f524f4c4560b81b81565b60006105b5836001600160a01b038416610e28565b3390565b6000828152602081905260409020610b4c9082610b1b565b1561032757610b59610b30565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152602081905260409020610bb59082610e72565b1561032757610bc2610b30565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6000828201838110156105b55760405162461bcd60e51b81526004016103959061185c565b610c83846323b872dd60e01b858585604051602401610c4c939291906116c0565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610e87565b50505050565b60006105b58383610f16565b60006105b5836001600160a01b038416610f5b565b60006105b882610f73565b600081604051602001610cc891906115f2565b604051602081830303815290604052805190602001209050919050565b60008151604114610cf8575060006105b8565b60208201516040830151606084015160001a601b017f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610d4157600093505050506105b8565b8060ff16601b14158015610d5957508060ff16601c14155b15610d6a57600093505050506105b8565b600060018783868660405160008152602001604052604051610d8f9493929190611711565b6020604051602081039080840390855afa158015610db1573d6000803e3d6000fd5b5050604051601f19015198975050505050505050565b60006105b583836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610f77565b61068b8363a9059cbb60e01b8484604051602401610c4c9291906116e4565b6000610e348383610f5b565b610e6a575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105b8565b5060006105b8565b60006105b5836001600160a01b038416610fa3565b6060610edc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166110699092919063ffffffff16565b80519091501561068b5780806020019051810190610efa9190611496565b61068b5760405162461bcd60e51b815260040161039590611a32565b81546000908210610f395760405162461bcd60e51b81526004016103959061176f565b826000018281548110610f4857fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b60008184841115610f9b5760405162461bcd60e51b8152600401610395919061172f565b505050900390565b6000818152600183016020526040812054801561105f5783546000198083019190810190600090879083908110610fd657fe5b9060005260206000200154905080876000018481548110610ff357fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061102357fe5b600190038181906000526020600020016000905590558660010160008781526020019081526020016000206000905560019450505050506105b8565b60009150506105b8565b60606110788484600085611080565b949350505050565b606061108b85611144565b6110a75760405162461bcd60e51b8152600401610395906119fb565b60006060866001600160a01b031685876040516110c491906115d6565b60006040518083038185875af1925050503d8060008114611101576040519150601f19603f3d011682016040523d82523d6000602084013e611106565b606091505b5091509150811561111a5791506110789050565b80511561112a5780518082602001fd5b8360405162461bcd60e51b8152600401610395919061172f565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590611078575050151592915050565b80356001600160a01b03811681146105b857600080fd5b600082601f8301126111a4578081fd5b81356111b76111b282611af2565b611acb565b8181529150602080830190848101818402860182018710156111d857600080fd5b60005b848110156111ff576111ed888361117d565b845292820192908201906001016111db565b505050505092915050565b600082601f83011261121a578081fd5b813567ffffffffffffffff811115611230578182fd5b611243601f8201601f1916602001611acb565b915080825283602082850101111561125a57600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215611284578081fd5b6105b5838361117d565b600080604083850312156112a0578081fd5b6112aa848461117d565b91506112b9846020850161117d565b90509250929050565b600080600080600080600060e0888a0312156112dc578283fd5b87356112e781611b3e565b965060208801356112f781611b3e565b9550604088013561130781611b3e565b945060608801359350608088013567ffffffffffffffff8082111561132a578485fd5b6113368b838c0161120a565b945060a08a0135935060c08a0135915080821115611352578283fd5b508801601f81018a13611363578182fd5b80356113716111b282611af2565b818152602080820191908401855b848110156113ac576113978f6020843589010161120a565b8452602093840193919091019060010161137f565b50508094505050505092959891949750929550565b6000806000606084860312156113d5578283fd5b83356113e081611b3e565b92506020840135915060408401356113f781611b3e565b809150509250925092565b600060208284031215611413578081fd5b813567ffffffffffffffff811115611429578182fd5b61107884828501611194565b60008060408385031215611447578182fd5b823567ffffffffffffffff8082111561145e578384fd5b61146a86838701611194565b9350602085013591508082111561147f578283fd5b5061148c85828601611194565b9150509250929050565b6000602082840312156114a7578081fd5b815180151581146105b5578182fd5b6000602082840312156114c7578081fd5b5035919050565b600080604083850312156114e0578182fd5b8235915060208301356114f281611b3e565b809150509250929050565b6000806040838503121561150f578182fd5b50508035926020909101359150565b60006020828403121561152f578081fd5b813567ffffffffffffffff811115611545578182fd5b6110788482850161120a565b60008151808452611569816020860160208601611b12565b601f01601f19169290920160200192915050565b60006bffffffffffffffffffffffff19808860601b168352808760601b166014840152808660601b1660288401525083603c83015282516115c581605c850160208701611b12565b91909101605c019695505050505050565b600082516115e8818460208701611b12565b9190910192915050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b6001600160a01b0391909116815260200190565b6001600160a01b0396871681529486166020860152928516604085015293166060830152608082019290925260a081019190915260c00190565b6001600160a01b03878116825286811660208301528581166040830152841660608201526080810183905260c060a082018190526000906116b490830184611551565b98975050505050505050565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b90815260200190565b93845260ff9290921660208401526040830152606082015260800190565b6000602082526105b56020830184611551565b60208082526013908201527231b0b63632b91034b9903737ba1030b236b4b760691b604082015260600190565b60208082526022908201527f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e604082015261647360f01b606082015260800190565b6020808252602f908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201526e0818591b5a5b881d1bc819dc985b9d608a1b606082015260800190565b6020808252601690820152750a8ded6cadc40d8cadccee8d040dcdee840dac2e8c6d60531b604082015260600190565b6020808252601290820152711cda59db985d1d5c995cc81a5b9d985b1a5960721b604082015260600190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252600b908201526a1d1e081d5b9b1bd8dad95960aa1b604082015260600190565b602080825260139082015272151bdad95b881b9bdd0814dd5c1c1bdc9d1959606a1b604082015260600190565b60208082526030908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201526f2061646d696e20746f207265766f6b6560801b606082015260800190565b60208082526015908201527431b0b63632b91034b9903737ba1031b937b9b9b2b960591b604082015260600190565b6020808252600f908201526e0d2dcc8caf040dcdee840dac2e8c6d608b1b604082015260600190565b60208082526017908201527f546f6b656e20616c726561647920537570706f72746564000000000000000000604082015260600190565b60208082526017908201527f4c6f636b3a3a4e6f7420537570706f727420546f6b656e000000000000000000604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b6020808252602f908201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560408201526e103937b632b9903337b91039b2b63360891b606082015260800190565b60405181810167ffffffffffffffff81118282101715611aea57600080fd5b604052919050565b600067ffffffffffffffff821115611b08578081fd5b5060209081020190565b60005b83811015611b2d578181015183820152602001611b15565b83811115610c835750506000910152565b6001600160a01b0381168114611b5357600080fd5b5056fea2646970667358221220d2334e801909e9d192a97dee8aae56305d1a775bd53bb4ed6567bbc9b725306264736f6c634300060c0033"

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

// LockAmount is a free data retrieval call binding the contract method 0xe2095ab4.
//
// Solidity: function lockAmount(address , address ) view returns(uint256)
func (_Escrows *EscrowsCaller) LockAmount(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrows.contract.Call(opts, &out, "lockAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockAmount is a free data retrieval call binding the contract method 0xe2095ab4.
//
// Solidity: function lockAmount(address , address ) view returns(uint256)
func (_Escrows *EscrowsSession) LockAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Escrows.Contract.LockAmount(&_Escrows.CallOpts, arg0, arg1)
}

// LockAmount is a free data retrieval call binding the contract method 0xe2095ab4.
//
// Solidity: function lockAmount(address , address ) view returns(uint256)
func (_Escrows *EscrowsCallerSession) LockAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Escrows.Contract.LockAmount(&_Escrows.CallOpts, arg0, arg1)
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

// Lock is a paid mutator transaction binding the contract method 0x4bbc170a.
//
// Solidity: function lock(address token, uint256 amount, address recipient) returns()
func (_Escrows *EscrowsTransactor) Lock(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Escrows.contract.Transact(opts, "lock", token, amount, recipient)
}

// Lock is a paid mutator transaction binding the contract method 0x4bbc170a.
//
// Solidity: function lock(address token, uint256 amount, address recipient) returns()
func (_Escrows *EscrowsSession) Lock(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.Lock(&_Escrows.TransactOpts, token, amount, recipient)
}

// Lock is a paid mutator transaction binding the contract method 0x4bbc170a.
//
// Solidity: function lock(address token, uint256 amount, address recipient) returns()
func (_Escrows *EscrowsTransactorSession) Lock(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Escrows.Contract.Lock(&_Escrows.TransactOpts, token, amount, recipient)
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
	Recipient     common.Address
	Amount        *big.Int
	AppchainIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x5470ff67bafe441e81f97ab58498aee8999bf604f37158490e83c01753cd6e21.
//
// Solidity: event Lock(address ethToken, address relayToken, address locker, address recipient, uint256 amount, uint256 appchainIndex)
func (_Escrows *EscrowsFilterer) FilterLock(opts *bind.FilterOpts) (*EscrowsLockIterator, error) {

	logs, sub, err := _Escrows.contract.FilterLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return &EscrowsLockIterator{contract: _Escrows.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x5470ff67bafe441e81f97ab58498aee8999bf604f37158490e83c01753cd6e21.
//
// Solidity: event Lock(address ethToken, address relayToken, address locker, address recipient, uint256 amount, uint256 appchainIndex)
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

// ParseLock is a log parse operation binding the contract event 0x5470ff67bafe441e81f97ab58498aee8999bf604f37158490e83c01753cd6e21.
//
// Solidity: event Lock(address ethToken, address relayToken, address locker, address recipient, uint256 amount, uint256 appchainIndex)
func (_Escrows *EscrowsFilterer) ParseLock(log types.Log) (*EscrowsLock, error) {
	event := new(EscrowsLock)
	if err := _Escrows.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
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
	return event, nil
}
