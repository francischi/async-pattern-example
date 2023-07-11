package interfaces

import(
	"golang/pkg/repos/models"
)

type ProductRepo interface {

	GetProducts(productIds []*string)(products []*models.ProductModel , err error)

}