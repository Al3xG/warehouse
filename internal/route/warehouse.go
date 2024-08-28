package route

import (
	"github.com/labstack/echo/v4"
	"warehouse/internal/controller/warehouse"
)

func InitWarehouseRoutes(e *echo.Echo) {
	e.POST("/warehouse", warehouse.CreateWarehouse)
	e.GET("/warehouse/:id", warehouse.GetWarehouse)
	e.PUT("/warehouse/:id", warehouse.EditWarehouse)
	e.DELETE("/warehouse/:id", warehouse.DeleteWarehouse)
}
