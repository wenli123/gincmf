package router

import (
	"gincmf/app/controller/api/admin"
	"github.com/gincmf/cmf"
)

//web路由初始化
func ApiListenRouter() {
	cmf.Rest("/", new(admin.IndexController))
	cmf.Rest("/settings", new(admin.SettingsController))
	cmf.Rest("/assets", new(admin.AssetController))
	cmf.Rest("/slide", new(admin.SlideController))
	cmf.Rest("/slide_item", new(admin.SlideItemController))
}
