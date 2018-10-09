package dao

import (
	"time"
	"util"
	"strconv"
	"fmt"
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