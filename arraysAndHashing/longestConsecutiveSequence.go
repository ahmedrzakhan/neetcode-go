package main

import "fmt"

// TC - O(N), SC - O(N)
func longestConsecutive(A []int) int {
	numsMap := make(map[int]bool)

	for _, num := range A {
		numsMap[num] = true
	}

	maxLen := 0
	for num := range numsMap {
		if !numsMap[num-1] {
			length := 1
			for numsMap[num+length] {
				length++
			}
			maxLen = max(length, maxLen)
		}
	}
	return maxLen
}

func mainlongestConsecutive() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
