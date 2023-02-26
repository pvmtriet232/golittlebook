package main

import "fmt"

func main() {
	goku := &Saiyan{"Goku", 9001}
	goku.Super()
	goku.ChangeName()
	fmt.Println(goku.Power) // will print 19001
	fmt.Println(goku.Name)  // will print 19001
}

// func Super(s *Saiyan) {
// 	s.Power += 10000
// }

func (s *Saiyan) Super() {
	s.Power += 10000
}
func (s *Saiyan) ChangeName() {
	s.Name = "Kakarotto!"
}

type Saiyan struct {
	Name  string
	Power int
}
