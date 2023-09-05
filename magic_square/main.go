package main

import "fmt"

func main() {
	A := [][]int{
		{7, 2, 4},
		{2, 7, 6},
		{9, 5, 1},
		{4, 3, 8},
		{3, 5, 4},
	}
	fmt.Print(solution(A))
}

func solution(A [][]int) int {
	rows := len(A)
	cols := len(A[0])
	var rs int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i+2 < rows && j+2 < cols {
				row1 := A[i][j] + A[i][j+1] + A[i][j+2]
				row2 := A[i+1][j] + A[i+1][j+1] + A[i+1][j+2]
				row3 := A[i+2][j] + A[i+2][j+1] + A[i+2][j+2]
				col1 := A[i][j] + A[i+1][j] + A[i+2][j]
				col2 := A[i][j+1] + A[i+1][j+1] + A[i+2][j+1]
				col3 := A[i][j+2] + A[i+1][j+2] + A[i+2][j+2]
				mainDiag := A[i][j] + A[i+1][j+1] + A[i+2][j+2]
				secDiag := A[i][j+2] + A[i+1][j+1] + A[i+2][j]

				if row1 <= 18 && row2 <= 18 && row3 <= 18 && col1 <= 18 && col2 <= 18 && col3 <= 18 && mainDiag <= 18 && secDiag <= 18 {
					rs = 3
				}
			} else if i+1 < rows && j+1 < cols {
				row1 := A[i][j] + A[i][j+1]
				row2 := A[i+1][j] + A[i+1][j+1]
				col1 := A[i][j] + A[i+1][j]
				col2 := A[i][j+1] + A[i+1][j+1]
				mainDiag := A[i][j] + A[i+1][j+1]
				secDiag := A[i][j+1] + A[i+1][j]

				if row1 == 4 && row2 == 4 && col1 == 4 && col2 == 4 && mainDiag == 4 && secDiag == 4 {
					rs = 2
				}
			}
		}
	}
	return rs
}
