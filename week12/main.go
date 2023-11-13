package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5)

	s := [5]int{0, 0, 0, 0, 0} // slice literal

	s[4] = 99
	s[2] = 11

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// }

	for _, value := range s {
		fmt.Println(value)
	}
	fmt.Println()

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

}
