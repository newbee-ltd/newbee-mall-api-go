package manage

import (
	"main.go/model/common"
)

// 轮播图
type MallCarousel struct {
	CarouselId   int             `json:"carouselId" form:"carouselId" gorm:"primarykey;AUTO_INCREMENT"`
	CarouselUrl  string          `json:"carouselUrl" form:"carouselUrl" gorm:"column:carousel_url;comment:轮播图;type:varchar(100);"`
	RedirectUrl  string          `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;comment:点击后的跳转地址(默认不跳转);type:varchar(100);"`
	CarouselRank int             `json:"carouselRank" form:"carouselRank" gorm:"column:carousel_rank;comment:排序值(字段越大越靠前);type:int"`
	IsDeleted    int             `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	CreateTime   common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
	CreateUser   int             `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建者id;type:int"`
	UpdateTime   common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"`
	UpdateUser   int             `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者id;type:int"`
}

// TableName MallCarousel 表名
func (MallCarousel) TableName() string {
	return "tb_newbee_mall_carousel"
}
