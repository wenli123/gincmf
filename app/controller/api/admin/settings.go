package admin

import (
	"encoding/json"
	"gincmf/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf"
	"github.com/gincmf/cmf/controller"
)

// AssetsController 图片资源控制器，定义了资源文件增删改查接口
type SettingsController struct {
	rc controller.RestController
}

func (rest *SettingsController) Get(c *gin.Context) {

	option := &model.Option{}
	siteResult := cmf.Db.First(option, "option_name = ?", "site_info") // 查询
	if !siteResult.RecordNotFound() {
		rest.rc.Success(c, "获取成功", option)
	}else{
		rest.rc.Success(c, "获取失败","")
	}


}

func (rest *SettingsController) Show(c *gin.Context) {
	var rewrite struct {
		id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *SettingsController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *SettingsController) Store(c *gin.Context) {

	option := &model.Option{}
	siteResult := cmf.Db.First(option, "option_name = ?", "site_info") // 查询
	if !siteResult.RecordNotFound() {
		siteInfoStr := option.OptionValue
		json.Unmarshal([]byte(siteInfoStr), &model.SiteInfo{})
		siteInfoValue, _ := json.Marshal(&model.SiteInfo{})
		cmf.Db.Model(option).Where("id = ?", option.Id).Update("option_value", string(siteInfoValue))
	}
	rest.rc.Success(c, "修改成功", &model.SiteInfo{})
}

func (rest *SettingsController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}
