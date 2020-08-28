package main

//110. 平衡二叉树

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var deep func(node *TreeNode) int
	deep = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := deep(node.Left)
		right := deep(node.Right)
		if abs(left-right) > 1 || left == -1 || right == -1 {
			return -1
		}
		return max(left, right) + 1
	}
	return deep(root) >= 0
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	isBalanced(nil)
}
