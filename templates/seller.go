package templates

import (
	"fmt"
)

type Seller struct {
	Name, Cpf, Phone string
	Plants           []Plant
}

func (s Seller) Seller() {
	fmt.Print("Nome: ", s.Name, "\nCPF: ", s.Cpf, "\nTelefone: ", s.Phone)
}
