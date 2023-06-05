package userdao

import (
	"NetworkDisk/dao"
)

func AddOne(account string, password string, cover string, maxCapacity int64) (UserTableStruct, error) {
	item := UserTableStruct{Account: account, Password: password, Cover: cover, MaxCapacity: maxCapacity}
	err := dao.MysqlDb.Create(&item)
	return item, err.Error
}

func GetById(id int) (UserTableStruct, error) {
	var user UserTableStruct
	result := dao.MysqlDb.First(&user, id)
	return user, result.Error
}

func GetByAccount(account string) (UserTableStruct, error) {
	user := UserTableStruct{}
	result := dao.MysqlDb.Where("account = ?", account).First(&user)
	return user, result.Error
}
