package main

import (
	"fmt"
)

// ---------------------------------------------------------------------------------------------------------------------
// quickSort сортирует массив используя алгоритм быстрой сортировки.
// В качестве опорного элемента берется последний элемент массива.
// Данный алгоритм, как и алгоритм сортировки слиянием использует принцип разделяй и влавствуй.
// В функции partition мы можем понять, что сперва мы берем опорный элемент, который является последним элементов
// массива, после этого все элемент, которые больше этого опорного элемент перемещаются вправ, а все остальные
// соответсвенно влево. Таким образом этот опрный элемент оказывается элемент на основе которого массив делится на две
// части и далее для двух часте массива происходит тоже самоей: берется опорный элемент, большие элементы вправо,
// меньшие влево и так далее до тех пор, пока сиходный массив не разделится до одного элемент. А массив из одного
// элемента является отсортированным массивом.
// ---------------------------------------------------------------------------------------------------------------------
func quickSort(arr []int, start int, end int) []int {
	if start < end {
		arr, q := partition(arr, start, end)
		arr = quickSort(arr, start, q)
		arr = quickSort(arr, q+1, end)
	}

	return arr
}

func partition(arr []int, start int, end int) ([]int, int) {
	x := arr[end-1]
	i := start-1

	for j := start; j < end-1; j++ {
		if start == 0 && end == 6 {
		}
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end-1] = arr[end-1], arr[i+1]

	return arr, i+1
}

func main() {
	arr := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	fmt.Printf("before sort: %v\n", arr)
	arr = quickSort(arr, 0, len(arr))
	fmt.Printf("after sort: %v\n", arr)
}
