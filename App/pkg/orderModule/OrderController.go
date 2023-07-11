package orderModule

import (
	// "fmt"
   "golang/pkg/orderModule/dtos"
   "golang/pkg/tokenModule"  
   "github.com/gin-gonic/gin"
   "golang/pkg/base"
)

type OrderController struct{
	OrderService OrderService
	TokenService tokenModule.TokenService
	base.Controller
}

func NewOrderController(orderService *OrderService ,tokenService *tokenModule.TokenService) *OrderController{
	return &OrderController{
		OrderService : *orderService,
		TokenService : *tokenService,
	}
}

func (c *OrderController)Add(g *gin.Context){
	jwtToken := g.Request.Header["Bearer-Token"][0]

	payload,err := c.TokenService.ParsePayload(jwtToken)
	if(err!=nil){
		c.HandleError(g,err)
	}

	var dto dtos.AddDto
	g.Bind(&dto)
	dto.MemberId = payload.MemberId

	err = c.OrderService.Add(&dto)
	if err!=nil{
		c.HandleError(g , err)
	}else {
		c.SuccessRes(g,nil)
	}
}