// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "./interfaces/IDecentralizedLottery.sol";

/**
 * @title DecentralizedLottery
 * @dev A lottery smart contract where participants buy tickets, and a random winner is selected.
 * The owner collects a fee, and the winner receives the prize pool.
 */
contract DecentralizedLottery is 
    Initializable, 
    UUPSUpgradeable,
    OwnableUpgradeable, 
    IDecentralizedLottery 
{
    // Constants
    uint256 constant MAX_OWNER_FEE = 2; // Maximum allowable owner fee in percentage.
    uint256 constant MIN_DURATION = 1 days; // Minimum duration for a lottery round.

    // State variables
    uint256 public endTime; // Timestamp when the current lottery round ends.
    uint256 public duration; // Duration of a single lottery round.
    uint256 public round; // Current lottery round number.
    uint256 public ticketPrice; // Price of a single lottery ticket.
    uint256 public ownerFee; // Fee percentage taken by the owner.
    mapping (address => uint256) public balances; // Tracks withdrawable balances of users (winner/owner).
    mapping (uint256 => mapping (address => uint256)) public weights;
    mapping (uint256 => mapping (address => bool)) public participantExist;
    mapping (uint256 => uint256) public totalWeight;
    address[] public participants;

    // Modifiers
    modifier enoughEthersSent(uint256 amount) {
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

    modifier validRound(uint256 _round) {
        require(_round <= round, "invalid round");
        _;
    }

    /**
     * @dev Contract constructor initializes lottery parameters.
     * @param _ownerFee Fee percentage taken by the owner (must not exceed MAX_OWNER_FEE).
     * @param _duration Duration of each lottery round (must not be less than MIN_DURATION).
     * @param _ticketPrice Price of a single lottery ticket.
     */
    function initialize(
        uint256 _ownerFee, 
        uint256 _duration, 
        uint256 _ticketPrice
    ) external initializer {
        require(_ownerFee <= MAX_OWNER_FEE, "illegal owner fee");
        require(_duration >= MIN_DURATION, "bad duration");

        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();

        ownerFee = _ownerFee;
        ticketPrice = _ticketPrice;
        duration = _duration;
        endTime = block.timestamp + _duration;
    }


    /**
     * @dev Allows users to buy lottery tickets by sending ETH.
     * @param amount Number of tickets the user wants to purchase.
     */
    function bid(uint256 amount) external payable enoughEthersSent(amount) lotteryNotFinished {
        uint256 refund = msg.value - amount * ticketPrice;
        if (refund > 0) {
            payable(msg.sender).transfer(refund);
        }

        if (!participantExist[round][msg.sender]) {
            participantExist[round][msg.sender] = true;
            participants.push(msg.sender);
        }

        totalWeight[round] += amount;
        weights[round][msg.sender] += amount;

        emit Bid(msg.sender, amount, round);
    }

    /**
     * @dev Restarts a new lottery round with a given duration if the current one is empty.
     * @param newDuration Optional new duration for the next round (defaults to current duration).
     */
    function restartEmpty(uint256 newDuration) 
        external 
        onlyOwner 
        lotteryFinished {
        require(participants.length == 0, "have participants");

        duration = newDuration != 0 ? newDuration : duration;
        endTime = block.timestamp + duration;
    }

    /**
     * @dev Generates a random number using keccak256 (not truly random, for demonstration purposes).
     * @return Random number.
     */
    function generateRandom() private view returns (uint256) {
        return uint256(keccak256(abi.encodePacked(block.prevrandao, block.timestamp)));
    }

    /**
     * @dev Selects a winner, calculates rewards, and starts a new lottery round.
     */
    function start() external onlyOwner lotteryFinished {
        require(participants.length > 0, "not enough participants"); 

        uint256 randomValue = generateRandom() % totalWeight[round]; 
        uint256 cumulativeWeight = 0;
        address winner;

        for (uint256 i = 0; i < participants.length; i++) {
            cumulativeWeight += weights[round][participants[i]];
            if (randomValue < cumulativeWeight) {
                winner = participants[i];
                break;
            }
        }

        uint256 totalBid = totalWeight[round] * ticketPrice;
        uint256 fee = (totalBid * ownerFee) / 100;
        uint256 reward = totalBid - fee;
        
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
    function withdraw(uint256 amount) external {
        require(balances[msg.sender] >= amount && amount > 0, "nothing to withdraw");

        unchecked {
            balances[msg.sender] -= amount;
        }

        payable(msg.sender).transfer(amount);

        emit Withdraw(msg.sender, amount);
    }

    /**
     * @dev Gets the number of participants in current round.
     * @return Number of participants.
     */
    function getParticipantsNumber() external view returns (uint256) {
        return participants.length;
    }

    /**
     * @dev Gets the time remaining for the current round.
     * @return Time remaining in seconds, or 0 if the round has ended.
     */
    function getTimeLeft() external view returns (uint256) {
        if (endTime < block.timestamp) {
            return 0;
        }

        return endTime - block.timestamp;
    }

    /**
     * @dev Allows users to buy lottery tickets by changing internal ETH balance.
     */
    function bidFromBalance(uint256 amount) external lotteryNotFinished {
        uint256 value = amount * ticketPrice;
        require(balances[msg.sender] >= value, "insufficient balance");
        
        unchecked {
            balances[msg.sender] -= value;
        }

        if (!participantExist[round][msg.sender]) {
            participantExist[round][msg.sender] = true;
            participants.push(msg.sender);
        }

        totalWeight[round] += amount;
        weights[round][msg.sender] += amount;

        emit Bid(msg.sender, amount, round);
    }

    /**
     * @dev Allows users to deposit ETH to balance.
     */
    function deposit() external payable {
        // Check if the deposit amount is greater than 0
        require(msg.value > 0, "deposit must be greater than 0");

        // Add the deposited msg.value to the user's balance
        balances[msg.sender] += msg.value;

        emit Deposit(msg.sender, msg.value);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}