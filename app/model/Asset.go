package model

type Asset struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id";gorm:"type:bigint(20);not null"`
	FileSize   int    `json:"file_size";gorm:"type:bigint(20);not null"`
	CreateTime int    `json:"create_time";gorm:"type:int(10);default:0"`
	Status     int    `json:"status";gorm:"type:tinyint(3);default:1"`
	FileKey    string `json:"file_key";gorm:"type:varchar(64);not null`
	FileName   string `json:"file_name";gorm:"type:varchar(100);not null`
	FilePath   string `json:"file_path";gorm:"type:varchar(100);not null`
	Suffix     string `json:"suffix";gorm:"type:varchar(10);not null`
	AssetType  int    `json:"asset_type";gorm:"column:type;type:tinyint(3);not null"`
	more       string `json:"more";gorm:"type:text`
}
