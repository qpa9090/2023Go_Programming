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
	fmt.Println(test[3]) // zero value
	fmt.Println(test)
	// fmt.print(test[5]) // comfile error, invalid argument index 5 out of bounds

	i := 0
	for i < len(test) { // while // if 크기를 넘어가면 panic: runtime error: index out of range [5] with length 5
		fmt.Println(test[i])
		i++
	}

	for idx, prime := range primes {
		fmt.Println(idx, prime)
	}

	for _, prime := range primes { //언더스코어 사용하여 값만 나오게 함
		fmt.Println(prime)
	}
}
