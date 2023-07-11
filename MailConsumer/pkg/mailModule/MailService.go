package mailModule

import (
	"fmt"
	"mailConsumer/pkg/base"
	"mailConsumer/pkg/repos/implement"
)

type MailService struct{
	base.Service
	MailRepo *implement.MailRepo
}

func NewMailService(mailRepo *implement.MailRepo) *MailService{
	var MailService MailService
	MailService.MailRepo = mailRepo
	return &MailService
}

func(s *MailService) Send(dto *SendDto)(err error){
	if err := dto.Check();err!=nil{
		return s.InvalidArgument(err.Error())
	}
	if err = s.MailRepo.SendMail(dto.Context , dto.To);err!=nil{
		fmt.Println(dto)
		return err
	}

	return nil
}