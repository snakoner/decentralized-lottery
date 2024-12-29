import React from "react";
import LotteryStatus from "./lottery/LotteryStatus";
import EntrySection from "./lottery/EntrySection";
import HistoryGrid from "./lottery/HistoryGrid";
import FeatureCards from "./lottery/FeatureCards";
import Header from "./layout/Header";
import Footer from "./layout/Footer";
import { useEffect, useState } from "react";
import { ALCHEMY_RPC_URL, CONTRACT_ADDRESS, CONTRACT_ABI } from "./constants";
import { ethers } from "ethers"; 

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

interface Winner {
	id: string;
	address: string;
	amount: string;
	timestamp: string;  
};

const HomePage = ({
	// poolAmount = "100 ETH",
	entryFee = "0.1 ETH",
	isProcessing = false,
	error = "",
	onEntrySubmit = () => {},
}: HomePageProps) => {

	const [allTimeReward, setAllTimeReward] = useState<string|null>("0");
	const [timeRemaining, setTimeRemaining] = useState<TimeRemaining>();
	const [ticketPrice, setTicketPrice] = useState<string|null>("0");
	const [secondsTimeRemaining, setSecondsTimeRemaining] = useState<number>(0);
	const [poolAmount, setPoolAmount] = useState<string|null>("0.0");
	const [winners, setWinners] = useState<Winner[]|null>();

    const browserProvider = new ethers.BrowserProvider(window.ethereum);
    const providerRpc = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
    const contractRpc = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, providerRpc);

	const getAllTimeReward = async() => {
		try {
			const response = await fetch("http://localhost:8000/all-time-reward", {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
				},     
			});
			
			const data = await response.json();
			console.log("all time reward: ", data['reward']);
			setAllTimeReward(data['reward'].slice(0, 7));
		} catch(error) {
			console.log(error);
		}
	};

	const getCurrentRound = async() => {
		try {
            const num: bigint = await contractRpc.round();
            return Number(num);    
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
		const pastWinners: Winner[] = [];
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
				if (data['winner'] !== undefined) {
					console.log('winner: ', data['winner']);	
					pastWinners.push({
						id: `w${data['winner']['round']}`,
						address: data['winner']['account'],
						amount: ethers.formatUnits(data['winner']['amount']),
						timestamp: "17657374543",
					});				
				}
			}

			setWinners(pastWinners);
		} catch(error) {
			console.log(error);
		}

		console.log(pastWinners);
	}

	const getCurrentPool = async() => {
		try {
			const currentRound = await getCurrentRound();
			const ticketPrice = await getTicketPrice();
		
			const response = await fetch(`http://localhost:8000/round/${currentRound}`, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
				},     
			});
			
			const data = await response.json();
			if (data['events']) {
				let poolSize = ethers.toBigInt(0);
				for (const event of data['events'])
					poolSize += ethers.toBigInt(event['amount']) * ticketPrice;				

				setPoolAmount(ethers.formatUnits(poolSize));
			}
		} catch(error) {
			console.log(error);
		}
	}

	const formatTime = (seconds: number): TimeRemaining => {
		const hours = Math.floor(seconds / 3600);
		const minutes = Math.floor((seconds % 3600) / 60);
		const secs = seconds % 60;
		
		return {
			hours: hours,
			minutes: minutes,
			seconds: secs,
		};
	};

	useEffect(() => {
		getAllTimeReward();
		getCurrentPool();
		getHistoricalWinners();
    }, []);
	
	return (
		<div className="min-h-screen flex flex-col bg-gradient-to-b from-gray-50 to-gray-100 dark:from-gray-950 dark:to-gray-900 transition-colors duration-300">
		<Header />
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
				poolAmount={allTimeReward}
				timeRemaining={timeRemaining}
				/>
			</div>

			<div className="transform hover:scale-[1.01] transition-transform">
				<EntrySection
				poolAmount={poolAmount}
				ticketPrice={ticketPrice}
				isProcessing={isProcessing}
				error={error}
				onEntrySubmit={onEntrySubmit}
				/>
			</div>

			<div className="transform hover:scale-[1.01] transition-transform">
				<HistoryGrid winners={winners}/>
			</div>
			</div>
		</main>
		<Footer />
		</div>
	);
};

export default HomePage;
