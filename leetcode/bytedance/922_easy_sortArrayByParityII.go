package main

//922. 按奇偶排序数组 II
//给定一个非负整数数组A， A 中一半整数是奇数，一半整数是偶数。
//
//对数组进行排序，以便当A[i] 为奇数时，i也是奇数；当A[i]为偶数时， i 也是偶数。
//
//你可以返回任何满足上述条件的数组作为答案。
//示例：
//
//输入：[4,2,5,7]
//输出：[4,5,2,7]
//解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
//
//
//提示：
//
//2 <= A.length <= 20000
//A.length % 2 == 0
//0 <= A[i] <= 1000

func sortArrayByParityII(A []int) []int {

	n := len(A)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if A[i]%2 != 0 {
				for j := i + 1; j < n; j += 2 {
					if A[j]%2 == 0 {
						A[i], A[j] = A[j], A[i]
						break
					}
				}
			}
		} else {
			if A[i]%2 == 0 {
				for j := i + 1; j < n; j += 2 {
					if A[j]%2 != 0 {
						A[i], A[j] = A[j], A[i]
						break
					}
				}
			}
		}

	}

	for i, j := 0, 1; i < n; i += 2 {
		if A[i]%2 != 0 {
			for A[j]%2 != 0 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}

	}
	return A
}
