package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"strconv"
)

type MallUserAddressApi struct {
}

func (m *MallUserAddressApi) AddressList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddressList := mallUserAddressService.GetMyAddress(token); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		response.FailWithMessage("获取地址失败:"+err.Error(), c)
	} else if len(userAddressList) == 0 {
		response.OkWithData(nil, c)
	} else {
		response.OkWithData(userAddressList, c)
	}
}

func (m *MallUserAddressApi) SaveUserAddress(c *gin.Context) {
	var req mallReq.AddAddressParam
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	err := mallUserAddressService.SaveUserAddress(token, req)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
	}
	response.OkWithMessage("创建成功", c)

}

func (m *MallUserAddressApi) UpdateMallUserAddress(c *gin.Context) {
	var req mallReq.UpdateAddressParam
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	err := mallUserAddressService.UpdateUserAddress(token, req)
	if err != nil {
		global.GVA_LOG.Error("更新用户地址失败", zap.Error(err))
		response.FailWithMessage("更新用户地址失败:"+err.Error(), c)
	}
	response.OkWithMessage("更新用户地址成功", c)
}

func (m *MallUserAddressApi) GetMallUserAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("addressId"))
	token := c.GetHeader("token")
	if err, userAddress := mallUserAddressService.GetMallUserAddressById(token, id); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		response.FailWithMessage("获取地址失败:"+err.Error(), c)
	} else {
		response.OkWithData(userAddress, c)
	}
}

func (m *MallUserAddressApi) GetMallUserDefaultAddress(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddress := mallUserAddressService.GetMallUserDefaultAddress(token); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		response.FailWithMessage("获取地址失败:"+err.Error(), c)
	} else {
		response.OkWithData(userAddress, c)
	}
}

func (m *MallUserAddressApi) DeleteUserAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("addressId"))
	token := c.GetHeader("token")
	err := mallUserAddressService.DeleteUserAddress(token, id)
	if err != nil {
		global.GVA_LOG.Error("删除用户地址失败", zap.Error(err))
		response.FailWithMessage("删除用户地址失败:"+err.Error(), c)
	}
	response.OkWithMessage("删除用户地址成功", c)

}
