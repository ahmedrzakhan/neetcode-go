package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/validate-binary-search-tree

type TreeNodeVB struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(N)
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(curr *TreeNode, min, max int) bool {
	if curr == nil {
		return true
	}

	if curr.Val <= min || curr.Val >= max {
		return false
	}

	return isValid(curr.Left, min, curr.Val) && isValid(curr.Right, curr.Val, max)
}

func mainVB() {
	root := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 8}}
	result := isValidBST(root)
	fmt.Println(result) // Should print "true"

	root = &TreeNode{Val: 24,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 33}}},
		Right: &TreeNode{Val: 26}}
	result = isValidBST(root)
	fmt.Println(result) // Should print "true"
}
