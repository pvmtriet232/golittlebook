package main

import (
	"fmt"

	"github.com/pvmtriet232/golittlebook/db"
)

func main() {
	var SaiyanList = make([]*db.Saiyan, 10)

	SaiyanList[0] = &db.Saiyan{
		Name:  "goku",
		Power: 9001,
	}
	SaiyanList[1] = &db.Saiyan{
		Name:  "Vegeta",
		Power: 19000,
	}
	SaiyanList[2] = &db.Saiyan{
		Name:  "Nappa",
		Power: 12002,
	}

	A := &SaiyanList[1].Power
	fmt.Println(A)

}
