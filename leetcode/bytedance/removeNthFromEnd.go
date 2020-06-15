package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	next, peek, pre := head, head, head

	i := 1
	for next != nil {
		next = next.Next
		if i > n {
			if peek.Next != nil {
				pre = peek
				peek = peek.Next
			}
		}
		i++
	}

	if pre == peek {
		head = pre.Next
	} else if peek.Next == nil {
		peek = nil
		pre.Next = peek
	} else {
		pre.Next = pre.Next.Next
	}

	return head

}

func main() {
	removeNthFromEnd(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}, 2)

}
