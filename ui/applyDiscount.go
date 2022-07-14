package ui

import (
	"awesomeProject/templates"
	"fmt"
)

func ApplyDiscount(plant templates.Plant) float32 {
	var discount float32
	fmt.Print("Insira o percentual de desconto: ")
	fmt.Scan(&discount)
	return plant.Price * discount / 100
}
