package handlers

import (
    "database/sql"
    "net/http"

    "github.com/labstack/echo/v4"
    
    "api-go/models"
)

type QuestionFetcher func(*sql.DB) (models.Question, error)

func GetQuestionHandler(fetcher QuestionFetcher, database *sql.DB) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        question, err := fetcher(database)
        if err != nil {
            // Return a 500 Internal Server Error response if there was an error
            return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
        }

        // Construct a response JSON object with the question and answer
        response := map[string]interface{}{
            "id":      question.ID,
            "topic":   question.Topic,
            "question": question.Ques,
            "answer":  question.Ans,
        }

        // Return the response as JSON with a 200 OK status code
        return ctx.JSON(http.StatusOK, response)
    }
}

