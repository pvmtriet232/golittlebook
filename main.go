package main

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/pvmtriet232/golittlebook/models"
)

func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)
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
