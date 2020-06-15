package main

import "fmt"

// 滑动窗口最大值 ，双端队列

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}
	win := make([]int, 0)
	result := make([]int, 0)
	for i, v := range nums {
		if i >= k && win[0] <= i-k {
			win = win[1:]
		}
		for len(win) > 0 && nums[win[len(win)-1]] < v {
			win = win[:len(win)-1]
		}
		//for j := 0; j < len(win); j++ {
		//	if nums[win[j]] < v {
		//		win = win[:j]
		//		break
		//	}
		//}
		win = append(win, i)
		if i >= k-1 {
			result = append(result, nums[win[0]])
		}
	}
	return result

}

func main() {
	println(fmt.Sprintf("%v", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)))
}
