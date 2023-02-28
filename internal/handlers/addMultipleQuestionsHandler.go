package handlers

import (
	"net/http"
	"database/sql"

	"github.com/labstack/echo/v4"

	"api-go/internal/models"
)

func AddMultipleQuestionsHandler(addQuestionsFunc func(*sql.DB, []models.Question) error, db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the question data from the request body
		var req struct {
			Questions []models.Question `json:"questions"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "failed to parse request body"})
		}

		// Call the function to add the questions to the database
		if err := addQuestionsFunc(db, req.Questions); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to add questions to database"})
		}

		return c.JSON(http.StatusCreated, echo.Map{"message": "questions added to database"})
	}
}
