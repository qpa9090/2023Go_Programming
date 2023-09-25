package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1 //1~100

	fmt.Println("Guess Game (1~100): ")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("You chance: ", 10-i)
		fmt.Print("Guess number: ")
		inputString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputString = strings.TrimSpace(inputString)
		inputNumber, err := strconv.Atoi(inputString)
		if err != nil {
			log.Fatal(err)
		}
		if inputNumber == answer {
			fmt.Println("Great You got the number! Congratulations") //right answer
		}
		if inputNumber < answer {
			fmt.Println("Your guess number is lower than answer.") //answer is higher!
		} else if inputNumber > answer {
			fmt.Println("Your guess number is higher than answer.") // answer is low
		}

	}
}
