package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	"strconv"
)

type MallGoodsInfoApi struct {
}

// 商品搜索
func (m *MallGoodsInfoApi) GoodsSearch(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	goodsCategoryId, _ := strconv.Atoi(c.Query("goodsCategoryId"))
	keyword := c.Query("keyword")
	orderBy := c.Query("orderBy")
	if err, list, total := mallGoodsInfoService.MallGoodsListBySearch(pageNumber, goodsCategoryId, keyword, orderBy); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   10,
		}, "获取成功", c)
	}
}

func (m *MallGoodsInfoApi) GoodsDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsInfo := mallGoodsInfoService.GetMallGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	}
	response.OkWithData(goodsInfo, c)
}
