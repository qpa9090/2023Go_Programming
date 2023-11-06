package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가'
	// var a int16 = 7
	// var a = 7
	// a := 7
	a := 7 // 타입이 한번 결정되면 정수가 들어와도 실수형으로 됨
	var b float64 = 5.3
	a = int(b) // type Conversion, Casting
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c)        // 유니코드(UTF-8) 값 출력
	fmt.Printf("%c\n", c) // 한 글자 출력
	fmt.Printf("%T\n", c) // rune은 결국 int32의 별명(alias)

	fmt.Println(math.Round(2.71))
	fmt.Println("Hello Go!")
	fmt.Println(strings.Title("go git github java"))
}
