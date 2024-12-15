import React from 'react';
import Header from './Header.tsx';
import './App.css';
import { ethers } from 'ethers';
import {useState, useEffect} from 'react';
import './index.css';
import LotteryStatus from './Lottery.tsx';
import LotteryInfo from './LotteryInfo.tsx';
import WinnersList from './WinnersList.tsx';

const supportedChains: ethers.Network[] = [
    new ethers.Network('sepolia', 11155111),
];

const ALCHEMY_RPC_URL = `https://eth-sepolia.g.alchemy.com/v2/QtPw5bLONCtW00agVEhE66pb1Vsv9RnC`;

function App() {
    const [account, setAccount] = useState<string|null>(null);
    const [network, setNetwork] = useState<ethers.Network|null>(null);
    const [walletError, setWalletError] = useState<string|null>(null);
    const [walletBalance, setWalletBalance] = useState<string|null>(null);
    const [connected, setConnected] = useState<boolean>(false);

    // New state to manage views
    const [view, setView] = useState<'home' | 'history'>('home');

    const walletBalanceFormat = (balance: bigint) => {
        return ethers.formatUnits(balance).slice(0, 6);        
    }

    const connectWallet = async () => {
        if (typeof window.ethereum === undefined) {
            setWalletError("MetaMask is not installed. Please install it to use this feature.");
            console.log(walletError);
            return;
        }

        try {
            const accounts = await window.ethereum.request({
                method: "eth_requestAccounts",
            });

            if (accounts.length === 0) {
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

        return localStorage.getItem('walletConnected') === 'true' ? true : false;
    }

    useEffect(() => {
        const walletConnected = localStorageWalletConnectHandler();
        if (walletConnected) {
            connectWallet();
        }
    }, []);

    // Handlers for switching views
    const handleHomeClick = () => setView('home'); // Navigate to Home
    const handleHistoryClick = () => setView('history'); // Navigate to History


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
                onHomeClick={handleHomeClick} // Pass navigation handlers to Header
                onHistoryClick={handleHistoryClick}
            />
            {view === 'home' && (
                <>
                    <LotteryInfo />
                    <LotteryStatus connected={connected} account={account} />
                </>
            )}
            {view === 'history' && <WinnersList />}
        </div>
    );
}

export default App;
