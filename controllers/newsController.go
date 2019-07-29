package controllers

import (
	"first/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func All(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var news []models.News

	db.Table("news").Select("date, content").Scan(&news)
	if len(news) <= 0 {
		error(c)
	}
	success(c, news)
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
