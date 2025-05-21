//imports
const { ethers } = require("hardhat")

//async main
async function main() {

    //合约地址可从区块链网络终端中复制获取
    // 地址通常推荐使用 校验和地址（Checksum Address）校验和地址是通过特定算法对地址进行处理，使地址中的部分字符变为大写，这样可以在一定程度上防止地址输入错误
    // 可以在代码中使用一些库来自动将地址转换为校验和地址。例如，使用 ethereumjs - util 库中的 toChecksumAddress 函数
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