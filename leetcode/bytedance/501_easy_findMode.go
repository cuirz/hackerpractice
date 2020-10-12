package main

//501. 二叉搜索树中的众数
//给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
//
//假定 BST 有如下定义：
//
//结点左子树中所含结点的值小于等于当前结点的值
//结点右子树中所含结点的值大于等于当前结点的值
//左子树和右子树都是二叉搜索树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	var pre, count, max int
	update := func(val int) {
		if pre == val{
			count ++
		}else{
			pre = val
			count = 1
		}
		if count == max{
			result = append(result,val)
		}else if count > max{
			max = count
			result = []int{val}
		}
	}
	var find func(node *TreeNode)
	find = func(node *TreeNode) {
		if node == nil {
			return
		}
		find(node.Left)
		update(node.Val)
		find(node.Right)
	}
	find(root)
	return result
}
