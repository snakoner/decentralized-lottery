import React from 'react';

const ParticipantsList: React.FC = () => {
    const dummyParticipants = [
        '0x1234...abcd',
        '0x5678...efgh',
        '0x9abc...ijkl',
        '0xdef0...mnop',
    ];

    return (
        <div className="lottery-container">
            <h1 className="lottery-title">Current Participants</h1>
            <ul className="list">
                {dummyParticipants.map((participant, index) => (
                    <li key={index} className="list-item">
                        {participant}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ParticipantsList;
