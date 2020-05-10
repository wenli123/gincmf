package migrate

import (
	"fmt"
	"github.com/gincmf/cmf"
	"time"
)

type User struct {
	MigrateStruct `gorm:"-"`
	Id int
	UserType int `gorm:"type:tinyint(3);not null"`
	Gender int `gorm:"type:tinyint(2);not null"`
	Birthday time.Time
	LastLoginTime time.Time
	Score int
	Coin int
	Balance float64 `gorm:"type:decimal(10,2);not null"`
	CreateTime time.Time
	UserStatus int `gorm:"type:tinyint(3);not null"`
	UserLogin string `gorm:"type:varchar(60);not null"`
	UserPass string `gorm:"type:varchar(64);not null"`
	UserNickname string `gorm:"type:varchar(50);not null"`
	UserEmail string `gorm:"type:varchar(100);not null"`
	UserUrl string `gorm:"type:varchar(100);not null"`
	Avatar string
	Signature string
	LastLoginip string
	UserActivationKey string
	Mobile string `gorm:"type:varchar(20);not null"`
	more string `gorm:"type:text"`
}

func (u User) TableName() string {
	return cmf.Config.Datebase.Prefix + "user"
}

func (u *User) AutoMigrate() {
	fmt.Println("用户数据迁移")
	cmf.Db.AutoMigrate(User{})
}
