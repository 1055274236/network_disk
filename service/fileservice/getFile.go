package fileservice

import (
	"NetworkDisk/dao/fileindexdao"
	"NetworkDisk/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFile(ctx *gin.Context) {
	parentIdStr := ctx.PostForm("parentId")
	parentId, atoiErr := strconv.Atoi(parentIdStr)
	userId, hasUserId := ctx.Get("userId")

	if atoiErr != nil || !hasUserId {
		service.SendBadRequestJson(ctx, nil, "错误参数！")
		return
	}

	result, daoErr := fileindexdao.GetByUserIdAndParentIdShow(userId.(int), parentId)
	if daoErr != nil {
		panic("数据库错误！")
	}

	service.SendSuccessJson(ctx, result, "操作成功！")
}
