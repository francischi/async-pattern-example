package implement
import (
	// "fmt"
	"encoding/json"
	"golang/pkg/base"
	"golang/pkg/helpers"
	"github.com/streadway/amqp"
)

type MailRepo struct{
	base.Repository
	ConnPool *helpers.ConnPool
}

type MailContent struct{
	Context string
	To 		string
}

func NewMailRepo(ConnPool *helpers.ConnPool) *MailRepo{
	return &MailRepo{ConnPool:ConnPool}
}

func(m *MailRepo) PushMailIntoQueue(context string , to string)(err error){
	conn ,err := m.ConnPool.GetConn()

	ch , queue ,err := m.prepareQueueChannel(conn , helpers.GetEnvStr("amqp.mailchannel"))
	if err!=nil{
		return m.SystemError("mail_repo_error :"+err.Error())
	}

	mail := MailContent{
		Context: context,
		To: to,
	}

	encodedQueueMail, err := json.Marshal(mail)

	err = ch.Publish(
		"",     // exchange
		queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "application/json",
		  Body:        encodedQueueMail,
	})
	if err!=nil{
		return m.SystemError("mail_repo_error :"+err.Error())
	}

	ch.Close() 
	if err = m.ConnPool.ReturnConn(conn);err!=nil{
		return m.SystemError("mail_repo_error :"+err.Error())
	}

	return nil
}

func (m *MailRepo)prepareQueueChannel(conn *amqp.Connection , queueName string)( *amqp.Channel ,amqp.Queue ,  error) {
	var channel *amqp.Channel
	var queue amqp.Queue
	
	channel, err := conn.Channel()
	if err!=nil{
		return channel, queue ,err
	}

	queue , err = channel.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err!=nil{
		return channel , queue , err
	}
	return channel , queue , nil
}