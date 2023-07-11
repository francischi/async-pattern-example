package common

import (
	"golang/pkg/base"
	"golang/pkg/repos/interfaces"
)

type MailService struct{
	base.Service
	MailRepo interfaces.MailRepo
}

func NewMailService(mailRepo interfaces.MailRepo) *MailService{
	var MailService MailService
	MailService.MailRepo = mailRepo
	return &MailService
}

func (s *MailService) SendOrderConfirmMail(context string , to string)( err error) {
	if err:=s.MailRepo.PushMailIntoQueue(context, to);err!=nil{
		return err
	}
	return nil
}