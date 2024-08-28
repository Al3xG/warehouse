package main

import App "warehouse/internal/app"

func main() {
	app := App.NewApp()
	app.Start()

	//e := echo.New()
	//e.Use(middleware.Logger()) // 👈 log all requests
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.Logger.Fatal(e.Start(":1323"))
}
