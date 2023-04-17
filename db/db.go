package db

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

var SaiyanList = make([]*Saiyan, 10)

func LoadSaiyan(id int) (string, int) {
	name := *&SaiyanList[id].Name
	power := *&SaiyanList[id].Power
	fmt.Printf("name: %v, power: %v", name, power)
}
