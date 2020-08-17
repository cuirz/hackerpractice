package main

//99. 恢复二叉搜索树
//二叉搜索树中的两个节点被错误地交换。
//
//请在不改变其结构的情况下，恢复这棵树。
//示例 1:
//
//输入: [1,3,null,null,2]
//
//   1
//  /
// 3
//  \
//   2
//
//输出: [3,1,null,null,2]
//
//   3
//  /
// 1
//  \
//   2
//示例 2:
//输入: [3,1,4,null,null,2]
//
//3
/// \
//1   4
//   /
//  2
//
//输出: [2,1,4,null,null,3]
//
//2
/// \
//1   4
//   /
// 3
//进阶:
//
//使用 O(n) 空间复杂度的解法很容易实现。
//你能想出一个只使用常数空间的解决方案吗？

//思路 中序遍历
// 假如正常是[1,2,3,4,5,6,7]
// 遍历出结果[1,6,3,4,5,2,7]
// 6>3,5>2  调换 6和2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	var pre,x,y *TreeNode
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && pre.Val > root.Val{
			y = root
			if x == nil{
				x = pre
			}else{
				break
			}

		}

		pre = root
		root = root.Right
	}
	x.Val,y.Val = y.Val,x.Val
}
