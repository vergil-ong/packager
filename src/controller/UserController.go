package controller

import (
	"github.com/kataras/iris/context"
	"dao"
	"util"
	"github.com/kataras/iris"
	"strconv"
	"strings"
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
	id := ctx.URLParam("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic("parse id int error")
	}

	delRes := dao.DelUser(idInt)
	userMap := iris.Map{
		"del_res":delRes,
	}
	ctx.JSON(util.BuildIrisMap(true, "删除用户信息成功", userMap))
}

func BatchRemoveUsers(ctx context.Context) {
	ids := ctx.FormValue("ids")
	idArray := strings.Split(ids, ",")
	for _,id := range idArray {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			panic("parse id int error")
		}
		delRes := dao.DelUser(idInt)
		if(!delRes){
			//删除失败则退出
			break;
		}
	}

}

func EditUser(ctx context.Context) {
	username := ctx.FormValue("name")
	sex := ctx.FormValue("sex")
	age := ctx.FormValue("age")
	birth := ctx.FormValue("birth")
	addr := ctx.FormValue("addr")
	slug := ctx.FormValue("slug")
	password := ctx.FormValue("password")
	id := ctx.FormValue("id")

	user := dao.BuildUser(id,username, sex, age, birth, addr, slug, password)

	updateRes := dao.UpdateUser(user)

	userMap := iris.Map{
		"update_res":updateRes,
	}
	ctx.JSON(util.BuildIrisMap(true, "修改用户信息成功", userMap))
}

func AdUser(ctx context.Context) {
	username := ctx.FormValue("name")
	sex := ctx.FormValue("sex")
	age := ctx.FormValue("age")
	birth := ctx.FormValue("birth")
	addr := ctx.FormValue("addr")
	slug := ctx.FormValue("slug")
	password := ctx.FormValue("password")

	user := dao.BuildUser("",username, sex, age, birth, addr, slug, password)

	insertUser := dao.InsertUser(user)

	userMap := iris.Map{
		"user":insertUser,
	}
	ctx.JSON(util.BuildIrisMap(true, "插入用户信息成功", userMap))
}
