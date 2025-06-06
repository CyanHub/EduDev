require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks: {
    hardhat: {
      chainId: 31337,  // HardHat默认链ID
    },
    local: {
      url: "http://127.0.0.1:8545", // 本地节点RPC地址
      chainId: 31337,  // 本地网络默认链ID
    },
  },
};