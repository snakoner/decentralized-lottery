import React, {useState} from "react";
import copyLogo from "./assets/copy.svg";
import etherscanLogo from "./assets/etherscan.svg";
import ethLogo from "./assets/eth.png";

import "./Header.css";
import { float } from "hardhat/internal/core/params/argumentTypes";

const Modal = ({ account, walletBalance, isOpen, onClose, onDisconnect }) => {
	if (!isOpen) 
		return null;
  
	const copyWalletAddress = async() => {
		const walletAddress = document.getElementById('wallet-address')?.innerHTML;
		console.log(walletAddress);

		if (walletAddress) {
			try {
				await navigator.clipboard.writeText(walletAddress);
				const info = document.getElementById('modal-wallet-info');
				if (info) {
					info.innerHTML = 'Copied';
					setTimeout(()=>{
						info.innerHTML = '';
					}, 1000);
				}

			} catch (err) {
				console.error('Error copying text: ', err);
			}	
		}
	}

	return (
	  <div className="modal-overlay" onClick={onClose}>
		<div className="modal-content" onClick={(e) => e.stopPropagation()}>
			<div className="modal-close">
				<button onClick={onClose}>✖</button>
			</div>
			<div className='modal-wallet-label'>
				<p>Your wallet</p>
			</div>
			<div className='modal-wallet-address'>
				<p id='wallet-address'>{account}</p>
			</div>
			<div className='modal-wallet-balance'>
				<p id='wallet-balance'><strong>Balance:</strong> {walletBalance}</p>
				<div className="modal-wallet-balance-eth-logo" 
					style={{
							backgroundImage: `url(${ethLogo})`
					}}
				></div>
			</div>
			<div className='modal-wallet-copy-show'>
				<div onClick={copyWalletAddress}>
					<div className='modal-wallet-copy'>
						<div className='modal-wallet-copy-logo' style={{
							backgroundImage: `url(${copyLogo})`
						}}></div>
						<p>Copy address</p>
					</div>
				</div>
				<a href={`https://etherscan.io/address/${account}`} target="_blank">
				<div className='modal-wallet-show'>
					<div className='modal-wallet-show-logo' style={{
							backgroundImage: `url(${etherscanLogo})`
						}}></div>
					<p>Show on Etherscan</p>
				</div>
				</a>
			</div>
			<div className='modal-wallet-info-area'>
				<p id='modal-wallet-info'></p>
			</div>
			<div className='modal-disconnect-wallet'>
				<button onClick={onDisconnect}>Disconnect wallet</button>
			</div>
		</div>
	  </div>
	);
};

const Header = ({ connected, account, network, error, connectWallet, disconnectWallet, walletBalance }) => {
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


	const disconnectWalletWrapper = ()=> {
		setModalIsOpen(false);
		disconnectWallet();
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
