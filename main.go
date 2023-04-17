package main

import (
	"fmt"

	"github.com/pvmtriet232/golittlebook/db"
)

func main() {

	db.SaiyanList[0] = &db.Saiyan{
		Name:  "goku",
		Power: 9001,
	}
	db.SaiyanList[1] = &db.Saiyan{
		Name:  "Vegeta",
		Power: 19000,
	}
	db.SaiyanList[2] = &db.Saiyan{
		Name:  "Nappa",
		Power: 12002,
	}

	// A := *&SaiyanList[1].Power
	// fmt.Println(A)
	a := db.LoadSaiyan(2)
	fmt.Println(a)

	// Name: vegeta, power: 19000

}
