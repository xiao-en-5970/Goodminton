// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2

package types

type IdReq struct {
	UserId int64 `json:"user_id"` // 用户ID
}

type PageReq struct {
	Page     int64 `form:"page,default=1"`       // 页码
	PageSize int64 `form:"page_size,default=20"` // 每页大小
}

type PageResp struct {
	Total int64      `json:"total"` // 总数
	List  []UserInfo `json:"list"`  // 用户列表
}

type RegisterReq struct {
	Nickname            string   `json:"nickname,optional"`              // 用户昵称
	AvatarPath          string   `json:"avatar_path,optional"`           // 头像路径
	SelfEvaluatedLevel  string   `json:"self_evaluated_level,optional"`  // 自评技术水平
	PersonalityTags     []string `json:"personality_tags,optional"`      // 性格标签
	PlayStyleTags       []string `json:"play_style_tags,optional"`       // 打球风格
	PreferredSkillLevel string   `json:"preferred_skill_level,optional"` // 希望对手技术水平
	PreferredTimeSlots  []string `json:"preferred_time_slots,optional"`  // 时间偏好
	PreferredRegions    []string `json:"preferred_regions,optional"`     // 常活动区域
	MaxCost             int64    `json:"max_cost,optional,range=[0:]"`   // 可接受花销
}

type UpdateReq struct {
	Nickname            string   `json:"nickname,optional"`              // 用户昵称
	AvatarPath          string   `json:"avatar_path,optional"`           // 头像路径
	SelfEvaluatedLevel  string   `json:"self_evaluated_level,optional"`  // 自评技术水平
	PersonalityTags     []string `json:"personality_tags,optional"`      // 性格标签
	PlayStyleTags       []string `json:"play_style_tags,optional"`       // 打球风格
	PreferredSkillLevel string   `json:"preferred_skill_level,optional"` // 希望对手技术水平
	PreferredTimeSlots  []string `json:"preferred_time_slots,optional"`  // 时间偏好
	PreferredRegions    []string `json:"preferred_regions,optional"`     // 常活动区域
	MaxCost             int64    `json:"max_cost,optional,range=[0:]"`   // 可接受花销
}

type UserInfo struct {
	UserId              int64    `json:"user_id"`               // 用户ID
	Nickname            string   `json:"nickname"`              // 用户昵称
	AvatarPath          string   `json:"avatar_path"`           // 头像路径
	SelfEvaluatedLevel  string   `json:"self_evaluated_level"`  // 自评技术水平
	SystemScore         int64    `json:"system_score"`          // 系统评分
	PersonalityTags     []string `json:"personality_tags"`      // 性格标签
	PlayStyleTags       []string `json:"play_style_tags"`       // 打球风格
	PreferredSkillLevel string   `json:"preferred_skill_level"` // 希望对手技术水平
	PreferredTimeSlots  []string `json:"preferred_time_slots"`  // 时间偏好
	PreferredRegions    []string `json:"preferred_regions"`     // 常活动区域
	MaxCost             int64    `json:"max_cost"`              // 可接受花销
	CreateTime          string   `json:"create_time"`           // 创建时间
	UpdateTime          string   `json:"update_time"`           // 更新时间
}
