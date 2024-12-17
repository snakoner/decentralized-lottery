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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticketPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"WinnerSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipantsNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"restartEmpty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timeFinished\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalParticipants\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051611e78380380611e7883398181016040528101906100319190610360565b335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161009991906103ef565b60405180910390fd5b6100b18161017b60201b60201c565b5060028311156100f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ed90610462565b60405180910390fd5b6201518082101561013c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610133906104ca565b60405180910390fd5b8260a08181525050806080818152505081600281905550814261015f9190610515565b60018190555061017361023c60201b60201c565b505050610548565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60086040518060a001604052805f81526020015f81526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f1515815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015560208201518160010155604082015181600201556060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff0219169083151502179055505050565b5f5ffd5b5f819050919050565b61033f8161032d565b8114610349575f5ffd5b50565b5f8151905061035a81610336565b92915050565b5f5f5f6060848603121561037757610376610329565b5b5f6103848682870161034c565b93505060206103958682870161034c565b92505060406103a68682870161034c565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103d9826103b0565b9050919050565b6103e9816103cf565b82525050565b5f6020820190506104025f8301846103e0565b92915050565b5f82825260208201905092915050565b7f696c6c6567616c206f776e6572206665650000000000000000000000000000005f82015250565b5f61044c601183610408565b915061045782610418565b602082019050919050565b5f6020820190508181035f83015261047981610440565b9050919050565b7f626164206475726174696f6e00000000000000000000000000000000000000005f82015250565b5f6104b4600c83610408565b91506104bf82610480565b602082019050919050565b5f6020820190508181035f8301526104e1816104a8565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61051f8261032d565b915061052a8361032d565b9250828201905080821115610542576105416104e8565b5b92915050565b60805160a0516118f36105855f395f8181610d0e0152610f0801525f81816103cb015281816105e3015281816106960152610cd901526118f35ff3fe6080604052600436106100fd575f3560e01c8063454a2ab3116100945780638da5cb5b116100635780638da5cb5b146102fb578063be9a655514610325578063c7e284b81461033b578063d5b2a01a14610365578063f2fde38b1461038f576100fd565b8063454a2ab3146102615780636bc1fb8a1461027d578063715018a6146102a55780638c65c81f146102bb576100fd565b806327e235e3116100d057806327e235e3146101a95780633197cbb6146101e557806335c1d3491461020f5780633ccfd60b1461024b576100fd565b80630abb4fe3146101015780630fb5a6b41461012b5780631209b1f614610155578063146ca5311461017f575b5f5ffd5b34801561010c575f5ffd5b506101156103b7565b604051610122919061126c565b60405180910390f35b348015610136575f5ffd5b5061013f6103c3565b60405161014c919061126c565b60405180910390f35b348015610160575f5ffd5b506101696103c9565b604051610176919061126c565b60405180910390f35b34801561018a575f5ffd5b506101936103ed565b6040516101a0919061126c565b60405180910390f35b3480156101b4575f5ffd5b506101cf60048036038101906101ca91906112e3565b6103f3565b6040516101dc919061126c565b60405180910390f35b3480156101f0575f5ffd5b506101f9610408565b604051610206919061126c565b60405180910390f35b34801561021a575f5ffd5b5061023560048036038101906102309190611338565b61040e565b6040516102429190611372565b60405180910390f35b348015610256575f5ffd5b5061025f610449565b005b61027b60048036038101906102769190611338565b6105e0565b005b348015610288575f5ffd5b506102a3600480360381019061029e9190611338565b6108ef565b005b3480156102b0575f5ffd5b506102b96109d1565b005b3480156102c6575f5ffd5b506102e160048036038101906102dc9190611338565b6109e4565b6040516102f29594939291906113a5565b60405180910390f35b348015610306575f5ffd5b5061030f610a51565b60405161031c9190611372565b60405180910390f35b348015610330575f5ffd5b50610339610a78565b005b348015610346575f5ffd5b5061034f610edf565b60405161035c919061126c565b60405180910390f35b348015610370575f5ffd5b50610379610f06565b604051610386919061126c565b60405180910390f35b34801561039a575f5ffd5b506103b560048036038101906103b091906112e3565b610f2a565b005b5f600780549050905090565b60025481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60035481565b6004602052805f5260405f205f915090505481565b60015481565b6007818154811061041d575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054116104c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bf90611450565b60405180910390fd5b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f1935050505015801561058e573d5f5f3e3d5ffd5b503373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364826040516105d5919061126c565b60405180910390a250565b807f00000000000000000000000000000000000000000000000000000000000000008161060d919061149b565b34101561064f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064690611526565b60405180910390fd5b6001544210610693576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068a9061158e565b60405180910390fd5b5f7f0000000000000000000000000000000000000000000000000000000000000000836106c0919061149b565b346106cb91906115ac565b90505f81111561071a573373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610718573d5f5f3e3d5ffd5b505b5f600860035481548110610731576107306115df565b5b905f5260205f209060040201905060065f60035481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661081857600733908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806002015f8154809291906108129061160c565b91905055505b83816001015f82825461082b9190611653565b925050819055508360055f60035481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461088f9190611653565b925050819055506003543373ffffffffffffffffffffffffffffffffffffffff167f4dcc013473324698bfbe263facec4ea4b1bc43624236542deabec62c2122b30586426040516108e1929190611686565b60405180910390a350505050565b6108f7610fae565b60015442101561093c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610933906116f7565b60405180910390fd5b5f600860035481548110610953576109526115df565b5b905f5260205f20906004020160020154146109a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099a9061175f565b60405180910390fd5b5f81036109b2576002546109b4565b805b600281905550600254426109c89190611653565b60018190555050565b6109d9610fae565b6109e25f611035565b565b600881815481106109f3575f80fd5b905f5260205f2090600402015f91509050805f015490806001015490806002015490806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160149054906101000a900460ff16905085565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610a80610fae565b600154421015610ac5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abc906116f7565b60405180910390fd5b5f600860035481548110610adc57610adb6115df565b5b905f5260205f2090600402016002015411610b2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b23906117c7565b60405180910390fd5b5f600860035481548110610b4357610b426115df565b5b905f5260205f20906004020190505f600780549050610b606110f6565b610b6a9190611812565b90505f5f90505f5f90505b600780549050811015610cd55760055f60035481526020019081526020015f205f60078381548110610baa57610ba96115df565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205482610c179190611653565b915081831015610cc85760078181548110610c3557610c346115df565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018460030160146101000a81548160ff02191690831515021790555042845f0181905550610cd5565b8080600101915050610b75565b505f7f00000000000000000000000000000000000000000000000000000000000000008460010154610d07919061149b565b90505f60647f000000000000000000000000000000000000000000000000000000000000000083610d38919061149b565b610d429190611842565b90505f8183610d5191906115ac565b90508160045f610d5f610a51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610da69190611653565b925050819055508060045f886003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610e1c9190611653565b92505081905550600354866003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b9383604051610e8f919061126c565b60405180910390a3610e9f611127565b60035f815480929190610eb19061160c565b919050555060025442610ec49190611653565b60018190555060075f610ed7919061121b565b505050505050565b5f426001541015610ef2575f9050610f03565b42600154610f0091906115ac565b90505b90565b7f000000000000000000000000000000000000000000000000000000000000000081565b610f32610fae565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610fa2575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610f999190611372565b60405180910390fd5b610fab81611035565b50565b610fb6611214565b73ffffffffffffffffffffffffffffffffffffffff16610fd4610a51565b73ffffffffffffffffffffffffffffffffffffffff161461103357610ff7611214565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161102a9190611372565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f444260405160200161110a929190611892565b604051602081830303815290604052805190602001205f1c905090565b60086040518060a001604052805f81526020015f81526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f1515815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015560208201518160010155604082015181600201556060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff0219169083151502179055505050565b5f33905090565b5080545f8255905f5260205f20908101906112369190611239565b50565b5b80821115611250575f815f90555060010161123a565b5090565b5f819050919050565b61126681611254565b82525050565b5f60208201905061127f5f83018461125d565b92915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112b282611289565b9050919050565b6112c2816112a8565b81146112cc575f5ffd5b50565b5f813590506112dd816112b9565b92915050565b5f602082840312156112f8576112f7611285565b5b5f611305848285016112cf565b91505092915050565b61131781611254565b8114611321575f5ffd5b50565b5f813590506113328161130e565b92915050565b5f6020828403121561134d5761134c611285565b5b5f61135a84828501611324565b91505092915050565b61136c816112a8565b82525050565b5f6020820190506113855f830184611363565b92915050565b5f8115159050919050565b61139f8161138b565b82525050565b5f60a0820190506113b85f83018861125d565b6113c5602083018761125d565b6113d2604083018661125d565b6113df6060830185611363565b6113ec6080830184611396565b9695505050505050565b5f82825260208201905092915050565b7f6e6f7468696e6720746f207769746864726177000000000000000000000000005f82015250565b5f61143a6013836113f6565b915061144582611406565b602082019050919050565b5f6020820190508181035f8301526114678161142e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6114a582611254565b91506114b083611254565b92508282026114be81611254565b915082820484148315176114d5576114d461146e565b5b5092915050565b7f6e6f7420656e6f756768206574686572000000000000000000000000000000005f82015250565b5f6115106010836113f6565b915061151b826114dc565b602082019050919050565b5f6020820190508181035f83015261153d81611504565b9050919050565b7f6c6f747465727920616c72656164792066696e697368656400000000000000005f82015250565b5f6115786018836113f6565b915061158382611544565b602082019050919050565b5f6020820190508181035f8301526115a58161156c565b9050919050565b5f6115b682611254565b91506115c183611254565b92508282039050818111156115d9576115d861146e565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61161682611254565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116485761164761146e565b5b600182019050919050565b5f61165d82611254565b915061166883611254565b92508282019050808211156116805761167f61146e565b5b92915050565b5f6040820190506116995f83018561125d565b6116a6602083018461125d565b9392505050565b7f6c6f7474657279206e6f742066696e69736865640000000000000000000000005f82015250565b5f6116e16014836113f6565b91506116ec826116ad565b602082019050919050565b5f6020820190508181035f83015261170e816116d5565b9050919050565b7f68617665207061727469636970616e74730000000000000000000000000000005f82015250565b5f6117496011836113f6565b915061175482611715565b602082019050919050565b5f6020820190508181035f8301526117768161173d565b9050919050565b7f6e6f7420656e6f756768207061727469636970616e74730000000000000000005f82015250565b5f6117b16017836113f6565b91506117bc8261177d565b602082019050919050565b5f6020820190508181035f8301526117de816117a5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61181c82611254565b915061182783611254565b925082611837576118366117e5565b5b828206905092915050565b5f61184c82611254565b915061185783611254565b925082611867576118666117e5565b5b828204905092915050565b5f819050919050565b61188c61188782611254565b611872565b82525050565b5f61189d828561187b565b6020820191506118ad828461187b565b602082019150819050939250505056fea26469706673582212209b5cc9ac2cfb9758665fdde736654db256117c564bbbb04f643f270159f1cf9764736f6c634300081c0033",
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

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 timeFinished, uint256 totalWeight, uint256 totalParticipants, address winner, bool finished)
func (_Lottery *LotteryCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TimeFinished      *big.Int
	TotalWeight       *big.Int
	TotalParticipants *big.Int
	Winner            common.Address
	Finished          bool
}, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "rounds", arg0)

	outstruct := new(struct {
		TimeFinished      *big.Int
		TotalWeight       *big.Int
		TotalParticipants *big.Int
		Winner            common.Address
		Finished          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TimeFinished = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalWeight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalParticipants = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Winner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Finished = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 timeFinished, uint256 totalWeight, uint256 totalParticipants, address winner, bool finished)
func (_Lottery *LotterySession) Rounds(arg0 *big.Int) (struct {
	TimeFinished      *big.Int
	TotalWeight       *big.Int
	TotalParticipants *big.Int
	Winner            common.Address
	Finished          bool
}, error) {
	return _Lottery.Contract.Rounds(&_Lottery.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 timeFinished, uint256 totalWeight, uint256 totalParticipants, address winner, bool finished)
func (_Lottery *LotteryCallerSession) Rounds(arg0 *big.Int) (struct {
	TimeFinished      *big.Int
	TotalWeight       *big.Int
	TotalParticipants *big.Int
	Winner            common.Address
	Finished          bool
}, error) {
	return _Lottery.Contract.Rounds(&_Lottery.CallOpts, arg0)
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
