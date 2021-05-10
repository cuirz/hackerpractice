package main

//872. 叶子相似的树
//请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列
//举个例子，如上图所示，给定一棵叶值序列为 (6, 7, 4, 9, 8) 的树。
//
//如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
//
//如果给定的两个根结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。
//输入：root1 = [1,2,3], root2 = [1,3,2]
//输出：false
// 
//
//提示：
//
//给定的两棵树可能会有 1 到 200 个结点。
//给定的两棵树上的值介于 0 到 200 之间。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	stack := make([]int, 0)
	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			stack = append(stack, node.Val)
			return
		}
		search(node.Left)
		search(node.Right)
	}

	search(root1)
	array1 := append([]int{}, stack...)
	stack = make([]int,0)
	search(root2)
	array2 := append([]int{}, stack...)
	if len(array1) != len(array2) {
		return false
	}
	for i, v := range array1 {
		if v != array2[i] {
			return false
		}
	}
	return true
}

func main() {
	leafSimilar(nil, nil)
}
