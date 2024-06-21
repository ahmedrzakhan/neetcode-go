package main

import (
	"fmt"
)

// https://leetcode.com/problems/binary-tree-right-side-view

type TreeNodeRS struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(N)
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			// If it's the last node in the current level, add it to the result.
			if i == levelSize-1 {
				result = append(result, currentNode.Val)
			}

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}

	return result
}

func mainBR() {
	// Example tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}}}
	root.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}

	fmt.Println("Right Side View:", rightSideView(root))
}
