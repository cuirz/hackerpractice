package main

import (
	"math"
	"math/rand"
)

//384. 打乱数组
//给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。
//
//实现 Solution class:
//
//Solution(int[] nums) 使用整数数组 nums 初始化对象
//int[] reset() 重设数组到它的初始状态并返回
//int[] shuffle() 返回数组随机打乱后的结果
// 
//
//示例：
//
//输入
//["Solution", "shuffle", "reset", "shuffle"]
//[[[1, 2, 3]], [], [], []]
//输出
//[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
//
//解释
//Solution solution = new Solution([1, 2, 3]);
//solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3, 1, 2]
//solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
//solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]
// 
//
//提示：
//
//1 <= nums.length <= 200
//-10^6 <= nums[i] <= 10^6
//nums 中的所有元素都是 唯一的
//最多可以调用 5 * 10^4 次 reset 和 shuffle

type Solution struct {
	array,origin []int
}


func Constructor(nums []int) Solution {
	return Solution{
		array: append([]int{},nums...),
		origin: nums,
	}
}


func (this *Solution) Reset() []int {
	copy(this.array,this.origin)
	return this.array
}


func (this *Solution) Shuffle() []int {
	n:=len(this.array)
	for i,_ := range this.array{
		j := i+rand.Intn(n-i)
		this.array[i],this.array[j] = this.array[j],this.array[i]
	}
	return this.array
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
