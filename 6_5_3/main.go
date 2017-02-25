package main

import (
	"errors"
	"math"
	"fmt"
)

type MinHeap struct {
	arr      []int
	heapSize int
}

func NewMinHeap(arr []int) MinHeap {
	return MinHeap{arr: arr, heapSize: len(arr)}
}

func (h *MinHeap) Min() int {
	return h.arr[0]
}

func (h *MinHeap) ExtractMin() (int, error) {
	if h.heapSize < 1 {
		return int(0), errors.New("Очередь пуста")
	}

	min := h.arr[0]
	h.arr[0] = h.arr[h.heapSize-1]
	h.heapSize -= 1

	h.minHeapify(1)

	return min, nil
}

func (h *MinHeap) DecreaseKey(i, key int) error {
	if key > h.arr[i] {
		return errors.New("Новый ключ больше текущего")
	}

	h.arr[i] = key

	for i > 0 && h.arr[parent(i)] > h.arr[i] {
		h.arr[i], h.arr[parent(i)] = h.arr[parent(i)], h.arr[i]
		i = parent(i)
	}

	return nil
}

func (h *MinHeap) Insert(key int) {
	h.heapSize += 1

	min := math.MaxInt32
	if h.heapSize > len(h.arr) {
		h.arr = append(h.arr, min)
	} else {
		h.arr[h.heapSize-1] = min
	}

	h.DecreaseKey(h.heapSize-1, key)
}

func (h *MinHeap) MinHeap() []int {
	minHeap := h.arr[0:h.heapSize]

	return minHeap
}

func (h *MinHeap) minHeapify(i int) {
	l := left(i)
	r := right(i)

	smallest := i
	if l <= h.heapSize && h.arr[l-1] < h.arr[i-1] {
		smallest = l
	}

	if r <= h.heapSize && h.arr[r-1] < h.arr[smallest-1] {
		smallest = r
	}

	if smallest != i {
		buf := h.arr[i-1]
		h.arr[i-1] = h.arr[smallest-1]
		h.arr[smallest-1] = buf

		h.minHeapify(smallest)
	}
}

// left возвращает левый лист
func left(i int) int {
	return 2 * i
}

// возвращает правый лист
func right(i int) int {
	return 2*i + 1
}

// parent возвращает индекс вершины
func parent(i int) int {
	return i / 2
}

func main() {
	// Создание очереди, распечатка очереди и распечатка максимального
	minHeap := NewMinHeap([]int{10})
	fmt.Println("inserting...")
	minHeap.Insert(13)
	minHeap.Insert(16)
	minHeap.Insert(21)
	minHeap.Insert(33)
	minHeap.Insert(40)
	minHeap.Insert(44)
	minHeap.Insert(54)
	minHeap.Insert(55)
	minHeap.Insert(60)
	minHeap.Insert(90)

	fmt.Printf("Heap: %v\n", minHeap.MinHeap())
	fmt.Printf("Min: %v\n", minHeap.Min())

	// Извлечение максимального и распечатка без максимального
	max, _ := minHeap.ExtractMin()
	fmt.Printf("Extract min: %v\n", max)
	fmt.Printf("Heap: %v\n", minHeap.MinHeap())

	// Еще одно извлечение максимального и распечатка без этого
	max, _ = minHeap.ExtractMin()
	fmt.Printf("Extract min: %v\n", max)
	fmt.Printf("Heap: %v\n", minHeap.MinHeap())

	// Вставка нового элемента в очередь
	/*minHeap.Insert(5)
	fmt.Printf("Insert %v\n", 5)
	fmt.Printf("Heap: %v\n", minHeap.MinHeap())
	fmt.Printf("Min: %v\n", minHeap.Min())*/

	// Изменение приоритета у элемента с индексом 3
	/*minHeap.DecreaseKey(3, 31)
	fmt.Printf("Insert %v\n", 31)
	fmt.Printf("Heap: %v\n", minHeap.MinHeap())
	fmt.Printf("Min: %v\n", minHeap.Min())*/

	fmt.Println("sorted output...")
	lengh := len(minHeap.MinHeap())
	for i := 0; i < lengh; i++ {
		fmt.Println(minHeap.ExtractMin())
	}
}
