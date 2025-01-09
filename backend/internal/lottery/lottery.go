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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"WinnerSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bidFromBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipantsNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticketPrice\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"participantExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"restartEmpty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"weights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156041575f5ffd5b50608051612b276100685f395f8181611689015281816116de01526118980152612b275ff3fe608060405260043610610165575f3560e01c806352d1902d116100d05780638da5cb5b11610089578063c7e284b811610063578063c7e284b8146104e9578063d0e30db014610513578063d5b2a01a1461051d578063f2fde38b1461054757610165565b80638da5cb5b1461047f578063ad3cb1cc146104a9578063be9a6555146104d357610165565b806352d1902d146103775780635c95bd86146103a15780636bc1fb8a146103dd578063715018a61461040557806372e78e1c1461041b57806380d859111461045757610165565b80633197cbb6116101225780633197cbb61461027557806334e4d2d91461029f57806335c1d349146102c75780634164c36714610303578063454a2ab31461033f5780634f1ef2861461035b57610165565b80630abb4fe3146101695780630fb5a6b4146101935780631209b1f6146101bd578063146ca531146101e757806327e235e3146102115780632e1a7d4d1461024d575b5f5ffd5b348015610174575f5ffd5b5061017d61056f565b60405161018a9190611f70565b60405180910390f35b34801561019e575f5ffd5b506101a761057b565b6040516101b49190611f70565b60405180910390f35b3480156101c8575f5ffd5b506101d1610581565b6040516101de9190611f70565b60405180910390f35b3480156101f2575f5ffd5b506101fb610587565b6040516102089190611f70565b60405180910390f35b34801561021c575f5ffd5b5061023760048036038101906102329190611ff4565b61058d565b6040516102449190611f70565b60405180910390f35b348015610258575f5ffd5b50610273600480360381019061026e9190612049565b6105a2565b005b348015610280575f5ffd5b5061028961070c565b6040516102969190611f70565b60405180910390f35b3480156102aa575f5ffd5b506102c560048036038101906102c09190612049565b610711565b005b3480156102d2575f5ffd5b506102ed60048036038101906102e89190612049565b610a37565b6040516102fa9190612083565b60405180910390f35b34801561030e575f5ffd5b506103296004803603810190610324919061209c565b610a72565b6040516103369190611f70565b60405180910390f35b61035960048036038101906103549190612049565b610a92565b005b61037560048036038101906103709190612216565b610d98565b005b348015610382575f5ffd5b5061038b610db7565b6040516103989190612288565b60405180910390f35b3480156103ac575f5ffd5b506103c760048036038101906103c2919061209c565b610de8565b6040516103d491906122bb565b60405180910390f35b3480156103e8575f5ffd5b5061040360048036038101906103fe9190612049565b610e12565b005b348015610410575f5ffd5b50610419610ed2565b005b348015610426575f5ffd5b50610441600480360381019061043c9190612049565b610ee5565b60405161044e9190611f70565b60405180910390f35b348015610462575f5ffd5b5061047d600480360381019061047891906122d4565b610efa565b005b34801561048a575f5ffd5b50610493611135565b6040516104a09190612083565b60405180910390f35b3480156104b4575f5ffd5b506104bd61116a565b6040516104ca9190612384565b60405180910390f35b3480156104de575f5ffd5b506104e76111a3565b005b3480156104f4575f5ffd5b506104fd6114f3565b60405161050a9190611f70565b60405180910390f35b61051b611518565b005b348015610528575f5ffd5b506105316115fd565b60405161053e9190611f70565b60405180910390f35b348015610552575f5ffd5b5061056d60048036038101906105689190611ff4565b611603565b005b5f600980549050905090565b60015481565b60035481565b60025481565b6005602052805f5260405f205f915090505481565b8060055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101580156105ee57505f81115b61062d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610624906123ee565b60405180910390fd5b8060055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156106ba573d5f5f3e3d5ffd5b503373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364826040516107019190611f70565b60405180910390a250565b5f5481565b5f544210610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074b90612456565b60405180910390fd5b5f6003548261076391906124a1565b90508060055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156107e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107dc9061252c565b60405180910390fd5b8060055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254039250508190555060075f60025481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661095557600160075f60025481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600933908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b8160085f60025481526020019081526020015f205f828254610977919061254a565b925050819055508160065f60025481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109db919061254a565b925050819055506002543373ffffffffffffffffffffffffffffffffffffffff167f19421268847f42dd61705778018ddfc43bcdce8517e7a630acb12f122c70948184604051610a2b9190611f70565b60405180910390a35050565b60098181548110610a46575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6006602052815f5260405f20602052805f5260405f205f91509150505481565b8060035481610aa191906124a1565b341015610ae3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ada906125c7565b60405180910390fd5b5f544210610b26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1d90612456565b60405180910390fd5b5f60035483610b3591906124a1565b34610b4091906125e5565b90505f811115610b8f573373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610b8d573d5f5f3e3d5ffd5b505b60075f60025481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610cb557600160075f60025481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600933908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b8260085f60025481526020019081526020015f205f828254610cd7919061254a565b925050819055508260065f60025481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610d3b919061254a565b925050819055506002543373ffffffffffffffffffffffffffffffffffffffff167f19421268847f42dd61705778018ddfc43bcdce8517e7a630acb12f122c70948185604051610d8b9190611f70565b60405180910390a3505050565b610da0611687565b610da98261176d565b610db38282611778565b5050565b5f610dc0611896565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b6007602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b610e1a61191d565b5f54421015610e5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e5590612662565b60405180910390fd5b5f60098054905014610ea5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9c906126ca565b60405180910390fd5b5f8103610eb457600154610eb6565b805b60018190555060015442610eca919061254a565b5f8190555050565b610eda61191d565b610ee35f6119a4565b565b6008602052805f5260405f205f915090505481565b5f610f03611a75565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff16148015610f4b5750825b90505f60018367ffffffffffffffff16148015610f7e57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610f8c575080155b15610fc3576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315611010576001855f0160086101000a81548160ff0219169083151502179055505b6002881115611054576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104b90612732565b60405180910390fd5b6201518087101561109a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110919061279a565b60405180910390fd5b6110a333611a9c565b6110ab611ab0565b87600481905550856003819055508660018190555086426110cc919061254a565b5f81905550831561112b575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051611122919061280d565b60405180910390a15b5050505050505050565b5f5f61113f611aba565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6111ab61191d565b5f544210156111ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e690612662565b60405180910390fd5b5f60098054905011611236576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122d90612870565b60405180910390fd5b5f60085f60025481526020019081526020015f2054611253611ae1565b61125d91906128bb565b90505f5f90505f5f5f90505b6009805490508110156113655760065f60025481526020019081526020015f205f6009838154811061129e5761129d6128eb565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548361130b919061254a565b9250828410156113585760098181548110611329576113286128eb565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150611365565b8080600101915050611269565b505f60035460085f60025481526020019081526020015f205461138891906124a1565b90505f60646004548361139b91906124a1565b6113a59190612918565b90505f81836113b491906125e5565b90508160055f6113c2611135565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611409919061254a565b925050819055508060055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461145c919061254a565b925050819055506002548473ffffffffffffffffffffffffffffffffffffffff167f866efae43947560fe9d0de9013bc323d5718730d5c1543261b48a7bcb0717b93836040516114ac9190611f70565b60405180910390a360025f8154809291906114c690612948565b9190505550600154426114d9919061254a565b5f8190555060095f6114eb9190611f1f565b505050505050565b5f425f541015611505575f9050611515565b425f5461151291906125e5565b90505b90565b5f341161155a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611551906129d9565b60405180910390fd5b3460055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546115a6919061254a565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516115f39190611f70565b60405180910390a2565b60045481565b61160b61191d565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361167b575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016116729190612083565b60405180910390fd5b611684816119a4565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148061173457507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661171b611b12565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561176b576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b61177561191d565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156117e057506040513d601f19601f820116820180604052508101906117dd9190612a21565b60015b61182157816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016118189190612083565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b811461188757806040517faa1d49a400000000000000000000000000000000000000000000000000000000815260040161187e9190612288565b60405180910390fd5b6118918383611b65565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161461191b576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b611925611bd7565b73ffffffffffffffffffffffffffffffffffffffff16611943611135565b73ffffffffffffffffffffffffffffffffffffffff16146119a257611966611bd7565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016119999190612083565b60405180910390fd5b565b5f6119ad611aba565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b611aa4611bde565b611aad81611c1e565b50565b611ab8611bde565b565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f4442604051602001611af5929190612a6c565b604051602081830303815290604052805190602001205f1c905090565b5f611b3e7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b611ca2565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611b6e82611cab565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115611bca57611bc48282611d74565b50611bd3565b611bd2611df4565b5b5050565b5f33905090565b611be6611e30565b611c1c576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b611c26611bde565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611c96575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401611c8d9190612083565b60405180910390fd5b611c9f816119a4565b50565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b03611d0657806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401611cfd9190612083565b60405180910390fd5b80611d327f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b611ca2565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051611d9d9190612adb565b5f60405180830381855af49150503d805f8114611dd5576040519150601f19603f3d011682016040523d82523d5f602084013e611dda565b606091505b5091509150611dea858383611e4e565b9250505092915050565b5f341115611e2e576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f611e39611a75565b5f0160089054906101000a900460ff16905090565b606082611e6357611e5e82611edb565b611ed3565b5f8251148015611e8957505f8473ffffffffffffffffffffffffffffffffffffffff163b145b15611ecb57836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401611ec29190612083565b60405180910390fd5b819050611ed4565b5b9392505050565b5f81511115611eed5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5080545f8255905f5260205f2090810190611f3a9190611f3d565b50565b5b80821115611f54575f815f905550600101611f3e565b5090565b5f819050919050565b611f6a81611f58565b82525050565b5f602082019050611f835f830184611f61565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611fc382611f9a565b9050919050565b611fd381611fb9565b8114611fdd575f5ffd5b50565b5f81359050611fee81611fca565b92915050565b5f6020828403121561200957612008611f92565b5b5f61201684828501611fe0565b91505092915050565b61202881611f58565b8114612032575f5ffd5b50565b5f813590506120438161201f565b92915050565b5f6020828403121561205e5761205d611f92565b5b5f61206b84828501612035565b91505092915050565b61207d81611fb9565b82525050565b5f6020820190506120965f830184612074565b92915050565b5f5f604083850312156120b2576120b1611f92565b5b5f6120bf85828601612035565b92505060206120d085828601611fe0565b9150509250929050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b612128826120e2565b810181811067ffffffffffffffff82111715612147576121466120f2565b5b80604052505050565b5f612159611f89565b9050612165828261211f565b919050565b5f67ffffffffffffffff821115612184576121836120f2565b5b61218d826120e2565b9050602081019050919050565b828183375f83830152505050565b5f6121ba6121b58461216a565b612150565b9050828152602081018484840111156121d6576121d56120de565b5b6121e184828561219a565b509392505050565b5f82601f8301126121fd576121fc6120da565b5b813561220d8482602086016121a8565b91505092915050565b5f5f6040838503121561222c5761222b611f92565b5b5f61223985828601611fe0565b925050602083013567ffffffffffffffff81111561225a57612259611f96565b5b612266858286016121e9565b9150509250929050565b5f819050919050565b61228281612270565b82525050565b5f60208201905061229b5f830184612279565b92915050565b5f8115159050919050565b6122b5816122a1565b82525050565b5f6020820190506122ce5f8301846122ac565b92915050565b5f5f5f606084860312156122eb576122ea611f92565b5b5f6122f886828701612035565b935050602061230986828701612035565b925050604061231a86828701612035565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61235682612324565b612360818561232e565b935061237081856020860161233e565b612379816120e2565b840191505092915050565b5f6020820190508181035f83015261239c818461234c565b905092915050565b7f6e6f7468696e6720746f207769746864726177000000000000000000000000005f82015250565b5f6123d860138361232e565b91506123e3826123a4565b602082019050919050565b5f6020820190508181035f830152612405816123cc565b9050919050565b7f6c6f747465727920616c72656164792066696e697368656400000000000000005f82015250565b5f61244060188361232e565b915061244b8261240c565b602082019050919050565b5f6020820190508181035f83015261246d81612434565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6124ab82611f58565b91506124b683611f58565b92508282026124c481611f58565b915082820484148315176124db576124da612474565b5b5092915050565b7f696e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f61251660148361232e565b9150612521826124e2565b602082019050919050565b5f6020820190508181035f8301526125438161250a565b9050919050565b5f61255482611f58565b915061255f83611f58565b925082820190508082111561257757612576612474565b5b92915050565b7f6e6f7420656e6f756768206574686572000000000000000000000000000000005f82015250565b5f6125b160108361232e565b91506125bc8261257d565b602082019050919050565b5f6020820190508181035f8301526125de816125a5565b9050919050565b5f6125ef82611f58565b91506125fa83611f58565b925082820390508181111561261257612611612474565b5b92915050565b7f6c6f7474657279206e6f742066696e69736865640000000000000000000000005f82015250565b5f61264c60148361232e565b915061265782612618565b602082019050919050565b5f6020820190508181035f83015261267981612640565b9050919050565b7f68617665207061727469636970616e74730000000000000000000000000000005f82015250565b5f6126b460118361232e565b91506126bf82612680565b602082019050919050565b5f6020820190508181035f8301526126e1816126a8565b9050919050565b7f696c6c6567616c206f776e6572206665650000000000000000000000000000005f82015250565b5f61271c60118361232e565b9150612727826126e8565b602082019050919050565b5f6020820190508181035f83015261274981612710565b9050919050565b7f626164206475726174696f6e00000000000000000000000000000000000000005f82015250565b5f612784600c8361232e565b915061278f82612750565b602082019050919050565b5f6020820190508181035f8301526127b181612778565b9050919050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f6127f76127f26127ed846127b8565b6127d4565b6127c1565b9050919050565b612807816127dd565b82525050565b5f6020820190506128205f8301846127fe565b92915050565b7f6e6f7420656e6f756768207061727469636970616e74730000000000000000005f82015250565b5f61285a60178361232e565b915061286582612826565b602082019050919050565b5f6020820190508181035f8301526128878161284e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6128c582611f58565b91506128d083611f58565b9250826128e0576128df61288e565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61292282611f58565b915061292d83611f58565b92508261293d5761293c61288e565b5b828204905092915050565b5f61295282611f58565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361298457612983612474565b5b600182019050919050565b7f6465706f736974206d7573742062652067726561746572207468616e203000005f82015250565b5f6129c3601e8361232e565b91506129ce8261298f565b602082019050919050565b5f6020820190508181035f8301526129f0816129b7565b9050919050565b612a0081612270565b8114612a0a575f5ffd5b50565b5f81519050612a1b816129f7565b92915050565b5f60208284031215612a3657612a35611f92565b5b5f612a4384828501612a0d565b91505092915050565b5f819050919050565b612a66612a6182611f58565b612a4c565b82525050565b5f612a778285612a55565b602082019150612a878284612a55565b6020820191508190509392505050565b5f81519050919050565b5f81905092915050565b5f612ab582612a97565b612abf8185612aa1565b9350612acf81856020860161233e565b80840191505092915050565b5f612ae68284612aab565b91508190509291505056fea26469706673582212205c6ac159deea3a416e424ea3de6f9dd40047fcbcb311ae91dc13dd775f72b86364736f6c634300081c0033",
}

// LotteryABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryMetaData.ABI instead.
var LotteryABI = LotteryMetaData.ABI

// LotteryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LotteryMetaData.Bin instead.
var LotteryBin = LotteryMetaData.Bin

// DeployLottery deploys a new Ethereum contract, binding an instance of Lottery to it.
func DeployLottery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lottery, error) {
	parsed, err := LotteryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LotteryBin), backend)
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

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Lottery *LotteryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Lottery *LotterySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Lottery.Contract.UPGRADEINTERFACEVERSION(&_Lottery.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Lottery *LotteryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Lottery.Contract.UPGRADEINTERFACEVERSION(&_Lottery.CallOpts)
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

// ParticipantExist is a free data retrieval call binding the contract method 0x5c95bd86.
//
// Solidity: function participantExist(uint256 , address ) view returns(bool)
func (_Lottery *LotteryCaller) ParticipantExist(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "participantExist", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParticipantExist is a free data retrieval call binding the contract method 0x5c95bd86.
//
// Solidity: function participantExist(uint256 , address ) view returns(bool)
func (_Lottery *LotterySession) ParticipantExist(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Lottery.Contract.ParticipantExist(&_Lottery.CallOpts, arg0, arg1)
}

// ParticipantExist is a free data retrieval call binding the contract method 0x5c95bd86.
//
// Solidity: function participantExist(uint256 , address ) view returns(bool)
func (_Lottery *LotteryCallerSession) ParticipantExist(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Lottery.Contract.ParticipantExist(&_Lottery.CallOpts, arg0, arg1)
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

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Lottery *LotteryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Lottery *LotterySession) ProxiableUUID() ([32]byte, error) {
	return _Lottery.Contract.ProxiableUUID(&_Lottery.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Lottery *LotteryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Lottery.Contract.ProxiableUUID(&_Lottery.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _ownerFee, uint256 _duration, uint256 _ticketPrice) returns()
func (_Lottery *LotteryTransactor) Initialize(opts *bind.TransactOpts, _ownerFee *big.Int, _duration *big.Int, _ticketPrice *big.Int) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "initialize", _ownerFee, _duration, _ticketPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _ownerFee, uint256 _duration, uint256 _ticketPrice) returns()
func (_Lottery *LotterySession) Initialize(_ownerFee *big.Int, _duration *big.Int, _ticketPrice *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.Initialize(&_Lottery.TransactOpts, _ownerFee, _duration, _ticketPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _ownerFee, uint256 _duration, uint256 _ticketPrice) returns()
func (_Lottery *LotteryTransactorSession) Initialize(_ownerFee *big.Int, _duration *big.Int, _ticketPrice *big.Int) (*types.Transaction, error) {
	return _Lottery.Contract.Initialize(&_Lottery.TransactOpts, _ownerFee, _duration, _ticketPrice)
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

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lottery *LotteryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lottery *LotterySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lottery.Contract.UpgradeToAndCall(&_Lottery.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Lottery *LotteryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Lottery.Contract.UpgradeToAndCall(&_Lottery.TransactOpts, newImplementation, data)
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
	Account common.Address
	Amount  *big.Int
	Round   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0x19421268847f42dd61705778018ddfc43bcdce8517e7a630acb12f122c709481.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 indexed round)
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

// WatchBid is a free log subscription operation binding the contract event 0x19421268847f42dd61705778018ddfc43bcdce8517e7a630acb12f122c709481.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 indexed round)
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

// ParseBid is a log parse operation binding the contract event 0x19421268847f42dd61705778018ddfc43bcdce8517e7a630acb12f122c709481.
//
// Solidity: event Bid(address indexed account, uint256 amount, uint256 indexed round)
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
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
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

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Lottery *LotteryFilterer) ParseDeposit(log types.Log) (*LotteryDeposit, error) {
	event := new(LotteryDeposit)
	if err := _Lottery.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lottery contract.
type LotteryInitializedIterator struct {
	Event *LotteryInitialized // Event containing the contract specifics and raw log

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
func (it *LotteryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryInitialized)
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
		it.Event = new(LotteryInitialized)
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
func (it *LotteryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryInitialized represents a Initialized event raised by the Lottery contract.
type LotteryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lottery *LotteryFilterer) FilterInitialized(opts *bind.FilterOpts) (*LotteryInitializedIterator, error) {

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LotteryInitializedIterator{contract: _Lottery.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lottery *LotteryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LotteryInitialized) (event.Subscription, error) {

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryInitialized)
				if err := _Lottery.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lottery *LotteryFilterer) ParseInitialized(log types.Log) (*LotteryInitialized, error) {
	event := new(LotteryInitialized)
	if err := _Lottery.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// LotteryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Lottery contract.
type LotteryUpgradedIterator struct {
	Event *LotteryUpgraded // Event containing the contract specifics and raw log

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
func (it *LotteryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryUpgraded)
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
		it.Event = new(LotteryUpgraded)
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
func (it *LotteryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryUpgraded represents a Upgraded event raised by the Lottery contract.
type LotteryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lottery *LotteryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LotteryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LotteryUpgradedIterator{contract: _Lottery.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lottery *LotteryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LotteryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryUpgraded)
				if err := _Lottery.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Lottery *LotteryFilterer) ParseUpgraded(log types.Log) (*LotteryUpgraded, error) {
	event := new(LotteryUpgraded)
	if err := _Lottery.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
