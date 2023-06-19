package fileindexdao

import (
	"NetworkDisk/dao"

	"gorm.io/gorm"
)

// @params:
// name 文件名称
// isDir 是否为文件夹
// static_id 文件对应静态文件Id
// parentId 父文件夹Id
// holdingUser 拥有者用户ID
func Add(name string, isDir bool, staticId int, parentId int, holdingUser int, isShow bool) (FileIndexTableStruct, error) {
	item := FileIndexTableStruct{FileName: name, ParentId: parentId, HoldingUser: holdingUser, IsDir: 1, IsShow: 0}
	if !isDir {
		item.StaticId = staticId
		item.IsDir = 0
	}
	if isShow {
		item.IsShow = 1
	}

	trErr := dao.MysqlDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&item).Error; err != nil {
			return err
		}
		if err := tx.Model(&FileIndexTableStruct{Id: parentId}).UpdateColumn("file_num", gorm.Expr("file_num + ?", 1)).Error; err != nil {
			return err
		}
		tempFileStore := FileStoreTableStruct{Id: staticId}
		if err := tx.First(&tempFileStore, staticId).Error; err != nil {
			return err
		}
		if err := tx.Model(&UserTableStruct{Id: holdingUser}).UpdateColumn("now_capacity", gorm.Expr("now_capacity + ?", tempFileStore.Size)).Error; err != nil {
			return err
		}
		return nil
	})
	return item, trErr
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
		dao.MysqlDb.First(&parent, parentId)
		result = dao.MysqlDb.Where(&FileIndexTableStruct{HoldingUser: userId, ParentId: parentId}).Limit(parent.FileNum).Find(&value)
	}
	return value, result.Error
}

// @params
// userId int 拥有者ID
// parentId int 父文件夹ID
func GetByUserIdAndParentIdShow(userId int, parentId int) ([]FileIndexTableStruct, error) {
	value := []FileIndexTableStruct{}
	var result *gorm.DB
	// 适当优化搜索
	if parentId == 0 {
		result = dao.MysqlDb.Where(&FileIndexTableStruct{HoldingUser: userId, ParentId: parentId, IsShow: 1}).Find(&value)
	} else {
		parent := FileIndexTableStruct{Id: parentId}
		dao.MysqlDb.First(&parent, parentId)
		result = dao.MysqlDb.Where(&FileIndexTableStruct{HoldingUser: userId, ParentId: parentId, IsShow: 1}).Limit(parent.FileNum).Find(&value)
	}
	return value, result.Error
}

func GetById(id int) (FileIndexTableStruct, error) {
	value := FileIndexTableStruct{}
	result := dao.MysqlDb.First(&value, id)
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
