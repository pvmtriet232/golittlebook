package models

// declare a struct
type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

// method of a struct
func (s *Saiyan) Hello() {
	s.Power = s.Power + 10000
	// fmt.Printf("power inside a function %v\n", s.Power)
	// fmt.Println(&s.Power, *&s.Power)
}

func NewSaiyan(name string, power int, Father *Saiyan) *Saiyan {
	return &Saiyan{
		Name:   name,
		Power:  power,
		Father: &Saiyan{},
	}
}

func NSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}
