// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {DecentralizedLottery} from "./DecentralizedLottery.sol";

contract DecentralizedLotteryERC20 is DecentralizedLottery {
    using SafeERC20 for IERC20;
    IERC20 public token;

    function _initialize(
        address _token,
        uint64 _ownerFee, 
        uint64 _duration, 
        uint256 _ticketPrice
    ) public initializer {
        super.initialize(_ownerFee, _duration, _ticketPrice);
        token = IERC20(_token);
    }

    function bid(uint256 ticketsNum) public payable override lotteryNotFinished {
        address account = msg.sender;

        if (!participantExist[round][account]) {
            participantExist[round][account] = true;
            participants.push(account);
        }

        totalWeight[round] += ticketsNum;
        weights[round][account] += ticketsNum;

        token.safeTransferFrom(account, address(this), ticketsNum * ticketPrice);

        emit Bid(account, ticketsNum, round);
    }

    function withdraw(uint256 amount) public override {
        address account = msg.sender;
        require(amount > 0 && balances[account] >= amount, NothingToWithdraw());

        unchecked {
            balances[account] -= amount;
        }

        token.safeTransfer(account, amount);

        emit Withdraw(account, amount);
    }
}