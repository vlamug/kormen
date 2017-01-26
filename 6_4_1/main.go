package main

import (
	"fmt"
)

type HeapSortStruct struct {
	Arr []int
	HeapSize int
}

// HeapSort это пирамидальная сортировка.
// Сперва строится невозрастающее дерево.
// После этого так как в корне(в нулевом элементе массива) лежит самый максимальный, этот корневой элемент меняется
// с последним и больше не участвует в сотировке.
// Далее просиходит перестроение дерева без последнего элемента так, что пирамида была невозрастающей.
// На второй итерации также меняется нулевой элемент с предпоследним и снова происходит перестроение.
// Так продолжается до тех пор пока не останется один элемент, который будет самым минимальным и будет занимать нулевое
// значение.
// Готово!
func HeapSort(arr []int) []int {
	heapSortStruct := HeapSortStruct{Arr: arr}

	heapSortStruct = BuildMaxHeap(heapSortStruct)

	for i := len(heapSortStruct.Arr); i >= 2; i-- {
		heapSortStruct.Arr[0], heapSortStruct.Arr[i-1] = heapSortStruct.Arr[i-1], heapSortStruct.Arr[0]
		heapSortStruct.HeapSize = heapSortStruct.HeapSize - 1
		heapSortStruct = MaxHeapify(heapSortStruct, 1)
	}

	return heapSortStruct.Arr
}

// BuildMaxHeap строит невозрастающую пирамиду
// Цикл для построения состоит из половины от длины массива итераций. При этом каждый индекс в цикле это узел и он
// является корнем.
// Время работы данного алгоритма: каждый вызов функции MaxHepify имеет сложность O(lgn), а всего имеется O(n) таких
// вызовов, следовательно время алгоритма равно O(nlgn).
func BuildMaxHeap(heapSortStruct HeapSortStruct) HeapSortStruct {
	heapSortStruct.HeapSize = len(heapSortStruct.Arr)

	for i := len(heapSortStruct.Arr)/2; i >= 1; i-- {
		heapSortStruct = MaxHeapify(heapSortStruct, i)
	}

	return heapSortStruct
}

// MaxHeapify поддерживает свойство невозрастающей пирамиды.
// В данном случае значение вершины должно быть меньше значений листьев этой вершины.
func MaxHeapify(heapSortStruct HeapSortStruct, i int) HeapSortStruct {
	l := left(i)
	r := right(i)

	largest := i
	if l <= heapSortStruct.HeapSize && heapSortStruct.Arr[l-1] > heapSortStruct.Arr[i-1] {
		largest = l
	}

	if r <= heapSortStruct.HeapSize && heapSortStruct.Arr[r-1] > heapSortStruct.Arr[largest-1] {
		largest = r
	}

	if largest != i {
		buf := heapSortStruct.Arr[i-1]
		heapSortStruct.Arr[i-1] = heapSortStruct.Arr[largest-1]
		heapSortStruct.Arr[largest-1] = buf

		heapSortStruct = MaxHeapify(heapSortStruct, largest)
	}

	return heapSortStruct
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
	// result: [1 2 3 4 7 8 9 10 14 16]
	arr = HeapSort(arr)
	fmt.Println(arr)
}