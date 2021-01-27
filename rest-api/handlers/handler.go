package handlers

import (
	"database/sql"
	"net/http"

	"github.com/pratamawijaya/Golang101/rest-api/models"

	"github.com/pratamawijaya/Golang101/rest-api/repository"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, repository.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task

		c.Bind(&task)

		id, err := repository.PutTask(db, task.Name, task.Status)

		if err != nil {
			return c.JSON(http.StatusCreated, H{"created": id})
		} else {
			return err
		}

	}
}
