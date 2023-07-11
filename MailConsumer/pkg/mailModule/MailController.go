package mailModule

import (
	"log"
	"encoding/json"
)

type MailController struct{
	MailService MailService
}

func NewMailController(mailService *MailService) *MailController{
	return &MailController{
		MailService : *mailService,
	}
}

func failOnError(err error) {
	if err != nil {
	  log.Fatalf("%s",err)
	}
}

func (c *MailController)SendMail(body []byte){

	var dto SendDto
	
	err := json.Unmarshal(body, &dto)
	failOnError(err)

	err = c.MailService.Send(&dto)
	failOnError(err)
}