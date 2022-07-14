package ui

func AplicarMais(precoTotal, discount float32) float32 {
	discount = 30
	return precoTotal * discount / 100
}
