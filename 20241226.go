package main

import "fmt"

func product_array(arr []int) []int {
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * arr[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * arr[i+1]
	}
	for i := 0; i < n; i++ {
		arr[i] = left[i] * right[i]
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(product_array(arr))
}
