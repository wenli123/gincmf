package home

import (
	"gincmf/app/controller/web"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/view"
)

type IndexController struct {
	c *web.ControllerStruct
}

//首页控制器
func (this *IndexController) Index(c *gin.Context) {
	view.Fetch("index.html")
}
