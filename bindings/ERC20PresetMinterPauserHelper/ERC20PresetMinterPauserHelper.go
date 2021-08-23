// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20PresetMinterPauserHelper

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

// ERC20PresetMinterPauserHelperMetaData contains all meta data concerning the ERC20PresetMinterPauserHelper contract.
var ERC20PresetMinterPauserHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002c7b38038062002c7b833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b5060405250505081818160049080519060200190620001cf92919062000490565b508060059080519060200190620001e892919062000490565b506012600660006101000a81548160ff021916908360ff16021790555050506000600660016101000a81548160ff021916908315150217905550620002466000801b6200023a620002fa60201b60201c565b6200030260201b60201c565b6200029c60405180807f4d494e5445525f524f4c45000000000000000000000000000000000000000000815250600b019050604051809103902062000290620002fa60201b60201c565b6200030260201b60201c565b620002f260405180807f5041555345525f524f4c45000000000000000000000000000000000000000000815250600b0190506040518091039020620002e6620002fa60201b60201c565b6200030260201b60201c565b50506200053f565b600033905090565b6200031482826200031860201b60201c565b5050565b6200034681600080858152602001908152602001600020600001620003bb60201b620020e21790919060201c565b15620003b7576200035c620002fa60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000620003eb836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620003f360201b60201c565b905092915050565b60006200040783836200046d60201b60201c565b6200046257826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905062000467565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004d357805160ff191683800117855562000504565b8280016001018555821562000504579182015b8281111562000503578251825591602001919060010190620004e6565b5b50905062000513919062000517565b5090565b6200053c91905b80821115620005385760008160009055506001016200051e565b5090565b90565b61272c806200054f6000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c806370a08231116100f9578063a457c2d711610097578063d539139311610071578063d539139314610888578063d547741f146108a6578063dd62ed3e146108f4578063e63ab1e91461096c576101a9565b8063a457c2d71461077a578063a9059cbb146107e0578063ca15c87314610846576101a9565b80639010d07c116100d35780639010d07c146105fb57806391d148541461067357806395d89b41146106d9578063a217fddf1461075c576101a9565b806370a082311461054b57806379cc6790146105a35780638456cb59146105f1576101a9565b8063313ce567116101665780633f4ba83a116101405780633f4ba83a146104a357806340c10f19146104ad57806342966c68146104fb5780635c975abb14610529576101a9565b8063313ce567146103cb57806336568abe146103ef578063395093511461043d576101a9565b806306fdde03146101ae578063095ea7b31461023157806318160ddd1461029757806323b872dd146102b5578063248a9ca31461033b5780632f2ff15d1461037d575b600080fd5b6101b661098a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f65780820151818401526020810190506101db565b50505050905090810190601f1680156102235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61027d6004803603604081101561024757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a2c565b604051808215151515815260200191505060405180910390f35b61029f610a4a565b6040518082815260200191505060405180910390f35b610321600480360360608110156102cb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a54565b604051808215151515815260200191505060405180910390f35b6103676004803603602081101561035157600080fd5b8101908080359060200190929190505050610b2d565b6040518082815260200191505060405180910390f35b6103c96004803603604081101561039357600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b4c565b005b6103d3610bd5565b604051808260ff1660ff16815260200191505060405180910390f35b61043b6004803603604081101561040557600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bec565b005b6104896004803603604081101561045357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c85565b604051808215151515815260200191505060405180910390f35b6104ab610d38565b005b6104f9600480360360408110156104c357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ddd565b005b6105276004803603602081101561051157600080fd5b8101908080359060200190929190505050610e86565b005b610531610e9a565b604051808215151515815260200191505060405180910390f35b61058d6004803603602081101561056157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610eb1565b6040518082815260200191505060405180910390f35b6105ef600480360360408110156105b957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610efa565b005b6105f9610f5c565b005b6106316004803603604081101561061157600080fd5b810190808035906020019092919080359060200190929190505050611001565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6106bf6004803603604081101561068957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611032565b604051808215151515815260200191505060405180910390f35b6106e1611063565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610721578082015181840152602081019050610706565b50505050905090810190601f16801561074e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610764611105565b6040518082815260200191505060405180910390f35b6107c66004803603604081101561079057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061110c565b604051808215151515815260200191505060405180910390f35b61082c600480360360408110156107f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506111d9565b604051808215151515815260200191505060405180910390f35b6108726004803603602081101561085c57600080fd5b81019080803590602001909291905050506111f7565b6040518082815260200191505060405180910390f35b61089061121d565b6040518082815260200191505060405180910390f35b6108f2600480360360408110156108bc57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611256565b005b6109566004803603604081101561090a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112df565b6040518082815260200191505060405180910390f35b610974611366565b6040518082815260200191505060405180910390f35b606060048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a225780601f106109f757610100808354040283529160200191610a22565b820191906000526020600020905b815481529060010190602001808311610a0557829003601f168201915b5050505050905090565b6000610a40610a3961139f565b84846113a7565b6001905092915050565b6000600354905090565b6000610a6184848461159e565b610b2284610a6d61139f565b610b1d8560405180606001604052806028815260200161255660289139600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610ad361139f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118639092919063ffffffff16565b6113a7565b600190509392505050565b6000806000838152602001908152602001600020600201549050919050565b610b7260008084815260200190815260200160002060020154610b6d61139f565b611032565b610bc7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180612454602f913960400191505060405180910390fd5b610bd18282611923565b5050565b6000600660009054906101000a900460ff16905090565b610bf461139f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c77576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f81526020018061269e602f913960400191505060405180910390fd5b610c8182826119b6565b5050565b6000610d2e610c9261139f565b84610d298560026000610ca361139f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a4990919063ffffffff16565b6113a7565b6001905092915050565b610d7e60405180807f5041555345525f524f4c45000000000000000000000000000000000000000000815250600b0190506040518091039020610d7961139f565b611032565b610dd3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260398152602001806124a56039913960400191505060405180910390fd5b610ddb611ad1565b565b610e2360405180807f4d494e5445525f524f4c45000000000000000000000000000000000000000000815250600b0190506040518091039020610e1e61139f565b611032565b610e78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603681526020018061257e6036913960400191505060405180910390fd5b610e828282611bda565b5050565b610e97610e9161139f565b82611da3565b50565b6000600660019054906101000a900460ff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000610f39826040518060600160405280602481526020016125b460249139610f2a86610f2561139f565b6112df565b6118639092919063ffffffff16565b9050610f4d83610f4761139f565b836113a7565b610f578383611da3565b505050565b610fa260405180807f5041555345525f524f4c45000000000000000000000000000000000000000000815250600b0190506040518091039020610f9d61139f565b611032565b610ff7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001806126426037913960400191505060405180910390fd5b610fff611f69565b565b600061102a8260008086815260200190815260200160002060000161207390919063ffffffff16565b905092915050565b600061105b8260008086815260200190815260200160002060000161208d90919063ffffffff16565b905092915050565b606060058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110fb5780601f106110d0576101008083540402835291602001916110fb565b820191906000526020600020905b8154815290600101906020018083116110de57829003601f168201915b5050505050905090565b6000801b81565b60006111cf61111961139f565b846111ca85604051806060016040528060258152602001612679602591396002600061114361139f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118639092919063ffffffff16565b6113a7565b6001905092915050565b60006111ed6111e661139f565b848461159e565b6001905092915050565b60006112166000808481526020019081526020016000206000016120bd565b9050919050565b60405180807f4d494e5445525f524f4c45000000000000000000000000000000000000000000815250600b019050604051809103902081565b61127c6000808481526020019081526020016000206002015461127761139f565b611032565b6112d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806125266030913960400191505060405180910390fd5b6112db82826119b6565b5050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60405180807f5041555345525f524f4c45000000000000000000000000000000000000000000815250600b019050604051809103902081565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561142d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602481526020018061261e6024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156114b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124de6022913960400191505060405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611624576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806125f96025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156116aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806124316023913960400191505060405180910390fd5b6116b58383836120d2565b6117218160405180606001604052806026815260200161250060269139600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118639092919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506117b681600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a4990919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290611910576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156118d55780820151818401526020810190506118ba565b50505050905090810190601f1680156119025780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b61194a816000808581526020019081526020016000206000016120e290919063ffffffff16565b156119b25761195761139f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6119dd8160008085815260200190815260200160002060000161211290919063ffffffff16565b15611a45576119ea61139f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600080828401905083811015611ac7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600660019054906101000a900460ff16611b53576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b6000600660016101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611b9761139f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611c7d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b611c89600083836120d2565b611c9e81600354611a4990919063ffffffff16565b600381905550611cf681600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a4990919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611e29576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806125d86021913960400191505060405180910390fd5b611e35826000836120d2565b611ea18160405180606001604052806022815260200161248360229139600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118639092919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611ef98160035461214290919063ffffffff16565b600381905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600660019054906101000a900460ff1615611fec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6001600660016101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861203061139f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b6000612082836000018361218c565b60001c905092915050565b60006120b5836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61220f565b905092915050565b60006120cb82600001612232565b9050919050565b6120dd838383612243565b505050565b600061210a836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6122b1565b905092915050565b600061213a836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612321565b905092915050565b600061218483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611863565b905092915050565b6000818360000180549050116121ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061240f6022913960400191505060405180910390fd5b8260000182815481106121fc57fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b61224e838383612409565b612256610e9a565b156122ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806126cd602a913960400191505060405180910390fd5b505050565b60006122bd838361220f565b61231657826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905061231b565b600090505b92915050565b600080836001016000848152602001908152602001600020549050600081146123fd576000600182039050600060018660000180549050039050600086600001828154811061236c57fe5b906000526020600020015490508087600001848154811061238957fe5b90600052602060002001819055506001830187600101600083815260200190815260200160002081905550866000018054806123c157fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050612403565b60009150505b92915050565b50505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e647345524332303a207472616e7366657220746f20746865207a65726f2061646472657373416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e7445524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332305072657365744d696e7465725061757365723a206d75737420686176652070617573657220726f6c6520746f20756e706175736545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332305072657365744d696e7465725061757365723a206d7573742068617665206d696e74657220726f6c6520746f206d696e7445524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332305072657365744d696e7465725061757365723a206d75737420686176652070617573657220726f6c6520746f20706175736545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c6645524332305061757361626c653a20746f6b656e207472616e73666572207768696c6520706175736564a2646970667358221220b9ef5457fd9ca214139d013589c0363f527cd36d77bf5965995415789c2ce7fc64736f6c63430006040033",
}

// ERC20PresetMinterPauserHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20PresetMinterPauserHelperMetaData.ABI instead.
var ERC20PresetMinterPauserHelperABI = ERC20PresetMinterPauserHelperMetaData.ABI

// ERC20PresetMinterPauserHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20PresetMinterPauserHelperMetaData.Bin instead.
var ERC20PresetMinterPauserHelperBin = ERC20PresetMinterPauserHelperMetaData.Bin

// DeployERC20PresetMinterPauserHelper deploys a new Ethereum contract, binding an instance of ERC20PresetMinterPauserHelper to it.
func DeployERC20PresetMinterPauserHelper(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *ERC20PresetMinterPauserHelper, error) {
	parsed, err := ERC20PresetMinterPauserHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20PresetMinterPauserHelperBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20PresetMinterPauserHelper{ERC20PresetMinterPauserHelperCaller: ERC20PresetMinterPauserHelperCaller{contract: contract}, ERC20PresetMinterPauserHelperTransactor: ERC20PresetMinterPauserHelperTransactor{contract: contract}, ERC20PresetMinterPauserHelperFilterer: ERC20PresetMinterPauserHelperFilterer{contract: contract}}, nil
}

// ERC20PresetMinterPauserHelper is an auto generated Go binding around an Ethereum contract.
type ERC20PresetMinterPauserHelper struct {
	ERC20PresetMinterPauserHelperCaller     // Read-only binding to the contract
	ERC20PresetMinterPauserHelperTransactor // Write-only binding to the contract
	ERC20PresetMinterPauserHelperFilterer   // Log filterer for contract events
}

// ERC20PresetMinterPauserHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PresetMinterPauserHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PresetMinterPauserHelperSession struct {
	Contract     *ERC20PresetMinterPauserHelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ERC20PresetMinterPauserHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PresetMinterPauserHelperCallerSession struct {
	Contract *ERC20PresetMinterPauserHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ERC20PresetMinterPauserHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PresetMinterPauserHelperTransactorSession struct {
	Contract     *ERC20PresetMinterPauserHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ERC20PresetMinterPauserHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PresetMinterPauserHelperRaw struct {
	Contract *ERC20PresetMinterPauserHelper // Generic contract binding to access the raw methods on
}

// ERC20PresetMinterPauserHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserHelperCallerRaw struct {
	Contract *ERC20PresetMinterPauserHelperCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PresetMinterPauserHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserHelperTransactorRaw struct {
	Contract *ERC20PresetMinterPauserHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20PresetMinterPauserHelper creates a new instance of ERC20PresetMinterPauserHelper, bound to a specific deployed contract.
func NewERC20PresetMinterPauserHelper(address common.Address, backend bind.ContractBackend) (*ERC20PresetMinterPauserHelper, error) {
	contract, err := bindERC20PresetMinterPauserHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelper{ERC20PresetMinterPauserHelperCaller: ERC20PresetMinterPauserHelperCaller{contract: contract}, ERC20PresetMinterPauserHelperTransactor: ERC20PresetMinterPauserHelperTransactor{contract: contract}, ERC20PresetMinterPauserHelperFilterer: ERC20PresetMinterPauserHelperFilterer{contract: contract}}, nil
}

// NewERC20PresetMinterPauserHelperCaller creates a new read-only instance of ERC20PresetMinterPauserHelper, bound to a specific deployed contract.
func NewERC20PresetMinterPauserHelperCaller(address common.Address, caller bind.ContractCaller) (*ERC20PresetMinterPauserHelperCaller, error) {
	contract, err := bindERC20PresetMinterPauserHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperCaller{contract: contract}, nil
}

// NewERC20PresetMinterPauserHelperTransactor creates a new write-only instance of ERC20PresetMinterPauserHelper, bound to a specific deployed contract.
func NewERC20PresetMinterPauserHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PresetMinterPauserHelperTransactor, error) {
	contract, err := bindERC20PresetMinterPauserHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperTransactor{contract: contract}, nil
}

// NewERC20PresetMinterPauserHelperFilterer creates a new log filterer instance of ERC20PresetMinterPauserHelper, bound to a specific deployed contract.
func NewERC20PresetMinterPauserHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PresetMinterPauserHelperFilterer, error) {
	contract, err := bindERC20PresetMinterPauserHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperFilterer{contract: contract}, nil
}

// bindERC20PresetMinterPauserHelper binds a generic wrapper to an already deployed contract.
func bindERC20PresetMinterPauserHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PresetMinterPauserHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PresetMinterPauserHelper.Contract.ERC20PresetMinterPauserHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.ERC20PresetMinterPauserHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.ERC20PresetMinterPauserHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PresetMinterPauserHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserHelper.Contract.DEFAULTADMINROLE(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserHelper.Contract.DEFAULTADMINROLE(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) MINTERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserHelper.Contract.MINTERROLE(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) MINTERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserHelper.Contract.MINTERROLE(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserHelper.Contract.PAUSERROLE(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserHelper.Contract.PAUSERROLE(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Allowance(&_ERC20PresetMinterPauserHelper.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Allowance(&_ERC20PresetMinterPauserHelper.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauserHelper.Contract.BalanceOf(&_ERC20PresetMinterPauserHelper.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauserHelper.Contract.BalanceOf(&_ERC20PresetMinterPauserHelper.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Decimals() (uint8, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Decimals(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) Decimals() (uint8, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Decimals(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20PresetMinterPauserHelper.Contract.GetRoleAdmin(&_ERC20PresetMinterPauserHelper.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20PresetMinterPauserHelper.Contract.GetRoleAdmin(&_ERC20PresetMinterPauserHelper.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC20PresetMinterPauserHelper.Contract.GetRoleMember(&_ERC20PresetMinterPauserHelper.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC20PresetMinterPauserHelper.Contract.GetRoleMember(&_ERC20PresetMinterPauserHelper.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC20PresetMinterPauserHelper.Contract.GetRoleMemberCount(&_ERC20PresetMinterPauserHelper.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC20PresetMinterPauserHelper.Contract.GetRoleMemberCount(&_ERC20PresetMinterPauserHelper.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20PresetMinterPauserHelper.Contract.HasRole(&_ERC20PresetMinterPauserHelper.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20PresetMinterPauserHelper.Contract.HasRole(&_ERC20PresetMinterPauserHelper.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Name() (string, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Name(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) Name() (string, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Name(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Paused() (bool, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Paused(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) Paused() (bool, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Paused(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Symbol() (string, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Symbol(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) Symbol() (string, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Symbol(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserHelper.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) TotalSupply() (*big.Int, error) {
	return _ERC20PresetMinterPauserHelper.Contract.TotalSupply(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20PresetMinterPauserHelper.Contract.TotalSupply(&_ERC20PresetMinterPauserHelper.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Approve(&_ERC20PresetMinterPauserHelper.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Approve(&_ERC20PresetMinterPauserHelper.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Burn(&_ERC20PresetMinterPauserHelper.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Burn(&_ERC20PresetMinterPauserHelper.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.BurnFrom(&_ERC20PresetMinterPauserHelper.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.BurnFrom(&_ERC20PresetMinterPauserHelper.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.DecreaseAllowance(&_ERC20PresetMinterPauserHelper.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.DecreaseAllowance(&_ERC20PresetMinterPauserHelper.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.GrantRole(&_ERC20PresetMinterPauserHelper.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.GrantRole(&_ERC20PresetMinterPauserHelper.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.IncreaseAllowance(&_ERC20PresetMinterPauserHelper.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.IncreaseAllowance(&_ERC20PresetMinterPauserHelper.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Mint(&_ERC20PresetMinterPauserHelper.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Mint(&_ERC20PresetMinterPauserHelper.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Pause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Pause(&_ERC20PresetMinterPauserHelper.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Pause(&_ERC20PresetMinterPauserHelper.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.RenounceRole(&_ERC20PresetMinterPauserHelper.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.RenounceRole(&_ERC20PresetMinterPauserHelper.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.RevokeRole(&_ERC20PresetMinterPauserHelper.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.RevokeRole(&_ERC20PresetMinterPauserHelper.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Transfer(&_ERC20PresetMinterPauserHelper.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Transfer(&_ERC20PresetMinterPauserHelper.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.TransferFrom(&_ERC20PresetMinterPauserHelper.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.TransferFrom(&_ERC20PresetMinterPauserHelper.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperSession) Unpause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Unpause(&_ERC20PresetMinterPauserHelper.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauserHelper.Contract.Unpause(&_ERC20PresetMinterPauserHelper.TransactOpts)
}

// ERC20PresetMinterPauserHelperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperApprovalIterator struct {
	Event *ERC20PresetMinterPauserHelperApproval // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserHelperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserHelperApproval)
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
		it.Event = new(ERC20PresetMinterPauserHelperApproval)
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
func (it *ERC20PresetMinterPauserHelperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserHelperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserHelperApproval represents a Approval event raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20PresetMinterPauserHelperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperApprovalIterator{contract: _ERC20PresetMinterPauserHelper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserHelperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserHelperApproval)
				if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) ParseApproval(log types.Log) (*ERC20PresetMinterPauserHelperApproval, error) {
	event := new(ERC20PresetMinterPauserHelperApproval)
	if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserHelperPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperPausedIterator struct {
	Event *ERC20PresetMinterPauserHelperPaused // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserHelperPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserHelperPaused)
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
		it.Event = new(ERC20PresetMinterPauserHelperPaused)
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
func (it *ERC20PresetMinterPauserHelperPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserHelperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserHelperPaused represents a Paused event raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20PresetMinterPauserHelperPausedIterator, error) {

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperPausedIterator{contract: _ERC20PresetMinterPauserHelper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserHelperPaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserHelperPaused)
				if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) ParsePaused(log types.Log) (*ERC20PresetMinterPauserHelperPaused, error) {
	event := new(ERC20PresetMinterPauserHelperPaused)
	if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserHelperRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperRoleGrantedIterator struct {
	Event *ERC20PresetMinterPauserHelperRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserHelperRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserHelperRoleGranted)
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
		it.Event = new(ERC20PresetMinterPauserHelperRoleGranted)
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
func (it *ERC20PresetMinterPauserHelperRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserHelperRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserHelperRoleGranted represents a RoleGranted event raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20PresetMinterPauserHelperRoleGrantedIterator, error) {

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

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperRoleGrantedIterator{contract: _ERC20PresetMinterPauserHelper.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserHelperRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserHelperRoleGranted)
				if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) ParseRoleGranted(log types.Log) (*ERC20PresetMinterPauserHelperRoleGranted, error) {
	event := new(ERC20PresetMinterPauserHelperRoleGranted)
	if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserHelperRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperRoleRevokedIterator struct {
	Event *ERC20PresetMinterPauserHelperRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserHelperRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserHelperRoleRevoked)
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
		it.Event = new(ERC20PresetMinterPauserHelperRoleRevoked)
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
func (it *ERC20PresetMinterPauserHelperRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserHelperRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserHelperRoleRevoked represents a RoleRevoked event raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20PresetMinterPauserHelperRoleRevokedIterator, error) {

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

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperRoleRevokedIterator{contract: _ERC20PresetMinterPauserHelper.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserHelperRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserHelperRoleRevoked)
				if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) ParseRoleRevoked(log types.Log) (*ERC20PresetMinterPauserHelperRoleRevoked, error) {
	event := new(ERC20PresetMinterPauserHelperRoleRevoked)
	if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserHelperTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperTransferIterator struct {
	Event *ERC20PresetMinterPauserHelperTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserHelperTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserHelperTransfer)
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
		it.Event = new(ERC20PresetMinterPauserHelperTransfer)
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
func (it *ERC20PresetMinterPauserHelperTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserHelperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserHelperTransfer represents a Transfer event raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20PresetMinterPauserHelperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperTransferIterator{contract: _ERC20PresetMinterPauserHelper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserHelperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserHelperTransfer)
				if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) ParseTransfer(log types.Log) (*ERC20PresetMinterPauserHelperTransfer, error) {
	event := new(ERC20PresetMinterPauserHelperTransfer)
	if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserHelperUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperUnpausedIterator struct {
	Event *ERC20PresetMinterPauserHelperUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserHelperUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserHelperUnpaused)
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
		it.Event = new(ERC20PresetMinterPauserHelperUnpaused)
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
func (it *ERC20PresetMinterPauserHelperUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserHelperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserHelperUnpaused represents a Unpaused event raised by the ERC20PresetMinterPauserHelper contract.
type ERC20PresetMinterPauserHelperUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20PresetMinterPauserHelperUnpausedIterator, error) {

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserHelperUnpausedIterator{contract: _ERC20PresetMinterPauserHelper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserHelperUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PresetMinterPauserHelper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserHelperUnpaused)
				if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PresetMinterPauserHelper *ERC20PresetMinterPauserHelperFilterer) ParseUnpaused(log types.Log) (*ERC20PresetMinterPauserHelperUnpaused, error) {
	event := new(ERC20PresetMinterPauserHelperUnpaused)
	if err := _ERC20PresetMinterPauserHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
