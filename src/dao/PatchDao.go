package dao

import (
	"util"
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
