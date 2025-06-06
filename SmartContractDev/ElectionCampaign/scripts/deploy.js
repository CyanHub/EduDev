const { ethers } = require("hardhat");

async function main() {
    // 获取合约工厂
    const Election = await ethers.getContractFactory("Election");

    // 部署合约
    const election = await Election.deploy();

    // 等待合约部署完成
    await election.waitForDeployment();

    console.log("部署的竞选合同为:", await election.getAddress());
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });