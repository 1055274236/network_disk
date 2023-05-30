package userservice

import (
	"NetworkDisk/config"
	"NetworkDisk/dao/userdao"
	"NetworkDisk/service"
	"NetworkDisk/utils/verifyuser"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	// 参数查找
	if len(account) == 0 || len(password) == 0 {
		service.SendErrorJson(ctx, nil, "账户名或者密码缺失！")
		return
	}

	user, rows := userdao.GetByAccount(ctx, account)
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	fmt.Println("password: ", password)
	// 账密校验
	if rows == 0 {
		service.SendErrorJson(ctx, nil, "该账户不存在！")
		return
	} else if password != user.Password {
		service.SendErrorJson(ctx, nil, "密码错误！")
		return
	}

	// 返回登陆成功并修改cookie
	ip := ctx.ClientIP()
	token := verifyuser.EncodeUser(verifyuser.UserMessage{Id: int64(user.Id), Account: account,
		Ip: ip, Ext: time.Now().Unix() + config.GlobalConfig.Gin.Login.Ext})
	ctx.SetCookie("token", base64.StdEncoding.EncodeToString(token), 0, "/", "localhost", false, true)

	service.SendSuccessJson(ctx, "登陆成功！")
}
