export const ALCHEMY_RPC_URL = 'https://eth-holesky.g.alchemy.com/v2/UtjFRzFoEQd533NSskUCCCKEpW7z93t2';
export const CONTRACT_ADDRESS = '0x5B4Ac9c70E46e787df176c31d152c582Af27be76';
export const CONTRACT_ABI = [
    "event Bid(address indexed account, uint amount, uint timestamp, uint indexed round)",
    "event WinnerSelected(address indexed account, uint amount, uint indexed round)",
    "event Withdraw(address indexed account, uint amount)",
    "function balances(address) public view returns (uint256)",
    "function weights(uint256, address) public view returns (uint256)",
    "function totalWeight(uint256) public view returns (uint256)",
    "function participants(uint256) public view returns (address)",
    "function bid(uint amount) external payable",
    "function withdraw() external",
    "function getParticipantsNumber() external view returns (uint)",
    "function getTimeLeft() external view returns (uint)",
    "function bidFromBalance(uint amount) external",
    "function ticketPrice() public view returns (uint256)",
    "function round() public view returns (uint256)"
];