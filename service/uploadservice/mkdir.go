package uploadservice

import (
	"NetworkDisk/dao/fileindexdao"
	"NetworkDisk/service"
	"strconv"
	"strings"

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

	// isRepetition, err := fileindexdao.GetIsRepetition(userId.(int), parentId, name)
	// if isRepetition || err != nil {
	// 	service.SendErrorJson(ctx, nil, "文件名重复！")
	// 	return
	// }

	resultArr, resultErr := fileindexdao.GetByUserIdAndParentId(userId.(int), parentId)
	if resultErr != nil {
		panic("数据库错误！")
	}
	for {
		if findSome(resultArr, name) {
			i := strings.Index(name, ".")
			if i == -1 {
				i = len(name)
			}
			name = name[:i] + "(1)" + name[i:]
		} else {
			break
		}
	}

	dir, err := fileindexdao.Add(name, true, 0, parentId, userId.(int), true)
	if err != nil {
		service.SendErrorJson(ctx, err, "数据错误！")
		return
	}
	service.SendSuccessJson(ctx, dir, "文件夹创建成功！")
}
