package main

import "sort"

//1300. 转变数组后最接近目标值的数组和
//给你一个整数数组arr 和一个目标值target ，请你返回一个整数value，使得将数组中所有大于value 的值变成value 后，数组的和最接近 target（最接近表示两者之差的绝对值最小）。
//
//如果有多种使得和最接近target的方案，请你返回这些整数中的最小值。
//
//请注意，答案不一定是arr 中的数字。
//
//
//示例 1：
//
//输入：arr = [4,9,3], target = 10
//输出：3
//解释：当选择 value 为 3 时，数组会变成 [3, 3, 3]，和为 9 ，这是最接近 target 的方案。
//示例 2：
//
//输入：arr = [2,3,5], target = 10
//输出：5
//示例 3：
//
//输入：arr = [60864,25176,27249,21296,20204], target = 56803
//输出：11361
//
//提示：
//
//1 <= arr.length <= 10^4
//1 <= arr[i], target <= 10^5

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	sum := target
	mid := sum / n
	for i := 0; i < n; i++ {
		if arr[i] > mid {
			mid = sum / (n - i)
			mod := sum % (n - i)
			if arr[i] >= mid {
				if mod*2 > n-i {
					return mid + 1
				}
				return mid
			}
		}
		sum -= arr[i]
	}
	return arr[n-1]

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(findBestValue([]int{1, 2, 23, 24, 34, 36}, 110))
	//println(findBestValue([]int{2,3,5}, 11))
}
