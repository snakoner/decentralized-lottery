// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const contractParams = {
  ownerFee: 2,
  duration: 60 * 60 * 24,
  ticketPrice: 1000000000000000, // 10 ^ 15 wei = 0.001 eth
};

const DecentralizedLottery = buildModule("DecentralizedLottery", (m) => {
	const lotteryContract = m.contract("DecentralizedLottery", 
		[
		contractParams.ownerFee,
		contractParams.duration,
		contractParams.ticketPrice
	], {});

	return { lotteryContract };
});

export default DecentralizedLottery;