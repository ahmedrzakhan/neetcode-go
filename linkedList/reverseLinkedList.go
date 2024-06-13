package main

import "fmt"

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	head *ListNode
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

func mainRLL() {
	fmt.Println("main")
}
