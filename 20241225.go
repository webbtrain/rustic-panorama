package main

import (
	"fmt"
)

func has_pair_with_sum(data []int, sum int) bool {
	// O(n) time complexity
	// O(1) space complexity
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == sum {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(has_pair_with_sum([]int{1, 2, 3, 9}, 8)) // false
	fmt.Println(has_pair_with_sum([]int{1, 2, 4, 4}, 8)) // true
}
