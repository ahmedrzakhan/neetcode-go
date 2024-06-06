package main

import "fmt"

// https://leetcode.com/problems/contains-duplicate/

// TC - O(N), SC - O(N)
func containsDuplicate(A []int) bool {
	checkMap := map[int]int{}

	for i := range A {
		if _, ok := checkMap[A[i]]; ok {
			return true
		} else {
			checkMap[A[i]] = 1
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(arr))
}
