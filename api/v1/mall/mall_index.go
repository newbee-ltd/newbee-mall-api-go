package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/enum"
	"main.go/model/common/response"
)

type MallIndexApi struct {
}

// MallIndexInfo 加载首页信息
func (m *MallIndexApi) MallIndexInfo(c *gin.Context) {
	err, _, mallCarouseInfo := mallCarouselService.GetCarouselsForIndex(5)
	if err != nil {
		global.GVA_LOG.Error("轮播图获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("轮播图获取失败", c)
	}
	err, hotGoodses := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsHot.Code(), 4)
	if err != nil {
		global.GVA_LOG.Error("热销商品获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("热销商品获取失败", c)
	}
	err, newGoodses := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsNew.Code(), 5)
	if err != nil {
		global.GVA_LOG.Error("新品获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("新品获取失败", c)
	}
	err, recommendGoodses := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsRecommond.Code(), 10)
	if err != nil {
		global.GVA_LOG.Error("推荐商品获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("推荐商品获取失败", c)
	}
	indexResult := make(map[string]interface{})
	indexResult["carousels"] = mallCarouseInfo
	indexResult["hotGoodses"] = hotGoodses
	indexResult["newGoodses"] = newGoodses
	indexResult["recommendGoodses"] = recommendGoodses
	response.OkWithData(indexResult, c)

}
