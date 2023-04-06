package main

import "fmt"

func main() {
	var triet = Saiyan{
		Name:  "Triet",
		Power: 9000,
	}
	var ptr = &triet
	ptr.Power = 5000

	fmt.Println(triet)
	fmt.Println(ptr)
}

type Saiyan struct {
	Name  string
	Power int
}

func (s *Saiyan) Hello() {
	s.Power = s.Power + 10000
}
