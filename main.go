package main

import (
	"cal/app"
	"cal/cmd/handlers"
	"cal/db"

	"github.com/labstack/echo/v4"
)

func main() {

	//link to db
	db.InitDB()

	//start app
	app.App()

	e := echo.New()

	//call api to get db
	e.GET("/app", handlers.InD)

	//to direct to root directory
	e.GET("/", handlers.Home)

	//to add new data
	e.POST("/n", handlers.Add)

	//to fetch specific data
	e.GET("/g/:iso_date", handlers.Get)

	//to fetch all data
	e.GET("/ga", handlers.GetA)

	//to update data
	e.PUT("/u/:id", handlers.Up)

	//to delete specific data
	e.DELETE("/d/:iso_date", handlers.Del)

	//to delete all data
	e.DELETE("/da", handlers.DelA)

	e.Logger.Fatal(e.Start(":8080"))
}
