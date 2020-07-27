package main

//141. 环形链表

//思路 快慢指针
// 当有环时，快慢指针会碰面

//思路 哈希表
// 碰到已存储的哈希表里的节点时说明有环
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow, quick := head, head
	t := 0
	for quick != nil {
		t++
		quick = quick.Next
		if t%2 == 0 {
			slow = slow.Next
		}
		if quick == slow {
			return true
		}
	}
	return false
}
