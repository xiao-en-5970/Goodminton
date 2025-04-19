# Goodminton Backend

# API 文档

## 目录
- [基础信息](#基础信息)
- [认证方式](#认证方式)
- [全局响应](#全局响应)
- [接口列表](#接口列表)
  - [用户模块](#用户模块)
- [错误代码](#错误代码)

---

## 基础信息
- ​**Base URL**: `https://api.example.com/v1`
- ​**协议**: HTTPS
- ​**数据格式**: JSON
- ​**编码**: UTF-8

## 认证方式
```http
Authorization: Bearer {access_token}
```

## 全局响应
### 成功响应
```http
{
  "code": 200,
  "data": {},
  "msg": "success"
}
```
### 错误响应
```http
{
  "code": 503,
  "msg": "服务不可用..."
}
```

## 接口列表
### 通用

### 用户模块
#### 1. 用户登录
Endpoint: POST /auth/login

请求参数:

参数名	类型	必填	说明
username	string	是	用户名，4-20位字符
password	string	是	密码，6位以上
请求示例:

json
复制
{
  "username": "testuser",
  "password": "mypassword123"
}
响应字段:

json
复制
{
  "code": 200,
  "data": {
    "token": "eyJhbGciOi...",
    "expires_in": 86400
  }
}


