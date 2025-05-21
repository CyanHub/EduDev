//imports
const { ethers } = require("hardhat");

//async main
async function main() {
    const contractFactory = await ethers.getContractFactory("Demo02HelloContract")

    const Demo02HelloContract = await contractFactory.deploy()

    console.log(await Demo02HelloContract.getMessage())
}

//main
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });