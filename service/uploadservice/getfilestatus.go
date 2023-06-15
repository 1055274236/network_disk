package uploadservice

import (
	"NetworkDisk/dao/filestoredao"
	"NetworkDisk/service"
	"errors"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type baseReturn struct {
	Status int   `json:"status"`
	Size   int64 `json:"size"`
}

func GetFileStatus(ctx *gin.Context) {
	md5 := ctx.Query("md5")
	sha1 := ctx.Query("sha1")

	if md5 != "" && sha1 != "" {
		result, daoErr := filestoredao.GetByMd5AndSha1One(md5, sha1)
		if daoErr != nil {
			if errors.Is(daoErr, gorm.ErrRecordNotFound) {
				service.SendJson(ctx, http.StatusNotFound, nil, "未找到相关数据！")
				return
			}
			panic("数据库错误！")
		}

		if result.Status == 1 {
			service.SendSuccessJson(ctx, baseReturn{1, result.Size}, "操作成功！")
			return
		} else {
			filtStat, statErr := os.Stat(path.Join("file", result.Folder, result.File))
			if statErr != nil {
				service.SendJson(ctx, http.StatusNotFound, nil, "文件已删除！")
				filestoredao.DeleteById(result.Id)
				return
			}

			service.SendSuccessJson(ctx, baseReturn{0, filtStat.Size()}, "操作成功！")
		}
	} else {
		service.SendErrorJson(ctx, nil, "参数错误！")
	}
}
