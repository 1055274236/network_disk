package uploadservice

import (
	"NetworkDisk/dao/fileindexdao"
	"NetworkDisk/dao/filestoredao"
	"NetworkDisk/service"
	"errors"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func CreateFileIndex(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	parentIdStr := ctx.PostForm("parentId")
	name := ctx.PostForm("name")
	parentId, atoiErr := strconv.Atoi(parentIdStr)

	md5Str := ctx.Query("md5")
	sha1Str := ctx.Query("sha1")
	if atoiErr != nil || md5Str == "" || sha1Str == "" {
		service.SendErrorJson(ctx, nil, "参数错误！")
		return
	}

	fileSome, isRErr := filestoredao.GetByMd5AndSha1One(md5Str, sha1Str)
	var fileStore filestoredao.FileStoreTableStruct
	if isRErr != nil {
		if errors.Is(isRErr, gorm.ErrRecordNotFound) {
			staticFolderName := time.Now().Format("20060102")
			staticFileName := uuid.NewV4().String()
			if os.MkdirAll(path.Join("file", staticFolderName), os.ModePerm) != nil {
				panic("文件夹创建错误！")
			}
			f, fErr := os.Create(path.Join("file", staticFolderName, staticFileName))
			if fErr != nil {
				panic("文件创建失败！")
			}
			f.Close()
			fileSome, _ = filestoredao.Add(staticFolderName, staticFileName, "", 0, "", "", userId.(int), 0)
		}
	}
	fileStore = fileSome

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

	result, resultErr := fileindexdao.Add(name, false, fileStore.Id, parentId, userId.(int), false)
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
