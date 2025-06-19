# 基于 Go 语言的文件系统开发文档

## 一、项目概述

本项目是一个基于 Go 语言和相关框架（如 Gin、GORM）设计并实现的文档管理系统，涵盖文档的上传、下载、权限分配和配置管理功能。系统旨在提供高效的文档存储、访问控制和操作记录功能。

## 二、功能要求实现情况

### 1. 基于 JWT 的用户身份认证

* 在 **auth/auth.go** 文件中实现了 `GenerateToken` 和 `ParseToken` 函数，用于生成和验证 JWT 令牌。
* 在 **handlers/handlers.go** 文件的 `LoginUser` 函数中，用户登录成功后会生成 JWT 令牌返回给客户端。

### 2. 基于 RBAC 模型设计权限管理

* 在 **models/models.go** 文件中定义了 `User`、`Role` 和 `Permission` 模型，通过多对多关系实现角色和权限的分配。
* 在 **handlers/handlers.go** 文件的 `UploadFile` 和 `DownloadFile` 函数中，会检查用户的角色和权限，以决定是否允许操作。

### 3. 使用 GORM 实现数据库操作

* 在 **main.go** 文件中，使用 GORM 连接数据库，并通过 `db.AutoMigrate` 自动迁移数据库表结构。
* 在 **handlers/handlers.go** 文件中，使用 GORM 进行用户、角色、权限和文件的增删改查操作。

### 4. 集成 Zap 日志库

* 在 **logging/logger.go** 文件中，使用 Zap 日志库记录用户操作日志，如登录、权限变更等。

### 5. 使用 Viper 库实现多格式配置管理

* 在 **config/config.go** 文件中，使用 Viper 库加载 **config.yaml** 配置文件，动态加载数据库和日志配置。

### 6. 实现文件上传、下载、文件权限管理功能

* 在 **handlers/handlers.go** 文件中，实现了 `UploadFile` 和 `DownloadFile` 函数，分别处理文件上传和下载请求，并进行权限检查。

## 三、技术要求实现情况

### 1. 数据库表结构合理性

* 在 **sql/file_system.sql** 文件中，定义了用户表、角色表、权限表、文件表以及关联表，使用了合适的字段类型和索引。

### 2. 实现事务操作

目前项目中未明确实现事务操作，可在涉及多表操作（如权限分配）时，使用 GORM 的事务功能来确保数据一致性。示例代码如下：

```go
db := c.MustGet("db").(*gorm.DB)
tx := db.Begin()
if tx.Error != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "开始事务失败"})
    return
}

// 执行多个数据库操作
if err := tx.Create(&role).Error; err != nil {
    tx.Rollback()
    c.JSON(http.StatusBadRequest, gin.H{"error": "角色创建失败"})
    return
}

if err := tx.Create(&permission).Error; err != nil {
    tx.Rollback()
    c.JSON(http.StatusBadRequest, gin.H{"error": "权限创建失败"})
    return
}

// 提交事务
if err := tx.Commit().Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "提交事务失败"})
    return
}

c.JSON(http.StatusCreated, gin.H{"message": "角色和权限创建成功"})
```

## 四、项目启动与使用

### 1. 启动 Gin 框架 web 服务

* 确保 Go 环境已安装，项目依赖已下载（使用 `go mod tidy` 下载依赖）。
* 运行 `go run main.go` 启动 Gin 框架 web 服务。

### 2. 启动 Python 静态文件服务器

* 打开终端，进入项目的 `web` 目录。
* 运行 `python -m http.server 4408` 启动 Python 静态文件服务器。

### 3. 用户注册与登录

* 打开浏览器，访问 `http://localhost:4408/pages/register.html` 进行用户注册。
* 访问 `http://localhost:4408/pages/login.html` 进行用户登录，登录成功后会返回 JWT 令牌。

### 4. 文件上传与下载

* 访问 `http://localhost:4408/pages/upload.html` 进行文件上传，上传时会进行权限检查。
* 访问 `http://localhost:4408/pages/download.html` 进行文件下载，下载时会进行权限检查。

### 5. 权限和角色管理

* 使用 ApiFox 发送 POST 请求到 `/roles` 和 `/permissions` 接口，创建角色和权限。
* 在 NaviCat 数据库查询页面输入 SQL 语句，将权限赋予给对应角色，示例 SQL 如下：

```sql
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 1);
```

## 五、代码结构说明

### 1. `handlers` 目录

* **handlers.go**：处理各种 HTTP 请求，如用户注册、登录、文件上传、下载、角色和权限创建等。

### 2. `models` 目录

* **models.go**：定义数据库模型，包括用户、角色、权限和文件等。

### 3. `auth` 目录

* **auth.go**：实现 JWT 令牌的生成和验证功能。

### 4. `config` 目录

* **config.go**：使用 Viper 库加载配置文件。
* **config.yaml**：配置文件，包含数据库、日志和服务器等配置信息。

### 5. `logging` 目录

* **logger.go**：使用 Zap 日志库记录用户操作日志。

### 6. `web` 目录

* `pages` 目录：包含前端页面，如注册、登录、文件上传和下载页面。
* `static` 目录：包含静态资源，如 CSS 和 JavaScript 文件。

### 7. `sql` 目录

* **file_system.sql**：数据库表结构定义。

### 8. **main.go**

* 项目入口文件，初始化配置、连接数据库、启动 Gin 框架 web 服务。

## 六、注意事项

* 确保数据库服务已启动，并且 **config.yaml** 中的数据库配置信息正确。
* 确保 Python 环境已安装，并且可以正常运行 `python -m http.server` 命令。
* 在使用 ApiFox 进行接口测试时，注意请求的 URL、请求方法和请求参数。
* 在 NaviCat 中执行 SQL 语句时，注意语法正确性和数据一致性。
