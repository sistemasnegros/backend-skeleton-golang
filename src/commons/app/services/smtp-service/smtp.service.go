package smtpService

import (
	smtpDomain "backend-skeleton-golang/commons/domain/smtp"
	gomailInfra "backend-skeleton-golang/commons/infra/gomail"
)

type SmtpService struct {
	mailer gomailInfra.ISmtp
}

type ISmtpService interface {
	Send(smtpDomain.SendArgs) error
}

func New(mailer gomailInfra.ISmtp) ISmtpService {
	return &SmtpService{mailer: mailer}

}

func (s *SmtpService) Send(args smtpDomain.SendArgs) error {
	err := s.mailer.Send(args)
	return err
}
