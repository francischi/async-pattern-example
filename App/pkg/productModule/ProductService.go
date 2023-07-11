package productModule

import(
	"golang/pkg/base"
	"golang/pkg/repos/interfaces"
)

type ProductService struct{
	base.Service
	ProductRepo interfaces.ProductRepo
}

func NewProductService(productRepo interfaces.ProductRepo) *ProductService{
	var ProductService ProductService
	ProductService.ProductRepo = productRepo
	return &ProductService
}

func (s *ProductService) ReduceProductsQuantity(productIds map[string]int)(err error){
	// s.ProductRepo.GetProducts(productIds)
	
	return nil
}