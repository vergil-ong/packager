package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"controller"
)

func HandlePingGet(ctx context.Context) {
	testString := ctx.Params().Get("test")
	fmt.Println(testString)
	ctx.JSON(iris.Map{
		"message": "pong",
	})
}

func main() {
	app := iris.Default()
	baseUrl := "/api"
	app.Get(baseUrl+"/ping", HandlePingGet)
	app.Post(baseUrl+"/login", controller.HandleLogin)
	//user
	app.Get(baseUrl+"/user/list", controller.ListUsers)
	app.Get(baseUrl+"/user/listpage", controller.ListUserPages)
	app.Delete(baseUrl+"/user/remove", controller.RemoveUser)
	app.Post(baseUrl+"/user/batchremove", controller.BatchRemoveUsers)
	app.Post(baseUrl+"/user/edit", controller.EditUser)
	app.Put(baseUrl+"/user/add", controller.AdUser)

	//patcher
	patcher := app.Party(baseUrl+"/patch")
	{
		patcher.Get("/list_page",controller.ListPatchesPages)
		patcher.Delete("/remove",controller.RemovePatch)
		//patcher.Delete("batch_remove")
		//patcher.Put("edit")
		patcher.Post("/add",controller.AddPatch)
	}
	// listen and serve on http://0.0.0.0:8080.

	app.Post("/meta/Upload",controller.Upload)
	app.Options("/meta/Upload", func(ctx context.Context) {
		ctx.StatusCode(204)
		ctx.Header("access-control-allow-credentials","true")
		ctx.Header("access-control-allow-headers","content-type")
		ctx.Header("access-control-allow-methods","GET,HEAD,PUT,PATCH,POST,DELETE")
		ctx.Header("Access-Control-Allow-Origin","*")
	})

	app.Run(iris.Addr(":8652"))
}
