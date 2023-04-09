package main

import (
	"fmt"

	"github.com/pvmtriet232/golittlebook/models"
)

func main() {
	Vegeta := &models.Saiyan{Name: "Vegeta", Power: 9000, Father: &models.Saiyan{Name: "King Vegeta", Power: 15000}}
	Goku := &models.Saiyan{Name: "Goku", Power: 10000, Father: &models.Saiyan{Name: "Bardock", Power: 8000}}
	Broly := &models.Saiyan{Name: "Broly", Power: 50000, Father: &models.Saiyan{Name: "Paragus", Power: 3000}}
	Gohan := &models.Saiyan{Name: "Gohan", Power: 8000, Father: nil}
	warriors := []*models.Saiyan{Vegeta, Goku, Gohan, Broly}

	Cell := NewSaiyan("Cell", 9000, nil)
	warriors = append(warriors, &Cell)
	fmt.Println(warriors[1].Name)
	for index, _ := range warriors {

		fmt.Printf("name: %v\n", warriors[index].Name)
	}

}

func NewSaiyan(name string, power int, father *models.Saiyan) models.Saiyan {
	return models.Saiyan{
		Name:   name,
		Power:  power,
		Father: nil,
	}
}

func extractPowers(s []*models.Saiyan) []int {
	powers := make([]int, len(s))
	for index, value := range s {
		powers[index] = value.Power
	}
	return powers
}
