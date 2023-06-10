package uploadservice

import (
	"NetworkDisk/dao/filestoredao"
	"NetworkDisk/service"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func BaseUpload(ctx *gin.Context) {
	// 根据账号确定该上传是否合法
	userId, ok := ctx.Get("userId")
	if !ok {
		service.SendErrorJson(ctx, nil, "错误用户！")
		return
	}
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
	boundaryFirst := append(append([]byte{45, 45}, boundary...), []byte{13, 10}...)
	// 多文件分隔符
	boundaryBase := append([]byte{13, 10}, boundaryFirst...)
	boundaryEnd := append(append([]byte{13, 10, 45, 45}, boundary...), []byte{45, 45, 13, 10}...)

	// 读取配置
	readBuf := make([]byte, 1024000)
	var jointChar []byte
	read_len, readErr := ctx.Request.Body.Read(readBuf)
	headEndIndex := 0   // 请求报文头结束位置
	var charTemp []byte // 暂时存储
	staticFolderName := time.Now().Format("20060102")
	staticFileName := uuid.NewV4().String()
	staticFolderPath := path.Join("file", staticFolderName)
	staticFilePath := path.Join(staticFolderPath, staticFileName)
	if os.MkdirAll(staticFolderPath, os.ModePerm) != nil {
		panic("文件夹创建错误！")
	}
	f, _ := os.Create(staticFilePath)

	// 非文件结束错误，暂时不管
	if readErr != nil && !errors.Is(readErr, io.EOF) {
		if f != nil {
			f.Close()
			deleteFile(staticFilePath)
		}
		return
	}

	// 文件名称
	// var sourceFileName string = ""

	headEndIndex = bytes.Index(readBuf[:read_len], []byte{13, 10, 13, 10})
	// sourceFileName = parseFileName(readBuf[len(boundaryFirst):headEndIndex])
	jointChar = bytes.Clone(readBuf[headEndIndex+4 : read_len])
	charTemp = bytes.Clone(jointChar)

	if errors.Is(readErr, io.EOF) {
		for {
			bIndex := bytes.Index(jointChar, boundaryBase)
			if bIndex == -1 {
				endIndex := bytes.Index(jointChar, boundaryEnd)
				if endIndex == -1 {
					f.Write(jointChar)
				} else {
					f.Write(jointChar[:endIndex])
				}
				f.Close()
				loggerUpload(staticFolderName, staticFileName, userId.(int), staticFilePath)
				break
			} else {
				f.Write(jointChar[:bIndex])
				f.Close()
				loggerUpload(staticFolderName, staticFileName, userId.(int), staticFilePath)

				jointChar = jointChar[bIndex:]

				headEndIndex = bytes.Index(readBuf[:read_len], []byte{13, 10, 13, 10})
				// sourceFileName = parseFileName(readBuf[len(boundaryBase):headEndIndex])
				staticFileName = uuid.NewV4().String()
				staticFilePath = path.Join("file", staticFolderName, staticFileName)
				f, _ = os.Create(staticFilePath)
				jointChar = jointChar[headEndIndex+4:]
			}
		}
	} else {

		// 读取请求
		for {
			read_len, readErr = ctx.Request.Body.Read(readBuf)
			if readErr != nil && !errors.Is(readErr, io.EOF) {
				if f != nil {
					f.Close()
					deleteFile(staticFilePath)
				}
				return
			}
			jointChar = append(charTemp, readBuf[:read_len]...)

			// 这个循环用来寻找间隔符
			for {
				bIndex := bytes.Index(jointChar, boundaryBase)
				if bIndex == -1 {
					if errors.Is(readErr, io.EOF) {
						endIndex := bytes.Index(jointChar, boundaryEnd)
						if endIndex != -1 {
							f.Write(jointChar[:endIndex])
						} else {
							f.Write(jointChar)
						}
						f.Close()
						loggerUpload(staticFolderName, staticFileName, userId.(int), staticFilePath)
					} else {

						if jointCharLen := len(jointChar); jointCharLen > read_len {
							f.Write(jointChar[:jointCharLen-read_len])
							charTemp = bytes.Clone(readBuf[:read_len])
						} else {
							charTemp = bytes.Clone(jointChar)
						}
					}
					break
				} else {
					headEndIndex = bytes.Index(jointChar[bIndex:], []byte{13, 10, 13, 10})
					if headEndIndex == -1 {
						charTemp = bytes.Clone(jointChar)
						break
					}
					f.Write(jointChar[:bIndex])
					f.Close()
					loggerUpload(staticFolderName, staticFileName, userId.(int), staticFilePath)

					jointChar = jointChar[bIndex:]

					// sourceFileName = parseFileName(readBuf[len(boundaryBase):headEndIndex])
					staticFileName = uuid.NewV4().String()
					staticFilePath = path.Join("file", staticFolderName, staticFileName)
					f, _ = os.Create(staticFilePath)
					jointChar = jointChar[headEndIndex+4:]
				}
			}
			if errors.Is(readErr, io.EOF) {
				break
			}
		}
	}
}

// 解析数据报文头，获取文件名称
// func parseFileName(source []byte) string {
// 	fileName := ""
// 	headArr := bytes.Split(source, []byte{13, 10})

// 	for _, headItem := range headArr {
// 		headItemKeyValueArr := strings.Split(string(headItem), ":")
// 		if strings.ToLower(headItemKeyValueArr[0]) == "content-disposition" {
// 			_, params, err := mime.ParseMediaType(headItemKeyValueArr[1])
// 			if err != nil {
// 				panic("数据解析错误！联系开发人员处理！")
// 			}

// 			fileName = params["filename"]
// 			break
// 		}
// 	}

// 	return fileName
// }

func loggerUpload(folder string, file string, createdUser int, filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic("文件错误！")
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	buffer := make([]byte, 512)
	n, _ := f.Read(buffer)
	contentType := http.DetectContentType(buffer[:n])
	f.Close()

	fs, _ := os.Stat(filePath)

	f, _ = os.Open(filePath)
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		panic("系统解析文件错误！")
	}
	md5str := fmt.Sprintf("%x", md5hash.Sum(nil))
	f.Close()

	f, _ = os.Open(filePath)
	sha1hash := sha1.New()
	if _, err := io.Copy(sha1hash, f); err != nil {
		panic("系统解析文件错误！")
	}
	sha1str := fmt.Sprintf("%x", sha1hash.Sum(nil))
	f.Close()

	// 数据库中未查找到相同文件特征
	_, resultErr := filestoredao.GetByMd5AndSha1(md5str, sha1str)
	if errors.Is(resultErr, gorm.ErrRecordNotFound) {
		filestoredao.Add(folder, file, contentType, fs.Size(), md5str, sha1str, createdUser, 1)
	} else {
		deleteFile(filePath)
	}
}

func deleteFile(filePath string) {
	os.Remove(filePath)
	folderPath, _ := path.Split(filePath)
	if files, _ := os.ReadDir(folderPath); len(files) == 0 {
		os.Remove(folderPath)
	}
}
