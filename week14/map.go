package main

import "fmt"

func main() {
	// var games map[int]string 일반적
	// games = make(map[int]string)

	// games := make(map[int]string) 줄여서

	//리터럴
	games := map[int]string{456: "성기훈", 218: "박해수", 067: "강새벽", 001: "오일남", 199: "알리"}

	// games[456] = "성기훈"
	// games[218] = "박해수"
	// games[067] = "강새벽"
	// games[199] = "오일남"
	// games[101] = "알리"
	for _, v := range games {
		fmt.Println(v)
	}
	games[023] = "장덕수"
	delete(games, 101)
	for k, v := range games {
		fmt.Println(k, v)
	}
}
