package filestoredao

import (
	"time"
)

type FileStoreTableStruct struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Folder      string    `gorm:"column:folder;type:varchar(255);NOT NULL" json:"folder"`
	File        string    `gorm:"column:file;type:varchar(255);NOT NULL" json:"file"`
	Type        string    `gorm:"column:type;type:varchar(255);NOT NULL" json:"type"`
	Size        int64     `gorm:"column:size;type:bigint(20);NOT NULL" json:"size"`
	Md5         string    `gorm:"column:md5;type:varchar(255);NOT NULL" json:"md5"`
	Sha1        string    `gorm:"column:sha1;type:varchar(255);NOT NULL" json:"sha1"`
	Status      int       `gorm:"column:status;type:tinyint(1);NOT NULL" json:"status"`
	CreatedUser int       `gorm:"column:created_user;type:int(11);NOT NULL" json:"created_user"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;NOT NULL" json:"created_at"`
}

func (m *FileStoreTableStruct) TableName() string {
	return "file_store"
}
