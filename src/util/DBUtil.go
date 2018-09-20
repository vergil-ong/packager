package util

import (
	"github.com/jinzhu/gorm"
)

var db_username string = "root"
var db_password string = "root"
var db_host string = "127.0.0.1"
var db_port = "3306"
var db_name = "packager"

func GetDBConnection() (db *gorm.DB, err error)  {
	return gorm.Open("mysql", db_username+":"+db_password+"@tcp("+db_host+":"+db_port+")/"+db_name+"?charset=utf8&parseTime=True&loc=Local")
}

type Page struct {
	limit int
	offset int
}

func (page *Page) GetLimit() int {
	return page.limit
}
func (page *Page) GetOffset() int {
	return page.offset
}
func (page *Page) SetLimit(limit int) {
	page.limit = limit
}
func (page *Page) SetOffset(offset int) {
	page.offset = offset
}
func BuildPageUnlimited() Page  {
	var page = new(Page)
	page.limit = -1
	page.offset = 20
	return  *page
}

func BuildPage(limit int , offset int) Page  {
	var page = new(Page)
	page.limit = limit
	page.offset = offset
	return  *page
}

