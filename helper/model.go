package helper

import (
	"halo-suster/model/domain"
	"halo-suster/model/web"
)

func ToCategoryResponseStaff(staff domain.Staff, token string) web.StaffRes {
	return web.StaffRes{
		PhoneNumber: staff.PhoneNumber,
		Name:        staff.Name,
		Token:       token,
	}
}

func ToCategoryResponseCreateProduct(product domain.Products) web.ProductCreateRes {
	return web.ProductCreateRes{
		Id:        product.Id,
		CreatedAt: product.CreatedAt,
	}
}
