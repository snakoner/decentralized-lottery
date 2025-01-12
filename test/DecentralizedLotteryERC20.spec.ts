import { bigint } from "hardhat/internal/core/params/argumentTypes";
import { loadFixture, ethers, expect } from "./setup";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/src/signers";
import { token } from "../typechain-types/@openzeppelin/contracts";
const { upgrades } = require("hardhat");

const contractDeployData = {
	ticketPrice: BigInt(100), // in tokens
	duration: 60 * 60 * 25, // 25 hours
	ownerCommission: 2,
	feePaidDelta: ethers.parseEther("4.0"),	
};

function getIndexOfAccount(accounts: User[], account: string) {
	for (let i = 0; i < accounts.length; i++) {
		if (accounts[i].signer.address === account) {
			return i;
		}
	}

	return -1;
}

function getRandomInt(min: number, max: number): number {
	min = Math.ceil(min);
	max = Math.floor(max);

	return Math.floor(Math.random() * (max - min + 1)) + min;
}

interface User {
	signer: HardhatEthersSigner;
	ticketsNumber: bigint;
};

describe("DecentralizedLotteryERC20 contract", function() {
	/*
		helpers
	*/
	const getImplementation = async(proxy: any) => {
		const implSlot = "0x360894A13BA1A3210667C828492DB98DCA3E2076CC3735A920A3CA505D382BBC";
		const implementationAddress = await ethers.provider.getStorage(await proxy.getAddress(), implSlot);

		return "0x" + implementationAddress.slice(2 + 24,);
	}

	const getTotalWeight = (users: User[]) => {
		let number = BigInt(0);
		for (const user of users) {
			number += user.ticketsNumber;
		}

		return number;
	}

	async function deploy() {
		const owner = (await ethers.getSigners())[0];
		const signers = (await ethers.getSigners()).slice(1,);
		
		// ERC-20 token
		const tokenFactory = await ethers.getContractFactory("Token");
		const token = await tokenFactory.deploy("My Token", "Token");
		await token.waitForDeployment();

		// UUPS proxy lottery
		const Factory = await ethers.getContractFactory("DecentralizedLotteryERC20");
		const proxy = await upgrades.deployProxy(
			Factory,
			[
				await token.getAddress(),
				contractDeployData.ownerCommission,
				contractDeployData.duration,
				contractDeployData.ticketPrice,
			],
			{
				initializer: "_initialize",
			}
		);

		const contract = Factory.attach(await proxy.getAddress());
		await proxy.waitForDeployment();

		// for future use: create users
		let users: User[] = [];
		for (const s of (await ethers.getSigners()).slice(1,)) {
			const user: User = {
				signer: s,
				ticketsNumber: ethers.toBigInt(getRandomInt(1, 20)),
			};
			users.push(user);
		}

		console.log('contract: ', await contract.getAddress());

		// for future use: mint tokens to users
		const tokensToMintPerAcc = 100000;
		for (const user of users) {
			await token.mint(await user.signer.getAddress(), tokensToMintPerAcc);
			await token.connect(user.signer).approve(await proxy.getAddress(), tokensToMintPerAcc);
			expect(await token.balanceOf(await user.signer.getAddress())).to.be.eq(tokensToMintPerAcc);
		}

		console.log('totalSupply: ', await token.totalSupply());

		return {contract, token, owner, users};
	}

	it ("should bid", async function() {
		const {contract, users} = await loadFixture(deploy);

		for (const user of users) {
			const value = contractDeployData.ticketPrice * user.ticketsNumber;
			await contract.connect(user.signer).bid(user.ticketsNumber, { value: value});
			expect(
				await contract.weights(0, user.signer.address)
			).to.be.eq(user.ticketsNumber);

			expect(
				await contract.participantExist(0, user.signer.address)
			).to.be.eq(true);
		}

		expect(
			await contract.totalWeight(0)
		).to.be.eq(getTotalWeight(users));

		expect(
			await contract.getParticipantsNumber()
		).to.be.eq(users.length);

		// check events
		const eventAbi = ["event Bid(address indexed account, uint amount, uint64 indexed round)",];
		const eventContract = new ethers.Contract(await contract.getAddress(), eventAbi, ethers.provider);
		const events = await eventContract.queryFilter(eventContract.filters.Bid(), 0, "latest");

		events.forEach((event) => {
			const e = ({
				account: event.args?.account,
				amount: event.args?.amount,
			}
		);
		
		expect(e.amount).to.be.eq(users[getIndexOfAccount(users, e.account)].ticketsNumber);
		});
	});

	it ("should bid more than once", async function() {
		const {contract, owner, users} = await loadFixture(deploy);

		for (const user of users) {
			const value = contractDeployData.ticketPrice * user.ticketsNumber;
			await contract.connect(user.signer).bid(user.ticketsNumber, { value: value});
			await contract.connect(user.signer).bid(user.ticketsNumber, { value: value});
			expect(
				await contract.weights(0, user.signer.address)
			).to.be.eq(user.ticketsNumber * BigInt(2));
		}
	});

	it ("should start", async function() {
		const {contract, token, users} = await loadFixture(deploy);

		for (const user of users) {
			const value = contractDeployData.ticketPrice * user.ticketsNumber;
			await contract.connect(user.signer).bid(user.ticketsNumber, { value: value});
		}

		await ethers.provider.send("evm_increaseTime", [60 * 60 * 26]);
		await ethers.provider.send("evm_mine")

		await contract.start();
		expect(await contract.round()).to.be.eq(1);

		// get event
		const eventAbi = ["event WinnerSelected(address indexed account, uint256 amount, uint64 indexed round)",];
		const eventContract = new ethers.Contract(await contract.getAddress(), eventAbi, ethers.provider);
		const events = await eventContract.queryFilter(eventContract.filters.WinnerSelected(), 0, "latest");
		expect(events.length).to.be.eq(1);
		
		const address = events[0].args[0];
		const reward = events[0].args[1];
		const index = getIndexOfAccount(users, address);

		expect(await contract.balances(address)).to.be.eq(reward);
		
		// check withdraw
		const initialBalance = await token.balanceOf(address);

		await contract.connect(users[index].signer).withdraw(reward);

		const finalBalance = await token.balanceOf(address);
		expect(finalBalance - initialBalance).to.be.eq(reward);
	});

	it ("should restart", async function() {
		const {contract, users} = await loadFixture(deploy);

		await ethers.provider.send("evm_increaseTime", [60 * 60 * 26]);
		await ethers.provider.send("evm_mine")

		expect(await contract.getTimeLeft()).to.be.eq(0);

		await contract.restartEmpty(0);

		expect(await contract.round()).to.be.eq(0);
		expect(await contract.getTimeLeft()).not.to.be.eq(0);
	});    
})
