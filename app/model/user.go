package model

type User struct {
	Id                int     `json:"id"`
	UserType          int     `gorm:"type:tinyint(3);not null" json:"user_type"`
	Gender            int     `gorm:"type:tinyint(2)" json:"gender"`
	Birthday          int64   `gorm:"type:int(11)" json:"birthday"`
	LastLoginTime     int64   `gorm:"type:int(11)" json:"last_login_time"`
	Score             int     `json:"score"`
	Coin              int     `json:"coin"`
	Balance           float64 `gorm:"type:decimal(10,2);not null" json:"balance"`
	CreateTime        int64   `gorm:"type:int(11)" json:"create_time"`
	UserStatus        int     `gorm:"type:tinyint(3);not null" json:"user_status"`
	UserLogin         string  `gorm:"type:varchar(60);not null" json:"user_login"`
	UserPass          string  `gorm:"type:varchar(64);not null" json:"user_pass"`
	UserNickname      string  `gorm:"type:varchar(50);not null" json:"user_nickname"`
	UserEmail         string  `gorm:"type:varchar(100);not null" json:"user_email"`
	UserUrl           string  `gorm:"type:varchar(100);not null" json:"user_url"`
	Avatar            string  `json:"avatar"`
	Signature         string  `json:"signature"`
	LastLoginip       string  `json:"last_loginip"`
	UserActivationKey string  `json:"user_activation_key"`
	Mobile            string  `gorm:"type:varchar(20);not null" json:"mobile"`
	more              string  `gorm:"type:text" json:"more"`
}
