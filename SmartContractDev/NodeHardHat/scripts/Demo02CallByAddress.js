//imports
const { ethers } = require("hardhat")

//async main
async function main() {

    //合约地址可从区块链网络终端中复制获取
    //contract address 0xe7f1725e7734ce288f8367e1bb143e90bb3f0512

    const Demo02HelloContract = await ethers.getContractAt("Demo02HelloContract", "0xe7f1725e7734ce288f8367e1bb143e90bb3f0512")

    console.log(await Demo02HelloContract.getMessage())

}

//main
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });