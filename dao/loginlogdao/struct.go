package loginlogdao

import (
	"time"
)

type LoginLogTableStruct struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UserId      int       `gorm:"column:user_id;type:int(11);comment:用户id;NOT NULL" json:"user_id"`
	UserAccount string    `gorm:"column:user_account;type:varchar(255);comment:用户账号;NOT NULL" json:"user_account"`
	Ip          string    `gorm:"column:ip;type:varchar(63);comment:用户登录ip;NOT NULL" json:"ip"`
	Device      string    `gorm:"column:device;type:varchar(255);comment:用户登录设备" json:"device"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;NOT NULL" json:"created_at"`
}

func (m *LoginLogTableStruct) TableName() string {
	return "login_log"
}
