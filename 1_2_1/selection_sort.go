package main

import (
	"fmt"
)

// selectionSortByAsc sorts array by asc using additional array.
func selectionSortByAsc(arr []int) []int {
	var (
		sortedArr []int
		length = len(arr)
	)

	for i := 0; i < length; i++ {
		min := arr[0]
		index := 0
		for j := 0; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				index = j
			}
		}

		arr = append(arr[:index], arr[index+1:]...)
		sortedArr = append(sortedArr, min)
	}

	return sortedArr
}

func main() {
	var lengthOfArray = 0

	fmt.Printf("enter length of array: ")
	fmt.Scanf("%d", &lengthOfArray)

	fmt.Println("enter items of array:")

	var unsortedArr []int
	var enteredNumber int
	for i := 0; i < lengthOfArray; i++ {
		fmt.Scanf("%d", &enteredNumber)

		unsortedArr = append(unsortedArr, enteredNumber)
	}

	fmt.Println(selectionSortByAsc(unsortedArr))
}
