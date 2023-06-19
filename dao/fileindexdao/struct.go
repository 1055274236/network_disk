package fileindexdao

import (
	"time"
)

type FileIndexTableStruct struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	FileName    string    `gorm:"column:file_name;type:varchar(255);NOT NULL" json:"file_name"`
	IsDir       int       `gorm:"column:is_dir;type:tinyint(1);default:0;NOT NULL" json:"is_dir"`
	FileNum     int       `gorm:"column:file_num;type:int(11);default:0;NOT NULL" json:"file_num"`
	StaticId    int       `gorm:"column:static_id;type:int(11);NOT NULL" json:"static_id"`
	ParentId    int       `gorm:"column:parent_id;type:int(11);default:-1;NOT NULL" json:"parent_id"`
	HoldingUser int       `gorm:"column:holding_user;type:int(11);NOT NULL" json:"holding_user"`
	IsShow      int       `gorm:"column:is_show;type:tinyint(1);NOT NULL" json:"is_show"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;NOT NULL" json:"created_at"`
	UpdatedAd   time.Time `gorm:"column:updated_ad;type:datetime;NOT NULL" json:"updated_ad"`
}

func (m *FileIndexTableStruct) TableName() string {
	return "file_index"
}

type UserTableStruct struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	NowCapacity int64     `gorm:"column:now_capacity;type:bigint(20);default:0;comment:当前容量;NOT NULL" json:"now_capacity"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间;NOT NULL" json:"updated_at"`
}

func (m *UserTableStruct) TableName() string {
	return "user"
}

type FileStoreTableStruct struct {
	Id   int   `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Size int64 `gorm:"column:size;type:bigint(20);NOT NULL" json:"size"`
}

func (m *FileStoreTableStruct) TableName() string {
	return "file_store"
}
