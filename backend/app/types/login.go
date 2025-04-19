package types

// LoginReq 登录请求参数
// swagger:parameters login
type LoginReq struct {
    // 用户名，4-20位字母数字组合
    // Required: true
    // Example: admin123
    Username string `json:"username" binding:"required,alphanum,min=4,max=20"`
    
    // 密码，最少6位
    // Required: true
    // Example: password123
    Password string `json:"password" binding:"required,min=6"`
}

// LoginResp 登录响应数据
// swagger:model LoginResp
type LoginResp struct {
    // 用户ID
    // Example: 123456
    UserId int64 `json:"user_id"`
    
    // 用户名
    // Example: admin123
    Username string `json:"username"`
    
    // 电子邮箱
    // Example: user@example.com
    Email string `json:"email,omitempty"`
    
    // 手机号码
    // Example: 13800138000
    Phone string `json:"phone,omitempty"`
    
    // 用户昵称
    // Example: 游戏达人
    Nickname string `json:"nickname,omitempty"`
    
    // 头像路径
    // Example: /avatars/default.png
    AvatarPath string `json:"avatar_path,omitempty"`
    
    // 自评等级
    // Example: 高级玩家
    SelfEvaluatedLevel string `json:"self_evaluatedLevel,omitempty"`
    
    // 系统评分
    // Example: 85
    SystemScore int32 `json:"system_score,omitempty"`
    
    // 个性标签
    // Example: ["竞技","团队合作"]
    PersonalityTags []string `json:"personality_tags,omitempty"`
    
    // 玩法偏好
    // Example: ["PVP","副本"]
    PlayStyleTags []string `json:"playStyle_tags,omitempty"`
    
    // 创建时间
    // Example: 2023-01-01 12:00:00
    CreateTime string `json:"createTime,omitempty"`
}