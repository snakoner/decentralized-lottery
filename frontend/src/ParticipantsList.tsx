import React, { useEffect, useState, useMemo } from 'react';
import { ethers } from 'ethers';

const contractABI = [
    "function participants(uint256) external view returns (address)",
    "function getParticipantsNum() external view returns (uint)",
    "event Bid(address indexed account, uint256 amount, uint256 timestamp, uint256 round)"
];
const provider = new ethers.JsonRpcProvider(
    `https://eth-sepolia.g.alchemy.com/v2/QtPw5bLONCtW00agVEhE66pb1Vsv9RnC`
);
const contractAddress = "0xE8f0b7144F2be28FE6Af69f15658d7b197Bf9f11";

const ParticipantsList: React.FC = () => {
    const [participants, setParticipants] = useState<{ address: string; bid: number }[]>([]);

    const contract = useMemo(() => new ethers.Contract(contractAddress, contractABI, provider), []);

    const fetchParticipants = async () => {
        try {
            console.log("Fetching participants...");
            const participantsCount: bigint = await contract.getParticipantsNum();
            const participantMap = new Map<string, number>();

            for (let i = 0; i < participantsCount; i++) {
                const participant = await contract.participants(i);
                participantMap.set(
                    participant,
                    (participantMap.get(participant) || 0) + 1
                );
            }

            const participantsArray = Array.from(participantMap.entries()).map(([address, bid]) => ({
                address,
                bid,
            }));

            setParticipants(participantsArray);
            console.log("Updated Participants List:", participantsArray);
        } catch (error) {
            console.error("Error fetching participants:", error);
        }
    };

    useEffect(() => {
        fetchParticipants();

        const handleBidEvent = (account: string, amount: ethers.BigNumber, timestamp: number, round: number) => {
            console.log("Bid event detected:", { account, amount: amount.toString(), timestamp, round });

            setParticipants((prevParticipants) => {
                const existingParticipant = prevParticipants.find(
                    (participant) => participant.address === account
                );

                if (existingParticipant) {
                    return prevParticipants.map((participant) =>
                        participant.address === account
                            ? { ...participant, bid: participant.bid + amount.toNumber() }
                            : participant
                    );
                } else {
                    return [...prevParticipants, { address: account, bid: amount.toNumber() }];
                }
            });
        };

        contract.on("Bid", handleBidEvent);

        return () => {
            contract.off("Bid", handleBidEvent);
        };
    }, [contract]);

    const truncateAddress = (address: string) => {
        return `${address.slice(0, 6)}...${address.slice(-4)}`;
    };

    const copyToClipboard = (address: string) => {
        navigator.clipboard.writeText(address).then(() => {
            alert(`Copied: ${address}`);
        });
    };

    return (
        <div className="lottery-container">
            <h1 className="lottery-title">Current Participants</h1>
            <ul className="list">
                {participants.map((participant, index) => (
                    <li key={index} className="list-item" onClick={() => copyToClipboard(participant.address)}>
                        {truncateAddress(participant.address)} - {participant.bid} ticket(s)
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ParticipantsList;
