package route

import (
	"github.com/labstack/echo/v4"
	"warehouse/internal/controller/device"
)

func InitDeviceRoutes(e *echo.Echo) {
	e.POST("/device", device.CreateDevice)
	e.GET("/device/:id", device.GetDevice)
	e.PUT("/device/:id", device.EditDevice)
	e.DELETE("/device/:id", device.DeleteDevice)
}
