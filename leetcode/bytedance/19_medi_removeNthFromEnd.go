package main

//19. 删除链表的倒数第N个节点
//给定一个链表，删除链表的倒数第n个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//说明：
//
//给定的 n保证是有效的。
//
//进阶：
//
//你能尝试使用一趟扫描实现吗？

//思路 快慢指针

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	index := 0
	first := head
	second := first
	var pre *ListNode
	for first != nil {
		first = first.Next
		if index < n{
			index ++
		}else{
			pre = second
			second = second.Next
		}
	}
	if pre != nil{
		pre.Next = second.Next
	}else{
		head = second.Next
	}

	return head
}
