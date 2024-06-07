package main

import "fmt"

// https://leetcode.com/problems/two-sum/

// TC - O(N), SC - O(N)
func twoSum(A []int, target int) []int {

	idxMap := make(map[int]int)

	for i, num := range A {
		if idx, exists := idxMap[target-num]; exists {
			return []int{idx, i}
		}
		idxMap[num] = i
	}

	return nil
}

func maintwoSum() {
	arr := []int{2, 7, 11, 15}
	// arr := []int{-3, -3, 11, 15}
	// target := -6
	// arr := []int{3, -3, 11, 15}
	// target := 0
	// arr := []int{-3, 3, 11, 15}
	// target := 0
	target := 9
	fmt.Println(twoSum(arr, target))
}
