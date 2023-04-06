package main

import (
	"fmt"

	"github.com/pvmtriet232/golittlebook/models"
)

func main() {
	var triet = &models.Saiyan{
		Name:  "Triet",
		Power: 9009,
	}
	triet.Hello()
	fmt.Println(triet.Power)
	// fmt.Printf("triet power currently is: %v\n", triet.Power)
	// fmt.Println(&triet.Power, *&triet.Power)
	khoa := models.NewSaiyan("khoa", 9007)
	fmt.Println(khoa.Power)
}

// declare a struct
