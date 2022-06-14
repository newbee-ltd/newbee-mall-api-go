package example

import "main.go/model/common"

type ExaFileUploadAndDownload struct {
	ID         int             `gorm:"primarykey"`                                                      // 主键ID
	Name       string          `json:"name" gorm:"comment:文件名"`                                         // 文件名
	Url        string          `json:"url" gorm:"comment:文件地址"`                                         // 文件地址
	Tag        string          `json:"tag" gorm:"comment:文件标签"`                                         // 文件标签
	Key        string          `json:"key" gorm:"comment:编号"`                                           // 编号
	CreateTime common.JSONTime `json:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"` // 创建时间
	UpdateTime common.JSONTime `json:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"` // 更新时间
}
