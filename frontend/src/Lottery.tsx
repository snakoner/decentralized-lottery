import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import { ethers } from "ethers";
import './index.css';
import "./Modal.css";
import robotLogo from "./assets/robot2.jpeg";
import ParticipantsList from "./ParticipantsList.tsx";
import { ALCHEMY_RPC_URL, CONTRACT_ADDRESS } from './constants.tsx';

const Modal = ({ isOpen, onClose, modalContent }) => {
    if (!isOpen) return null; // Не отображаем, если окно закрыто
    return ReactDOM.createPortal(        
        <div className="modal-overlay">
            <div className="modal-content">
                <div className="modal-wallet-disconnected-close">
                    <button onClick={onClose}>✖</button>
                </div>
                <div className="modal-wallet-disconnected-label">
                    <p>{modalContent}</p>
                </div>
                <div className="modal-wallet-disconnected-image" style={{
                    backgroundImage: `url(${robotLogo})`
                }}></div>
            </div>
        </div>,
        document.body
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
    const [modalContent, setModalContent] = useState<string|null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [withdrawButtonDisabled, setWithdrawButtonDisabled] = useState<boolean>(false);

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

    // contract get state
    const getTimeLeft = async() => {
        try {
            const num: bigint = await contractRpc.getTimeLeft();        
            setTimeLeft(Number(num));
        } catch (error) {
            console.log(error);
        }
    }

    const getParticipantsNum = async() => {
        console.log("beginning operation");
        try {
            const tempList: string[] = [];
            const num: bigint = await contractRpc.getParticipantsNum();
            console.log("Number received: ", num);
            setParticipants(Number(num));
            // Fetch participants list
            console.log("beginning that one operation");

            for (let i = 0; i < Number(num); i++) {
                const myAddress: string =  await contractRpc.participants(i);
                console.log(myAddress);
                tempList.push(myAddress);
            }

            setParticipantsList(tempList);
        } catch (error) {
            console.log("getParticipantsNum(): ", error);
        }
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
            if (Number(num) == 0) {
                setWithdrawButtonDisabled(true);
            }
            setUnlockedBalance(ethers.formatUnits(num));
        }
    }

    const withdraw = async() => {
        if (!connected) {
            setModalContent('Connect your wallet to withdraw');
            openModal();
            return;
        }

        try {
            const signer = await browserProvider.getSigner();
            const contract = new ethers.Contract(CONTRACT_ADDRESS, contractABI, signer);
            const tx = await contract.withdraw(account);
            await tx.wait();
            getBalanceToWithdraw();
        } catch(error) {
            console.log(error);
        }
    }

    const buyTicket = async() => {
        if (!connected) {
            setModalContent('Connect your wallet to buy tickets');
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
        getTimeLeft();
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
                    <p>👥 <strong>Tickets bought:</strong> {participants}</p>
                    <p>⏳ <strong>Time Left:</strong> {formatTime(timeLeft)}</p>
                    <p>⏳ <strong>Your tickets:</strong> {ticketNum}</p>
                </div>
                <input id="input-lottery" className="lottery-input" placeholder="1"></input>
                <button className="lottery-button" onClick={buyTicket}>
                    Buy Ticket
                </button>
                <Modal isOpen={isModalOpen} onClose={closeModal} modalContent={modalContent} />
                <div>
                    {loading ? <Spinner /> : <h1 id="tx-result"></h1>}
                </div>
            </div>
            <div className="lottery-container">
                <h1 className="lottery-title">Withdraw</h1>
                <div className="lottery-info">
                    <p>⏳ <strong>Available value:</strong> {unlockedBalance} ETH</p>
                </div>
                <button id="withdraw-button" className="lottery-button" onClick={withdraw} disabled={withdrawButtonDisabled}>
                    Withdraw
                </button>
            </div>
            {/* Add the new lists */}
            <ParticipantsList />
        </div>        
    );
};

export default LotteryStatus;