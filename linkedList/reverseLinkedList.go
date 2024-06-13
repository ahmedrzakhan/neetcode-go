package main

import "fmt"

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// TC - O(N), SC - O(1)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func mainRLL() {
	// Example: 1 -> 2 -> 3 -> 4 -> 5 becomes 5 -> 4 -> 3 -> 2 -> 1
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println("Original List:")
	printList(head)

	head = reverseList(head)

	fmt.Println("Reversed List:")
	printList(head)
}
