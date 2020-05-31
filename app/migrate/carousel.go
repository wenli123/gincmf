package migrate

import (
	"fmt"
	"github.com/gincmf/cmf"
)

type Slide struct {
	MigrateStruct `gorm:"-"`
	Id            int
	Status        int    `gorm:"type:tinyint(3);default:1"`
	Name          string `gorm:"type:varchar(50);not null"`
	Remark        string `gorm:"type:varchar(255);not null"`
	DeleteAt      int    `gorm:"type:int(11)";default:0`
}

type SlideItem struct {
	MigrateStruct `gorm:"-"`
	Id            int
	SlideId       int     `gorm:"type:tinyint(3);not null"`
	Status        int     `gorm:"type:tinyint(3);default:1"`
	ListOrder     float64 `gorm:"type:float;default:10000"`
	Title         string  `gorm:"type:varchar(50);not null"`
	Image         string  `gorm:"type:varchar(255);not null"`
	Url           string  `gorm:"type:varchar(255)"`
	Target        string  `gorm:"type:varchar(10)"`
	Description   string  `gorm:"type:varchar(255)"`
	Content       string  `gorm:"type:text"`
	More          string  `gorm:"type:text"`
}

func (m *Slide) AutoMigrate() {
	fmt.Println("轮播图数据迁移")
	cmf.Db.AutoMigrate(Slide{})
	cmf.Db.AutoMigrate(SlideItem{})
}
