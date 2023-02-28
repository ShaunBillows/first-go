package queries

import (
    "database/sql"
    "fmt"
    
	"api-go/models"
)

func GetRandomQuestion(db *sql.DB) (models.Question, error) {
    question := models.Question{}
    query := `SELECT id, topic, question, answer FROM questions ORDER BY RAND() LIMIT 1`

    row := db.QueryRow(query)
    err := row.Scan(&question.ID, &question.Topic, &question.Ques, &question.Ans)
    if err != nil {
        return models.Question{}, fmt.Errorf("failed to get question: %v", err)
    }

    return question, nil
}
