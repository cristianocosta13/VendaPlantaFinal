package ui

import (
	"awesomeProject/data"
	"awesomeProject/templates"
	"errors"
	"fmt"
)

func RegisterSeller() (*templates.Seller, error) {
	var name, cpf, phone string

	fmt.Print("\nSobre o vendedor, insira:\n")
	fmt.Print("Nome: ")
	fmt.Scan(&name)
	fmt.Print("CPF: ")
	fmt.Scan(&cpf)
	fmt.Print("Telefone: ")
	fmt.Scan(&phone)

	seller := templates.Seller{
		Name:  name,
		Cpf:   cpf,
		Phone: phone,
	}
	if name != "" && phone != "" && cpf != "" {
		data.SaveSeller(seller)
		return &seller, nil
	}
	return nil, errors.New("Dados incompativeis ou não totalmente preenchidos")
}
