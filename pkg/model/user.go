package model

import (
	"github.com/skyfour/go-model/pkg/utils"
)

// UserPortfolios - user's portfolio information
type Users struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Nickname string `json:"nickname"`
	Timestamp
}

// 默认表名为type 后面定义的结构体名，这样可以自己按照需求生成表名
func (Users) TableName() string {
	return "users1"
}

//初始化，可以直接生成表
func init() {
	utils.DB.AutoMigrate(&Users{})
}
