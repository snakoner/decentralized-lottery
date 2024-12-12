import React, { useState, useEffect } from "react";
import { ethers } from "ethers";

const Header = ({ connected, account, network, error, connectWallet, disconnectWallet, walletBalance }) => {
	const slicedWalletAddress = ()=> {
		if (!account) {
			return "";
		}

		return account.slice(0, 6) + "..." + account.slice(-4);
	};

	return (
    <header style={styles.header}>
      <div style={styles.leftSection}>
        <h1 style={styles.title}>🎲 Lottery DApp</h1>
        <p style={styles.info}>
          <strong>Wallet balance:</strong> {walletBalance || "Connect wallet"}
        </p>
        <p style={styles.info}>
          <strong>Network:</strong> {network || "Unknown"}
        </p>
      </div>
      <div style={styles.rightSection}>
        {!connected ? (
          <button style={styles.button} onClick={connectWallet}>
            Connect Wallet
          </button>
        ) : (
          <div>
            <p style={styles.walletInfo}>
              <strong>Wallet:</strong> {slicedWalletAddress()}
            </p>
            <button style={styles.disconnectButton} onClick={disconnectWallet}>
              Disconnect
            </button>
          </div>
        )}
      </div>
    </header>
  );
};

const styles = {
  header: {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
    padding: "20px",
    backgroundColor: "#282c34",
    color: "#ffffff",
    borderBottom: "2px solid #444",
    boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)",
  },
  leftSection: {
    flex: 1,
  },
  title: {
    margin: 0,
    fontSize: "1.8rem",
    fontWeight: "bold",
    color: "#ffcc00",
  },
  info: {
    margin: "5px 0",
    fontSize: "1rem",
    color: "#d1d5db",
  },
  rightSection: {
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-end",
  },
  walletInfo: {
    marginBottom: "5px",
    fontSize: "0.9rem",
    color: "#d1d5db",
  },
  button: {
    padding: "10px 15px",
    fontSize: "1rem",
    backgroundColor: "#4caf50",
    color: "#ffffff",
    border: "none",
    borderRadius: "5px",
    cursor: "pointer",
    transition: "background-color 0.3s",
  },
  buttonHover: {
    backgroundColor: "#45a049",
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