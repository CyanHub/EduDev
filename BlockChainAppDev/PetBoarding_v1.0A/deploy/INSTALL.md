# 宠物寄养系统安装部署手册

## 环境要求
- Go 1.20+
- MySQL 8.0+
- Redis 6.0+
- Nginx 1.18+
- Docker 20.10+ (可选)

## 一、基础环境安装

### 1. 安装Go语言环境
```bash
# Windows
choco install golang

# Linux
wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 2. 安装MySQL数据库
```bash
# Windows
choco install mysql

# Linux
sudo apt install mysql-server
sudo systemctl start mysql
```

### 3. 安装Redis
```bash
# Windows
choco install redis-64

# Linux
sudo apt install redis-server
sudo systemctl start redis
```

## 二、项目部署

### 1. 克隆代码库
```bash
git clone https://github.com/your-repo/pet-boarding.git
cd pet-boarding
```

### 2. 初始化数据库
```bash
mysql -u root -p < database/pet_boarding.sql
```

### 3. 配置数据库连接
修改config/database.yaml：
```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  user: root
  password: your_password
  dbname: pet_boarding

redis:
  addr: 127.0.0.1:6379
  password: ""
  db: 0
```

### 4. 安装依赖
```bash
go mod tidy
```

## 三、微服务启动

### 启动顺序：
1. 基础服务
```bash
# Redis
redis-server &

# MySQL
mysqld &
```

2. 微服务（每个服务独立终端）
```bash
# 用户服务
cd cmd/user_service && go run main.go

# 宠物服务
cd cmd/pet_service && go run main.go

# 订单服务
cd cmd/order_service && go run main.go

# 评价服务
cd cmd/review_service && go run main.go

# API网关
cd cmd/gateway && go run main.go
```

## 四、Nginx负载均衡配置
修改deploy/nginx.conf：
```nginx
http {
    upstream user_service {
        server 127.0.0.1:8081;
        server 127.0.0.1:8082;
    }

    upstream pet_service {
        server 127.0.0.1:8083;
    }

    server {
        listen 80;

        location /api/users {
            proxy_pass http://user_service;
        }

        location /api/pets {
            proxy_pass http://pet_service;
        }
    }
}
```
启动Nginx：
```bash
nginx -c deploy/nginx.conf
```

## 五、容器化部署（可选）

### 1. 构建Docker镜像
```bash
docker build -t user-service -f cmd/user_service/Dockerfile .
```

### 2. docker-compose.yaml示例
```yaml
version: '3'

services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root
    volumes:
      - ./database:/docker-entrypoint-initdb.d

  redis:
    image: redis:6-alpine

  user_service:
    build: ./cmd/user_service
    ports:
      - "8081:8080"

  nginx:
    image: nginx:1.21
    ports:
      - "80:80"
    volumes:
      - ./deploy/nginx.conf:/etc/nginx/nginx.conf
```

启动容器集群：
```bash
docker-compose up -d
```

## 六、验证部署
```bash
curl http://localhost/api/users/health
curl http://localhost/api/pets/1
```

## 常见问题
1. 端口冲突：检查8080-8083端口占用情况
2. 数据库连接失败：确认database.yaml配置
3. Redis连接异常：检查redis-server是否运行