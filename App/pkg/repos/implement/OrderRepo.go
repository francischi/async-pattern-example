package implement

import (
	// "fmt"
	"golang/pkg/base"
	"golang/pkg/helpers"
)

type OrderRepo struct{
	base.Repository
	ConnPool *helpers.ConnPool
}

func NewOrderRepo(ConnPool *helpers.ConnPool) *OrderRepo{
	return &OrderRepo{ConnPool:ConnPool}
}

func(o *OrderRepo) AddOrder(orderId string , orderProducts map[string]int)(err error){

	// do something

	return nil
}

