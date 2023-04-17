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

	ShowSaiyan(2)

}

func ShowSaiyan(id int) {
	item := db.LoadSaiyan(id)
	fmt.Printf("Saiyan name: %v, Power :%v \n", item.Name, item.Power)
}
