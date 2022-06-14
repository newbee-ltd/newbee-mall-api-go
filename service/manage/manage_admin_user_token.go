package manage

import (
	"main.go/global"
	"main.go/model/manage"
)

type ManageAdminUserTokenService struct {
}

func (m *ManageAdminUserTokenService) ExistAdminToken(token string) (err error, mallAdminUserToken manage.MallAdminUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallAdminUserToken).Error
	return
}

func (m *ManageAdminUserTokenService) DeleteMallAdminUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]manage.MallAdminUserToken{}, "token =?", token).Error
	return err
}
