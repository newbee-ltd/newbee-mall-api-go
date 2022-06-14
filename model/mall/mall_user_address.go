package mall

import (
	"main.go/model/common"
)

// MallUserAddress 结构体
// 如果含有time.Time 请自行import time包
type MallUserAddress struct {
	AddressId     int             `json:"addressId" form:"addressId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId        int             `json:"userId" form:"userId" gorm:"column:user_id;comment:用户主键id;type:bigint"`
	UserName      string          `json:"userName" form:"userName" gorm:"column:user_name;comment:收货人姓名;type:varchar(30);"`
	UserPhone     string          `json:"userPhone" form:"userPhone" gorm:"column:user_phone;comment:收货人手机号;type:varchar(11);"`
	DefaultFlag   int             `json:"defaultFlag" form:"defaultFlag" gorm:"column:default_flag;comment:是否为默认 0-非默认 1-是默认;type:tinyint"`
	ProvinceName  string          `json:"provinceName" form:"provinceName" gorm:"column:province_name;comment:省;type:varchar(32);"`
	CityName      string          `json:"cityName" form:"cityName" gorm:"column:city_name;comment:城;type:varchar(32);"`
	RegionName    string          `json:"regionName" form:"regionName" gorm:"column:region_name;comment:区;type:varchar(32);"`
	DetailAddress string          `json:"detailAddress" form:"detailAddress" gorm:"column:detail_address;comment:收件详细地址(街道/楼宇/单元);type:varchar(64);"`
	IsDeleted     int             `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	CreateTime    common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:添加时间;type:datetime"`
	UpdateTime    common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"`
}

// TableName MallUserAddress 表名
func (MallUserAddress) TableName() string {
	return "tb_newbee_mall_user_address"
}
