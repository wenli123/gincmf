package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
	"github.com/gincmf/cmf/router"
)

func ValidationBearerToken(c *gin.Context)  {
	s := router.Srv
	token, err := s.ValidationBearerToken(c.Request)
	if err != nil {
		controller.RestControllerStruct{}.Error(c,err.Error())
		fmt.Println("err",err.Error())
	}
	//c.Set("token",token)
	fmt.Println("token",token)
}