//+build wireinject

package orderModule

import (
	"golang/pkg/helpers"
	"golang/pkg/tokenModule"
	"golang/pkg/productModule"
	"golang/pkg/common"
	"github.com/google/wire"
	impl "golang/pkg/repos/implement"
	"golang/pkg/repos/interfaces"
)

var ProductRepo = wire.NewSet(impl.NewProductRepo,wire.Bind(new(interfaces.ProductRepo), new(*impl.ProductRepo)))

var OrderRepo = wire.NewSet(impl.NewOrderRepo,wire.Bind(new(interfaces.OrderRepo), new(*impl.OrderRepo)))

var MailRepo = wire.NewSet(impl.NewMailRepo,wire.Bind(new(interfaces.MailRepo), new(*impl.MailRepo)))

func InitialOrderController() *OrderController{
	wire.Build(
		NewOrderController,
		NewOrderService,
		OrderRepo,

		tokenModule.NewTokenService , 

		productModule.NewProductService,
		ProductRepo,

		common.NewMailService,
		MailRepo,

		helpers.GetConnPool,
		helpers.NewSqlSession ,
	)
	return &OrderController{}
}