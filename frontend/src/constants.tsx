import { network } from "hardhat";

export const ALCHEMY_RPC_URL = 'https://eth-sepolia.g.alchemy.com/v2/UtjFRzFoEQd533NSskUCCCKEpW7z93t2';
export const CONTRACT_ADDRESS = '0xfC8118820706C3EF0C47Ff0d9d7C2b67fDAf23b9';
export const CONTRACT_ABI = [
    "event Bid(address indexed account, uint amount, uint indexed round)",
    "event WinnerSelected(address indexed account, uint amount, uint indexed round)",
    "event Withdraw(address indexed account, uint amount)",
    "event Deposit(address indexed user, uint amount)",
    "function balances(address) public view returns (uint256)",
    "function weights(uint256, address) public view returns (uint256)",
    "function totalWeight(uint256) public view returns (uint256)",
    "function participants(uint256) public view returns (address)",
    "function bid(uint amount) external payable",
    "function withdraw(uint amount) external",
    "function getParticipantsNumber() external view returns (uint)",
    "function getTimeLeft() external view returns (uint)",
    "function bidFromBalance(uint amount) external",
    "function ticketPrice() public view returns (uint256)",
    "function round() public view returns (uint256)",
    "function deposit() external payable"
];

export const localStorageWalletConnectHandler = () => {
    if (localStorage.getItem('walletConnected') === undefined) {
        localStorage.setItem('walletConnected', 'false');
    }

    return localStorage.getItem('walletConnected') === 'true' ? true : false;
}

export const supportedChains = [
    {
        network: "sepolia",
        chainId: 11155111,
    },
];