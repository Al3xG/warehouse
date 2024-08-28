package device

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateDevice(c echo.Context) error {
	return c.String(http.StatusCreated, "/device/ (Create device)")
}

func DeleteDevice(c echo.Context) error {
	return c.String(http.StatusOK, "/device/id (Delete device)")
}

func GetDevice(c echo.Context) error {
	// get param value
	fmt.Println(c.Param("id"))
	return c.String(http.StatusOK, "/device/id (Get device)")
}

func EditDevice(c echo.Context) error {
	return c.String(http.StatusOK, "/device/id (Edit device)")
}
