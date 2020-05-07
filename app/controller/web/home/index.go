package home

import (
	"gincmf/app/controller/web"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	c *web.ControllerStruct
}

//首页控制器
func (this *IndexController) Index(c *gin.Context) {
	web.Assign("hello","Hello World")
	web.Assign("world","my name is dai fu yang")
	web.Fetch("index.html")
}
