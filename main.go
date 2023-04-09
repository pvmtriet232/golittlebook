package main

import (
	"fmt"
	"strings"

	"github.com/pvmtriet232/golittlebook/models"
)

func main() {
	haystack := "th spice must flow on the table of the supermarket"
	findSpace := strings.Index(haystack[4:], " ")
	fmt.Println(findSpace)
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
