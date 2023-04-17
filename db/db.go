package db

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

var SaiyanList = make([]*Saiyan, 10)

func LoadSaiyan(id int) *Saiyan {
	return *&SaiyanList[id]
}
func (s *Saiyan) TestSaiyan() {
	fmt.Printf("Name %v, Power:%v \n", s.Name, s.Power)
}
