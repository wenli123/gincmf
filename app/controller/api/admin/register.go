package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
)

//注册账号
type RegisterController struct {
	rc controller.RestController
}

func (s *RegisterController) Get(c *gin.Context) {
	s.rc.Success(c, "test方法操作成功Get", nil)
}

func (s *RegisterController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	s.rc.Success(c, "操作成功show", nil)
}

func (s *RegisterController) Edit(c *gin.Context) {
	s.rc.Success(c, "操作成功Edit", nil)
}

func (s *RegisterController) Store(c *gin.Context) {
	s.rc.Success(c, "操作成功Store", nil)
}

func (s *RegisterController) Delete(c *gin.Context) {
	s.rc.Success(c, "操作成功Delete", nil)
}
