package filetempdao

import "NetworkDisk/dao"

func Add(name string, id int) (FileTempTableStruct, error) {
	temp := FileTempTableStruct{FileName: name, Id: id}
	result := dao.MysqlDb.Create(&temp)
	return temp, result.Error
}

func GetById(id int) (FileTempTableStruct, error) {
	temp := FileTempTableStruct{}
	result := dao.MysqlDb.First(&temp, id)
	return temp, result.Error
}
