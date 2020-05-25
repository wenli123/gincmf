package model

type User struct {
	Id                int     `json:"id"`
	UserType          int     `json:"user_type";gorm:"type:tinyint(3);not null"`
	Gender            int     `json:"gender";gorm:"type:tinyint(2);"`
	Birthday          int64   `json:"birthday;"gorm:"type:int(11)"`
	LastLoginTime     int64   `json:"last_login_time;"gorm:"type:int(11)"`
	Score             int     `json:"score"`
	Coin              int     `json:"coin"`
	Balance           float64 `json:"balance";gorm:"type:decimal(10,2);not null"`
	CreateTime        int64   `json:"create_time";gorm:"type:int(11)"`
	UserStatus        int     `json:"user_status";gorm:"type:tinyint(3);not null"`
	UserLogin         string  `json:"user_login";gorm:"type:varchar(60);not null"`
	userPass          string  `json:"user_pass";gorm:"type:varchar(64);not null"`
	UserNickname      string  `json:"user_nickname";gorm:"type:varchar(50);not null"`
	UserEmail         string  `json:"user_email";gorm:"type:varchar(100);not null"`
	UserUrl           string  `json:"user_url";gorm:"type:varchar(100);not null"`
	Avatar            string  `json:"avatar"`
	Signature         string  `json:"signature"`
	LastLoginip       string  `json:"last_loginip"`
	UserActivationKey string  `json:"user_activation_key"`
	Mobile            string  `json:"mobile;"gorm:"type:varchar(20);not null"`
	more              string  `json:"more;"gorm:"type:text"`
}
