package userservice

import (
	"NetworkDisk/dao/loginlogdao"
	"NetworkDisk/dao/userdao"
	"NetworkDisk/service"
	"NetworkDisk/utils/verifyuser"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	// 参数查找
	if len(account) == 0 || len(password) == 0 {
		service.SendErrorJson(ctx, nil, "账户名或者密码缺失！")
		return
	}

	user, err := userdao.GetByAccount(account)
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	// 账密校验
	if errors.Is(err, gorm.ErrRecordNotFound) {
		service.SendErrorJson(ctx, nil, "该账户不存在！")
		return
	} else if password != user.Password {
		service.SendErrorJson(ctx, nil, "密码错误！")
		return
	}

	// 返回登陆成功并修改cookie
	ip := ctx.ClientIP()
	token, err := verifyuser.EncodeUser(verifyuser.UserMessage{Id: user.Id, Account: account,
		Ip: ip, CreatedAt: time.Now().Unix()})
	if err != nil {
		panic("系统生成token失败，请联系开发人员处理！")
	}
	ctx.SetCookie("token", base64.StdEncoding.EncodeToString(token), 0, "/", "localhost", false, true)

	service.SendSuccessJson(ctx, struct {
		Account     string `json:"account"`
		Name        string `json:"name"`
		Cover       string `json:"cover"`
		MaxCapacity int64  `json:"maxCapacity"`
		NowCapacity int64  `json:"nowCapacity"`
	}{user.Account, user.Name, user.Cover, user.MaxCapacity, user.NowCapacity}, "登陆成功！")

	go loginlogdao.AddOne(user.Id, account, ip, "web")
}

func SignIn(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	cover := ctx.PostForm("cover")

	// 参数校对
	if len(account) == 0 || len(password) == 0 {
		service.SendErrorJson(ctx, nil, "信息校验失败！")
		return
	}
	user, err := userdao.GetByAccount(account)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		service.SendErrorJson(ctx, nil, "当前账号重复！")
		return
	}

	user, err = userdao.AddOne(account, password, cover, 52428800)

	if err != nil {
		service.SendErrorJson(ctx, nil, "数据库添加失败！请联系开发人员处理！")
		return
	}

	// 生成token
	ip := ctx.ClientIP()
	token, err := verifyuser.EncodeUser(verifyuser.UserMessage{Id: user.Id, Account: account,
		Ip: ip, CreatedAt: time.Now().Unix()})
	if err != nil {
		panic("系统生成token失败，请联系开发人员处理！")
	}
	ctx.SetCookie("token", base64.StdEncoding.EncodeToString(token), 0, "/", "localhost", false, true)

	service.SendSuccessJson(ctx, struct {
		Account     string `json:"account"`
		Name        string `json:"name"`
		Cover       string `json:"cover"`
		MaxCapacity int64  `json:"maxCapacity"`
		NowCapacity int64  `json:"nowCapacity"`
	}{user.Account, user.Name, user.Cover, user.MaxCapacity, user.NowCapacity}, "注册成功！")

	go loginlogdao.AddOne(user.Id, account, ip, "web")
}
