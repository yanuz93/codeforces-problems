package main

import (
	"bytes"
	"fmt"
)

func main() {
	var numOfInput int64
	_, err := fmt.Scanln(&numOfInput)
	if err != nil {
		fmt.Println("invalid input", err)
		return
	}

	words := make([]string, 0, numOfInput)
	for i := int64(0); i < numOfInput; i++ {
		var word string
		_, err := fmt.Scanln(&word)
		if err != nil {
			fmt.Println("invalid input", err)
			return
		}
		words = append(words, word)
	}

	var result bytes.Buffer
	for _, word := range words {
		length := len(word)
		if length > 10 {
			fmt.Fprintf(&result, "%s%d%s\n", string(word[0]), length-2, string(word[length-1]))
			continue
		}
		fmt.Fprintf(&result, "%s\n", word)
	}

	fmt.Print(result.String())
}
