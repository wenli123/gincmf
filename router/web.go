package router

import (
	"gincmf/app/controller/web/home"
	"gincmf/app/middleware"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf"
	"github.com/gincmf/cmf/controller"
)

//web路由初始化

func WebListenRouter() {
	cmf.Get("/", middleware.BaseController, new(home.IndexController).Index)

	cmf.Get("/api/currentUser",middleware.ValidationBearerToken,middleware.ValidationAdmin, func(c *gin.Context) {
		var currentUser = util.CurrentUser(c)
		controller.RestController{}.Success(c,"获取成功",currentUser)

	} )

}
