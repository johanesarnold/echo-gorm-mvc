package main

import (
	"echo-gorm-mvc/config"
	"echo-gorm-mvc/db"
	"echo-gorm-mvc/routes"
	"os"

	"github.com/labstack/echo/middleware"
)

func main() {
	configuration := config.GetConfig()
	db.Init()
	e := routes.Init()

	if configuration.APP_DEBUG == "True" {
		e.Debug = true
	}

	//set log file
	f, _ := os.OpenFile("log/echo.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer f.Close()
	loggerConfig := middleware.DefaultLoggerConfig
	loggerConfig.Output = f
	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Use(middleware.Recover())
	e.Use(routes.FirstEncounter)

	e.Logger.Fatal(e.Start(":1234"))
}
