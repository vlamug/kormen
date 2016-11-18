package main

import (
	"fmt"
)

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		buf := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > buf {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = buf
	}

	return arr
}

func main() {
	arr := []int{4, 1, 45, 90, 2, 10}
	fmt.Printf("unsorted array: %+v\n", arr)

	sortedArr := insertionSort(arr)
	fmt.Printf("sorted array: %+v\n", sortedArr)
}
