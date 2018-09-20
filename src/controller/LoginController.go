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
	resultBol := dao.CheckUserPassword(username, password)
	if resultBol {
		page := util.BuildPageUnlimited()
		users := dao.ListUsers(username, page)
		userMap := iris.Map{
			"user":users[0],
		}
		//登陆成功
		ctx.JSON(util.BuildIrisMap(true, "登陆成功", userMap))
	} else {
		//登陆失败
		ctx.JSON(util.BuildIrisMap(false,"账号或密码错误",nil))
	}
}
