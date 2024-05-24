package main

import "fmt"

func main() {
	var input int8
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if input > 2 && input%2 == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
	return
}
