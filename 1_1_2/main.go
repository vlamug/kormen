package main

import "fmt"

// insertionSort sorts the array descending.
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		buf := arr[i]

		for j >= 0 && buf > arr[j] {
			arr[j+1] = arr[j]
			j = j-1
		}

		arr[j+1] = buf
	}

	return arr
}

func main() {
	arr := []int{3, 90, 10, 11, 0, 56, 78, 4}
	fmt.Printf("unsoreted array: %+v\n", arr)

	sortedArr := insertionSort(arr)
	fmt.Printf("sorted array: %+v\n", sortedArr)
}
