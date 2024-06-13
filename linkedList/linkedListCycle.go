package main

import "fmt"

// https://leetcode.com/problems/linked-list-cycle/

type ListNodeC struct {
	Val  int
	Next *ListNode
}

// TC - O(N), SC - O(1)
func hasCycle(head *ListNode) bool {

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}

func mainLLC() {
	fmt.Println("main")
}
