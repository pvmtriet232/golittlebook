package main

import "fmt"

func main() {
	var triet = &Saiyan{
		Name:  "Triet",
		Power: 9000,
	}
	triet.Hello()
	fmt.Println(triet.Power)
	// fmt.Printf("triet power currently is: %v\n", triet.Power)
	// fmt.Println(&triet.Power, *&triet.Power)
	khoa := NewSaiyan("Khoa", 9003)
	fmt.Println(khoa)
}

type Saiyan struct {
	Name  string
	Power int
}

func (s *Saiyan) Hello() {
	s.Power = s.Power + 10000
	// fmt.Printf("power inside a function %v\n", s.Power)
	// fmt.Println(&s.Power, *&s.Power)
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
