package orderModule

import (
	// "fmt"
	"golang/pkg/base"
	"golang/pkg/helpers"
	"golang/pkg/orderModule/dtos"
	"golang/pkg/productModule"
	"golang/pkg/common"
	"golang/pkg/repos/interfaces"
)

type OrderService struct{
	base.Service
	ProductService  *productModule.ProductService 
	MailService  *common.MailService 
	OrderRepo interfaces.OrderRepo
}

func NewOrderService(productService *productModule.ProductService , OrderRepo interfaces.OrderRepo , MailService  *common.MailService ) *OrderService{
	var OrderService OrderService
	OrderService.ProductService = productService
	OrderService.OrderRepo = OrderRepo
	OrderService.MailService = MailService
	return &OrderService
}

func(s *OrderService)Add(dto *dtos.AddDto)(err error){
	if err := dto.Check();err!=nil{
		return s.InvalidArgument(err.Error())
	}

	orderProducts := make(map[string]int)
	for _,product := range dto.Products{
		orderProducts[product.ProductId] = product.Quantity
	}

	if err:=s.ProductService.ReduceProductsQuantity(orderProducts);err!=nil{
		return err
	}

	orderId := helpers.CreateUuid()

	if err:=s.OrderRepo.AddOrder( orderId , orderProducts);err!=nil{
		return err
	}

	if err:= s.MailService.SendOrderConfirmMail("Thanks for making an order!" , dto.Buyer);err!=nil{
		return err
	}

	return nil
}