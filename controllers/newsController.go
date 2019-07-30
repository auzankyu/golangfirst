package controllers

import (
	"first/models"
	models2 "ginproject/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func All(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var news []models2.News

	db.Table("news").Select("date, content").Scan(&news)
	if len(news) <= 0 {
		error(c)
	}
	success(c, news)
}

func Allv2(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var news models.News
	var arrNews []models.News

	rows, err := db.Table("news").Select("content").Rows()
	defer rows.Close()

	if err != nil {
		error(c)
	}
	for rows.Next() {
		rows.Scan(&news.Content)
		arrNews = append(arrNews, news)
	}
	success(c, arrNews)
}

func Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var (
		news models.News
	)

	err := db.First(&news).Error
	if err != nil {
		error(c)
	}
	success(c, news)
}
