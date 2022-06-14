package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"strconv"
)

type ManageGoodsInfoApi struct {
}

func (m *ManageGoodsInfoApi) CreateGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manageReq.GoodsInfoAddParam
	_ = c.ShouldBindJSON(&mallGoodsInfo)
	if err := mallGoodsInfoService.CreateMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败!"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMallGoodsInfo 删除MallGoodsInfo
func (m *ManageGoodsInfoApi) DeleteGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manage.MallGoodsInfo
	_ = c.ShouldBindJSON(&mallGoodsInfo)
	if err := mallGoodsInfoService.DeleteMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// ChangeMallGoodsInfoByIds 批量删除MallGoodsInfo
func (m *ManageGoodsInfoApi) ChangeGoodsInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	sellStatus := c.Param("status")
	if err := mallGoodsInfoService.ChangeMallGoodsInfoByIds(IDS, sellStatus); err != nil {
		global.GVA_LOG.Error("修改商品状态失败!", zap.Error(err))
		response.FailWithMessage("修改商品状态失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改商品状态成功", c)
	}
}

// UpdateMallGoodsInfo 更新MallGoodsInfo
func (m *ManageGoodsInfoApi) UpdateGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manageReq.GoodsInfoUpdateParam
	_ = c.ShouldBindJSON(&mallGoodsInfo)
	if err := mallGoodsInfoService.UpdateMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMallGoodsInfo 用id查询MallGoodsInfo
func (m *ManageGoodsInfoApi) FindGoodsInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsInfo := mallGoodsInfoService.GetMallGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	}
	goodsInfoRes := make(map[string]interface{})
	goodsInfoRes["goods"] = goodsInfo
	if _, thirdCategory := mallGoodsCategoryService.SelectCategoryById(goodsInfo.GoodsCategoryId); thirdCategory != (manage.MallGoodsCategory{}) {
		goodsInfoRes["thirdCategory"] = thirdCategory
		if _, secondCategory := mallGoodsCategoryService.SelectCategoryById(thirdCategory.ParentId); secondCategory != (manage.MallGoodsCategory{}) {
			goodsInfoRes["secondCategory"] = secondCategory
			if _, firstCategory := mallGoodsCategoryService.SelectCategoryById(secondCategory.ParentId); firstCategory != (manage.MallGoodsCategory{}) {
				goodsInfoRes["firstCategory"] = firstCategory
			}
		}
	}
	response.OkWithData(goodsInfoRes, c)

}

// GetMallGoodsInfoList 分页获取MallGoodsInfo列表
func (m *ManageGoodsInfoApi) GetGoodsInfoList(c *gin.Context) {
	var pageInfo manageReq.MallGoodsInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	goodsName := c.Query("goodsName")
	goodsSellStatus := c.Query("goodsSellStatus")
	if err, list, total := mallGoodsInfoService.GetMallGoodsInfoInfoList(pageInfo, goodsName, goodsSellStatus); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
