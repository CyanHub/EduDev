# WebSocket编程

1. 理解WebSocket协议原理：掌握 WebSocket 如何建立连接、传输数据。
2. 掌握WebSocket连接的生命周期：理解 WebSocket 连接从创建到关闭的各个阶段。
3. 实践WebSocket在线协作工具开发：能够使用 github.com/gorilla/websocket 库建立 WebSocket 连接，处理数据传输，以及管理连接的生命周期，能够使用
   WebSocket 实现一个简单的在线协作文档功能。

## 实验原理

本次实验基于WebSocket 协议，利用其在客户端和服务器之间建立持久连接并进行全双工通信的特性，使用 Go 语言的 gorilla/websocket 库实现一个简单的在线协作工具。

客户端通过 HTTP握手升级为 WebSocket 连接后，利用Goroutine 实现并发读写，将编辑操作实时同步到服务器，服务器再将更新广播给其他客户端，以此实现多人在线协作。

## 实验仪器与材料

1. 计算机：运行Windows、Linux或macOS操作系统的计算机。
2. Go语言环境：已安装并配置好Go语言环境，包括GOPATH、GOROOT等环境变量的设置。
3. 文本编辑器或IDE：如VS Code、GoLand等，用于编写Go代码。
4. Go语言基础知识：对Go语言的基本语法和编程概念有一定的了解。

## 实验内容步骤

1. 创建全局变量
   1. 定义upgrader变量，用于将 HTTP 连接升级为 WebSocket 连接。
   2. 定义clients变量，用于存储所有连接的客户端。
   3. 定义content变量，用于存储共享文档的内容。
2. 实现handleWebSocket函数：
   1. 升级HTTP连接：将HTTP连接升级为WebSocket连接
   2. 启动读写协程
3. 实现读函数：循环遍历客户端消息，将接收到的消息广播给其他客户端。
4. 实现广播函数：遍历所有的客户端连接，排除发送者，发送消息。


## 实验内容概要

这里面主要是通过ApiFox的配合，结合自身开发的各类包，再根据业务需求，自定义序列化器，实现灵活的数据存储格式转换。能够使用GORM 内置的 JSON 和 GOB 序列化器，以及自定义序列化器，处理复杂数据类型的存储。

## 实验环境

- 操作系统：Windows 10
- 开发环境：GoLand 2024.1.4
- 数据库：MySQL 8.0.36
- 第三方库：Zap 1.27.0
