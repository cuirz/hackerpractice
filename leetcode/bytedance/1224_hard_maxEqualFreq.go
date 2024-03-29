package main

//1224. 最大相等频率
//给你一个正整数数组nums，请你帮忙从该数组中找出能满足下面要求的 最长 前缀，并返回该前缀的长度：
//
//从前缀中 恰好删除一个 元素后，剩下每个数字的出现次数都相同。
//如果删除这个元素后没有剩余元素存在，仍可认为每个数字都具有相同的出现次数（也就是 0 次）。
//
//
//
//示例 1：
//
//输入：nums = [2,2,1,1,5,3,3,5]
//输出：7
//解释：对于长度为 7 的子数组 [2,2,1,1,5,3,3]，如果我们从中删去 nums[4] = 5，就可以得到 [2,2,1,1,3,3]，里面每个数字都出现了两次。
//示例 2：
//
//输入：nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
//输出：13
//
//
//提示：
//
//2 <= nums.length <= 10^5
//1 <= nums[i] <= 10^5

func maxEqualFreq(nums []int) int {
	freq := map[int]int{}
	count := map[int]int{}
	var result, maxFreq int
	for i, v := range nums {
		if count[v] > 0 {
			freq[count[v]]--
		}
		count[v]++
		maxFreq = max(maxFreq, count[v])
		freq[count[v]]++
		if maxFreq == 1 ||
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			result = max(result, i+1)
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(maxEqualFreq([]int{1, 1, 1}))
}
