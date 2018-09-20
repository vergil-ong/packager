package dao

import (
	"fmt"
	"time"
	"util"
)

type User struct {
	ID       uint      `gorm:"column:id" json:"id"`
	NAME     string    `gorm:"column:name" json:"name"`
	ADDR     string    `gorm:"column:addr" json:"addr"`
	BIRTH    time.Time `gorm:"column:birth" json:"birth"`
	GENDER   int       `gorm:"column:gender" json:"sex"`
	SLUG     string    `gorm:"column:slug" json:"slug"`
	PASSWORD string    `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "user"
}

func CheckUserPassword(slug string, password string) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var users []User
	db.Where("slug = ? and password = ?", slug, password).Find(&users)

	if len(users) != 0 {
		return true
	} else {
		return false
	}
	defer db.Close()

	return false
}

func ListUsers(slug string, page util.Page) []User {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var users []User
	where := db.Where("slug = ? ", slug,)
	pageDB := where
	if page.GetOffset() != -1 {
		pageDB = where.Offset(page.GetOffset()).Limit(page.GetLimit())
	}
	pageDB.Find(&users)

	defer db.Close()

	return users
}

func DelUser(id int) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	if err := db.Where("id = ?", id).Delete(User{}).Error; err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer db.Close()

	return true
}

func BatchDelUser(ids []int) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var users []User

	if err := db.Where("id in (?)", ids).Delete(&users).Error; err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer db.Close()

	return true
}

func UpdateUser(user User) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	if err = db.Model(&user).Update(&user).Error; err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer db.Close()

	return true
}

func InsertUser(user User) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}
	e := db.Create(&user).Error
	if e != nil {
		err := e.Error()
		fmt.Println(err)
		return false
	}
	defer db.Close()

	return true
}
