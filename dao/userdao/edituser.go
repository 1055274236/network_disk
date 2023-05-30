package userdao

import (
	"NetworkDisk/dao"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context, user UserTableStruct) error {
	err := dao.MysqlDb.Create(&user)
	return err.Error
}

func GetById(ctx *gin.Context, id int) (UserTableStruct, int64) {
	var user UserTableStruct
	result := dao.MysqlDb.Limit(1).Where(UserTableStruct{Id: id}, "Id").Find(&user)
	return user, result.RowsAffected
}

func GetByAccount(ctx *gin.Context, account string) (UserTableStruct, int64) {
	var user UserTableStruct
	result := dao.MysqlDb.Where(UserTableStruct{Account: account}, "Account").Take(&user)
	return user, result.RowsAffected
}
