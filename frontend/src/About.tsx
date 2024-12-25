import React from 'react';
import "./About.css";
import about1 from "./assets/about1.jpg";
import about2 from "./assets/about2.jpg";
import about3 from "./assets/about3.jpg";


// 1.	Smart Contracts: Automated, self-executing contracts manage ticket sales, winner selection, and prize distribution, eliminating the need for a central authority. ￼
// 2.	Randomness and Fairness: Utilizing blockchain-based random number generators ensures unbiased and tamper-proof winner selection, fostering trust among participants. ￼
// 3.	Transparency and Security: All transactions are recorded on an immutable ledger, allowing participants to verify processes and outcomes, thereby enhancing security and reducing fraud. ￼


const About = () => {
  return (
    <div className="about-container">
        <div className='about-general-container'>
            <div className='about-general-text-container'>
                <p>A decentralized lottery operates transparently and autonomously through blockchain technology, ensuring fairness and security without intermediaries</p>
            </div>
        </div>
        <div className='about-explanation-contrainer'>
            <div className='about-explanation-sub-contrainer'>
                <img src={about1}></img>
                <div className="about-explanation-text-contrainer">
                    <p>Smart Contracts: Automated, self-executing contracts manage ticket sales, winner selection, and prize distribution, eliminating the need for a central authority</p>
                </div>
            </div>
            <div className='about-explanation-sub-contrainer'>
                <img src={about2}></img>
                <div className="about-explanation-text-contrainer">
                    <p>Randomness and Fairness: Utilizing blockchain-based random number generators ensures unbiased and tamper-proof winner selection, fostering trust among participants</p>
                </div>
            </div>
            <div className='about-explanation-sub-contrainer' style={{}}>
                <img src={about3}></img>
                <div className="about-explanation-text-contrainer">
                    <p>Transparency and Security: All transactions are recorded on an immutable ledger, allowing participants to verify processes and outcomes, thereby enhancing security and reducing fraud</p>
                </div>
            </div>
        </div>
    </div>
  );
};

export default About;