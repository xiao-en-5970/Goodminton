以下是三个接口的Markdown格式文档，包含完整的请求/响应示例和参数说明：

---

# 用户服务API文档

## 1. 用户注册

### 请求
`POST /api/v1/user/register`

```json
{
  "username": "2022210857",
  "password": "P@ssw0rd123",
  "nickname": "羽毛球小将",
  "avatar_path": "avatars/2022210857.png",
  "self_evaluated_level": "intermediate",
  "system_score": 75,
  "personality_tags": ["外向", "进攻型"],
  "play_style_tags": ["快攻", "网前"],
  "preferred_skill_level": "advanced",
  "preferred_time_slots": ["weekday_evening", "weekend_morning"],
  "preferred_regions": ["海淀区", "朝阳区"],
  "max_cost": 100
}
```

### 参数说明
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 学号/用户名(4-20位字母数字) |
| password | string | 是 | 密码(最少6位) |
| nickname | string | 否 | 显示名称 |
| avatar_path | string | 否 | 头像文件路径 |
| self_evaluated_level | string | 否 | 枚举: beginner/intermediate/advanced |
| personality_tags | string[] | 否 | 最多5个性格标签 |
| preferred_time_slots | string[] | 否 | 枚举值见附录A |

### 成功响应
`201 Created`
```json
{
  "user_id": 123456,
  "username": "2022210857",
  "register_time": "2023-08-20T14:30:00Z"
}
```

### 错误代码
| 代码 | 说明 |
|------|------|
| 40001 | 用户名已存在 |
| 40002 | 密码强度不足 |

---

## 2. 用户登录

### 请求
`POST /api/v1/user/login`

```json
{
  "username": "2022210857",
  "password": "P@ssw0rd123"
}
```

### 参数说明
| 字段 | 类型 | 必填 | 验证规则 |
|------|------|------|----------|
| username | string | 是 | alphanum, min=4, max=20 |
| password | string | 是 | min=6 |

### 成功响应
`200 OK`
```json
{
  "cookie": "auth_token=eyJhbG...",
  "expires_in": 86400
}
```

### 响应头
```
Set-Cookie: auth_token=eyJhbG...; Path=/; HttpOnly; Max-Age=86400
```

### 错误代码
| 代码 | 说明 |
|------|------|
| 40101 | 用户名不存在 |
| 40102 | 密码错误 |

---

## 3. 获取用户信息

### 请求
`GET /api/v1/user/auth/get_userinfo`

### 请求头
```
Cookie: auth_token=eyJhbG...
```

### 成功响应
`200 OK`
```json
{
  "user_id": 123456,
  "username": "2022210857",
  "nickname": "羽毛球小将",
  "avatar_url": "https://cdn.example.com/avatars/2022210857.png",
  "skill_level": "intermediate",
  "tags": {
    "personality": ["外向", "进攻型"],
    "play_style": ["快攻", "网前"]
  },
  "preferences": {
    "regions": ["海淀区", "朝阳区"],
    "time_slots": ["weekday_evening"]
  }
}
```

### 错误代码
根据提供的错误代码，以下是完整的Markdown接口文档补充错误代码部分：

# 错误代码规范

## 全局错误代码 (10xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 10001 | 200 OK | 操作成功 |
| 10002 | 500 Internal Server Error | 服务器内部错误 |
| 10003 | 400 Bad Request | 请求格式错误 |
| 10004 | 500 Internal Server Error | 未知错误 |
| 10005 | 502 Bad Gateway | 错误网关 |

## 用户相关错误 (20xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 20001 | 200 OK | 登录成功 |
| 20002 | 401 Unauthorized | 信息门户校验失败 |
| 20003 | 200 OK | 信息门户校验成功，请完成注册 |
| 20004 | 401 Unauthorized | 密码错误 |
| 20005 | 503 Service Unavailable | 信息门户服务不可用 |
| 20006 | 500 Internal Server Error | 信息门户未知错误 |
| 20007 | 201 Created | 注册成功 |
| 20008 | 404 Not Found | 用户不存在 |
| 20009 | 409 Conflict | 用户已存在 |
| 20010 | 500 Internal Server Error | 用户信息更新失败 |

## 鉴权错误 (30xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 30001 | 401 Unauthorized | 未授权访问 |
| 30002 | 401 Unauthorized | 无效Token |

## 使用示例

### 成功响应
```json
{
  "code": 10001,
  "message": "成功",
  "data": {
    "user_id": 123
  }
}
```

### 错误响应
```json
{
  "code": 20004,
  "message": "密码错误",
  "data": null
}
```
