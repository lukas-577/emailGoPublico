package main

import (
	correosgo "correosgo/email"
	"fmt"
)

func main() {
	toEmails := []string{"lukasmedinar@gmail.com", "lukasmedinar@gmail.com"}
	subject := "Subject of the email"
	body := "This is the email body."

	from := correosgo.GetEmailFrom()
	auth := correosgo.PlainAuth()
	message := correosgo.JoinMessageStructure(toEmails, subject, body)

	result := correosgo.SendEmail(correosgo.GetAddressSMTP(), auth, from, toEmails, message)

	fmt.Println(result)
}
