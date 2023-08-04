package main

import (
	"example/config"
	"example/controller"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Connect To Database
	db := config.InitDatabase(os.Getenv("DB_URL"))
	bookController := controller.NewBookController(db)

	bookRoute := e.Group("/books")
	bookRoute.POST("", bookController.Create)
	bookRoute.GET("", bookController.Index)
	bookRoute.GET("/:id", bookController.Detail)
	bookRoute.DELETE("/:id", bookController.Delete)

	e.Logger.Fatal(e.Start(":8000"))
}
