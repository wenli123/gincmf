package migrate

import (
	"fmt"
	"github.com/gincmf/cmf"
)

type Asset struct {
	MigrateStruct `gorm:"-"`
	Id            int
	UserId        int    `gorm:"type:bigint(20);not null"`
	FileSize      int    `gorm:"type:bigint(20);not null"`
	CreateTime    int    `gorm:"type:int(10);default:0"`
	Status        int    `gorm:"type:tinyint(3);default:1"`
	FileKey       string `gorm:"type:varchar(64);not null`
	FileName      string `gorm:"type:varchar(100);not null`
	FilePath      string `gorm:"type:varchar(100);not null`
	suffix        string `gorm:"type:varchar(10);not null`
	AssetType     int    `gorm:"column:type;type:tinyint(3);not null"`
	more          string `gorm:"type:text`
}

func (m *Asset) AutoMigrate() {
	fmt.Println("资源文件管理迁移文件")
	cmf.Db.AutoMigrate(&Asset{})
}
