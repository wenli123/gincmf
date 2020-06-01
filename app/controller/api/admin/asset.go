package admin

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
	"github.com/nu7hatch/gouuid"
	"io"
	"log"
	"os"
	"strings"
)

// AssetsController 图片资源控制器，定义了资源文件增删改查接口
type AssetController struct {
	rc controller.RestController
}

func (rest *AssetController) Get(c *gin.Context) {
	rest.rc.Success(c, "操作成功Get", nil)
}

func (rest *AssetController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *AssetController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *AssetController) Store(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		rest.rc.Error(c, "图片不能为空！", nil)
		return
	}

	tempFile,tempErr := file.Open()
	defer tempFile.Close()

	if tempErr != nil {
		fmt.Println("tempErr",tempErr)
	}

	path := "public/uploads/"

	filePath := path + file.Filename

	_, err = os.Stat(path)
	if err != nil {
		os.Mkdir(path, os.ModePerm)
	}

	suffixArr := strings.Split(file.Filename, ".")

	suffix := suffixArr[len(suffixArr)-1]

	fmt.Println("suffix", suffix)
	log.Println(file.Filename)

	id, err := uuid.NewV4()
	log.Println("uuid", id)

	md5h := md5.New()
	io.Copy(md5h, tempFile)
	fileMd5 :=  hex.EncodeToString(md5h.Sum([]byte("")))
	log.Println("md5", fileMd5)

	// 上传文件至指定目录

	c.SaveUploadedFile(file, filePath)

	rest.rc.Success(c, "上传成功", file.Filename)
}

func (rest *AssetController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}
