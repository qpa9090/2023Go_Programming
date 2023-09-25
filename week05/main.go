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

	fmt.Println(inputScore)
}
