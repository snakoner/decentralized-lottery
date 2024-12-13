import { ethers as hardhatEthers } from "hardhat";
import { parseEther } from "ethers"; // Direct import for ethers@6.x

async function main() {
    const [deployer] = await hardhatEthers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);

    // Constructor arguments
    const ownerFee = 2; // Maximum owner fee (percentage)
    const duration = 3 * 24 * 60 * 60; // 3 days in seconds
    const ticketPrice = parseEther("0.01"); // Converts 0.01 ETH to Wei

    // Deploy the DecentralizedLottery contract
    const DecentralizedLottery = await hardhatEthers.getContractFactory("DecentralizedLottery");
    const decentralizedLottery = await DecentralizedLottery.deploy(ownerFee, duration, ticketPrice);

    // The deployed contract address
    console.log("DecentralizedLottery deployed to:", decentralizedLottery.target || decentralizedLottery.address);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
