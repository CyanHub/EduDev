//imports
const { ethers } = require("hardhat");

//async main
async function main() {
    const contractFactory = await ethers.getContractFactory("Demo01HelloWorld")

    const Demo01HelloWorld = await contractFactory.deploy()
}

//main
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });