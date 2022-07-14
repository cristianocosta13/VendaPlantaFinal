package ui

import (
	"awesomeProject/data"
	"awesomeProject/templates"
	"errors"
	"fmt"
)

func RegisterClient() (*templates.Client, error) {
	var name, phone, address, cpf string

	fmt.Print("\nSobre o cliente, insira:\n")
	fmt.Print("Nome: ")
	fmt.Scan(&name)
	fmt.Print("Telefone: ")
	fmt.Scan(&phone)
	fmt.Print("CPF: ")
	fmt.Scan(&cpf)
	fmt.Print("Endereço: ")
	fmt.Scan(&address)

	client := templates.Client{
		Name:    name,
		Address: address,
		Phone:   phone,
		Cpf:     cpf,
	}
	if name != "" && address != "" && phone != "" && cpf != "" {
		data.SaveClient(client)
		return &client, nil
	}

	return nil, errors.New("Dados incompativeis ou não totalmente preenchidos")
}
