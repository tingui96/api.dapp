package lib

import (
	"github.com/ic-matcom/api.dapp/schema/dto"
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
)

func SendSingleMessage(to, subject, message string) error {
	d := gomail.Dialer{Host: "smtp.gmail.com", Port: 587, Username:"user@gmail.com", Password: "password aqui"}
	// config used in postfix local
	// d := gomail.Dialer{Host: "localhost", Port: 25}
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	s, err := d.Dial()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", "user@gmail.com")
	m.SetAddressHeader("To", to, "")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	if err := gomail.Send(s, m); err != nil {
		return err
	}
	m.Reset()

	return nil
}

//nolint:unparam
func SendToRecipients(recipients []dto.User, subject, message string) error {
	mail := gomail.NewMessage()

	addresses := make([]string, len(recipients))
	for i, recipient := range recipients {
		addresses[i] = mail.FormatAddress(recipient.Username, recipient.Name)
	}

	// begin: template
	t := template.New("templateRecipients.html")
	t, err := t.ParseFiles("templateRecipients.html")
	if err != nil {
		return err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		return err
	}
	// end: template

	mail.SetHeader("From", "user@gmail.com")
	mail.SetHeader("To", addresses...)
	mail.SetHeader("Subject", subject)
	//mail.SetHeader("Subject", "votación")
	//mail.SetBody("text/html", message)
	mail.SetBody("text/html", tpl.String())
	mail.Attach("manual_de_acceso.pdf")

	dialer := gomail.Dialer{Host: "smtp.gmail.com", Port: 587, Username:"user@gmail.com", Password: ""}
	if err := dialer.DialAndSend(mail); err != nil {
		return err
	}
	mail.Reset()
	return nil
}
