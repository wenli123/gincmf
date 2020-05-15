package admin

import (
	"fmt"
	"gincmf/app/middleware"
	"github.com/gincmf/cmf/controller"
	"github.com/gin-gonic/gin"
	"time"
)

type TestControllerStruct struct {
	rc controller.RestControllerStruct
}

func (i *TestControllerStruct) Get(c *gin.Context) {

	token := middleware.Token

	var data struct{
		ExpiresIn int64 `json:"expires_in"`
		ClientId string `json:"client_id"`
		UserId string `json:"user_id"`
	}

	data.ExpiresIn = int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds())
	data.ClientId = token.GetClientID()
	data.UserId = token.GetUserID()

	fmt.Println("data：",data)

	i.rc.Success(c, "test方法操作成功Get", data)
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