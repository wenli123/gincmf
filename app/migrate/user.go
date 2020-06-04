package migrate

import (
	"fmt"
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/util"
	"time"
)

type User struct {
	Migrate
}

func (u *User) AutoMigrate() {
	fmt.Println("用户数据迁移")
	cmf.Db.AutoMigrate(&model.User{})

	userResult := cmf.Db.First(&model.User{}, "user_login = ?", "admin") // 查询

	if userResult.RecordNotFound() {
		fmt.Println("user", userResult)
		//新增一条管理员数据
		cmf.Db.Create(&model.User{UserType: 1, UserLogin: "admin", UserPass: util.GetMd5("123456"), CreateTime: time.Now().Unix()})
	}
}
