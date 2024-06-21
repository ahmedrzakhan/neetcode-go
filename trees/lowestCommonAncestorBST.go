package main

import "fmt"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNodeLC struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(H), SC - O(H)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Base case: Empty tree
	if root == nil {
		return nil
	}

	// Exploit BST property:
	// * If both p and q are smaller than root, LCA is in the left subtree
	// * If both p and q are larger than root, LCA is in the right subtree
	// * Otherwise, the current root is the LCA
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

func mainLC() {
	// Example BST
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}

	p := root.Left
	q := root.Right

	lca := lowestCommonAncestor(root, p, q)
	fmt.Println("LCA of", p.Val, "and", q.Val, "is:", lca.Val)
}
