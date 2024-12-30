import { Button } from "@/components/ui/button";
import { Wallet } from "lucide-react";
import { ThemeToggle } from "@/components/theme/theme-toggle";
import { Logo } from "@/components/ui/logo";
import { useState, useEffect } from "react";
import { ethers } from "ethers";
import { ALCHEMY_RPC_URL, supportedChains, localStorageWalletConnectHandler } from "../constants.tsx";
import { add } from "date-fns";

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

export default function Header({account: string, setAccount}) {
	const [walletError, setWalletError] = useState<string|null>(null);
	// const [account, setAccount] = useState<string|null>(null);
	const [network, setNetwork] = useState<ethers.Network|null>(null);
    const [walletBalance, setWalletBalance] = useState<string|null>(null);
    const [connected, setConnected] = useState<boolean>(false);


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
				setConnected(false);
				localStorage.setItem('walletConnected', 'false');
				return;
			}

			setAccount(accounts[0]);
			setConnected(true);
			localStorage.setItem('walletConnected', 'true');
			connectButtonSet(slicedWalletAddress(accounts[0]));

			const remoteProvider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
			const provider = new ethers.BrowserProvider(window.ethereum);

			const chainID =  (await provider.getNetwork()).chainId;
			if (convertSupportedChains().get(chainID) === undefined) {
				await window.ethereum.request({
					method: 'wallet_switchEthereumChain',
					params: [{ chainId: '0x' + supportedChains[0].chainId.toString(16) }], // chainId в формате hex (например, '0x1' для Ethereum Mainnet)
				});
			}

			setNetwork((await provider.getNetwork()));
			setWalletBalance(walletBalanceFormat(await remoteProvider.getBalance(accounts[0])));
		} catch (err) {
			setWalletError("Failed to connect wallet. Please try again.");
			console.log(walletError);
		}
	}

	const connectButtonSet = (data: string) => {
		const element = document.getElementById('wallet-connect-button-text');
		if (element !== undefined) {
			element.innerText = data;
		}
	}

	const slicedWalletAddress = (acc: string) => {
		return acc.slice(0, 6) + "..." + acc.slice(-4);
	};

	const disconnectWallet = () => {
        setConnected(false);
        localStorage.setItem('walletConnected', 'false');
        setAccount(null);
        setWalletBalance(null);
        setNetwork(null);
		connectButtonSet("Connect wallet");
    }

	useEffect(() => {
        const walletConnected = localStorageWalletConnectHandler();
		setConnected(walletConnected);
		if (walletConnected) {
            connectWallet();
        }
    }, []);

	return (
		<header className="sticky top-0 z-50 border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
		<div className="container flex h-16 items-center justify-between">
			<Logo />
			<nav className="flex items-center space-x-4">
			<ThemeToggle />
			<Button
				onClick={connected ? disconnectWallet: connectWallet}
				variant="outline"
				className="gap-2 hover:scale-105 transition-all hover:bg-purple-50 hover:border-purple-200 hover:text-purple-600 dark:hover:bg-purple-900 dark:hover:border-purple-700 dark:hover:text-purple-400"
			>
				<Wallet className="h-4 w-4" />
				<div id='wallet-connect-button-text'>Connect wallet</div>
			</Button>
			</nav>
		</div>
		</header>
	);
}
