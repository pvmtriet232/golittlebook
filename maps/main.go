package main

import "fmt"

func main() {
	saiyanMap := make(map[string]int)
	saiyanMap["goku"] = 9001
	saiyanMap["vegeta"] = 19000
	saiyanMap["gohan"] = 5000
	fmt.Println(saiyanMap)
	for key, value := range saiyanMap {
		fmt.Printf("Saiyan name: %v , power: %v\n", key, value)
	}

	saiyanMapTwo := map[string]int{
		"goku":   9001,
		"vegeta": 19000,
	}
	fmt.Println(saiyanMapTwo)

	type saiyan struct {
		Name string
	}
	goku := &saiyan{"goku"}
	vegeta := &saiyan{"vegeta"}
	a := make([]saiyan, 2)
	b := make([]*saiyan, 2)
	fmt.Println(a, b, goku, vegeta)
	b = append(b, goku)
	fmt.Println(a, b, goku, vegeta)

}
