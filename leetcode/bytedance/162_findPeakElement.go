package main

func findPeakElement(nums []int) int {
	var search func(l, r int) int
	search = func(l, r int) int {
		if l == r {
			return l
		}
		mid := (r + l) / 2
		if nums[mid+1] > nums[mid] {
			return search(mid+1, r)
		}
		return search(l, r)

	}
	return search(0, len(nums)-1)
}
