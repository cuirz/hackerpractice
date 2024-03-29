package main

//94. 二叉树的中序遍历
//给定一个二叉树，返回它的中序遍历。
//
//示例:
//
//输入: [1,null,2,3]
//1
//\
//2
///
//3
//
//输出: [1,3,2]
//进阶:递归算法很简单，你可以通过迭代算法完成吗？

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	var result []int

	left := inorderTraversal(root.Left)
	if left != nil {
		result = append(result, left...)
	}
	result = append(result, root.Val)
	right := inorderTraversal(root.Right)
	if right != nil {
		result = append(result, right...)
	}
	return result
}

func main() {
	var result []int
	result = append(result, []int{1}...)

}
