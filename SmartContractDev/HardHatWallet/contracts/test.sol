// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Test {
    uint[] public numbers; // 公开的整数数组

    // 构造函数：初始化数组（可自定义初始值）
    constructor() {
        numbers = [1, 2, 3, 4, 5]; // 示例数组，和为15
    }

    // 获取数组元素之和
    function getSum() public view returns (uint) {
        uint sum = 0;
        for (uint i = 0; i < numbers.length; i++) {
            sum += numbers[i];
        }
        return sum;
    }

    // 可选：添加元素到数组（用于测试交易）
    function addNumber(uint _num) public {
        numbers.push(_num);
    }
}