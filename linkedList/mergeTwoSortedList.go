package main

import "fmt"

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNodeM struct {
	Val  int
	Next *ListNode
}

// TC - O(N), SC - O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else if l2 != nil {
		curr.Next = l2
	}

	return head.Next
}

func printListM(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	mergedList := mergeTwoLists(l1, l2)
	printList(mergedList) // Output should be 1 1 2 3 4 4
}
