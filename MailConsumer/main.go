package main

import (
	// "fmt"
	"log"
	"mailConsumer/pkg/helpers"
	"mailConsumer/pkg/mailModule"
)

type QueueOrder struct{
	OrderId string 
	OrderProducts map[string]int 
}

func main(){

	if err := helpers.InitEnvSetting();err !=nil{
		panic(err)
	}

	msgs , err := helpers.InitQueueConnect()
	if err !=nil{
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			go mailModule.InitialMailController().SendMail(d.Body)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}