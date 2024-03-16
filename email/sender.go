package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type Message struct {
	Title string
	Value string
}

func Send(data []Message, email string) error {
	from := "tivetind23@gmx.com"
	password := "Pu*zYNSyV2C2cv2"

	to := []string{
		email,
	}

	smtpHost := "mail.gmx.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("./email/template.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("From: tivetind23@gmx.com\r\nTo: tivetind23@gmail.com\r\nSubject: This is a test subject \r\n \r\n%s \r\n", mimeHeaders)))

	err = t.Execute(&body, data)
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}
