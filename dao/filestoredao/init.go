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
func Add(folder string, file string, fileType string, size int64, md5 string, sha1 string, createdUser int, status int) (FileStoreTableStruct, error) {
	item := FileStoreTableStruct{Folder: folder, File: file, Type: fileType, Size: size,
		Md5: md5, Sha1: sha1, CreatedUser: createdUser, Status: status}
	result := dao.MysqlDb.Create(&item)
	return item, result.Error
}

func GetByMd5AndSha1(md5 string, sha1 string) ([]FileStoreTableStruct, error) {
	temp := []FileStoreTableStruct{}
	result := dao.MysqlDb.Where(&FileStoreTableStruct{Md5: md5, Sha1: sha1}).Find(&temp)
	return temp, result.Error
}

func GetByMd5AndSha1One(md5 string, sha1 string) (FileStoreTableStruct, error) {
	temp := FileStoreTableStruct{}
	result := dao.MysqlDb.Where(&FileStoreTableStruct{Md5: md5, Sha1: sha1, Status: 1}).First(&temp)
	return temp, result.Error
}

func GetById(id int) (FileStoreTableStruct, error) {
	temp := FileStoreTableStruct{Id: id}
	result := dao.MysqlDb.First(&temp, id)
	return temp, result.Error
}

func ChangeType(id int, fileType string) (FileStoreTableStruct, error) {
	temp := FileStoreTableStruct{Id: id}
	result := dao.MysqlDb.Model(&temp).UpdateColumn("type", fileType)
	return temp, result.Error
}

func DeleteById(id int) error {
	temp := FileStoreTableStruct{Id: id}
	result := dao.MysqlDb.Delete(&temp)
	return result.Error
}

func ChangeModifiableData(id int, fileType string, size int64, status int, user int) (FileStoreTableStruct, error) {
	temp := FileStoreTableStruct{Id: id}
	result := dao.MysqlDb.Model(&temp).Updates(FileStoreTableStruct{Type: fileType, Size: size, Status: status, CreatedUser: user})
	return temp, result.Error
}
