# 宠物寄养系统数据库设计

## 概述

本文档详细说明宠物寄养系统的数据库设计，包括表结构、字段定义、索引设计和分表策略。系统采用MySQL作为关系型数据库，结合Redis缓存提升性能。

## 数据库分表策略

根据业务需求和数据量预估，我们采用以下分表策略：

### 水平分表

- **用户表**：按用户ID取模分表（user_0, user_1, ...）
- **订单表**：按时间范围分表（orders_2023, orders_2024, ...）
- **通知表**：按时间范围分表（notifications_2023, notifications_2024, ...）

### 垂直分表

- **用户基本信息表**和**用户详细信息表**：将不常用的大字段拆分到详细信息表
- **寄养服务基本信息表**和**寄养服务详细描述表**：将详细描述等大字段拆分到详细描述表

## 表结构设计

### 用户相关表

#### 用户基本信息表 (user_%d)

```sql
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
```

#### 用户详细信息表 (user_detail)

```sql
CREATE TABLE `user_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` varchar(64) NOT NULL COMMENT '用户唯一ID',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像URL',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  `bio` text COMMENT '个人简介',
  `preferences` json DEFAULT NULL COMMENT '偏好设置',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### 宠物相关表

#### 宠物信息表 (pet)

```sql
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
```

### 寄养服务相关表

#### 寄养服务基本信息表 (boarding_service)

```sql
CREATE TABLE `boarding_service` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `service_id` varchar(64) NOT NULL COMMENT '服务唯一ID',
  `provider_id` varchar(64) NOT NULL COMMENT '服务提供者ID',
  `title` varchar(128) NOT NULL COMMENT '服务标题',
  `price` decimal(10,2) NOT NULL COMMENT '价格/天',
  `location` varchar(255) NOT NULL COMMENT '服务地点',
  `pet_type` varchar(32) NOT NULL COMMENT '适用宠物类型',
  `capacity` int NOT NULL COMMENT '容纳数量',
  `available` tinyint NOT NULL DEFAULT '1' COMMENT '是否可用：0-不可用，1-可用',
  `rating` decimal(2,1) DEFAULT NULL COMMENT '平均评分',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_service_id` (`service_id`),
  KEY `idx_provider_id` (`provider_id`),
  KEY `idx_pet_type` (`pet_type`),
  KEY `idx_location` (`location`(20))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

#### 寄养服务详细描述表 (boarding_service_detail)

```sql
CREATE TABLE `boarding_service_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `service_id` varchar(64) NOT NULL COMMENT '服务唯一ID',
  `description` text COMMENT '详细描述',
  `facilities` json DEFAULT NULL COMMENT '设施列表',
  `services` json DEFAULT NULL COMMENT '提供的服务列表',
  `rules` text COMMENT '规则和要求',
  `photos` json DEFAULT NULL COMMENT '照片URL列表',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_service_id` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### 订单相关表

#### 订单表 (orders_%d)

```sql
CREATE TABLE `orders_%d` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_id` varchar(64) NOT NULL COMMENT '订单唯一ID',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `service_id` varchar(64) NOT NULL COMMENT '服务ID',
  `pet_id` varchar(64) NOT NULL COMMENT '宠物ID',
  `start_date` date NOT NULL COMMENT '开始日期',
  `end_date` date NOT NULL COMMENT '结束日期',
  `total_price` decimal(10,2) NOT NULL COMMENT '总价',
  `status` tinyint NOT NULL COMMENT '状态：0-待确认，1-已确认，2-进行中，3-已完成，4-已取消',
  `payment_status` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态：0-未支付，1-已支付，2-已退款',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_service_id` (`service_id`),
  KEY `idx_status` (`status`),
  KEY `idx_date_range` (`start_date`, `end_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

#### 订单详情表 (order_detail)

```sql
CREATE TABLE `order_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_id` varchar(64) NOT NULL COMMENT '订单唯一ID',
  `special_requests` text COMMENT '特殊要求',
  `care_notes` text COMMENT '护理记录',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### 评价相关表

#### 评价表 (review)

```sql
CREATE TABLE `review` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `review_id` varchar(64) NOT NULL COMMENT '评价唯一ID',
  `order_id` varchar(64) NOT NULL COMMENT '订单ID',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `service_id` varchar(64) NOT NULL COMMENT '服务ID',
  `rating` tinyint NOT NULL COMMENT '评分（1-5）',
  `content` text COMMENT '评价内容',
  `photos` json DEFAULT NULL COMMENT '照片URL列表',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_review_id` (`review_id`),
  UNIQUE KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_service_id` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### 通知相关表

#### 通知表 (notifications_%d)

```sql
CREATE TABLE `notifications_%d` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `notification_id` varchar(64) NOT NULL COMMENT '通知唯一ID',
  `user_id` varchar(64) NOT NULL COMMENT '接收用户ID',
  `type` tinyint NOT NULL COMMENT '类型：0-系统通知，1-订单通知，2-评价通知',
  `title` varchar(128) NOT NULL COMMENT '通知标题',
  `content` text NOT NULL COMMENT '通知内容',
  `related_id` varchar(64) DEFAULT NULL COMMENT '相关ID（订单ID等）',
  `read` tinyint NOT NULL DEFAULT '0' COMMENT '是否已读：0-未读，1-已读',
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_notification_id` (`notification_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_type` (`type`),
  KEY `idx_read` (`read`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### 管理员相关表

#### 管理员表 (admin)

```sql
CREATE TABLE `admin` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `admin_id` varchar(64) NOT NULL COMMENT '管理员唯一ID',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(128) NOT NULL COMMENT '密码（加密存储）',
  `email` varchar(64) NOT NULL COMMENT '邮箱',
  `role` tinyint NOT NULL DEFAULT '0' COMMENT '角色：0-普通管理员，1-超级管理员',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0-禁用，1-正常',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_admin_id` (`admin_id`),
  UNIQUE KEY `idx_username` (`username`),
  UNIQUE KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

## 数据库索引设计

为了提高查询效率，我们为各个表设计了以下索引：

1. **主键索引**：每个表都有自增的id作为主键
2. **唯一索引**：对于业务ID（如user_id, pet_id, service_id等）创建唯一索引
3. **普通索引**：对于常用的查询条件创建普通索引，如user_id, status, type等
4. **复合索引**：对于常用的多条件查询创建复合索引，如订单表的日期范围索引

## 缓存设计

使用Redis缓存热点数据，提高系统性能：

1. **用户信息缓存**：缓存用户基本信息，减少数据库查询
   - 键格式：`user:{user_id}`
   - 过期时间：1小时

2. **宠物信息缓存**：缓存宠物信息
   - 键格式：`pet:{pet_id}`
   - 过期时间：1小时

3. **寄养服务缓存**：缓存热门寄养服务信息
   - 键格式：`service:{service_id}`
   - 过期时间：30分钟

4. **订单状态缓存**：缓存订单状态
   - 键格式：`order:{order_id}:status`
   - 过期时间：10分钟

5. **通知计数缓存**：缓存用户未读通知数量
   - 键格式：`notification:unread:{user_id}`
   - 过期时间：5分钟

## 分布式锁

使用Redis实现分布式锁，解决并发问题：

1. **订单创建锁**：防止重复下单
   - 键格式：`lock:order:create:{user_id}:{service_id}:{start_date}:{end_date}`
   - 过期时间：10秒

2. **库存锁**：防止超卖
   - 键格式：`lock:service:capacity:{service_id}:{date}`
   - 过期时间：5秒

3. **用户注册锁**：防止用户名或邮箱重复注册
   - 键格式：`lock:user:register:{username/email}`
   - 过期时间：5秒

## 数据库迁移计划

1. **初始化脚本**：创建数据库表结构的SQL脚本
2. **数据迁移工具**：使用Flyway或Liquibase管理数据库版本和迁移
3. **分表实现**：使用代码中的分表策略，根据用户ID或时间范围动态选择表

## 性能优化策略

1. **读写分离**：主库写入，从库读取，减轻主库压力
2. **分表策略**：根据业务特点进行水平分表和垂直分表
3. **索引优化**：为常用查询创建合适的索引
4. **缓存机制**：使用Redis缓存热点数据
5. **批量操作**：使用批量插入和更新，减少数据库交互次数
6. **定期归档**：将历史数据归档到历史表，提高活跃表的查询效率