package main

import (
	"math"
	"math/rand"
)

//363. 矩形区域不超过 K 的最大数值和
//给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
//
//题目数据保证总会存在一个数值和不超过 k 的矩形区域。
//输入：matrix = [[1,0,1],[0,-2,3]], k = 2
//输出：2
//解释：蓝色边框圈出来的矩形区域[[0, 1], [-2, 3]]的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
//输入：matrix = [[2,2,-1]], k = 3
//输出：3
//思路 有序集合 二分查找

func maxSumSubmatrix(matrix [][]int, k int) int {
	result := math.MinInt64
	for i := range matrix {
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] {
			for c, v := range row {
				sum[c] += v
			}
			t := &TreeSet{}
			t.put(0)
			s := 0
			for _, v := range sum {
				s += v
				if lb := t.lowerbound(s - k); lb != nil {
					result = max(result, s-lb.val)
				}
				t.put(s)

			}
		}
	}
	return result
}

type node struct {
	ch       [2]*node
	val      int
	priority int
}

func (n *node) cmp(b int) int {
	if n.val > b {
		return 0
	} else if n.val < b {
		return 1
	}
	return -1
}
func (n *node) rotate(d int) *node {
	x := n.ch[d^1]
	n.ch[d^1] = x.ch[d]
	x.ch[d] = n
	return x
}

type TreeSet struct {
	root *node
}

func (t *TreeSet) _put(n *node, val int) *node {
	if n == nil {
		return &node{
			val:      val,
			priority: rand.Int(),
		}
	}
	if d := n.cmp(val); d >= 0 {
		n.ch[d] = t._put(n.ch[d], val)
		if n.ch[d].priority > n.priority {
			n = n.rotate(d ^ 1)
		}
	}
	return n
}
func (t *TreeSet) put(val int) {
	t.root = t._put(t.root, val)

}
func (t *TreeSet) lowerbound(val int) (lb *node) {
	for n := t.root; n != nil; {
		d := n.cmp(val)
		if d == 0 {
			lb = n
			n = n.ch[0]
		} else if d > 0 {
			n = n.ch[1]
		} else {
			return n
		}
	}
	return
}
