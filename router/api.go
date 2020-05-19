package router

import (
	"gincmf/app/controller/api/admin"
	"gincmf/app/middleware"
	"github.com/gincmf/cmf"
)

//web路由初始化
func ApiListenRouter() {
	cmf.Rest("/", new(admin.IndexController))
	cmf.Rest("/assets", new(admin.AssetsController),middleware.ValidationBearerToken,middleware.ValidationAdmin)
	cmf.Rest("/carousel", new(admin.SlideController),middleware.ValidationBearerToken,middleware.ValidationAdmin)
}
