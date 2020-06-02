package middleware

import (
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/view"
)

func BaseController(c *gin.Context) {
	cmf.LoadTemplate()
	view.Template.Context = c
	view.Assign("tmpl", cmf.TemplateMap.ThemePath + "/" + cmf.TemplateMap.Theme) //静态资源路径
	c.Next()
}
