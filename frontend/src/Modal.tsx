import React from "react";
import ReactDOM from "react-dom";
import copyLogo from "./assets/copy.svg";
import etherscanLogo from "./assets/etherscan.svg";
import ethLogo from "./assets/eth.png";

import "./Modal.css";

interface ModalProps {
    account: string;
    walletBalance: string;
    isOpen: boolean;
    onClose: () => void;
    onDisconnect?: () => void; // Optional onDisconnect function
    title: string; // Added: Allows dynamic titles, e.g., "Your Wallet" or "Participant Details"
    showDisconnect?: boolean; // Added: Controls the visibility of the Disconnect button
}

const Modal: React.FC<ModalProps> = ({
                                         account,
                                         walletBalance,
                                         isOpen,
                                         onClose,
                                         onDisconnect,
                                         title, // Added: Dynamic title
                                         showDisconnect = true, // Added: Default is true, but can be toggled off
                                     }) => {
    if (!isOpen) return null;

    const copyWalletAddress = async () => {
        const walletAddress = document.getElementById('wallet-address')?.innerText;
        console.log(walletAddress);

        if (walletAddress) {
            try {
                await navigator.clipboard.writeText(walletAddress);
                const info = document.getElementById('modal-wallet-info');
                if (info) {
                    info.innerText = 'Copied';
                    setTimeout(() => {
                        info.innerText = '';
                    }, 1000);
                }
            } catch (err) {
                console.error('Error copying text: ', err);
            }
        }
    };

    return ReactDOM.createPortal(
        <div className="modal-overlay" onClick={onClose}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
                {/* Close Button */}
                <div className="modal-close">
                    <button onClick={onClose}>✖</button>
                </div>

                {/* Dynamic Title */}
                <div className="modal-wallet-label">
                    <p>{title}</p> {/* Updated: Dynamic title */}
                </div>

                {/* Address Section */}
                <div className="modal-wallet-address">
                    <p id="wallet-address">{account}</p>
                </div>

                {/* Balance Section */}
                <div className="modal-wallet-balance">
                    <p id="wallet-balance">
                        <strong>Balance:</strong> {walletBalance}
                    </p>
                    <div
                        className="modal-wallet-balance-eth-logo"
                        style={{ backgroundImage: `url(${ethLogo})` }}
                    ></div>
                </div>

                {/* Copy and Etherscan Links */}
                <div className="modal-wallet-copy-show">
                    <div onClick={copyWalletAddress}>
                        <div className="modal-wallet-copy">
                            <div
                                className="modal-wallet-copy-logo"
                                style={{ backgroundImage: `url(${copyLogo})` }}
                            ></div>
                            <p>Copy address</p>
                        </div>
                    </div>
                    <a
                        href={`https://etherscan.io/address/${account}`}
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <div className="modal-wallet-show">
                            <div
                                className="modal-wallet-show-logo"
                                style={{ backgroundImage: `url(${etherscanLogo})` }}
                            ></div>
                            <p>Show on Etherscan</p>
                        </div>
                    </a>
                </div>

                <div className="modal-wallet-info-area">
                    <p id="modal-wallet-info"></p>
                </div>

                {/* Disconnect Wallet Button */}
                {showDisconnect && onDisconnect && (
                    <div className="modal-disconnect-wallet">
                        <button onClick={onDisconnect}>Disconnect wallet</button>
                    </div>
                )}
            </div>
        </div>, document.body);
};

export default Modal;
