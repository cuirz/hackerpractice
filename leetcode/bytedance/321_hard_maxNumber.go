package main

//321. 拼接最大数
//给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n) 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
//
//求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
//
//说明: 请尽可能地优化你算法的时间和空间复杂度。
//
//示例 1:
//
//输入:
//nums1 = [3, 4, 6, 5]
//nums2 = [9, 1, 2, 5, 8, 3]
//k = 5
//输出:
//[9, 8, 6, 5, 3]

//单调栈

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	result := make([]int, 0)
	start := 0
	if k > len(nums2){
		start = k - len(nums2)
	}

	biggest := func(array []int, s int) []int {
		n := len(array)
		stack := make([]int, 0)
		for i, v := range array {
			for len(stack) > 0 && n-i-1 >= s-len(stack) && v > stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) < s {
				stack = append(stack, v)
			}
		}
		return stack
	}
	compareLess := func(x, y []int) bool {
		for i := 0; i < len(x) && i < len(y); i++ {
			if x[i] != y[i] {
				return x[i] < y[i]
			}
		}
		return len(x) < len(y)
	}
	merge := func(x, y []int) []int {
		res := make([]int, len(x)+len(y))
		for i, _ := range res {
			if compareLess(x,y)	{
				res[i],y = y[0],y[1:]
			}else{
				res[i],x = x[0],x[1:]
			}
		}
		return res
	}

	for i := start; i <= len(nums1)&& i <= k; i++ {
		x := biggest(nums1, i)
		y := biggest(nums2, k-i)
		res := merge(x, y)
		if compareLess(result,res){
			result = res
		}
	}
	return result
}

