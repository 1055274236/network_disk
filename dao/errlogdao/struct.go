package errlogdao

import "time"

type ErrLogTableStruct struct {
	Id        int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Url       string    `gorm:"column:url;type:varchar(255);NOT NULL" json:"url"`
	Header    string    `gorm:"column:header;type:text;NOT NULL" json:"header"`
	Params    string    `gorm:"column:params;type:varchar(255);NOT NULL" json:"params"`
	Err       string    `gorm:"column:err;type:varchar(255);NOT NULL" json:"err"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;NOT NULL" json:"created_at"`
	Status    int       `gorm:"column:status;type:tinyint(1);default:0;NOT NULL" json:"status"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;NOT NULL" json:"updated_at"`
}

func (m *ErrLogTableStruct) TableName() string {
	return "err_log"
}
