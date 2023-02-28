package handlers

import (
	"net/http"
	"database/sql"

	"github.com/labstack/echo/v4"

	"api-go/models"
)

func AddOneQuestionHandler(addQuestionFunc func(*sql.DB, models.Question) error, db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
	// Parse the question data from the request body
	question := new(models.Question)
	if err := c.Bind(question); err != nil {
	return c.JSON(http.StatusBadRequest, echo.Map{"error": "failed to parse request body"})
	}
	// Call the function to add the question to the database
	if err := addQuestionFunc(db, *question); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to add question to database"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "question added to database"})
	}
}
