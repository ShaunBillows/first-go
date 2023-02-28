package db

import (
    "database/sql"
	"fmt"
    "os"

    "github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
  )

  func Connect() (*sql.DB, error) {

    err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("failed to load .env file: %v", err)
    }

    dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

    // Open a connection to the database
    db, err := sql.Open("mysql", dbConnectionString)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    // Check the connection
    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("error pinging database: %v", err)
    }

    fmt.Println("Successfully connected to the database!")
    return db, nil
}
