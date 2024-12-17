// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./IDecentralizedLottery.sol";

/**
 * @title DecentralizedLottery
 * @dev A lottery smart contract where participants buy tickets, and a random winner is selected.
 * The owner collects a fee, and the winner receives the prize pool.
 */
contract DecentralizedLottery is IDecentralizedLottery, Ownable {
    // Constants
    uint constant MAX_OWNER_FEE = 2; // Maximum allowable owner fee in percentage.
    uint constant MIN_DURATION = 1 days; // Minimum duration for a lottery round.

    // State variables
    uint public endTime; // Timestamp when the current lottery round ends.
    uint public duration; // Duration of a single lottery round.
    uint public round; // Current lottery round number.
    uint public totalBid; // Total amount of ETH collected in the current round.
    uint immutable public ticketPrice; // Price of a single lottery ticket.
    uint immutable public ownerFee; // Fee percentage taken by the owner.
    uint public participantsNum; // Number of unique participants in the current round.
    address[] public participants; // Array of participant addresses for the current round.

    mapping (address => uint) public balances; // Tracks withdrawable balances of users (winner/owner).
    mapping (address => mapping(uint => uint)) ticketNum; // Tracks the number of tickets per participant per round.

    /**
     * @dev Contract constructor initializes lottery parameters.
     * @param _ownerFee Fee percentage taken by the owner (must not exceed MAX_OWNER_FEE).
     * @param _duration Duration of each lottery round (must not be less than MIN_DURATION).
     * @param _ticketPrice Price of a single lottery ticket.
     */
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

    // Modifiers
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

    /**
     * @dev Allows users to buy lottery tickets by sending ETH.
     * @param amount Number of tickets the user wants to purchase.
     */
    function bid(uint amount) 
        external 
        payable 
        enoughEthersSent(amount) 
        lotteryNotFinished 
    {
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

    /**
     * @dev Restarts a new lottery round with a given duration if the current one is empty.
     * @param newDuration Optional new duration for the next round (defaults to current duration).
     */
    function restartEmpty(uint newDuration) 
        external 
        onlyOwner 
        lotteryFinished 
    {
        require(participants.length == 0, "have participants");

        duration = newDuration != 0 ? newDuration : duration;
        endTime = block.timestamp + duration;
    }

    /**
     * @dev Generates a random number using keccak256 (not truly random, for demonstration purposes).
     * @return Random number.
     */
    function generateRandom() private view returns (uint) {
        return uint(keccak256(abi.encodePacked(block.prevrandao, block.timestamp)));
    }

    /**
     * @dev Selects a winner, calculates rewards, and starts a new lottery round.
     */
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
        participantsNum = 0;

        delete participants;
    }

    /**
     * @dev Allows a user to withdraw their available balance.
     */
    function withdraw() external {
        require(balances[msg.sender] > 0, "nothing to withdraw");

        uint value = balances[msg.sender];

        balances[msg.sender] = 0;
        payable(msg.sender).transfer(value);

        emit Withdraw(msg.sender, value);
    }

    /**
     * @dev Gets the number of tickets a user purchased in the current round.
     * @param account Address of the user.
     * @return Number of tickets.
     */
    function getTicketNum(address account) external view returns (uint) {
        return ticketNum[account][round];
    }

    /**
     * @dev Gets the withdrawable balance of a user.
     * @param account Address of the user.
     * @return Withdrawable balance.
     */
    function getUnlockedBalance(address account) external view returns (uint) {
        return balances[account];
    }

    /**
     * @dev Gets the time remaining for the current round.
     * @return Time remaining in seconds, or 0 if the round has ended.
     */
    function getTimeLeft() external view returns (uint) {
        if (endTime < block.timestamp) {
            return 0;
        }

        return endTime - block.timestamp;
    }
}