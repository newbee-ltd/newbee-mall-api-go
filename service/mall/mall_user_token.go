package mall

import (
	"main.go/global"
	"main.go/model/mall"
)

type MallUserTokenService struct {
}

func (m *MallUserTokenService) ExistUserToken(token string) (err error, mallUserToken mall.MallUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallUserToken).Error
	return
}

func (m *MallUserTokenService) DeleteMallUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]mall.MallUserToken{}, "token =?", token).Error
	return err
}
