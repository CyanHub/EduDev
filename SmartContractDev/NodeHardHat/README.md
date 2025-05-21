# Hardhat工程化

## 智能合约编辑工具

> 常见的合约编辑工具有 `remix`及VS `code IDE`

### Remix IDE简介

> `Remix` 是以太坊智能合约编程语言 `Solidity IDE`，其实基于浏览器的IDE，有一个很大的好处就是不用安装，打开即用。

官网 `https://remix.ethereum.org`

### VS code IDE

为避免翻墙，使用vs code编写solidity智能合约

安装solidity插件

`https://marketplace.visualstudio.com/items?itemName=NomicFoundation.hardhat-solidity`

## Hardhat

> 为了规范以太坊智能合约开发流程,也为了可以模拟复杂场景,简化部署流程 我们可以使用智能合约开发的工程化工具 `hardhat`

### Hardhat简介

官网: `https://hardhat.org/`

> Hardhat是一个编译、部署、测试和调试以太坊应用的开发环境。它可以帮助开发人员管理和自动化构建智能合约和dApps过程中固有的重复性任务，并围绕这一工作流程轻松引入更多功能。这意味着hardhat在最核心的地方是编译、运行和测试智能合约。
> Hardhat内置了Hardhat网络，这是一个专为开发设计的本地以太坊网络。主要功能有Solidity调试，跟踪调用堆栈、console.log()和交易失败时的明确错误信息提示等。

> Hardhat是目前最好的框架之一，支持快速测试，同时提供了最好的教程和最简单的集成。 老实说，每个喜欢JS框架的人都应该在某个时候试用Hardhat。它真的很容易上手，具有快速的测试， 而且入门非常简单。Hardhat的Discord也总是非常迅速地回答问题，因此，如果遇到问题，你 总是可以寻求帮助。Hathat使用Waffle和Ethers.js进行测试 —— 可以说是更好的JavaScript 智能合约框架 —— 开发人员的生活质量确实能得到一些改善。

### Hardhat项目初始化

```shell
//初始化Node项目
npm init

//安装Hardhat
npm install --save-dev hardhat

//启动并初始化Hardhat工程
npx hardhat
```

项目欢迎信息

```shell
888    888                      888 888               888
888    888                      888 888               888
888    888                      888 888               888
8888888888  8888b.  888d888 .d88888 88888b.   8888b.  888888
888    888     "88b 888P"  d88" 888 888 "88b     "88b 888
888    888 .d888888 888    888  888 888  888 .d888888 888
888    888 888  888 888    Y88b 888 888  888 888  888 Y88b.
888    888 "Y888888 888     "Y88888 888  888 "Y888888  "Y888

Welcome to Hardhat v2.22.15
```

### Hardhat项目结构

核心目录如下：

- `contracts`：智能合约目录
- `scripts` ：部署脚本文件
- `test`：智能合约测试用例文件夹。
- `hardhat.config.js`：配置文件，配置hardhat连接的网络及编译选项。

### 创建本地区块链网络

运行本地节点网络

```shell
-- 启动hardhat本地测试网络(以太坊协议标准)
npx hardhat node

-- 注意观察本地区块链节点监听端口
```

修改./hardhat.config.js配置文件，配置本地区块链网络

注意：本地区块链网络的链ID固定为31337(隐式创建，未日志提示)

```javascript
require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.27",
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545/",
      chainId: 31337,
    }
  }
};
```

### 编写合约

**强调**solidity智能合约的必要要素和编写规范

- 必须以许可证声明开头，不可插入其他内容
- 必须声明solidity编译版本
- 必须以contract语句块声明合约

*提出*在合约中可以使用hardhat自带的console进行日志输出

- hardhat/console.sol
- 显然此console与window.console不同
- 显然智能合约的运行时环境并非浏览器环境

创建 ./contracts/Demo01HelloWorld.sol 文件以编写合约

```js
//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Demo01HelloWorld {
    constructor(){
        console.log("hello world");
    }
}
```

### 编译合约

源码编写完成后，需要编译

```shell
npx hardhat compile
```

编译的产出将放置在 ./artifacts目录下

### 部署合约

智能合约区别于传统程序，它运行在区块链网络中，注意以下概念：

- 合约地址：每当合约安装到区块链网络中，会确定合约的唯一地址
- 区块链网络是开放的，网络的参与者都可以安装自己的功能
- 可以通过合约地址找到指定合约并调用

创建 ./scripts/Demo01Deploy.js文件用以部署合约

- ethers模块具有操作链上事务的能力
- ehters.getContractFactory()获取合约工厂，用以创建复数合约实例
- 合约工厂.deploy()用以部署合约，每部署一次会创建新的合约实例(工厂模式)
- 部署完成后，将获得该合约实例，后续可以通过这一实例调用合约方法

```javascript
//imports
const {ethers} = require("hardhat");

//async main
async function main() {
    const contractFactory = await ethers.getContractFactory("Demo01HelloWorld")

    const Demo01HelloWorld = await contractFactory.deploy()
}

//main
main()
.then(()=>process.exit(0))
.catch((error)=>{
    console.error(error);
    process.exit(1);
});
```

部署脚本封装完成后，执行脚本，将合约部署到本地区块链网络

```shell
npx hardhat run ./scripts/Demo01Deploy.js --network localhost
```

部署完成后，留意观察本地区块链网络终端，尤其注意获取**合约地址**

![image-20241108110319117](asset/image-20241108110319117.png)

### 调用合约

在本小节中，一方面快速重现合约的编写和部署流程，一方面引入进一步的特性：

- 采用getter和setter的经典模式引入变量、合约方法
- 在部署合约时立刻通过合约实例调用
- 在部署合约后，在其他文件中通过合约地址获取合约实例调用

创建 ./contracts/Demo02HelloContract.sol 文件以编写合约

```js
//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Demo02HelloContract {
    constructor(){
        //合约已部署
        console.log("Contract deployed");

        //初始化成员属性
        setMessage();
    }

    string message;

    function setMessage() public {
        message = "hello Contract!";
    }

    function getMessage() public view returns (string memory){
        return message;
    }
}
```

### 部署时调用合约

创建 ./scripts/Demo02CallByDeploy.js文件用以部署并调用合约

- 为避免反复获取合约地址，减少文件数量，这种形式在后文会频繁出现

```javascript
//imports
const {ethers} = require("hardhat");

//async main
async function main() {
    const contractFactory = await ethers.getContractFactory("Demo02HelloContract")

    const Demo02HelloContract = await contractFactory.deploy()
  
    console.log(await Demo02HelloContract.getMessage())
}

//main
main()
.then(()=>process.exit(0))
.catch((error)=>{
    console.error(error);
    process.exit(1);
});
```

脚本编写完成后，投放到本地区块链网络执行，观测结果。

```shell
npx hardhat run ./scripts/Demo02CallByDeploy.js --network localhost
```

注意到：

- 将Demo02合约的地址复制备用
- 区块链网络收到一个合约调用，名为Demo02#getMessage
- 脚本成功获取了消息文本并输出在右侧命令行
- 至此，完成了使用JS通过智能合约与区块链网络的交互

![image-20241108110917231](asset/image-20241108110917231.png)

### 通过地址调用合约

创建 ./scripts/Demo02CallByAddress.js文件用以通过合约地址调用合约

- ethers.getContractAt(合约名称,合约地址) 可从指定地址获取合约实例

```javascript
//imports
const { ethers } = require("hardhat")

//async main
async function main() {

    //合约地址可从区块链网络终端中复制获取
    //contract address 0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
  
    const Demo02HelloContract = await ethers.getContractAt("Demo02HelloContract","0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0")

    console.log(await Demo02HelloContract.getMessage())

}

//main
main()
.then(()=>process.exit(0))
.catch((error)=>{
    console.error(error);
    process.exit(1);
});
```

脚本编写完成后，投放到本地区块链网络执行，观测结果。

```shell
npx hardhat run ./scripts/Demo02CallByAddress.js --network localhost
```

得到相同执行效果

# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a Hardhat Ignition module that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat ignition deploy ./ignition/modules/Lock.js
```
