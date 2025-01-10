// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IDecentralizedLottery {
    error NotEnoughtEther();
    
    error LotteryAlreadyFinished();

    error LotteryNotFinished();

    error InvalidRoundNumber();

    error HaveParticipants();

    error NotEnoughParticipants();

    error NothingToWithdraw();

    error WithdrawalFailed();

    /**
     * @dev Emitted when account buys one ticket
     */
    event Bid(address indexed account, uint256 amount, uint64 indexed round);

    /**
     * @dev Emitted when owner picked winner
     */
    event WinnerSelected(address indexed account, uint256 amount, uint64 indexed round);

    /**
     * @dev Emitted when owner picked winner
     */
    event Withdraw(address indexed account, uint256 amount);

    /**
     * @dev Emitted when owner restarts lottery
     */
    event RestartLottery(uint64 prevDuration, uint64 newDuration);

    /**
     * @dev Writes bid to mapping
     */
    function bid(uint256 amount) external payable;

    /**
     * @dev Function to choose winner
     */
    function start() external;

    /**
     * @dev Function to withdraw reward for winner and owner to withdraw commission
     */
    function withdraw(uint256 amount) external;

    /**
     * @dev Function to restart lottery if amount of participants is zero
     */
    function restartEmpty(uint64 newDuration) external;

}
