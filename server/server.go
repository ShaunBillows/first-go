package server

import (
	"log"

	"github.com/labstack/echo/v4"

	"api-go/handlers"
	"api-go/queries"
	"api-go/db"
)

func Start() {
	e := echo.New()

    database, err := db.Connect()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer database.Close()

	v1Group := e.Group("/v1")

	// GET endpoints
	v1Group.GET("/questions/sql", handlers.GetQuestionHandler(queries.GetSQLQuestion, database))
	v1Group.GET("/questions/go", handlers.GetQuestionHandler(queries.GetGoQuestion, database))
	v1Group.GET("/questions/random", handlers.GetQuestionHandler(queries.GetRandomQuestion, database))
	v1Group.GET("/questions", handlers.GetAllQuestionsHandler(queries.GetAllQuestionsFromDB, database))
	v1Group.GET("/docs", handlers.ApiDocsHandler)

	// POST endpoints
	v1Group.POST("/questions", handlers.AddOneQuestionHandler(queries.AddQuestion, database))
	v1Group.POST("/questions/batch", handlers.AddMultipleQuestionsHandler(queries.AddQuestions, database))
	
	e.Logger.Fatal(e.Start(":1323"))
}
