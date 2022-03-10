package main

//373. 查找和最小的K对数字
//给定两个以升序排列的整数数组 nums1 和 nums2,以及一个整数 k。
//
//定义一对值(u,v)，其中第一个元素来自nums1，第二个元素来自 nums2。
//
//请找到和最小的 k个数对(u1,v1), (u2,v2) ... (uk,vk)。
//
//
//
//示例 1:
//
//输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//输出: [1,2],[1,4],[1,6]
//解释: 返回序列中的前 3 对数：
//[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//示例 2:
//
//输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//输出: [1,1],[1,1]
//解释: 返回序列中的前 2 对数：
//    [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//示例 3:
//
//输入: nums1 = [1,2], nums2 = [3], k = 3
//输出: [1,3],[2,3]
//解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//
//
//提示:
//
//1 <= nums1.length, nums2.length <= 10^4
//-10^9 <= nums1[i], nums2[i] <= 10^9
//nums1, nums2 均为升序排列
//1 <= k <= 1000

//思路 二分查找
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res [][]int
	m := len(nums1)
	n := len(nums2)
	l, r := nums1[0]+nums2[0], nums1[m-1]+nums2[n-1]
	sum := r
	for l <= r {
		mid := l + (r-l)>>1
		count, start, end := 0, 0, n-1
		for start < m && end >= 0 {
			if nums1[start]+nums2[end] > mid {
				end--
			} else {
				count += end + 1
				start++
			}
		}
		if count < k {
			l = mid + 1
		} else {
			sum = mid
			r = mid - 1
		}
	}
	pos := n - 1
	for _, first := range nums1 {
		for pos >= 0 && first+nums2[pos] >= sum {
			pos--
		}
		for _, second := range nums2[:pos+1] {
			res = append(res, []int{first, second})
			if len(res) == k {
				return res
			}
		}
	}
	pos = n - 1
	for _, first := range nums1 {
		for pos >= 0 && first+nums2[pos] > sum {
			pos--
		}
		for j := pos; j >= 0 && first+nums2[j] == sum; j-- {
			res = append(res, []int{first, nums2[j]})
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
