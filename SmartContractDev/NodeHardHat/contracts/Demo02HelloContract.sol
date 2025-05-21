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