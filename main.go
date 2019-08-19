package main

import (
	"golangapi/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := initDB()
	app := gin.Default()
	app.Use(inject(db))

	app.GET("/all", controllers.All)
	app.GET("/allv2", controllers.Allv2)
	app.GET("/find/:id", controllers.Get)
	app.POST("/create", controllers.Create)
	app.PUT("/update/:content", controllers.Update)

	app.GET("/api", controllers.GetApi)
	app.GET("/soa", controllers.InquiryBalance)
	app.Run(":1235")
}
