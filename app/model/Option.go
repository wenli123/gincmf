package model

type Option struct {
	Id          int    `json:"id"`
	AutoLoad    int    `json:"autoload";gorm:"type:tinyint(3);default:1;not null"`
	OptionName  string `json:"option_name";gorm:"type:varchar(64);not null"`
	OptionValue string `json:"option_value";gorm:"type:text"`
}

//定义site_info类型
type SiteInfo struct {
	SiteName           string `json:"site_name"`
	AdminPassword      string `json:"admin_password"`
	SiteSeoTitle       string `json:"site_seo_title"`
	SiteSeoKeywords    string `json:"site_seo_keywords"`
	SiteSeoDescription string `json:"site_seo_description"`
	SiteIcp            string `json:"site_icp"`
	SiteGwa            string `json:"site_gwa"`
	SiteAdminEmail     string `json:"site_admin_email"`
	SiteAnalytics      string `json:"site_analytics"`
	OpenRegistration   string `json:"open_registration"`
}

//定义upload_setting类型
type UploadSetting struct {
	MaxFiles  string `json:"max_files"`
	ChunkSize string `json:"chunk_size"`
	FileTypes `json:"file_types"`
}

type FileTypes struct {
	Image TypeValues `json:"image"`
	video TypeValues `json:"video"`
	Audio TypeValues `json:"audio"`
	File  TypeValues `json:"file"`
}

type TypeValues struct {
	UploadMaxFileSize string `json:"upload_max_file_size"`
	Extensions        string `json:"extensions"`
}
