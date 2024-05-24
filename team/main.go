package main

import (
	"fmt"
)

func main() {
	var numOfProblems int16
	_, err := fmt.Scanln(&numOfProblems)
	if err != nil {
		fmt.Println("invalid input: ", err)
		return
	}

	votes := make([]int8, 0, numOfProblems)
	for i := int16(0); i < numOfProblems*3; i++ {
		var vote int8
		_, err := fmt.Scan(&vote)
		if err != nil {
			fmt.Println("invalid input: ", err)
			return
		}
		votes = append(votes, vote)
	}

	var acceptedProblems int16
	for i := int16(0); i < numOfProblems; i++ {
		var count int8
		for j := int8(0); j < 3; j++ {
			if votes[i*3+int16(j)] == 1 {
				count++
			}

			// increment acceptedProblems when count > 1
			// and break the inner loop
			// so there will be no problem that accepted more than once
			if count > 1 {
				acceptedProblems++
				break
			}
		}

	}

	fmt.Println(acceptedProblems)
}
