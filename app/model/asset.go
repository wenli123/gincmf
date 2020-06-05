package model

type Asset struct {
	Id         int    `json:"id"`
	UserId     int    `gorm:"type:bigint(20);comment:'所属用户id';not null" json:"user_id"`
	FileSize   int64    `gorm:"type:bigint(20);comment:'文件大小';not null" json:"file_size"`
	CreateTime int64    `gorm:"type:int(10);comment:'上传时间';default:0" json:"create_time"`
	Status     int    `gorm:"type:tinyint(3);comment:'文件状态';default:1" json:"status"`
	FileKey    string `gorm:"type:varchar(64);comment:'文件惟一码';not null" json:"file_key"`
	FileName   string `gorm:"type:varchar(100);comment:'文件名';not null" json:"file_name"`
	FilePath   string `gorm:"type:varchar(100);comment:'文件路径';not null" json:"file_path"`
	Suffix     string `gorm:"type:varchar(10);comment:'文件后缀';not null" json:"suffix"`
	AssetType  int    `gorm:"column:type;type:tinyint(3);comment:'资源类型';not null" json:"asset_type"`
	More       string `gorm:"type:text;comment:'更多配置'" json:"more"`
}
