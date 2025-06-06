const { ethers } = require("hardhat");

async function main() {
    // 获取合约工厂
    const Test = await ethers.getContractFactory("Test");
    // 部署合约
    const testContract = await Test.deploy();
    // 等待合约部署完成并获取部署收据
    const deploymentReceipt = await testContract.waitForDeployment();
    // 输出合约地址
    console.log("Test合约已部署至地址:", await testContract.getAddress());
}

main()
    // 无需额外的then/catch处理，直接调用main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("部署失败:", error);
        process.exit(1);
    });