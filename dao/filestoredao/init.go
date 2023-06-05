package filestoredao

import "NetworkDisk/dao"

// @params:
// folder 文件夹名称
// file 文件名称
// fileType 文件类型
// size 文件大小
// md5 md5
// sha1 sha1
// createdUser 最初创建这个文件的用户
func Add(folder string, file string, fileType string, size int64, md5 string, sha1 string, createdUser int) (FileStoreTableStruct, error) {
	item := FileStoreTableStruct{Folder: folder, File: file, Type: fileType, Size: size, Md5: md5, Sha1: sha1, CreatedUser: createdUser}
	result := dao.MysqlDb.Create(&item)
	return item, result.Error
}

func GetByMd5AndSha1(md5 string, sha1 string) (FileStoreTableStruct, error) {
	temp := FileStoreTableStruct{}
	result := dao.MysqlDb.Limit(1).Where(&FileStoreTableStruct{Md5: md5, Sha1: sha1}).Find(&temp)
	return temp, result.Error
}

func GetById(id int) (FileStoreTableStruct, error) {
	temp := FileStoreTableStruct{Id: id}
	result := dao.MysqlDb.First(&temp, id)
	return temp, result.Error
}

func ChangeType(id int, fileType string) (FileStoreTableStruct, error) {
	temp := FileStoreTableStruct{Id: id, Type: fileType}
	result := dao.MysqlDb.Model(&temp).UpdateColumn("type", fileType)
	return temp, result.Error
}
