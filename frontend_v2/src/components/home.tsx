import React from "react";
import LotteryStatus from "./lottery/LotteryStatus";
import EntrySection from "./lottery/EntrySection";
import HistoryGrid from "./lottery/HistoryGrid";
import FeatureCards from "./lottery/FeatureCards";
import Header from "./layout/Header";
import Footer from "./layout/Footer";
import { useEffect, useState } from "react";
import { ethers } from "ethers"; 
import { ALCHEMY_RPC_URL, CONTRACT_ADDRESS, CONTRACT_ABI, localStorageWalletConnectHandler } from "./constants";

interface TimeRemaining {
    hours: number;
    minutes: number;
    seconds: number;
};

interface HomePageProps {
  poolAmount?: string;
  timeRemaining?: TimeRemaining;
  entryFee?: string;
  isProcessing?: boolean;
  error?: string;
  onEntrySubmit?: () => void;
}

interface Participant {
	id: string;
	address: string;
	amount: string;
	timestamp: string;  
};

function convertUnixTimestampToDate(timestamp: number): string {
	const date = new Date(timestamp * 1000);
  
	// Example format: 'YYYY-MM-DD HH:mm'
	const year = date.getFullYear();
	const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are zero-indexed
	const day = String(date.getDate()).padStart(2, '0');
	const hours = String(date.getHours()).padStart(2, '0');
	const minutes = String(date.getMinutes()).padStart(2, '0');
  
	return `${year}-${month}-${day} ${hours}:${minutes}`;
  }

const HomePage = ({
	entryFee = "0.1 ETH",
	isProcessing = false,
	error = "",
	onEntrySubmit = () => {},
}: HomePageProps) => {

	const [allTimeReward, setAllTimeReward] = useState<string|null>("0");
	const [currentRound, setCurrentRound] = useState<number>(0);
	const [timeRemaining, setTimeRemaining] = useState<TimeRemaining>();
	const [ticketPrice, setTicketPrice] = useState<string|null>("0");
	const [ticketsPerAccount, setTicketsPerAccount] = useState<number>(0);
	const [secondsTimeRemaining, setSecondsTimeRemaining] = useState<number>(0);
	const [poolAmount, setPoolAmount] = useState<string|null>("0.0");
	const [winners, setWinners] = useState<Participant[]|null>();
	const [participants, setParticipants] = useState<Participant[]|null>();
	const [account, setAccount] = useState<string|null>();

    // const browserProvider = new ethers.BrowserProvider(window.ethereum);
    const providerRpc = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
    const contractRpc = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, providerRpc);

	const getTimeRemaining = async() => {
		try {
            const num = Number(await contractRpc.getTimeLeft());
			setSecondsTimeRemaining(num);
            return num;    
        } catch (error) {
			throw error;
        }
	}

	const getAllTimeReward = async() => {
		try {
			const response = await fetch("http://localhost:8000/all-time-reward", {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
				},     
			});
			
			const data = await response.json();
			setAllTimeReward(data['reward'].slice(0, 7));
		} catch(error) {
			console.log(error);
		}
	};

	const getCurrentRound = async() => {
		try {
            const num = Number(await contractRpc.round());
			setCurrentRound(num);
			return num;    
        } catch (error) {
			throw error;
        }
	}

	const getTicketPrice = async() => {
		try {
			const num: bigint = await contractRpc.ticketPrice();
			setTicketPrice(ethers.formatUnits(num));

			return num;
		} catch (error) {
			throw error;
		}
	}

	const getHistoricalWinners = async() => {
		const pastWinners: Participant[] = [];
		let poolSize = ethers.toBigInt(0);

		try {
			const num: bigint = await contractRpc.round();

			for (let i = 0; i < Number(num); i++) {
				const response = await fetch(`http://localhost:8000/winner/${i}`, {
					method: 'GET',
					headers: {
						'Content-Type': 'application/json',
					},     
				});

				const data = await response.json();
				const winner = data['winner'];
				if (winner !== undefined) {
					pastWinners.push({
						id: `w${winner?.round}`,
						address: winner?.account,
						amount: ethers.formatUnits(winner?.amount),
						timestamp: convertUnixTimestampToDate(winner?.timestamp),
					});
				}
			}

			setWinners(pastWinners);
		} catch(error) {
			console.log(error);
		}
	}


	const getCurrentParticipants = async() => {
		const participants: Participant[] = [];
		try {
			let round = await getCurrentRound();
			let poolSize = ethers.toBigInt(0);
			let ticketsPerAcc: number = 0;

			const ticketPrice: bigint = await contractRpc.ticketPrice();
			const response = await fetch(`http://localhost:8000/round/${round}`, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
				},     
			});

			const data = await response.json();
			const events = data['events'];
			let _account: string;
			if (!account) {
				if (localStorageWalletConnectHandler()) {
					const accounts = await window.ethereum.request({
						method: "eth_requestAccounts",
					});

					if (accounts.length === 0) {
						localStorage.setItem('walletConnected', 'false');
					}
					
					_account = accounts[0];
				}
			}

			if (events) {
				for (let i = 0; i < events.length; i++) {
					if (_account && events[i].account.toUpperCase() === _account.toUpperCase()) {
						ticketsPerAcc += events[i].amount;
					}
					poolSize += ethers.toBigInt(events[i]['amount']) * ticketPrice;				
					participants.push({
						id: `w${i}`,
						address: events[i]?.account,
						amount: ethers.formatUnits(ethers.toBigInt(events[i].amount) * ticketPrice),
						timestamp: convertUnixTimestampToDate(events[i]?.timestamp),
					});				
				}
			}
			console.log('poolSize:', ethers.formatUnits(poolSize));
			console.log('tickets: ', ticketsPerAcc);
			setTicketsPerAccount(ticketsPerAcc);
			setParticipants(participants);
			setPoolAmount(ethers.formatUnits(poolSize));
		} catch(error) {
			console.log(error);
		}
	}

	useEffect(() => {
		getAllTimeReward();
		getTimeRemaining();
		getTicketPrice();
		getHistoricalWinners();
		getCurrentParticipants();

		const timer = setInterval(() => {
            setSecondsTimeRemaining((prevTime) => (prevTime > 0 ? prevTime - 1 : 0));
        }, 1000);
        return () => clearInterval(timer);
    }, []);
	
	return (
		<div className="min-h-screen flex flex-col bg-gradient-to-b from-gray-50 to-gray-100 dark:from-gray-950 dark:to-gray-900 transition-colors duration-300">
		<Header account={account} setAccount={setAccount}/>
		<main className="flex-1 py-8 px-4 animate-fade-in">
			<div className="max-w-[1200px] mx-auto space-y-12">
			<div className="text-center space-y-4 mb-8">
				<h1 className="text-4xl font-bold bg-gradient-to-r from-purple-500 via-blue-500 to-purple-500 dark:from-purple-400 dark:via-blue-400 dark:to-purple-400 bg-clip-text text-transparent">
				Welcome to ETH Lottery
				</h1>
				<p className="text-xl text-muted-foreground">
				Your chance to win big with Ethereum
				</p>
			</div>

			<div className="transform hover:scale-[1.01] transition-transform">
				<FeatureCards />
			</div>

			<div className="transform hover:scale-[1.01] transition-transform">
				<LotteryStatus
				currentRound={currentRound}
				poolAmount={allTimeReward}
				secondsTimeRemaining={secondsTimeRemaining}
				/>
			</div>

			<div className="transform hover:scale-[1.01] transition-transform">
				<EntrySection
				poolAmount={poolAmount}
				ticketPrice={ticketPrice}
				ticketsPerAccount={ticketsPerAccount}
				isProcessing={isProcessing}
				error={error}
				onEntrySubmit={onEntrySubmit}
				/>
			</div>

			<div className="transform hover:scale-[1.01] transition-transform">
				<HistoryGrid winners={winners} participants={participants}/>
			</div>
			</div>
		</main>
		<Footer />
		</div>
	);
};

export default HomePage;
