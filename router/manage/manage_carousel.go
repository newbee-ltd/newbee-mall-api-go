package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type ManageCarouselRouter struct {
}

func (r *ManageCarouselRouter) InitManageCarouselRouter(Router *gin.RouterGroup) {
	mallCarouselRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var mallCarouselApi = v1.ApiGroupApp.ManageApiGroup.ManageCarouselApi
	{
		mallCarouselRouter.POST("carousels", mallCarouselApi.CreateCarousel)   // 新建MallCarousel
		mallCarouselRouter.DELETE("carousels", mallCarouselApi.DeleteCarousel) // 删除MallCarousel
		mallCarouselRouter.PUT("carousels", mallCarouselApi.UpdateCarousel)    // 更新MallCarousel
		mallCarouselRouter.GET("carousels/:id", mallCarouselApi.FindCarousel)  // 根据ID获取轮播图
		mallCarouselRouter.GET("carousels", mallCarouselApi.GetCarouselList)   // 获取轮播图列表
	}
}
