package main

import (
	"fmt"
)

func countingSort(arr []int) []int {
	var tmp = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		tmp[arr[i]]++
	}

	for i := 1; i < len(arr); i++ {
		tmp[i] += tmp[i-1]
	}

	var res = make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		res[tmp[arr[i]] - 1] = arr[i]
		tmp[arr[i]] -= 1
	}

	return res
}

func main() {
	arr := []int{6, 0, 2, 0, 1, 3, 4, 6, 1, 3, 2}
	fmt.Printf("before: %v\n", arr)
	fmt.Printf("after: %v\n", countingSort(arr))
}
