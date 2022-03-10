package main

//75. 颜色分类
//给定一个包含红色、白色和蓝色，一共n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
//
//注意:
//不能使用代码库中的排序函数来解决这道题。
//
//示例:
//
//输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]
//进阶：
//
//一个直观的解决方案是使用计数排序的两趟扫描算法。
//首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
//你能想出一个仅使用常数空间的一趟扫描算法吗？

//思路 荷兰旗问题
// 三色 左，右分配2种色以后，剩下的另一种就在中间
func sortColors(nums []int) {
	n := len(nums)
	seek, l, r := 0, 0, n-1
	for seek <= r {
		if nums[seek] == 0 {
			nums[l], nums[seek] = nums[seek], nums[l]
			seek++
			l++
		} else if nums[seek] == 2 {
			nums[r], nums[seek] = nums[seek], nums[r]
			r--
		} else {
			seek++
		}
	}
}

func main() {
	sortColors([]int{
		0, 1, 0, 2, 1, 0, 1, 2,
		//0, 2,
	})
}
