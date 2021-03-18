package main

//92. 反转链表 II
//给你单链表的头节点 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
// 
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//示例 2：
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
// 
//
//提示：
//
//链表中节点数目为 n
//1 <= n <= 500
//-500 <= Node.val <= 500
//1 <= left <= right <= n
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	root := &ListNode{Next: head}
	pre := root
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	pos := pre.Next
	for i := 0; i < right-left; i++ {
		next := pos.Next
		pos.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return root.Next
}
