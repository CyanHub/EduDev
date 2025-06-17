# 宠物寄养系统API文档

## 概述

本文档详细说明宠物寄养系统的API接口设计，包括各个微服务提供的接口、请求参数、响应格式和错误码。所有API均通过API网关（http://localhost:8080）进行访问。

## 通用规范

### 请求格式

- 所有API请求均使用HTTP协议
- GET请求参数通过URL查询字符串传递
- POST/PUT/DELETE请求参数通过JSON格式的请求体传递
- 请求头需包含Content-Type: application/json
- 认证信息通过Authorization请求头传递，格式为Bearer {token}

### 响应格式

所有API响应均使用JSON格式，基本结构如下：

```json
{
  "code": 200,          // 状态码，200表示成功，非200表示失败
  "message": "success", // 状态描述
  "data": {}            // 响应数据，可能是对象、数组或null
}
```

### 错误码

| 错误码 | 描述 |
| ------ | ---- |
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 409 | 资源冲突 |
| 500 | 服务器内部错误 |

## 用户服务 (User Service)

基础路径：`/api/users`

### 注册用户

- **URL**: `/api/users/register`
- **方法**: POST
- **描述**: 注册新用户
- **请求参数**:

```json
{
  "username": "johndoe",       // 用户名，必填
  "password": "password123",   // 密码，必填
  "email": "john@example.com", // 邮箱，必填
  "phone": "1234567890",       // 手机号，可选
  "role": 0                     // 角色，0-普通用户，1-服务提供者，默认0
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_id": "u123456789",
    "username": "johndoe",
    "email": "john@example.com",
    "phone": "1234567890",
    "role": 0,
    "status": 0,
    "created_at": "2023-06-01T12:00:00Z"
  }
}
```

### 用户登录

- **URL**: `/api/users/login`
- **方法**: POST
- **描述**: 用户登录
- **请求参数**:

```json
{
  "username": "johndoe",     // 用户名，必填
  "password": "password123"  // 密码，必填
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_id": "u123456789",
    "username": "johndoe",
    "role": 0,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", // JWT令牌
    "expires_at": "2023-06-02T12:00:00Z"                 // 令牌过期时间
  }
}
```

### 获取用户信息

- **URL**: `/api/users/{user_id}`
- **方法**: GET
- **描述**: 获取指定用户的信息
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_id": "u123456789",
    "username": "johndoe",
    "email": "john@example.com",
    "phone": "1234567890",
    "role": 0,
    "status": 1,
    "created_at": "2023-06-01T12:00:00Z",
    "updated_at": "2023-06-01T12:00:00Z"
  }
}
```

### 更新用户信息

- **URL**: `/api/users/{user_id}`
- **方法**: PUT
- **描述**: 更新用户信息
- **请求参数**:

```json
{
  "email": "newemail@example.com", // 新邮箱，可选
  "phone": "9876543210",           // 新手机号，可选
  "password": "newpassword123"     // 新密码，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_id": "u123456789",
    "username": "johndoe",
    "email": "newemail@example.com",
    "phone": "9876543210",
    "role": 0,
    "status": 1,
    "updated_at": "2023-06-01T13:00:00Z"
  }
}
```

## 宠物服务 (Pet Service)

基础路径：`/api/pets`

### 添加宠物

- **URL**: `/api/pets`
- **方法**: POST
- **描述**: 添加新宠物
- **请求参数**:

```json
{
  "name": "Buddy",           // 宠物名称，必填
  "type": "Dog",             // 宠物类型，必填
  "breed": "Golden Retriever", // 品种，可选
  "age": 3,                   // 年龄，可选
  "gender": 1,                // 性别，0-未知，1-雄性，2-雌性，可选
  "weight": 25.5,             // 体重(kg)，可选
  "description": "Friendly and energetic" // 描述，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "pet_id": "p123456789",
    "user_id": "u123456789",
    "name": "Buddy",
    "type": "Dog",
    "breed": "Golden Retriever",
    "age": 3,
    "gender": 1,
    "weight": 25.5,
    "description": "Friendly and energetic",
    "created_at": "2023-06-01T14:00:00Z"
  }
}
```

### 获取宠物信息

- **URL**: `/api/pets/{pet_id}`
- **方法**: GET
- **描述**: 获取指定宠物的信息
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "pet_id": "p123456789",
    "user_id": "u123456789",
    "name": "Buddy",
    "type": "Dog",
    "breed": "Golden Retriever",
    "age": 3,
    "gender": 1,
    "weight": 25.5,
    "description": "Friendly and energetic",
    "created_at": "2023-06-01T14:00:00Z",
    "updated_at": "2023-06-01T14:00:00Z"
  }
}
```

### 获取用户的所有宠物

- **URL**: `/api/pets/user/{user_id}`
- **方法**: GET
- **描述**: 获取指定用户的所有宠物
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "pet_id": "p123456789",
      "user_id": "u123456789",
      "name": "Buddy",
      "type": "Dog",
      "breed": "Golden Retriever",
      "age": 3,
      "gender": 1,
      "weight": 25.5,
      "description": "Friendly and energetic",
      "created_at": "2023-06-01T14:00:00Z",
      "updated_at": "2023-06-01T14:00:00Z"
    },
    {
      "pet_id": "p987654321",
      "user_id": "u123456789",
      "name": "Whiskers",
      "type": "Cat",
      "breed": "Siamese",
      "age": 2,
      "gender": 2,
      "weight": 4.2,
      "description": "Calm and independent",
      "created_at": "2023-06-01T15:00:00Z",
      "updated_at": "2023-06-01T15:00:00Z"
    }
  ]
}
```

### 更新宠物信息

- **URL**: `/api/pets/{pet_id}`
- **方法**: PUT
- **描述**: 更新宠物信息
- **请求参数**:

```json
{
  "name": "Buddy Jr",        // 新名称，可选
  "age": 4,                  // 新年龄，可选
  "weight": 26.2,            // 新体重，可选
  "description": "Very friendly and energetic" // 新描述，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "pet_id": "p123456789",
    "user_id": "u123456789",
    "name": "Buddy Jr",
    "type": "Dog",
    "breed": "Golden Retriever",
    "age": 4,
    "gender": 1,
    "weight": 26.2,
    "description": "Very friendly and energetic",
    "updated_at": "2023-06-01T16:00:00Z"
  }
}
```

### 删除宠物

- **URL**: `/api/pets/{pet_id}`
- **方法**: DELETE
- **描述**: 删除宠物
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "Pet deleted successfully",
  "data": null
}
```

## 寄养服务 (Boarding Service)

基础路径：`/api/boarding`

### 创建寄养服务

- **URL**: `/api/boarding`
- **方法**: POST
- **描述**: 创建新的寄养服务
- **请求参数**:

```json
{
  "title": "Luxury Pet Boarding",  // 服务标题，必填
  "price": 50.00,                 // 价格/天，必填
  "location": "123 Main St, City", // 服务地点，必填
  "pet_type": "Dog",               // 适用宠物类型，必填
  "capacity": 5,                   // 容纳数量，必填
  "description": "Luxury boarding with premium care", // 详细描述，可选
  "facilities": ["Air Conditioning", "Heated Floors", "Outdoor Play Area"], // 设施列表，可选
  "services": ["Daily Walks", "Grooming", "Training"], // 提供的服务列表，可选
  "rules": "Pets must be vaccinated and friendly" // 规则和要求，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "service_id": "s123456789",
    "provider_id": "u123456789",
    "title": "Luxury Pet Boarding",
    "price": 50.00,
    "location": "123 Main St, City",
    "pet_type": "Dog",
    "capacity": 5,
    "available": 1,
    "rating": null,
    "description": "Luxury boarding with premium care",
    "facilities": ["Air Conditioning", "Heated Floors", "Outdoor Play Area"],
    "services": ["Daily Walks", "Grooming", "Training"],
    "rules": "Pets must be vaccinated and friendly",
    "created_at": "2023-06-01T17:00:00Z"
  }
}
```

### 获取寄养服务信息

- **URL**: `/api/boarding/{service_id}`
- **方法**: GET
- **描述**: 获取指定寄养服务的信息
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "service_id": "s123456789",
    "provider_id": "u123456789",
    "provider_name": "John Doe",
    "title": "Luxury Pet Boarding",
    "price": 50.00,
    "location": "123 Main St, City",
    "pet_type": "Dog",
    "capacity": 5,
    "available": 1,
    "rating": 4.8,
    "review_count": 25,
    "description": "Luxury boarding with premium care",
    "facilities": ["Air Conditioning", "Heated Floors", "Outdoor Play Area"],
    "services": ["Daily Walks", "Grooming", "Training"],
    "rules": "Pets must be vaccinated and friendly",
    "photos": ["url1.jpg", "url2.jpg"],
    "created_at": "2023-06-01T17:00:00Z",
    "updated_at": "2023-06-01T17:00:00Z"
  }
}
```

### 搜索寄养服务

- **URL**: `/api/boarding/search`
- **方法**: GET
- **描述**: 搜索寄养服务
- **请求参数**:
  - `pet_type` (string, 可选): 宠物类型
  - `location` (string, 可选): 位置关键词
  - `min_price` (number, 可选): 最低价格
  - `max_price` (number, 可选): 最高价格
  - `start_date` (string, 可选): 开始日期 (YYYY-MM-DD)
  - `end_date` (string, 可选): 结束日期 (YYYY-MM-DD)
  - `page` (number, 可选): 页码，默认1
  - `limit` (number, 可选): 每页数量，默认10
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 50,
    "page": 1,
    "limit": 10,
    "services": [
      {
        "service_id": "s123456789",
        "provider_id": "u123456789",
        "provider_name": "John Doe",
        "title": "Luxury Pet Boarding",
        "price": 50.00,
        "location": "123 Main St, City",
        "pet_type": "Dog",
        "capacity": 5,
        "rating": 4.8,
        "review_count": 25,
        "created_at": "2023-06-01T17:00:00Z"
      },
      // 更多服务...
    ]
  }
}
```

### 更新寄养服务

- **URL**: `/api/boarding/{service_id}`
- **方法**: PUT
- **描述**: 更新寄养服务信息
- **请求参数**:

```json
{
  "title": "Premium Luxury Pet Boarding", // 新标题，可选
  "price": 55.00,                        // 新价格，可选
  "capacity": 6,                          // 新容量，可选
  "available": 0,                         // 是否可用，0-不可用，1-可用，可选
  "description": "Updated description"     // 新描述，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "service_id": "s123456789",
    "title": "Premium Luxury Pet Boarding",
    "price": 55.00,
    "capacity": 6,
    "available": 0,
    "description": "Updated description",
    "updated_at": "2023-06-01T18:00:00Z"
  }
}
```

## 订单服务 (Order Service)

基础路径：`/api/orders`

### 创建订单

- **URL**: `/api/orders`
- **方法**: POST
- **描述**: 创建新订单
- **请求参数**:

```json
{
  "service_id": "s123456789",    // 服务ID，必填
  "pet_id": "p123456789",        // 宠物ID，必填
  "start_date": "2023-07-01",   // 开始日期，必填
  "end_date": "2023-07-05",     // 结束日期，必填
  "special_requests": "Please give medication twice daily" // 特殊要求，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "order_id": "o123456789",
    "user_id": "u123456789",
    "service_id": "s123456789",
    "pet_id": "p123456789",
    "start_date": "2023-07-01",
    "end_date": "2023-07-05",
    "total_price": 250.00,
    "status": 0,
    "payment_status": 0,
    "special_requests": "Please give medication twice daily",
    "created_at": "2023-06-01T19:00:00Z"
  }
}
```

### 获取订单信息

- **URL**: `/api/orders/{order_id}`
- **方法**: GET
- **描述**: 获取指定订单的信息
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "order_id": "o123456789",
    "user_id": "u123456789",
    "service_id": "s123456789",
    "pet_id": "p123456789",
    "service_title": "Luxury Pet Boarding",
    "pet_name": "Buddy",
    "provider_id": "u987654321",
    "provider_name": "Jane Smith",
    "start_date": "2023-07-01",
    "end_date": "2023-07-05",
    "total_price": 250.00,
    "status": 1,
    "payment_status": 1,
    "special_requests": "Please give medication twice daily",
    "care_notes": "Pet is doing well, medication given as requested",
    "created_at": "2023-06-01T19:00:00Z",
    "updated_at": "2023-06-01T20:00:00Z"
  }
}
```

### 获取用户的所有订单

- **URL**: `/api/orders/user/{user_id}`
- **方法**: GET
- **描述**: 获取指定用户的所有订单
- **请求参数**:
  - `status` (number, 可选): 订单状态过滤
  - `page` (number, 可选): 页码，默认1
  - `limit` (number, 可选): 每页数量，默认10
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 15,
    "page": 1,
    "limit": 10,
    "orders": [
      {
        "order_id": "o123456789",
        "service_id": "s123456789",
        "service_title": "Luxury Pet Boarding",
        "pet_id": "p123456789",
        "pet_name": "Buddy",
        "start_date": "2023-07-01",
        "end_date": "2023-07-05",
        "total_price": 250.00,
        "status": 1,
        "payment_status": 1,
        "created_at": "2023-06-01T19:00:00Z"
      },
      // 更多订单...
    ]
  }
}
```

### 获取服务提供者的所有订单

- **URL**: `/api/orders/provider/{provider_id}`
- **方法**: GET
- **描述**: 获取指定服务提供者的所有订单
- **请求参数**:
  - `status` (number, 可选): 订单状态过滤
  - `page` (number, 可选): 页码，默认1
  - `limit` (number, 可选): 每页数量，默认10
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 25,
    "page": 1,
    "limit": 10,
    "orders": [
      {
        "order_id": "o123456789",
        "user_id": "u123456789",
        "user_name": "John Doe",
        "service_id": "s123456789",
        "service_title": "Luxury Pet Boarding",
        "pet_id": "p123456789",
        "pet_name": "Buddy",
        "start_date": "2023-07-01",
        "end_date": "2023-07-05",
        "total_price": 250.00,
        "status": 1,
        "payment_status": 1,
        "created_at": "2023-06-01T19:00:00Z"
      },
      // 更多订单...
    ]
  }
}
```

### 更新订单状态

- **URL**: `/api/orders/{order_id}/status`
- **方法**: PUT
- **描述**: 更新订单状态
- **请求参数**:

```json
{
  "status": 2,  // 新状态，必填，0-待确认，1-已确认，2-进行中，3-已完成，4-已取消
  "care_notes": "Pet is doing well, enjoying daily walks" // 护理记录，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "order_id": "o123456789",
    "status": 2,
    "care_notes": "Pet is doing well, enjoying daily walks",
    "updated_at": "2023-06-01T21:00:00Z"
  }
}
```

### 取消订单

- **URL**: `/api/orders/{order_id}/cancel`
- **方法**: PUT
- **描述**: 取消订单
- **请求参数**:

```json
{
  "reason": "Change of plans" // 取消原因，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "Order cancelled successfully",
  "data": {
    "order_id": "o123456789",
    "status": 4,
    "updated_at": "2023-06-01T22:00:00Z"
  }
}
```

## 评价服务 (Review Service)

基础路径：`/api/reviews`

### 创建评价

- **URL**: `/api/reviews`
- **方法**: POST
- **描述**: 创建新评价
- **请求参数**:

```json
{
  "order_id": "o123456789",    // 订单ID，必填
  "rating": 5,                // 评分（1-5），必填
  "content": "Excellent service, my pet was very happy", // 评价内容，可选
  "photos": ["url1.jpg", "url2.jpg"] // 照片URL列表，可选
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "review_id": "r123456789",
    "order_id": "o123456789",
    "user_id": "u123456789",
    "service_id": "s123456789",
    "rating": 5,
    "content": "Excellent service, my pet was very happy",
    "photos": ["url1.jpg", "url2.jpg"],
    "created_at": "2023-06-01T23:00:00Z"
  }
}
```

### 获取评价信息

- **URL**: `/api/reviews/{review_id}`
- **方法**: GET
- **描述**: 获取指定评价的信息
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "review_id": "r123456789",
    "order_id": "o123456789",
    "user_id": "u123456789",
    "user_name": "John Doe",
    "service_id": "s123456789",
    "service_title": "Luxury Pet Boarding",
    "rating": 5,
    "content": "Excellent service, my pet was very happy",
    "photos": ["url1.jpg", "url2.jpg"],
    "created_at": "2023-06-01T23:00:00Z",
    "updated_at": "2023-06-01T23:00:00Z"
  }
}
```

### 获取服务的所有评价

- **URL**: `/api/reviews/service/{service_id}`
- **方法**: GET
- **描述**: 获取指定服务的所有评价
- **请求参数**:
  - `page` (number, 可选): 页码，默认1
  - `limit` (number, 可选): 每页数量，默认10
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 30,
    "page": 1,
    "limit": 10,
    "average_rating": 4.8,
    "reviews": [
      {
        "review_id": "r123456789",
        "user_id": "u123456789",
        "user_name": "John Doe",
        "rating": 5,
        "content": "Excellent service, my pet was very happy",
        "photos": ["url1.jpg", "url2.jpg"],
        "created_at": "2023-06-01T23:00:00Z"
      },
      // 更多评价...
    ]
  }
}
```

### 获取用户的所有评价

- **URL**: `/api/reviews/user/{user_id}`
- **方法**: GET
- **描述**: 获取指定用户的所有评价
- **请求参数**:
  - `page` (number, 可选): 页码，默认1
  - `limit` (number, 可选): 每页数量，默认10
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 15,
    "page": 1,
    "limit": 10,
    "reviews": [
      {
        "review_id": "r123456789",
        "service_id": "s123456789",
        "service_title": "Luxury Pet Boarding",
        "rating": 5,
        "content": "Excellent service, my pet was very happy",
        "created_at": "2023-06-01T23:00:00Z"
      },
      // 更多评价...
    ]
  }
}
```

## 通知服务 (Notification Service)

基础路径：`/api/notifications`

### 获取用户的所有通知

- **URL**: `/api/notifications/user/{user_id}`
- **方法**: GET
- **描述**: 获取指定用户的所有通知
- **请求参数**:
  - `read` (number, 可选): 是否已读，0-未读，1-已读
  - `page` (number, 可选): 页码，默认1
  - `limit` (number, 可选): 每页数量，默认10
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 20,
    "page": 1,
    "limit": 10,
    "unread_count": 5,
    "notifications": [
      {
        "notification_id": "n123456789",
        "type": 1,
        "title": "Order Confirmed",
        "content": "Your order #o123456789 has been confirmed",
        "related_id": "o123456789",
        "read": 0,
        "created_at": "2023-06-02T10:00:00Z"
      },
      // 更多通知...
    ]
  }
}
```

### 获取通知详情

- **URL**: `/api/notifications/{notification_id}`
- **方法**: GET
- **描述**: 获取指定通知的详情
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "notification_id": "n123456789",
    "user_id": "u123456789",
    "type": 1,
    "title": "Order Confirmed",
    "content": "Your order #o123456789 has been confirmed",
    "related_id": "o123456789",
    "read": 0,
    "created_at": "2023-06-02T10:00:00Z"
  }
}
```

### 标记通知为已读

- **URL**: `/api/notifications/{notification_id}/read`
- **方法**: PUT
- **描述**: 标记指定通知为已读
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "Notification marked as read",
  "data": null
}
```

### 标记所有通知为已读

- **URL**: `/api/notifications/user/{user_id}/read-all`
- **方法**: PUT
- **描述**: 标记指定用户的所有通知为已读
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "All notifications marked as read",
  "data": null
}
```

## 管理员服务 (Admin Service)

基础路径：`/api/admin`

### 管理员登录

- **URL**: `/api/admin/login`
- **方法**: POST
- **描述**: 管理员登录
- **请求参数**:

```json
{
  "username": "admin",      // 用户名，必填
  "password": "admin123"   // 密码，必填
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "admin_id": "a123456789",
    "username": "admin",
    "role": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", // JWT令牌
    "expires_at": "2023-06-03T10:00:00Z"                 // 令牌过期时间
  }
}
```

### 获取系统统计数据

- **URL**: `/api/admin/dashboard`
- **方法**: GET
- **描述**: 获取系统统计数据
- **请求参数**: 无
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_count": 1000,
    "provider_count": 200,
    "service_count": 500,
    "order_count": 2000,
    "pending_orders": 50,
    "completed_orders": 1500,
    "cancelled_orders": 100,
    "total_revenue": 75000.00,
    "average_rating": 4.7,
    "recent_users": [
      {
        "user_id": "u123456789",
        "username": "johndoe",
        "created_at": "2023-06-01T12:00:00Z"
      },
      // 更多用户...
    ],
    "recent_orders": [
      {
        "order_id": "o123456789",
        "user_id": "u123456789",
        "service_id": "s123456789",
        "total_price": 250.00,
        "status": 1,
        "created_at": "2023-06-01T19:00:00Z"
      },
      // 更多订单...
    ]
  }
}
```

### 获取所有用户

- **URL**: `/api/admin/users`
- **方法**: GET
- **描述**: 获取所有用户
- **请求参数**:
  - `role` (number, 可选): 角色过滤
  - `status` (number, 可选): 状态过滤
  - `page` (number, 可选): 页码，默认1
  - `limit` (number, 可选): 每页数量，默认10
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 1000,
    "page": 1,
    "limit": 10,
    "users": [
      {
        "user_id": "u123456789",
        "username": "johndoe",
        "email": "john@example.com",
        "phone": "1234567890",
        "role": 0,
        "status": 1,
        "created_at": "2023-06-01T12:00:00Z"
      },
      // 更多用户...
    ]
  }
}
```

### 审核用户

- **URL**: `/api/admin/users/{user_id}/status`
- **方法**: PUT
- **描述**: 更新用户状态
- **请求参数**:

```json
{
  "status": 1  // 新状态，必填，0-待审核，1-正常，2-禁用
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "User status updated successfully",
  "data": {
    "user_id": "u123456789",
    "status": 1,
    "updated_at": "2023-06-02T11:00:00Z"
  }
}
```

### 获取所有服务

- **URL**: `/api/admin/services`
- **方法**: GET
- **描述**: 获取所有寄养服务
- **请求参数**:
  - `pet_type` (string, 可选): 宠物类型过滤
  - `available` (number, 可选): 可用状态过滤
  - `page` (number, 可选): 页码，默认1
  - `limit` (number, 可选): 每页数量，默认10
- **响应**:

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total": 500,
    "page": 1,
    "limit": 10,
    "services": [
      {
        "service_id": "s123456789",
        "provider_id": "u987654321",
        "provider_name": "Jane Smith",
        "title": "Luxury Pet Boarding",
        "price": 50.00,
        "location": "123 Main St, City",
        "pet_type": "Dog",
        "capacity": 5,
        "available": 1,
        "rating": 4.8,
        "created_at": "2023-06-01T17:00:00Z"
      },
      // 更多服务...
    ]
  }
}
```

### 审核服务

- **URL**: `/api/admin/services/{service_id}/status`
- **方法**: PUT
- **描述**: 更新服务状态
- **请求参数**:

```json
{
  "available": 1  // 新状态，必填，0-不可用，1-可用
}
```

- **响应**:

```json
{
  "code": 200,
  "message": "Service status updated successfully",
  "data": {
    "service_id": "s123456789",
    "available": 1,
    "updated_at": "2023-06-02T12:00:00Z"
  }
}
```

## 健康检查

所有服务都提供健康检查接口：

- **URL**: `/health`
- **方法**: GET
- **描述**: 检查服务健康状态
- **请求参数**: 无
- **响应**:

```json
{
  "status": "UP",
  "service": "user-service",
  "timestamp": "2023-06-02T13:00:00Z"
}
```

## 错误响应示例

当API请求失败时，将返回以下格式的错误响应：

```json
{
  "code": 400,          // 错误码
  "message": "Invalid parameters", // 错误描述
  "data": null          // 数据为null
}
```

## API版本控制

未来API版本更新时，将通过URL路径进行版本控制，例如：

- `/api/v1/users`
- `/api/v2/users`

当前文档描述的是v1版本的API。