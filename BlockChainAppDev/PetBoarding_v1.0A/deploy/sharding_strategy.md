# 数据库分表及缓存策略

## 1. 水平分表方案（订单表）
- 分表依据：用户ID哈希值
- 分表数量：8个
- 表名格式：order_0, order_1, ..., order_7
- 路由规则：user_id % 8

## 2. 垂直分表方案（宠物信息表）
- 基础信息表：pet_basic (id, name, type, age, owner_id)
- 健康档案表：pet_health (id, medical_history, vaccination, last_checkup)

## 3. Redis缓存策略
- 缓存热点订单数据（最近7天的订单）
- 缓存宠物基础信息（TTL: 1小时）
- 使用Redis集群保证高可用

## 4. 中间件配置
- 使用ShardingSphere实现分库分表
- Redis哨兵模式配置