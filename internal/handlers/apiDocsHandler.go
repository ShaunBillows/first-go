package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func ApiDocsHandler(c echo.Context) error {
    docs := []map[string]string{
        {
            "title": "API Overview",
            "description": "This API provides endpoints for accessing and managing resources",
        },
        {
            "title": "Get SQL Question",
            "description": "Get a random SQL question from the database",
            "endpoint": "/v1/questions/sql",
            "method": "GET",
            "parameters": "",
            "response": "{\n    \"question\": \"What is a JOIN statement in SQL?\",\n    \"answer\": \"A JOIN statement is used to combine data from two or more tables based on a related column between them.\"\n}",
        },
        {
            "title": "Get Go Question",
            "description": "Get a random Go programming question from the database",
            "endpoint": "/v1/questions/go",
            "method": "GET",
            "parameters": "",
            "response": "{\n    \"question\": \"What is a pointer in Go?\",\n    \"answer\": \"A pointer is a variable that stores the memory address of another variable.\"\n}",
        },
        {
            "title": "Get Random Question",
            "description": "Get a random question from either the SQL or Go programming category",
            "endpoint": "/v1/questions/random",
            "method": "GET",
            "parameters": "",
            "response": "{\n    \"question\": \"What is a pointer in Go?\",\n    \"answer\": \"A pointer is a variable that stores the memory address of another variable.\"\n}",
        },
        {
            "title": "Get All Questions",
            "description": "Get all questions from the database",
            "endpoint": "/v1/questions",
            "method": "GET",
            "parameters": "",
            "response": "[{\n    \"question\": \"What is a JOIN statement in SQL?\",\n    \"answer\": \"A JOIN statement is used to combine data from two or more tables based on a related column between them.\"\n},\n{\n    \"question\": \"What is a pointer in Go?\",\n    \"answer\": \"A pointer is a variable that stores the memory address of another variable.\"\n}]",
        },
        {
            "title": "API Documentation",
            "description": "Get the API documentation",
            "endpoint": "/v1/docs",
            "method": "GET",
            "parameters": "",
            "response": "Returns the documentation for the API endpoints.",
        },
        {
            "title": "Add One Question",
            "description": "Add a single question to the database",
            "endpoint": "/v1/questions",
            "method": "POST",
            "parameters": "{\n    \"topic\": \"SQL\",\n    \"question\": \"What is a SELECT statement in SQL?\",\n    \"answer\": \"A SELECT statement is used to retrieve data from one or more tables in a database.\"\n}",
            "response": "Returns a success message if the question was added successfully.",
        },
        {
            "title": "Add Multiple Questions",
            "description": "Add multiple questions to the database",
            "endpoint": "/v1/questions/batch",
            "method": "POST",
            "parameters": "[{\n    \"topic\": \"SQL\",\n    \"question\": \"What is a SELECT statement in SQL?\",\n    \"answer\": \"A SELECT statement is used to retrieve data from one or more tables in a database.\"\n},\n{\n    \"topic\": \"Go\",\n    \"question\": \"What is a struct in Go?\",\n    \"answer\": \"A struct in Go is a composite data type that groups together zero or more values with different types. It is a way to define your own custom data types by grouping related values together.\"\n}", "response": "Returns a success message if the questions were added successfully.",
        },
    }
    

    return c.JSON(http.StatusOK, docs)
}
