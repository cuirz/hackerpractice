package bst

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return check(root, sum, false)
}

func check(root *TreeNode, sum int, isLeaf bool) bool {
	if root == nil {
		if sum != 0 || !isLeaf {
			return false
		}
		return true
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		isLeaf = true
	} else {
		isLeaf = false
	}
	return check(root.Left, sum, isLeaf) || check(root.Right, sum, isLeaf)
}

func hasPathSum2(root *TreeNode, sum int) bool {
	// 1.如果当前跟节点为空则返回False
	if root == nil {
		return false
	}
	// 2.进行减值 如果已经小于0则直接返回false 不进行继续探测 剪枝操作
	// 3.如果sum已经为0 且有一个子树为空直接返回true
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	// 递归
	sum -= root.Val
	return hasPathSum2(root.Left, sum) || hasPathSum2(root.Right, sum)
}
