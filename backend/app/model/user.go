package model

import (
	"database/sql"
	"time"
	
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AccountStatus string



type User struct {
	ID                  uint64         `gorm:"primaryKey;autoIncrement;comment:用户编号"`
	Username            string         `gorm:"type:varchar(50);not null;uniqueIndex;comment:登录用户名（唯一）"`
	PasswordHash        string         `gorm:"type:varchar(255);not null;comment:加密后的密码"`
	Email               sql.NullString `gorm:"type:varchar(100);uniqueIndex;comment:邮箱（可选）"`
	Phone               sql.NullString `gorm:"type:varchar(20);uniqueIndex;comment:手机号（可选）"`
	AccountStatus       AccountStatus  `gorm:"type:ENUM('active', 'locked', 'disabled');default:'active';comment:账号状态"`
	LastLoginTime       sql.NullTime   `gorm:"comment:最后登录时间"`
	
	Nickname            sql.NullString `gorm:"type:varchar(50);comment:可选，仅供查看"`
	AvatarPath          string         `gorm:"type:varchar(255);default:'default-avatar.png';comment:头像路径(相对路径)"`
	SelfEvaluatedLevel  sql.NullString `gorm:"type:varchar(20);comment:自评技术水平"`
	SystemScore         sql.NullInt32  `gorm:"comment:系统评估得分(0~100)"`
	PersonalityTags     datatypes.JSON `gorm:"comment:性格标签"`
	PlayStyleTags       datatypes.JSON `gorm:"comment:打球风格"`
	PreferredSkillLevel sql.NullString `gorm:"type:varchar(20);comment:希望对手技术水平"`
	PreferredTimeSlots  datatypes.JSON `gorm:"comment:时间偏好"`
	PreferredRegions    datatypes.JSON `gorm:"comment:常活动区域"`
	MaxCost             sql.NullInt32  `gorm:"comment:可接受的花销（单位：元）"`
	HistoricalPartners  datatypes.JSON `gorm:"comment:历史搭档ID列表"`
	RatingsGiven        datatypes.JSON `gorm:"comment:对别人的评价(用户ID:评分)"`
	CreateTime          time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime          time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
}

// TableName 设置表名
func (User) TableName() string {
	return "user"
}

// BeforeCreate 钩子 - 创建前设置默认值
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.AvatarPath == "" {
		u.AvatarPath = "default-avatar.png"
	}
	return nil
}

// CheckConstraints 自定义检查约束
func (u *User) CheckConstraints() error {
	if u.SystemScore.Valid && (u.SystemScore.Int32 < 0 || u.SystemScore.Int32 > 100) {
		return gorm.ErrInvalidValue
	}
	if u.MaxCost.Valid && u.MaxCost.Int32 < 0 {
		return gorm.ErrInvalidValue
	}
	return nil
}

// BeforeSave 钩子 - 保存前验证
func (u *User) BeforeSave(tx *gorm.DB) error {
	return u.CheckConstraints()
}

// BeforeUpdate 钩子 - 更新前验证
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.CheckConstraints()
}