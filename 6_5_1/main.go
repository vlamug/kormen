package main

import (
	"errors"
	"fmt"
)

// MaxHeap структура очереди с приоритетом
type MaxHeap struct {
	arr []int
	heapSize int
}

// NewMaxHeap создает новую структуру очереди с приоритетом
func NewMaxHeap(arr []int) MaxHeap {
	return MaxHeap{arr: arr, heapSize: len(arr)}
}

// Max возвращает максимальный элемент из очереди. Элемент не удаляется из очереди
func (h *MaxHeap) Max() int {
	return h.arr[0]
}

// ExtractMax извлекает максимальный элемент из очереди. Элемент удаляется из очереди
func (h *MaxHeap) ExtractMax() (int, error) {
	if h.heapSize < 1 {
		return int(0), errors.New("Очередь пуста")
	}

	max := h.arr[0]
	h.arr[0] = h.arr[h.heapSize-1]
	h.heapSize -= 1

	h.maxHeapify(1)

	return max, nil
}

// IncreaseKey повышает приоритет в очереди
func (h *MaxHeap) IncreaseKey(i int, key int) error {
	if key < h.arr[i] {
		return errors.New("Новый ключ меньше текущего")
	}

	h.arr[i] = key

	for i > 0 && h.arr[parent(i)] < h.arr[i] {
		h.arr[i], h.arr[parent(i)] = h.arr[parent(i)], h.arr[i]
		i = parent(i)
	}

	return nil
}

func (h *MaxHeap) Insert(key int) {
	h.heapSize += 1

	h.arr = append(h.arr, 0)
	h.IncreaseKey(h.heapSize, key)
}

// maxHeapify поддерживает свойство невозрастающей пирамиды.
// В данном случае значение вершины должно быть меньше значений листьев этой вершины.
func (h *MaxHeap) maxHeapify(i int) {
	l := left(i)
	r := right(i)

	largest := i
	if l <= h.heapSize && h.arr[l-1] > h.arr[i-1] {
		largest = l
	}

	if r <= h.heapSize && h.arr[r-1] > h.arr[largest-1] {
		largest = r
	}

	if largest != i {
		buf := h.arr[i-1]
		h.arr[i-1] = h.arr[largest-1]
		h.arr[largest-1] = buf

		h.maxHeapify(largest)
	}
}

// left возвращает левый лист
func left(i int) int {
	return 2*i
}

// возвращает правый лист
func right(i int) int {
	return 2*i + 1
}

// parent возвращает индекс вершины
func parent(i int) int {
	return i/2
}

func main() {
	// Создание очереди, распечатка очереди и распечатка максимального
	maxHeap := NewMaxHeap([]int{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1})
	fmt.Printf("Heap: %v\n", maxHeap.arr)
	fmt.Printf("Max: %v\n", maxHeap.Max())

	// Извлечение максимального и распечатка без максимального
	max, _ := maxHeap.ExtractMax()
	fmt.Printf("Extract max: %v\n", max)
	fmt.Printf("Heap: %v\n", maxHeap.arr)

	// Еще одно извлечение максимального и распечатка без этого
	max, _ = maxHeap.ExtractMax()
	fmt.Printf("Extract max: %v\n", max)
	fmt.Printf("Heap: %v\n", maxHeap.arr)

	// Вставка нового элемента в очередь
	maxHeap.Insert(30)
	fmt.Printf("Heap: %v\n", maxHeap.arr)
	fmt.Printf("Max: %v\n", maxHeap.Max())

	// Изменение приоритета у элемента с индексом 3
	maxHeap.IncreaseKey(3, 31)
	fmt.Printf("Heap: %v\n", maxHeap.arr)
	fmt.Printf("Max: %v\n", maxHeap.Max())
}
