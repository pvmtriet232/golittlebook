package main

import "fmt"

func main() {
	goku := &saiyan{
		Name:  "goku",
		Power: 9000,
	}
	Super(goku)
	fmt.Println(goku.Power)
}

func Super(s *saiyan) {
	s.Power += 10000
}

type saiyan struct {
	Name  string
	Power int
}
