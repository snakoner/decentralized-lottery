
import React, { useEffect, useState, useMemo } from 'react';
import Modal from "./Modal.tsx"; // Added: Import the Modal component
import { ethers } from 'ethers';

const contractABI = [
    "function participants(uint256) external view returns (address)",
    "function getParticipantsNum() external view returns (uint)",
    "event Bid(address indexed account, uint amount, uint timestamp, uint indexed round)"
];
const provider = new ethers.JsonRpcProvider(
    `https://eth-sepolia.g.alchemy.com/v2/QtPw5bLONCtW00agVEhE66pb1Vsv9RnC`
);
const contractAddress = "0xE8f0b7144F2be28FE6Af69f15658d7b197Bf9f11";

const ParticipantsList: React.FC = () => {
    const [participants, setParticipants] = useState<{ address: string; bid: number }[]>([]);
    const [modalIsOpen, setModalIsOpen] = useState<boolean>(false); // Added: State for modal visibility
    const [selectedParticipant, setSelectedParticipant] = useState<{ address: string; bid: number } | null>(null); // Added: State to store selected participant

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

        const handleBidEvent = () => {
            console.log("Bid event detected, refetching participants...");
            fetchParticipants();
        };

        // Set up the listener
        contract.on("Bid", handleBidEvent);

        // Cleanup listener on component unmount
        return () => {
            contract.off("Bid", handleBidEvent);
        };
    }, [contract]);

    const truncateAddress = (address: string) => {
        return `${address.slice(0, 6)}...${address.slice(-4)}`.toLowerCase();
    };

    // Added: Function to show modal with participant details
    const showParticipantInfo = (participant: { address: string; bid: number }) => {
        console.log("Opening modal for participant:", participant);
        setSelectedParticipant(participant);
        setModalIsOpen(true);
    };

    // Added: Function to hide the modal
    const hideParticipantInfo = () => {
        setModalIsOpen(false);
        setSelectedParticipant(null);
    };

    return (
        <div className="lottery-container">
            <h1 className="lottery-title">Current Participants</h1>
            <ul className="list">
                {participants.map((participant, index) => (
                    <li
                        key={index}
                        className="list-item"
                        onClick={() => showParticipantInfo(participant)} // Updated: Show modal on click
                        style={{marginBottom: "8px", cursor: "pointer"}}
                    >
                        {truncateAddress(participant.address)} - {participant.bid} ticket(s)
                    </li>
                ))}
            </ul>
            {/* Added: Modal for showing participant details */}
            <Modal
                account={selectedParticipant?.address || ""}
                walletBalance={`${selectedParticipant?.bid || 0} Tickets`}
                isOpen={modalIsOpen}
                onClose={hideParticipantInfo}
                onDisconnect={() => {}} // No disconnect functionality for participants
                title="Participant Details" // Updated: Title for participants
                showDisconnect={false} // Updated: Hide Disconnect button
            />
        </div>
    );
};

export default ParticipantsList;