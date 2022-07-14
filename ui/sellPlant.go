package ui

import (
	"awesomeProject/data"
	"fmt"
)

func SellPlant() {
	var cod int
	var cpfSeller, cpfClient string

	fmt.Print("Insira o CPF do vendedor responsável pela venda: ")
	fmt.Scan(&cpfSeller)
	fmt.Print("Insira o CPF do cliente: ")
	fmt.Scan(&cpfClient)
	fmt.Print("Insira o código da planta a ser comprada: ")
	fmt.Scan(&cod)

	plant, errorPlant := data.SearchByCod(cod)
	_, errorClient := data.SearchClienteByCod(cpfClient)
	_, errorSeller := data.SearchSellerByCpf(cpfSeller)

	if errorPlant == nil && errorClient == nil && errorSeller == nil {
		plant.Inventory-- //não está decrementando, problema no ponteiro
		amountDebilled := ApplyDiscount(*plant)
		payback, errorBanckroll, pendingValue := Payment(amountDebilled, *plant)
		if errorBanckroll != nil {
			fmt.Print(errorBanckroll, "\nO cliente deve pagar mais R$", pendingValue)
		} else {
			fmt.Print("O troco do cliente é R$", payback)
		}
	} else {
		fmt.Println("ERRO PLANTA: ", errorPlant)
		fmt.Println("ERRO CLIENTE: ", errorClient)
		fmt.Println("ERRO VENDEDOR: ", errorSeller)
	}
}
