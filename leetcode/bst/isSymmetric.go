package bst

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := make([]*TreeNode, 0)
	right := make([]*TreeNode, 0)

	left = append(left, root.Left)
	right = append(right, root.Right)
	for compare(left, right) {
		l := make([]*TreeNode, 0)
		r := make([]*TreeNode, 0)
		for _, v := range left {
			if v != nil {
				l = append(l, v.Left)
				l = append(l, v.Right)
			}
		}
		for _, v := range right {
			if v != nil {
				r = append(r, v.Left)
				r = append(r, v.Right)
			}
		}
		left = l
		right = r
		if len(left) == 0 && len(right) == 0 {
			return true
		}
	}
	return false
}

func compare(left, right []*TreeNode) bool {
	if len(left) != len(right) {
		return false
	}
	size := len(right)
	for i, v := range left {
		index := size - 1 - i
		if v == nil {
			if right[index] != nil {
				return false
			}
		} else if right[index] == nil {
			return false
		} else if v.Val != right[index].Val {
			return false
		}
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	return isSym(root, root)
}

func isSym(r1, r2 *TreeNode) bool {
	if r1 == nil || r2 == nil {
		return r1 == r2
	}
	if r1.Val != r2.Val {
		return false
	}
	return isSym(r1.Left, r2.Right) && isSym(r1.Right, r2.Left)
}
