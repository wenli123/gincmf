package model

type Slide struct {
	Id       int    `json:"id"`
	Status   string `gorm:"type:tinyint(3);default:0" json:"status"`
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Remark   string `gorm:"type:varchar(255);not null" json:"remark"`
	DeleteAt int    `json:"delete_at"`
}

type SlideItem struct {
	Id          int
	SlideId     int     `gorm:"type:tinyint(3);not null" json:"slide_id"`
	Status      int     `gorm:"type:tinyint(3);default:1" json:"status"`
	ListOrder   float64 `gorm:"type:float;default:10000" json:"list_order"`
	Title       string  `gorm:"type:varchar(50);not null" json:"title"`
	Image       string  `gorm:"type:varchar(255);not null" json:"image"`
	Url         string  `gorm:"type:varchar(255)" json:"url"`
	Target      string  `gorm:"type:varchar(10)" json:"target"`
	Description string  `gorm:"type:varchar(255)" json:"description"`
	Content     string  `gorm:"type:text" json:"content"`
	More        string  `gorm:"type:text" json:"more"`
}
