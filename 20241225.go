package main

import (
	"fmt"
)

func has_pair_with_sum(data []int, sum int) bool {
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
	nums := []int{1, 2, 3, 9}
	sum := 8
	nums2 := []int{1, 2, 4, 4}
	sum2 := 8
	fmt.Println(has_pair_with_sum(nums, sum))   // false
	fmt.Println(has_pair_with_sum(nums2, sum2)) // true
}
