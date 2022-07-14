package templates

import "fmt"

type Plant struct {
	Species, Sized, Color, Description string
	Price                              float32
	Inventory, Cod                     int
}

func (p Plant) Plant() {
	fmt.Print("Espécie: ", p.Species, "\nPorte: ", p.Sized, "\nCor: ", p.Color, "\nPreço: ",
		p.Price, "\nCódigo: ", p.Cod, "\nDescrição: ", p.Description, "\nEstoque: ", p.Inventory)
}
