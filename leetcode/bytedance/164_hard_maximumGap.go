package main

import "sort"

//164. 最大间距
//给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
//
//如果数组元素个数小于 2，则返回 0。
//
//示例1:
//
//输入: [3,6,9,1]
//输出: 3
//解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
//示例2:
//
//输入: [10]
//输出: 0
//解释: 数组元素个数小于 2，因此返回 0。
//说明:
//
//你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
//请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。

//思路1 直接用排序，暴力找出
//思路2 使用基础桶的算法
func maximumGap(nums []int) int {
	sort.Ints(nums)

	result, pre := 0, 0
	if len(nums) > 0 {
		pre = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		result = max(result, nums[i]-pre)
		pre = nums[i]
	}
	return result

}

func maximumGapBase(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	buf := make([]int, n)
	maxVal := maxarry(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			println(cnt[digit] - 1)
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		//for _,v :=range nums{
		//	digit := v / exp % 10
		//	buf[cnt[digit]-1] = v
		//	cnt[digit]--
		//}
		copy(nums, buf)
	}

	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return
}

func maxarry(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(maximumGapBase([]int{73, 22, 93, 43, 55, 14, 28, 65, 39, 81, 99, 45, 70}))
}