package interfaces

type MailRepo interface{

	PushMailIntoQueue(context string , to string)(err error)

}