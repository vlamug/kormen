package main

import "fmt"

// BuildMaxHeap строит невозрастающую пирамиду
// Цикл для построения состоит из половины от длины массива итераций. При этом каждый индекс в цикле это узел и он
// является корнем.
// Время работы данного алгоритма: каждый вызов функции MaxHepify имеет сложность O(lgn), а всего имеется O(n) таких
// вызовов, следовательно время алгоритма равно O(nlgn).
func BuildMaxHeap(arr []int) []int {
	for i := len(arr)/2; i >= 1; i-- {
		arr = MaxHeapify(arr, i)
	}

	return arr
}

// MaxHeapify поддерживает свойство невозрастающей пирамиды.
// В данном случае значение вершины должно быть меньше значений листьев этой вершины.
func MaxHeapify(arr []int, i int) []int {
	l := left(i)
	r := right(i)

	largest := i
	if l <= len(arr) && arr[l-1] > arr[i-1] {
		largest = l
	}

	if r <= len(arr) && arr[r-1] > arr[largest-1] {
		largest = r
	}

	if largest != i {
		buf := arr[i-1]
		arr[i-1] = arr[largest-1]
		arr[largest-1] = buf

		arr = MaxHeapify(arr, largest)
	}

	return arr
}

// parent возвращает индекс вершины
func parent(i int) int {
	return i/2
}

// left возвращает левый лист
func left(i int) int {
	return 2*i
}

// возвращает правый лист
func right(i int) int {
	return 2*i + 1
}

func main() {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fmt.Println(arr)
	// result: [16, 14, 10, 8, 7, 9, 3, 2, 4, 1]
	arr = BuildMaxHeap(arr)
	fmt.Println(arr)

	fmt.Println("another exmaple:")

	arr = []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	fmt.Println(arr)
	arr = BuildMaxHeap(arr)
	fmt.Println(arr)
}
