package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	recursive(&result, l1, l2)
	return result
}

func recursive(result **ListNode, l1 *ListNode, l2 *ListNode) {
	if l1 == nil && l2 == nil {
		if (*result).Val == 0 {
			*result = nil
		}
		return
	}
	(*result).Next = new(ListNode)
	var n1, n2 *ListNode
	if l1 != nil {
		(*result).Val += l1.Val
		n1 = l1.Next
	}
	if l2 != nil {
		(*result).Val += l2.Val
		n2 = l2.Next
	}
	if (*result).Val > 9 {
		(*result).Next.Val = (*result).Val / 10
		(*result).Val = (*result).Val % 10
	}

	recursive(&(*result).Next, n1, n2)
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val:  9,
		Next: nil,
	}

	p := addTwoNumbers(l1, l2)
	next := p
	for {
		if next == nil {
			break
		}
		println(next.Val)
		next = next.Next
	}
	println("3.10" > "20")
	//println(fmt.Sprintf("%v",addTwoNumbers(l1,l2)))
}
