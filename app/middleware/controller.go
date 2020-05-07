package middleware

import (
	"gincmf/app/controller/web"
	"github.com/daifuyang/cmf"
	"github.com/gin-gonic/gin"
)

func BaseController(c *gin.Context) {
	cmf.LoadTemplate()
	web.Template.Context = c
	web.Assign("tmpl",cmf.TemplateMap.Theme)
	c.Next()
}
