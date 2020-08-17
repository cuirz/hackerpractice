package main

import "sort"

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2{
		return 0
	}
	sort.Ints(nums)
	result := 0
	pre := nums[0]
	for i:=1;i<n;i++{
		result = max(result,nums[i]-pre)
		pre = nums[i]
	}
	return result

}

func max(x,y int)int{
	if x > y{
		return x
	}
	return y
}

func main(){
	println(maximumGap([]int{3,6,9,1}))
}
