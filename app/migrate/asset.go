package migrate

import (
	"fmt"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type Asset struct {
	MigrateStruct `gorm:"-"`
	Id            int
	UserId        int    `gorm:"type:bigint(20);comment:所属用户id;not null"`
	FileSize      int    `gorm:"type:bigint(20);comment:文件大小;not null"`
	CreateTime    int    `gorm:"type:int(10);comment:上传时间;default:0"`
	Status        int    `gorm:"type:tinyint(3);comment:文件状态;default:1"`
	FileKey       string `gorm:"type:varchar(64);comment:文件惟一码;not null`
	FileName      string `gorm:"type:varchar(100);comment:文件名;not null`
	FilePath      string `gorm:"type:varchar(100);comment:文件路径;not null`
	Suffix        string `gorm:"type:varchar(10);comment:文件后缀;not null`
	AssetType     int    `gorm:"column:type;type:tinyint(3);comment:资源类型;not null"`
	more          string `gorm:"type:text;comment:更多配置'`
}

func (m *Asset) AutoMigrate() {
	fmt.Println("资源文件管理迁移文件")
	cmf.Db.AutoMigrate(&Asset{})
}
