package alert

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmailAlert(server string) {
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	to := []string{"mradzy328@gmail.com"}
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	message := []byte(fmt.Sprintf("Subject: Server Down Alert\r\n\r\nServer %s is DOWN!", server))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	fmt.Println("Alert email sent successfully")
}
