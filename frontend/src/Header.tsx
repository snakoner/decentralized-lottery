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


  return (
      <header style={styles.header}>
        {/* Left section: Logo */}
        <div style={styles.leftSection}>
          <h1 style={styles.logo}>🎲 Lottery DApp</h1>
        </div>

        {/* Center section: Tabs */}
        <nav style={styles.centerSection}>
          <button style={styles.tabButton}>History</button>
          <button style={styles.tabButton}>About</button>
        </nav>

        {/* Right section: Wallet info */}
        <div style={styles.rightSection}>
          {!connected ? (
              <button style={styles.button} onClick={connectWallet}>
                Connect Wallet
              </button>
          ) : 
		  ( 
		  	<div>
				<Modal account={account} walletBalance={walletBalance} isOpen={modalIsOpen} onClose={hideWalletInfo} onDisconnect={disconnectWalletWrapper}></Modal>
				<div style={{display: "inline-block"}}>
					<button style={styles.button} onClick={showWalletInfo}>
						{slicedWalletAddress(account)}
					</button>
				</div>
			</div>
			)
          }
        </div>
      </header>
  );
};

const styles = {
  header: {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
    padding: "10px 20px",
    backgroundColor: "#282c34",
    color: "#ffffff",
	height: "100px",
    borderBottom: "2px solid #444",
    boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)",
  },

  leftSection: {
    flex: 1,
    display: "flex",
    alignItems: "center",
  },

  logo: {
    margin: 0,
    fontSize: "1.8rem",
    fontWeight: "bold",
    color: "#ffcc00",
    cursor: "pointer",
  },

  centerSection: {
    flex: 2,
    display: "flex",
    justifyContent: "center",
    gap: "20px",
  },

  tabButton: {
    background: "none",
    border: "none",
    color: "#ffffff",
    fontSize: "1rem",
    cursor: "pointer",
    padding: "5px 10px",
    transition: "color 0.3s",
  },

  tabButtonHover: {
    color: "#ffcc00",
    textDecoration: "underline",
  },

  rightSection: {
    flex: 1,
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-end",
  },

  walletInfo: {
    margin: "5px 0",
    fontSize: "0.9rem",
    color: "#d1d5db",
  },

  button: {
    padding: "10px 15px",
    fontSize: "1rem",
    backgroundColor: "#1263f1",
    color: "#ffffff",
    border: "none",
    borderRadius: "5px",
    cursor: "pointer",
    transition: "background-color 0.3s",
	width: "180px",
},

  disconnectButton: {
    padding: "5px 10px",
    fontSize: "0.9rem",
    backgroundColor: "#f44336",
    color: "#ffffff",
    border: "none",
    borderRadius: "5px",
    cursor: "pointer",
  },
};

export default Header;
