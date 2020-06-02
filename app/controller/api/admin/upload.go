package admin

import (
	"fmt"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
)

// AssetsController 图片资源控制器，定义了资源文件增删改查接口
type UploadController struct {
	rc controller.RestController
}

func (rest *UploadController) Get(c *gin.Context) {
	uploadSetting := util.UploadSetting(c)
	rest.rc.Success(c, "获取成功！", uploadSetting)
}

func (rest *UploadController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *UploadController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *UploadController) Store(c *gin.Context) {

	c.Request.ParseForm()
	formData := c.Request.PostForm
	fmt.Println("formData",formData)

	chunkSiseMap , ok := formData["chunk_size"]
	chunkSise := ""
	if ok {
		chunkSise = chunkSiseMap[0]
	}

	if ok {
		chunkSise = chunkSiseMap[0]
	}

	fmt.Println("chunkSise",chunkSise)


	rest.rc.Success(c, "修改成功", nil)
}

func (rest *UploadController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}
