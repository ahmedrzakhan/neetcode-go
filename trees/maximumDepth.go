package main

import "fmt"

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(H)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	// Example binary tree
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	maxDepth := maxDepth(root)
	fmt.Println("Maximum depth:", maxDepth) // Output: Maximum depth: 3
}
