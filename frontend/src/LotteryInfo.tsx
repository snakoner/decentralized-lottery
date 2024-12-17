import React, {useState} from "react";
import "./LotteryInfo.css";
import robotImage from "./assets/robot.jpeg";
import trophyLogo from "./assets/trophy.png";
import ethereumLogo from "./assets/eth.png";
import { float } from "hardhat/internal/core/params/argumentTypes";

const LotteryInfo = () => {
    const [allTimeReward, setAllTimeReward] = useState<string|null>(null);
    const getAllTimeReward = async() => {
        try {
            const response = await fetch("http://0.0.0.0:8000/all-time-reward", {
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
            {/* Left Side */}
            <div className="lottery-welcome-left-side">
                <div className="lottery-welcome-label">
                    <p>Welcome to Decentralized Lottery</p>
                </div>
                <div className="lottery-welcome-description">
                    <p>
                        Experience now
                    </p>
                </div>
                <div className="lottery-welcome-statistics">
                    <div className="lottery-welcome-statistics-info">
                        <div
                            className="lottery-welcome-statistics-info-logo"
                            style={{ backgroundImage: `url(${trophyLogo})` }}
                        ></div>
                        <div className="lottery-welcome-statistics-info-text">
                            <p>All time reward:</p>
                        </div>
                        <div className="lottery-welcome-statistics-info-total-reward">
                            <p>{allTimeReward} ETH</p>
                        </div>
                    </div>
                    <div
                        className="lottery-welcome-statistics-info"
                        style={{ float: "right" }}
                    >
                        <div
                            className="lottery-welcome-statistics-info-logo"
                            style={{ backgroundImage: `url(${ethereumLogo})` }}
                        ></div>
                        <div className="lottery-welcome-statistics-info-text-powered">
                            <p>Build on Ethereum</p>
                        </div>
                    </div>
                </div>
            </div>

            {/* Image */}
            <img src={robotImage} className="lottery-welcome-image" alt="Robot" />
        </div>
    );
}

export default LotteryInfo;