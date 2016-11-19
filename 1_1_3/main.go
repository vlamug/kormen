package main

import (
	"fmt"
	"os"
)

// linearSearch searches the entered number linearly
func linearSearch(arr []int, search int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			return i
		}
	}

	return -1
}

func main() {
	length := 0
	fmt.Print("enter length of array: ")
	_, err := fmt.Scanf("%d\n", &length)
	if err != nil {
		fmt.Printf("error of typing array length: %s\n", err)
		os.Exit(1)
	}

	num := 0
	arr := []int{}
	for i := 0; i < length; i++ {
		fmt.Scanf("%d\n", &num)
		if err != nil {
			fmt.Printf("error of typing array item: %s\n", err)
			os.Exit(1)
		}

		arr = append(arr, num)
	}

	searchNum := 0
	fmt.Print("enter search number: ")
	fmt.Scanf("%d\n", &searchNum)
	if err != nil {
		fmt.Printf("error of typing search item: %s\n", err)
		os.Exit(1)
	}

	position := linearSearch(arr, searchNum)
	if position == -1 {
		fmt.Println("item not found")
	} else {
		fmt.Printf("item is located at position: %d\n", position)
	}
}
