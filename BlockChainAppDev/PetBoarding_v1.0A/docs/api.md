# Zooer API 文档

## 用户服务

### 用户注册
`POST /register`
请求参数:
- username: string
- password: string

响应:
- message: string

### 用户登录
`POST /login`
请求参数:
- username: string
- password: string

响应:
- message: string

## 宠物服务

### 添加宠物
`POST /pets`
请求参数:
- name: string
- type: string

响应:
- id: string
- name: string
- type: string

### 获取宠物详情
`GET /pets/:id`

响应:
- id: string
- name: string
- type: string

### 更新宠物
`PUT /pets/:id`
请求参数:
- name: string
- type: string

响应:
- id: string
- name: string
- type: string

### 删除宠物
`DELETE /pets/:id`

响应:
- message: string

## 订单服务

### 创建订单
`POST /orders`
请求参数:
- amount: int

响应:
- id: string
- amount: int
- status: string

### 支付订单
`PUT /orders/:id/pay`

响应:
- id: string
- amount: int
- status: string

## 评价服务

### 添加评价
`POST /reviews`
请求参数:
- content: string
- rating: int

响应:
- id: string
- content: string
- rating: int

### 获取评价详情
`GET /reviews/:id`

响应:
- id: string
- content: string
- rating: int

### 更新评价
`PUT /reviews/:id`
请求参数:
- content: string
- rating: int

响应:
- id: string
- content: string
- rating: int

### 删除评价
`DELETE /reviews/:id`

响应:
- message: string