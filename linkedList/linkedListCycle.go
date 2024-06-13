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

func createListss(vals []int, cycleIndex int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy
	var cycleNode *ListNode

	for i, val := range vals {
		current.Next = &ListNode{Val: val}
		current = current.Next
		if i == cycleIndex {
			cycleNode = current
		}
	}

	if cycleNode != nil {
		current.Next = cycleNode // Create a cycle
	}

	return dummy.Next
}

func mainLLC() {
	// Example: Create a list with a cycle
	// List: 1 -> 2 -> 3 -> 4 -> 2 (cycle starts at index 1)
	head := createListss([]int{1, 2, 3, 4}, 1)
	fmt.Println("Does the list have a cycle?", hasCycle(head))
}
