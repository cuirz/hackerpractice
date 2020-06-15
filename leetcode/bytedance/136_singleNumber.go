package main

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

func singleNumber(nums []int) int {
	// 均出现两次是 point  异或会中合掉相同数字
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
