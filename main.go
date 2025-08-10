package main

import (
	"insider/src"
	"insider/src/config"
	"insider/src/infra/middleware"

	_ "insider/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Insider API
// @version 1.0
// @description This is the Insider messaging API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @schemes http https

func main() {
	cfg := config.LoadConfig()

	e := echo.New()

	e.Use(middleware.ExceptionMiddleware)
	e.Use(middleware.CreateLogMiddleware())
	e.Use(middleware.RecoverMiddleware)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	handler := src.NewHandler(cfg)

	e.POST("/messages", handler.MessageController.CreateMessage)
	e.POST("/messages/toggle", handler.MessageController.ToggleMessageService)
	e.GET("/sentmessages", handler.MessageController.GetSentMessages)

	e.GET("/healthcheck", handler.ProbeController.HealthCheck)
	e.GET("/ping", handler.ProbeController.PingPong)

	e.Logger.Fatal(e.Start(":" + cfg.APP_PORT))
}
