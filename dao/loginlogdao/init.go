package loginlogdao

import (
	"NetworkDisk/dao"
)

func Add(userId int, userAccount string, ip string, device string) (int64, error) {
	item := LoginLogTableStruct{UserId: userId, UserAccount: userAccount, Ip: ip, Device: device}
	result := dao.MysqlDb.Create(&item)
	return result.RowsAffected, result.Error
}
