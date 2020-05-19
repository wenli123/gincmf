package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
)

//注册账号
type RegisterController struct {
	rc controller.RestController
}

func (i *RegisterController) Get(c *gin.Context) {
	i.rc.Success(c, "test方法操作成功Get", nil)
}

func (i *RegisterController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	i.rc.Success(c, "操作成功show", nil)
}

func (i *RegisterController) Edit(c *gin.Context) {
	i.rc.Success(c, "操作成功Edit", nil)
}

func (i *RegisterController) Store(c *gin.Context) {
	i.rc.Success(c, "操作成功Store", nil)
}

func (i *RegisterController) Delete(c *gin.Context) {
	i.rc.Success(c, "操作成功Delete", nil)
}
