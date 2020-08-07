package main

//142. 环形链表 II
//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//说明：不允许修改给定的链表。

//思路 快慢指针
// 快慢指针第一次交汇的地方 假设 慢指针走了 k步，那快指针走了2k步
// 慢指针从交汇处开始继续走，创建一个新慢指针从头部开始同步走
// 当两慢指针交汇时就是环的节点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	peek, slow, fast := head, head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != peek {
				slow = slow.Next
				peek = peek.Next
			}
			return peek
		}
	}
	return nil
}
