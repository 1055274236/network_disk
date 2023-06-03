package errlogdao

import "NetworkDisk/dao"

func Add(url string, header string, params string, err string) (ErrLogTableStruct, error) {
	item := ErrLogTableStruct{Url: url, Header: header, Params: params, Err: err}
	result := dao.MysqlDb.Create(&item)
	return item, result.Error
}
