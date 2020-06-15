package main

func search2(nums []int, target int) int {
	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	}

	if mid > 0 {
		if index := search(nums[0:mid], target); index > -1 {
			return index
		}
	}
	if mid < len(nums)-1 {
		if index := search(nums[mid+1:], target); index > -1 {
			return mid + 1 + index
		}
	}

	return -1

}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	println(search([]int{-1, 0, 3, 5, 9, 12}, -2))
}
