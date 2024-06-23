package main

import "fmt"

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

type TreeNodeCB struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(N)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// The first element in preorder is always the root.
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// Find the root in inorder to split into left and right subtrees.
	midIndex := findIndex(rootVal, inorder)

	// Recursively build the left and right subtrees.
	root.Left = buildTree(preorder[1:midIndex+1], inorder[:midIndex])
	root.Right = buildTree(preorder[midIndex+1:], inorder[midIndex+1:])
	return root
}

// find returns the index of the given value in the slice.
func findIndex(val int, A []int) int {
	for i, v := range A {
		if v == val {
			return i
		}
	}
	return -1
}

func mainCB() {
	// preorder := []int{3, 9, 20, 15, 7}
	// inorder := []int{9, 3, 15, 20, 7}

	preorder := []int{3, 9, 1, 4, 2, 20, 15, 7, 5, 8}
	inorder := []int{1, 9, 2, 4, 3, 15, 20, 5, 7, 8}
	root := buildTree(preorder, inorder)
	fmt.Println("Root of the constructed tree:", root.Val)
	// Add more print statements if you want to check the structure of the tree

	/**
		3
		/ \
	   9   20
	  / \   / \
	 1   4 15  7
		/      / \
	   2      5   8
	*/
}
