package db

type Saiyan struct {
	Name  string
	Power int
}

var SaiyanList = make([]*Saiyan, 10)

func LoadSaiyan(id int) *Saiyan {
	return *&SaiyanList[id]
}
