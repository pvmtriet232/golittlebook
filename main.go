package main

import (
	"fmt"

	"github.com/pvmtriet232/golittlebook/models"
)

func main() {
	scores := []int{1,2,3,4,5}
	slice := scores[0:4]
	slice[0] = 999
	fmt.Println(scores)
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
