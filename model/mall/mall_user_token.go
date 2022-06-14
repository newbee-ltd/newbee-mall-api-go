package mall

import (
	"time"
)

// MallUserToken 结构体
// 如果含有time.Time 请自行import time包
type MallUserToken struct {
	UserId     int       `json:"userId" form:"userId" gorm:"primarykey;AUTO_INCREMENT"`
	Token      string    `json:"token" form:"token" gorm:"column:token;comment:token值(32位字符串);type:varchar(32);"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"`
	ExpireTime time.Time `json:"expireTime" form:"expireTime" gorm:"column:expire_time;comment:token过期时间;type:datetime"`
}

// TableName MallUserToken 表名
func (MallUserToken) TableName() string {
	return "tb_newbee_mall_user_token"
}
