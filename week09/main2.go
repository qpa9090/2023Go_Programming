package main

import "fmt"

func double(number int) {
	number = number * 2
}

func main() {
	// 107p
	var amount int = 5
	double(amount)
	fmt.Printf("%d\n", amount)

}
