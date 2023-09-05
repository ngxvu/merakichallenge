package main

import (
	"fmt"
	"strconv"
)

func main() {
	A := Solution(1, 111)
	fmt.Print(A)
}

func Solution(A int, B int) int {
	if A > B {
		return 0
	}
	var numb int
	if A == B {
		repeatDigit := make(map[rune]int)
		S := strconv.Itoa(A)
		for _, digit := range S {
			if _, ok := repeatDigit[digit]; !ok {
				repeatDigit[digit] = 1
			} else {
				repeatDigit[digit]++
			}
		}

		for _, value := range repeatDigit {
			if value >= 3 {
				numb = 0
			}
			if value < 3 {
				numb = 1
			}
		}
	} else if A < B {
		for i := A; i <= B; i++ {
			repeatDigit := make(map[rune]int)
			S := strconv.Itoa(i)
			for _, digit := range S {
				if _, ok := repeatDigit[digit]; !ok {
					repeatDigit[digit] = 1
				} else {
					repeatDigit[digit]++
				}
			}
			for _, value := range repeatDigit {
				if value >= 3 {
					return numb
				}
			}
			numb++
		}
	}
	return numb
}
