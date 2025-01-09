import { bigint } from "hardhat/internal/core/params/argumentTypes";
import { loadFixture, ethers, expect } from "./setup";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/src/signers";
const { upgrades } = require("hardhat");

const contractDeployData = {
    ticketPrice: ethers.parseEther("1.0"),
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

describe("DecentralizedLottery contract", function() {
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
        const Factory = await ethers.getContractFactory("DecentralizedLottery");
        const proxy = await upgrades.deployProxy(
            Factory,
            [contractDeployData.ownerCommission,
            contractDeployData.duration,
            contractDeployData.ticketPrice],
            {
                initializer: "initialize",
            }
        );

        const contract = Factory.attach(await proxy.getAddress());
        await proxy.waitForDeployment();

        // for future use create users
        let users: User[] = [];
        for (const s of (await ethers.getSigners()).slice(1,)) {
            const user: User = {
                signer: s,
                ticketsNumber: ethers.toBigInt(getRandomInt(1, 20)),
            };
            users.push(user);
        }

        return {contract, owner, users};
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
		const eventAbi = ["event Bid(address indexed account, uint amount, uint indexed round)",];
		const eventContract = new ethers.Contract(await contract.getAddress(), eventAbi, ethers.provider);
		const events = await eventContract.queryFilter(eventContract.filters.Bid(), 0, "latest");

        console.log(events.length);
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
        const {contract, users} = await loadFixture(deploy);

        for (const user of users) {
            const value = contractDeployData.ticketPrice * user.ticketsNumber;
            await contract.connect(user.signer).bid(user.ticketsNumber, { value: value});
        }

        await ethers.provider.send("evm_increaseTime", [60 * 60 * 26]);
        await ethers.provider.send("evm_mine")

        await contract.start();
        expect(await contract.round()).to.be.eq(1);

        // get event
		const eventAbi = ["event WinnerSelected(address indexed account, uint amount, uint indexed round)",];
		const eventContract = new ethers.Contract(await contract.getAddress(), eventAbi, ethers.provider);
		const events = await eventContract.queryFilter(eventContract.filters.WinnerSelected(), 0, "latest");
        expect(events.length).to.be.eq(1);
        
        const address = events[0].args[0];
        const reward = events[0].args[1];
        const index = getIndexOfAccount(users, address);

        expect(await contract.balances(address)).to.be.eq(reward);
        // check withdraw
        const initialBalance = await ethers.provider.getBalance(address);

        const tx = await contract.connect(users[index].signer).withdraw(reward);
        const receipt = await tx.wait();

        const finalBalance = await ethers.provider.getBalance(address);
        const transactionFee = receipt.gasUsed * (tx.gasPrice || receipt.effectiveGasPrice);
        expect(finalBalance - initialBalance).to.be.eq(reward - transactionFee);
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
