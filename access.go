package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Email struct {
	ID           int
	EmailAddress string
	Body         string
	Status       string
}

var emails []Email

func AccessDatabase() sql.DB {
	var db *sql.DB

	error := godotenv.Load()
	if error != nil {
		fmt.Println("Error loading .env file:", error)
	}

	password := os.Getenv("DATABASE_PASSWORD")

	//Connect to the database
	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/Outbox", password)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	defer db.Close()

	return *db
}

func QueryDatabase() []Email {
	db := AccessDatabase()
	var email Email

	//Query the database
	query, err := db.Query("SELECT * FROM outbox WHERE status = 'pending'")
	if err != nil {
		fmt.Println("Error querying database:", err)
	}

	for query.Next() {

		err = query.Scan(&email.ID, &email.EmailAddress, &email.Body, &email.Status)
		if err != nil {
			fmt.Println("Error scanning query results:", err)
		}

		emails = append(emails, email)
	}

	return emails
}

func UpdateDatabase(id int, status string) {
	db := AccessDatabase()

	//Update the database
	query, err := db.Exec("UPDATE outbox SET status = ? WHERE id = ?", status, id)
	if err != nil {
		fmt.Println("Error querying database:", err)
	}

	fmt.Println(query)
}
