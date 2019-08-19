package controllers

import (
	"golangapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func All(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var news []models.News

	db.Table("news").Select("date, content").Scan(&news)
	if len(news) <= 0 {
		error(c, "error")
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
		error(c, "error")
	}
	for rows.Next() {
		rows.Scan(&news.Content)
		arrNews = append(arrNews, news)
	}
	success(c, arrNews)
}

func Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var news models.News

	err := db.Where("category_id = ?", id).First(&news).Error
	if err != nil {
		error(c, err)
	}
	success(c, news)
}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var news models.News
	c.BindJSON(&news)
	response := db.Create(&news)
	if response.Error != nil {
		error(c, "gagal menyimpan data")
	} else {
		success(c, news)
	}
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	content := c.Params.ByName("content")
	var news models.News

	if err := db.Where("content = ?", content).First(&news).Error; err != nil {
		error(c, "gagal update")
	}
	c.BindJSON(&news)
	db.Save(&news)
	success(c, news)
}
