package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("input score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') // solution option 2
	if err != nil {                            // conditionals
		log.Fatal(err)
	}
	if inputScore >= 90 { // mismatched types string and untyped int
		grade := "A grade"
	} else {
		grade := "under A grade"
	}

	fmt.Println(inputScore)
}
