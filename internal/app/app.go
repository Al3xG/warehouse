package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"warehouse/internal/route"
)

type App struct {
	echo *echo.Echo
}

func NewApp() *App {
	return &App{echo: echo.New()}
}

func (app *App) Start() {
	app.echo.Use(middleware.Logger())
	route.InitWarehouseRoutes(app.echo)
	route.InitDeviceRoutes(app.echo)
	if err := app.echo.Start(":8080"); err != nil {
		panic(err)
	}
}
