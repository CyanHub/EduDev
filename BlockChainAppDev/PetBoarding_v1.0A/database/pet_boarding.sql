-- 创建数据库
CREATE DATABASE pet_boarding;

-- 使用数据库

USE pet_boarding;

-- 用户表
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 宠物基础信息表
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

-- 宠物健康档案表
CREATE TABLE pet_health (
    id BIGINT PRIMARY KEY,
    medical_history TEXT,
    vaccination TEXT,
    last_checkup DATE,
    FOREIGN KEY (id) REFERENCES pet_basic (id),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 订单分表 (共8个)
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

-- 创建其余7个订单分表 (order_1 到 order_7)
CREATE TABLE order_1 LIKE order_0;

CREATE TABLE order_2 LIKE order_0;

CREATE TABLE order_3 LIKE order_0;

CREATE TABLE order_4 LIKE order_0;

CREATE TABLE order_5 LIKE order_0;

CREATE TABLE order_6 LIKE order_0;

CREATE TABLE order_7 LIKE order_0;

-- 服务提供者表
CREATE TABLE providers (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    service_type VARCHAR(50) NOT NULL,
    rating DECIMAL(3, 2),
    FOREIGN KEY (user_id) REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 创建索引
CREATE INDEX idx_pet_owner ON pet_basic (owner_id);

CREATE INDEX idx_order_user ON order_0 (user_id);

CREATE INDEX idx_order_pet ON order_0 (pet_id);