package main

import (
	"gincmf/router"
	"github.com/gincmf/cmf"
)

func main() {

	//初始化配置设置
	cmf.Initialize("./conf/config.json")
	//初始化路由设置
	router.WebListenRouter()
	router.ApiListenRouter()
	//启动服务
	cmf.Start()
}
