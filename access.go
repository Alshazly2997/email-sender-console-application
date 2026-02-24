package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Email struct {
	ID           int
	EmailAddress string
	Body         string
	Status       string
}

var emails []Email

func AccessDatabase() []Email {
	var db *sql.DB
	var email Email

	//Connect to the database
	db, err := sql.Open("mysql", "root:death notemysql@tcp(127.0.0.1:3306)/Outbox")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	defer db.Close()

	//Query the database
	query, err := db.Query("SELECT * FROM outbox")
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
