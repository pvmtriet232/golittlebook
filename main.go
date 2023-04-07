package main

import (
	"fmt"
)

func main() {

	scores := []int{2, 5, 2, 6, 67, 7, 1245, 234}
	for index, num := range scores {
		fmt.Printf("index : %v, number : %v\n", index, num)
	}

}

// declare a struct
