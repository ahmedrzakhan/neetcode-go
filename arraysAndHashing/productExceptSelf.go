package main

import "fmt"

// https://leetcode.com/problems/product-of-array-except-self/

// TC - O(N), SC - O(N)
func productExceptSelf(A []int) []int {
	N := len(A)
	result := make([]int, N)

	prefix := 1
	for i := 0; i < N; i++ {
		result[i] = prefix
		prefix *= A[i]
	}

	postfix := 1
	for i := N - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= A[i]
	}

	return result
}

func mainproductExceptSelf() {
	nums := []int{1, 2, 3, 4}
	productArray := productExceptSelf(nums)
	fmt.Println("productArray", productArray)
}
