package main

//143. 重排链表
//给定一个单链表L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//示例1:
//
//给定链表 1->2->3->4, 重新排列为 1->4->2->3.
//示例 2:
//
//给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid, tail := head, head
	for tail != nil && tail.Next != nil {
		mid = mid.Next
		tail = tail.Next.Next
	}
	l2 := mid.Next
	mid.Next = nil
	var pre, cur *ListNode = nil, l2
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	l1 := head
	l2 = pre
	var t1, t2 *ListNode
	for l1 != nil && l2 != nil {
		t1 = l1.Next
		t2 = l2.Next
		l1.Next = l2
		l1 = t1
		l2.Next = l1
		l2 = t2
	}
}
