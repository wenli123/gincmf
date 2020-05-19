package middleware

import (
	"fmt"
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
	"github.com/gincmf/cmf/router"
	"gopkg.in/oauth2.v3"
)

var Token oauth2.TokenInfo

//验证oauth token值
func ValidationBearerToken(c *gin.Context) {
	s := router.Srv
	t, err := s.ValidationBearerToken(c.Request)
	Token = t
	if err != nil {
		controller.RestController{}.Error(c, err.Error())
		fmt.Println("err", err.Error())
		return
	}
	c.Set("user_id", t.GetUserID())
	c.Next()
}

//验证是否为管理员
func ValidationAdmin(c *gin.Context) {
	currentUser := util.CurrentUser(c)
	userType := currentUser.UserType
	if userType != 1 {
		controller.RestController{}.Error(c, "您不是管理员，无权访问！")
		return
	}
	fmt.Println("user：",userType)
	c.Next()
}
