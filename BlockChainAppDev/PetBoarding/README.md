# 课程设计项目：

## 选题一：宠物寄养系统后端开发

### 项目概述

1. **宠物寄养系统**旨在为**宠物主人**和**寄养服务提供者**搭建一个便捷的**在线平台**，主要功能包**括宠物信息管理**、**寄养服务预订**、**订单处理**以及**评价反馈**。对于宠物主人来说，他们可以**通过该系统注册账号**，录入**宠物的基本信息**如**品种**、**年龄**等，并**根据自身需求选择合适的寄养服务**进行**预订及在线支付**。另外，宠物主人还能在**服务结束后对服务进行评价**。
2. **寄养服务提供者**则可以**通过系统发布他们的服务详情**，比如**价格**、**服务内容**和**位置**等信息，并管理**来自宠物主人**的**预订请求**。提供者还需要**记录每日宠物的护理情况**，确保宠物得到妥善照顾。**管理员角色**负责**审核用户**和**服务提供者**的**注册信息**，保证**所有发布的服务符合平台规则**，并**监控系统的运行状态**以优化用户体验。
3. 系统的**核心功能**涵盖了**搜索**与**筛选寄养服务**，**订单变动提醒**、**日程安排**帮助协调**最佳寄养时间的各种服务相关的通知**，以及**建立公平透明的评价体系**促进服务质量提升。同时，**系统还需考虑**、**可靠性**、**扩展性**和**响应速度**等**非功能性需求**，确保**良好的用户体验**和**系统稳定性**。**`<u>`最终目标是创建一个既安全又高效的宠物寄养服务平台，满足宠物主人和服务提供者的多样化需求。`</u>`**

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

#### 1、`<u>`课程设计报告 `</u>`；

#### 2、`<u>`源代码 `</u>`；

#### 3、`<u>`源代码安装部署说明文档 `</u>`（根据这个文档 `<u>`要能安装部署运行成功 `</u>`）；

#### 4、`<u>`测试、演示视频 `</u>`。



---



# 宠物寄养系统 (Pet Boarding System)

## 项目实现概述

本项目是基于微服务架构的宠物寄养系统后端实现，采用Go语言和Gin框架开发。系统划分为多个微服务，每个服务负责特定的业务功能，通过API网关统一对外提供服务。

### 已实现的微服务

1. **用户服务 (User Service)**：管理用户注册、登录、个人信息等功能
2. **宠物服务 (Pet Service)**：管理宠物信息的添加、修改、查询等功能
3. **寄养服务 (Boarding Service)**：处理寄养预订、查询、修改、取消等功能
4. **评价服务 (Review Service)**：管理用户对寄养服务的评价和评分
5. **通知服务 (Notification Service)**：处理系统通知和消息推送
6. **管理员服务 (Admin Service)**：提供系统管理和监控功能
7. **API网关 (API Gateway)**：统一接口入口，负责请求路由和负载均衡

### 技术栈

- **后端**：Go语言 + Gin框架
- **容器化**：Docker + Docker Compose
- **服务通信**：RESTful API

### 服务端口分配

- API网关：8080
- 用户服务：8081
- 宠物服务：8082
- 寄养服务：8083
- 评价服务：8084
- 通知服务：8085
- 管理员服务：8086

## 快速开始

### 前提条件

- 安装Docker和Docker Compose
- 安装Go 1.21或更高版本（如需本地开发）

### 启动服务

1. 克隆仓库
   ```bash
   git clone https://github.com/yourusername/pet-boarding.git
   cd pet-boarding
   ```

2. 使用Docker Compose启动所有服务
   ```bash
   docker-compose up -d
   ```

3. 访问API网关
   ```
   http://localhost:8080
   ```

### 本地开发

1. 进入服务目录
   ```bash
   cd services/user-service
   ```

2. 安装依赖
   ```bash
   go mod download
   ```

3. 运行服务
   ```bash
   go run cmd/user/main.go
   ```

## 项目结构

```
├── api-gateway/            # API网关服务
├── services/               # 微服务目录
│   ├── user-service/       # 用户服务
│   ├── pet-service/        # 宠物服务
│   ├── boarding-service/   # 寄养服务
│   ├── review-service/     # 评价服务
│   ├── notification-service/ # 通知服务
│   └── admin-service/      # 管理员服务
├── docker/                 # Docker相关配置
├── config/                 # 配置文件
└── docker-compose.yml      # Docker Compose配置
```

## 服务间通信

服务间通信目前通过HTTP REST API实现。未来计划引入gRPC进行服务间的高效通信。

## 数据持久化

当前版本使用内存存储数据，仅用于演示。生产环境将使用MySQL数据库进行数据持久化。

## 未来计划

1. 集成MySQL数据库
2. 添加Redis缓存
3. 实现Kafka消息队列
4. 添加用户认证和授权
5. 实现服务发现和注册
6. 添加监控和日志系统
7. 开发前端界面

---

# 课程设计项目：

## 选题一：宠物寄养系统后端开发

### 项目概述

1. **宠物寄养系统**旨在为**宠物主人**和**寄养服务提供者**搭建一个便捷的**在线平台**，主要功能包**括宠物信息管理**、**寄养服务预订**、**订单处理**以及**评价反馈**。对于宠物主人来说，他们可以**通过该系统注册账号**，录入**宠物的基本信息**如**品种**、**年龄**等，并**根据自身需求选择合适的寄养服务**进行**预订及在线支付**。另外，宠物主人还能在**服务结束后对服务进行评价**。
2. **寄养服务提供者**则可以**通过系统发布他们的服务详情**，比如**价格**、**服务内容**和**位置**等信息，并管理**来自宠物主人**的**预订请求**。提供者还需要**记录每日宠物的护理情况**，确保宠物得到妥善照顾。**管理员角色**负责**审核用户**和**服务提供者**的**注册信息**，保证**所有发布的服务符合平台规则**，并**监控系统的运行状态**以优化用户体验。
3. 系统的**核心功能**涵盖了**搜索**与**筛选寄养服务**，**订单变动提醒**、**日程安排**帮助协调**最佳寄养时间的各种服务相关的通知**，以及**建立公平透明的评价体系**促进服务质量提升。同时，**系统还需考虑**、**可靠性**、**扩展性**和**响应速度**等**非功能性需求**，确保**良好的用户体验**和**系统稳定性**。**`<u>`最终目标是创建一个既安全又高效的宠物寄养服务平台，满足宠物主人和服务提供者的多样化需求。`</u>`**

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

#### 1、`<u>`课程设计报告 `</u>`；

#### 2、`<u>`源代码 `</u>`；

#### 3、`<u>`源代码安装部署说明文档 `</u>`（根据这个文档 `<u>`要能安装部署运行成功 `</u>`）；

#### 4、`<u>`测试、演示视频 `</u>`。



---



# 宠物寄养系统后端开发方案

## 一、系统架构设计

### 1. 微服务划分

根据宠物寄养系统的业务需求，我将系统划分为以下微服务：

- 用户服务 (User Service) ：负责用户注册、登录、认证和用户信息管理
- 宠物服务 (Pet Service) ：管理宠物信息，包括品种、年龄等基本信息
- 寄养服务 (Boarding Service) ：管理服务提供者发布的寄养服务信息
- 订单服务 (Order Service) ：处理预订、支付和订单状态管理
- 评价服务 (Review Service) ：管理用户对寄养服务的评价和反馈
- 通知服务 (Notification Service) ：处理系统消息、订单变动提醒等
- 管理服务 (Admin Service) ：管理员审核用户和服务提供者信息





### 2. 系统架构图

```plaintext
                  ┌─────────────────┐
                  │   API Gateway    │
                  │     (Nginx)      │
                  └─────────────────┘
                          │
          ┌───────────────┼───────────────┐
          │               │               │
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│   User Service  │ │   Pet Service   │ │ Boarding Service│
└─────────────────┘ └─────────────────┘ └─────────────────┘
          │               │               │
          └───────────────┼───────────────┘
                          │
          ┌───────────────┼───────────────┐
          │               │               │
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│  Order Service  │ │  Review Service │ │Notification Svc │
└─────────────────┘ └─────────────────┘ └─────────────────┘
          │               │               │
          └───────────────┼───────────────┘
                          │
                  ┌─────────────────┐
                  │  Admin Service   │
                  └─────────────────┘
                          │
          ┌───────────────┼───────────────┐
          │               │               │
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│  MySQL Cluster  │ │  Redis Cluster  │ │  Kafka Cluster  │
└─────────────────┘ └─────────────────┘ └─────────────────┘
```



### 3. 微服务划分原由
- 按业务功能划分 ：每个微服务对应系统中的一个核心业务功能，保持单一职责原则
- 数据自治 ：每个微服务管理自己的数据，减少服务间强依赖
- 团队协作 ：不同团队可以独立开发和维护不同的微服务
- 技术异构 ：允许不同微服务使用最适合其业务需求的技术栈
- 独立部署 ：每个微服务可以独立部署和扩展，提高系统弹性
## 二、技术选型
### 1. 开发语言和框架
- 编程语言 ：Go语言（基于实验中已使用的技术栈）
- Web框架 ：Gin（用于创建RESTful API端点）
- ORM框架 ：GORM（用于数据库操作）
### 2. 微服务通信
- 同步通信 ：RESTful API（基于HTTP协议的服务间通信）
- 异步通信 ：Kafka消息队列（用于解耦服务、处理异步任务）
- 服务发现 ：通过API网关（Nginx）进行请求路由
### 3. 数据存储
- 关系型数据库 ：MySQL（用于存储结构化数据）
- 缓存 ：Redis（用于缓存热点数据、实现分布式锁）
- 消息队列 ：Kafka（用于异步消息处理）
### 4. 部署和运维
- 容器化 ：Docker（用于应用打包和部署）
- 负载均衡 ：Nginx（实现服务的高可用和负载均衡）
- 服务编排 ：Docker Compose（管理多容器应用）
## 三、数据库设计
### 1. 数据库分表策略
根据业务需求，我们采用以下分表策略：

- 水平分表 ：对用户表和订单表进行水平分表
  
  - 用户表按用户ID取模分表（user_1, user_2, ...）
  - 订单表按时间范围分表（orders_2023, orders_2024, ...）
- 垂直分表 ：将不常用的大字段拆分到扩展表
  
  - 用户详细信息表（user_details）
  - 服务详细描述表（service_details）

### 2. 主要数据表设计

```sql
-- 用户表
CREATE TABLE `user_%d` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` varchar(64) NOT NULL COMMENT '用户唯一ID',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(128) NOT NULL COMMENT '密码（加密存储）',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
  `role` tinyint NOT NULL DEFAULT '0' COMMENT '角色：0-普通用户，1-服务提供者，2-管理员',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-待审核，1-正常，2-禁用',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  KEY `idx_phone` (`phone`),
  KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 宠物信息表
CREATE TABLE `pet` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `pet_id` varchar(64) NOT NULL COMMENT '宠物唯一ID',
  `user_id` varchar(64) NOT NULL COMMENT '所属用户ID',
  `name` varchar(64) NOT NULL COMMENT '宠物名称',
  `type` varchar(32) NOT NULL COMMENT '宠物类型（猫/狗等）',
  `breed` varchar(64) DEFAULT NULL COMMENT '品种',
  `age` int DEFAULT NULL COMMENT '年龄',
  `gender` tinyint DEFAULT NULL COMMENT '性别：0-未知，1-雄性，2-雌性',
  `weight` decimal(5,2) DEFAULT NULL COMMENT '体重(kg)',
  `description` text COMMENT '描述',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_pet_id` (`pet_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 寄养服务表
CREATE TABLE `boarding_service` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `service_id` varchar(64) NOT NULL COMMENT '服务唯一ID',
  `provider_id` varchar(64) NOT NULL COMMENT '服务提供者ID',
  `title` varchar(128) NOT NULL COMMENT '服务标题',
  `price` decimal(10,2) NOT NULL COMMENT '价格/天',
  `location` varchar(255) NOT NULL COMMENT '服务地点',
  `pet_type` varchar(32) NOT NULL COMMENT '适用宠物类型',
  `capacity` int NOT NULL DEFAULT '1' COMMENT '可接待数量',
  `available` tinyint NOT NULL DEFAULT '1' COMMENT '是否可用：0-不可用，1-可用',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_service_id` (`service_id`),
  KEY `idx_provider_id` (`provider_id`),
  KEY `idx_pet_type` (`pet_type`),
  KEY `idx_location` (`location`(20))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 订单表
CREATE TABLE `order_%d` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_id` varchar(64) NOT NULL COMMENT '订单唯一ID',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `pet_id` varchar(64) NOT NULL COMMENT '宠物ID',
  `service_id` varchar(64) NOT NULL COMMENT '服务ID',
  `provider_id` varchar(64) NOT NULL COMMENT '服务提供者ID',
  `start_date` date NOT NULL COMMENT '开始日期',
  `end_date` date NOT NULL COMMENT '结束日期',
  `total_amount` decimal(10,2) NOT NULL COMMENT '总金额',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-待支付，1-已支付，2-进行中，3-已完成，4-已取消',
  `payment_id` varchar(64) DEFAULT NULL COMMENT '支付ID',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_service_id` (`service_id`),
  KEY `idx_provider_id` (`provider_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 评价表
CREATE TABLE `review` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `review_id` varchar(64) NOT NULL COMMENT '评价唯一ID',
  `order_id` varchar(64) NOT NULL COMMENT '订单ID',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `provider_id` varchar(64) NOT NULL COMMENT '服务提供者ID',
  `service_id` varchar(64) NOT NULL COMMENT '服务ID',
  `rating` tinyint NOT NULL COMMENT '评分(1-5)',
  `content` text COMMENT '评价内容',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_review_id` (`review_id`),
  UNIQUE KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_provider_id` (`provider_id`),
  KEY `idx_service_id` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```



## 四、核心功能实现
### 1. API网关实现
使用Nginx作为API网关，实现请求路由和负载均衡：

```nginx
# nginx.conf
http {
    # 上游服务器配置
    upstream user_service {
        server user-service:8081;
        server user-service-replica:8081 backup; # 备份服务器
    }
    
    upstream pet_service {
        server pet-service:8082;
        server pet-service-replica:8082 backup;
    }
    
    upstream boarding_service {
        server boarding-service:8083;
        server boarding-service-replica:8083 backup;
    }
    
    upstream order_service {
        server order-service:8084;
        server order-service-replica:8084 backup;
    }
    
    upstream review_service {
        server review-service:8085;
        server review-service-replica:8085 backup;
    }
    
    upstream notification_service {
        server notification-service:8086;
    }
    
    upstream admin_service {
        server admin-service:8087;
    }
    
    # API网关配置
    server {
        listen 80;
        server_name api.petboarding.com;
        
        # 用户服务
        location /api/users/ {
            proxy_pass http://user_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
        
        # 宠物服务
        location /api/pets/ {
            proxy_pass http://pet_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
        
        # 寄养服务
        location /api/boarding/ {
            proxy_pass http://boarding_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
        
        # 订单服务
        location /api/orders/ {
            proxy_pass http://order_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
        
        # 评价服务
        location /api/reviews/ {
            proxy_pass http://review_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
        
        # 通知服务
        location /api/notifications/ {
            proxy_pass http://notification_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
        
        # 管理服务
        location /api/admin/ {
            proxy_pass http://admin_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
        }
    }
}
```



### 2. 分布式锁实现（基于Redis）
使用Redis实现分布式锁，用于处理并发预订请求：

```go
// boarding/service/booking.go
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type BookingService struct {
	RedisClient *redis.Client
}

// 尝试预订服务，使用分布式锁确保并发安全
func (s *BookingService) TryBookService(ctx context.Context, serviceID string, startDate, endDate time.Time) (bool, error) {
	// 生成锁的key
	lockKey := fmt.Sprintf("booking_lock:%s:%s_%s", serviceID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	
	// 尝试获取锁，设置过期时间为10秒
	lockValue := time.Now().String()
	lockAcquired, err := s.RedisClient.SetNX(ctx, lockKey, lockValue, 10*time.Second).Result()
	
	if err != nil {
		return false, fmt.Errorf("获取锁失败: %w", err)
	}
	
	if !lockAcquired {
		return false, fmt.Errorf("服务正在被其他用户预订，请稍后再试")
	}
	
	// 确保锁会被释放
	defer s.RedisClient.Del(ctx, lockKey)
	
	// 检查服务在指定日期是否可用
	isAvailable, err := s.checkServiceAvailability(ctx, serviceID, startDate, endDate)
	if err != nil {
		return false, err
	}
	
	if !isAvailable {
		return false, fmt.Errorf("服务在所选日期不可用")
	}
	
	// 预订逻辑...
	
	return true, nil
}

// 检查服务在指定日期是否可用
func (s *BookingService) checkServiceAvailability(ctx context.Context, serviceID string, startDate, endDate time.Time) (bool, error) {
	// 实现检查逻辑...
	return true, nil
}
```



### 3. 分布式事务实现（基于Kafka消息队列）
使用Kafka实现分布式事务，确保订单创建和支付过程的数据一致性：

```go
// order/service/order.go
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/petboarding/order/model"
	"gorm.io/gorm"
)

type OrderService struct {
	DB           *gorm.DB
	KafkaProducer sarama.SyncProducer
}

// 创建订单并发送消息到Kafka
func (s *OrderService) CreateOrder(ctx context.Context, order *model.Order) error {
	// 开始数据库事务
	tx := s.DB.Begin()
	
	// 创建订单记录
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("创建订单失败: %w", err)
	}
	
	// 准备Kafka消息
	orderJSON, err := json.Marshal(order)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("序列化订单失败: %w", err)
	}
	
	// 发送消息到Kafka
	msg := &sarama.ProducerMessage{
		Topic: "order_created",
		Key:   sarama.StringEncoder(order.OrderID),
		Value: sarama.ByteEncoder(orderJSON),
	}
	
	_, _, err = s.KafkaProducer.SendMessage(msg)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("发送Kafka消息失败: %w", err)
	}
	
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}
	
	return nil
}
```



### 4. 缓存实现（基于Redis）
使用Redis缓存热点数据，提升系统性能：

```go
// boarding/service/service.go
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/petboarding/boarding/model"
	"gorm.io/gorm"
)

type BoardingService struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

// 获取寄养服务详情，优先从缓存获取
func (s *BoardingService) GetServiceByID(ctx context.Context, serviceID string) (*model.BoardingService, error) {
	// 尝试从Redis缓存获取
	cacheKey := fmt.Sprintf("service:%s", serviceID)
	cachedData, err := s.RedisClient.Get(ctx, cacheKey).Result()
	
	// 缓存命中
	if err == nil {
		var service model.BoardingService
		if err := json.Unmarshal([]byte(cachedData), &service); err == nil {
			return &service, nil
		}
	}
	
	// 缓存未命中或解析失败，从数据库获取
	var service model.BoardingService
	if err := s.DB.Where("service_id = ?", serviceID).First(&service).Error; err != nil {
		return nil, fmt.Errorf("获取服务失败: %w", err)
	}
	
	// 将数据写入缓存
	serviceJSON, _ := json.Marshal(service)
	s.RedisClient.Set(ctx, cacheKey, serviceJSON, 30*time.Minute) // 缓存30分钟
	
	return &service, nil
}

// 更新服务信息并清除缓存
func (s *BoardingService) UpdateService(ctx context.Context, service *model.BoardingService) error {
	// 更新数据库
	if err := s.DB.Save(service).Error; err != nil {
		return fmt.Errorf("更新服务失败: %w", err)
	}
	
	// 删除缓存
	cacheKey := fmt.Sprintf("service:%s", service.ServiceID)
	s.RedisClient.Del(ctx, cacheKey)
	
	return nil
}
```



## 五、部署方案

### 1. Docker容器化
为每个微服务创建Dockerfile：

```dockerfile
# 以用户服务为例
FROM golang:1.18-alpine AS builder

WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service ./cmd/user

# 运行阶段使用轻量级镜像
FROM alpine:latest

WORKDIR /app

# 从builder阶段复制编译好的二进制文件
COPY --from=builder /app/user-service .

# 设置时区
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

# 暴露端口
EXPOSE 8081

# 运行服务
CMD ["./user-service"]
```



### 2. Docker Compose配置
使用Docker Compose管理多容器应用：

```yaml
# docker-compose.yml
version: '3'

services:
  # API网关
  nginx:
    image: nginx:latest
    ports:
      - "80:80"
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/nginx.conf
    depends_on:
      - user-service
      - pet-service
      - boarding-service
      - order-service
      - review-service
      - notification-service
      - admin-service
    networks:
      - petboarding-network

  # MySQL数据库
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: petboarding
      MYSQL_USER: petuser
      MYSQL_PASSWORD: petpassword
    volumes:
      - mysql-data:/var/lib/mysql
      - ./mysql/init:/docker-entrypoint-initdb.d
    ports:
      - "3306:3306"
    networks:
      - petboarding-network

  # Redis缓存
  redis:
    image: redis:6.2
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data
    networks:
      - petboarding-network

  # Kafka消息队列
  zookeeper:
    image: wurstmeister/zookeeper
    ports:
      - "2181:2181"
    networks:
      - petboarding-network

  kafka:
    image: wurstmeister/kafka
    ports:
      - "9092:9092"
    environment:
      KAFKA_ADVERTISED_HOST_NAME: kafka
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_CREATE_TOPICS: "order_created:1:1,payment_completed:1:1,notification:1:1"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    depends_on:
      - zookeeper
    networks:
      - petboarding-network

  # 用户服务
  user-service:
    build:
      context: ./user-service
      dockerfile: Dockerfile
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: petuser
      DB_PASSWORD: petpassword
      DB_NAME: petboarding
      REDIS_HOST: redis
      REDIS_PORT: 6379
    depends_on:
      - mysql
      - redis
    networks:
      - petboarding-network

  # 宠物服务
  pet-service:
    build:
      context: ./pet-service
      dockerfile: Dockerfile
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: petuser
      DB_PASSWORD: petpassword
      DB_NAME: petboarding
      REDIS_HOST: redis
      REDIS_PORT: 6379
    depends_on:
      - mysql
      - redis
    networks:
      - petboarding-network

  # 寄养服务
  boarding-service:
    build:
      context: ./boarding-service
      dockerfile: Dockerfile
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: petuser
      DB_PASSWORD: petpassword
      DB_NAME: petboarding
      REDIS_HOST: redis
      REDIS_PORT: 6379
      KAFKA_BROKERS: kafka:9092
    depends_on:
      - mysql
      - redis
      - kafka
    networks:
      - petboarding-network

  # 订单服务
  order-service:
    build:
      context: ./order-service
      dockerfile: Dockerfile
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: petuser
      DB_PASSWORD: petpassword
      DB_NAME: petboarding
      REDIS_HOST: redis
      REDIS_PORT: 6379
      KAFKA_BROKERS: kafka:9092
    depends_on:
      - mysql
      - redis
      - kafka
    networks:
      - petboarding-network

  # 评价服务
  review-service:
    build:
      context: ./review-service
      dockerfile: Dockerfile
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: petuser
      DB_PASSWORD: petpassword
      DB_NAME: petboarding
      REDIS_HOST: redis
      REDIS_PORT: 6379
    depends_on:
      - mysql
      - redis
    networks:
      - petboarding-network

  # 通知服务
  notification-service:
    build:
      context: ./notification-service
      dockerfile: Dockerfile
    environment:
      KAFKA_BROKERS: kafka:9092
      REDIS_HOST: redis
      REDIS_PORT: 6379
    depends_on:
      - kafka
      - redis
    networks:
      - petboarding-network

  # 管理服务
  admin-service:
    build:
      context: ./admin-service
      dockerfile: Dockerfile
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: petuser
      DB_PASSWORD: petpassword
      DB_NAME: petboarding
      REDIS_HOST: redis
      REDIS_PORT: 6379
    depends_on:
      - mysql
      - redis
    networks:
      - petboarding-network

networks:
  petboarding-network:
    driver: bridge

volumes:
  mysql-data:
  redis-data:
```



## 六、微服务示例代码

### 1. 用户服务（User Service）

```go
// user-service/handler/user.go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petboarding/user/model"
	"github.com/petboarding/user/service"
)

type UserHandler struct {
	UserService *service.UserService
}

// 注册新用户
func (h *UserHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.UserService.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "用户注册成功", "user_id": user.UserID})
}

// 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	token, err := h.UserService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// 获取用户信息
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	
	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// 更新用户信息
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	userID := c.Param("id")
	
	var updateReq model.UserUpdate
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.UserService.UpdateUser(userID, &updateReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}
```



### 2.订单服务（Order Service）

```go
// order-service/handler/order.go
package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/petboarding/order/model"
	"github.com/petboarding/order/service"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

// 创建订单
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderReq struct {
		UserID     string    `json:"user_id" binding:"required"`
		PetID      string    `json:"pet_id" binding:"required"`
		ServiceID  string    `json:"service_id" binding:"required"`
		ProviderID string    `json:"provider_id" binding:"required"`
		StartDate  time.Time `json:"start_date" binding:"required"`
		EndDate    time.Time `json:"end_date" binding:"required"`
		Amount     float64   `json:"amount" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// 创建订单对象
	order := &model.Order{
		OrderID:    generateOrderID(),
		UserID:     orderReq.UserID,
		PetID:      orderReq.PetID,
		ServiceID:  orderReq.ServiceID,
		ProviderID: orderReq.ProviderID,
		StartDate:  orderReq.StartDate,
		EndDate:    orderReq.EndDate,
		TotalAmount: orderReq.Amount,
		Status:     0, // 待支付
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	
	// 保存订单并发送消息到Kafka
	if err := h.OrderService.CreateOrder(c, order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message":  "订单创建成功",
		"order_id": order.OrderID,
	})
}

// 支付订单
func (h *OrderHandler) PayOrder(c *gin.Context) {
	orderID := c.Param("id")
	
	var payReq struct {
		PaymentID string `json:"payment_id" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&payReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.OrderService.PayOrder(c, orderID, payReq.PaymentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "订单支付成功"})
}

// 获取订单详情
func (h *OrderHandler) GetOrderDetail(c *gin.Context) {
	orderID := c.Param("id")
	
	order, err := h.OrderService.GetOrderByID(c, orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "订单不存在"})
		return
	}
	
	c.JSON(http.StatusOK, order)
}

// 获取用户的订单列表
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID := c.Param("user_id")
	status := c.Query("status")
	
	orders, err := h.OrderService.GetOrdersByUserID(c, userID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, orders)
}

// 生成订单ID
func generateOrderID() string {
	// 实现订单ID生成逻辑
	return "ORD" + time.Now().Format("20060102150405") + randomString(6)
}

// 生成随机字符串
func randomString(n int) string {
	// 实现随机字符串生成逻辑
	return "123456" // 简化示例
}
```



## 七、总结
本宠物寄养系统后端开发方案**基于微服务架构设计**，充分考虑了系统的**可扩展性**、**高可用性**和**性能需求**。通过合理的**微服务划分**、**数据库分表设计**、**分布式锁**和**分布式事务**的实现，以及**缓存中间件**的使用，满足了项目的**所有考核重点要求**：

1. **分布式微服务架构** ：清晰划分了**7个核心微服务**，并提供了详细的**架构图**和**划分原由**
2. **微服务通信** ：实现了**RESTful API**和**Kafka消息队列**两种**通信方式**，并**使用Nginx作为API网关进行请求路由**
3. **分布式事务与分布式锁** ：使用**Redis实现分布式锁**，**保证并发安全**；使用**Kafka消息队列实现分布式事务**，**确保数据一致性**
4. **集群部署** ：通过**Docker容器化**和**Nginx负载均衡**实现服务的**高可用部署**
5. **数据库分表及中间件使用** ：实现了**水平分表**和**垂直分表**，并**结合Redis缓存提升系统性能**

该方案不仅满足了宠物寄养系统的**功能需求**，还充分**利用了分布式系统的优势**，为系统提供了**良好的性能**、**可靠性**和**可扩展性**。
