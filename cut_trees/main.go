package main

import "fmt"

func Solution(A []int) int {
	count := 0
	canNotCut := -1
	firstIndex := A[0]
	secondIndex := A[1]
	if len(A) <= 3 {
		return count
	} else if len(A) == 4 && firstIndex < secondIndex {
		for j := 1; j < len(A)-1; j++ {
			if (A[j-1] < A[j] && A[j] > A[j+1]) || (A[j-1] > A[j] && A[j] > A[j+1]) {
				return count
			}
		}
		return canNotCut
	} else if len(A) == 4 && firstIndex > secondIndex {
		for j := 1; j < len(A)-1; j++ {
			if (A[j-1] < A[j] && A[j] > A[j+1]) || (A[j-1] < A[j] && A[j] < A[j+1]) {
				return count
			}
		}
		return canNotCut
	} else if len(A) > 4 {
		for i := 0; i < len(A)-1; i++ {
			copyA := make([]int, len(A))
			copy(copyA, A)
			for j := i; j < len(copyA)-1; j++ {
				copyA[j] = copyA[j+1]
			}
			newA := copyA[:len(copyA)-1]
			if len(newA) < 2 {
				return count
			}
			first := newA[0]
			second := newA[1]
			if first < second {
				for j := 1; j < len(newA)-1; j++ {
					if (newA[j-1] < newA[j] && newA[j] > newA[j+1]) || (newA[j-1] > newA[j] && newA[j] > newA[j+1]) {
						count++
					}
				}
			} else if first > second {
				for j := 1; j < len(newA)-1; j++ {
					if (newA[j-1] < newA[j] && newA[j] > newA[j+1]) || (newA[j-1] < newA[j] && newA[j] < newA[j+1]) {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println(Solution([]int{3, 4, 5, 3, 7})) // Should print 3
	//fmt.Println(Solution([]int{1, 2, 3, 4}))    // Should print -1
	//fmt.Println(Solution([]int{1, 3, 1, 2}))    // Should print 0
}
