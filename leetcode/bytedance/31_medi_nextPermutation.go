package main

//31. 下一个排列
//实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须原地修改，只允许使用额外常数空间。
//
//以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1

//思路
func nextPermutation(nums []int) {
	n := len(nums)
	index := 0
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {

			for j := i; j < n; j++ {
				if nums[i-1] < nums[j] {
					index = j
				} else {
					break
				}
			}
			nums[i-1], nums[index] = nums[index], nums[i-1]
			index = i
			break
		}
	}

	var end int
	for i := index; i < index+(n-index)>>1; i++ {
		end = n - 1 - i + index
		nums[i], nums[end] = nums[end], nums[i]
	}
}

func main() {

	nextPermutation([]int{1, 3, 2})
}
