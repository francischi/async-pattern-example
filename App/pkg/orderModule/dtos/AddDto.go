package dtos

import (
	"errors"
	"golang/pkg/helpers"
)

type AddDto struct{
	MemberId string
	Products []product
	Buyer string
}

type product struct{
	ProductId string
	Quantity int
}

func (dto *AddDto) Check()(error){
	if len(dto.MemberId)==0{
		return errors.New("member_id_required")
	}
	if len(dto.Products)==0{
		return errors.New("products_required")
	}
	if len(dto.Buyer)==0{
		return errors.New("buyer_required")
	}
	if !helpers.IsValidEmail(dto.Buyer){
		return errors.New("invalid_buyer")
	}
	for _, element := range dto.Products {
		if err := element.Check(); err!=nil{
			return err
		}
	}
	return nil
}

func (p *product) Check()(error){
	if len(p.ProductId) == 0{
		return errors.New("productId_required")
	}
	if p.Quantity <= 0{
		return errors.New("invalid_quantity")
	}
	return nil
}