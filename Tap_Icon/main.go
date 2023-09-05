package main

import (
	"fmt"
	"math"
)

func main() {
	A := []int{100, 200, 100}
	B := []int{50, 100, 100}
	X := 100
	Y := 70
	fmt.Print(Solution(A, B, X, Y))

}

//loop qua A và lấy theo

func Solution(A []int, B []int, X int, Y int) int {
	var listIcon [][]int
	var stt int
	tap := []int{X, Y}
	if len(A) != len(B) {
	}
	for i := 0; i < len(A); i++ {
		var icon []int
		icon = append(icon, A[i], B[i])
		listIcon = append(listIcon, icon)
		continue
	}

	for i := 0; i < len(listIcon); i++ {
		match := true
		for j := 0; j < len(listIcon[i]); j++ {
			if tap[j] != listIcon[i][j] && math.Abs(float64(tap[j]-listIcon[i][j])) > 20 {
				match = false
			}
		}
		if match {
			stt = i
			break
		}
		if !match {
			stt = -1
		}
	}
	return stt
}
