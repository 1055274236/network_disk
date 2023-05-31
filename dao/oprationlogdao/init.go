package oprationlogdao

import (
	"NetworkDisk/dao"
)

func Add(userId int, params string, url string, duration int, ip string) (int64, error) {
	item := OperationLogTableStruct{UserId: userId, Params: params, Url: url, Duration: duration, Ip: ip}
	result := dao.MysqlDb.Create(&item)
	return result.RowsAffected, result.Error
}
