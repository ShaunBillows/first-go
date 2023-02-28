# Go/MySQL/Echo API

A simple Go/MySQL/Echo API for learning Go.

## Endpoints

- GET /v1/questions/sql: Get a random SQL question from the database.
- GET /v1/questions/go: Get a random Go programming question from the database.
- GET /v1/questions/random: Get a random question from either the SQL or Go programming category.
- GET /v1/questions: Get all questions from the database.
- GET /v1/docs: Get the API documentation.
- POST /v1/questions: Add a single question to the database.
- POST /v1/questions/batch: Add multiple questions to the database.

## Environment Variables

The following environment variable should be set in a .env file::

- `DB_CONNECTION_STRING`: MySQL connection string

## Database Schema

Table Name: questions

| Column Name | Data Type   | Constraints                 |
| ----------- | ----------- | --------------------------- |
| id          | INT         | AUTO_INCREMENT, PRIMARY KEY |
| topic       | VARCHAR(50) | NOT NULL                    |
| question    | TEXT        | NOT NULL                    |
| answer      | TEXT        | NOT NULL                    |

To create the table:

CREATE TABLE questions (
id INT AUTO_INCREMENT PRIMARY KEY,
topic VARCHAR(50) NOT NULL,
question TEXT NOT NULL,
answer TEXT NOT NULL
);
