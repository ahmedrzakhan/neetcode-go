package main

import "fmt"

type TreeNodeKth struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/kth-smallest-element-in-a-bst

// TC  - O(N), SC - O(H)
func kthSmallest(root *TreeNode, K int) int {
	count := 0
	return inOrderTraversal(root, &count, K)
}

// inOrderTraversal performs the in-order traversal of the BST.
func inOrderTraversal(curr *TreeNode, count *int, K int) int {
	if curr == nil {
		return -1
	}

	if curr.Left != nil {
		val := inOrderTraversal(curr.Left, count, K)
		if *count == K {
			return val
		}
	}

	*count++
	if *count == K {
		return curr.Val
	}

	if curr.Right != nil {
		return inOrderTraversal(curr.Right, count, K)
	}

	return -1
}

// TC - (H+k), SC - O(H)
func kthSmallestIt(root *TreeNode, K int) int {
	stack := []*TreeNode{}
	curr := root
	count := 0

	for curr != nil || len(stack) > 0 {
		// Push all left children onto the stack
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// Process the node
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		// If count is k, return the value
		if count == K {
			return curr.Val
		}

		// Move to the right child
		curr = curr.Right
	}

	// If k is out of bounds
	return -1
}

// helper function to create a new TreeNode.
func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func mainKth() {
	root := newNode(3)
	root.Left = newNode(1)
	root.Right = newNode(4)
	root.Left.Right = newNode(2)

	k := 1
	fmt.Printf("The %dth smallest element in the BST is: %d\n", k, kthSmallest(root, 3))
}
