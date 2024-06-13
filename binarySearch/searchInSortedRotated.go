package main

import "fmt"

// https://leetcode.com/problems/search-in-rotated-sorted-array/

// TC - O(N), SC - O(1)
func searchR(A []int, target int) int {
	L, R := 0, len(A)-1
	for L <= R {
		mid := L + (R-L)/2
		if A[mid] == target {
			return mid
		}

		if A[mid] >= A[L] {
			if target >= A[L] && target < A[mid] {
				R = mid - 1
			} else {
				L = mid + 1
			}
		} else {
			if target > A[mid] && target <= A[R] {
				L = mid + 1
			} else {
				R = mid - 1
			}
		}
	}

	return -1
}

func mainsearchR() {
	arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	target := 5
	result := searchR(arr, target)
	fmt.Println("Index of target is:", result) // Output: Index of target is: 4

	arr = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	result = searchR(arr, target)
	fmt.Println("Index of target is:", result) // Output: Index of target is: 4

	// target = 3
	// result = searchR(arr, target)
	// fmt.Println("Index of target is:", result) // Output: Index of target is: -1

	arr = []int{6, 7, 0, 1, 2, 4, 5}
	target = 5
	result = searchR(arr, target)
	fmt.Println("Index of target is:", result) // Output: Index of target is: -1

}
