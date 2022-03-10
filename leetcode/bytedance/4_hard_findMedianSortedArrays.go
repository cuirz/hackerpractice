package main

//4. 寻找两个正序数组的中位数
//给定两个大小为 m 和 n 的正序（从小到大）数组nums1 和nums2。
//
//请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为O(log(m + n))。
//
//你可以假设nums1和nums2不会同时为空。
//
//
//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5

//思路
//找到第 k (k>1) 小的元素，那么就取 A = nums1[k/2-1] 和 B = nums2[k/2-1] 进行比较
//nums1 中小于等于 A 的元素有 nums1[0..k/2-2] 共计 k/2-1 个
//nums2 中小于等于 B 的元素有 nums2[0..k/2-2] 共计 k/2-1 个
//取 v = min(A, A)，两个数组中小于等于 v 的元素共计不会超过 (k/2-1) + (k/2-1) <= k-2 个
//这样 v 本身最大也只能是第 k-1 小的元素
//如果 v = A，那么 nums1[0..k/2-1] 都不可能是第 k 小的元素。把这些元素全部 "删除"，剩下的作为新的 nums1 数组
//如果 v = B，那么 nums2[0..k/2-1] 都不可能是第 k 小的元素。把这些元素全部 "删除"，剩下的作为新的 nums2 数组
//由于 "删除" 了一些元素（这些元素都比第 k 小的元素要小），因此需要修改 k 的值，减去删除的数的个数

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	size := n + m
	if size%2 == 1 {
		mid := size>>1 + 1
		return float64(getKthElement(nums1, nums2, mid))
	} else {
		mid1, mid2 := size>>1, size>>1+1
		return float64(getKthElement(nums1, nums2, mid1)+getKthElement(nums1, nums2, mid2)) / 2

	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	offset1, offset2 := 0, 0
	for {
		if offset1 == len(nums1) {
			return nums2[offset2+k-1]
		}
		if offset2 == len(nums2) {
			return nums1[offset1+k-1]
		}
		if k == 1 {
			return min(nums1[offset1], nums2[offset2])
		}
		mid := k >> 1
		index1 := min(offset1+mid-1, len(nums1)-1)
		index2 := min(offset2+mid-1, len(nums2)-1)
		x, y := nums1[index1], nums2[index2]
		if x <= y {
			k -= index1 - offset1 + 1
			offset1 = index1 + 1
		} else {
			k -= index2 - offset2 + 1
			offset2 = index2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
