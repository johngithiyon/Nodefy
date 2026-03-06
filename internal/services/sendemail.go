package services

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to string, otp string) error {


	// Sender Gmail address
	from := os.Getenv("MAIL_FROM")

	// Gmail App Password (16 characters)
	pass := os.Getenv("MAIL_PASS")

	// Receiver email list (Outlook mail)
	recipients := []string{to}

	// SMTP server info for Gmail
	host := "smtp.gmail.com"
	port := "587"

	// SMTP full host string
	address := host + ":" + port

	// SMTP Authentication
	auth := smtp.PlainAuth("Nodefy", from, pass, host)

	// Email message (MUST contain Subject and two newlines)
	message := []byte(
		"Subject: Your OTP for Nodefy\n" +
			"\n" +
			"Your OTP is: " + otp + "\n",
	)

	// Sending the email
	senderr := smtp.SendMail(address, auth, from, recipients, message)
	if senderr != nil {
		fmt.Println("Error sending email:", senderr)
		return senderr
	}

	return senderr
}