package router

import (
	"gincmf/app/controller/api/admin"
	"gincmf/app/middleware"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
)

//web路由初始化
func ApiListenRouter() {
	cmf.Rest("/", new(admin.IndexController))

	cmf.Get("/api/currentUser", middleware.ValidationBearerToken, middleware.ValidationAdmin, func(c *gin.Context) {
		var currentUser = util.CurrentUser(c)
		controller.RestController{}.Success(c, "获取成功", currentUser)
	})

	cmf.Rest("/settings", new(admin.SettingsController), middleware.ValidationBearerToken, middleware.ValidationAdmin)
	cmf.Rest("/assets", new(admin.AssetController), middleware.ValidationBearerToken, middleware.ValidationAdmin)
	cmf.Rest("/upload", new(admin.UploadController), middleware.ValidationBearerToken, middleware.ValidationAdmin)
	cmf.Rest("/slide", new(admin.SlideController), middleware.ValidationBearerToken, middleware.ValidationAdmin)
	cmf.Rest("/slide_item", new(admin.SlideItemController), middleware.ValidationBearerToken, middleware.ValidationAdmin)
}
