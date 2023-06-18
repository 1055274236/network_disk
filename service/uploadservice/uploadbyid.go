package uploadservice

import (
	"NetworkDisk/dao/fileindexdao"
	"NetworkDisk/dao/filestoredao"
	"NetworkDisk/service"
	"bytes"
	"errors"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UploadById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userId, _ := ctx.Get("userId")
	fileId, fileIdErr := strconv.Atoi(idStr)
	if fileIdErr != nil {
		service.SendErrorJson(ctx, nil, "参数错误！")
		return
	}
	// proofread
	fileIndex, fileIndexErr := fileindexdao.GetById(fileId)
	if fileIndexErr != nil {
		if errors.Is(fileIndexErr, gorm.ErrRecordNotFound) {
			service.SendNotFoundJson(ctx, nil, "未找到相关数据！")
			return
		}
		panic("数据库错误！")
	}
	if fileIndex.HoldingUser != userId {
		service.SendErrorJson(ctx, nil, "权限错误！")
		return
	}
	fileStore, fileStoreErr := filestoredao.GetById(fileIndex.StaticId)
	if fileStoreErr != nil {
		if errors.Is(fileStoreErr, gorm.ErrRecordNotFound) {
			service.SendNotFoundJson(ctx, nil, "未找到相关数据！")
			return
		}
		panic("数据库错误！")
	}

	// Read And Write
	staticFilePath := path.Join("file", fileStore.Folder, fileStore.File)
	// 获取文件分隔符
	ct := ctx.Request.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(ct)
	_, hasb := params["boundary"]
	if err != nil || !hasb {
		service.SendErrorJson(ctx, nil, "错误请求方式！")
		return
	}
	// boundary 作为文件分隔符
	// 从请求报文来说，这玩意儿应该都是\r\n的，但是他删掉了一个，只能修改了
	boundary := []byte(params["boundary"])
	boundaryEnd := append(append([]byte{13, 10, 45, 45}, boundary...), []byte{45, 45, 13, 10}...)

	// 读取配置
	readBuf := make([]byte, 1024000)
	var jointChar []byte
	read_len, readErr := ctx.Request.Body.Read(readBuf)
	headEndIndex := 0   // 请求报文头结束位置
	var charTemp []byte // 暂时存储
	f, _ := os.OpenFile(staticFilePath, os.O_WRONLY, os.ModeAppend)

	// 非文件结束错误，暂时不管
	if readErr != nil && !errors.Is(readErr, io.EOF) {
		if f != nil {
			f.Close()
		}
		return
	}

	headEndIndex = bytes.Index(readBuf[:read_len], []byte{13, 10, 13, 10})
	jointChar = bytes.Clone(readBuf[headEndIndex+4 : read_len])
	charTemp = bytes.Clone(jointChar)

	if errors.Is(readErr, io.EOF) {
		endIndex := bytes.Index(jointChar, boundaryEnd)
		if endIndex == -1 {
			f.Write(jointChar)
		} else {
			f.Write(jointChar[:endIndex])
		}
		f.Close()
		if uploadByIdComplete(fileStore.Id, fileStore.Md5, fileStore.Sha1, userId.(int), staticFilePath) {
			service.SendSuccessJson(ctx, nil, "操作成功！")
		} else {
			service.SendErrorJson(ctx, nil, "数据基础信息不匹配，请重新上传！")
		}
	} else {
		// 读取请求
		for {
			read_len, readErr = ctx.Request.Body.Read(readBuf)
			if readErr != nil && !errors.Is(readErr, io.EOF) {
				if f != nil {
					f.Close()
				}
				return
			}
			jointChar = append(charTemp, readBuf[:read_len]...)

			if !errors.Is(readErr, io.EOF) {
				f.Write(charTemp)
				charTemp = bytes.Clone(readBuf[:read_len])
			} else {
				endIndex := bytes.Index(jointChar, boundaryEnd)
				if endIndex != -1 {
					f.Write(jointChar[:endIndex])
				} else {
					f.Write(jointChar)
				}
				f.Close()
				if uploadByIdComplete(fileStore.Id, fileStore.Md5, fileStore.Sha1, userId.(int), staticFilePath) {
					service.SendSuccessJson(ctx, nil, "操作成功！")
				} else {
					service.SendErrorJson(ctx, nil, "数据基础信息不匹配，请重新上传！")
				}
				break
			}
		}
	}
}

func uploadByIdComplete(storeId int, md5Str string, sha1Str string, userId int, staticFilePath string) bool {
	// type
	f, err := os.Open(staticFilePath)
	if err != nil {
		panic("文件错误！")
	}
	defer func() {
		f.Close()
	}()
	buffer := make([]byte, 261)
	n, _ := f.Read(buffer)
	contentType := http.DetectContentType(buffer[:n])

	// size
	fs, _ := os.Stat(staticFilePath)

	// md5 and sha1
	md5Chan := make(chan string)
	sha1Chan := make(chan string)
	go getFileFeatureCode(staticFilePath, "md5", md5Chan)
	go getFileFeatureCode(staticFilePath, "sha1", sha1Chan)
	newMd5Str := <-md5Chan
	newSHa1Str := <-sha1Chan

	if newMd5Str != md5Str || newSHa1Str != sha1Str {
		deleteFile(staticFilePath)
		return false
	}

	filestoredao.ChangeModifiableData(storeId, contentType, fs.Size(), 1, userId)
	return true
}
