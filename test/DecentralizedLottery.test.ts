import { loadFixture, ethers, expect } from "./setup";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/src/signers";

const _ticketPrice = ethers.parseEther("0.1");
const _duration = 60 * 60 * 25; // 25 hours
const _ownerCommission = 2;

async function getRawLogs(contractAddress: string) {
    const filter = {
        address: contractAddress,
        topics: [
			ethers.id("event WinnerSelected(address indexed account, uint amount, uint round)"),
			ethers.id("event Bid(address indexed account, uint amount, uint timestamp, uint round)"),
        ],
        fromBlock: 0,
        toBlock: "latest",
    };

    const logs = await ethers.provider.getLogs(filter);
	console.log("getRaw: ", logs.length)

    logs.forEach((log) => {
        console.log("Raw log:", log);
    });
}

async function handleWinnerSelected(account: string, amount: bigint, round: bigint) {
	console.log("WWWWW")
}

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
			_ownerCommission, 
			_duration,
			_ticketPrice);

		await contract.waitForDeployment();

        return {contract, owner, parts};
    }

    it ("should be depoyed", async function() {
        const {contract} = await loadFixture(deploy);
        expect(contract.target).to.be.properAddress;
    });

    it ("contract logic", async function() {
        const {contract, owner, parts} = await loadFixture(deploy);
		let participants: Participant[] = [];
		let addedAddressed = new Map<string, number>();
		let i = 0;
		let totalParticipants = 0;

		for (const part of parts) {
			const _part = new Participant(part, Math.floor(Math.random() * (10 + 1)) + 1);
			participants.push(_part);
			const weiValue = ethers.toBigInt(_part.ticketsNumber) * _ticketPrice;

			await contract.connect(_part.signer).bid(_part.ticketsNumber, {
				value: weiValue,
			})

			totalParticipants += _part.ticketsNumber;
			// console.log(_part.getAddress(), _part.ticketsNumber)
		}

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
		})

		expect(await contract.getTimeLeft()).not.to.be.eq(0)

		// 26 hours
		await ethers.provider.send("evm_increaseTime", [60 * 60 * 26]);
		await ethers.provider.send("evm_mine")

		console.log("participants: ", await contract.participantsNum());

		expect(await contract.getTimeLeft()).to.be.eq(0)

		const tx = await contract.start()
		await tx.wait()
		
		const _contract = new ethers.Contract(
			await contract.getAddress(), 
			["event WinnerSelected(address indexed account, uint amount, uint round)"],
			ethers.provider)

		_contract.on("WinnerSelected", handleWinnerSelected)
	});
})