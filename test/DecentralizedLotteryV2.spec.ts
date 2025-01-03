import { loadFixture, ethers, expect } from "./setup";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/src/signers";

const contractDeployData = {
    ticketPrice: ethers.parseEther("1.0"),
    duration: 60 * 60 * 25, // 25 hours
    ownerCommission: 2,
    feePaidDelta: ethers.parseEther("4.0"),	
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

describe("DecentralizedLotteryV2 contract", function() {
        async function deploy() {
        const owner = (await ethers.getSigners())[0];
        const parts = (await ethers.getSigners()).slice(1, 10);
        const Factory = await ethers.getContractFactory("DecentralizedLotteryV2");
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
        let totalWeight = 0;

        // 1. participants should be able to bid()
        for (const participant of parts) {
            const _part = new Participant(participant, Math.floor(Math.random() * (10 + 1)) + 1);
            participants.push(_part);
            const weiValue = ethers.toBigInt(_part.ticketsNumber) * contractDeployData.ticketPrice;

            await contract.connect(_part.signer).bid(_part.ticketsNumber, {
                value: weiValue,
            });
            
            totalWeight += _part.ticketsNumber;
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

        // check initial state
        let weight = await contract.totalWeight(0);
        let partNumber = await contract.getParticipantsNumber();

        expect(weight).to.be.eq(totalWeight);
        expect(partNumber).to.be.eq(parts.length);
        expect(await contract.getParticipantsNumber()).to.be.eq(parts.length);

        // 2. should be start lottery
        await ethers.provider.send("evm_increaseTime", [60 * 60 * 26]);
        await ethers.provider.send("evm_mine")

        // 3. try to bid() if time is over
        try {
            const weiValue = ethers.toBigInt(participants[0].ticketsNumber) * contractDeployData.ticketPrice;
            await contract.connect(participants[0].signer).bid(participants[0].ticketsNumber, {
                value: weiValue,
            });

            expect(false).to.be.eq(true);
        } catch {}

        // 4. try to start
        const tx = await contract.start()
        await tx.wait()

        // 5. withdraw
        const getWinnerSelectedEvent = async () => {
            const eventFilter = contract.filters.WinnerSelected();
            const events = await contract.queryFilter(eventFilter, 0, "latest");
            const winnerAddress = events[0].args[0];
            const rewardWei = events[0].args[1];
            const round = events[0].args[2];
            
            return {winnerAddress, rewardWei, round};
        }

        const {winnerAddress} = await getWinnerSelectedEvent();

        const winner = getSignerByAddress(winnerAddress, parts)
        const winnerBalanceBefore = await ethers.provider.getBalance(winner.address);

        // 6. bid from balance
        await contract.connect(winner).bidFromBalance(2);
        const ticketWeight = await contract.weights(ethers.toBigInt(await contract.round()), winner.address);

        const _nextRound = await contract.round();
        expect(_nextRound).to.be.eq(1);

        weight = await contract.totalWeight(_nextRound);
        expect(weight).to.be.eq(2);
        expect(await contract.getParticipantsNumber()).to.be.eq(1);

        expect(ticketWeight).to.be.eq(ethers.toBigInt(2));

        // 7. withdraw
        await contract.connect(winner).withdraw(await contract.balances(winner.address));

        const winnerBalanceAfter = await ethers.provider.getBalance(winner.address);
        const winnerBalanceChange = winnerBalanceAfter - winnerBalanceBefore;
        const ownerFee = ethers.toBigInt(totalWeight) * contractDeployData.ticketPrice * (await contract.ownerFee()) / ethers.toBigInt(100);
        const totalSum = ownerFee + winnerBalanceChange;

        expect(totalSum).to.be.greaterThan(ethers.toBigInt(totalWeight) * contractDeployData.ticketPrice - contractDeployData.feePaidDelta);

        const ownerBalanceBefore = await ethers.provider.getBalance(owner.address);
        await contract.withdraw(await contract.balances(owner.address));  
        const ownerBalanceAfter = await ethers.provider.getBalance(owner.address);

        expect(ownerBalanceAfter - ownerBalanceBefore).to.be.greaterThan(ownerFee - contractDeployData.feePaidDelta);
    });
})
