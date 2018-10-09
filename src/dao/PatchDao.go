package dao

import (
	"util"
	"strconv"
	"fmt"
)

type Patch struct {
	ID 	uint	`gorm:"column:id" json:"id"`
	NAME string `gorm:"column:name" json:"patch_name"`
	PatchType string `gorm:"column:patch_type" json:"patch_type"`
	PatchVersion string `gorm:"column:patch_version" json:"patch_version"`
	PatchMeta string `gorm:"column:patch_meta" json:"patch_meta"`
	PatchShell string `gorm:"column:patch_shell" json:"patch_shell"`
	PatchFile string `gorm:"column:patch_file" json:"patch_file"`
}

func (Patch) TableName() string {
	return "patch"
}

func ListPatches(patchName string, page util.Page) []Patch {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var patches []Patch
	where := db
	if patchName != "" {
		where = where.Where("name like ?", "%"+patchName+"%",)
	}
	pageDB := where
	if page.GetOffset() != -1 {
		pageDB = where.Offset(page.GetOffset()).Limit(page.GetLimit())
	}
	pageDB.Find(&patches)

	defer db.Close()

	return patches
}

func InsertPatch(patch Patch) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}
	e := db.Create(&patch).Error
	if e != nil {
		err := e.Error()
		fmt.Println(err)
		return false
	}
	defer db.Close()

	return true
}

func DelPatch(id int) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	if err := db.Where("id = ?", id).Delete(Patch{}).Error; err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer db.Close()

	return true
}

func BuildPatch(
	id string,
	name string,
	patchType string,
	patchVersion string,
	patchMeta string,
	patchShell string,
	patchFile string,
) Patch{

	patch := new(Patch)

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err!=nil {
			fmt.Println(err)
			panic("parse id int error")
		}
		patch.ID = uint(idInt)
	}

	patch.NAME = name
	patch.PatchType = patchType
	patch.PatchVersion = patchVersion
	patch.PatchMeta = patchMeta
	patch.PatchShell = patchShell
	patch.PatchFile = patchFile

	return *patch

}