package fileindexdao

import (
	"NetworkDisk/dao"

	"gorm.io/gorm"
)

// @params:
// name 文件名称
// size 文件尺寸
// isDir 是否为文件夹
// static_id 文件对应静态文件Id
// parentId 父文件夹Id
// holdingUser 拥有者用户ID
func Add(name string, isDir bool, size int64, staticId int, parentId int, holdingUser int) (FileIndexTableStruct, error) {
	item := FileIndexTableStruct{FileName: name, ParentId: parentId, HoldingUser: holdingUser, IsDir: 1}
	if !isDir {
		item.Size = size
		item.StaticId = staticId
		item.IsDir = 0
	}
	result := dao.MysqlDb.Create(&item)
	dao.MysqlDb.Model(&FileIndexTableStruct{}).Where("id = ?", parentId).UpdateColumn("file_num", gorm.Expr("file_num + ?", 1))
	return item, result.Error
}

func DeleteByIdAndAccount(ids []string, holdingUser int) (int64, error) {
	deleteArr := []FileIndexTableStruct{}
	parentMap := make(map[int]int)
	for id := range ids {
		deleteArr = append(deleteArr, FileIndexTableStruct{Id: id})
	}
	result := dao.MysqlDb.Where("holding_user = ?", holdingUser).Delete(&deleteArr)

	// 修改父文件夹文件数量
	for _, item := range deleteArr {
		if item.ParentId == 0 {
			continue
		}
		if value, ok := parentMap[item.ParentId]; ok {
			parentMap[item.ParentId] = value + 1
		} else {
			parentMap[item.ParentId] = 1
		}
	}
	for key, value := range parentMap {
		dao.MysqlDb.Model(&FileIndexTableStruct{}).Where(
			&FileIndexTableStruct{HoldingUser: holdingUser, Id: key}).UpdateColumn("file_num", gorm.Expr("file_num - ?", value))
	}
	return result.RowsAffected, result.Error
}

// @params
// userId int 拥有者ID
// parentId int 父文件夹ID
func GetByUserIdAndParentId(userId int, parentId int) ([]FileIndexTableStruct, error) {
	value := []FileIndexTableStruct{}
	var result *gorm.DB
	// 适当优化搜索
	if parentId == 0 {
		result = dao.MysqlDb.Where(&FileIndexTableStruct{HoldingUser: userId, ParentId: parentId}).Find(&value)
	} else {
		parent := FileIndexTableStruct{Id: parentId}
		dao.MysqlDb.Limit(1).Find(&parent, parentId)
		result = dao.MysqlDb.Where(&FileIndexTableStruct{HoldingUser: userId, ParentId: parentId}).Limit(parent.FileNum).Find(&value)
	}
	return value, result.Error
}

func GetById(id int) (FileIndexTableStruct, error) {
	value := FileIndexTableStruct{}
	result := dao.MysqlDb.Limit(1).Find(&value, id)
	return value, result.Error
}

// 检查当前文件夹下，是否名称相同
func GetIsRepetition(userId int, parentId int, name string) (bool, error) {
	result, err := GetByUserIdAndParentId(userId, parentId)
	if err != nil {
		return true, err
	}
	for _, item := range result {
		if item.FileName == name {
			return true, nil
		}
	}
	return false, nil
}
