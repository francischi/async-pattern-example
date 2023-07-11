package implement

import (
	// "fmt"
	"gorm.io/gorm"
	"golang/pkg/repos/models"
)

type ProductRepo struct {
	DBconn *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo{
	return &ProductRepo{DBconn:db}
}

func (p *ProductRepo) GetProducts(productIds []*string)(productmodels []*models.ProductModel , err error){
	var products []*models.ProductModel
	err = p.DBconn.Where("productId IN ?", productIds).Find(&products).Error
	return products,err
}