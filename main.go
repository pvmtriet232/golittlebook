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
	fmt.Printf("power:%v\n", extractPowers(warriors))
	fmt.Printf("power:%v\n", extractPowers(warriors))

}

func NewSaiyan(name string, power int, father *models.Saiyan) models.Saiyan {
	return models.Saiyan{
		Name:   name,
		Power:  power,
		Father: nil,
	}
}

// declare a slice, loop through the slice and extract index,power store
// in a slice name , then return power
func extractPowers(saiyans []*models.Saiyan) []int {
	powers := make([]int, len(saiyans))
	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}

// declare a slice , loop through the slice of saiyans , append the value
// of individual saiyan.power to the power slice, return the power slice
func extractPowersNew(saiyans []*models.Saiyan) []int {
	powers := make([]int, 0, 10)
	for _, saiyan := range saiyans {
		powers = append(powers, saiyan.Power)
	}
	return powers
}
