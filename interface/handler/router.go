package handler

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, taskHandler TaskHandler, reviewHandler ReviewHandler) {

	e.POST("/task", taskHandler.Post())
	e.GET("/task/:id", taskHandler.Get())
	e.PUT("/task/:id", taskHandler.Put())
	e.DELETE("/task/:id", taskHandler.Delete())

	// review
	e.POST("/reviews", reviewHandler.Post())
	e.GET("/reviews/:id", reviewHandler.Get())
}
