package controller

import (
	"github.com/kataras/iris/context"
	"dao"
	"util"
	"github.com/kataras/iris"
	"strconv"
	"strings"
)

func ListPatchesPages(ctx context.Context) {
	patchName := ctx.URLParam("name")
	page := util.BuildPageUnlimited()

	patches := dao.ListPatches(patchName,page)
	resultMap := iris.Map{
		"patches":patches,
	}
	ctx.JSON(util.BuildIrisMap(true, "获取patch信息成功", resultMap))
}

func RemovePatch(ctx context.Context) {
	id := ctx.URLParam("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic("parse id int error")
	}

	delRes := dao.DelPatch(idInt)
	userMap := iris.Map{
		"del_res":delRes,
	}
	ctx.JSON(util.BuildIrisMap(true, "删除patch信息成功", userMap))
}

func BatchRemovePatches(ctx context.Context) {
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

func EditPatch(ctx context.Context) {
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

func AddPatch(ctx context.Context) {
	patchName := ctx.FormValue("patch_name")
	meta := ctx.FormValue("meta")
	patchShell := ctx.FormValue("patch_shell")
	patchType := ctx.FormValue("patch_type")
	patchVersion := ctx.FormValue("patch_version")

	patch := dao.BuildPatch("", patchName, patchType, patchVersion, meta, patchShell, "")

	insertPatch := dao.InsertPatch(patch)

	resultMap := iris.Map{
		"patch":insertPatch,
	}
	ctx.JSON(util.BuildIrisMap(true, "添加patch成功", resultMap))
}
