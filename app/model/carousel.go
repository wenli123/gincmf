package model

type Slide struct {
	Id       int    `json:"id"`
	Status   string `json:"status";gorm:"type:tinyint(3);default:0"`
	Name     string `json:"name";gorm:"type:varchar(50);not null"`
	Remark   string `json:"remark";gorm:"type:varchar(255);not null"`
	DeleteAt int    `json:"delete_at"`
}

type SlideItem struct {
	Id          int
	SlideId     int     `gorm:"type:tinyint(3);not null"`
	Status      int     `gorm:"type:tinyint(3);default:1"`
	ListOrder   float64 `gorm:"type:float;default:10000"`
	Title       string  `gorm:"type:varchar(50);not null"`
	Image       string  `gorm:"type:varchar(255);not null"`
	Url         string  `gorm:"type:varchar(255)"`
	Target      string  `gorm:"type:varchar(10)"`
	Description string  `gorm:"type:varchar(255)"`
	Content     string  `gorm:"type:text"`
	More        string  `gorm:"type:text"`
}
