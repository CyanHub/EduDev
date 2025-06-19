# 课程设计项目：

## 选题一：宠物寄养系统后端开发

### 项目概述

1. **宠物寄养系统**旨在为**宠物主人**和**寄养服务提供者**搭建一个便捷的**在线平台**，主要功能包**括宠物信息管理**、**寄养服务预订**、**订单处理**以及**评价反馈**。对于宠物主人来说，他们可以**通过该系统注册账号**，录入**宠物的基本信息**如**品种**、**年龄**等，并**根据自身需求选择合适的寄养服务**进行**预订及在线支付**。另外，宠物主人还能在**服务结束后对服务进行评价**。
2. **寄养服务提供者**则可以**通过系统发布他们的服务详情**，比如**价格**、**服务内容**和**位置**等信息，并管理**来自宠物主人**的**预订请求**。提供者还需要**记录每日宠物的护理情况**，确保宠物得到妥善照顾。**管理员角色**负责**审核用户**和**服务提供者**的**注册信息**，保证**所有发布的服务符合平台规则**，并**监控系统的运行状态**以优化用户体验。
3. 系统的**核心功能**涵盖了**搜索**与**筛选寄养服务**，**订单变动提醒**、**日程安排**帮助协调**最佳寄养时间等各种服务相关的通知**，以及**建立公平透明的评价体系**促进服务质量提升。同时，**系统还需考虑**、**可靠性**、**扩展性**和**响应速度**等**非功能性需求**，确保**良好的用户体验**和**系统稳定性**。**``最终目标是创建一个既安全又高效的宠物寄养服务平台，满足宠物主人和服务提供者的多样化需求。``**

## 选题二：自立项目。要求项目要有一定的复杂度，涵盖本门课所学知识。

**无论选题一还是选题二 都要能体现以下考核重点：**

### 01-分布式微服务架构考核重点：

- **能清晰划分出微服务**
- **能根据划分出的微服务来画出正确的架构图**
- **能清晰写出划分原由**

### 02-基于微服务通信考核重点：

- **能正确选择并实现服务间通信方式（如RestfulApi、GRPC、消息队列）**
- **能合理使用API网关进行请求路由**

### 03-分布式事务与分布式锁考核重点：

- **能识别需要保证数据一致性的业务场景**
- **能正确应用分布式事务解决方案**
- **能合理使用分布式锁（如Redis锁）防止并发冲突**

### 04-集群部署考核重点：

- **能搭建多节点集群环境（如Nginx负载均衡）**
- **能实现服务的高可用部署**
- **能对部署过程进行文档记录和总结**

### 05-数据库分表及中间件使用考核重点：

- **能识别出需要分表的应用场景**
- **能根据业务需求进行水平/垂直分表设计**
- **能结合缓存中间件提升系统性能**

## 需要提交的考查项目资料

### （注意，无论是选题一还是选题二，都要能体现考核重点，详情看下面考核重点）：

#### 1、``课程设计报告 ``；

#### 2、``源代码 ``；

#### 3、``源代码安装部署说明文档 ``（根据这个文档 ``要能安装部署运行成功 ``）；

#### 4、``测试、演示视频 ``。

---

# 开发文档

## 一、项目概述

宠物寄养系统后端开发项目旨在为宠物主人和寄养服务提供者搭建一个便捷的在线平台。系统主要功能包括宠物信息管理、寄养服务预订、订单处理以及评价反馈。同时考虑了系统的可靠性、扩展性和响应速度等非功能性需求，确保良好的用户体验和系统稳定性。

## 二、系统架构设计

1. ### **微服务划分**

   * 用户服务 (User Service)：负责用户注册、登录、认证和用户信息管理。
   * 宠物服务 (Pet Service)：管理宠物信息，包括品种、年龄等基本信息。
   * 订单服务 (Order Service)：处理预订、支付和订单状态管理。
   * 评价服务 (Review Service)：管理用户对寄养服务的评价和反馈。
   * 网关服务 (Gateway Service)：通过API网关（Nginx）进行请求路由。
2. ### **系统架构图**

```plaintext
                  ┌─────────────────┐
                  │   API Gateway    │
                  │     (Nginx)      │
                  └─────────────────┘
                          │
          ┌───────────────┼───────────────┐
          │               │               │
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│   User Service  │ │   Pet Service   │ │  Order Service  │
└─────────────────┘ └─────────────────┘ └─────────────────┘
          │               │               │
          └───────────────┼───────────────┘
                          │
          ┌───────────────┼───────────────┐
          │               │               │
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│  Review Service │ │  MySQL Cluster  │ │  Redis Cluster  │
└─────────────────┘ └─────────────────┘ └─────────────────┘
```

3. ### **微服务划分原由**

   * 按业务功能划分：每个微服务对应系统中的一个核心业务功能，保持单一职责原则。
   * 数据自治：每个微服务管理自己的数据，减少服务间强依赖。
   * 团队协作：不同团队可以独立开发和维护不同的微服务。
   * 技术异构：允许不同微服务使用最适合其业务需求的技术栈。
   * 独立部署：每个微服务可以独立部署和扩展，提高系统弹性。

## 三、技术选型

1. ### **开发语言和框架**

   * 编程语言：Go语言
   * Web框架：Gin（用于创建RESTful API端点）
   * 数据库驱动：`go-sql-driver/mysql`
2. ### **微服务通信**

   * 同步通信：RESTful API（基于HTTP协议的服务间通信）
3. ### **数据存储**

   * 关系型数据库：MySQL（用于存储结构化数据）
   * 缓存：Redis（用于缓存热点数据、实现分布式锁）
4. ### **部署和运维**

   * 负载均衡：Nginx（实现服务的高可用和负载均衡）

## 四、数据库设计

1. ### **数据库分表策略**

   * **水平分表** ：订单表按用户ID哈希值分表，共8个表，表名格式为 `order_0` 到 `order_7`，路由规则为 `user_id % 8`。
   * **垂直分表** ：宠物信息表分为基础信息表 `pet_basic` 和健康档案表 `pet_health`。
2. ### **主要数据表设计**

- **用户表 (`users`)**

```sql
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

- **宠物基础信息表 (`pet_basic`)**

```sql
CREATE TABLE pet_basic (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    type VARCHAR(20) NOT NULL,
    age INT,
    owner_id BIGINT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

- **订单分表 (`order_0` 到 `order_7`)**

```sql
CREATE TABLE order_0 (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    pet_id BIGINT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (pet_id) REFERENCES pet_basic (id)
);
```

## 五、API 文档

详细的API文档可参考 `docs/api.md` 文件，主要API接口如下：

1. ### **用户服务**
   
   - 用户注册：`POST /register`
   - 用户登录：`POST /login`
2. ### **宠物服务**
   
   - 添加宠物：`POST /pets`
   - 获取宠物详情：`GET /pets/:id`
   - 更新宠物：`PUT /pets/:id`
   - 删除宠物：`DELETE /pets/:id`
3. ### **订单服务**
   
   - 创建订单：`POST /orders`
   - 支付订单：`PUT /orders/:id/pay`
4. ### **评价服务**
   
   - 添加评价：`POST /reviews`
   - 获取评价详情：`GET /reviews/:id`
   - 更新评价：`PUT /reviews/:id`
   - 删除评价：`DELETE /reviews/:id`

## 六、项目结构

```plaintext
.vscode/
README.md
api/
`-- pet.html
cmd/
|-- gateway/
|-- order_service/
|-- pet_service/
|-- review_service/
`-- user_service/
config/
`-- database.yaml
database/
`-- pet_boarding.sql
deploy/
|-- nginx.conf
`-- sharding_strategy.md
docs/
`-- api.md
go.mod
tools/
`-- replace_module.go
```

## 七、开发环境搭建

1. **安装依赖** 确保已经安装Go语言环境，然后在项目根目录下执行以下命令下载依赖：

```go
go mod tidy
```

2. **配置数据库** 修改 `config/database.yaml` 文件中的数据库连接信息，确保MySQL和Redis服务正常运行。
3. **启动服务** 分别进入各个微服务目录，执行以下命令启动服务：

```
go run main.go
```

## 八、部署说明

1. **容器化部署** 可以使用Docker将各个微服务打包成镜像，使用Docker Compose进行多容器管理。
2. **负载均衡** 使用Nginx作为负载均衡器，配置文件可参考 `deploy/nginx.conf`。



