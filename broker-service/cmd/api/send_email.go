package main

import (
	"log"
	"net/smtp"
)

func send_mail(body string) {
	from := "ha>>>@gmail.com"
	pass := "biln xpgz uzgp sgjr"
	to := "val??@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("email sent")
}
