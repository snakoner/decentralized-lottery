import React, { useEffect, useState } from 'react';
import { ethers } from 'ethers';
import Modal from './Modal.tsx';
import { ALCHEMY_RPC_URL, CONTRACT_ADDRESS } from './constants.tsx';
const contractABI = [
    "function round() external view returns (uint)",
];
const provider = new ethers.JsonRpcProvider(
    ALCHEMY_RPC_URL
);

const WinnersList: React.FC = () => {
    const [winners, setWinners] = useState<{ round: number; address: string; prize: string }[]>([]);
    const [currentRound, setCurrentRound] = useState<number | null>(null);
    const [modalIsOpen, setModalIsOpen] = useState<boolean>(false);
    const [selectedWinner, setSelectedWinner] = useState<{ round: number; address: string; prize: string } | null>(null);

    const contract = new ethers.Contract(CONTRACT_ADDRESS, contractABI, provider);

    // Fetch the current round number from the contract
    const fetchCurrentRound = async () => {
        try {
            const round = await contract.round();
            setCurrentRound(Number(round));
        } catch (error) {
            console.error("Error fetching current round:", error);
        }
    };

    // Fetch winner details for all rounds
    const fetchWinners = async () => {
        if (currentRound === null) return;

        const fetchedWinners = [];
        for (let i = 0; i < currentRound; i++) {
            try {
                console.log('try fetching round: ', i);
                const response = await fetch(`http://localhost:8000/winner/${i}`);
                if (!response.ok) {
                    console.error(`Error fetching winner for round ${i}:`, response.statusText);
                    continue;
                }
                const data = await response.json();
                if (data['winner']['round-finished']) {
                    fetchedWinners.push({
                        round: data['winner']['round'],
                        address: data['winner']['account'],
                        prize: ethers.formatUnits(data['winner']['amount'].toString()) + " ETH", // Format the prize amount
                    });
                }
            } catch (error) {
                console.error(`Error fetching winner for round ${i}:`, error);
            }
        }

        console.log('setWinners:', fetchedWinners);
        setWinners(fetchedWinners);
    };

    useEffect(() => {
        fetchCurrentRound();
    }, []);

    useEffect(() => {
        if (currentRound !== null) {
            fetchWinners();
        }
    }, [currentRound]);

    const truncateAddress = (address: string) => `${address.slice(0, 6)}...${address.slice(-4)}`.toLowerCase();

    const showWinnerInfo = (winner: { round: number; address: string; prize: string }) => {
        setSelectedWinner(winner);
        setModalIsOpen(true);
    };

    const hideWinnerInfo = () => {
        setSelectedWinner(null);
        setModalIsOpen(false);
    };

    return (
        <div className='main-container'>
        <div className="lottery-container" style={{margin:"0 auto"}}>
            <h1 className="lottery-title">Past Winners</h1>
            <ul className="list">
                {winners.map((winner, index) => (
                    <li
                        key={index}
                        className="list-item"
                        onClick={() => showWinnerInfo(winner)}
                        style={{ cursor: "pointer", marginBottom: "10px" }}
                    >
                        Round {winner.round}: {truncateAddress(winner.address)} won {winner.prize}
                    </li>
                ))}
            </ul>

            {selectedWinner && (
                <Modal
                    isOpen={modalIsOpen}
                    onClose={hideWinnerInfo}
                    title="Winner Details"
                    content={
                        <div>
                            <p><strong>Round:</strong> {selectedWinner.round}</p>
                            <p><strong>Address:</strong> {selectedWinner.address}</p>
                            <p><strong>Prize:</strong> {selectedWinner.prize}</p>
                        </div>
                    }
                />
            )}
        </div>
        </div>
    );
};

export default WinnersList;
