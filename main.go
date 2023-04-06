package main

import (
	"fmt"
)

func main() {
	var triet = &models.Saiyan{
		Name:  "Triet",
		Power: 9000,
	}
	triet.Hello()
	fmt.Println(triet.Power)
	// fmt.Printf("triet power currently is: %v\n", triet.Power)
	// fmt.Println(&triet.Power, *&triet.Power)
	khoa := NewSaiyan("khoa", 9007)
	fmt.Println(khoa.Power)
}

// declare a struct

func (s *models.Saiyan) Hello() {
	s.Power = s.Power + 10000
	// fmt.Printf("power inside a function %v\n", s.Power)
	// fmt.Println(&s.Power, *&s.Power)
}

func NewSaiyan(name string, power int) models.Saiyan {
	return models.Saiyan{
		Name:  name,
		Power: power,
	}
}
