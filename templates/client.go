package templates

import "fmt"

type Client struct {
	Name, Phone, Address, Cpf string
}

func (c Client) Client() {
	fmt.Print("Nome: ", c.Name, "\nCPF: ", c.Cpf, "\nTelefone: ", c.Phone, "\nEndereço: ", c.Address)
}
