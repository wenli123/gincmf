package admin

import (
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
)

type IndexController struct {
	rc controller.RestController
}

func (s *IndexController) Get(c *gin.Context) {

	util.SetLog(c,"admin","index","get","查看了首页列表！")
	s.rc.Success(c, "操作成功Get", nil)
}

func (s *IndexController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	s.rc.Success(c, "操作成功show", nil)
}

func (s *IndexController) Edit(c *gin.Context) {
	s.rc.Success(c, "操作成功Edit", nil)
}

func (s *IndexController) Store(c *gin.Context) {
	s.rc.Success(c, "操作成功Store", nil)
}

func (s *IndexController) Delete(c *gin.Context) {
	s.rc.Success(c, "操作成功Delete", nil)
}