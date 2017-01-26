package util

import (
	"strings"
	"strconv"
)

// StrToArr преобразует строку вида "4 6 2 9 1 90 33 0" в массив вида "int[]{4, 6, 2, 9, 1, 90, 33, 0}"
func StrToArr(arrStr string) ([]int, error) {
	var arr []int

	for _, value := range strings.Split(arrStr, " ") {
		num, err := strconv.Atoi(strings.Trim(value, "\n"))
		if err != nil {
			return []int{}, err
		}

		arr = append(arr, num)
	}

	return arr, nil
}
