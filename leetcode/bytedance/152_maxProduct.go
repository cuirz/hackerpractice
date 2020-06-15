package main

//152. 乘积最大子数组
//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

//动态规划
// 思路：类似 53.最大子序和，但区别是有负数情况下 负负得正最大值根据负数最小值而变化
// 状态转移时判断前子数组乘积最大、最小值和当前值乘积
// fmax(i) = 最大值(max[i-1]*nums[i],min[i-1]*nums[i],nums[i])
// fmin(i) = 最小值(max[i-1]*nums[i],min[i-1]*nums[i],nums[i])

func maxProduct(nums []int) int {
	n := len(nums)

	maxv := make([]int, n)
	minv := make([]int, n)
	var result int
	result, maxv[0], minv[0] = nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		maxv[i], minv[i] = maxmin(maxv[i-1]*nums[i], minv[i-1]*nums[i], nums[i])
		if result < maxv[i] {
			result = maxv[i]
		}

	}
	return result
}

func maxmin(x, y, z int) (int, int) {
	if x > y {
		if y > z {
			return x, z
		} else {
			if z > x {
				return z, y
			}
			return x, y
		}
	} else {
		if x > z {
			return y, z
		} else {
			if z > y {
				return z, x
			}
			return y, x
		}
	}
}
