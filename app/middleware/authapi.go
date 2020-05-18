package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
	"github.com/gincmf/cmf/router"
	"gopkg.in/oauth2.v3"
)

var Token oauth2.TokenInfo

func ValidationBearerToken(c *gin.Context)  {
	s := router.Srv
	t, err := s.ValidationBearerToken(c.Request)
	Token = t
	if err != nil {
		controller.RestControllerStruct{}.Error(c,err.Error())
		fmt.Println("err",err.Error())
	}

	c.Set("user_id",t.GetUserID())

	c.Next()
}