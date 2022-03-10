package main

//82. 删除排序链表中的重复元素 II
//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中没有重复出现的数字。
//
//示例1:
//
//输入: 1->2->3->3->4->4->5
//输出: 1->2->5
//示例2:
//
//输入: 1->1->1->2->3
//输出: 2->3

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	root := &ListNode{}
	pre := root
	cur, next := head, head.Next
	for next != nil {
		isSkip := false
		for next != nil && cur.Val == next.Val {
			isSkip = true
			next = next.Next
		}
		if !isSkip {
			pre.Next = cur
			pre = cur
		}
		cur = next
		if next != nil {
			next = next.Next
		}

	}
	pre.Next = cur
	return root.Next
}
