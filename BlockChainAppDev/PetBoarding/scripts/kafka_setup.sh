#!/bin/bash

# 宠物寄养系统Kafka配置脚本
# 此脚本用于初始化Kafka主题和配置

# 确保脚本在错误时退出
set -e

echo "开始配置Kafka..."

# Kafka服务器地址
KAFKA_BROKER="kafka:9092"

# 等待Kafka启动
echo "等待Kafka启动..."
sleep 10

# 创建主题函数
create_topic() {
    topic_name=$1
    partitions=$2
    replication_factor=$3
    retention_ms=$4
    
    echo "创建主题: $topic_name (分区: $partitions, 副本: $replication_factor, 保留时间: $retention_ms)"
    
    kafka-topics --create \
        --bootstrap-server $KAFKA_BROKER \
        --topic $topic_name \
        --partitions $partitions \
        --replication-factor $replication_factor \
        --config retention.ms=$retention_ms \
        --if-not-exists
}

# 创建用户相关主题
echo "创建用户相关主题..."
create_topic "user-events" 3 1 604800000  # 7天保留期
create_topic "user-notifications" 3 1 604800000  # 7天保留期

# 创建宠物相关主题
echo "创建宠物相关主题..."
create_topic "pet-events" 3 1 604800000  # 7天保留期

# 创建寄养服务相关主题
echo "创建寄养服务相关主题..."
create_topic "service-events" 3 1 604800000  # 7天保留期
create_topic "service-availability" 3 1 86400000  # 1天保留期

# 创建订单相关主题
echo "创建订单相关主题..."
create_topic "order-events" 5 1 2592000000  # 30天保留期
create_topic "order-status-changes" 5 1 2592000000  # 30天保留期
create_topic "payment-events" 3 1 2592000000  # 30天保留期

# 创建评价相关主题
echo "创建评价相关主题..."
create_topic "review-events" 3 1 7776000000  # 90天保留期

# 创建通知相关主题
echo "创建通知相关主题..."
create_topic "email-notifications" 3 1 604800000  # 7天保留期
create_topic "sms-notifications" 3 1 604800000  # 7天保留期
create_topic "push-notifications" 3 1 604800000  # 7天保留期

# 创建系统相关主题
echo "创建系统相关主题..."
create_topic "system-events" 1 1 2592000000  # 30天保留期
create_topic "audit-logs" 1 1 7776000000  # 90天保留期

# 列出所有主题
echo "已创建的主题列表:"
kafka-topics --list --bootstrap-server $KAFKA_BROKER

echo "Kafka配置完成!"

# 主题使用说明
cat << EOF

宠物寄养系统Kafka主题使用说明：

1. 用户相关主题：
   - user-events: 用户注册、登录、更新等事件
   - user-notifications: 发送给用户的通知事件

2. 宠物相关主题：
   - pet-events: 宠物添加、更新、删除等事件

3. 寄养服务相关主题：
   - service-events: 服务创建、更新、状态变更等事件
   - service-availability: 服务可用性变更事件

4. 订单相关主题：
   - order-events: 订单创建、更新等事件
   - order-status-changes: 订单状态变更事件
   - payment-events: 支付相关事件

5. 评价相关主题：
   - review-events: 评价创建、回复等事件

6. 通知相关主题：
   - email-notifications: 邮件通知事件
   - sms-notifications: 短信通知事件
   - push-notifications: 推送通知事件

7. 系统相关主题：
   - system-events: 系统级事件
   - audit-logs: 审计日志事件

消息格式规范：
所有消息应使用JSON格式，包含以下基本字段：
- event_id: 事件唯一标识
- event_type: 事件类型
- source: 事件源服务
- timestamp: 事件时间戳
- data: 事件数据

示例：
{
  "event_id": "evt123456789",
  "event_type": "order_created",
  "source": "order-service",
  "timestamp": "2023-06-01T12:00:00Z",
  "data": {
    "order_id": "o123456789",
    "user_id": "u123456789",
    "service_id": "s123456789",
    "pet_id": "p123456789",
    "total_price": 250.00
  }
}

EOF