package model

type Option struct {
	Id            int
	Autoload        int    `gorm:"type:tinyint(3);default:1;not null"`
	OptionName	string `gorm:"type:varchar(64);not null"`
	OptionValue string `gorm:"type:text"`
}

//定义site_info类型
type SiteInfo struct {
	SiteName           string `json:"site_name"`
	SiteSeoTitle       string `json:"site_seo_title"`
	SiteSeoKeywords    string `json:"site_seo_keywords"`
	SiteSeoDescription string `json:"site_seo_description"`
	SiteIcp            string `json:"site_icp"`
	SiteGwa            string `json:"site_gwa"`
	SiteAdminEmail     string `json:"site_admin_email"`
	siteAnalytics      string `json:"site_analytics"`
}
