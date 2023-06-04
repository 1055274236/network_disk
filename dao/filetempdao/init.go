package filetempdao

import "NetworkDisk/dao"

func Add(name string, id int) (FileTempTableStruct, error) {
	temp := FileTempTableStruct{FileName: name, Id: id}
	result := dao.MysqlDb.Create(&temp)
	return temp, result.Error
}

func GetById(id int) (FileTempTableStruct, error) {
	temp := FileTempTableStruct{}
	result := dao.MysqlDb.Limit(1).Find(&temp, id)
	return temp, result.Error
}
