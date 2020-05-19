package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
	"log"
)

// AssetsController 图片资源控制器，定义了资源文件增删改查接口
type AssetsController struct {
	rc controller.RestController
}

func (rest *AssetsController) Get(c *gin.Context) {
	rest.rc.Success(c, "操作成功Get", nil)
}

func (rest *AssetsController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *AssetsController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *AssetsController) Store(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		rest.rc.Error(c,"图片不能为空！")
		return
	}

	fmt.Println("file")

	log.Println(file.Filename)
	// 上传文件至指定目录

	c.SaveUploadedFile(file,"public/uploads/" + file.Filename )

	rest.rc.Success(c, "上传成功", file.Filename)
}

func (rest *AssetsController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}
