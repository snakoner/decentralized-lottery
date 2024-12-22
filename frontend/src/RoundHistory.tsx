import React, { useEffect, useState } from 'react';
import { ethers } from 'ethers';
import Modal from './Modal.tsx';
import { ALCHEMY_RPC_URL, CONTRACT_ADDRESS, CONTRACT_ABI } from './constants.tsx';

const provider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
const contract = new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, provider);

interface Participant {
    address: string;
    tickets: number;
}

interface RoundDetail {
    round: number;
    participants: Participant[];
    pool: string;
    winner: string;
    winnings: string;
}

const RoundHistory: React.FC = () => {
    const [rounds, setRounds] = useState<{ round: number; participants: number; pool: string }[]>([]);
    const [currentRound, setCurrentRound] = useState<number | null>(null);
    const [roundDetailPopup, setRoundDetailPopup] = useState<RoundDetail | null>(null);
    const [modalIsOpen, setModalIsOpen] = useState<boolean>(false);
    const [participantModalContent, setParticipantModalContent] = useState<string | null>(null);

    // Fetch current round number from the contract
    const fetchCurrentRound = async () => {
        try {
            const round = await contract.round();
            setCurrentRound(Number(round));
        } catch (error) {
            console.error("Error fetching current round:", error);
        }
    };

    // Fetch round summaries
    const fetchRounds = async () => {
        if (currentRound === null) return;

        const fetchedRounds = [];
        for (let i = 0; i < currentRound; i++) {
            try {
                const response = await fetch(`http://localhost:8000/round/${i}`);
                const data = await response.json();
                const uniqueParticipants = new Set(data.events.map((e: any) => e.account));
                const totalPool = ethers.formatEther(
                    ethers.toBigInt(
                        data.events.reduce((sum: bigint, e: any) => sum + BigInt(e.amount), BigInt(0))
                    )
                );

                fetchedRounds.push({
                    round: i,
                    participants: uniqueParticipants.size,
                    pool: `${totalPool} ETH`,
                });
            } catch (error) {
                console.error(`Error fetching round ${i}:`, error);
            }
        }

        setRounds(fetchedRounds);
    };

    // Fetch detailed information for a specific round
    const fetchRoundDetails = async (round: number) => {
        try {
            // Fetch participants
            const roundResponse = await fetch(`http://localhost:8000/round/${round}`);
            const roundData = await roundResponse.json();
            const participantMap: { [address: string]: number } = {};

            roundData.events.forEach((e: any) => {
                participantMap[e.account] = (participantMap[e.account] || 0) + e.amount;
            });

            const participants = Object.entries(participantMap).map(([address, tickets]) => ({
                address,
                tickets,
            }));

            // Fetch winner
            const winnerResponse = await fetch(`http://localhost:8000/winner/${round}`);
            const winnerData = await winnerResponse.json();
            const winner = winnerData.winner.account;
            const winnings = ethers.formatEther(ethers.toBigInt(winnerData.winner.amount)) + " ETH";

            setRoundDetailPopup({
                round,
                participants,
                pool: rounds.find((r) => r.round === round)?.pool || "0 ETH",
                winner,
                winnings,
            });
        } catch (error) {
            console.error(`Error fetching details for round ${round}:`, error);
        }
    };

    useEffect(() => {
        fetchCurrentRound();
    }, []);

    useEffect(() => {
        if (currentRound !== null) {
            fetchRounds();
        }
    }, [currentRound]);

    const truncateAddress = (address: string) => `${address.slice(0, 6)}...${address.slice(-4)}`.toLowerCase();

    return (
        <div className="main-container">
            <div className="lottery-container">
                <h1 className="lottery-title">Round History</h1>
                <ul className="list">
                    {rounds.map((round, index) => (
                        <li
                            key={index}
                            className="list-item"
                            onClick={() => fetchRoundDetails(round.round)}
                            style={{ cursor: "pointer", marginBottom: "10px" }}
                        >
                            Round {round.round}: {round.participants} participants, pool: {round.pool}
                        </li>
                    ))}
                </ul>
            </div>

            {/* Round Details Popup */}
            {roundDetailPopup && (
                <div className="modal-overlay">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h2>Round {roundDetailPopup.round} Details</h2>
                            <button onClick={() => setRoundDetailPopup(null)}>✖</button>
                        </div>
                        <div className="modal-body">
                            <p><strong>Total Pool:</strong> {roundDetailPopup.pool}</p>
                            <p><strong>Winner:</strong> {truncateAddress(roundDetailPopup.winner)}</p>
                            <p><strong>Winnings:</strong> {roundDetailPopup.winnings}</p>
                            <p><strong>Participants:</strong></p>
                            <ul>
                                {roundDetailPopup.participants.map((p, idx) => (
                                    <li
                                        key={idx}
                                        style={{ cursor: "pointer", marginBottom: "8px" }}
                                        onClick={() => {
                                            setParticipantModalContent(`Address: ${p.address}, Tickets: ${p.tickets}`);
                                            setModalIsOpen(true);
                                        }}
                                    >
                                        {truncateAddress(p.address)} - {p.tickets} ticket(s)
                                    </li>
                                ))}
                            </ul>
                        </div>
                    </div>
                </div>
            )}

            {/* Participant Modal */}
            {modalIsOpen && (
                <Modal
                    isOpen={modalIsOpen}
                    onClose={() => setModalIsOpen(false)}
                    title="Participant Details"
                    content={<div>{participantModalContent}</div>}
                />
            )}
        </div>
    );
};

export default RoundHistory;
