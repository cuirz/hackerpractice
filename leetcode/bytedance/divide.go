package main

import "math"

func divide(dividend int, divisor int) int {
	minus := false
	if dividend < 0 {
		dividend = -dividend
		minus = !minus
	}
	if divisor < 0 {
		divisor = -divisor
		minus = !minus
	}
	result := div(dividend, divisor)
	if minus {
		result = -result
	}
	if result < math.MinInt32 {
		return math.MinInt32
	} else if result > math.MaxInt32 {
		return math.MaxInt32
	}
	return result

}

func div(x, y int) int {
	if x < y {
		return 0
	}
	count := 1
	s := y
	for x-s >= s {
		count <<= 1
		s <<= 1
	}
	return count + div(x-s, y)

}

func main() {
	//8
	//1
	println(divide(-2147483648, -1))
}
