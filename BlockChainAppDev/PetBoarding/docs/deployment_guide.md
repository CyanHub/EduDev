# 宠物寄养系统部署指南

## 概述

本文档提供宠物寄养系统的详细部署指南，包括环境要求、Docker部署、本地开发部署以及常见问题解决方案。

## 环境要求

### 基础环境

- Docker 20.10.x 或更高版本
- Docker Compose 2.x 或更高版本
- Git
- 至少 4GB 内存
- 至少 10GB 磁盘空间

### 开发环境（可选）

- Go 1.21 或更高版本
- MySQL 8.0 或更高版本（未来计划）
- Redis 6.0 或更高版本（未来计划）
- Kafka 3.0 或更高版本（未来计划）

## Docker 部署（推荐）

使用 Docker Compose 是部署宠物寄养系统最简单的方式，它会自动构建和启动所有必要的服务。

### 步骤 1：克隆代码仓库

```bash
git clone https://github.com/cyanhub/petboarding.git
cd petboarding
```

### 步骤 2：配置环境变量（可选）

系统默认使用 `docker-compose.yml` 中的环境变量配置。如果需要自定义配置，可以创建 `.env` 文件：

```bash
# 创建环境变量文件
cp .env.example .env

# 编辑环境变量
# 根据需要修改 .env 文件中的配置
```

### 步骤 3：构建和启动服务

```bash
# 构建并启动所有服务（后台运行）
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f api-gateway
```

### 步骤 4：访问系统

服务启动后，可以通过以下地址访问系统：

- API 网关：http://localhost:8080
- 用户服务：http://localhost:8081
- 宠物服务：http://localhost:8082
- 寄养服务：http://localhost:8083
- 评价服务：http://localhost:8084
- 通知服务：http://localhost:8085
- 管理员服务：http://localhost:8086

### 步骤 5：停止服务

```bash
# 停止所有服务但不删除容器
docker-compose stop

# 停止并删除所有容器、网络
docker-compose down

# 停止并删除所有容器、网络和卷（慎用）
docker-compose down -v
```

## 本地开发部署

如果需要进行开发或调试，可以选择本地部署方式。

### 步骤 1：克隆代码仓库

```bash
git clone https://github.com/cyanhub/petboarding.git
cd petboarding
```

### 步骤 2：安装依赖

```bash
go mod download
```

### 步骤 3：启动各个服务

需要打开多个终端窗口，分别启动各个服务：

**用户服务**
```bash
cd services/user-service
go run main.go
```

**宠物服务**
```bash
cd services/pet-service
go run main.go
```

**寄养服务**
```bash
cd services/boarding-service
go run main.go
```

**评价服务**
```bash
cd services/review-service
go run main.go
```

**通知服务**
```bash
cd services/notification-service
go run main.go
```

**管理员服务**
```bash
cd services/admin-service
go run main.go
```

**API 网关**
```bash
cd api-gateway
go run main.go
```

## 环境变量配置

以下是系统中使用的主要环境变量及其说明：

### 通用环境变量

| 环境变量 | 描述 | 默认值 |
| ------- | ---- | ------ |
| `PORT` | 服务监听端口 | 各服务不同 |
| `GIN_MODE` | Gin框架运行模式 | `debug` |
| `LOG_LEVEL` | 日志级别 | `info` |

### API 网关环境变量

| 环境变量 | 描述 | 默认值 |
| ------- | ---- | ------ |
| `USER_SERVICE_URL` | 用户服务URL | `http://user-service:8081` |
| `PET_SERVICE_URL` | 宠物服务URL | `http://pet-service:8082` |
| `BOARDING_SERVICE_URL` | 寄养服务URL | `http://boarding-service:8083` |
| `REVIEW_SERVICE_URL` | 评价服务URL | `http://review-service:8084` |
| `NOTIFICATION_SERVICE_URL` | 通知服务URL | `http://notification-service:8085` |
| `ADMIN_SERVICE_URL` | 管理员服务URL | `http://admin-service:8086` |

### 数据库环境变量（未来计划）

| 环境变量 | 描述 | 默认值 |
| ------- | ---- | ------ |
| `DB_HOST` | 数据库主机 | `mysql` |
| `DB_PORT` | 数据库端口 | `3306` |
| `DB_USER` | 数据库用户名 | `petboarding` |
| `DB_PASSWORD` | 数据库密码 | `petboarding123` |
| `DB_NAME` | 数据库名称 | `petboarding` |

### Redis环境变量（未来计划）

| 环境变量 | 描述 | 默认值 |
| ------- | ---- | ------ |
| `REDIS_HOST` | Redis主机 | `redis` |
| `REDIS_PORT` | Redis端口 | `6379` |
| `REDIS_PASSWORD` | Redis密码 | `` |

### Kafka环境变量（未来计划）

| 环境变量 | 描述 | 默认值 |
| ------- | ---- | ------ |
| `KAFKA_BROKERS` | Kafka代理地址列表 | `kafka:9092` |

## 服务扩展

### 水平扩展

可以通过修改 `docker-compose.yml` 文件中的 `replicas` 参数来水平扩展服务：

```yaml
services:
  user-service:
    deploy:
      replicas: 3  # 启动3个实例
```

或者使用 Docker Compose 命令：

```bash
docker-compose up -d --scale user-service=3
```

### 负载均衡

系统使用API网关进行负载均衡。当服务实例增加时，API网关会自动分发请求到不同的服务实例。

## 监控和日志

### 查看容器日志

```bash
# 查看所有服务的日志
docker-compose logs -f

# 查看特定服务的日志
docker-compose logs -f user-service
```

### 健康检查

所有服务都提供了健康检查接口，可以通过以下方式访问：

```
GET http://localhost:{PORT}/health
```

## 常见问题解决

### 服务无法启动

1. 检查端口是否被占用：
   ```bash
   netstat -ano | findstr "8080 8081 8082 8083 8084 8085 8086"
   ```

2. 检查Docker服务是否正常运行：
   ```bash
   docker info
   ```

3. 检查日志查找错误：
   ```bash
   docker-compose logs
   ```

### 服务之间无法通信

1. 检查网络配置：
   ```bash
   docker network ls
   docker network inspect petboarding-network
   ```

2. 检查环境变量是否正确配置：
   ```bash
   docker-compose config
   ```

3. 检查服务健康状态：
   ```bash
   curl http://localhost:{PORT}/health
   ```

### 容器内存不足

1. 增加Docker可用内存：
   - Windows/Mac：在Docker Desktop设置中增加内存限制
   - Linux：修改Docker守护进程配置

2. 优化服务内存使用：
   ```yaml
   services:
     user-service:
       deploy:
         resources:
           limits:
             memory: 256M
   ```

## 生产环境部署建议

### 安全性配置

1. 使用HTTPS：配置API网关使用SSL证书
2. 设置强密码：修改所有默认密码
3. 限制网络访问：只开放必要的端口
4. 定期更新：保持系统和依赖库的更新

### 高可用性配置

1. 使用多个服务实例
2. 配置数据库主从复制
3. 使用Redis集群
4. 实施自动故障转移

### 备份策略

1. 定期备份数据库
2. 备份配置文件
3. 设置自动备份计划

## 未来部署计划

1. Kubernetes部署支持
2. CI/CD流水线集成
3. 自动扩缩容配置
4. 完整监控系统集成（Prometheus + Grafana）
5. 分布式追踪（Jaeger/Zipkin）

## 支持和反馈

如有部署问题或建议，请通过以下方式联系我们：

- 提交GitHub Issue
- 发送邮件至：support@petboarding.example.com（示例）