package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type Asset struct {
	Migrate
}

func (_ *Asset) AutoMigrate() {
	cmf.Db.AutoMigrate(&model.Asset{})
}

