package queries

import (
	"database/sql"
	"fmt"
	
	"api-go/models"
)

func AddQuestions(db *sql.DB, questions []models.Question) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare the SQL statement for inserting a question into the database
	query := `INSERT INTO questions (topic, question, answer) VALUES (?, ?, ?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %v", err)
	}
	defer stmt.Close()

	// Iterate through the questions and execute the prepared statement for each question
	for _, q := range questions {
		_, err = stmt.Exec(q.Topic, q.Ques, q.Ans)
		if err != nil {
			return fmt.Errorf("failed to insert question into database: %v", err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
