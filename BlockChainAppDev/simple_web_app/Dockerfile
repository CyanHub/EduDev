# 分阶段构建

#源镜像
FROM golang:1.24.3 AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
GO111MODULE=on \
CGO_ENABLED=1

COPY . .
RUN go mod download

#go构建可执行文件
RUN go build -o simple_web_app

# 第二阶段构建
FROM docker.1ms.run/xiaoxijin/apline

#作者
LABEL maintainer="test"
LABEL version="1.0"
LABEL description="This is a docker image for my web application."
# MAINTAINER Razil "test"
## 在docker的根目录下创建相应的使用目录
RUN mkdir -p /app
## 设置工作目录
WORKDIR /app
## 从前一个阶段复制构建好的可执行文件
COPY --from=builder /app/simple_web_app .

#应用端口
EXPOSE 8080

RUN chmod +x simple_web_app
ENTRYPOINT ["./simple_web_app"]

# 设置容器启动时运行的命令
# 这里是运行编译好的可执行文件 simple_web_app
# CMD ["/app/simple_web_app"]