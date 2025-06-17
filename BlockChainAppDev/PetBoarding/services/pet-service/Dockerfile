# 使用官方Go镜像作为构建环境
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pet-service ./cmd/pet/main.go

# 使用轻量级的alpine镜像作为运行环境
FROM alpine:latest

# 安装必要的CA证书
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/pet-service .

# 暴露端口
EXPOSE 8082

# 运行应用
CMD ["./pet-service"]