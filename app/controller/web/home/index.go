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
	web.Fetch("index.html")
}
