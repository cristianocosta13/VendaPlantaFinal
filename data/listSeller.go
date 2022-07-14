package data

import (
	"awesomeProject/templates"
	"errors"
)

var sellers []templates.Seller

//func SalvarVendedor(vendedor modelos.Vendedor) {
//	vendedores = append(vendedores, vendedor)
//}

func SaveSeller(seller templates.Seller) []templates.Seller {
	sellers = append(sellers, seller)
	return sellers
}

func SearchSellerByCpf(cpf string) (*templates.Seller, error) {
	for _, v := range sellers {
		if v.Cpf == cpf {
			return &v, nil
		}
	}
	return nil, errors.New("Vendedor não encontrado!")
}
