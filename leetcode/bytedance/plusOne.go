package main

import "fmt"

func plusOne(digits []int) []int {
	size := len(digits)
	for i := size - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] > 9 && i > 0 {
			digits[i] %= 10
		} else {
			break
		}
	}
	if digits[0] == 10 {
		digits[0] = 0
		result := make([]int, 0)
		result = append(result, 1)
		result = append(result, digits...)
		digits = result
	}
	return digits

}

func main() {
	print(fmt.Sprintf("%v", plusOne([]int{9, 9, 9})))
}
