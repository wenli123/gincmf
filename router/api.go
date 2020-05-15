package router

import (
	"gincmf/app/controller/api/admin"
	"gincmf/app/middleware"
	"github.com/gincmf/cmf"
)

//web路由初始化
func ApiListenRouter() {
	cmf.Rest("/", new(admin.IndexControllerStruct))
	cmf.Rest("/test", new(admin.TestControllerStruct),middleware.ValidationBearerToken)
}
