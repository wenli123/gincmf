package model

type User struct {
	Id                int
	UserType          int   `gorm:"type:tinyint(3);not null"`
	Gender            int   `gorm:"type:tinyint(2);"`
	Birthday          int64 `gorm:"type:int(11)"`
	LastLoginTime     int64 `gorm:"type:int(11)"`
	Score             int
	Coin              int
	Balance           float64 `gorm:"type:decimal(10,2);not null"`
	CreateTime        int64   `gorm:"type:int(11)"`
	UserStatus        int     `gorm:"type:tinyint(3);not null"`
	UserLogin         string  `gorm:"type:varchar(60);not null"`
	UserPass          string  `gorm:"type:varchar(64);not null"`
	UserNickname      string  `gorm:"type:varchar(50);not null"`
	UserEmail         string  `gorm:"type:varchar(100);not null"`
	UserUrl           string  `gorm:"type:varchar(100);not null"`
	Avatar            string
	Signature         string
	LastLoginip       string
	UserActivationKey string
	Mobile            string `gorm:"type:varchar(20);not null"`
	more              string `gorm:"type:text"`
}

