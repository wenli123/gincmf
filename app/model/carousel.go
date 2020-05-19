package model

type Slide struct {
	Id            int `json:"id"`
	Status        int    `json:"status";gorm:"type:tinyint(3);default:0"`
	Name          string `json:"name";gorm:"type:varchar(50);not null"`
	Remark        string `json:"remark";gorm:"type:varchar(255);not null"`
	DeleteAt      int `json:"delete_at"`
}
