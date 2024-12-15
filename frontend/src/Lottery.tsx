import React, { useState, useEffect } from "react";
import { ethers } from "ethers";
import './index.css';
// import { bigint } from "hardhat/internal/core/params/argumentTypes";
import WinnersList from "./WinnersList.tsx";
import ParticipantsList from "./ParticipantsList.tsx";

const ALCHEMY_RPC_URL = `https://eth-sepolia.g.alchemy.com/v2/QtPw5bLONCtW00agVEhE66pb1Vsv9RnC`;
const CONTRACT_ADDRESS = "0xE8f0b7144F2be28FE6Af69f15658d7b197Bf9f11";

const Modal = ({ isOpen, onClose }) => {
    if (!isOpen) return null; // Не отображаем, если окно закрыто
    return (
        <div className="modal-overlay">
            <div className="modal-content">
                {/* <h2>Подключите кошелек</h2> */}
                <p>Подключите кошелек</p>
                <button onClick={onClose}>  ❌  </button>
            </div>
        </div>
    );
};

const Spinner = () => {
    return (
        <div className="spinner"></div>
    );
};

interface LotteryProps {
    connected: boolean;
    account: string;
}

const LotteryStatus: React.FC<LotteryProps> = ({connected, account}) => {
    const [ticketPriceWei, setTicketPriceWei] = useState<bigint>(0);
    const [ticketPrice, setTicketPrice] = useState<string>(""); 
    const [participants, setParticipants] = useState<number>(0);
    const [participantsList, setParticipantsList] = useState<string[]>([]); //ADDED
    const [timeLeft, setTimeLeft] = useState<number>(0);
    const [ticketNum, setTicketNum] = useState<number>(0);
    const [unlockedBalance, setUnlockedBalance] = useState<bigint>(0);
    const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
    const [loading, setLoading] = useState<boolean>(true);

    const openModal = () => setIsModalOpen(true);
    const closeModal = () => setIsModalOpen(false);

    const contractABI = [
        "function getTimeLeft() external view returns (uint)",
        "function getParticipantsNum() external view returns (uint)",
        //"function getParticipants() external view returns (address)",
        "function participants(uint256) external view returns (address)",
        "function getUnlockedBalance(address account) external view returns (uint)",
        "function getTicketNum(address account) external view returns (uint)",
        "function ticketPrice() public view returns (uint)",
        "function round() public view returns (uint)",        
        "function bid(uint amount) external payable",
        "function withdraw(address payable _to) external",
    ];

    const browserProvider = new ethers.BrowserProvider(window.ethereum);
    const providerRpc = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
    const contractRpc = new ethers.Contract(CONTRACT_ADDRESS, contractABI, providerRpc);

    const formatTime = (seconds: number) => {
        const hrs = Math.floor(seconds / 3600);
        const mins = Math.floor((seconds % 3600) / 60);
        const secs = seconds % 60;
        return `${hrs.toString().padStart(2, "0")}:${mins
            .toString()
            .padStart(2, "0")}:${secs.toString().padStart(2, "0")}`;
    };

    // const checkProviderConnection = async() => {
    //     const blockNumber = await providerRpc.getBlockNumber();
    //     console.log("current block: ", blockNumber);
    // }

    // contract get state
    const getParticipantsNumber = async() => {
        const num: bigint = await contractRpc.getTimeLeft();        
        setTimeLeft(Number(num));
    }

    const getParticipantsNum = async() => {
        console.log("beginning operation");
        const num: bigint = await contractRpc.getParticipantsNum();
        console.log("Number received");
        setParticipants(Number(num));
        // Fetch participants list
        console.log("beginning that one operation");
        const myAddress: string =  await contractRpc.participants(0);
        const tempList: string[] = [myAddress];
        console.log(myAddress);

        // for (let i = 0; i < participants; i++) {
        //     console.log("attempting to get a participant (retarded way)");
        //     const participantAddr = await contractRpc.getParticipants(i);
        //     console.log(participantAddr);
        //     tempList.push(participantAddr);
        // }
        setParticipantsList(tempList);
    }



    const getTicketPrice = async() => {
        const num: bigint = await contractRpc.ticketPrice();
        setTicketPriceWei(num);
        setTicketPrice(ethers.formatEther(num) + " ETH");
    }

    // const getRound = async() => {
    //     const num: bigint = await contractRpc.round();
    //     return num;
    // }

    const getTicketsByUser = async() => {
        // let _account;
        // _account = await getAccount();
        
        if (account !== null) {
            const num: bigint = await contractRpc.getTicketNum(account);
            setTicketNum(Number(num));
        }
    }

    const getBalanceToWithdraw = async() => {
        if (account !== null) {
            const num: bigint = await contractRpc.getUnlockedBalance(account);
            setUnlockedBalance(ethers.formatUnits(num));
        }
    }

    // const postBuyTicket = async(fromAddress: string, ticketNum: number) => {
    //     try {
    //         console.log(JSON.stringify({
    //             fromAddress: fromAddress,
    //             ticketNum: ticketNum,
    //         }));
    //         const response = await fetch("http://0.0.0.0:8000/bid", {
    //             method: 'POST',
    //             headers: {
    //                 'Content-Type': 'application/json',
    //             },     
    //             body: JSON.stringify({
    //                 fromAddress: fromAddress,
    //                 ticketNum: ticketNum,
    //             }),
    //         });
            
    //         // const data = await response.json();
    //         console.log('transaction sent');
    //     } catch(error) {
    //         console.log(error);
    //     }
    // };

    const withdraw = async() => {
        if (!connected) {
            openModal();
            return;
        }

        if (Number(unlockedBalance) === 0) {
            console.log("zero balance to withdraw");
            return;
        }

        const signer = await browserProvider.getSigner();
        const contract = new ethers.Contract(CONTRACT_ADDRESS, contractABI, signer);
        
        try {
            const tx = await contract.withdraw(account);
            await tx.wait();
            getBalanceToWithdraw();
        } catch(error) {
            console.log(error);
        }
    }

    const buyTicket = async() => {
        if (!connected) {
            openModal();
            return;
        }

        const input = document.getElementById("input-lottery");
        let ticketNumber: number = 0;
        if (input) {
            const inputValue = input.value;
            if (inputValue === "") {
                ticketNumber = 1;
            } else {
                ticketNumber = Number(inputValue);
            }
        }

        const txParams = {
            value: ticketPriceWei * ethers.toBigInt(ticketNumber),
        };

        console.log(txParams);
        // postBuyTicket(account, ticketNumber);
        const signer = await browserProvider.getSigner();
        const contract = new ethers.Contract(CONTRACT_ADDRESS, contractABI, signer);
        try {
            setLoading(true);
            const tx = await contract.bid(ethers.toBigInt(ticketNumber), txParams);        
            await tx.wait();
            const txResult = document.getElementById('tx-result');
            if (txResult) {
                txResult.innerHTML = '✅ Tx completed';
            }

            getTicketsByUser();
        } catch (error) {
            console.log(error);
        }    

        setLoading(false);
    }

    // internal
    useEffect(() => {
        console.log("lotteryStatus() = ", connected);

        // contract interaction
        setLoading(false);
        getParticipantsNumber();
        getParticipantsNum();
        getTicketPrice();
        getTicketsByUser();
        getBalanceToWithdraw();

        const timer = setInterval(() => {
            setTimeLeft((prevTime) => (prevTime > 0 ? prevTime - 1 : 0));
        }, 1000);
        return () => clearInterval(timer);
    }, [account]);

    return (
        <div className="main-container">
            <div className="lottery-container">
                <h1 className="lottery-title">Lottery Status</h1>
                <div className="lottery-info">
                    <p>🎟️ <strong>Ticket Price:</strong> {ticketPrice}</p>
                    <p>👥 <strong>Participants:</strong> {participants}</p>
                    <p>⏳ <strong>Time Left:</strong> {formatTime(timeLeft)}</p>
                    <p>⏳ <strong>Your tickets:</strong> {ticketNum}</p>
                </div>
                <input id="input-lottery" className="lottery-input" placeholder="1"></input>
                <button className="lottery-button" onClick={buyTicket}>
                    Buy Ticket
                </button>
                <Modal isOpen={isModalOpen} onClose={closeModal} />
                <div>
                    {loading ? <Spinner /> : <h1 id="tx-result"></h1>}
                </div>
            </div>
            <div className="lottery-container">
                <h1 className="lottery-title">Withdraw</h1>
                <div className="lottery-info">
                    <p>⏳ <strong>Available value:</strong> {unlockedBalance}</p>
                </div>
                <button id="withdraw-buttion" className="lottery-button" onClick={withdraw}>
                    Withdraw
                </button>
            </div>
            {/* Add the new lists */}
            <ParticipantsList />
            <WinnersList />
        </div>        
    );
};

export default LotteryStatus;