package main

import "fmt"

func main() {
	fmt.Print(Solution("beans", "banes"))
}

func Solution(S string, T string) string {

	if S == T {
		return "EQUAL"
	}
	var rs string
	if len(S) != len(T) {
		if len(T) == len(S)+1 {
			for i := range S {
				if S[i] != T[i] {
					return "ADD " + string(T[i])
				}
			}
			return "ADD " + string(T[len(T)-1])
		} else {
			return "IMPOSSIBLE"
		}
	} else {
		diff := 0
		index := -1
		for i := range S {
			if S[i] != T[i] {
				diff++
				if index == -1 {
					index = i
				}
			}
		}
		if diff == 1 {
			return "CHANGE " + string(S[index]) + " " + string(T[index])
		}
		if diff >= 2 {
			for i := index; i < len(T); i++ {
				if T[i] != S[i+1] {
					rs = "MOVE " + string(T[i])
					break
				}
			}
		}
		return rs
	}
}
