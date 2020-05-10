package main

import (
	"gincmf/app/migrate"
	"gincmf/router"
	"github.com/gincmf/cmf"
)

func main() {
	//初始化配置设置
	cmf.Initialize("./conf/config.json")
	//初始化路由设置
	router.WebListenRouter()
	router.ApiListenRouter()
	migrate.AutoMigrate()
	//启动服务
	cmf.Start()
	defer cmf.Db.Close()
	//执行数据库迁移
}
