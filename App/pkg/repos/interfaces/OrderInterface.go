package interfaces

type OrderRepo interface{

	AddOrder(orderId string , orderProducta map[string]int)(err error)

}