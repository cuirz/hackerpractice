package main

//337. 打家劫舍 III
//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
//计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
//示例 1:
//
//输入: [3,2,3,null,3,null,1]
//
//3
/// \
//2   3
//\   \
//3   1
//
//输出: 7
//解释:小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//示例 2:
//
//输入: [3,4,5,1,3,null,1]
//
//    3
/// \
//4   5
/// \   \
//1   3   1
//
//输出: 9
//解释:小偷一晚能够盗取的最高金额= 4 + 5 = 9.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// dfs
// 当前root 可以偷 和 不可以偷时
// 可以时 left不可以时的最大 + right不可以时的最大 + root
// 不可以时 max(left可以，left不可以) + max(right可以，right不可以）

func rob(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// 返回 []int 0.可以时   1.不可以时
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		if root.Left == nil && root.Right == nil {
			return []int{root.Val, 0}
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		// 后序
		return []int{root.Val + left[1] + right[1],
			max(left[0], left[1]) + max(right[0], right[1])}
	}
	result := dfs(root)
	return max(result[0], result[1])

}
