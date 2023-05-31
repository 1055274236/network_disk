package loginlogdao

import (
	"NetworkDisk/dao"
)

func AddOne(userId int, userAccount string, ip string, device string) (LoginLogTableStruct, error) {
	item := LoginLogTableStruct{UserId: userId, UserAccount: userAccount, Ip: ip, Device: device}
	result := dao.MysqlDb.Create(&item)
	return item, result.Error
}
