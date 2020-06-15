package main

func minSubArrayLen(s int, nums []int) int {
	array := make([]int, 0)
	result := 0
	sum := 0
	for _, w := range nums {
		array = append(array, w)
		sum += w
		for sum >= s {
			if result == 0 || result > len(array) {
				result = len(array)
			}
			sum -= array[0]
			array = array[1:]

		}

	}
	return result
}

func main() {
	minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
}
