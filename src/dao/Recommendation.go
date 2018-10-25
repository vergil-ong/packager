package dao

import (
	"time"
	"util"
	"strconv"
	"fmt"
	"github.com/kataras/iris/core/errors"
)

type Recommendation struct {
	ID           		uint   `gorm:"column:id" json:"id"`
	FileName         	string `gorm:"column:file_name" json:"file_name"`
	Path    			string `gorm:"column:path" json:"path"`
	AppearTimes 		int `gorm:"column:appear_times" json:"appear_times"`
	Mtime    			time.Time `gorm:"column:mtime" json:"mtime"`
}

func (Recommendation) TableName() string {
	return "file_path_recommendation"
}

func ListRecommendations(fileName string, page util.Page) []Recommendation {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var recommendations []Recommendation
	where := db
	if fileName != "" {
		where = where.Where("file_name = ?", fileName,)
	}
	where.Where("appear_times")
	pageDB := where
	if page.GetOffset() != -1 {
		pageDB = where.Offset(page.GetOffset()).Limit(page.GetLimit())
	}
	pageDB.Find(&recommendations)

	defer db.Close()

	return recommendations
}

func GetRecommendation(fileName string, path string) (*Recommendation,error) {
	db, dbErr := util.GetDBConnection()
	if dbErr != nil {
		panic("connection failure")
	}

	var recommendation Recommendation
	where := db
	if fileName != "" {
		where = where.Where("file_name = ?", fileName,)
	}
	if path != "" {
		where = where.Where("path = ?", path,)
	}
	where.Find(&recommendation)

	defer db.Close()

	var err error
	if &recommendation == nil || recommendation.ID == 0 {
		err = errors.New("cannot find recommendation")
	}

	return &recommendation,err
}

func ListRecommendationsMaxAppear(fileName string, page util.Page) []Recommendation {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	var recommendations []Recommendation

	db.Raw("SELECT tb1.* FROM file_path_recommendation tb1 " +
		"INNER JOIN " +
		"(SELECT file_name,MAX(appear_times) AS m_app FROM file_path_recommendation GROUP BY file_name) tb2 " +
		"ON tb1.file_name = tb2.file_name AND tb1.appear_times = tb2.m_app WHERE tb1.file_name = ? ", fileName).Scan(&recommendations)

	defer db.Close()

	return recommendations
}

func UpdateRecommendation(recommendation *Recommendation) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	if err = db.Model(recommendation).Update(recommendation).Error; err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer db.Close()

	return true
}

func InsertRecommendation(recommendation *Recommendation) bool {
	db, err := util.GetDBConnection()
	if err != nil {
		panic("connection failure")
	}

	e := db.Create(recommendation).Error
	if e != nil {
		err := e.Error()
		fmt.Println(err)
		return false
	}

	defer db.Close()

	return true
}

func BuildRecommendation(
	id string,
	fileName string,
	path string,
	appearTimes string,
) Recommendation{

	recommendation := new(Recommendation)

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err!=nil {
			fmt.Println(err)
			panic("parse id int error")
		}
		recommendation.ID = uint(idInt)
	}

	recommendation.FileName = fileName
	recommendation.Path = path
	if appearTimes!=""{
		appearTimesInt, err := strconv.Atoi(appearTimes)
		if err!=nil {
			fmt.Println(err)
			panic("parse AppearTimes to int error")
		}
		recommendation.AppearTimes = appearTimesInt
	}

	return *recommendation

}