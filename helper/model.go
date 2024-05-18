package helper

import (
	"halo-suster/model/domain"
	"halo-suster/model/web"
)

func ToCategoryResponseUser(user domain.User, token string) web.UserRes {
	return web.UserRes{
		Nip:   user.Nip,
		Name:  user.Name,
		Token: token,
	}
}

// func ToCategoryResponseCreateProduct(product domain.Products) web.ProductCreateRes {
// 	return web.ProductCreateRes{
// 		Id:        product.Id,
// 		CreatedAt: product.CreatedAt,
// 	}
// }
