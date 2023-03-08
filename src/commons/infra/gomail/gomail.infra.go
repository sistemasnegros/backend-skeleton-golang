package gomailInfra

import (
	configService "backend-skeleton-golang/commons/app/services/config-service"
	logService "backend-skeleton-golang/commons/app/services/log-service"
	smtpDomain "backend-skeleton-golang/commons/domain/smtp"
	"bytes"
	"crypto/tls"
	"path/filepath"
	"text/template"

	"gopkg.in/gomail.v2"
)

type ISmtp interface {
	Send(smtpDomain.SendArgs) error
}

type GomailInfra struct {
	Smtp *gomail.Dialer
}

func New() ISmtp {
	d := gomail.NewDialer(
		configService.GetSmtpHost(),
		configService.GetSmtpPort(),
		configService.GetSmtpUser(),
		configService.GetSmtpPass(),
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &GomailInfra{Smtp: d}
}

func (s *GomailInfra) Send(args smtpDomain.SendArgs) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", configService.GetSmtpFrom())
	msg.SetHeader("To", args.To)
	if args.Cc != "" {
		msg.SetAddressHeader("Cc", args.Cc, "")
	}

	msg.SetHeader("Subject", args.Subject)

	pathTemplate, _ := filepath.Abs("commons/infra/html/emails/" + args.Template)

	templateEmail, errParse := template.ParseFiles(pathTemplate)

	if errParse != nil {
		logService.Error(errParse.Error())
		return errParse
	}

	var html bytes.Buffer

	templateEmail.Execute(&html, args.Data)

	msg.SetBody("text/html", string(html.Bytes()))

	errSmtp := s.Smtp.DialAndSend(msg)

	if errSmtp != nil {
		logService.Error(errSmtp.Error())
	}

	logService.Info("email sent successfully: " + args.To)

	return errSmtp

}
