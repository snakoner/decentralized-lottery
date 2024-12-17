// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const contractParams = {
  ownerFee: 2,
  duration: 60 * 60 * 24,
  ticketPrice: 100000000000000, // 10 ^ 15 wei = 0.0001 eth
};

const DecentralizedLottery = buildModule("DecentralizedLotteryV2", (m) => {
	const lotteryContract = m.contract("DecentralizedLotteryV2", 
		[
		contractParams.ownerFee,
		contractParams.duration,
		contractParams.ticketPrice
	], {});

	return { lotteryContract };
});

export default DecentralizedLottery;
