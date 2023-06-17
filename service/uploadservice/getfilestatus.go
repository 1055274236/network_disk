package uploadservice

import (
	"NetworkDisk/dao/fileindexdao"
	"NetworkDisk/dao/filestoredao"
	"NetworkDisk/service"
	"errors"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFileStatus(ctx *gin.Context) {
	md5 := ctx.Query("md5")
	sha1 := ctx.Query("sha1")

	if md5 != "" && sha1 != "" {
		_, daoErr := filestoredao.GetByMd5AndSha1One(md5, sha1)
		if daoErr != nil {
			if errors.Is(daoErr, gorm.ErrRecordNotFound) {
				service.SendNotFoundJson(ctx, nil, "未找到相关数据！")
				return
			}
			panic("数据库错误！")
		}

		service.SendSuccessJson(ctx, struct {
			IsExist bool `json:"isExist"`
		}{true}, "操作成功！")

	} else {
		service.SendErrorJson(ctx, nil, "参数错误！")
	}
}

func GetFileStatusById(ctx *gin.Context) {
	idStr := ctx.Query("id")
	userId, hasUserId := ctx.Get("userId")

	id, atoiErr := strconv.Atoi(idStr)
	if atoiErr != nil || !hasUserId {
		service.SendErrorJson(ctx, nil, "错误参数！")
		return
	}

	fileIndex, fileIndexErr := fileindexdao.GetById(id)
	if fileIndexErr != nil {
		if errors.Is(fileIndexErr, gorm.ErrRecordNotFound) {
			service.SendNotFoundJson(ctx, nil, "未找到相关数据！")
			return
		}
		panic("数据库错误！")
	}
	if fileIndex.HoldingUser != userId.(int) {
		service.SendErrorJson(ctx, nil, "无此权限！")
		return
	}

	fileStore, fileStoreErr := filestoredao.GetById(fileIndex.Id)
	if fileIndexErr != nil {
		if errors.Is(fileStoreErr, gorm.ErrRecordNotFound) {
			service.SendNotFoundJson(ctx, nil, "未找到相关数据！")
			return
		}
		panic("数据库错误！")
	}

	if fileStore.Status == 1 {
		service.SendSuccessJson(ctx, struct {
			Status int   `json:"status"`
			Size   int64 `json:"size"`
		}{1, fileStore.Size}, "操作成功！")
	} else {
		fs, fsErr := os.Stat(path.Join("file", fileStore.Folder, fileStore.File))
		if fsErr != nil {
			service.SendNotFoundJson(ctx, nil, "未找到相关文件！")
			return
		}

		service.SendSuccessJson(ctx, struct {
			Status int   `json:"status"`
			Size   int64 `json:"size"`
		}{0, fs.Size()}, "操作成功！")
	}
}
