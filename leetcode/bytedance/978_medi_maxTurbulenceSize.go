package main

//978. 最长湍流子数组
//当 A的子数组A[i], A[i+1], ..., A[j]满足下列条件时，我们称其为湍流子数组：
//
//若i <= k < j，当 k为奇数时，A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
//或 若i <= k < j，当 k 为偶数时，A[k] > A[k+1]，且当 k为奇数时，A[k] < A[k+1]。
//也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。
//
//返回 A 的最大湍流子数组的长度。
//
//
//
//示例 1：
//
//输入：[9,4,2,10,7,8,8,1,9]
//输出：5
//解释：(A[1] > A[2] < A[3] > A[4] < A[5])
//示例 2：
//
//输入：[4,8,12,16]
//输出：2
//示例 3：
//
//输入：[100]
//输出：1
//
//
//提示：
//
//1 <= A.length <= 40000
//0 <= A[i] <= 10^9

//思路 双指针
func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	l, r := 0, 0
	result := 1
	for r < n-1 {
		if l == r {
			if arr[l] == arr[l+1] {
				l++
			}
			r++
		} else {
			if (arr[r-1] < arr[r] && arr[r] > arr[r+1]) || (arr[r-1] > arr[r] && arr[r] < arr[r+1]) {
				r++
			} else {
				l = r
			}
		}

		result = max(result, r-l+1)
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
	//println(maxTurbulenceSize([]int{0,8,45,88,48,68,28,55,17,24}))
	println(maxTurbulenceSize([]int{0, 8}))
}
