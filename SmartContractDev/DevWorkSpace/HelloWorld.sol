// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract HelloWorld {
    // 私有状态变量，初始化赋值
    string private message = "Hello, Blockchain!"; 

    // 获取message值的函数
    function getMessage() public view returns (string memory) {
        return message;
    }

    // 修改message值的函数
    function setMessage(string memory newMessage) public {
        message = newMessage;
    }
}