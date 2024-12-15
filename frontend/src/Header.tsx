import React, { useState } from "react";

import "./Header.css";
import Modal from "./Modal.tsx"; // Import the Modal component

interface HeaderProps {
    connected: boolean;
    account: string;
    network: string;
    error: string;
    connectWallet: () => void;
    disconnectWallet: () => void;
    walletBalance: string;
}

const Header: React.FC<HeaderProps> = ({ connected, account, network, error, connectWallet, disconnectWallet, walletBalance }) => {
    const [modalIsOpen, setModalIsOpen] = useState<boolean>(false);

    const slicedWalletAddress = () => {
        if (!account) {
            return "";
        }

        return account.slice(0, 6) + "..." + account.slice(-4);
    };

    const showWalletInfo = () => {
        console.log("show ");
        setModalIsOpen(true);
    };

    const hideWalletInfo = () => {
        console.log("close");
        setModalIsOpen(false);
    }

    const disconnectWalletWrapper = () => {
        setModalIsOpen(false);
        disconnectWallet();
    };

    return (
        <header className="header">
            {/* Left section: Logo */}
            <div className="left-section">
                <h1 className="logo">🎲 Lottery DApp</h1>
            </div>

            {/* Center section: Tabs */}
            <nav className="center-section">
                <button className="tab-button">History</button>
                <button className="tab-button">About</button>
            </nav>

            {/* Right section: Wallet info */}
            <div className="right-section">
                {!connected ? (
                    <button className="button" onClick={connectWallet}>
                        Connect Wallet
                    </button>
                ) : (
                    <div>
                        <Modal
                            account={account}
                            walletBalance={walletBalance}
                            isOpen={modalIsOpen}
                            onClose={hideWalletInfo}
                            onDisconnect={disconnectWalletWrapper}
                            title="Your Wallet"
                            showDisconnect={true}
                        />
                        <div className="wallet-button-container">
                            <button className="button" onClick={showWalletInfo}>
                                {slicedWalletAddress()}
                            </button>
                        </div>
                    </div>
                )}
            </div>
        </header>
    );
};

export default Header;
