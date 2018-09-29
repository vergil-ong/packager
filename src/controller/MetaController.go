package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"mime/multipart"
	"strings"
	"os"
	"io"
	"fmt"
	"util"
)

const UPLOAD_DIR string = "D:/uploads/"

func Upload(ctx context.Context) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		fmt.Println(err.Error())
		return
	}

	defer file.Close()
	fname := info.Filename

	group := ctx.FormValue("group")
	if group == "" {
		group = util.GenerateSerial()
	}

	out, err := os.OpenFile(UPLOAD_DIR+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)

	/*resultMap := iris.Map{
		"patches":patches,
	}*/
	resultMap := iris.Map{
		"fname": fname,
		"group": group,
	}
	ctx.Header("access-control-allow-origin","*")
	ctx.JSON(resultMap)
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=	time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}