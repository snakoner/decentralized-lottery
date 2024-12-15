import React from "react";

const Header = ({ connected, account, network, error, connectWallet, disconnectWallet, walletBalance }) => {
  const slicedWalletAddress = () => {
    if (!account) {
      return "";
    }

    return account.slice(0, 6) + "..." + account.slice(-4);
  };

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
          ) : (
              <div>
                <p style={styles.walletInfo}>
                  <strong>Wallet:</strong> {slicedWalletAddress()}
                </p>
                <p style={styles.walletInfo}>
                  <strong>Balance:</strong> {walletBalance || "N/A"} ETH
                </p>
                <p style={styles.walletInfo}>
                  <strong>Network:</strong> {network || "Unknown"}
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
    padding: "10px 20px",
    backgroundColor: "#282c34",
    color: "#ffffff",
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
    backgroundColor: "#4caf50",
    color: "#ffffff",
    border: "none",
    borderRadius: "5px",
    cursor: "pointer",
    transition: "background-color 0.3s",
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
