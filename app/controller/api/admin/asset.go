package admin

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"gincmf/app/model"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
	cmfUtil "github.com/gincmf/cmf/util"
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

// AssetsController 图片资源控制器，定义了资源文件增删改查接口
type AssetController struct {
	rc controller.RestController
}

type Size interface {
	Size() int64
}

var context *gin.Context

var FilesData []string

func (rest *AssetController) Get(c *gin.Context) {

	var assets []model.Asset

	query := "status = ?"
	queryArgs := []interface{}{"1"}

	paramType := c.DefaultPostForm("type","0")

	query += " AND type = ?"
	queryArgs = append(queryArgs,paramType)


	current := c.DefaultQuery("current","1")
	pageSize := c.DefaultQuery("pageSize","10")

	intCurrent ,_ := strconv.Atoi(current)
	intPageSize,_ := strconv.Atoi(pageSize)

	if intCurrent <= 0 {
		rest.rc.Error(c,"当前页码需大于0！",nil)
	}

	if intPageSize <= 0 {
		rest.rc.Error(c,"每页数需大于0！",nil)
	}

	total := 0
	cmf.Db.Limit(pageSize).Where(query, queryArgs...).Find(&assets).Count(&total)
	notFount := cmf.Db.Limit(intPageSize).Where(query, queryArgs...).Offset((intCurrent - 1) * intPageSize).Order("id desc").Find(&assets).RecordNotFound()

	if notFount {
		rest.rc.Error(c,"该页码内容不存在！",assets)
	}

	paginationData := &model.Paginate{Data: &assets, Current: current, PageSize: pageSize, Total: total}
	rest.rc.Success(c, "轮播图列表获取成功！", paginationData)
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

	form, _ := c.MultipartForm()
	files := form.File["file[]"]

	fileType := c.DefaultPostForm("type","0")

	if len(files) <= 0 {
		rest.rc.Error(c, "图片不能为空！",nil)
		return
	}
	context = c

	FilesData = FilesData[0:0]

	errCount := 0

	for _, fileItem := range files{
		err := handleUpload(fileItem,fileType)
		if err != nil {
			errCount++
		}
	}

	msg := "上传成功！"

	if errCount > 0 {
		msg = msg + "其中忽略" + strconv.Itoa(errCount) + "条！"
	}

	rest.rc.Success(c, msg,FilesData)
}

func (rest *AssetController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}

// 根据文件处理上传逻辑
// 1.判断上传类型，验证后缀合理性 type [0 => "图片" 1 => "视频" 2 => "文件"]
func handleUpload(file *multipart.FileHeader,fileType string)  error{
	tempFile,tempErr := file.Open()
	defer tempFile.Close()
	
	if tempErr != nil {
		fmt.Println("tempErr",tempErr)
	}

	var fileSize int64 = 0

	if sizeInterface, ok := tempFile.(Size); ok {
		fileSize = sizeInterface.Size()
	}

	fmt.Println("fileSize",fileSize)

	suffixArr := strings.Split(file.Filename, ".")

	suffix := suffixArr[len(suffixArr)-1]

	fmt.Println("suffix", suffix)
	
	//获取后缀列表
	uploadSetting := util.UploadSetting(context)
	switch fileType {
		case "0":
			if !strings.Contains(uploadSetting.Image.Extensions,suffix) {
				return errors.New("【"+ suffix + "】不是合法的图片后缀！")
				fmt.Println("0000",suffix)
			}
		case "1":
			if !strings.Contains(uploadSetting.Audio.Extensions,suffix) {
				return errors.New("【"+ suffix + "】不是合法的音频后缀！")
			}
		case "2":
			if !strings.Contains(uploadSetting.Video.Extensions,suffix) {
				return errors.New("【"+ suffix + "】不是合法的视频后缀！")
			}
		case "3":
			if !strings.Contains(uploadSetting.File.Extensions,suffix) {
				return errors.New("【"+ suffix + "】不是合法的附件后缀！")
			}
	}

	path := "public/uploads"
	t := time.Now()
	timeArr := []int{t.Year(), int(t.Month()), t.Day() }

	var timeDir string
	for key,timeInt := range timeArr{

		current := strconv.Itoa(timeInt)
		if key > 0 {
			if len(current) <= 1 {
				current = "0" + current
			}
		}
		tempStr  :=  "/" + current
		timeDir += tempStr
	}

	path += timeDir + "/"

	fileUuid, err := uuid.NewV4()

	fileName := cmfUtil.GetMd5(fileUuid.String() + suffixArr[0])
	fileNameSuffix := fileName + "." +suffix
	FilesData = append(FilesData,fileNameSuffix)

	filePath := path + fileNameSuffix

	_, err = os.Stat(path)
	if err != nil {
		os.MkdirAll( path, os.ModePerm)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, tempFile); err != nil {
		fmt.Println(err)
	}

	md5h := md5.New()
	md5h.Write(buf.Bytes())

	fileMd5 :=  hex.EncodeToString(md5h.Sum([]byte("")))
	log.Println("md5", fileMd5)

	sha1h := sha1.New()
	sha1h.Write(buf.Bytes())

	fileSha1 := hex.EncodeToString(sha1h.Sum([]byte("")))
	log.Println("sha1", fileSha1)
	// 上传文件至指定目录

	context.SaveUploadedFile(file, filePath)

	userId,_ := context.Get("user_id")
	if userId == nil {
		userId = "0"
	}
	userIdInt , _ := strconv.Atoi(userId.(string))

	fileTypeInt,_ := strconv.Atoi(fileType)
	//保存到数据库
	cmf.Db.Create(&model.Asset{
		UserId:userIdInt,
		FileSize:fileSize,
		CreateTime:time.Now().Unix(),
		FileKey:fileUuid.String(),
		FileName:fileNameSuffix,
		FilePath:filePath,
		Suffix:suffix,
		AssetType:fileTypeInt,
	})

	return nil
}
