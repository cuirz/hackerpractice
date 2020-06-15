package main

import "math"

func lengthOfLongestSubstringTwoDistinctBest(s string) int {
	n := len(s)
	if n < 3 {
		return n
	}
	max := 2
	l, r := 0, 0
	mmap := make(map[byte]int, 3)
	for r < n {
		if len(mmap) < 3 {
			mmap[s[r]] = r
			r++
		}
		if len(mmap) == 3 {
			index := math.MaxInt32
			for _, v := range mmap {
				index = Min(index, v)
			}
			delete(mmap, s[index])
			l = index + 1
		}
		max = Max(max, r-l)
	}
	return max
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	l := 0
	max := 0
	array := make(map[byte]bool, 3)
	array[s[l]] = true
	for r := 1; r < len(s); r++ {
		if array[s[r]] {
			max = Max(max, r-l+1)
			continue
		}
		array[s[r]] = true
		if len(array) > 2 {
			tmp := r - 1
			for l = tmp - 1; l >= 0; l-- {
				if s[l] != s[tmp] {
					delete(array, s[l])
					l += 1
					break
				}
			}
		}
		max = Max(max, r-l+1)
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(lengthOfLongestSubstringTwoDistinctBest("abc"))
	println(lengthOfLongestSubstringTwoDistinct("abc"))
}
