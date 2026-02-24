package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func emailSender() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth(
		"",
		user,
		password,
		host,
	)

	emails := AccessDatabase()
	for _, email := range emails {
		fmt.Println("Sending email to:", email.EmailAddress)
		err := smtp.SendMail(
			fmt.Sprintf("%s:%s", host, port),
			auth,
			user,
			[]string{email.EmailAddress},
			[]byte(fmt.Sprintf("Subject: Hello There!\n\n%s", email.Body)),
		)

		if err == nil {
			UpdateDatabase(email.ID, "send")
		} else if err != nil {
			UpdateDatabase(email.ID, "fail")
		}

	}

}
