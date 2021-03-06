package main

import (
	"fmt"
)

// ---------------------------------------------------------------------------------------------------------------------
// Алгоритм сортировка слиянием("merge sort"). Построен на принципе "разделя-и-властвуй". Данный алгоритм используя
// рекурсию сперва делит массив пополам до тех пор, пока массив не будет состоять из одного элемента. А массив из
// одного элемента является отсортированным. Затем происходит слияние двух разделенных массивов. При слиянии создается
// дополнительный буферный массив, в который и выбираются из двух массивов эелементы по возрастанию или по убыванию.
// ---------------------------------------------------------------------------------------------------------------------

func mergeSort(mas []int, p, r int) {
	if p < r {
		// Определение позиции по которой делается деление массива.
		q := (p + r) / 2
		// Сортируем левую половинку массива.
		mergeSort(mas, p, q)
		// Сортируем правую половинку массива.
		mergeSort(mas, q + 1, r)
		// Сливаем две отсортированные половинки массива.
		merge(mas, p, q, r)
	}
}

func merge(mas []int, p, q, r int) {
	// Буферный массив, в который из двух отсортированных половинок массива будут ложится элементы также сортируясь.
	var sortedMass []int
	var firstIndex = p
	var secondIndex = q + 1

	// Из двух массивов собираются эелменты в буферный массив.
	for firstIndex <= q && secondIndex <= r {
		if mas[firstIndex-1] < mas[secondIndex-1] {
			sortedMass = append(sortedMass, mas[firstIndex-1])
			firstIndex++
		} else {
			sortedMass = append(sortedMass, mas[secondIndex-1])
			secondIndex++
		}
	}

	// Дописываются элементы в буферный массив.
	for i := firstIndex-1; i < q; i++ {
		sortedMass = append(sortedMass, mas[i])
	}
	for i := secondIndex-1; i < r; i++ {
		sortedMass = append(sortedMass, mas[i])
	}

	// Отсортированный буферный массив записываем в основной массив.
	var index = 0
	for i := p-1; i < r; i++ {
		mas[i] = sortedMass[index]
		index++
	}
}

func main() {
	mas := []int{333, 40, 52, 26, 38, 57, 9, 49}
	fmt.Println("Исходный массив: ", mas)

	mergeSort(mas, 1, len(mas))

	fmt.Println("Отсортированный массив: ", mas)
}
