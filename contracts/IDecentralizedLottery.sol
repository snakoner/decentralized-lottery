// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

interface IDecentralizerLottery {
    /**
     * @dev Emitted when account buys one ticket
     */
    event Bid(address indexed account, uint amount, uint timestamp, uint indexed round);

    /**
     * @dev Emitted when owner picked winner
     */
    event WinnerSelected(address indexed account, uint amount, uint indexed round);

    /**
     * @dev Emitted when owner picked winner
     */
    event Withdraw(address indexed account, address indexed to, uint amount);

    /**
     * @dev Writes bid to mapping
     */
    function bid(uint amount) external payable;

    /**
     * @dev Function to choose winner
     */
    function start() external;

    /**
     * @dev Function to withdraw reward for winner and owner to withdraw commission
     */
    function withdraw(address payable to) external;

    /**
     * @dev Function to restart lottery if amount of participants is zero
     */
    function restartEmpty(uint newDuration) external;

}
