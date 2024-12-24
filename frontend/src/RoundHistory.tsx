import React, { useEffect, useState } from 'react';
import { ethers } from 'ethers';
import { ALCHEMY_RPC_URL, CONTRACT_ADDRESS, CONTRACT_ABI } from './constants.tsx';
import copyLogo from './assets/copy.svg';
import etherscanLogo from './assets/etherscan.svg';
import './RoundHistory.css';

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
    start: string;
    end: string;
}

const RoundHistory: React.FC = () => {
    const [rounds, setRounds] = useState<RoundDetail[]>([]);
    const [expandedRound, setExpandedRound] = useState<number | null>(null); // To manage expanded dropdown
    const [itemsPerPage, setItemsPerPage] = useState<number>(10); // Items per page
    const [currentPage, setCurrentPage] = useState<number>(1);

    // Fetch round data
    const fetchRounds = async () => {
        try {
            const roundCount = await contract.round();
            const fetchedRounds: RoundDetail[] = [];

            for (let i = 0; i < roundCount; i++) {
                const response = await fetch(`http://localhost:8000/round/${i}`);
                const roundData = await response.json();

                const participantMap: { [address: string]: number } = {};
                roundData.events.forEach((e: any) => {
                    participantMap[e.account] = (participantMap[e.account] || 0) + e.amount;
                });

                const participants = Object.entries(participantMap).map(([address, tickets]) => ({
                    address,
                    tickets,
                }));

                const winnerResponse = await fetch(`http://localhost:8000/winner/${i}`);
                const winnerData = await winnerResponse.json();



                const totalPool = roundData.events.reduce((sum: bigint, e: any) => {
                    return sum + BigInt(e.amount || 0);
                }, BigInt(0));

                console.log("Raw totalPool (in wei):", totalPool.toString()); // Debugging

// Adjust by dividing by 10^4 instead of 10^14
                const correctedPool = (Number(totalPool) / 10 ** 4).toFixed(4); // Scale to ETH
                console.log("Corrected Pool (ETH):", correctedPool); // Debugging/ Debugging


                fetchedRounds.push({
                    round: i,
                    participants,
                    pool: `${correctedPool} ETH`,
                    winner: winnerData.winner.account,
                    winnings: ethers.formatEther(BigInt(winnerData.winner.amount)) + ' ETH',
                    start: roundData.start,
                    end: roundData.end,
                });
            }

            setRounds(fetchedRounds);
        } catch (error) {
            console.error('Error fetching rounds:', error);
        }
    };

    useEffect(() => {
        fetchRounds();
    }, []);

    const toggleRound = (round: number) => {
        setExpandedRound(expandedRound === round ? null : round);
        setCurrentPage(1); // Reset pagination when toggling
    };

    const paginate = (items: Participant[]) => {
        const start = (currentPage - 1) * itemsPerPage;
        return items.slice(start, start + itemsPerPage);
    };

    const handleItemsPerPageChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setItemsPerPage(Number(event.target.value));
        setCurrentPage(1); // Reset to the first page
    };

    const copyToClipboard = (address: string) => {
        navigator.clipboard.writeText(address);
        alert('Address copied to clipboard!');
    };

    return (
        <div className="round-history-container">
            <h1 className="lottery-title">Round History</h1>
            <div className="round-list">
                {rounds.map((round) => (
                    <div key={round.round} className="round-item">
                        <div className="round-header" onClick={() => toggleRound(round.round)}>
                            <p>
                                <strong>Round {round.round}:</strong> Pool: {round.pool}, Winner: {round.winner}
                            </p>
                            <p>Start: {round.start}, End: {round.end}</p>
                        </div>
                        {expandedRound === round.round && (
                            <div className="round-details">
                                <p><strong>Pool:</strong> {round.pool}</p>
                                <p>
                                    <strong>Winner:</strong> {round.winner}{' '}
                                    <img
                                        src={copyLogo}
                                        alt="Copy"
                                        className="icon"
                                        onClick={() => copyToClipboard(round.winner)}
                                    />
                                    <a
                                        href={`https://etherscan.io/address/${round.winner}`}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                    >
                                        <img src={etherscanLogo} alt="Etherscan" className="icon" />
                                    </a>
                                </p>
                                <p><strong>Winnings:</strong> {round.winnings}</p>
                                <p><strong>Participants:</strong></p>

                                {/* Pagination Controls */}
                                <div className="pagination-controls">
                                    <label>
                                        Show:
                                        <select value={itemsPerPage} onChange={handleItemsPerPageChange}>
                                            <option value={10}>10</option>
                                            <option value={50}>50</option>
                                            <option value={100}>100</option>
                                        </select>
                                    </label>
                                    <p>
                                        Page {currentPage} of {Math.ceil(round.participants.length / itemsPerPage)}
                                    </p>
                                    <button
                                        disabled={currentPage === 1}
                                        onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
                                    >
                                        Previous
                                    </button>
                                    <button
                                        disabled={currentPage === Math.ceil(round.participants.length / itemsPerPage)}
                                        onClick={() =>
                                            setCurrentPage((prev) => Math.min(prev + 1, Math.ceil(round.participants.length / itemsPerPage)))
                                        }
                                    >
                                        Next
                                    </button>
                                </div>

                                {/* Participants List */}
                                <ul className="participant-list">
                                    {paginate(round.participants).map((participant, index) => (
                                        <li key={index}>
                                            {participant.address} - {participant.tickets} ticket(s)
                                            <img
                                                src={copyLogo}
                                                alt="Copy"
                                                className="icon"
                                                onClick={() => copyToClipboard(participant.address)}
                                            />
                                            <a
                                                href={`https://etherscan.io/address/${participant.address}`}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                            >
                                                <img src={etherscanLogo} alt="Etherscan" className="icon" />
                                            </a>
                                        </li>
                                    ))}
                                </ul>
                            </div>
                        )}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default RoundHistory;
