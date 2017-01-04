package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"errors"
)

// ---------------------------------------------------------------------------------------------------------------------
// Алгоритм бинарного поиска ("binary search"). Построен на принципе "разделя-и-властвуй". Данный алгоритм используя
// рекусрию на каждом этапе вызова функции определяет в какой половинке массива располагается элемент до тех пор, пока
// массив не сузится до 2-ух элементов. После этого из двух элементов выбирается искомый, иначе выдается сообщение об
// отсутсвии элемента в массиве.
// На вход необходимо подавать два параметра:
// 1. отсортированный массив по возрастанию, элементы которого необходимо вводить через пробел,
// 2. искомый элемент
// ---------------------------------------------------------------------------------------------------------------------

func main() {
	fmt.Println("enter array elements separated by commas:")
	reader := bufio.NewReader(os.Stdin)
	arrStr, _ := reader.ReadString('\n')

	fmt.Println("enter search item:")
	searchNumberStr, _ := reader.ReadString('\n')

	arr := strToArr(arrStr)
	searchNumber := strToInt(searchNumberStr)

	index, err := binarySearch(searchNumber, arr, 0, len(arr))
	if err != nil {
		fmt.Printf("error while search number: %s\n", err)
	} else {
		fmt.Printf("index of founded number: %d\n", index)
	}
}

func strToArr(arrStr string) []int {
	var arr []int

	for _, value := range strings.Split(arrStr, " ") {
		num, err := strconv.Atoi(strings.Trim(value, "\n"))
		if err != nil {
			fmt.Printf("Error by convert string to ints array. Error: %s\n", err)
			os.Exit(1)
		}

		arr = append(arr, num)
	}

	return arr
}

func strToInt(searchNumberStr string) int {
	searchNumber, err := strconv.Atoi(strings.Trim(searchNumberStr, "\n"))
	if err != nil {
		fmt.Printf("Error by convert string to int. Error: %s\n", err)
		os.Exit(1)
	}

	return searchNumber
}

func binarySearch(searchNumber int, arr []int, start int, end int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}

	if (end - start) <= 2 {
		if searchNumber == arr[start] {
			return start, nil
		} else if searchNumber == arr[end - 1] {
			return end - 1, nil
		} else {
			return 0, errors.New("search numbder do not exists")
		}
	}

	middle := ((end - start) / 2) + start

	if searchNumber >= arr[middle] {
		return binarySearch(searchNumber, arr, middle, end)
	} else {
		return binarySearch(searchNumber, arr, start, middle)
	}
}
