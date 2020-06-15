package bst

func maxDepth(root *TreeNode) int {
	return loop(root, 0)
}

func loop(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	left := loop(root.Left, depth+1)
	right := loop(root.Right, depth+1)
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
