package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type Slide struct {
	Migrate
}

type SlideItem struct {}

func (_ *Slide) AutoMigrate() {
	cmf.Db.AutoMigrate(model.Slide{})
	cmf.Db.AutoMigrate(model.SlideItem{})
}
