package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution([]int{1, 8, 3, 2}))    // Output: 6
	fmt.Println(Solution([]int{2, 3, 3, 2}))    // Output: 3
	fmt.Println(Solution([]int{6, 2, 4, 7}))    // Output: 0
	fmt.Println(Solution([]int{0, 0, 0, 0, 0})) // Output: 1
	fmt.Println(Solution([]int{0, 0, 0, 0, 1})) // Output: 5
}

func Solution(digits []int) int {
	count := 0
	seen := make(map[int]bool) // map để lưu những giá trị đã xuất hiện
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			for k := 0; k < len(digits); k++ {
				for l := 0; l < len(digits); l++ {
					if i == j || i == k || i == l || j == k || j == l || k == l {
						continue
					}
					hour := digits[i]*10 + digits[j]
					minute := digits[k]*10 + digits[l]
					if hour < 24 && minute < 60 {
						value := hour*100 + minute // giá trị biểu diễn cho giờ và phút
						if !seen[value] {          // nếu chưa xuất hiện trong map
							count++
							seen[value] = true // đánh dấu là đã xuất hiện
						}
					}
				}
			}
		}
	}
	return count
}
