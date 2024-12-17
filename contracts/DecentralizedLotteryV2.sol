// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./IDecentralizedLottery.sol";

/**
 * @title DecentralizedLottery
 * @dev A lottery smart contract where participants buy tickets, and a random winner is selected.
 * The owner collects a fee, and the winner receives the prize pool.
 */
contract DecentralizedLotteryV2 is Ownable, IDecentralizedLottery {
    // Constants
    uint constant MAX_OWNER_FEE = 2; // Maximum allowable owner fee in percentage.
    uint constant MIN_DURATION = 1 days; // Minimum duration for a lottery round.

    // State variables
    uint public endTime; // Timestamp when the current lottery round ends.
    uint public duration; // Duration of a single lottery round.
    uint public round; // Current lottery round number.
    uint immutable public ticketPrice; // Price of a single lottery ticket.
    uint immutable public ownerFee; // Fee percentage taken by the owner.
    mapping (address => uint) public balances; // Tracks withdrawable balances of users (winner/owner).
    mapping (uint => mapping (address => uint)) weights;
    mapping (uint => mapping (address => bool)) participantExist;
    mapping (uint => uint) public totalWeight;
    address[] public participants;

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

    modifier validRound(uint _round) {
        require(_round <= round, "invalid round");
        _;
    }

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


    /**
     * @dev Allows users to buy lottery tickets by sending ETH.
     * @param amount Number of tickets the user wants to purchase.
     */
    function bid(uint amount) external payable enoughEthersSent(amount) lotteryNotFinished {
        uint refund = msg.value - amount * ticketPrice;
        if (refund > 0) {
            payable(msg.sender).transfer(refund);
        }

        if (!participantExist[round][msg.sender]) {
            participants.push(msg.sender);
        }

        totalWeight[round] += amount;
        weights[round][msg.sender] += amount;

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

        uint randomValue = generateRandom() % totalWeight[round]; 
        uint cumulativeWeight = 0;
        address winner;

        for (uint i = 0; i < participants.length; i++) {
            cumulativeWeight += weights[round][participants[i]];
            if (randomValue < cumulativeWeight) {
                winner = participants[i];
                break;
            }
        }

        uint totalBid = totalWeight[round] * ticketPrice;
        uint fee = (totalBid * ownerFee) / 100;
        uint reward = totalBid - fee;
        
        balances[owner()] += fee;
        balances[winner] += reward;

        emit WinnerSelected(winner, reward, round);

        round++;
        endTime = block.timestamp + duration;

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
     * @dev Gets the number of participants in current round.
     * @return Number of participants.
     */
    function getParticipantsNumber() external view returns (uint) {
        return participants.length;
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

    // function getTotalWeight(uint _round) external validRound(_round) view returns (uint) {
    //     return totalWeight[_round];
    // }
}