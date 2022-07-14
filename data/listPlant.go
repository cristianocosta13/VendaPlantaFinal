package data

import (
	"awesomeProject/templates"
	"errors"
)

var plants []templates.Plant

//func SalvarPlanta(planta modelos.Planta) {
//	plantas = append(plantas, planta)
//}

func SavePlant(plant templates.Plant) []templates.Plant {
	plants = append(plants, plant)
	return plants
}

func SearchByCod(cod int) (*templates.Plant, error) {
	for _, p := range plants {
		if p.Cod == cod {
			return &p, nil
		}
	}
	return nil, errors.New("Planta não encontrada!")
}
