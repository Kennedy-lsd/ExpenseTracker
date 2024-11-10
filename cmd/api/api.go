package api

import (
	"log"

	"github.com/Kennedy-lsd/ExpenseTracker/database"
	"github.com/Kennedy-lsd/ExpenseTracker/internal/handlers"
	"github.com/Kennedy-lsd/ExpenseTracker/internal/repos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Api() {
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, //Update with your frontend :)
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	repository := repos.NewRepository(db)
	handler := handlers.NewHandler(repository)

	purchaseHandler := app.Group("api")

	purchaseHandler.GET("/purchase", handler.GetAllTasks)
	purchaseHandler.POST("/purchase", handler.CreateTask)
	purchaseHandler.DELETE("/purchase/:id", handler.DeleteTask)

	app.Start(":8080")
}
