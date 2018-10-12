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
	"dao"
	"strconv"
)

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

	fileSep := "/"
	localPath := util.FileStorePath+group+fileSep
	os.MkdirAll(util.FileStorePath+group+fileSep,0666)
	fileInfo := dao.BuildFileInfo("", "", localPath, group)
	insertFileInfo, e := dao.InsertFileInfo(fileInfo)
	if e!=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	fileIDStr := strconv.Itoa(int(insertFileInfo.ID))
	localPath = util.FileStorePath+group+fileSep+fileIDStr+fileSep
	os.MkdirAll(localPath,0666)
	localPath = util.FileStorePath+group+fileSep+fileIDStr+fileSep+fname


	out, err := os.OpenFile(localPath,
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
	pageUnlimited := util.BuildPageUnlimited()
	recommendations := dao.ListRecommendationsMaxAppear(fname, pageUnlimited)
	var path string;
	if recommendations != nil && len(recommendations)>0 {
		path = recommendations[0].Path
	}
	//fileInfo = dao.BuildFileInfo("", path, localPath, group)
	//insertFileInfo, e := dao.InsertFileInfo(fileInfo)
	//if e!=nil {
	//	ctx.StatusCode(iris.StatusInternalServerError)
	//	ctx.JSON("Error while uploading: <b>" + err.Error() + "</b>")
	//	return
	//}
	//更新文件存储路径
	fileInfo.ID = insertFileInfo.ID
	fileInfo.TargetPath = path
	//fileInfo.LocalPath = strings.Replace(localPath,".","\\.",-1)
	fileInfo.LocalPath = localPath
	dao.UpdateFileInfo(fileInfo)

	//fmt.Println(insertFileInfo.ID)

	resultMap := iris.Map{
		"fname": fname,
		"group": group,
		"file_path":path,
		"file_info_id":insertFileInfo.ID,
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