package controller

import (
	"github.com/kataras/iris/context"
	"dao"
	"util"
	"github.com/kataras/iris"
	"strconv"
)

func ListUsers(ctx context.Context) {
	//username := ctx.PostValue("username")
	username := ctx.URLParam("name")
	page := util.BuildPageUnlimited()
	users := dao.ListUsers(username,page)
	userMap := iris.Map{
		"users":users,
	}
	ctx.JSON(util.BuildIrisMap(true, "获取用户信息成功", userMap))
}

func ListUserPages(ctx context.Context) {
	username := ctx.URLParam("name")
	pageStr := ctx.URLParam("page")
	pageInt,err := strconv.Atoi(pageStr)
	if err != nil {
		panic("parse page int error")
	}
	if pageInt >= 1 {
		pageInt = pageInt - 1
	}
	page := util.BuildPage(20,pageInt)
	users := dao.ListUsers(username,page)
	userMap := iris.Map{
		"users":users,
	}
	ctx.JSON(util.BuildIrisMap(true, "获取用户信息成功", userMap))
}

func RemoveUser(ctx context.Context) {

}

func BatchRemoveUsers(ctx context.Context) {

}

func EditUser(ctx context.Context) {

}

func AdUser(ctx context.Context) {
	username := ctx.FormValue("name")
	sex := ctx.URLParam("sex")
	pageInt,err := strconv.Atoi(pageStr)
	if err != nil {
		panic("parse page int error")
	}
	if pageInt >= 1 {
		pageInt = pageInt - 1
	}
	page := util.BuildPage(20,pageInt)
	users := dao.ListUsers(username,page)
	userMap := iris.Map{
		"users":users,
	}
	ctx.JSON(util.BuildIrisMap(true, "获取用户信息成功", userMap))
}
