import React from 'react';
import {useState, useEffect} from 'react';
import { ethers } from 'ethers';
import Header from './Header.tsx';
import LotteryStatus from './Lottery.tsx';
import LotteryInfo from './LotteryInfo.tsx';
import WinnersList from './WinnersList.tsx';
import RoundHistory from './RoundHistory.tsx';
import Footer from './Footer.tsx';
import { ALCHEMY_RPC_URL, localStorageWalletConnectHandler, supportedChains } from './constants.tsx';
import './App.css';
import './index.css';
import './Background.css';

const convertSupportedChains = (): Map<bigint, ethers.Network> => {
    let map = new Map<bigint, ethers.Network>();

    for (const suppChains of supportedChains) {
        map.set(
            ethers.toBigInt(suppChains.chainId), 
            new ethers.Network(suppChains.network, suppChains.chainId)
        );
    }

    return map;
}

function App() {
    const [account, setAccount] = useState<string|null>(null);
    const [network, setNetwork] = useState<ethers.Network|null>(null);
    const [walletError, setWalletError] = useState<string|null>(null);
    const [walletBalance, setWalletBalance] = useState<string|null>(null);
    const [connected, setConnected] = useState<boolean>(false);
    const fallingStars: JSX.Element[] = [];
    const minTimeFall = 2.0;

    const generateFallingStars = () => {
        fallingStars.length = 0;
        for (let i = 0; i < 20; i++) {
            const randValue = Math.random() * 6;;
            const timeFall = randValue < minTimeFall ? minTimeFall : randValue;
            const marginLeft = Math.random() * 1400;
            fallingStars.push(
                <div className='falling-star' style={{marginLeft: `${marginLeft}px`, animation: `fall ${timeFall}s linear infinite`}}></div>
            );
        }    
    }

    generateFallingStars();

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

            const chainID =  (await provider.getNetwork()).chainId;
            if (convertSupportedChains().get(chainID) === undefined) {
                await window.ethereum.request({
                    method: 'wallet_switchEthereumChain',
                    params: [{ chainId: '0x' + supportedChains[0].chainId.toString(16) }], // chainId в формате hex (например, '0x1' для Ethereum Mainnet)
                });
            }

            setNetwork((await provider.getNetwork()).name);
            setWalletBalance(walletBalanceFormat(await remoteProvider.getBalance(accounts[0])));
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
            {view === 'history' && (
                <>
                        <WinnersList />
                        <RoundHistory />
                </>
                
            )}
        <Footer></Footer>
        {fallingStars}
        </div>
    );
}

export default App;
