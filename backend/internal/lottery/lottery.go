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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticketPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"WinnerSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bidFromBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipantsNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"restartEmpty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"weights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161201138038061201183398181016040528101906100319190610265565b335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161009991906102f4565b60405180910390fd5b6100b18161016d60201b60201c565b5060028311156100f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ed90610367565b60405180910390fd5b6201518082101561013c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610133906103cf565b60405180910390fd5b8260a08181525050806080818152505081600281905550814261015f919061041a565b60018190555050505061044d565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f819050919050565b61024481610232565b811461024e575f5ffd5b50565b5f8151905061025f8161023b565b92915050565b5f5f5f6060848603121561027c5761027b61022e565b5b5f61028986828701610251565b935050602061029a86828701610251565b92505060406102ab86828701610251565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102de826102b5565b9050919050565b6102ee816102d4565b82525050565b5f6020820190506103075f8301846102e5565b92915050565b5f82825260208201905092915050565b7f696c6c6567616c206f776e6572206665650000000000000000000000000000005f82015250565b5f61035160118361030d565b915061035c8261031d565b602082019050919050565b5f6020820190508181035f83015261037e81610345565b9050919050565b7f626164206475726174696f6e00000000000000000000000000000000000000005f82015250565b5f6103b9600c8361030d565b91506103c482610385565b602082019050919050565b5f6020820190508181035f8301526103e6816103ad565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61042482610232565b915061042f83610232565b9250828201905080821115610447576104466103ed565b5b92915050565b60805160a051611b806104915f395f8181610f4c01526111df01525f81816104680152818161065c0152818161095401528181610a070152610f080152611b805ff3fe60806040526004361061011e575f3560e01c8063454a2ab31161009f578063be9a655511610063578063be9a6555146103b8578063c7e284b8146103ce578063d0e30db0146103f8578063d5b2a01a14610402578063f2fde38b1461042c5761011e565b8063454a2ab3146102f85780636bc1fb8a14610314578063715018a61461033c57806372e78e1c146103525780638da5cb5b1461038e5761011e565b80632e1a7d4d116100e65780632e1a7d4d146102065780633197cbb61461022e57806334e4d2d91461025857806335c1d349146102805780634164c367146102bc5761011e565b80630abb4fe3146101225780630fb5a6b41461014c5780631209b1f614610176578063146ca531146101a057806327e235e3146101ca575b5f5ffd5b34801561012d575f5ffd5b50610136610454565b6040516101439190611456565b60405180910390f35b348015610157575f5ffd5b50610160610460565b60405161016d9190611456565b60405180910390f35b348015610181575f5ffd5b5061018a610466565b6040516101979190611456565b60405180910390f35b3480156101ab575f5ffd5b506101b461048a565b6040516101c19190611456565b60405180910390f35b3480156101d5575f5ffd5b506101f060048036038101906101eb91906114cd565b610490565b6040516101fd9190611456565b60405180910390f35b348015610211575f5ffd5b5061022c60048036038101906102279190611522565b6104a5565b005b348015610239575f5ffd5b5061024261060f565b60405161024f9190611456565b60405180910390f35b348015610263575f5ffd5b5061027e60048036038101906102799190611522565b610615565b005b34801561028b575f5ffd5b506102a660048036038101906102a19190611522565b6108f6565b6040516102b3919061155c565b60405180910390f35b3480156102c7575f5ffd5b506102e260048036038101906102dd9190611575565b610931565b6040516102ef9190611456565b60405180910390f35b610312600480360381019061030d9190611522565b610951565b005b34801561031f575f5ffd5b5061033a60048036038101906103359190611522565b610c30565b005b348015610347575f5ffd5b50610350610cf2565b005b34801561035d575f5ffd5b5061037860048036038101906103739190611522565b610d05565b6040516103859190611456565b60405180910390f35b348015610399575f5ffd5b506103a2610d1a565b6040516103af919061155c565b60405180910390f35b3480156103c3575f5ffd5b506103cc610d41565b005b3480156103d9575f5ffd5b506103e26110cf565b6040516103ef9190611456565b60405180910390f35b6104006110f6565b005b34801561040d575f5ffd5b506104166111dd565b6040516104239190611456565b60405180910390f35b348015610437575f5ffd5b50610452600480360381019061044d91906114cd565b611201565b005b5f600880549050905090565b60025481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60035481565b6004602052805f5260405f205f915090505481565b8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101580156104f157505f81115b610530576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105279061160d565b60405180910390fd5b8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156105bd573d5f5f3e3d5ffd5b503373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364826040516106049190611456565b60405180910390a250565b60015481565b6001544210610659576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065090611675565b60405180910390fd5b5f7f00000000000000000000000000000000000000000000000000000000000000008261068691906116c0565b90508060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610708576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ff9061174b565b60405180910390fd5b8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254039250508190555060065f60035481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661081257600833908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b8160075f60035481526020019081526020015f205f8282546108349190611769565b925050819055508160055f60035481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108989190611769565b925050819055506003543373ffffffffffffffffffffffffffffffffffffffff167f4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b30584426040516108ea92919061179c565b60405180910390a35050565b60088181548110610905575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6005602052815f5260405f20602052805f5260405f205f91509150505481565b807f00000000000000000000000000000000000000000000000000000000000000008161097e91906116c0565b3410156109c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b79061180d565b60405180910390fd5b6001544210610a04576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109fb90611675565b60405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000083610a3191906116c0565b34610a3c919061182b565b90505f811115610a8b573373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610a89573d5f5f3e3d5ffd5b505b60065f60035481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610b4b57600833908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b8260075f60035481526020019081526020015f205f828254610b6d9190611769565b925050819055508260055f60035481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610bd19190611769565b925050819055506003543373ffffffffffffffffffffffffffffffffffffffff167f4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b3058542604051610c2392919061179c565b60405180910390a3505050565b610c38611285565b600154421015610c7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c74906118a8565b60405180910390fd5b5f60088054905014610cc4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbb90611910565b60405180910390fd5b5f8103610cd357600254610cd5565b805b60028190555060025442610ce99190611769565b60018190555050565b610cfa611285565b610d035f61130c565b565b6007602052805f5260405f205f915090505481565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610d49611285565b600154421015610d8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d85906118a8565b60405180910390fd5b5f60088054905011610dd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcc90611978565b60405180910390fd5b5f60075f60035481526020019081526020015f2054610df26113cd565b610dfc91906119c3565b90505f5f90505f5f5f90505b600880549050811015610f045760055f60035481526020019081526020015f205f60088381548110610e3d57610e3c6119f3565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205483610eaa9190611769565b925082841015610ef75760088181548110610ec857610ec76119f3565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150610f04565b8080600101915050610e08565b505f7f000000000000000000000000000000000000000000000000000000000000000060075f60035481526020019081526020015f2054610f4591906116c0565b90505f60647f000000000000000000000000000000000000000000000000000000000000000083610f7691906116c0565b610f809190611a20565b90505f8183610f8f919061182b565b90508160045f610f9d610d1a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610fe49190611769565b925050819055508060045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546110379190611769565b925050819055506003548473ffffffffffffffffffffffffffffffffffffffff167f866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b93836040516110879190611456565b60405180910390a360035f8154809291906110a190611a50565b9190505550600254426110b49190611769565b60018190555060085f6110c79190611405565b505050505050565b5f4260015410156110e2575f90506110f3565b426001546110f0919061182b565b90505b90565b5f3411611138576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112f90611ae1565b60405180910390fd5b3460045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546111849190611769565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a1534426040516111d392919061179c565b60405180910390a2565b7f000000000000000000000000000000000000000000000000000000000000000081565b611209611285565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611279575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401611270919061155c565b60405180910390fd5b6112828161130c565b50565b61128d6113fe565b73ffffffffffffffffffffffffffffffffffffffff166112ab610d1a565b73ffffffffffffffffffffffffffffffffffffffff161461130a576112ce6113fe565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401611301919061155c565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f44426040516020016113e1929190611b1f565b604051602081830303815290604052805190602001205f1c905090565b5f33905090565b5080545f8255905f5260205f20908101906114209190611423565b50565b5b8082111561143a575f815f905550600101611424565b5090565b5f819050919050565b6114508161143e565b82525050565b5f6020820190506114695f830184611447565b92915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61149c82611473565b9050919050565b6114ac81611492565b81146114b6575f5ffd5b50565b5f813590506114c7816114a3565b92915050565b5f602082840312156114e2576114e161146f565b5b5f6114ef848285016114b9565b91505092915050565b6115018161143e565b811461150b575f5ffd5b50565b5f8135905061151c816114f8565b92915050565b5f602082840312156115375761153661146f565b5b5f6115448482850161150e565b91505092915050565b61155681611492565b82525050565b5f60208201905061156f5f83018461154d565b92915050565b5f5f6040838503121561158b5761158a61146f565b5b5f6115988582860161150e565b92505060206115a9858286016114b9565b9150509250929050565b5f82825260208201905092915050565b7f6e6f7468696e6720746f207769746864726177000000000000000000000000005f82015250565b5f6115f76013836115b3565b9150611602826115c3565b602082019050919050565b5f6020820190508181035f830152611624816115eb565b9050919050565b7f6c6f747465727920616c72656164792066696e697368656400000000000000005f82015250565b5f61165f6018836115b3565b915061166a8261162b565b602082019050919050565b5f6020820190508181035f83015261168c81611653565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6116ca8261143e565b91506116d58361143e565b92508282026116e38161143e565b915082820484148315176116fa576116f9611693565b5b5092915050565b7f696e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f6117356014836115b3565b915061174082611701565b602082019050919050565b5f6020820190508181035f83015261176281611729565b9050919050565b5f6117738261143e565b915061177e8361143e565b925082820190508082111561179657611795611693565b5b92915050565b5f6040820190506117af5f830185611447565b6117bc6020830184611447565b9392505050565b7f6e6f7420656e6f756768206574686572000000000000000000000000000000005f82015250565b5f6117f76010836115b3565b9150611802826117c3565b602082019050919050565b5f6020820190508181035f830152611824816117eb565b9050919050565b5f6118358261143e565b91506118408361143e565b925082820390508181111561185857611857611693565b5b92915050565b7f6c6f7474657279206e6f742066696e69736865640000000000000000000000005f82015250565b5f6118926014836115b3565b915061189d8261185e565b602082019050919050565b5f6020820190508181035f8301526118bf81611886565b9050919050565b7f68617665207061727469636970616e74730000000000000000000000000000005f82015250565b5f6118fa6011836115b3565b9150611905826118c6565b602082019050919050565b5f6020820190508181035f830152611927816118ee565b9050919050565b7f6e6f7420656e6f756768207061727469636970616e74730000000000000000005f82015250565b5f6119626017836115b3565b915061196d8261192e565b602082019050919050565b5f6020820190508181035f83015261198f81611956565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6119cd8261143e565b91506119d88361143e565b9250826119e8576119e7611996565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f611a2a8261143e565b9150611a358361143e565b925082611a4557611a44611996565b5b828204905092915050565b5f611a5a8261143e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a8c57611a8b611693565b5b600182019050919050565b7f6465706f736974206d7573742062652067726561746572207468616e203000005f82015250565b5f611acb601e836115b3565b9150611ad682611a97565b602082019050919050565b5f6020820190508181035f830152611af881611abf565b9050919050565b5f819050919050565b611b19611b148261143e565b611aff565b82525050565b5f611b2a8285611b08565b602082019150611b3a8284611b08565b602082019150819050939250505056fea2646970667358221220024a78b98b6489ccfac8a865be3acd935c6b2e6ad61e8cfd9b9db440ee561e1d64736f6c634300081c0033",
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

// Weights is a free data retrieval call binding the contract method 0x4164c367.
//
// Solidity: function weights(uint256 , address ) view returns(uint256)
func (_Lottery *LotteryCaller) Weights(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "weights", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Weights is a free data retrieval call binding the contract method 0x4164c367.
//
// Solidity: function weights(uint256 , address ) view returns(uint256)
func (_Lottery *LotterySession) Weights(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Lottery.Contract.Weights(&_Lottery.CallOpts, arg0, arg1)
}

// Weights is a free data retrieval call binding the contract method 0x4164c367.
//
// Solidity: function weights(uint256 , address ) view returns(uint256)
func (_Lottery *LotteryCallerSession) Weights(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Lottery.Contract.Weights(&_Lottery.CallOpts, arg0, arg1)
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

// BidFromBalance is a paid mutator transaction binding the contract method 0x34e4d2d9.
//
// Solidity: function bidFromBalance(uint256 amount) returns()
func (_Lottery *LotteryTransactor) BidFromBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "bidFromBalance", amount)
}

// BidFromBalance is a paid mutator transaction binding the contract method 0x34e4d2d9.
//
// Solidity: function bidFromBalance(uint256 amount) returns()
func (_Lottery *LotterySession) BidFromBalance(amount *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.BidFromBalance(&_Lottery.TransactOpts, amount)
}

// BidFromBalance is a paid mutator transaction binding the contract method 0x34e4d2d9.
//
// Solidity: function bidFromBalance(uint256 amount) returns()
func (_Lottery *LotteryTransactorSession) BidFromBalance(amount *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.BidFromBalance(&_Lottery.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Lottery *LotteryTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Lottery *LotterySession) Deposit() (*types.Transaction, error) {
	return _Lottery.Contract.Deposit(&_Lottery.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Lottery *LotteryTransactorSession) Deposit() (*types.Transaction, error) {
	return _Lottery.Contract.Deposit(&_Lottery.TransactOpts)
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

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Lottery *LotteryTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Lottery *LotterySession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.Withdraw(&_Lottery.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Lottery *LotteryTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.Withdraw(&_Lottery.TransactOpts, amount)
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

// LotteryDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Lottery contract.
type LotteryDepositIterator struct {
	Event *LotteryDeposit // Event containing the contract specifics and raw log

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
func (it *LotteryDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryDeposit)
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
		it.Event = new(LotteryDeposit)
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
func (it *LotteryDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryDeposit represents a Deposit event raised by the Lottery contract.
type LotteryDeposit struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 timestamp)
func (_Lottery *LotteryFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*LotteryDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &LotteryDepositIterator{contract: _Lottery.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 timestamp)
func (_Lottery *LotteryFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LotteryDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryDeposit)
				if err := _Lottery.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 timestamp)
func (_Lottery *LotteryFilterer) ParseDeposit(log types.Log) (*LotteryDeposit, error) {
	event := new(LotteryDeposit)
	if err := _Lottery.contract.UnpackLog(event, "Deposit", log); err != nil {
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
