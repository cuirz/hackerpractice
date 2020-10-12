package main

import "math"

//968. 监控二叉树
//给定一个二叉树，我们在树的节点上安装摄像头。
//
//节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
//
//计算监控树的所有节点所需的最小摄像头数量。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	inf := math.MaxInt32  /2
	var minCamera func(node *TreeNode) (int,int,int)
	minCamera = func(node *TreeNode) (int,int,int) {
		if node == nil {
			return inf,0,0
		}
		la,lb,lc := minCamera(node.Left)
		ra,rb,rc := minCamera(node.Right)
		a := lc+rc+1
		b := min(a,min(la+rb,ra+lb))
		c := min(a,lb+rb)
		return a,b,c

	}
	_,res,_ := minCamera(root)
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
