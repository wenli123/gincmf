package util

import (
	"encoding/json"
	"fmt"
	"gincmf/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
)

//获取当前登录管理员id
func CurrentAdminId(c *gin.Context) string{
	userId,_ := c.Get("user_id")
	return userId.(string)
}

//获取当前用户信息
func CurrentUser(c *gin.Context) model.User {
	session := sessions.Default(c)
	u := &model.User{}
	currentUser := session.Get("current_user")
	//根据userId获取当前用户信息，不存在写入session
	if currentUser == nil {
		userId,_ := c.Get("user_id")
		result := cmf.Db.First(u,"user_login = ?", userId).RecordNotFound()
		if !result {
			session.Set("current_user",*u)
		}
		currentUser = u
	}
	return currentUser.(model.User)
}

//获取网站上传配置信息
func UploadSetting(c *gin.Context) model.UploadSetting{
	session := sessions.Default(c)
	uploadSettingStr := session.Get("uploadSetting")

	fmt.Println("uploadSetting Session",uploadSettingStr)

	option := &model.Option{}

	uploadSetting := model.UploadSetting{}


	if uploadSettingStr == nil {
		uploadResult := cmf.Db.First(option, "option_name = ?", "upload_setting") // 查询
		if !uploadResult.RecordNotFound() {
			uploadSettingStr = option.OptionValue

			fmt.Println("uploadSetting",uploadSettingStr)
			fmt.Printf(`%T`, uploadSettingStr)
			//存入session
			session.Set(uploadSetting,uploadSettingStr)
		}
	}

	//读取的数据为json格式，需要进行解码
	json.Unmarshal([]byte(uploadSettingStr.(string)), uploadSetting)

	return uploadSetting
}