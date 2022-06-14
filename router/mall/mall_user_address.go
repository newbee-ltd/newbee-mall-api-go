package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type MallUserAddressRouter struct {
}

func (m *MallUserRouter) InitMallUserAddressRouter(Router *gin.RouterGroup) {
	mallUserAddressRouter := Router.Group("v1").Use(middleware.UserJWTAuth())
	var mallUserAddressApi = v1.ApiGroupApp.MallApiGroup.MallUserAddressApi
	{
		mallUserAddressRouter.GET("/address", mallUserAddressApi.AddressList)                       //用户地址
		mallUserAddressRouter.POST("/address", mallUserAddressApi.SaveUserAddress)                  //添加地址
		mallUserAddressRouter.PUT("/address", mallUserAddressApi.UpdateMallUserAddress)             //修改用户地址
		mallUserAddressRouter.GET("/address/:addressId", mallUserAddressApi.GetMallUserAddress)     //获取地址详情
		mallUserAddressRouter.GET("/address/default", mallUserAddressApi.GetMallUserDefaultAddress) //获取默认地址
		mallUserAddressRouter.DELETE("/address/:addressId", mallUserAddressApi.DeleteUserAddress)   //删除地址
	}

}
