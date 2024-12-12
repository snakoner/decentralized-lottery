import React from 'react';
import Header from './Header.tsx';
import './App.css';
import { ethers } from 'ethers';
import {useState, useEffect} from 'react';
import './index.css';
import LotteryStatus from './Lottery.tsx';
import WalletConnect from './Header.tsx';

const supportedChains: ethers.Network[] = [
    new ethers.Network('sepolia', 11155111),
];

const CONTRACT_ADDRESS = "0x5908C0CD77e0FA105565a2BAF5c0F0C4ba60e978";
const ALCHEMY_RPC_URL = process.env.REACT_APP_ALCHEMY_PRC_URL;

function App() {
    const [account, setAccount] = useState<string|null>(null);
    const [network, setNetwork] = useState<ethers.Network|null>(null);
    const [walletError, setWalletError] = useState<string|null>(null);
    const [walletBalance, setWalletBalance] = useState<string|null>(null);
    const [connected, setConnected] = useState<boolean>(false);

    const walletBalanceFormat = (balance: bigint) => {
        return ethers.formatUnits(balance).slice(0, 6) + " ETH";        
    }

    const connectWallet = async () => {
        if (typeof window.ethereum == undefined) {
            setWalletError("MetaMask is not installed. Please install it to use this feature.");
            console.log(walletError);
            return;
        }

        try {
            const accounts = await window.ethereum.request({
                method: "eth_requestAccounts",
            });

            if (accounts.length == 0) {
                localStorage.setItem('walletConnected', 'false');
                return;
            }

            setAccount(accounts[0]);
            setConnected(true);
            localStorage.setItem('walletConnected', 'true');

            const remoteProvider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
            const provider = new ethers.BrowserProvider(window.ethereum);

            setWalletBalance(walletBalanceFormat(await remoteProvider.getBalance(accounts[0])));
            setNetwork((await provider.getNetwork()).name);
        } catch (err) {
            setWalletError("Failed to connect wallet. Please try again.");
            console.log(walletError);
        }
    }

    const disconnectWallet = () => {
        localStorage.setItem('walletConnected', 'false');
        setConnected(false);
        setAccount(null);
        setWalletBalance(null);
        setNetwork(null);
    }

    const localStorageWalletConnectHandler = () => {
        if (localStorage.getItem('walletConnected') === undefined) {
            localStorage.setItem('walletConnected', 'false');
        }

        return localStorage.getItem('walletConnected') == 'true' ? true : false;
    }

    useEffect(() => {
        const walletConnected = localStorageWalletConnectHandler();
        if (walletConnected) {
            connectWallet();
        }
    }, []);

    return (
        <div className="App">
            <Header
                connected={connected}
                account={account} 
                network={network} 
                error={walletError} 
                connectWallet={connectWallet}
                disconnectWallet={disconnectWallet}
                walletBalance={walletBalance}
            ></Header>
            <LotteryStatus connected={connected} account={account}></LotteryStatus>
        </div>
    );
}

export default App;
