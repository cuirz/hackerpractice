package main

//560. 和为K的子数组
//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

func subarraySum(nums []int, k int) int {
	dic := make(map[int]int)
	pre := 0
	dic[0] = 1
	count := 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := dic[pre-k]; ok {
			count += dic[pre-k]
		}
		dic[pre]++
	}
	return count
}

func main() {
	println(subarraySum([]int{1, 2, 3}, 3))
}
