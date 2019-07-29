package main

import (
	"ginproject/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := initDB()
	app := gin.Default()
	app.Use(inject(db))

	app.GET("/all", controllers.All)
	app.GET("/", controllers.Get)
	app.Run(":1234")
}
