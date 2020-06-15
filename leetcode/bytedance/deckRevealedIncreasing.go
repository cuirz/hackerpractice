package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Slice(deck, func(i, j int) bool {
		if deck[i] > deck[j] {
			return true
		}
		return false
	})
	// n := len(deck)
	result := make([]int, 0)

	for _, v := range deck {
		m := len(result)
		if m > 0 {
			result = append([]int{result[m-1]}, result[:m-1]...)
		}

		result = append([]int{v}, result...)
	}
	return result

}

func main() {
	deckRevealedIncreasing([]int{3, 4, 1, 2, 6})
}
