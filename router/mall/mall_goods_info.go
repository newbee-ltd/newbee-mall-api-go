package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallGoodsInfoIndexRouter struct {
}

func (m *MallGoodsInfoIndexRouter) InitMallGoodsInfoIndexRouter(Router *gin.RouterGroup) {
	mallGoodsRouter := Router.Group("v1")
	var mallGoodsInfoApi = v1.ApiGroupApp.MallApiGroup.MallGoodsInfoApi
	{
		mallGoodsRouter.GET("/search", mallGoodsInfoApi.GoodsSearch)           // 商品搜索
		mallGoodsRouter.GET("/goods/detail/:id", mallGoodsInfoApi.GoodsDetail) //商品详情
	}
}
