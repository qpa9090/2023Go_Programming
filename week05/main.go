package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("input score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputScore = strings.TrimSpace(inputScore)
	Score, err := strconv.ParseFloat(inputScore, 64)
	var grade string

	if Score >= 90.0 {
		grade = "A grade"
	} else {
		grade = "under A grade"
	}

	fmt.Println(inputScore)
	fmt.Println("Youu will get", grade) //undefined: grade
}
