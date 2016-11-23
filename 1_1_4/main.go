package main

import (
	"fmt"
	"strconv"
	"os"
)

// parseNumber parses number as string into number as array.
func parseNumber(number string) ([]int, error) {
	var numberArr []int
	for i := 0; i < len(number); i++ {
		number, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return []int{}, err
		}

		numberArr = append(numberArr, number)
	}

	return numberArr, nil
}

// sumNumbers sums two numbers, which are represented as arrays.
func sumNumbers(firstNumber []int , secondNumber []int) []int {
	var (
		longNumber []int
		shortNumber []int
	)

	if len(firstNumber) > len(secondNumber) {
		longNumber = firstNumber
		shortNumber = secondNumber
	} else {
		longNumber = secondNumber
		shortNumber = firstNumber
	}

	remainder := 0
	sum := 0
	sumOfNumbers := make([]int, len(longNumber))
	for i := 0; i < len(longNumber); i++ {
		if (len(shortNumber) - i - 1) >= 0 {
			sum = remainder + longNumber[len(longNumber) - i - 1] + shortNumber[len(shortNumber) - i - 1]
		} else {
			sum = remainder + longNumber[len(longNumber) - i - 1]
		}

		sumOfNumbers[len(longNumber) - i - 1] = sum % 10

		if sum >= 10 {
			remainder = 1
		} else {
			remainder = 0
		}
	}

	return sumOfNumbers
}

func main() {
	var (
		firstNumber string
		secondNumber string
	)

	fmt.Printf("enter the first number: ")
	fmt.Scanf("%s\n", &firstNumber)

	fmt.Printf("enter the second number: ")
	fmt.Scanf("%s\n", &secondNumber)

	firstNumberArr, err := parseNumber(firstNumber)
	if err != nil {
		fmt.Printf("cant parse number: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("first number in array: %+v\n", firstNumberArr)

	secondNumberArr, err := parseNumber(secondNumber)
	if err != nil {
		fmt.Printf("cant parse number: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("second number in array: %+v\n", secondNumberArr)

	fmt.Printf("sum of these numbers: %+v\n", sumNumbers(firstNumberArr, secondNumberArr))
}
