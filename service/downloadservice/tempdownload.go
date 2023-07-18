package downloadservice

import (
	"NetworkDisk/dao/fileindexdao"
	"NetworkDisk/dao/filestoredao"
	"NetworkDisk/dao/filetempdao"
	"NetworkDisk/service"
	"errors"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TempDownload(ctx *gin.Context) {
	idStr := ctx.Param("id")
	name := ctx.Param("name")
	id, err := strconv.Atoi(idStr)

	if len(idStr) == 0 || len(name) == 0 || err != nil {
		service.SendErrorJson(ctx, nil, "错误参数！")
		return
	}

	temp, err := filetempdao.GetById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		service.SendErrorJson(ctx, nil, "未查找到该数据！")
		return
	} else if err != nil {
		panic("文件缓存读取错误！")
	}

	file, err := filestoredao.GetById(temp.FileId)
	if err != nil {
		panic("文件缓存指定文件读取错误！")
	}

	ctx.File(path.Join("/", file.Folder, file.File))
}

func DownloadFileByIndex(ctx *gin.Context) {
	idStr := ctx.Param("id")
	name := ctx.Param("name")
	id, err := strconv.Atoi(idStr)

	if err != nil || name == "" {
		service.SendBadRequestJson(ctx, nil, "错误参数！")
		return
	}
	fileIndex, fileIndexErr := fileindexdao.GetById(id)
	if fileIndexErr != nil || name != fileIndex.FileName {
		service.SendBadRequestJson(ctx, nil, "错误参数！")
		return
	}
	filestore, fileStoreErr := filestoredao.GetById(fileIndex.StaticId)
	if fileStoreErr != nil {
		service.SendBadRequestJson(ctx, nil, "错误参数！")
		return
	}
	ctx.File(path.Join("file", filestore.Folder, filestore.File))
}
