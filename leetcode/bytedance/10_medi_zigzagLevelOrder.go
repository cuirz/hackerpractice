package main

//103. 二叉树的锯齿形层序遍历
//给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回锯齿形层序遍历如下：
//
//[
//[3],
//[20,9],
//[15,7]
//]

//思路  层序遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var search func(node *TreeNode, d int)
	search = func(node *TreeNode, d int) {
		if node == nil{
			return
		}
		if len(result) <= d {
			result = append(result, []int{})
		}
		result[d] = append(result[d], node.Val)
		search(node.Left, d+1)
		search(node.Right, d+1)
	}
	search(root, 0)
	for i := 1; i < len(result); i += 2 {
		end := len(result[i])
		mid := end >> 1
		for j := 0; j < mid; j++ {
			result[i][j], result[i][end-j-1] = result[i][end-j-1], result[i][j]
		}
	}
	return result
}
