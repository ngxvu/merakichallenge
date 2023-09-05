package main

import (
	"fmt"
)

func Solution(S string) string {
	var capsLock bool
	var result []byte

	for i := 0; i < len(S); i++ {
		switch S[i] {
		case 'C':
			capsLock = !capsLock
		case 'B':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		default:
			if capsLock {
				result = append(result, S[i]-'a'+'A')
			} else {
				result = append(result, S[i])
			}
		}
	}
	return string(result)
}

func main() {
	fmt.Println(Solution("CrCellBax"))          // Output: Relax
	fmt.Println(Solution("CgCoodlBClCuck"))     // Output: GoodLuck
	fmt.Println(Solution("aCaBBCCab"))          // Output: AB
	fmt.Println(Solution("aBB"))                // Output: ""
	fmt.Println(Solution("CnCguyenCxCuanCvCu")) // Output: "NguyenXuanVu"
}
