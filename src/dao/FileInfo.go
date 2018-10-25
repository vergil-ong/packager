package dao

import (
	"util"
	"strconv"
	"fmt"
)

type FileInfo struct {
	ID 			uint	`gorm:"column:id" json:"id"`
	TargetPath 	string `gorm:"column:target_path" json:"target_path"`
	LocalPath 	string `gorm:"column:local_path" json:"local_path"`
	Group 		string `gorm:"column:group" json:"group"`
	PatchID 	int `gorm:"column:patch_id" json:"patch_id"`
}

func (FileInfo) TableName() string {
	return "file_info"
}

func GetFileInfoByID(id int) FileInfo {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var fileInfo FileInfo
	where := db
	if id != -1 {
		where = where.Where("id = ?", id,)
	}
	where.Find(&fileInfo)

	defer db.Close()

	return fileInfo
}

func GetFileInfosByPatchID(patchID int) []FileInfo {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var fileInfo []FileInfo
	where := db
	if patchID != -1 {
		where = where.Where("patch_id = ?", patchID,)
	}
	where.Find(&fileInfo)

	defer db.Close()

	return fileInfo
}

func ListFileInfos(id int, targetPath string, localPath string, group string, page util.Page) []FileInfo {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var fileInfos []FileInfo
	where := db
	if id != -1 {
		where = where.Where("id = ?", id,)
	}
	if targetPath != "" {
		where = where.Where("target_path = ?", targetPath,)
	}
	if localPath != "" {
		where = where.Where("local_path = ?", localPath,)
	}
	if group != "" {
		where = where.Where("group = ?", group,)
	}
	pageDB := where
	if page.GetOffset() != -1 {
		pageDB = where.Offset(page.GetOffset()).Limit(page.GetLimit())
	}
	pageDB.Find(&fileInfos)

	defer db.Close()

	return fileInfos
}

func InsertFileInfo(fileInfo FileInfo) (FileInfo,error) {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}
	e := db.Create(&fileInfo).Error
	if e != nil {
		err := e.Error()
		fmt.Println(err)
		return FileInfo{},e
	}
	defer db.Close()

	return fileInfo,nil
}

func DelFileInfo(id int) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	if err := db.Where("id = ?", id).Delete(FileInfo{}).Error; err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer db.Close()

	return true
}

func UpdateFileInfo(fileInfo FileInfo) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	if err = db.Model(&fileInfo).Update(&fileInfo).Error; err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer db.Close()

	return true
}

func BuildFileInfo(
	id string,
	tagetPath string,
	localPath string,
	group string,
) FileInfo{

	fileInfo := new(FileInfo)

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err!=nil {
			fmt.Println(err)
			panic("parse id int error")
		}
		fileInfo.ID = uint(idInt)
	}

	fileInfo.TargetPath = tagetPath
	fileInfo.LocalPath = localPath
	fileInfo.Group = group

	return *fileInfo

}