package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
)

type MallGoodsCategoryApi struct {
}

//返回分类数据(首页调用)
func (m *MallGoodsCategoryApi) GetGoodsCategory(c *gin.Context) {
	err, list := mallGoodsCategoryService.GetCategoriesForIndex()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	}
	response.OkWithData(list, c)
}
