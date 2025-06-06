// 导入部署生成的合约JSON文件
const testContract = require("../artifacts/contracts/test.sol/test.json");

// 导出合约ABI和地址
module.exports = {
  abi: testContract.abi,
    contractAddress: "0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0" // 替换为实际合约地址
};
