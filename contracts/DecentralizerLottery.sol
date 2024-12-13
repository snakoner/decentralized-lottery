// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./IDecentralizedLottery.sol";

contract DecentralizedLottery is IDecentralizerLottery, Ownable {
    uint constant MAX_OWNER_FEE = 2; // percents
    uint constant MIN_DURATION = 1 days;
    uint public endTime;
    uint public duration;
    uint public round;
    uint public totalBid;
    uint immutable public ticketPrice;
    uint immutable public ownerFee;
    uint public participantsNum;   // in current round
    address[] public participants;
    mapping (address => uint) public balances;  // account => balance to withdraw
    mapping (address => mapping(uint => uint)) ticketNum; // account => (round => ticketNum);

    constructor(
        uint _ownerFee, 
        uint _duration, 
        uint _ticketPrice
        ) Ownable(msg.sender) {
        require(_ownerFee <= MAX_OWNER_FEE, "illegal owner fee");
        require(_duration >= MIN_DURATION, "bad duration");

        ownerFee = _ownerFee;
        ticketPrice = _ticketPrice;
        duration = _duration;
        endTime = block.timestamp + _duration;
    }

    modifier enoughEthersSent(uint amount) {
        require(msg.value >= amount * ticketPrice, "not enough ether");
        _;
    }

    modifier lotteryNotFinished() {
        require(block.timestamp < endTime, "lottery already finished");
        _;
    }

    modifier lotteryFinished() {
        require(block.timestamp >= endTime, "lottery not finished");
        _;
    }

    function bid(uint amount) external payable enoughEthersSent(amount) lotteryNotFinished {
        uint refund = msg.value - amount * ticketPrice;
        if (refund > 0) {
            payable(msg.sender).transfer(refund);
        }

        for (uint i = 0; i < amount; i++) {
            participants.push(msg.sender);
        }

        if (ticketNum[msg.sender][round] == 0) {
            participantsNum++;
        }

        ticketNum[msg.sender][round] += amount;
        totalBid += amount * ticketPrice;

        emit Bid(msg.sender, amount, block.timestamp, round);
    }

    function restartEmpty(uint newDuration) external onlyOwner lotteryFinished {
        require(participants.length == 0, "have participants");
    
        duration = newDuration != 0 ? newDuration : duration;
        endTime = block.timestamp + duration;
    }

    // todo chainlink vrf
    function generateRandom() private view returns (uint) {
        return uint(keccak256(abi.encodePacked(block.prevrandao, block.timestamp)));
    }

    function start() external onlyOwner lotteryFinished {
        require(participants.length > 0, "not enough participants");

        uint winnerIdx = generateRandom() % participants.length;
        address winner = participants[winnerIdx];
        uint fee = (totalBid * ownerFee) / 100;
        uint reward = totalBid - fee;
        
        balances[owner()] += fee;
        balances[winner] += reward;
        totalBid = 0;

        emit WinnerSelected(winner, reward, round);    

        round++;
        endTime = block.timestamp + duration;
        delete participants;
    }

    function withdraw(address payable _to) external {
        require(balances[msg.sender] > 0, "nothing to withdraw");
        uint value = balances[msg.sender];

        balances[msg.sender] = 0;
        _to.transfer(value);

        emit Withdraw(msg.sender, _to, value);
    }

    function getTicketNum(address account) external view returns (uint) {
        return ticketNum[account][round];
    }

    function getUnlockedBalance(address account) external view returns (uint) {
        return balances[account];
    }

    function getTimeLeft() external view returns (uint) {
        // time expired, can call start()
        if (endTime < block.timestamp) {
            return 0;
        }

        return endTime - block.timestamp;
    }
}