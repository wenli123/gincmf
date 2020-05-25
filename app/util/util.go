package util

import (
	"gincmf/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf"
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
	}
	user := session.Get("current_user")
	return user.(model.User)
}
