// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lottery

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

// LotteryMetaData contains all meta data concerning the Lottery contract.
var LotteryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticketPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"WinnerSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipantsNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"restartEmpty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051611aa0380380611aa083398181016040528101906100319190610265565b335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161009991906102f4565b60405180910390fd5b6100b18161016d60201b60201c565b5060028311156100f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ed90610367565b60405180910390fd5b6201518082101561013c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610133906103cf565b60405180910390fd5b8260a08181525050806080818152505081600281905550814261015f919061041a565b60018190555050505061044d565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f819050919050565b61024481610232565b811461024e575f5ffd5b50565b5f8151905061025f8161023b565b92915050565b5f5f5f6060848603121561027c5761027b61022e565b5b5f61028986828701610251565b935050602061029a86828701610251565b92505060406102ab86828701610251565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102de826102b5565b9050919050565b6102ee816102d4565b82525050565b5f6020820190506103075f8301846102e5565b92915050565b5f82825260208201905092915050565b7f696c6c6567616c206f776e6572206665650000000000000000000000000000005f82015250565b5f61035160118361030d565b915061035c8261031d565b602082019050919050565b5f6020820190508181035f83015261037e81610345565b9050919050565b7f626164206475726174696f6e00000000000000000000000000000000000000005f82015250565b5f6103b9600c8361030d565b91506103c482610385565b602082019050919050565b5f6020820190508181035f8301526103e6816103ad565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61042482610232565b915061042f83610232565b9250828201905080821115610447576104466103ed565b5b92915050565b60805160a05161161661048a5f395f8181610bd70152610d8301525f81816103c7015281816105df015281816106920152610b9301526116165ff3fe6080604052600436106100fd575f3560e01c8063454a2ab3116100945780638da5cb5b116100635780638da5cb5b146102f7578063be9a655514610321578063c7e284b814610337578063d5b2a01a14610361578063f2fde38b1461038b576100fd565b8063454a2ab3146102615780636bc1fb8a1461027d578063715018a6146102a557806372e78e1c146102bb576100fd565b806327e235e3116100d057806327e235e3146101a95780633197cbb6146101e557806335c1d3491461020f5780633ccfd60b1461024b576100fd565b80630abb4fe3146101015780630fb5a6b41461012b5780631209b1f614610155578063146ca5311461017f575b5f5ffd5b34801561010c575f5ffd5b506101156103b3565b6040516101229190610ffa565b60405180910390f35b348015610136575f5ffd5b5061013f6103bf565b60405161014c9190610ffa565b60405180910390f35b348015610160575f5ffd5b506101696103c5565b6040516101769190610ffa565b60405180910390f35b34801561018a575f5ffd5b506101936103e9565b6040516101a09190610ffa565b60405180910390f35b3480156101b4575f5ffd5b506101cf60048036038101906101ca9190611071565b6103ef565b6040516101dc9190610ffa565b60405180910390f35b3480156101f0575f5ffd5b506101f9610404565b6040516102069190610ffa565b60405180910390f35b34801561021a575f5ffd5b50610235600480360381019061023091906110c6565b61040a565b6040516102429190611100565b60405180910390f35b348015610256575f5ffd5b5061025f610445565b005b61027b600480360381019061027691906110c6565b6105dc565b005b348015610288575f5ffd5b506102a3600480360381019061029e91906110c6565b6108bb565b005b3480156102b0575f5ffd5b506102b961097d565b005b3480156102c6575f5ffd5b506102e160048036038101906102dc91906110c6565b610990565b6040516102ee9190610ffa565b60405180910390f35b348015610302575f5ffd5b5061030b6109a5565b6040516103189190611100565b60405180910390f35b34801561032c575f5ffd5b506103356109cc565b005b348015610342575f5ffd5b5061034b610d5a565b6040516103589190610ffa565b60405180910390f35b34801561036c575f5ffd5b50610375610d81565b6040516103829190610ffa565b60405180910390f35b348015610396575f5ffd5b506103b160048036038101906103ac9190611071565b610da5565b005b5f600880549050905090565b60025481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60035481565b6004602052805f5260405f205f915090505481565b60015481565b60088181548110610419575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054116104c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bb90611173565b60405180910390fd5b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f1935050505015801561058a573d5f5f3e3d5ffd5b503373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364826040516105d19190610ffa565b60405180910390a250565b807f00000000000000000000000000000000000000000000000000000000000000008161060991906111be565b34101561064b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064290611249565b60405180910390fd5b600154421061068f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610686906112b1565b60405180910390fd5b5f7f0000000000000000000000000000000000000000000000000000000000000000836106bc91906111be565b346106c791906112cf565b90505f811115610716573373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610714573d5f5f3e3d5ffd5b505b60065f60035481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166107d657600833908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b8260075f60035481526020019081526020015f205f8282546107f89190611302565b925050819055508260055f60035481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461085c9190611302565b925050819055506003543373ffffffffffffffffffffffffffffffffffffffff167f4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b30585426040516108ae929190611335565b60405180910390a3505050565b6108c3610e29565b600154421015610908576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ff906113a6565b60405180910390fd5b5f6008805490501461094f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109469061140e565b60405180910390fd5b5f810361095e57600254610960565b805b600281905550600254426109749190611302565b60018190555050565b610985610e29565b61098e5f610eb0565b565b6007602052805f5260405f205f915090505481565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6109d4610e29565b600154421015610a19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a10906113a6565b60405180910390fd5b5f60088054905011610a60576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5790611476565b60405180910390fd5b5f60075f60035481526020019081526020015f2054610a7d610f71565b610a8791906114c1565b90505f5f90505f5f5f90505b600880549050811015610b8f5760055f60035481526020019081526020015f205f60088381548110610ac857610ac76114f1565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205483610b359190611302565b925082841015610b825760088181548110610b5357610b526114f1565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150610b8f565b8080600101915050610a93565b505f7f000000000000000000000000000000000000000000000000000000000000000060075f60035481526020019081526020015f2054610bd091906111be565b90505f60647f000000000000000000000000000000000000000000000000000000000000000083610c0191906111be565b610c0b919061151e565b90505f8183610c1a91906112cf565b90508160045f610c286109a5565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610c6f9190611302565b925050819055508060045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610cc29190611302565b925050819055506003548473ffffffffffffffffffffffffffffffffffffffff167f866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b9383604051610d129190610ffa565b60405180910390a360035f815480929190610d2c9061154e565b919050555060025442610d3f9190611302565b60018190555060085f610d529190610fa9565b505050505050565b5f426001541015610d6d575f9050610d7e565b42600154610d7b91906112cf565b90505b90565b7f000000000000000000000000000000000000000000000000000000000000000081565b610dad610e29565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e1d575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610e149190611100565b60405180910390fd5b610e2681610eb0565b50565b610e31610fa2565b73ffffffffffffffffffffffffffffffffffffffff16610e4f6109a5565b73ffffffffffffffffffffffffffffffffffffffff1614610eae57610e72610fa2565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610ea59190611100565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f4442604051602001610f859291906115b5565b604051602081830303815290604052805190602001205f1c905090565b5f33905090565b5080545f8255905f5260205f2090810190610fc49190610fc7565b50565b5b80821115610fde575f815f905550600101610fc8565b5090565b5f819050919050565b610ff481610fe2565b82525050565b5f60208201905061100d5f830184610feb565b92915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61104082611017565b9050919050565b61105081611036565b811461105a575f5ffd5b50565b5f8135905061106b81611047565b92915050565b5f6020828403121561108657611085611013565b5b5f6110938482850161105d565b91505092915050565b6110a581610fe2565b81146110af575f5ffd5b50565b5f813590506110c08161109c565b92915050565b5f602082840312156110db576110da611013565b5b5f6110e8848285016110b2565b91505092915050565b6110fa81611036565b82525050565b5f6020820190506111135f8301846110f1565b92915050565b5f82825260208201905092915050565b7f6e6f7468696e6720746f207769746864726177000000000000000000000000005f82015250565b5f61115d601383611119565b915061116882611129565b602082019050919050565b5f6020820190508181035f83015261118a81611151565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111c882610fe2565b91506111d383610fe2565b92508282026111e181610fe2565b915082820484148315176111f8576111f7611191565b5b5092915050565b7f6e6f7420656e6f756768206574686572000000000000000000000000000000005f82015250565b5f611233601083611119565b915061123e826111ff565b602082019050919050565b5f6020820190508181035f83015261126081611227565b9050919050565b7f6c6f747465727920616c72656164792066696e697368656400000000000000005f82015250565b5f61129b601883611119565b91506112a682611267565b602082019050919050565b5f6020820190508181035f8301526112c88161128f565b9050919050565b5f6112d982610fe2565b91506112e483610fe2565b92508282039050818111156112fc576112fb611191565b5b92915050565b5f61130c82610fe2565b915061131783610fe2565b925082820190508082111561132f5761132e611191565b5b92915050565b5f6040820190506113485f830185610feb565b6113556020830184610feb565b9392505050565b7f6c6f7474657279206e6f742066696e69736865640000000000000000000000005f82015250565b5f611390601483611119565b915061139b8261135c565b602082019050919050565b5f6020820190508181035f8301526113bd81611384565b9050919050565b7f68617665207061727469636970616e74730000000000000000000000000000005f82015250565b5f6113f8601183611119565b9150611403826113c4565b602082019050919050565b5f6020820190508181035f830152611425816113ec565b9050919050565b7f6e6f7420656e6f756768207061727469636970616e74730000000000000000005f82015250565b5f611460601783611119565b915061146b8261142c565b602082019050919050565b5f6020820190508181035f83015261148d81611454565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6114cb82610fe2565b91506114d683610fe2565b9250826114e6576114e5611494565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61152882610fe2565b915061153383610fe2565b92508261154357611542611494565b5b828204905092915050565b5f61155882610fe2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361158a57611589611191565b5b600182019050919050565b5f819050919050565b6115af6115aa82610fe2565b611595565b82525050565b5f6115c0828561159e565b6020820191506115d0828461159e565b602082019150819050939250505056fea264697066735822122063b89c428a09892a3484a9c82603dbc2d684790e5cb8aed409397f1424daf18664736f6c634300081c0033",
}

// LotteryABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryMetaData.ABI instead.
var LotteryABI = LotteryMetaData.ABI

// LotteryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LotteryMetaData.Bin instead.
var LotteryBin = LotteryMetaData.Bin

// DeployLottery deploys a new Ethereum contract, binding an instance of Lottery to it.
func DeployLottery(auth *bind.TransactOpts, backend bind.ContractBackend, _ownerFee *big.Int, _duration *big.Int, _ticketPrice *big.Int) (common.Address, *types.Transaction, *Lottery, error) {
	parsed, err := LotteryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LotteryBin), backend, _ownerFee, _duration, _ticketPrice)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// Lottery is an auto generated Go binding around an Ethereum contract.
type Lottery struct {
	LotteryCaller     // Read-only binding to the contract
	LotteryTransactor // Write-only binding to the contract
	LotteryFilterer   // Log filterer for contract events
}

// LotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotterySession struct {
	Contract     *Lottery          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryCallerSession struct {
	Contract *LotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryTransactorSession struct {
	Contract     *LotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryRaw struct {
	Contract *Lottery // Generic contract binding to access the raw methods on
}

// LotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryCallerRaw struct {
	Contract *LotteryCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryTransactorRaw struct {
	Contract *LotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLottery creates a new instance of Lottery, bound to a specific deployed contract.
func NewLottery(address common.Address, backend bind.ContractBackend) (*Lottery, error) {
	contract, err := bindLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// NewLotteryCaller creates a new read-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryCaller(address common.Address, caller bind.ContractCaller) (*LotteryCaller, error) {
	contract, err := bindLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryCaller{contract: contract}, nil
}

// NewLotteryTransactor creates a new write-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryTransactor, error) {
	contract, err := bindLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryTransactor{contract: contract}, nil
}

// NewLotteryFilterer creates a new log filterer instance of Lottery, bound to a specific deployed contract.
func NewLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryFilterer, error) {
	contract, err := bindLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryFilterer{contract: contract}, nil
}

// bindLottery binds a generic wrapper to an already deployed contract.
func bindLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LotteryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.LotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Lottery *LotteryCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Lottery *LotterySession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Lottery.Contract.Balances(&_Lottery.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Lottery *LotteryCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Lottery.Contract.Balances(&_Lottery.CallOpts, arg0)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Lottery *LotteryCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Lottery *LotterySession) Duration() (*big.Int, error) {
	return _Lottery.Contract.Duration(&_Lottery.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Lottery *LotteryCallerSession) Duration() (*big.Int, error) {
	return _Lottery.Contract.Duration(&_Lottery.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Lottery *LotteryCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Lottery *LotterySession) EndTime() (*big.Int, error) {
	return _Lottery.Contract.EndTime(&_Lottery.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Lottery *LotteryCallerSession) EndTime() (*big.Int, error) {
	return _Lottery.Contract.EndTime(&_Lottery.CallOpts)
}

// GetParticipantsNumber is a free data retrieval call binding the contract method 0x0abb4fe3.
//
// Solidity: function getParticipantsNumber() view returns(uint256)
func (_Lottery *LotteryCaller) GetParticipantsNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getParticipantsNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParticipantsNumber is a free data retrieval call binding the contract method 0x0abb4fe3.
//
// Solidity: function getParticipantsNumber() view returns(uint256)
func (_Lottery *LotterySession) GetParticipantsNumber() (*big.Int, error) {
	return _Lottery.Contract.GetParticipantsNumber(&_Lottery.CallOpts)
}

// GetParticipantsNumber is a free data retrieval call binding the contract method 0x0abb4fe3.
//
// Solidity: function getParticipantsNumber() view returns(uint256)
func (_Lottery *LotteryCallerSession) GetParticipantsNumber() (*big.Int, error) {
	return _Lottery.Contract.GetParticipantsNumber(&_Lottery.CallOpts)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Lottery *LotteryCaller) GetTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Lottery *LotterySession) GetTimeLeft() (*big.Int, error) {
	return _Lottery.Contract.GetTimeLeft(&_Lottery.CallOpts)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Lottery *LotteryCallerSession) GetTimeLeft() (*big.Int, error) {
	return _Lottery.Contract.GetTimeLeft(&_Lottery.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lottery *LotteryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lottery *LotterySession) Owner() (common.Address, error) {
	return _Lottery.Contract.Owner(&_Lottery.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lottery *LotteryCallerSession) Owner() (common.Address, error) {
	return _Lottery.Contract.Owner(&_Lottery.CallOpts)
}

// OwnerFee is a free data retrieval call binding the contract method 0xd5b2a01a.
//
// Solidity: function ownerFee() view returns(uint256)
func (_Lottery *LotteryCaller) OwnerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "ownerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerFee is a free data retrieval call binding the contract method 0xd5b2a01a.
//
// Solidity: function ownerFee() view returns(uint256)
func (_Lottery *LotterySession) OwnerFee() (*big.Int, error) {
	return _Lottery.Contract.OwnerFee(&_Lottery.CallOpts)
}

// OwnerFee is a free data retrieval call binding the contract method 0xd5b2a01a.
//
// Solidity: function ownerFee() view returns(uint256)
func (_Lottery *LotteryCallerSession) OwnerFee() (*big.Int, error) {
	return _Lottery.Contract.OwnerFee(&_Lottery.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_Lottery *LotteryCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_Lottery *LotterySession) Participants(arg0 *big.Int) (common.Address, error) {
	return _Lottery.Contract.Participants(&_Lottery.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_Lottery *LotteryCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _Lottery.Contract.Participants(&_Lottery.CallOpts, arg0)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Lottery *LotteryCaller) Round(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "round")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Lottery *LotterySession) Round() (*big.Int, error) {
	return _Lottery.Contract.Round(&_Lottery.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Lottery *LotteryCallerSession) Round() (*big.Int, error) {
	return _Lottery.Contract.Round(&_Lottery.CallOpts)
}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_Lottery *LotteryCaller) TicketPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "ticketPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_Lottery *LotterySession) TicketPrice() (*big.Int, error) {
	return _Lottery.Contract.TicketPrice(&_Lottery.CallOpts)
}

// TicketPrice is a free data retrieval call binding the contract method 0x1209b1f6.
//
// Solidity: function ticketPrice() view returns(uint256)
func (_Lottery *LotteryCallerSession) TicketPrice() (*big.Int, error) {
	return _Lottery.Contract.TicketPrice(&_Lottery.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x72e78e1c.
//
// Solidity: function totalWeight(uint256 ) view returns(uint256)
func (_Lottery *LotteryCaller) TotalWeight(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "totalWeight", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x72e78e1c.
//
// Solidity: function totalWeight(uint256 ) view returns(uint256)
func (_Lottery *LotterySession) TotalWeight(arg0 *big.Int) (*big.Int, error) {
	return _Lottery.Contract.TotalWeight(&_Lottery.CallOpts, arg0)
}

// TotalWeight is a free data retrieval call binding the contract method 0x72e78e1c.
//
// Solidity: function totalWeight(uint256 ) view returns(uint256)
func (_Lottery *LotteryCallerSession) TotalWeight(arg0 *big.Int) (*big.Int, error) {
	return _Lottery.Contract.TotalWeight(&_Lottery.CallOpts, arg0)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) payable returns()
func (_Lottery *LotteryTransactor) Bid(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "bid", amount)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) payable returns()
func (_Lottery *LotterySession) Bid(amount *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.Bid(&_Lottery.TransactOpts, amount)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 amount) payable returns()
func (_Lottery *LotteryTransactorSession) Bid(amount *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.Bid(&_Lottery.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lottery *LotteryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lottery *LotterySession) RenounceOwnership() (*types.Transaction, error) {
	return _Lottery.Contract.RenounceOwnership(&_Lottery.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lottery *LotteryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lottery.Contract.RenounceOwnership(&_Lottery.TransactOpts)
}

// RestartEmpty is a paid mutator transaction binding the contract method 0x6bc1fb8a.
//
// Solidity: function restartEmpty(uint256 newDuration) returns()
func (_Lottery *LotteryTransactor) RestartEmpty(opts *bind.TransactOpts, newDuration *big.Int) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "restartEmpty", newDuration)
}

// RestartEmpty is a paid mutator transaction binding the contract method 0x6bc1fb8a.
//
// Solidity: function restartEmpty(uint256 newDuration) returns()
func (_Lottery *LotterySession) RestartEmpty(newDuration *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.RestartEmpty(&_Lottery.TransactOpts, newDuration)
}

// RestartEmpty is a paid mutator transaction binding the contract method 0x6bc1fb8a.
//
// Solidity: function restartEmpty(uint256 newDuration) returns()
func (_Lottery *LotteryTransactorSession) RestartEmpty(newDuration *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.RestartEmpty(&_Lottery.TransactOpts, newDuration)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Lottery *LotteryTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Lottery *LotterySession) Start() (*types.Transaction, error) {
	return _Lottery.Contract.Start(&_Lottery.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Lottery *LotteryTransactorSession) Start() (*types.Transaction, error) {
	return _Lottery.Contract.Start(&_Lottery.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lottery *LotteryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lottery *LotterySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.TransferOwnership(&_Lottery.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lottery *LotteryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lottery.Contract.TransferOwnership(&_Lottery.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lottery *LotteryTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lottery *LotterySession) Withdraw() (*types.Transaction, error) {
	return _Lottery.Contract.Withdraw(&_Lottery.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lottery *LotteryTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Lottery.Contract.Withdraw(&_Lottery.TransactOpts)
}

// LotteryBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Lottery contract.
type LotteryBidIterator struct {
	Event *LotteryBid // Event containing the contract specifics and raw log

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
func (it *LotteryBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryBid)
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
		it.Event = new(LotteryBid)
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
func (it *LotteryBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryBid represents a Bid event raised by the Lottery contract.
type LotteryBid struct {
	Account   common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Round     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0x4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b305.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 timestamp, uint256 indexed round)
func (_Lottery *LotteryFilterer) FilterBid(opts *bind.FilterOpts, account []common.Address, round []*big.Int) (*LotteryBidIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "Bid", accountRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &LotteryBidIterator{contract: _Lottery.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0x4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b305.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 timestamp, uint256 indexed round)
func (_Lottery *LotteryFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *LotteryBid, account []common.Address, round []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "Bid", accountRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryBid)
				if err := _Lottery.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0x4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b305.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 timestamp, uint256 indexed round)
func (_Lottery *LotteryFilterer) ParseBid(log types.Log) (*LotteryBid, error) {
	event := new(LotteryBid)
	if err := _Lottery.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lottery contract.
type LotteryOwnershipTransferredIterator struct {
	Event *LotteryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LotteryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryOwnershipTransferred)
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
		it.Event = new(LotteryOwnershipTransferred)
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
func (it *LotteryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryOwnershipTransferred represents a OwnershipTransferred event raised by the Lottery contract.
type LotteryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lottery *LotteryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LotteryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LotteryOwnershipTransferredIterator{contract: _Lottery.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lottery *LotteryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LotteryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryOwnershipTransferred)
				if err := _Lottery.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lottery *LotteryFilterer) ParseOwnershipTransferred(log types.Log) (*LotteryOwnershipTransferred, error) {
	event := new(LotteryOwnershipTransferred)
	if err := _Lottery.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryWinnerSelectedIterator is returned from FilterWinnerSelected and is used to iterate over the raw logs and unpacked data for WinnerSelected events raised by the Lottery contract.
type LotteryWinnerSelectedIterator struct {
	Event *LotteryWinnerSelected // Event containing the contract specifics and raw log

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
func (it *LotteryWinnerSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryWinnerSelected)
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
		it.Event = new(LotteryWinnerSelected)
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
func (it *LotteryWinnerSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryWinnerSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryWinnerSelected represents a WinnerSelected event raised by the Lottery contract.
type LotteryWinnerSelected struct {
	Account common.Address
	Amount  *big.Int
	Round   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWinnerSelected is a free log retrieval operation binding the contract event 0x866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b93.
//
// Solidity: event WinnerSelected(address indexed account, uint256 amount, uint256 indexed round)
func (_Lottery *LotteryFilterer) FilterWinnerSelected(opts *bind.FilterOpts, account []common.Address, round []*big.Int) (*LotteryWinnerSelectedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "WinnerSelected", accountRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &LotteryWinnerSelectedIterator{contract: _Lottery.contract, event: "WinnerSelected", logs: logs, sub: sub}, nil
}

// WatchWinnerSelected is a free log subscription operation binding the contract event 0x866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b93.
//
// Solidity: event WinnerSelected(address indexed account, uint256 amount, uint256 indexed round)
func (_Lottery *LotteryFilterer) WatchWinnerSelected(opts *bind.WatchOpts, sink chan<- *LotteryWinnerSelected, account []common.Address, round []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "WinnerSelected", accountRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryWinnerSelected)
				if err := _Lottery.contract.UnpackLog(event, "WinnerSelected", log); err != nil {
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

// ParseWinnerSelected is a log parse operation binding the contract event 0x866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b93.
//
// Solidity: event WinnerSelected(address indexed account, uint256 amount, uint256 indexed round)
func (_Lottery *LotteryFilterer) ParseWinnerSelected(log types.Log) (*LotteryWinnerSelected, error) {
	event := new(LotteryWinnerSelected)
	if err := _Lottery.contract.UnpackLog(event, "WinnerSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Lottery contract.
type LotteryWithdrawIterator struct {
	Event *LotteryWithdraw // Event containing the contract specifics and raw log

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
func (it *LotteryWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryWithdraw)
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
		it.Event = new(LotteryWithdraw)
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
func (it *LotteryWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryWithdraw represents a Withdraw event raised by the Lottery contract.
type LotteryWithdraw struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Lottery *LotteryFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address) (*LotteryWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return &LotteryWithdrawIterator{contract: _Lottery.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Lottery *LotteryFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *LotteryWithdraw, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryWithdraw)
				if err := _Lottery.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Lottery *LotteryFilterer) ParseWithdraw(log types.Log) (*LotteryWithdraw, error) {
	event := new(LotteryWithdraw)
	if err := _Lottery.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
