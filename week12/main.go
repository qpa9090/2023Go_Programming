package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5)

	s := [5]int{0, 0, 0, 0, 0}

	for _, value := range s {
		fmt.Println(value)
	}
}
