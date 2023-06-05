package filetempdao

import (
	"NetworkDisk/dao"
	"time"
)

func Add(name string, id int, timeout time.Time) (FileTempTableStruct, error) {
	temp := FileTempTableStruct{FileName: name, Id: id, Timeout: timeout}
	result := dao.MysqlDb.Create(&temp)
	return temp, result.Error
}

func GetById(id int) (FileTempTableStruct, error) {
	temp := FileTempTableStruct{}
	result := dao.MysqlDb.First(&temp, id)
	return temp, result.Error
}
