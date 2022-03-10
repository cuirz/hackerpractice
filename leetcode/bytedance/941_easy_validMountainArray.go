package main

//941. 有效的山脉数组
//给定一个整数数组A，如果它是有效的山脉数组就返回true，否则返回 false。
//
//让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
//
//A.length >= 3
//在0 < i< A.length - 1条件下，存在i使得：
//A[0] < A[1] < ... A[i-1] < A[i]
//A[i] > A[i+1] > ... > A[A.length - 1]

func validMountainArray(A []int) bool {
	peek := 0
	pre := -1
	strictly := 0
	for i, v := range A {
		if strictly == 0 {
			if v == pre {
				return false
			} else if v > pre {
				pre = v
			} else {
				peek = i - 1
				strictly = 1
				pre = v
			}
		} else {
			if v < pre {
				pre = v
			} else {
				return false
			}
		}
	}
	return peek != 0 && peek != len(A)-1
}
