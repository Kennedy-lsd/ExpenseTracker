package handlers

import (
	"net/http"
	"strconv"

	"github.com/Kennedy-lsd/ExpenseTracker/data"
	"github.com/Kennedy-lsd/ExpenseTracker/internal/repos"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Repository *repos.Repository
}

func NewHandler(r *repos.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h *Handler) GetAllTasks(c echo.Context) error {
	tasks, err := h.Repository.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(tasks) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Task not found",
		})
	}

	return c.JSON(http.StatusOK, tasks)

}

func (h *Handler) CreateTask(c echo.Context) error {
	newTask := new(data.SetPurchase)

	if err := c.Bind(newTask); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input: "+err.Error())
	}

	err := h.Repository.Create(newTask)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating task: "+err.Error())
	}

	return c.JSON(http.StatusCreated, newTask)
}

func (h *Handler) DeleteTask(c echo.Context) error {
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Id param"})
	}
	deletedTask := h.Repository.Delete(id)

	if deletedTask != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Not Found"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Stock was deleted"})
}
