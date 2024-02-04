package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strconv"

	"github.com/IainMcl/HereWeGo/internal/logging"
	"github.com/IainMcl/HereWeGo/internal/settings"
)

type MailService struct {
	Port    int
	Address string
	Auth    *smtp.Auth
}

func New() *MailService {
	auth := smtp.PlainAuth(
		"",
		settings.MailSettings.Username,
		settings.MailSettings.Password,
		settings.MailSettings.Host,
	)
	return &MailService{
		Port:    settings.MailSettings.Port,
		Address: settings.MailSettings.Host + ":" + strconv.Itoa(settings.MailSettings.Port),
		Auth:    &auth,
	}
}

func (m *MailService) SendMail(to, subject, body string) error {
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	err := smtp.SendMail(
		m.Address,
		*m.Auth,
		settings.MailSettings.Username,
		[]string{to},
		[]byte("Subject: "+subject+headers+"\r\n"+body),
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MailService) SendMailTemplate(recipient, templateFile string, data interface{}) error {
	logging.Info(fmt.Sprintf("Sending email to %s using template %s", recipient, templateFile))
	tmpl, err := template.New("email").ParseFiles("internal/mail/templates/" + templateFile)
	if err != nil {
		logging.Error(fmt.Sprintf("Error parsing template: %s", err))
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		logging.Error(fmt.Sprintf("Error executing template subject: %s", err))
		return err
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		logging.Error(fmt.Sprintf("Error executing template plainBody: %s", err))
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		logging.Error(fmt.Sprintf("Error executing template htmlBody: %s", err))
		return err
	}
	err = smtp.SendMail(
		m.Address,
		*m.Auth,
		settings.MailSettings.Username,
		[]string{recipient},
		[]byte("Subject: "+subject.String()+"\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"+htmlBody.String()),
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *MailService) SendGroupMail(to []string, subject, body string) error {
	err := smtp.SendMail(
		m.Address,
		*m.Auth,
		settings.MailSettings.Username,
		to,
		[]byte("Subject: "+subject+"\r\n"+body),
	)
	if err != nil {
		return err
	}
	return nil
}
