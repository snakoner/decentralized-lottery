import { loadFixture, ethers, expect } from "./setup";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/src/signers";

const contractDeployData = {
	ticketPrice: ethers.parseEther("1.0"),
	duration: 60 * 60 * 25, // 25 hours
	ownerCommission: 2,
	feePaidDelta: ethers.parseEther("0.01"),	
};

function getSignerByAddress(address: string, signers: HardhatEthersSigner[]) : HardhatEthersSigner {
	let ret: HardhatEthersSigner = signers[0];
	for (const s of signers) {
		if (s.address == address) {
			ret = s;
			break;
		}
	}

	return ret;
}

class Participant {
	signer: HardhatEthersSigner;
	ticketsNumber: number;

	constructor(signer: HardhatEthersSigner, ticketsNumber: number) {
		this.signer = signer;
		this.ticketsNumber = ticketsNumber;
	}

	getAddress(): string {
		return this.signer.address;
	}
};

describe("DecentralizedLottery contract", function() {
	    async function deploy() {
		const owner = (await ethers.getSigners())[0];
		const parts = (await ethers.getSigners()).slice(1, 10);

		console.log(ethers.formatUnits(await ethers.provider.getBalance(owner.address)));
        const Factory = await ethers.getContractFactory("DecentralizedLottery");
        const contract = await Factory.deploy(
			contractDeployData.ownerCommission,
			contractDeployData.duration,
			contractDeployData.ticketPrice
		);

		await contract.waitForDeployment();

        return {contract, owner, parts};
    }

    it ("should be depoyed", async function() {
        const {contract} = await loadFixture(deploy);
        expect(contract.target).to.be.properAddress;
    });

    it ("should be finished", async function() {
        const {contract, owner, parts} = await loadFixture(deploy);
		let participants: Participant[] = [];
		let addedAddressed = new Map<string, number>();
		let totalParticipants = 0;

		// 1. participants should be buy tickets
		for (const participant of parts) {
			const _part = new Participant(participant, Math.floor(Math.random() * (10 + 1)) + 1);
			participants.push(_part);
			const weiValue = ethers.toBigInt(_part.ticketsNumber) * contractDeployData.ticketPrice;

			await contract.connect(_part.signer).bid(_part.ticketsNumber, {
				value: weiValue,
			})

			totalParticipants += _part.ticketsNumber;
		}


		let totalTickets = 0;
		for (let i = 0; i < totalParticipants; i++) {
			const address = await contract.participants(i);
			const value = addedAddressed.get(address)
			if (value === undefined) {
				addedAddressed.set(address, 1)
			} else {
				addedAddressed.set(address, value + 1)
			}
		}

		
		addedAddressed.forEach((value, key) => {
			expect(value).to.be.eq(addedAddressed.get(key))
			totalTickets += value;
		})

		// check that participant num is eq to number of participants
		expect(await contract.participantsNum()).to.be.eq(participants.length);
		// check that time is not over
		expect(await contract.getTimeLeft()).not.to.be.eq(0)

		// 2. should be start lottery
		await ethers.provider.send("evm_increaseTime", [60 * 60 * 26]);
		await ethers.provider.send("evm_mine")

		const tx = await contract.start()
		await tx.wait()

		const getWinnerSelectedEvent = async () => {
			const eventFilter = contract.filters.WinnerSelected();
			const events = await contract.queryFilter(eventFilter, 0, "latest");
			const winnerAddress = events[0].args[0];
			const rewardWei = events[0].args[1];
			const round = events[0].args[2];
			
			return {winnerAddress, rewardWei, round};
		}

		const {winnerAddress, rewardWei, round} = await getWinnerSelectedEvent();

		const winner = getSignerByAddress(winnerAddress, parts)
		const winnerBalanceBefore = await ethers.provider.getBalance(winner.address);

		await contract.connect(winner).withdraw(winner.address);

		const winnerBalanceAfter = await ethers.provider.getBalance(winner.address);
		const winnerBalanceChange = winnerBalanceAfter - winnerBalanceBefore;
		const ownerFee = ethers.toBigInt(totalTickets) * contractDeployData.ticketPrice * (await contract.ownerFee()) / ethers.toBigInt(100);
		const totalSum = ownerFee + winnerBalanceChange;

		expect(totalSum).to.be.greaterThan(ethers.toBigInt(totalTickets) * contractDeployData.ticketPrice - contractDeployData.feePaidDelta);

		const ownerBalanceBefore = await ethers.provider.getBalance(owner.address);
		await contract.withdraw(owner.address);
		const ownerBalanceAfter = await ethers.provider.getBalance(owner.address);

		expect(ownerBalanceAfter - ownerBalanceBefore).to.be.greaterThan(ownerFee - contractDeployData.feePaidDelta);

		// check that new round started
		expect(await contract.round()).to.be.eq(1);
		expect(await contract.participantsNum()).to.be.eq(0);
	});
})
