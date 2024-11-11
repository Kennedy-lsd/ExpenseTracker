package api

import (
	"log"

	"github.com/Kennedy-lsd/ExpenseTracker/database"
	"github.com/Kennedy-lsd/ExpenseTracker/internal/handlers"
	"github.com/Kennedy-lsd/ExpenseTracker/internal/repos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Api() error {
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
		return err
	}

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, //Update with your frontend :)
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	app.Use(middleware.Secure())

	app.Use(middleware.BodyLimit("10M"))

	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(500)))

	app.Use(middleware.Gzip())

	repository := repos.NewRepository(db)
	handler := handlers.NewHandler(repository)

	purchaseHandler := app.Group("api")

	purchaseHandler.GET("/purchase", handler.GetAllTasks)
	purchaseHandler.POST("/purchase", handler.CreateTask)
	purchaseHandler.DELETE("/purchase/:id", handler.DeleteTask)
	purchaseHandler.PATCH("/purchase/:id", handler.UpdateTask)

	app.Start(":8080")

	return nil
}
