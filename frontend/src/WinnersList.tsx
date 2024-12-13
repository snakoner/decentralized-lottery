import React from 'react';

const WinnersList: React.FC = () => {
    const dummyWinners = [
        { round: 1, address: '0x1234...abcd', prize: '1.5 ETH' },
        { round: 2, address: '0x5678...efgh', prize: '0.8 ETH' },
        { round: 3, address: '0x9abc...ijkl', prize: '2.3 ETH' },
    ];

    return (
        <div className="lottery-container">
            <h1 className="lottery-title">Past Winners</h1>
            <ul className="list">
                {dummyWinners.map((winner, index) => (
                    <li key={index} className="list-item">
                        Round {winner.round}: {winner.address} won {winner.prize}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default WinnersList;
