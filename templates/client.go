package templates

import (
	"errors"
	"fmt"
)

type Client struct {
	name, Phone, Address, Cpf string
}

func (c Client) String() string {
	return fmt.Sprint("Nome: ", c.name, "\nCPF: ", c.Cpf, "\nTelefone: ", c.Phone, "\nEndereço: ", c.Address)
}

func (c Client) Name() string {
	return c.name
}

func (c Client) SetName(name string) error {
	if name == "" {
		return errors.New("nome inválido")
	}

	c.name = name
	return nil
}
