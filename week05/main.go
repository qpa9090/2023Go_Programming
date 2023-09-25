package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "cs r?cks~"
	replacdWords := strings.NewReplacer("?", "o")
	fixedWords := replacdWords.Replace(brokenWords)
	fmt.Println(fixedWords)
}
