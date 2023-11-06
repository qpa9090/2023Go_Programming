package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[2] = 5

	primes := [3]int{2, 3, 5} // initialize by array literal

	for i := 0; i < 3; i++ {
		fmt.Println(primes[i])
	}

	// 초기화 하지 않은 원소의 제로 값은 해당 원소 타입의 제로값으로 결정된다.
	test := [5]bool{true, true, true}
	fmt.Println(test[3])
}
