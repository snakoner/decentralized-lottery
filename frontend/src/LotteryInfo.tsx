import React, {useState} from "react";
import "./LotteryInfo.css";
import robotImage from "./assets/robot.jpeg";
import trophyLogo from "./assets/trophy.png";
import ethereumLogo from "./assets/eth.png";
import { float } from "hardhat/internal/core/params/argumentTypes";

const LotteryInfo = () => {
    const [allTimeReward, setAllTimeReward] = useState<string|null>("0");
    const getAllTimeReward = async() => {
        try {
            const response = await fetch("http://localhost:8000/all-time-reward", {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },     
            });
            
            const data = await response.json();
            console.log("all time reward: ", data['reward']);
            setAllTimeReward(data['reward'].slice(0, 7));
        } catch(error) {
            console.log(error);
        }

    }

    useState(() => {
        getAllTimeReward();
    }, []);

    return (
        <div className="lottery-welcome">
            <div className="lottery-welcome-left-side">
                <div className="lottery-welcome-label">
                    <p>Welcome to Decentralized Lottery</p>
                </div>
                <div className="lottery-welcome-description">
                    <p>
                        Experience a fair, transparent, and decentralized way to win big! Powered by blockchain, our lottery ensures trust and excitement for every participant. Join today and test your luck
                    </p>
                </div>
                <div className="lottery-welcome-statistics">
                    <div className="lottery-welcome-statistics-info">
                        <div className="lottery-welcome-statistics-info-logo" style={{
							backgroundImage: `url(${trophyLogo})`
						}}></div>
                        <div className="lottery-welcome-statistics-info-text">
                            <p>All time reward</p>
                            <p style={{fontSize: "16px", fontWeight: "100"}}>{allTimeReward} ETH</p>
                        </div>
                    </div>
                    <div className="lottery-welcome-statistics-info" style={{float:"right"}}>
                        <div className="lottery-welcome-statistics-info-logo" style={{
							backgroundImage: `url(${ethereumLogo})`
						}}></div>
                        <div className="lottery-welcome-statistics-info-text" style={{marginTop: "40px"}}>
                            <p>Build on Ethereum</p>
                        </div>

                    </div>
                    {/* <div className="lottery-welcome-statistics-info" style={{float:"right"}}></div> */}
                </div>
            </div>
            <img src={robotImage}></img>
        </div>
    );
}

export default LotteryInfo;