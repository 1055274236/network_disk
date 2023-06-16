package uploadservice

import (
	"NetworkDisk/dao/filestoredao"
	"NetworkDisk/service"
	"errors"

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
