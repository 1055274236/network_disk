package uploadservice

import (
	"NetworkDisk/dao/fileindexdao"
	"NetworkDisk/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Mkdir(ctx *gin.Context) {
	name := ctx.PostForm("name")
	parentStr := ctx.PostForm("parentId")
	parentId, err := strconv.Atoi(parentStr)

	if err != nil || len(name) == 0 {
		service.SendErrorJson(ctx, nil, "参数错误！")
		return
	}
	userId, _ := ctx.Get("userId")

	isRepetition, err := fileindexdao.GetIsRepetition(userId.(int), parentId, name)
	if isRepetition || err != nil {
		service.SendErrorJson(ctx, nil, "文件名重复！")
		return
	}

	dir, err := fileindexdao.Add(name, true, 0, 0, parentId, userId.(int))
	if err != nil {
		service.SendErrorJson(ctx, err, "数据错误！")
		return
	}
	service.SendSuccessJson(ctx, dir, "文件夹创建成功！")
}
