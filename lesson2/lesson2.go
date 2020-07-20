package main

import (
	"fmt"
	"math"
)

func main() {
}

func isEvenNumber(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func isMultipleOf3(num int) bool {
	if num%3 == 0 {
		return true
	}
	return false
}

func showFibonachchi() {
	result := make([]int64, 100)
	result[1] = 1
	for i := 2; i < 100; i++ {
		result[i] = result[i-2] + result[i-1]
	}
	fmt.Println(result)
}

func showSimpleNumbers100() {
	result := make([]int, 100)
	var (
		i = 0
		n = 2
	)
	for {
		if isNumberSimple(n) {
			result[i] = n
			i++
		}
		n++
		if i == 100 {
			break
		}
	}
	fmt.Println(result)
}

func isNumberSimple(number int) bool {
	number = int(math.Abs(float64(number)))
	if number < 3 {
		return true
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
