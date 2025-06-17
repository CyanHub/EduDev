-- 宠物寄养系统数据库初始化脚本
-- 创建数据库
CREATE DATABASE IF NOT EXISTS petboarding CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE petboarding;

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` VARCHAR(32) NOT NULL COMMENT '用户唯一标识',
  `username` VARCHAR(50) NOT NULL COMMENT '用户名',
  `password` VARCHAR(100) NOT NULL COMMENT '密码（加密存储）',
  `email` VARCHAR(100) NOT NULL COMMENT '邮箱',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',
  `role` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色：0-普通用户，1-服务提供者',
  `status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0-待审核，1-正常，2-禁用',
  `avatar` VARCHAR(255) DEFAULT NULL COMMENT '头像URL',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  UNIQUE KEY `idx_username` (`username`),
  UNIQUE KEY `idx_email` (`email`),
  KEY `idx_status` (`status`),
  KEY `idx_role` (`role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 宠物表
CREATE TABLE IF NOT EXISTS `pets` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `pet_id` VARCHAR(32) NOT NULL COMMENT '宠物唯一标识',
  `user_id` VARCHAR(32) NOT NULL COMMENT '所属用户ID',
  `name` VARCHAR(50) NOT NULL COMMENT '宠物名称',
  `type` VARCHAR(30) NOT NULL COMMENT '宠物类型',
  `breed` VARCHAR(50) DEFAULT NULL COMMENT '品种',
  `age` INT UNSIGNED DEFAULT NULL COMMENT '年龄',
  `gender` TINYINT UNSIGNED DEFAULT 0 COMMENT '性别：0-未知，1-雄性，2-雌性',
  `weight` DECIMAL(5,2) DEFAULT NULL COMMENT '体重(kg)',
  `description` TEXT DEFAULT NULL COMMENT '描述',
  `photo` VARCHAR(255) DEFAULT NULL COMMENT '照片URL',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_pet_id` (`pet_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='宠物表';

-- 寄养服务表
CREATE TABLE IF NOT EXISTS `boarding_services` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `service_id` VARCHAR(32) NOT NULL COMMENT '服务唯一标识',
  `provider_id` VARCHAR(32) NOT NULL COMMENT '服务提供者ID',
  `title` VARCHAR(100) NOT NULL COMMENT '服务标题',
  `price` DECIMAL(10,2) NOT NULL COMMENT '价格/天',
  `location` VARCHAR(255) NOT NULL COMMENT '服务地点',
  `pet_type` VARCHAR(30) NOT NULL COMMENT '适用宠物类型',
  `capacity` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT '容纳数量',
  `available` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否可用：0-不可用，1-可用',
  `rating` DECIMAL(2,1) DEFAULT NULL COMMENT '评分',
  `review_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '评价数量',
  `description` TEXT DEFAULT NULL COMMENT '详细描述',
  `facilities` JSON DEFAULT NULL COMMENT '设施列表',
  `services` JSON DEFAULT NULL COMMENT '提供的服务列表',
  `rules` TEXT DEFAULT NULL COMMENT '规则和要求',
  `photos` JSON DEFAULT NULL COMMENT '照片URL列表',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_service_id` (`service_id`),
  KEY `idx_provider_id` (`provider_id`),
  KEY `idx_pet_type` (`pet_type`),
  KEY `idx_location` (`location`(20)),
  KEY `idx_price` (`price`),
  KEY `idx_rating` (`rating`),
  KEY `idx_available` (`available`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='寄养服务表';

-- 订单表
CREATE TABLE IF NOT EXISTS `orders` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `order_id` VARCHAR(32) NOT NULL COMMENT '订单唯一标识',
  `user_id` VARCHAR(32) NOT NULL COMMENT '用户ID',
  `service_id` VARCHAR(32) NOT NULL COMMENT '服务ID',
  `pet_id` VARCHAR(32) NOT NULL COMMENT '宠物ID',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `total_price` DECIMAL(10,2) NOT NULL COMMENT '总价',
  `status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0-待确认，1-已确认，2-进行中，3-已完成，4-已取消',
  `payment_status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付状态：0-未支付，1-已支付，2-已退款',
  `special_requests` TEXT DEFAULT NULL COMMENT '特殊要求',
  `care_notes` TEXT DEFAULT NULL COMMENT '护理记录',
  `cancel_reason` VARCHAR(255) DEFAULT NULL COMMENT '取消原因',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_service_id` (`service_id`),
  KEY `idx_pet_id` (`pet_id`),
  KEY `idx_status` (`status`),
  KEY `idx_payment_status` (`payment_status`),
  KEY `idx_date_range` (`start_date`, `end_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

-- 评价表
CREATE TABLE IF NOT EXISTS `reviews` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `review_id` VARCHAR(32) NOT NULL COMMENT '评价唯一标识',
  `order_id` VARCHAR(32) NOT NULL COMMENT '订单ID',
  `user_id` VARCHAR(32) NOT NULL COMMENT '用户ID',
  `service_id` VARCHAR(32) NOT NULL COMMENT '服务ID',
  `rating` TINYINT UNSIGNED NOT NULL COMMENT '评分（1-5）',
  `content` TEXT DEFAULT NULL COMMENT '评价内容',
  `photos` JSON DEFAULT NULL COMMENT '照片URL列表',
  `reply` TEXT DEFAULT NULL COMMENT '商家回复',
  `reply_time` TIMESTAMP NULL DEFAULT NULL COMMENT '回复时间',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_review_id` (`review_id`),
  UNIQUE KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_service_id` (`service_id`),
  KEY `idx_rating` (`rating`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评价表';

-- 通知表
CREATE TABLE IF NOT EXISTS `notifications` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `notification_id` VARCHAR(32) NOT NULL COMMENT '通知唯一标识',
  `user_id` VARCHAR(32) NOT NULL COMMENT '用户ID',
  `type` TINYINT UNSIGNED NOT NULL COMMENT '类型：1-订单通知，2-系统通知，3-促销通知',
  `title` VARCHAR(100) NOT NULL COMMENT '标题',
  `content` TEXT NOT NULL COMMENT '内容',
  `related_id` VARCHAR(32) DEFAULT NULL COMMENT '相关ID（如订单ID）',
  `read` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否已读：0-未读，1-已读',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_notification_id` (`notification_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_type` (`type`),
  KEY `idx_read` (`read`),
  KEY `idx_related_id` (`related_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通知表';

-- 管理员表
CREATE TABLE IF NOT EXISTS `admins` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `admin_id` VARCHAR(32) NOT NULL COMMENT '管理员唯一标识',
  `username` VARCHAR(50) NOT NULL COMMENT '用户名',
  `password` VARCHAR(100) NOT NULL COMMENT '密码（加密存储）',
  `name` VARCHAR(50) DEFAULT NULL COMMENT '姓名',
  `role` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '角色：1-普通管理员，2-超级管理员',
  `status` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：0-禁用，1-正常',
  `last_login_time` TIMESTAMP NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_admin_id` (`admin_id`),
  UNIQUE KEY `idx_username` (`username`),
  KEY `idx_role` (`role`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- 服务可用性表（用于记录服务的可用日期）
CREATE TABLE IF NOT EXISTS `service_availability` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `service_id` VARCHAR(32) NOT NULL COMMENT '服务ID',
  `date` DATE NOT NULL COMMENT '日期',
  `available_count` INT UNSIGNED NOT NULL COMMENT '可用数量',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_service_date` (`service_id`, `date`),
  KEY `idx_date` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='服务可用性表';

-- 支付记录表
CREATE TABLE IF NOT EXISTS `payments` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `payment_id` VARCHAR(32) NOT NULL COMMENT '支付唯一标识',
  `order_id` VARCHAR(32) NOT NULL COMMENT '订单ID',
  `user_id` VARCHAR(32) NOT NULL COMMENT '用户ID',
  `amount` DECIMAL(10,2) NOT NULL COMMENT '金额',
  `payment_method` VARCHAR(20) NOT NULL COMMENT '支付方式',
  `transaction_id` VARCHAR(100) DEFAULT NULL COMMENT '交易ID',
  `status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0-处理中，1-成功，2-失败，3-退款',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_payment_id` (`payment_id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`),
  KEY `idx_transaction_id` (`transaction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付记录表';

-- 系统配置表
CREATE TABLE IF NOT EXISTS `system_configs` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `config_key` VARCHAR(50) NOT NULL COMMENT '配置键',
  `config_value` TEXT NOT NULL COMMENT '配置值',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '描述',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_config_key` (`config_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统配置表';

-- 初始化管理员账号
INSERT INTO `admins` (`admin_id`, `username`, `password`, `name`, `role`) VALUES
('a100000001', 'admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVKIUi', '系统管理员', 2);
-- 密码为 admin123，使用bcrypt加密

-- 初始化系统配置
INSERT INTO `system_configs` (`config_key`, `config_value`, `description`) VALUES
('site_name', '宠物寄养系统', '网站名称'),
('site_description', '提供专业的宠物寄养服务', '网站描述'),
('maintenance_mode', '0', '维护模式：0-关闭，1-开启'),
('version', '1.0.0', '系统版本'),
('max_order_days', '30', '最大预订天数'),
('commission_rate', '0.05', '平台佣金比例');