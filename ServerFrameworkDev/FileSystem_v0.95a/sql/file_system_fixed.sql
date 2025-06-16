-- 1. 先删除旧数据库（如果需要）
DROP DATABASE IF EXISTS `file_system`;

-- 2. 创建新数据库
CREATE DATABASE IF NOT EXISTS `file_system` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `file_system`;

-- 3. 创建基础表（按依赖顺序）
-- 修改users表的id字段类型为bigint unsigned
CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL COMMENT '用户名',
    `password` varchar(100) NOT NULL COMMENT '密码',
    `nick_name` varchar(50) DEFAULT '系统用户' COMMENT '昵称',
    `avatar` varchar(500) DEFAULT '/images/avatar/default.png' COMMENT '头像，增加长度以适应可能较长的路径',
    `status` tinyint(1) DEFAULT 1 COMMENT '状态(1:正常,0:禁用)',
    `email` varchar(100) NOT NULL COMMENT '用户邮箱',
    `phone` varchar(20) NOT NULL COMMENT '用户手机号',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`),
    UNIQUE KEY `idx_email` (`email`) COMMENT '添加邮箱唯一索引，保证邮箱唯一性'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

-- 角色表
CREATE TABLE IF NOT EXISTS `roles` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL COMMENT '角色名称',
    `description` varchar(255) DEFAULT NULL COMMENT '角色描述',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 权限表
CREATE TABLE IF NOT EXISTS `permissions` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL COMMENT '权限名称',
    `description` varchar(255) DEFAULT NULL COMMENT '权限描述',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 用户角色关联表
CREATE TABLE IF NOT EXISTS `user_roles` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL COMMENT '用户 ID',
    `role_id` bigint unsigned NOT NULL COMMENT '角色 ID',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_role_id` (`role_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户角色关联表';

-- 角色权限关联表
CREATE TABLE IF NOT EXISTS `role_permissions` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `role_id` bigint unsigned NOT NULL,
    `permission_id` bigint unsigned NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_role_permission` (`role_id`, `permission_id`),
    CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 文件表（修正外键引用）
-- 修改外键约束的表创建顺序
CREATE TABLE IF NOT EXISTS `files` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL COMMENT '文件名',
    `path` varchar(500) NOT NULL COMMENT '存储路径',
    `size` bigint NOT NULL COMMENT '文件大小(字节)',
    `type` varchar(50) DEFAULT NULL COMMENT '文件类型',
    `uploader_id` bigint unsigned NOT NULL COMMENT '上传者ID',
    `is_public` tinyint(1) DEFAULT 0 COMMENT '是否公开(1:是,0:否)',
    `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_uploader` (`uploader_id`),
    CONSTRAINT `fk_files_user` FOREIGN KEY (`uploader_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 文件权限表（修正外键引用）
CREATE TABLE IF NOT EXISTS `file_permissions` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `file_id` bigint unsigned NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    `permission` varchar(50) NOT NULL COMMENT '权限类型(download,edit,delete)',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_user_permission` (
        `file_id`,
        `user_id`,
        `permission`
    ),
    CONSTRAINT `fk_file_permissions_file` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_file_permissions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 操作记录表（修正外键引用）
CREATE TABLE IF NOT EXISTS `operation_records` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL,
    `operation_type` varchar(50) NOT NULL COMMENT '操作类型',
    `operation_detail` text COMMENT '操作详情',
    `ip_address` varchar(50) DEFAULT NULL COMMENT 'IP地址',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`),
    CONSTRAINT `fk_operation_records_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 5. 初始化数据（添加存在性检查）
INSERT IGNORE INTO
    `roles` (`name`, `description`)
VALUES ('admin', '系统管理员'),
    ('user', '普通用户');

INSERT IGNORE INTO
    `permissions` (`name`, `description`)
VALUES ('file_upload', '文件上传权限'),
    ('file_download', '文件下载权限'),
    ('file_delete', '文件删除权限'),
    ('file_share', '文件共享权限'),
    ('user_manage', '用户管理权限'),
    ('role_manage', '角色管理权限');

-- 管理员拥有所有权限
INSERT INTO
    `role_permissions` (`role_id`, `permission_id`)
SELECT r.id, p.id
FROM `roles` r, `permissions` p
WHERE
    r.name = 'admin';

-- 普通用户基础权限
INSERT INTO
    `role_permissions` (`role_id`, `permission_id`)
SELECT r.id, p.id
FROM `roles` r, `permissions` p
WHERE
    r.name = 'user'
    AND p.name IN (
        'file_upload',
        'file_download'
    );
