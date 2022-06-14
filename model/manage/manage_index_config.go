// 自动生成模板MallIndexConfig
package manage

import (
	"main.go/model/common"
)

// MallIndexConfig 结构体
// 如果含有time.Time 请自行import time包
type MallIndexConfig struct {
	ConfigId    int             `json:"configId" form:"configId" gorm:"primarykey;AUTO_INCREMENT"`
	ConfigName  string          `json:"configName" form:"configName" gorm:"column:config_name;comment:显示字符(配置搜索时不可为空，其他可为空);type:varchar(50);"`
	ConfigType  int             `json:"configType" form:"configType" gorm:"column:config_type;comment:1-搜索框热搜 2-搜索下拉框热搜 3-(首页)热销商品 4-(首页)新品上线 5-(首页)为你推荐;type:tinyint"`
	GoodsId     int             `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id 默认为0;type:bigint"`
	RedirectUrl string          `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;comment:点击后的跳转地址(默认不跳转);type:varchar(100);"`
	ConfigRank  int             `json:"configRank" form:"configRank" gorm:"column:config_rank;comment:排序值(字段越大越靠前);type:int"`
	IsDeleted   int             `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	CreateTime  common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
	CreateUser  int             `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建者id;type:int"`
	UpdateTime  common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:最新修改时间;type:datetime"`
	UpdateUser  int             `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者id;type:int"`
}

// TableName MallIndexConfig 表名
func (MallIndexConfig) TableName() string {
	return "tb_newbee_mall_index_config"
}
