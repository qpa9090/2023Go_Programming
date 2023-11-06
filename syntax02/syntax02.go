package main

import (
	"fmt"
	//	"reflect"
	"strings"
)

func main() {

	texts := "G@ M@ney"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

	// 	/* 변수명은 영문자로 시작
	// 	영문 대문자의 경우 다른 패키지에서 접근할 수 있다.
	// 	소문자로 시작하는 변수는 동일 패키지에서만 접근 가능*/
	// 	var a string
	// 	var b bool
	// 	var c rune
	// 	var d float64
	// 	var E int
	// 	// a := 7

	// 	//naming convention
	// 	var studentId string //student_id대신
	// 	var i int32          //index의 약자

	// 	fmt.Println(a, b, c, d, E, studentId, i)

	// 	fmt.Printf("%T\n", c)

	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(a))
}
