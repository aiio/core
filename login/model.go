package login

import (
	"gorm.io/gorm"
	"time"
)

// AuthUser 角色
type AuthUser struct {
	ID      uint   `gorm:"primarykey" json:"-"`
	UID     string `gorm:"type:char(32);uniqueIndex" json:"uid" validate:"required"`
	GroupId uint   `gorm:"type:bigint(20)" json:"group_id" validate:"required"`
}

// AuthGroup 管理组表
type AuthGroup struct {
	ID     uint   `gorm:"primarykey" json:"-"`
	Name   string `gorm:"type:varchar(128)" json:"name" validate:"required"`
	Rules  string `gorm:"type:varchar(128)" json:"rules" validate:"required"`
	Status int    `gorm:"type:tinyint(1);not null;default:0" json:"status"` // 状态 0暂存，1使用
}

// AuthRule 许可 权限
type AuthRule struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Pid    uint   `gorm:"type:bigint(20)" json:"pid"`
	Name   string `gorm:"type:varchar(128)" json:"name" validate:"required"`
	Host   string `gorm:"type:char(32);uniqueIndex:humudx" json:"host" validate:"required"`
	Uri    string `gorm:"type:char(32);uniqueIndex:humudx" json:"uri" validate:"required"`
	Method string `gorm:"type:char(32);uniqueIndex:humudx" json:"method" validate:"required"`
	Status int    `gorm:"type:tinyint(1);not null;default:0" json:"status"` // 状态 0暂存，1使用
}

// UserAuth 用户授权表
type UserAuth struct {
	ID          uint           `gorm:"primarykey" json:"-"`
	UID         string         `gorm:"type:char(32);uniqueIndex" json:"uid" validate:"required"`                          // 用户ID
	Scenes      string         `gorm:"type:char(32);uniqueIndex:idx_scenes_identity" json:"type" validate:"required"`     // 身份标识类型
	Identity    string         `gorm:"type:char(32);uniqueIndex:idx_scenes_identity" json:"identity" validate:"required"` // 身份标识 上述对应的值
	Certificate string         `gorm:"type:char(64)" json:"certificate" validate:"required"`                              // 密码
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// UserBase 用户基础资料表
type UserBase struct {
	ID        uint           `gorm:"primarykey" json:"-"`
	UID       string         `gorm:"type:char(36);uniqueIndex" json:"uid" validate:"required"` // 用户UID
	Nickname  string         `gorm:"type:char(16)" json:"nickname" validate:"required"`        // 用户昵称
	Avatar    string         `gorm:"type:char(128)" json:"avatar"`                             // 头像
	AuthName  string         `gorm:"type:char(16)" json:"auth_name"`                           // 授权名称
	Wechat    string         `gorm:"type:char(32);uniqueIndex" json:"wechat"`                  // 微信号
	Phone     string         `gorm:"type:char(16);uniqueIndex" json:"phone"`                   // 手机号
	Address   string         `gorm:"type:char(128)" json:"address"`                            // 地址
	Gender    int            `gorm:"type:tinyint(1);not null;default:0" json:"gender"`         // 状态 0未知/保密，1男，2女
	Status    int            `gorm:"type:tinyint(1);not null;default:0" json:"status"`         // 状态 0暂存，1使用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
