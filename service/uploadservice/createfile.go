package uploadservice

import (
	"NetworkDisk/dao/fileindexdao"
	"NetworkDisk/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateFileIndex(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	parentIdStr := ctx.PostForm("parentId")
	name := ctx.PostForm("name")
	parentId, atoiErr := strconv.Atoi(parentIdStr)
	if atoiErr != nil {
		service.SendErrorJson(ctx, nil, "参数错误！")
		return
	}

	resultArr, resultErr := fileindexdao.GetByUserIdAndParentId(userId.(int), parentId)
	if resultErr != nil {
		panic("数据库错误！")
	}
	for {
		if findSome(resultArr, name) {
			name = name + "(1)"
			continue
		}
		break
	}
	result, resultErr := fileindexdao.Add(name, false, 0, 0, parentId, userId.(int), false)
	if resultErr != nil {
		panic("数据库错误！")
	}
	service.SendSuccessJson(ctx, struct {
		Id int `json:"id"`
	}{result.Id}, "创建成功！")
}

func findSome(arr []fileindexdao.FileIndexTableStruct, name string) bool {
	for _, item := range arr {
		if item.FileName == name {
			return true
		}
	}
	return false
}
