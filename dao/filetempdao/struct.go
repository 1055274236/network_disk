package filetempdao

import (
	"time"
)

type FileTempTableStruct struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	FileName    string    `gorm:"column:file_name;type:varchar(255);NOT NULL" json:"file_name"`
	FileId      int       `gorm:"column:file_id;type:int(11);NOT NULL" json:"file_id"`
	CreatedUser int       `gorm:"column:created_user;type:int(11);NOT NULL" json:"created_user"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;NOT NULL" json:"created_at"`
	Timeout     time.Time `gorm:"column:timeout;type:datetime;NOT NULL" json:"timeout"`
}

func (m *FileTempTableStruct) TableName() string {
	return "file_temp"
}
