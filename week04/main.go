package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var a int = 9
	var b float32
	var c, d string
	var f string
	var g bool // zero value
	h := 'Z'   // 단축변수선언 short variable declare
	i := "문자열"
	J := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할 수 있음" // 외부에 노출되는 변수
	koreanZombie := "정찬성"                         // camelCase

	// a = 9
	b = 2.7
	c = "inha"
	d = "Go!"

	fmt.Println(a, b, c, d, f, g, h, i, J, koreanZombie)
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(float64(a) > b)
	fmt.Println(a * int(b))

	// fmt.Println(math.Floor("inha")) 에러 출력
	// fmt.Println(strings.Title(4.88))
	fmt.Println(reflect.TypeOf('B'))
	fmt.Println(reflect.TypeOf(100))
	fmt.Println(reflect.TypeOf(2.71))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf("Go!"))
	fmt.Println('A') // rune (Go uses the Unicode Standard)
	fmt.Println('허', '현', '도', '\n')
	math.Floor(2.75)
	fmt.Println(math.Floor(2.75))
	fmt.Println(math.Ceil(2.75))
	fmt.Println(strings.Title("오픈소스\t프로그래밍\n\"Go\""))
	fmt.Println(strings.Title("open source programming go!")) // 소문자를 대문자로 변경
} // 리터럴은 값을 할당할 수 없다. 소스로만 사용가능. 변수등에 저장되는 변할 수 없는 값
