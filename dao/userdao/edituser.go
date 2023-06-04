package userdao

import (
	"NetworkDisk/dao"
)

func AddOne(account string, password string, cover string, maxCapacity int64) (UserTableStruct, error) {
	item := UserTableStruct{Account: account, Password: password, Cover: cover, MaxCapacity: maxCapacity}
	err := dao.MysqlDb.Create(&item)
	return item, err.Error
}

func GetById(id int) (UserTableStruct, int64) {
	var user UserTableStruct
	result := dao.MysqlDb.Limit(1).Find(&user, id)
	return user, result.RowsAffected
}

func GetByAccount(account string) (UserTableStruct, int64) {
	var user UserTableStruct
	result := dao.MysqlDb.Where(UserTableStruct{Account: account}, "Account").Take(&user)
	return user, result.RowsAffected
}
