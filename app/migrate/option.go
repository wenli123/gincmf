package migrate

import (
	"encoding/json"
	"fmt"
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type Option struct {
	MigrateStruct `gorm:"-"`
	Id            int
	AutoLoad      int    `gorm:"type:tinyint(3);default:1;not null"`
	OptionName    string `gorm:"type:varchar(64);not null"`
	OptionValue   string `gorm:"type:text"`
}

func (m *Option) AutoMigrate() {
	fmt.Println("配置文件迁移文件")
	cmf.Db.AutoMigrate(&Option{})
	siteResult := cmf.Db.First(&Option{}, "option_name = ?", "site_info") // 查询
	if siteResult.RecordNotFound() {
		//初始化默认json
		siteInfo := &model.SiteInfo{}
		siteInfoValue, _ := json.Marshal(siteInfo)
		cmf.Db.Create(&Option{AutoLoad: 1, OptionName: "site_info", OptionValue: string(siteInfoValue)})
	}

	uploadResult := cmf.Db.First(&Option{}, "option_name = ?", "upload_setting") // 查询
	if uploadResult.RecordNotFound() {
		//初始化默认json
		uploadSetting := &model.UploadSetting{}
		uploadSettingValue, _ := json.Marshal(uploadSetting)
		fmt.Println("uploadSettingValue",string(uploadSettingValue))
		cmf.Db.Create(&Option{AutoLoad: 1, OptionName: "upload_setting", OptionValue: string(uploadSettingValue)})
	}
}
