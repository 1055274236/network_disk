package userdao

import (
	"time"
)

type UserTableStruct struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Account     string    `gorm:"column:account;type:varchar(63);NOT NULL" json:"account"`
	Password    string    `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	Name        string    `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Cover       string    `gorm:"column:cover;type:varchar(255);comment:头像" json:"cover"`
	MaxCapacity int64     `gorm:"column:max_capacity;type:bigint(20);default:52428800;comment:最大容量;NOT NULL" json:"max_capacity"`
	NowCapacity int64     `gorm:"column:now_capacity;type:bigint(20);default:0;comment:当前容量;NOT NULL" json:"now_capacity"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间;NOT NULL" json:"updated_at"`
}

func (m *UserTableStruct) TableName() string {
	return "user"
}
