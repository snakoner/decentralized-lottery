import React, {useState} from "react";
import "./LotteryInfo.css";
import robotImage from "./assets/robot.jpeg";
import trophyLogo from "./assets/trophy.png";
import ethereumLogo from "./assets/eth.png";
import { float } from "hardhat/internal/core/params/argumentTypes";

const LotteryInfo = () => {
    const getAllTimeReward = async() => {
        try {
            const response = await fetch("http://0.0.0.0:8000/all-time-reward", {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },     
                // body: JSON.stringify({}),
            });
            
            const data = await response.json();
            // console.log('transaction sent');
        } catch(error) {
            console.log(error);
        }

    }

    return (
        <div class="lottery-welcome">
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
                            <p>All time reward:</p>
                        </div>
                        <div className="lottery-welcome-statistics-info-total-reward">
                            <p>0.0001 ETH</p>
                        </div>
                    </div>

                    <div className="lottery-welcome-statistics-info" style={{float:"right"}}>
                        <div className="lottery-welcome-statistics-info-logo" style={{
							backgroundImage: `url(${ethereumLogo})`
						}}></div>
                        <div className="lottery-welcome-statistics-info-text-powered">
                            <p className="lottery-welcome-statistics-info-text-powered">Build on Ethereum</p>
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