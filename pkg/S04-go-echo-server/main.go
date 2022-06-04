package main

import (
	"github.com/wilsonify/S04-go-echo-server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Strength - summary: signal strength
	e.POST("/strength", c.Strength)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
