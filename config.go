package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		"root:R4h4514@tcp(warazo.id:3306)/news?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}

func inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
