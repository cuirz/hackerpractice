package main

import "sort"

//61. 扑克牌中的顺子
//从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

func isStraight(nums []int) bool {
	sort.Ints(nums)
	n := len(nums)
	min := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			min++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[n-1]-nums[min] < 5

}

func main() {
	println(isStraight([]int{0, 0, 4, 6, 8}))
}
