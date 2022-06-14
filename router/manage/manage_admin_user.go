package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type ManageAdminUserRouter struct {
}

func (r *ManageAdminUserRouter) InitManageAdminUserRouter(Router *gin.RouterGroup) {
	mallAdminUserRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	mallAdminUserWithoutRouter := Router.Group("v1")
	var mallAdminUserApi = v1.ApiGroupApp.ManageApiGroup.ManageAdminUserApi
	{
		mallAdminUserRouter.POST("createMallAdminUser", mallAdminUserApi.CreateAdminUser) // 新建MallAdminUser
		mallAdminUserRouter.PUT("adminUser/name", mallAdminUserApi.UpdateAdminUserName)   // 更新MallAdminUser
		mallAdminUserRouter.PUT("adminUser/password", mallAdminUserApi.UpdateAdminUserPassword)
		mallAdminUserRouter.GET("users", mallAdminUserApi.UserList)
		mallAdminUserRouter.PUT("users/:lockStatus", mallAdminUserApi.LockUser)
		mallAdminUserRouter.GET("adminUser/profile", mallAdminUserApi.AdminUserProfile) // 根据ID获取 admin详情
		mallAdminUserRouter.DELETE("logout", mallAdminUserApi.AdminLogout)
		mallAdminUserRouter.POST("upload/file", mallAdminUserApi.UploadFile) //上传图片

	}
	{
		mallAdminUserWithoutRouter.POST("adminUser/login", mallAdminUserApi.AdminLogin) //管理员登陆
	}
}
