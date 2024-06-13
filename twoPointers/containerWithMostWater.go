package main

import "fmt"

// https://leetcode.com/problems/container-with-most-water/

// TC - O(N), SC - O(N)
func maxArea(A []int) int {
	N := len(A)
	L, R := 0, N-1
	maxArea := 0

	for L < R {
		area := (R - L) * min(A[L], A[R])
		maxArea = max(area, maxArea)

		if A[L] < A[R] {
			L++
		} else {
			R--
		}
	}
	return maxArea
}

func mainmaxArea() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(arr))
}
