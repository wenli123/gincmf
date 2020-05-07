package admin

import (
	"github.com/daifuyang/cmf/controller"
	"github.com/gin-gonic/gin"
)

type IndexControllerStruct struct {
	Rc controller.RestControllerStruct
}

func (ic *IndexControllerStruct) Get(c *gin.Context) {
	controller.Success(c, "操作成功Get", nil)
}

func (ic *IndexControllerStruct) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	controller.Success(c, "操作成功show", nil)
}

func (ic *IndexControllerStruct) Edit(c *gin.Context) {
	controller.Success(c, "操作成功Edit", nil)
}

func (ic *IndexControllerStruct) Store(c *gin.Context) {
	controller.Success(c, "操作成功Store", nil)
}

func (ic *IndexControllerStruct) Delete(c *gin.Context) {
	controller.Success(c, "操作成功Delete", nil)
}
