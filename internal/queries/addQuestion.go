package queries

import (
	"database/sql"
	"fmt"
    
	"api-go/internal/models"
)

func AddQuestion(db *sql.DB, question models.Question) error {
    // Prepare the SQL statement for inserting a question into the database
    query := `INSERT INTO questions (topic, question, answer) VALUES (?, ?, ?)`
    stmt, err := db.Prepare(query)
    if err != nil {
        return fmt.Errorf("failed to prepare insert statement: %v", err)
    }
    defer stmt.Close()

    // Execute the prepared statement with the values of the new question
    _, err = stmt.Exec(question.Topic, question.Ques, question.Ans)
    if err != nil {
        return fmt.Errorf("failed to insert question into database: %v", err)
    }

    return nil
}
