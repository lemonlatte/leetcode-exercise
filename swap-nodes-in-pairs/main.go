package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(prev, head *ListNode) (*ListNode, *ListNode) {
	second := head.Next

	if second == nil {
		return nil, nil
	}

	third := second.Next

	head.Next = third
	second.Next = head

	if prev != nil {
		prev.Next = second
	}
	return head, third
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head.Next
	var prev *ListNode

	for head != nil {
		prev, head = swap(prev, head)
	}

	return first
}
