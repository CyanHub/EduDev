-- 角色(roles) id:					权限(permissions) id:
-- 1. 上传者;							1.上传文件权限
-- 2. 管理者;							2.下载文件权限
-- 3. 浏览者;

-- 先选定角色与权限ID， 将角色和权限关联起来。
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 1);

-- 再选定用户与角色ID， 将用户和角色关联起来。
INSERT INTO user_roles (user_id, role_id) VALUES (4, 1);

-- 当前数据库中“管理者”角色（id = 2）还未绑定权限，你可以使用以下 SQL 语句将“上传文件权限”（id = 1）和“下载文件权限”（id = 2）绑定到“管理者”角色：
INSERT INTO
    role_permissions (role_id, permission_id)
VALUES (2, 1),
    (2, 2);

-- 假设你想让 admin007 用户（id = 1）成为管理者，可使用以下 SQL 语句将该用户与“管理者”角色绑定：
INSERT INTO user_roles (user_id, role_id) VALUES (1, 2);

-- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --

-- 删除users表的deleted_at字段
ALTER TABLE users DROP COLUMN deleted_at;

-- 同时需要删除其他表的deleted_at（如果有）
ALTER TABLE roles DROP COLUMN deleted_at;

ALTER TABLE permissions DROP COLUMN deleted_at;

-- 创建一个新的关联表来存储文件和用户 ID 的关系
CREATE TABLE `file_access` (
    `file_id` bigint unsigned NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    PRIMARY KEY (`file_id`, `user_id`),
    FOREIGN KEY (`file_id`) REFERENCES `files` (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 使用 MySQL 客户端连接到数据库，执行以下 SQL 语句查看用户的密码哈希值：
SELECT password FROM users WHERE username = 'testuser';

-- 展示数据库表
SELECT
    TABLE_NAME,
    COLUMN_NAME,
    DATA_TYPE,
    IS_NULLABLE,
    COLUMN_KEY,
    COLUMN_DEFAULT
FROM INFORMATION_SCHEMA.COLUMNS
WHERE
    TABLE_SCHEMA = 'file_system'
    AND TABLE_NAME IN (
        'users',
        'user_roles',
        'roles',
        'permissions',
        'role_permissions',
        'files',
        'file_access'
    )
ORDER BY TABLE_NAME, ORDINAL_POSITION;

-- 查询数据库表具体内容
SELECT * FROM roles;

SELECT * FROM permissions;

SELECT * FROM role_permissions;

SELECT * FROM user_roles;

SELECT * FROM files;

SELECT * FROM file_access;