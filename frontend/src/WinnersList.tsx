import React, { useEffect, useState } from 'react';
import { ethers } from 'ethers';
import Modal from './Modal.tsx';

const contractABI = [
    "function round() external view returns (uint)",
];
const provider = new ethers.JsonRpcProvider(
    `https://eth-sepolia.g.alchemy.com/v2/QtPw5bLONCtW00agVEhE66pb1Vsv9RnC`
);
const contractAddress = "0xE8f0b7144F2be28FE6Af69f15658d7b197Bf9f11";

const WinnersList: React.FC = () => {
    const [winners, setWinners] = useState<{ round: number; address: string; prize: string }[]>([]);
    const [currentRound, setCurrentRound] = useState<number | null>(null);
    const [modalIsOpen, setModalIsOpen] = useState<boolean>(false);
    const [selectedWinner, setSelectedWinner] = useState<{ round: number; address: string; prize: string } | null>(null);

    const contract = new ethers.Contract(contractAddress, contractABI, provider);

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
        for (let i = 1; i <= currentRound; i++) {
            try {
                const response = await fetch(`/winner/${i}`);
                if (!response.ok) {
                    console.error(`Error fetching winner for round ${i}:`, response.statusText);
                    continue;
                }
                const data = await response.json();
                if (data?.Winner?.Account) {
                    fetchedWinners.push({
                        round: data.Winner.Round,
                        address: data.Winner.Account,
                        prize: ethers.formatUnits(data.Winner.Amount, 18) + " ETH", // Format the prize amount
                    });
                }
            } catch (error) {
                console.error(`Error fetching winner for round ${i}:`, error);
            }
        }
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
        <div className="lottery-container">
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
    );
};

export default WinnersList;
