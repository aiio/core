package modelx

import (
	"time"

	"gorm.io/gorm"
)

// ConfigGroup 配置分组
type ConfigGroup struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Owner     string         `gorm:"type:varchar(128);default:'';comment:所属者" json:"owner"`
	Name      string         `gorm:"type:varchar(128);default:'';comment:分组名称" json:"name"`
	Tag       string         `gorm:"type:char(32);default:'';uniqueIndex;comment:分组唯一tag" json:"tag"`
	Remark    string         `gorm:"type:varchar(128);default:'';comment:备注" json:"remark"`
	Sort      int            `gorm:"type:int(10);default:0;comment:排序" json:"sort"`
	Status    int            `gorm:"type:tinyint(4);default:0;comment:状态：0禁用，1启用" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Config 配置信息
type Config struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Owner     string         `gorm:"type:varchar(128);default:'';comment:所属者" json:"owner"`
	GroupTag  string         `gorm:"type:char(32);default:'';comment:分组标签" json:"group_tag"`
	Tag       string         `gorm:"type:char(32);default:'';uniqueIndex;comment:分组唯一tag" json:"tag"`
	Value     string         `gorm:"type:varchar(1024);default:'';comment:值" json:"value"`
	Remark    string         `gorm:"type:varchar(1024);default:'';comment:备注" json:"remark"`
	Sort      int            `gorm:"type:int(10);default:0;comment:排序" json:"sort"`
	Status    int            `gorm:"type:tinyint(4);default:0;comment:状态：0禁用，1启用" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
