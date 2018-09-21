package controller

import (
	"dao"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"util"
)

func HandleLogin(ctx context.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	user,resultBol := dao.CheckUserPassword(username, password)
	if resultBol {
		userMap := iris.Map{
			"user":user,
		}
		//登陆成功
		ctx.JSON(util.BuildIrisMap(true, "登陆成功", userMap))
	} else {
		//登陆失败
		ctx.JSON(util.BuildIrisMap(false,"账号或密码错误",nil))
	}
}
