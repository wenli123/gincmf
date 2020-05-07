package router

import (
	"gincmf/app/controller/web/home"
	"gincmf/app/middleware"
	"github.com/gincmf/cmf"
)

//web路由初始化

func WebListenRouter() {
	cmf.Get("/", middleware.BaseController, new(home.IndexController).Index)
}
