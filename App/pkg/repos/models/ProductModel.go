package models

type ProductModel struct{
	Id         	int    `json:"id"`
	ProductId   string `gorm:"type:varchar(36);not null;index:product_id"`
	Name       	string `gorm:"type:varchar(100);not null"`
	Price		int    `gorm:"type:int(8) unsigned;not null"`
	Frozen    	int	   `gorm:"type:int(10) unsigned;not null"`
	Available	int    `gorm:"type:int(10) unsigned;not null"`
	CreateTime  int    `gorm:"type:int(10);not null"`
	UpdateTime 	int    `gorm:"type:int(10);default:null"`
	DeleteTime  int    `gorm:"type:int(10);default:null"`
}

func (m *ProductModel) TableName() string {
	return "product"
}