import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

require('dotenv').config()

const API_KEY = process.env.API_KEY;
const PRIVATE_KEY = process.env.PRIVATE_KEY;
const ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY;

module.exports = {
  defaultNetwork: "sepolia",
  networks: {
    hardhat: {
      
    },
    sepolia: {
      url: `https://eth-sepolia.g.alchemy.com/v2/${API_KEY}`,
      accounts: [PRIVATE_KEY],
      saveDeployments: true,
    },
    // etherscan: {
    //   apiKey: {
    //     sepolia: ETHERSCAN_API_KEY,
    //   }
    // }
  },
  solidity: "0.8.20",
  paths: {
    sources: "./contracts", // Папка с вашими контрактами
    artifacts: "./artifacts", // Папка для скомпилированных контрактов
  },
};

const config: HardhatUserConfig = {
  solidity: "0.8.28",
};

export default config;
