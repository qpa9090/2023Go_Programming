package main

import (
	"fmt"
	"week10/src/greeting"
	"week10/src/mymath"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Println(mymath.MyAbs(-10))
	fmt.Println(mymath.MyPower(2, 10))
}
