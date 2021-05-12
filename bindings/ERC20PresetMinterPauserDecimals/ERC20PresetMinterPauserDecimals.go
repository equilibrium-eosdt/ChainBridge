// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20PresetMinterPauserDecimals

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

// ERC20PresetMinterPauserDecimalsABI is the input ABI used to generate the binding from.
const ERC20PresetMinterPauserDecimalsABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20PresetMinterPauserDecimalsBin is the compiled bytecode used for deploying new contracts.
var ERC20PresetMinterPauserDecimalsBin = "0x60806040523480156200001157600080fd5b5060405162002cb838038062002cb8833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b506040526020018051906020019092919050505082828160049080519060200190620001dc929190620004cd565b508060059080519060200190620001f5929190620004cd565b506012600660006101000a81548160ff021916908360ff16021790555050506000600660016101000a81548160ff021916908315150217905550620002536000801b620002476200031960201b60201c565b6200032160201b60201c565b620002a960405180807f4d494e5445525f524f4c45000000000000000000000000000000000000000000815250600b01905060405180910390206200029d6200031960201b60201c565b6200032160201b60201c565b620002ff60405180807f5041555345525f524f4c45000000000000000000000000000000000000000000815250600b0190506040518091039020620002f36200031960201b60201c565b6200032160201b60201c565b62000310816200033760201b60201c565b5050506200057c565b600033905090565b6200033382826200035560201b60201c565b5050565b80600660006101000a81548160ff021916908360ff16021790555050565b6200038381600080858152602001908152602001600020600001620003f860201b620020e21790919060201c565b15620003f457620003996200031960201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600062000428836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200043060201b60201c565b905092915050565b6000620004448383620004aa60201b60201c565b6200049f578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620004a4565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200051057805160ff191683800117855562000541565b8280016001018555821562000541579182015b828111156200054057825182559160200191906001019062000523565b5b50905062000550919062000554565b5090565b6200057991905b80821115620005755760008160009055506001016200055b565b5090565b90565b61272c806200058c6000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c806370a08231116100f9578063a457c2d711610097578063d539139311610071578063d539139314610888578063d547741f146108a6578063dd62ed3e146108f4578063e63ab1e91461096c576101a9565b8063a457c2d71461077a578063a9059cbb146107e0578063ca15c87314610846576101a9565b80639010d07c116100d35780639010d07c146105fb57806391d148541461067357806395d89b41146106d9578063a217fddf1461075c576101a9565b806370a082311461054b57806379cc6790146105a35780638456cb59146105f1576101a9565b8063313ce567116101665780633f4ba83a116101405780633f4ba83a146104a357806340c10f19146104ad57806342966c68146104fb5780635c975abb14610529576101a9565b8063313ce567146103cb57806336568abe146103ef578063395093511461043d576101a9565b806306fdde03146101ae578063095ea7b31461023157806318160ddd1461029757806323b872dd146102b5578063248a9ca31461033b5780632f2ff15d1461037d575b600080fd5b6101b661098a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f65780820151818401526020810190506101db565b50505050905090810190601f1680156102235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61027d6004803603604081101561024757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a2c565b604051808215151515815260200191505060405180910390f35b61029f610a4a565b6040518082815260200191505060405180910390f35b610321600480360360608110156102cb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a54565b604051808215151515815260200191505060405180910390f35b6103676004803603602081101561035157600080fd5b8101908080359060200190929190505050610b2d565b6040518082815260200191505060405180910390f35b6103c96004803603604081101561039357600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b4c565b005b6103d3610bd5565b604051808260ff1660ff16815260200191505060405180910390f35b61043b6004803603604081101561040557600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bec565b005b6104896004803603604081101561045357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c85565b604051808215151515815260200191505060405180910390f35b6104ab610d38565b005b6104f9600480360360408110156104c357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ddd565b005b6105276004803603602081101561051157600080fd5b8101908080359060200190929190505050610e86565b005b610531610e9a565b604051808215151515815260200191505060405180910390f35b61058d6004803603602081101561056157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610eb1565b6040518082815260200191505060405180910390f35b6105ef600480360360408110156105b957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610efa565b005b6105f9610f5c565b005b6106316004803603604081101561061157600080fd5b810190808035906020019092919080359060200190929190505050611001565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6106bf6004803603604081101561068957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611032565b604051808215151515815260200191505060405180910390f35b6106e1611063565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610721578082015181840152602081019050610706565b50505050905090810190601f16801561074e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610764611105565b6040518082815260200191505060405180910390f35b6107c66004803603604081101561079057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061110c565b604051808215151515815260200191505060405180910390f35b61082c600480360360408110156107f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506111d9565b604051808215151515815260200191505060405180910390f35b6108726004803603602081101561085c57600080fd5b81019080803590602001909291905050506111f7565b6040518082815260200191505060405180910390f35b61089061121d565b6040518082815260200191505060405180910390f35b6108f2600480360360408110156108bc57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611256565b005b6109566004803603604081101561090a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112df565b6040518082815260200191505060405180910390f35b610974611366565b6040518082815260200191505060405180910390f35b606060048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a225780601f106109f757610100808354040283529160200191610a22565b820191906000526020600020905b815481529060010190602001808311610a0557829003601f168201915b5050505050905090565b6000610a40610a3961139f565b84846113a7565b6001905092915050565b6000600354905090565b6000610a6184848461159e565b610b2284610a6d61139f565b610b1d8560405180606001604052806028815260200161255660289139600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610ad361139f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118639092919063ffffffff16565b6113a7565b600190509392505050565b6000806000838152602001908152602001600020600201549050919050565b610b7260008084815260200190815260200160002060020154610b6d61139f565b611032565b610bc7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180612454602f913960400191505060405180910390fd5b610bd18282611923565b5050565b6000600660009054906101000a900460ff16905090565b610bf461139f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c77576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f81526020018061269e602f913960400191505060405180910390fd5b610c8182826119b6565b5050565b6000610d2e610c9261139f565b84610d298560026000610ca361139f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a4990919063ffffffff16565b6113a7565b6001905092915050565b610d7e60405180807f5041555345525f524f4c45000000000000000000000000000000000000000000815250600b0190506040518091039020610d7961139f565b611032565b610dd3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260398152602001806124a56039913960400191505060405180910390fd5b610ddb611ad1565b565b610e2360405180807f4d494e5445525f524f4c45000000000000000000000000000000000000000000815250600b0190506040518091039020610e1e61139f565b611032565b610e78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603681526020018061257e6036913960400191505060405180910390fd5b610e828282611bda565b5050565b610e97610e9161139f565b82611da3565b50565b6000600660019054906101000a900460ff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000610f39826040518060600160405280602481526020016125b460249139610f2a86610f2561139f565b6112df565b6118639092919063ffffffff16565b9050610f4d83610f4761139f565b836113a7565b610f578383611da3565b505050565b610fa260405180807f5041555345525f524f4c45000000000000000000000000000000000000000000815250600b0190506040518091039020610f9d61139f565b611032565b610ff7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001806126426037913960400191505060405180910390fd5b610fff611f69565b565b600061102a8260008086815260200190815260200160002060000161207390919063ffffffff16565b905092915050565b600061105b8260008086815260200190815260200160002060000161208d90919063ffffffff16565b905092915050565b606060058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110fb5780601f106110d0576101008083540402835291602001916110fb565b820191906000526020600020905b8154815290600101906020018083116110de57829003601f168201915b5050505050905090565b6000801b81565b60006111cf61111961139f565b846111ca85604051806060016040528060258152602001612679602591396002600061114361139f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118639092919063ffffffff16565b6113a7565b6001905092915050565b60006111ed6111e661139f565b848461159e565b6001905092915050565b60006112166000808481526020019081526020016000206000016120bd565b9050919050565b60405180807f4d494e5445525f524f4c45000000000000000000000000000000000000000000815250600b019050604051809103902081565b61127c6000808481526020019081526020016000206002015461127761139f565b611032565b6112d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806125266030913960400191505060405180910390fd5b6112db82826119b6565b5050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60405180807f5041555345525f524f4c45000000000000000000000000000000000000000000815250600b019050604051809103902081565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561142d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602481526020018061261e6024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156114b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124de6022913960400191505060405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611624576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806125f96025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156116aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806124316023913960400191505060405180910390fd5b6116b58383836120d2565b6117218160405180606001604052806026815260200161250060269139600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118639092919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506117b681600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a4990919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290611910576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156118d55780820151818401526020810190506118ba565b50505050905090810190601f1680156119025780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b61194a816000808581526020019081526020016000206000016120e290919063ffffffff16565b156119b25761195761139f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6119dd8160008085815260200190815260200160002060000161211290919063ffffffff16565b15611a45576119ea61139f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600080828401905083811015611ac7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600660019054906101000a900460ff16611b53576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b6000600660016101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611b9761139f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611c7d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b611c89600083836120d2565b611c9e81600354611a4990919063ffffffff16565b600381905550611cf681600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a4990919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611e29576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806125d86021913960400191505060405180910390fd5b611e35826000836120d2565b611ea18160405180606001604052806022815260200161248360229139600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546118639092919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611ef98160035461214290919063ffffffff16565b600381905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600660019054906101000a900460ff1615611fec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6001600660016101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861203061139f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b6000612082836000018361218c565b60001c905092915050565b60006120b5836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61220f565b905092915050565b60006120cb82600001612232565b9050919050565b6120dd838383612243565b505050565b600061210a836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6122b1565b905092915050565b600061213a836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612321565b905092915050565b600061218483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611863565b905092915050565b6000818360000180549050116121ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061240f6022913960400191505060405180910390fd5b8260000182815481106121fc57fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b61224e838383612409565b612256610e9a565b156122ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806126cd602a913960400191505060405180910390fd5b505050565b60006122bd838361220f565b61231657826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905061231b565b600090505b92915050565b600080836001016000848152602001908152602001600020549050600081146123fd576000600182039050600060018660000180549050039050600086600001828154811061236c57fe5b906000526020600020015490508087600001848154811061238957fe5b90600052602060002001819055506001830187600101600083815260200190815260200160002081905550866000018054806123c157fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050612403565b60009150505b92915050565b50505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e647345524332303a207472616e7366657220746f20746865207a65726f2061646472657373416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e7445524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332305072657365744d696e7465725061757365723a206d75737420686176652070617573657220726f6c6520746f20756e706175736545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332305072657365744d696e7465725061757365723a206d7573742068617665206d696e74657220726f6c6520746f206d696e7445524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332305072657365744d696e7465725061757365723a206d75737420686176652070617573657220726f6c6520746f20706175736545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c6645524332305061757361626c653a20746f6b656e207472616e73666572207768696c6520706175736564a2646970667358221220b9595201a47f1707948570064910602f0b7e08b39fcc15c938414ba2ea6e75a064736f6c63430006040033"

// DeployERC20PresetMinterPauserDecimals deploys a new Ethereum contract, binding an instance of ERC20PresetMinterPauserDecimals to it.
func DeployERC20PresetMinterPauserDecimals(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *ERC20PresetMinterPauserDecimals, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PresetMinterPauserDecimalsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20PresetMinterPauserDecimalsBin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20PresetMinterPauserDecimals{ERC20PresetMinterPauserDecimalsCaller: ERC20PresetMinterPauserDecimalsCaller{contract: contract}, ERC20PresetMinterPauserDecimalsTransactor: ERC20PresetMinterPauserDecimalsTransactor{contract: contract}, ERC20PresetMinterPauserDecimalsFilterer: ERC20PresetMinterPauserDecimalsFilterer{contract: contract}}, nil
}

// ERC20PresetMinterPauserDecimals is an auto generated Go binding around an Ethereum contract.
type ERC20PresetMinterPauserDecimals struct {
	ERC20PresetMinterPauserDecimalsCaller     // Read-only binding to the contract
	ERC20PresetMinterPauserDecimalsTransactor // Write-only binding to the contract
	ERC20PresetMinterPauserDecimalsFilterer   // Log filterer for contract events
}

// ERC20PresetMinterPauserDecimalsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserDecimalsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserDecimalsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserDecimalsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserDecimalsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PresetMinterPauserDecimalsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PresetMinterPauserDecimalsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PresetMinterPauserDecimalsSession struct {
	Contract     *ERC20PresetMinterPauserDecimals // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ERC20PresetMinterPauserDecimalsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PresetMinterPauserDecimalsCallerSession struct {
	Contract *ERC20PresetMinterPauserDecimalsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// ERC20PresetMinterPauserDecimalsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PresetMinterPauserDecimalsTransactorSession struct {
	Contract     *ERC20PresetMinterPauserDecimalsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// ERC20PresetMinterPauserDecimalsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PresetMinterPauserDecimalsRaw struct {
	Contract *ERC20PresetMinterPauserDecimals // Generic contract binding to access the raw methods on
}

// ERC20PresetMinterPauserDecimalsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserDecimalsCallerRaw struct {
	Contract *ERC20PresetMinterPauserDecimalsCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PresetMinterPauserDecimalsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PresetMinterPauserDecimalsTransactorRaw struct {
	Contract *ERC20PresetMinterPauserDecimalsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20PresetMinterPauserDecimals creates a new instance of ERC20PresetMinterPauserDecimals, bound to a specific deployed contract.
func NewERC20PresetMinterPauserDecimals(address common.Address, backend bind.ContractBackend) (*ERC20PresetMinterPauserDecimals, error) {
	contract, err := bindERC20PresetMinterPauserDecimals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimals{ERC20PresetMinterPauserDecimalsCaller: ERC20PresetMinterPauserDecimalsCaller{contract: contract}, ERC20PresetMinterPauserDecimalsTransactor: ERC20PresetMinterPauserDecimalsTransactor{contract: contract}, ERC20PresetMinterPauserDecimalsFilterer: ERC20PresetMinterPauserDecimalsFilterer{contract: contract}}, nil
}

// NewERC20PresetMinterPauserDecimalsCaller creates a new read-only instance of ERC20PresetMinterPauserDecimals, bound to a specific deployed contract.
func NewERC20PresetMinterPauserDecimalsCaller(address common.Address, caller bind.ContractCaller) (*ERC20PresetMinterPauserDecimalsCaller, error) {
	contract, err := bindERC20PresetMinterPauserDecimals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsCaller{contract: contract}, nil
}

// NewERC20PresetMinterPauserDecimalsTransactor creates a new write-only instance of ERC20PresetMinterPauserDecimals, bound to a specific deployed contract.
func NewERC20PresetMinterPauserDecimalsTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PresetMinterPauserDecimalsTransactor, error) {
	contract, err := bindERC20PresetMinterPauserDecimals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsTransactor{contract: contract}, nil
}

// NewERC20PresetMinterPauserDecimalsFilterer creates a new log filterer instance of ERC20PresetMinterPauserDecimals, bound to a specific deployed contract.
func NewERC20PresetMinterPauserDecimalsFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PresetMinterPauserDecimalsFilterer, error) {
	contract, err := bindERC20PresetMinterPauserDecimals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsFilterer{contract: contract}, nil
}

// bindERC20PresetMinterPauserDecimals binds a generic wrapper to an already deployed contract.
func bindERC20PresetMinterPauserDecimals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PresetMinterPauserDecimalsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PresetMinterPauserDecimals.Contract.ERC20PresetMinterPauserDecimalsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.ERC20PresetMinterPauserDecimalsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.ERC20PresetMinterPauserDecimalsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20PresetMinterPauserDecimals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.DEFAULTADMINROLE(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.DEFAULTADMINROLE(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) MINTERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.MINTERROLE(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) MINTERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.MINTERROLE(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.PAUSERROLE(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.PAUSERROLE(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Allowance(&_ERC20PresetMinterPauserDecimals.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Allowance(&_ERC20PresetMinterPauserDecimals.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.BalanceOf(&_ERC20PresetMinterPauserDecimals.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.BalanceOf(&_ERC20PresetMinterPauserDecimals.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Decimals() (uint8, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Decimals(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) Decimals() (uint8, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Decimals(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.GetRoleAdmin(&_ERC20PresetMinterPauserDecimals.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.GetRoleAdmin(&_ERC20PresetMinterPauserDecimals.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.GetRoleMember(&_ERC20PresetMinterPauserDecimals.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.GetRoleMember(&_ERC20PresetMinterPauserDecimals.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.GetRoleMemberCount(&_ERC20PresetMinterPauserDecimals.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.GetRoleMemberCount(&_ERC20PresetMinterPauserDecimals.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.HasRole(&_ERC20PresetMinterPauserDecimals.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.HasRole(&_ERC20PresetMinterPauserDecimals.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Name() (string, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Name(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) Name() (string, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Name(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Paused() (bool, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Paused(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) Paused() (bool, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Paused(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Symbol() (string, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Symbol(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) Symbol() (string, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Symbol(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20PresetMinterPauserDecimals.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) TotalSupply() (*big.Int, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.TotalSupply(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.TotalSupply(&_ERC20PresetMinterPauserDecimals.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Approve(&_ERC20PresetMinterPauserDecimals.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Approve(&_ERC20PresetMinterPauserDecimals.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Burn(&_ERC20PresetMinterPauserDecimals.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Burn(&_ERC20PresetMinterPauserDecimals.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.BurnFrom(&_ERC20PresetMinterPauserDecimals.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.BurnFrom(&_ERC20PresetMinterPauserDecimals.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.DecreaseAllowance(&_ERC20PresetMinterPauserDecimals.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.DecreaseAllowance(&_ERC20PresetMinterPauserDecimals.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.GrantRole(&_ERC20PresetMinterPauserDecimals.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.GrantRole(&_ERC20PresetMinterPauserDecimals.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.IncreaseAllowance(&_ERC20PresetMinterPauserDecimals.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.IncreaseAllowance(&_ERC20PresetMinterPauserDecimals.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Mint(&_ERC20PresetMinterPauserDecimals.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Mint(&_ERC20PresetMinterPauserDecimals.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Pause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Pause(&_ERC20PresetMinterPauserDecimals.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Pause(&_ERC20PresetMinterPauserDecimals.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.RenounceRole(&_ERC20PresetMinterPauserDecimals.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.RenounceRole(&_ERC20PresetMinterPauserDecimals.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.RevokeRole(&_ERC20PresetMinterPauserDecimals.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.RevokeRole(&_ERC20PresetMinterPauserDecimals.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Transfer(&_ERC20PresetMinterPauserDecimals.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Transfer(&_ERC20PresetMinterPauserDecimals.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.TransferFrom(&_ERC20PresetMinterPauserDecimals.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.TransferFrom(&_ERC20PresetMinterPauserDecimals.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsSession) Unpause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Unpause(&_ERC20PresetMinterPauserDecimals.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20PresetMinterPauserDecimals.Contract.Unpause(&_ERC20PresetMinterPauserDecimals.TransactOpts)
}

// ERC20PresetMinterPauserDecimalsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsApprovalIterator struct {
	Event *ERC20PresetMinterPauserDecimalsApproval // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserDecimalsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserDecimalsApproval)
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
		it.Event = new(ERC20PresetMinterPauserDecimalsApproval)
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
func (it *ERC20PresetMinterPauserDecimalsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserDecimalsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserDecimalsApproval represents a Approval event raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20PresetMinterPauserDecimalsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsApprovalIterator{contract: _ERC20PresetMinterPauserDecimals.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserDecimalsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserDecimalsApproval)
				if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) ParseApproval(log types.Log) (*ERC20PresetMinterPauserDecimalsApproval, error) {
	event := new(ERC20PresetMinterPauserDecimalsApproval)
	if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserDecimalsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsPausedIterator struct {
	Event *ERC20PresetMinterPauserDecimalsPaused // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserDecimalsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserDecimalsPaused)
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
		it.Event = new(ERC20PresetMinterPauserDecimalsPaused)
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
func (it *ERC20PresetMinterPauserDecimalsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserDecimalsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserDecimalsPaused represents a Paused event raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20PresetMinterPauserDecimalsPausedIterator, error) {

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsPausedIterator{contract: _ERC20PresetMinterPauserDecimals.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserDecimalsPaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserDecimalsPaused)
				if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) ParsePaused(log types.Log) (*ERC20PresetMinterPauserDecimalsPaused, error) {
	event := new(ERC20PresetMinterPauserDecimalsPaused)
	if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserDecimalsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsRoleGrantedIterator struct {
	Event *ERC20PresetMinterPauserDecimalsRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserDecimalsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserDecimalsRoleGranted)
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
		it.Event = new(ERC20PresetMinterPauserDecimalsRoleGranted)
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
func (it *ERC20PresetMinterPauserDecimalsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserDecimalsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserDecimalsRoleGranted represents a RoleGranted event raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20PresetMinterPauserDecimalsRoleGrantedIterator, error) {

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

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsRoleGrantedIterator{contract: _ERC20PresetMinterPauserDecimals.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserDecimalsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserDecimalsRoleGranted)
				if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) ParseRoleGranted(log types.Log) (*ERC20PresetMinterPauserDecimalsRoleGranted, error) {
	event := new(ERC20PresetMinterPauserDecimalsRoleGranted)
	if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserDecimalsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsRoleRevokedIterator struct {
	Event *ERC20PresetMinterPauserDecimalsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserDecimalsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserDecimalsRoleRevoked)
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
		it.Event = new(ERC20PresetMinterPauserDecimalsRoleRevoked)
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
func (it *ERC20PresetMinterPauserDecimalsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserDecimalsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserDecimalsRoleRevoked represents a RoleRevoked event raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20PresetMinterPauserDecimalsRoleRevokedIterator, error) {

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

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsRoleRevokedIterator{contract: _ERC20PresetMinterPauserDecimals.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserDecimalsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserDecimalsRoleRevoked)
				if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) ParseRoleRevoked(log types.Log) (*ERC20PresetMinterPauserDecimalsRoleRevoked, error) {
	event := new(ERC20PresetMinterPauserDecimalsRoleRevoked)
	if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserDecimalsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsTransferIterator struct {
	Event *ERC20PresetMinterPauserDecimalsTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserDecimalsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserDecimalsTransfer)
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
		it.Event = new(ERC20PresetMinterPauserDecimalsTransfer)
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
func (it *ERC20PresetMinterPauserDecimalsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserDecimalsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserDecimalsTransfer represents a Transfer event raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20PresetMinterPauserDecimalsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsTransferIterator{contract: _ERC20PresetMinterPauserDecimals.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserDecimalsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserDecimalsTransfer)
				if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) ParseTransfer(log types.Log) (*ERC20PresetMinterPauserDecimalsTransfer, error) {
	event := new(ERC20PresetMinterPauserDecimalsTransfer)
	if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PresetMinterPauserDecimalsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsUnpausedIterator struct {
	Event *ERC20PresetMinterPauserDecimalsUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20PresetMinterPauserDecimalsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PresetMinterPauserDecimalsUnpaused)
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
		it.Event = new(ERC20PresetMinterPauserDecimalsUnpaused)
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
func (it *ERC20PresetMinterPauserDecimalsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PresetMinterPauserDecimalsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PresetMinterPauserDecimalsUnpaused represents a Unpaused event raised by the ERC20PresetMinterPauserDecimals contract.
type ERC20PresetMinterPauserDecimalsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20PresetMinterPauserDecimalsUnpausedIterator, error) {

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20PresetMinterPauserDecimalsUnpausedIterator{contract: _ERC20PresetMinterPauserDecimals.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20PresetMinterPauserDecimalsUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20PresetMinterPauserDecimals.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PresetMinterPauserDecimalsUnpaused)
				if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ERC20PresetMinterPauserDecimals *ERC20PresetMinterPauserDecimalsFilterer) ParseUnpaused(log types.Log) (*ERC20PresetMinterPauserDecimalsUnpaused, error) {
	event := new(ERC20PresetMinterPauserDecimalsUnpaused)
	if err := _ERC20PresetMinterPauserDecimals.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
