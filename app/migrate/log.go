package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type Log struct {
	Migrate
}

func (_ *Log) AutoMigrate() {
	cmf.Db.AutoMigrate(&model.Log{})
}
