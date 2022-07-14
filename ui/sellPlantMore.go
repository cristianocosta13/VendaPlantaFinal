package ui

import (
	"awesomeProject/data"
	"awesomeProject/templates"
	"fmt"
)

func SellPlantMore() {
	var cod, quant int
	var cpfSeller, cpfClient string
	var plants []templates.Plant
	var e error
	var somaPreco, desconto float32

	fmt.Print("Insira o CPF do vendedor responsável pela venda: ")
	fmt.Scan(&cpfSeller)
	fmt.Print("Insira o CPF do cliente: ")
	fmt.Scan(&cpfClient)
	fmt.Println("Quantas plantas deseja comprar?")
	fmt.Scan(&quant)
	if quant > 1 {
		for i := 0; i < quant; i++ {
			fmt.Print("Insira o código da planta a ser comprada: ")
			fmt.Scan(&cod)
			plant, errorPlant := data.SearchByCod(cod)
			if errorPlant == nil {
				plants = append(plants, *plant)
			} else {
				fmt.Print(errorPlant, " Não será possivél adicioná-la a lista de compras\n")
			}
		}
	} else {
		fmt.Print("Insira o código da planta a ser comprada: ")
		fmt.Scan(&cod)
		plant, errorPlant := data.SearchByCod(cod)
		if plant == nil {
			fmt.Println("ERRO PLANTA: ", errorPlant)
			plants = append(plants, *plant)
		}
		e = errorPlant
	}

	_, errorClient := data.SearchClienteByCod(cpfClient)
	_, errorSeller := data.SearchSellerByCpf(cpfSeller)

	if e == nil && errorClient == nil && errorSeller == nil {
		for _, p := range plants {
			somaPreco += p.Price
		}
		if len(plants) == 1 {
			desconto = 10
		} else {
			desconto = 30
		}
		amountDebilled := AplicarMais(somaPreco, desconto)
		payback, errorBanckroll, pendingValue := PagarPorMais(amountDebilled, somaPreco)

		if errorBanckroll != nil {
			fmt.Print(errorBanckroll, "\nO cliente deve pagar mais R$", pendingValue)
		} else {
			fmt.Print("O troco do cliente é R$", payback)
		}
	} else {
		fmt.Println("ERRO CLIENTE: ", errorClient)
		fmt.Println("ERRO VENDEDOR: ", errorSeller)
	}

}
