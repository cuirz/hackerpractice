package main

//515. 在每个树行中找最大值
//给定一棵二叉树的根节点root ，请找出该二叉树中每一层的最大值。
//
//
//
//示例1：
//
//
//
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
//示例2：
//
//输入: root = [1,2,3]
//输出: [1,3]
//
//
//提示：
//
//二叉树的节点个数的范围是 [0,10^4]
//-2^31<= Node.val <= 2^31- 1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var result []int
	var search func(*TreeNode, int)
	search = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		if dep >= len(result) {
			result = append(result, node.Val)
		} else {
			result[dep] = max(result[dep], node.Val)
		}
		search(node.Left, dep+1)
		search(node.Right, dep+1)
	}
	search(root, 0)
	return result
}
