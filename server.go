package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/** **********************************
 *          Main Process
 ********************************** */
func main() {
	router := echo.New()

	initRouting(router)

	router.Logger.Fatal(router.Start(":1323"))
}

func initRouting(router *echo.Echo) {
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
