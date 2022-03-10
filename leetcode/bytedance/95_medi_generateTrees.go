package main

//95. 不同的二叉搜索树 II
//给定一个整数 n，生成所有由 1 ...n 为节点所组成的 二叉搜索树 。
//
//
//
//示例：
//
//输入：3
//输出：
//[
// [1,null,3,2],
// [3,2,null,1],
// [3,1,null,null,2],
// [2,1,3],
// [1,null,2,null,3]
//]
//解释：
//以上的输出对应以下 5 种不同结构的二叉搜索树：
//
//1         3     3      2      1
//\       /     /      / \      \
//3     2     1      1   3      2
///     /       \                 \
//2     1         2                 3
//
//
//提示：
//
//0 <= n <= 8

//思路 枚举

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var create func(s, e int) []*TreeNode
	create = func(s, e int) []*TreeNode {
		if s > e {
			return []*TreeNode{nil}
		}
		treeArray := make([]*TreeNode, 0)
		for i := s; i <= e; i++ {
			leftTrees := create(s, i-1)
			rightTrees := create(i+1, e)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					root := &TreeNode{Val: i}
					root.Left = left
					root.Right = right
					treeArray = append(treeArray, root)
				}
			}
		}
		return treeArray
	}

	return create(1, n)

}

func main() {
	println(generateTrees(3))
}
