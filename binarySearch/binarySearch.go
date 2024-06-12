package main

import "fmt"

// https://leetcode.com/problems/binary-search/

// TC - O(N), SC - O(1)
func search(A []int, target int) int {
	N := len(A)
	L, R := 0, N-1

	for L <= R {
		mid := L + (R-L)/2
		if A[mid] == target {
			return mid
		} else if A[mid] < target {
			L = mid + 1
		} else {
			R = mid - 1
		}
	}

	return -1
}

func mainBinarySearch() {
	A := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(A, target))
}
