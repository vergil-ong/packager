package main

import (
	"controller"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
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
	app.Get(baseUrl+"/user/list", controller.ListUsers)
	app.Get(baseUrl+"/user/listpage", controller.ListUserPages)
	app.Post(baseUrl+"/user/remove", controller.RemoveUser)
	app.Post(baseUrl+"/user/batchremove", controller.BatchRemoveUsers)
	app.Post(baseUrl+"/user/edit", controller.EditUser)
	app.Post(baseUrl+"/user/add", controller.AdUser)
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8652"))
}
