package db


type Saiyan struct {
	Name  string
	Power int
}

SaiyanList := make([]*Saiyan, 10)

SaiyanList[0] = &Saiyan{
	Name:  "goku",
	Power: 9001,
}
SaiyanList[1] = &Saiyan{
	Name:  "Vegeta",
	Power: 19000,
}
SaiyanList[2] = &Saiyan{
	Name:  "Nappa",
	Power: 12002,
}

a := &SaiyanList[1].Power
fmt.Println(a)