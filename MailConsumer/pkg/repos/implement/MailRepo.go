package implement

import (
	// "fmt"
	"mailConsumer/pkg/helpers"
	"mailConsumer/pkg/base"
	"github.com/alexcesaro/mail/gomail"
)

type MailRepo struct {
	base.Repository
}

func NewMailRepo() *MailRepo{
	return &MailRepo{}
}

func (m *MailRepo)SendMail(context string, to string)(err error){
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", helpers.GetEnvStr("mail.sendaccount") , helpers.GetEnvStr("mail.name") )
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Order Confirmed!")
	msg.AddAlternative("text/html", "Hello! "+context)

	mailer := gomail.NewMailer("smtp.gmail.com", "francischiooo@gmail.com", "znvtosircworjbrf", 587)
	if err := mailer.Send(msg); err != nil {
		return m.SystemError(err.Error())
	}
	return nil
}
