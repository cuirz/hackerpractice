package main

//659. 分割数组为连续子序列
//给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个子序列，其中每个子序列都由连续整数组成且长度至少为 3 。
//
//如果可以完成上述分割，则返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//输入: [1,2,3,3,4,5]
//输出: True
//解释:
//你可以分割出这样两个连续子序列 :
//1, 2, 3
//3, 4, 5
//提示：
//
//输入的数组长度范围为 [1, 10000]

//贪心
func isPossible(nums []int) bool {
	count := make(map[int]int)
	bucket := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}

	for _, v := range nums {
		if count[v] == 0 {
			continue
		}
		if count[v] > 0 && bucket[v-1] > 0 {
			bucket[v-1]--
			bucket[v]++
			count[v]--
		} else if count[v] > 0 && count[v+1] > 0 && count[v+2] > 0 {
			count[v]--
			count[v+1]--
			count[v+2]--
			bucket[v+2]++
		} else {
			return false
		}
	}
	return true
}

//分桶
func isPossible2(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	bucket := make([][]int, 0)
	for _, v := range nums {
		isSkip := false
		for i := len(bucket) - 1; i > -1; i-- {
			s := len(bucket[i])
			if v == bucket[i][s-1]+1 {
				bucket[i] = append(bucket[i], v)
				isSkip = true
				break
			}
		}
		if !isSkip {
			bucket = append(bucket, []int{v})
		}
	}

	for _, v := range bucket {
		if len(v) < 3 {
			return false
		}
	}
	return true
}
func main() {
	println(isPossible([]int{1, 2, 3, 3, 4, 5}))
}
