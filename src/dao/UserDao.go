package dao

import (
	"fmt"
	"time"
	"util"
	"strconv"
)

type User struct {
	ID       uint      `gorm:"column:id" json:"id"`
	NAME     string    `gorm:"column:name" json:"name"`
	ADDR     string    `gorm:"column:addr" json:"addr"`
	BIRTH    time.Time `gorm:"column:birth" json:"birth"`
	GENDER   int       `gorm:"column:gender" json:"sex"`
	SLUG     string    `gorm:"column:slug" json:"slug"`
	PASSWORD string    `gorm:"column:password" json:"password"`
	AGE      int       `gorm:"column:age" json:"age"`
}

const DEFAULT_SEX_STR string  = "0"
const DEFAULT_AGE_STR string  = "18"
const DEFAULT_BIRTH_STR string  = "2018-09-21"
const DEFAULT_TIME_FORMAT string  = "2006-01-02"

func (User) TableName() string {
	return "user"
}

func CheckUserPassword(slug string, password string) (User,bool) {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var nilUser User
	var users []User
	db.Where("slug = ? and password = ?", slug, password).Find(&users)

	if len(users) != 0 {
		return users[0],true
	} else {
		return nilUser,false
	}
	defer db.Close()

	return nilUser,false
}

func ListUsers(slug string, page util.Page) []User {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var users []User
	where := db
	fmt.Println(slug)
	if slug != "" {
		where = where.Where("name like ?", "%"+slug+"%",)
	}
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

func BuildUser(
	id string,
	username string,
	sex string,
	age string,
	birth string,
	addr string,
	slug string,
	password string) User{

	user := new(User)
	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err!=nil {
			fmt.Println(err)
			panic("parse id int error")
		}
		user.ID = uint(idInt)
	}

	user.NAME = username
	if sex == "" {
		sex = DEFAULT_SEX_STR
	}
	sexInt, err := strconv.Atoi(sex)
	if err!=nil {
		fmt.Println(err)
		panic("parse sex int error")
	}
	user.GENDER = sexInt
	if age == "" {
		age = DEFAULT_AGE_STR
	}
	ageInt, err := strconv.Atoi(age)
	if err!=nil {
		fmt.Println(err)
		panic("parse age int error")
	}
	user.AGE = ageInt
	if birth == "" {
		birth = DEFAULT_BIRTH_STR
	}
	birthTime, err := time.Parse(DEFAULT_TIME_FORMAT, birth)
	if err != nil {
		fmt.Println(err)
		panic("parse birth to time error")
	}
	user.BIRTH = birthTime
	user.ADDR = addr
	user.SLUG = slug
	user.PASSWORD = password
	return *user
}