// Given inner = 1, outer = 3, points_x = [0, 1, 2, -2, 3] and points_y = [0, 1, 4, 1, 0]
package main

import (
	"fmt"
	"math"
)

func main() {
	inner := 1
	outer := 3
	points_x := []int{0, 1, 2, -2, 3}
	points_y := []int{0, 1, 4, 1, 0}
	rs := solution(inner, outer, points_x, points_y)
	fmt.Print(rs)
}

func solution(inner int, outer int, points_x []int, point_y []int) int {
	count := 0
	for i := 0; i < len(points_x); i++ {
		d := math.Sqrt(float64(points_x[i]*points_x[i] + point_y[i]*point_y[i]))
		if float64(inner) < d && d < float64(outer) {
			count += 1
		}
	}
	return count
}
