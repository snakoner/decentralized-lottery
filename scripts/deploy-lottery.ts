const { ethers, upgrades } = require("hardhat");

const contractParams = {
    ownerFee: 2,
    duration: 60 * 60 * 24,
    ticketPrice: 100000000000000, // 10 ^ 15 wei = 0.0001 eth
};
  
async function main() {
    const lottery = await ethers.getContractFactory("DecentralizedLottery");

    console.log("Deploying contract...");
    const proxy = await upgrades.deployProxy(lottery, [		
        contractParams.ownerFee,
		contractParams.duration,
		contractParams.ticketPrice
    ], {
        initializer: "initialize",
    });

    console.log("DecentralizedLottery:", await proxy.getAddress());
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});