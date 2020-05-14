package admin

import (
	"github.com/gincmf/cmf/controller"
	"github.com/gin-gonic/gin"
)

type TestControllerStruct struct {
	rc controller.RestControllerStruct
}

func (i *TestControllerStruct) Get(c *gin.Context) {
	i.rc.Success(c, "操作成功Get", nil)
}

func (i *TestControllerStruct) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	i.rc.Success(c, "操作成功show", nil)
}

func (i *TestControllerStruct) Edit(c *gin.Context) {
	i.rc.Success(c, "操作成功Edit", nil)
}

func (i *TestControllerStruct) Store(c *gin.Context) {
	i.rc.Success(c, "操作成功Store", nil)
}

func (i *TestControllerStruct) Delete(c *gin.Context) {
	i.rc.Success(c, "操作成功Delete", nil)
}
