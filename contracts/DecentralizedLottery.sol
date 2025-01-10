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
    // State variables
    uint64 public endTime; // Timestamp when the current lottery round ends.
    uint64 public duration; // Duration of a single lottery round.
    uint64 public round; // Current lottery round number.
    uint64 public ownerFee; // Fee percentage taken by the owner.
    uint256 public ticketPrice; // Price of a single lottery ticket.
    mapping (address => uint256) public balances; // Tracks withdrawable balances of users (winner/owner).
    mapping (uint256 => mapping (address => uint256)) public weights;
    mapping (uint256 => mapping (address => bool)) public participantExist;
    mapping (uint256 => uint256) public totalWeight;
    address[] public participants;

    // Modifiers
    modifier enoughEthersSent(uint256 amount) {
        require(msg.value == (amount * ticketPrice), NotEnoughtEther());
        _;
    }

    modifier lotteryNotFinished() {
        require(block.timestamp < endTime, LotteryAlreadyFinished());
        _;
    }

    modifier lotteryFinished() {
        require(block.timestamp >= endTime, LotteryNotFinished());
        _;
    }

    modifier validRound(uint256 _round) {
        require(_round <= round, InvalidRoundNumber());
        _;
    }

    /**
     * @dev Contract constructor initializes lottery parameters.
     * @param _ownerFee Fee percentage taken by the owner (must not exceed MAX_OWNER_FEE).
     * @param _duration Duration of each lottery round (must not be less than MIN_DURATION).
     * @param _ticketPrice Price of a single lottery ticket.
     */
    function initialize(
        uint64 _ownerFee, 
        uint64 _duration, 
        uint256 _ticketPrice
    ) external initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();

        ownerFee = _ownerFee;
        ticketPrice = _ticketPrice;
        duration = _duration;
        endTime = uint64(block.timestamp) + _duration;
    }


    /**
     * @dev Allows users to buy lottery tickets by sending ETH.
     * @param amount Number of tickets the user wants to purchase.
     */
    function bid(uint256 amount) external payable enoughEthersSent(amount) lotteryNotFinished {
        address account = msg.sender;
        if (!participantExist[round][account]) {
            participantExist[round][account] = true;
            participants.push(account);
        }

        totalWeight[round] += amount;
        weights[round][account] += amount;

        emit Bid(account, amount, round);
    }

    /**
     * @dev Restarts a new lottery round with a given duration if the current one is empty.
     * @param newDuration Optional new duration for the next round (defaults to current duration).
     */
    function restartEmpty(uint64 newDuration) 
        external 
        onlyOwner 
        lotteryFinished {
        require(participants.length == 0, HaveParticipants());

        duration = newDuration != 0 ? newDuration : duration;
        endTime = uint64(block.timestamp) + duration;
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
        require(participants.length > 0, NotEnoughParticipants()); 

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
        endTime = uint64(block.timestamp) + duration;

        delete participants;
    }

    /**
     * @dev Allows a user to withdraw their available balance.
     */
    function withdraw(uint256 amount) external {
        address account = msg.sender;
        require(balances[account] >= amount && amount > 0, NothingToWithdraw());

        unchecked {
            balances[account] -= amount;
        }

        payable(account).transfer(amount);

        emit Withdraw(account, amount);
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
        address account = msg.sender;
        require(balances[account] >= value, NotEnoughtEther());
        
        unchecked {
            balances[account] -= value;
        }

        if (!participantExist[round][account]) {
            participantExist[round][account] = true;
            participants.push(account);
        }

        totalWeight[round] += amount;
        weights[round][account] += amount;

        emit Bid(account, amount, round);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}