package main

//919. 完全二叉树插入器
//完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。
//
//设计一种算法，将一个新节点插入到一个完整的二叉树中，并在插入后保持其完整。
//
//实现 CBTInserter 类:
//
//CBTInserter(TreeNode root)使用头节点为root的给定树初始化该数据结构；
//CBTInserter.insert(int v) 向树中插入一个值为Node.val == val的新节点TreeNode。使树保持完全二叉树的状态，并返回插入节点TreeNode的父节点的值；
//CBTInserter.get_root() 将返回树的头节点。
//
//
//示例 1：
//
//
//
//输入
//["CBTInserter", "insert", "insert", "get_root"]
//[[[1, 2]], [3], [4], []]
//输出
//[null, 1, 2, [1, 2, 3, 4]]
//
//解释
//CBTInserter cBTInserter = new CBTInserter([1, 2]);
//cBTInserter.insert(3);  // 返回 1
//cBTInserter.insert(4);  // 返回 2
//cBTInserter.get_root(); // 返回 [1, 2, 3, 4]
//
//
//提示：
//
//树中节点数量范围为[1, 1000]
//0 <= Node.val <= 5000
//root是完全二叉树
//0 <= val <= 5000
//每个测试用例最多调用insert和get_root操作10^4次

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	*TreeNode
	candidate []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	candidate := []*TreeNode{}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
		if top.Left == nil || top.Right == nil {
			candidate = append(candidate, top)
		}

	}
	return CBTInserter{root, candidate}
}

func (this *CBTInserter) Insert(val int) int {
	child := &TreeNode{Val: val}
	top := this.candidate[0]
	if top.Left == nil {
		top.Left = child
	} else {
		top.Right = child
		this.candidate = this.candidate[1:]
	}
	this.candidate = append(this.candidate, child)
	return top.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.TreeNode
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
