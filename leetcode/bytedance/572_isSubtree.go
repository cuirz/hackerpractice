package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s.Val == t.Val {
		return isSubtree(s.Left, t.Left) && isSubtree(s.Right, t.Right)
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func compare(s *TreeNode, t *TreeNode) bool {
	if s == t {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		if compare(s.Left, t.Left, true) && compare(s.Right, t.Right, true) {
			return true
		}
		return false
	}
	return false
}

func main() {

}
