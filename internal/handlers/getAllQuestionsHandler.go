package handlers

import (
	"net/http"
	"database/sql"

	"github.com/labstack/echo/v4"

	"api-go/internal/models"
)

func GetAllQuestionsHandler(fetcher func(*sql.DB) ([]models.Question, error), db *sql.DB) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        questions, err := fetcher(db)
        if err != nil {
            // Return a 500 Internal Server Error response if there was an error
            return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
        }

        // Construct a response JSON object with all questions
        response := make([]map[string]interface{}, len(questions))
        for i, question := range questions {
            response[i] = map[string]interface{}{
                "id":      question.ID,
                "topic":   question.Topic,
                "question": question.Ques,
                "answer":  question.Ans,
            }
        }

        // Return the response as JSON with a 200 OK status code
        return ctx.JSON(http.StatusOK, response)
    }
}