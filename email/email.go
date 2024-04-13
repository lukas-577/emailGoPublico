package correosgo

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("email/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func SendEmail(addr string, auth smtp.Auth, from string, to []string, message []byte) string {
	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		log.Printf("Error!: %s", err)
		return "Error! sending mail"
	}
	return "email sent successfully"
}

func GetAddressSMTP() string {
	return os.Getenv("address")
}

func GetEmailFrom() string {
	return os.Getenv("username")
}

func PlainAuth() smtp.Auth {
	return smtp.PlainAuth(os.Getenv("identity"), os.Getenv("username"), os.Getenv("password"), os.Getenv("host"))
}

func JoinMessageStructure(emailList []string, subject string, body string) []byte {
	concatenate := strings.Join(emailList, ", ")
	toMsg := fmt.Sprintf("To: %v\r\n", concatenate)

	subjectMsg := fmt.Sprintf("Subject: %v\r\n", subject)

	bodyMsg := fmt.Sprintf("%v\r\n", body)

	return []byte(toMsg + subjectMsg + "\r\n" + bodyMsg)
}
