package manage

import (
	"main.go/model/common"
)

// MallUser 商城用户信息
type MallUser struct {
	UserId        int             `json:"userId" form:"userId" gorm:"primarykey;AUTO_INCREMENT"`
	NickName      string          `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;type:varchar(50);"`
	LoginName     string          `json:"loginName" form:"loginName" gorm:"column:login_name;comment:登陆名称(默认为手机号);type:varchar(11);"`
	PasswordMd5   string          `json:"passwordMd5" form:"passwordMd5" gorm:"column:password_md5;comment:MD5加密后的密码;type:varchar(32);"`
	IntroduceSign string          `json:"introduceSign" form:"introduceSign" gorm:"column:introduce_sign;comment:个性签名;type:varchar(100);"`
	IsDeleted     int             `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:注销标识字段(0-正常 1-已注销);type:tinyint"`
	LockedFlag    int             `json:"lockedFlag" form:"lockedFlag" gorm:"column:locked_flag;comment:锁定标识字段(0-未锁定 1-已锁定);type:tinyint"`
	CreateTime    common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:注册时间;type:datetime"`
}

// TableName MallUser 表名
func (MallUser) TableName() string {
	return "tb_newbee_mall_user"
}
