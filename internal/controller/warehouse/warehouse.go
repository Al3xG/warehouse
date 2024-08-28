package warehouse

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateWarehouse(c echo.Context) error {
	return c.String(http.StatusCreated, "/warehouse/ (CreateWarehouse)")
}

func GetWarehouse(c echo.Context) error {
	// get param value
	fmt.Println(c.Param("id"))
	return c.String(http.StatusOK, "/warehouse/:id (GetWarehouse)")
}

func DeleteWarehouse(c echo.Context) error {
	return c.String(http.StatusOK, "/warehouse/:id (DeleteWarehouse)")
}

func EditWarehouse(c echo.Context) error {
	return c.String(http.StatusOK, "/warehouse/:id (EditWarehouse)")
}
