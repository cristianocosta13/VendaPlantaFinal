package ui

import (
	"errors"
	"fmt"
)

func PagarPorMais(discount float32, preco float32) (float32, error, float32) {
	var valuePaid float32
	fmt.Print("Insira o valor a ser pago: ")
	fmt.Scan(&valuePaid)
	if valuePaid >= preco {
		return valuePaid - (preco - discount), nil, 0
	} else {
		return valuePaid, errors.New("Saldo insuficiente!"), preco - valuePaid
	}
}
