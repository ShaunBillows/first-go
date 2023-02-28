package queries

import (
	"database/sql"
	"fmt"
    
	"api-go/internal/models"
)

func GetAllQuestionsFromDB(db *sql.DB) ([]models.Question, error) {
    questions := []models.Question{}
    query := `SELECT id, topic, question, answer FROM questions`

    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to get questions: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        question := models.Question{}
        if err := rows.Scan(&question.ID, &question.Topic, &question.Ques, &question.Ans); err != nil {
            return nil, fmt.Errorf("failed to scan question row: %v", err)
        }
        questions = append(questions, question)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error while iterating question rows: %v", err)
    }

    return questions, nil
}
