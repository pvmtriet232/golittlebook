package main

import (
	"fmt"

	"github.com/pvmtriet232/golittlebook/models"
)

func main() {
	gohan := &models.Saiyan{
		Name:  "gohan",
		Power: 9000,
		Father: &models.Saiyan{
			Name:   "goku",
			Power:  10000,
			Father: nil,
		},
	}
	fmt.Println(*gohan.Father)

	Khoa := models.NewSaiyan("khoa", 7000, &models.Saiyan{
		Name:   "Lam",
		Power:  1000,
		Father: nil,
	})
	fmt.Println(Khoa.Power)
}

// declare a struct
