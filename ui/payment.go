package ui

import (
	"awesomeProject/templates"
	"errors"
	"fmt"
)

func Payment(discount float32, plant templates.Plant) (float32, error, float32) {
	var valuePaid float32
	fmt.Print("Insira o valor a ser pago: ")
	fmt.Scan(&valuePaid)
	if valuePaid > plant.Price {
		return valuePaid - (plant.Price - discount), nil, 0
	} else {
		return valuePaid, errors.New("Saldo insuficiente!"), plant.Price - valuePaid
	}
}
