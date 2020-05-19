package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
)

type IndexController struct {
	rc controller.RestController
}

func (i *IndexController) Get(c *gin.Context) {
	i.rc.Success(c, "操作成功Get", nil)
}

func (i *IndexController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	i.rc.Success(c, "操作成功show", nil)
}

func (i *IndexController) Edit(c *gin.Context) {
	i.rc.Success(c, "操作成功Edit", nil)
}

func (i *IndexController) Store(c *gin.Context) {
	i.rc.Success(c, "操作成功Store", nil)
}

func (i *IndexController) Delete(c *gin.Context) {
	i.rc.Success(c, "操作成功Delete", nil)
}