import React, { useEffect, useState, useMemo } from 'react';
import { ethers } from 'ethers';
import Modal from "./Modal.tsx"; // Import the Modal component
import { ALCHEMY_RPC_URL, CONTRACT_ABI, CONTRACT_ADDRESS } from './constants.tsx';

const provider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);

const ParticipantsList: React.FC = () => {
    const [participants, setParticipants] = useState<{ address: string; tickets: number }[]>([]);
    const [modalIsOpen, setModalIsOpen] = useState<boolean>(false); // Modal visibility state
    const [selectedParticipant, setSelectedParticipant] = useState<{ address: string; tickets: number } | null>(null);

    // Instantiate the contract
    const contract = useMemo(() => new ethers.Contract(CONTRACT_ADDRESS, CONTRACT_ABI, provider), []);

    // Fetch participants and their ticket counts
    const fetchParticipants = async () => {
        try {
            console.log("Fetching participants...");
            const participantsCount: bigint = await contract.getParticipantsNumber();
            const currentRound: bigint = await contract.round(); // Fetch the current round

            const participantMap = new Map<string, number>(); // Use a Map to ensure uniqueness

            for (let i = 0; i < Number(participantsCount); i++) {
                const participantAddress: string = await contract.participants(i);
                const ticketCount: bigint = await contract.weights(currentRound, participantAddress);

                // Add the participant and ticket count to the Map
                participantMap.set(participantAddress, Number(ticketCount));
            }

            // Convert the Map to an array for rendering
            const participantsArray = Array.from(participantMap.entries()).map(([address, tickets]) => ({
                address,
                tickets,
            }));

            setParticipants(participantsArray); // Update state with unique participants
            console.log("Deduplicated Participants List:", participantsArray);
        } catch (error) {
            console.error("Error fetching participants:", error);
        }
    };
    useEffect(() => {
        fetchParticipants();

        // Listen for "Bid" events to update participants list in real-time
        const handleBidEvent = () => {
            console.log("Bid event detected, refetching participants...");
            fetchParticipants();
        };

        contract.on("Bid", handleBidEvent);

        // Cleanup listener on unmount
        return () => {
            contract.off("Bid", handleBidEvent);
        };
    }, [contract]);

    // Truncate address for cleaner display
    const truncateAddress = (address: string) => {
        return `${address.slice(0, 6)}...${address.slice(-4)}`.toLowerCase();
    };

    // Open the modal with selected participant's details
    const showParticipantInfo = (participant: { address: string; tickets: number }) => {
        console.log("Opening modal for participant:", participant);
        setSelectedParticipant(participant);
        setModalIsOpen(true);
    };

    // Close the modal
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
                        onClick={() => showParticipantInfo(participant)}
                        style={{ marginBottom: "8px", cursor: "pointer" }}
                    >
                        {truncateAddress(participant.address)} - 🎟️ {participant.tickets} ticket(s)
                    </li>
                ))}
            </ul>

            {/* Modal to display participant details */}
            <Modal
                account={selectedParticipant?.address || ""}
                walletBalance={`${selectedParticipant?.tickets || 0} 🎟️ Tickets`}
                isOpen={modalIsOpen}
                onClose={hideParticipantInfo}
                onDisconnect={() => {}} // No disconnect functionality for participants
                title="Participant Details"
                content={
                    <div>
                        <p>🎟️ <strong>Tickets bought:</strong> {selectedParticipant?.tickets} ticket(s)</p>
                    </div>
                }// Updated: Title for participants
                showDisconnect={false} // Updated: Hide Disconnect button
            />
        </div>
    );
};

export default ParticipantsList;
