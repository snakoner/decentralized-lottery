export const ALCHEMY_RPC_URL = 'https://eth-sepolia.g.alchemy.com/v2/UtjFRzFoEQd533NSskUCCCKEpW7z93t2';
export const CONTRACT_ADDRESS = '0xFB0E88b0c92e3583E1d4028221608c1Cf35637D5';
export const CONTRACT_ABI = [
    "event Bid(address indexed account, uint amount, uint timestamp, uint indexed round)",
    "event WinnerSelected(address indexed account, uint amount, uint indexed round)",
    "event Withdraw(address indexed account, uint amount)",
    "event Deposit(address indexed user, uint amount, uint timestamp)",
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
    "function round() public view returns (uint256)",
    "function deposit() external payable"
];