export const ALCHEMY_RPC_URL = 'https://eth-sepolia.g.alchemy.com/v2/UtjFRzFoEQd533NSskUCCCKEpW7z93t2';
export const CONTRACT_ADDRESS = '0xfC8118820706C3EF0C47Ff0d9d7C2b67fDAf23b9';
export const CONTRACT_ABI = [
    "event Bid(address indexed account, uint256 amount, uint64 indexed round)",
    "event WinnerSelected(address indexed account, uint256 amount, uint64 indexed round)",
    "event Withdraw(address indexed account, uint256 amount)",
    "function balances(address) public view returns (uint256)",
    "function weights(uint256, address) public view returns (uint256)",
    "function totalWeight(uint256) public view returns (uint256)",
    "function participants(uint256) public view returns (address)",
    "function bid(uint256 amount) external payable",
    "function withdraw(uint256 amount) external",
    "function getParticipantsNumber() external view returns (uint256)",
    "function getTimeLeft() external view returns (uint256)",
    "function bidFromBalance(uint256 amount) external",
    "function ticketPrice() public view returns (uint256)",
    "function round() public view returns (uint256)",
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