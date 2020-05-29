package migrate

import (
	"encoding/json"
	"fmt"
	"gincmf/app/model"
	"github.com/gincmf/cmf"
)

type Option struct {
	MigrateStruct `gorm:"-"`
	Id            int
	Autoload        int    `gorm:"type:tinyint(3);default:1;not null"`
	OptionName	string `gorm:"type:varchar(64);not null"`
	OptionValue string `gorm:"type:text"`
}

func (m *Option) AutoMigrate() {
	fmt.Println("配置文件迁移文件")
	cmf.Db.AutoMigrate(&Option{})
	siteResult := cmf.Db.First(&Option{}, "option_name = ?", "site_info") // 查询
	if siteResult.RecordNotFound() {
		cmf.Db.Create(&Option{Autoload: 1,OptionName: "site_info"})

		option := &model.Option{}
		siteResult := cmf.Db.First(option, "option_name = ?", "site_info") // 查询
		if !siteResult.RecordNotFound() {
			siteInfoStr := option.OptionValue
			json.Unmarshal([]byte(siteInfoStr), &model.SiteInfo{})
			siteInfoValue, _ := json.Marshal(&model.SiteInfo{})
			cmf.Db.Model(option).Where("id = ?", option.Id).Update("option_value", string(siteInfoValue))
		}

	}

	uploadResult := cmf.Db.First(&Option{}, "option_name = ?", "upload_settings") // 查询
	if uploadResult.RecordNotFound() {
		cmf.Db.Create(&Option{Autoload: 1,OptionName: "upload_settings"})
	}
}
