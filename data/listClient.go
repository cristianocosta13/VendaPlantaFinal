package data

import (
	"awesomeProject/templates"
	"errors"
)

var clients []templates.Client

//func SalvarCliente(cliente modelos.Cliente) {
//	clientes = append(clientes, cliente)
//}

func SaveClient(client templates.Client) []templates.Client {
	clients = append(clients, client)
	return clients
}

func SearchClienteByCod(cpf string) (*templates.Client, error) {
	for _, c := range clients {
		if c.Cpf == cpf {
			return &c, nil
		}
	}
	return nil, errors.New("Cliente não encontrado!")
}
