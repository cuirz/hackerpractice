package main

//962. 最大宽度坡
//给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
//
//找出 A 中的坡的最大宽度，如果不存在，返回 0 。
//
//
//
//示例 1：
//
//输入：[6,0,8,2,1,5]
//输出：4
//解释：
//最大宽度的坡为 (i, j) = (1, 5): A[1] = 0 且 A[5] = 5.
//示例 2：
//
//输入：[9,8,1,0,1,9,4,0,4,1]
//输出：7
//解释：
//最大宽度的坡为 (i, j) = (2, 9): A[2] = 1 且 A[9] = 1.

//思路 1 ： 滑动窗口，双指针
//从左右两边开始制定指针，当寻找到底一个符合条件斜坡后，左右指针形成一个滑动窗口, pass比 r小的值

// 思路 2： 单调递减函数 排序
//单调栈应用范围
//求解数组中元素右边第一个比它小的元素的下标，从前往后，构造单调递增栈
//求解数组中元素右边第一个比它大的元素的下标，从前往后，构造单调递减栈
//求解数组中元素左边第一个比它小的元素的下标，从后往前，构造单调递减栈
//求解数组中元素左边第一个比它小的元素的下标，从后往前，构造单调递增栈

func maxWidthRamp(A []int) int {
	n := len(A)
	stack := make([]int, 0)
	stack = append(stack, 0)
	for i := 1; i < n; i++ {
		if A[i] <= A[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}

	result := 0
	for i := n - 1; i > -1; i-- {
		index := len(stack) - 1
		for index > -1 && A[i] >= A[stack[index]] {
			result = max(result, i-stack[index])
			stack = stack[:index]
			index = len(stack) - 1
		}
	}
	return result

}

func maxWidthRampOne(A []int) int {
	l := 0
	r := len(A) - 1
	width := 0
	for r > width {
		if A[l] <= A[r] {
			width = max(width, r-l)
			if l == 0 {
				return width
			}
		}
		l++
		if r-l < width {
			big := A[r]
			for r > 0 && A[r] <= big {
				r--
			}
			r--
			l = 0
		}
	}
	return width
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
	//println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
	println(maxWidthRamp([]int{3, 4, 1, 2}))
}
