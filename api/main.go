package main

import (
	"sample/config"
	"sample/infra"
	"sample/interface/handler"
	"sample/usecase"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

/** **********************************
 *          Main Process
 ********************************** */
func main() {
	taskRepository := infra.NewTaskRepository(config.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	reviewRepository := infra.NewReviewRepository(config.NewDB())
	reviewUsecase := usecase.NewReviewUsecase(reviewRepository)
	reviewHandler := handler.NewReviewHandler(reviewUsecase)

	e := echo.New()
	handler.InitRouting(e, taskHandler, reviewHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
