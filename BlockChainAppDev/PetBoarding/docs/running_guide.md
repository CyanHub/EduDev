# 宠物寄养系统运行指南

## 环境要求

- Docker 20.10.x 或更高版本
- Docker Compose 2.x 或更高版本
- Git
- 至少 4GB 内存
- 至少 10GB 磁盘空间

## 准备工作

### 1. 确认Docker服务正常运行

在Windows系统中，确保Docker Desktop已正常启动。可以通过以下命令检查Docker状态：

```powershell
docker info
```

如果显示错误，请尝试重启Docker Desktop应用。

### 2. 创建必要的目录结构

```powershell
# 创建日志目录
mkdir -p logs/user-service logs/pet-service logs/boarding-service logs/review-service logs/notification-service logs/admin-service logs/api-gateway
```

### 3. 初始化数据库

系统会在首次启动时自动初始化数据库，无需手动操作。

## 运行系统

### 方法一：使用基本配置（推荐初次使用）

```powershell
# 构建并启动基本服务
docker-compose up -d
```

这将启动核心微服务，但不包含监控和日志系统。

### 方法二：使用扩展配置（完整功能）

```powershell
# 构建并启动所有服务，包括监控和日志系统
docker-compose -f docker-compose.extended.yml up -d
```

## 常见问题解决

### 1. 无法拉取镜像问题

错误信息：`unable to get image 'hashicorp/consul:1.11': error during connect: Get "http://%2F%2F.%2Fpipe%2FdockerDesktopLinuxEngine/v1.49/images/hashicorp/consul:1.11/json": open //./pipe/dockerDesktopLinuxEngine: The system cannot find the file specified.`

解决方法：

1. **重启Docker服务**
   - 关闭Docker Desktop
   - 等待几秒钟
   - 重新启动Docker Desktop
   - 确保Docker服务完全启动（查看状态图标变为绿色）

2. **手动拉取镜像**
   ```powershell
   docker pull hashicorp/consul:1.11
   ```

3. **检查网络连接**
   - 确保您的计算机能够访问互联网
   - 如果使用代理，请确保Docker配置了正确的代理设置

4. **使用镜像加速器**
   - 在Docker Desktop设置中配置镜像加速器

### 2. 端口冲突问题

如果某些端口已被占用，可以修改`docker-compose.yml`或`docker-compose.extended.yml`文件中的端口映射。

### 3. 内存不足问题

如果系统报告内存不足，请在Docker Desktop设置中增加分配给Docker的内存。

## 访问系统

成功启动后，可以通过以下地址访问系统：

- 前端界面：http://localhost:80
- API网关：http://localhost:8080
- 用户服务：http://localhost:8081
- 宠物服务：http://localhost:8082
- 寄养服务：http://localhost:8083
- 评价服务：http://localhost:8084
- 通知服务：http://localhost:8085
- 管理员服务：http://localhost:8086

### 扩展服务（仅在使用docker-compose.extended.yml时可用）

- Consul UI：http://localhost:8500
- Kibana：http://localhost:5601
- Grafana：http://localhost:3000（默认用户名/密码：admin/admin123）
- Prometheus：http://localhost:9090

## 停止系统

```powershell
# 停止基本配置的服务
docker-compose down

# 或停止扩展配置的服务
docker-compose -f docker-compose.extended.yml down
```

## 查看日志

```powershell
# 查看所有服务的日志
docker-compose logs

# 查看特定服务的日志
docker-compose logs user-service
```

## 系统使用

### 用户功能

1. 注册/登录：访问前端界面，点击"注册"或"登录"按钮
2. 宠物管理：添加、编辑和删除您的宠物信息
3. 浏览服务：查看可用的寄养服务，按位置、价格等筛选
4. 预订服务：选择服务，指定日期和宠物，完成预订
5. 查看订单：跟踪订单状态，取消或评价已完成的订单

### 管理员功能

1. 管理员登录：访问管理员登录页面
2. 用户管理：查看、编辑和管理用户信息
3. 服务管理：添加、编辑和管理寄养服务
4. 订单管理：查看和处理订单