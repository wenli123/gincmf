package router

import (
	"github.com/gincmf/cmf"
	"gincmf/app/controller/api/admin"
)

//web路由初始化
func ApiListenRouter() {
	cmf.Rest("/", new(admin.IndexControllerStruct))
}
