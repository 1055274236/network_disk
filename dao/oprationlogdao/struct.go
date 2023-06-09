package oprationlogdao

import (
	"time"
)

type OperationLogTableStruct struct {
	Id        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UserId    int       `gorm:"column:user_id;type:int(11);comment:用户操作id;NOT NULL" json:"user_id"`
	Params    string    `gorm:"column:Params;type:text" json:"params"`
	Url       string    `gorm:"column:url;type:varchar(255);comment:api网址;NOT NULL" json:"url"`
	Duration  int       `gorm:"column:duration;type:int(11);comment:毫秒数;NOT NULL" json:"duration"`
	Ip        string    `gorm:"column:ip;type:varchar(255)" json:"ip"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;NOT NULL" json:"created_at"`
}

func (m *OperationLogTableStruct) TableName() string {
	return "operation_log"
}
