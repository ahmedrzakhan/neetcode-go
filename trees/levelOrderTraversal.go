package main

import (
	"fmt"
)

// https://leetcode.com/problems/binary-tree-level-order-traversal/

type TreeNodeLOT struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(N)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int

		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, currentNode.Val)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

func mainLOT() {
	// Example tree
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}

	fmt.Println("Level Order Traversal:", levelOrder(root))
}
