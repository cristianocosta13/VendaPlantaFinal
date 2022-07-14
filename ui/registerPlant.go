package ui

import (
	"awesomeProject/data"
	"awesomeProject/templates"
	"errors"
	"fmt"
)

func RegisterPlant() (*templates.Plant, error) {
	var species, color, sized, description string
	var inventory, cod int
	var price float32

	fmt.Print("\nSobre a planta, insira:\n")
	fmt.Print("Especie: ")
	fmt.Scan(&species)
	fmt.Print("Cor: ")
	fmt.Scan(&color)
	fmt.Print("Porte: ")
	fmt.Scan(&sized)
	fmt.Print("Preço: ")
	fmt.Scan(&price)
	fmt.Print("Código: ")
	fmt.Scan(&cod)
	fmt.Print("Estoque: ")
	fmt.Scan(&inventory)
	fmt.Print("Descrição: ")
	fmt.Scan(&description)

	plant := templates.Plant{
		Species:     species,
		Color:       color,
		Sized:       sized,
		Price:       price,
		Inventory:   inventory,
		Description: description,
		Cod:         cod,
	}
	if species != "" && color != "" && sized != "" && price != 0 && description != "" && cod != 0 {
		data.SavePlant(plant)
		return &plant, nil
	}

	return nil, errors.New("Dados incompativeis ou não totalmente preenchidos")
}
