// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./IDecentralizedLottery.sol";

contract DecentralizedLotteryV2 is Ownable, IDecentralizedLottery {
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
    address[] public participants;
    Round[] public rounds;

    struct Round {
        uint timeFinished;
        uint totalWeight;
        uint totalParticipants;
        address winner;
        bool finished;
    }

    // Modifiers
    modifier enoughEthersSent(uint amount) {
        require(msg.value >= amount * ticketPrice, "not enough ether");
        _;
    }

    // We have got endTime state variable, which shows if current round finished
    modifier lotteryNotFinished() {
        require(block.timestamp < endTime, "lottery already finished");
        _;
    }

    modifier lotteryFinished() {
        require(block.timestamp >= endTime, "lottery not finished");
        _;
    }

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

        _createNewRound();
    }

    function _createNewRound() internal {
        rounds.push(Round({
            timeFinished: 0,
            totalWeight: 0,
            totalParticipants: 0,
            winner: address(0),
            finished: false
        }));
    }

    function bid(uint amount) external payable enoughEthersSent(amount) lotteryNotFinished {
        uint refund = msg.value - amount * ticketPrice;
        if (refund > 0) {
            payable(msg.sender).transfer(refund);
        }

        Round storage currentRound = rounds[round];

        if (!participantExist[round][msg.sender]) {
            participants.push(msg.sender);
            currentRound.totalParticipants++;
        }

        currentRound.totalWeight += amount;
        weights[round][msg.sender] += amount;

        emit Bid(msg.sender, amount, block.timestamp, round);
    }

    function restartEmpty(uint newDuration) 
        external 
        onlyOwner 
        lotteryFinished 
    {
        require(rounds[round].totalParticipants == 0, "have participants");

        duration = newDuration != 0 ? newDuration : duration;
        endTime = block.timestamp + duration;
    }

    function generateRandom() private view returns (uint) {
        return uint(keccak256(abi.encodePacked(block.prevrandao, block.timestamp)));
    }

    function start() external onlyOwner lotteryFinished {
        require(rounds[round].totalParticipants > 0, "not enough participants"); 

        Round storage currentRound = rounds[round];
        uint randomValue = generateRandom() % participants.length; 
        uint cumulativeWeight = 0;

        for (uint i = 0; i < participants.length; i++) {
            cumulativeWeight += weights[round][participants[i]];
            if (randomValue < cumulativeWeight) {
                currentRound.winner = participants[i];
                currentRound.finished = true;
                currentRound.timeFinished = block.timestamp;
                break;
            }
        }

        uint totalBid = currentRound.totalWeight * ticketPrice;
        uint fee = (totalBid * ownerFee) / 100;
        uint reward = totalBid - fee;
        
        balances[owner()] += fee;
        balances[currentRound.winner] += reward;

        emit WinnerSelected(currentRound.winner, reward, round);

        _createNewRound();
        round++;
        endTime = block.timestamp + duration;

        delete participants;
    }

    function withdraw() external {
        require(balances[msg.sender] > 0, "nothing to withdraw");

        uint value = balances[msg.sender];

        balances[msg.sender] = 0;
        payable(msg.sender).transfer(value);

        emit Withdraw(msg.sender, value);
    }

    function getParticipantsNumber() external view returns (uint) {
        return participants.length;
    }
}