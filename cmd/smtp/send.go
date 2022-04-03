package smtp

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type BumMail struct {
	Smtp SmtpConfig
}

type EmailRequest struct {
	To      string `json:"to"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Image   string `json:"image"`
}

func (b *BumMail) Send(to []string, title, message, imgUrl string) error {
	auth := smtp.PlainAuth("", b.Smtp.From, b.Smtp.Password, b.Smtp.Host)

	t, _ := template.ParseFiles("cmd/smtp/template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", title, mimeHeaders)))

	t.Execute(&body, struct {
		Title    string
		Message  string
		ImageURL string
	}{
		Title:    title,
		Message:  message,
		ImageURL: imgUrl,
	})

	return smtp.SendMail(b.Smtp.Host+":"+b.Smtp.Port, auth, b.Smtp.From, to, body.Bytes())
}
