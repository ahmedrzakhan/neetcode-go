package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// TC - O(N), SC - O(N)
func lengthOfLongestSubstring(S string) int {
	charMap := make(map[rune]int)
	maxLen, L := 0, 0

	for R, ch := range S {
		if prevIndex, ok := charMap[ch]; ok && prevIndex >= L {
			L = prevIndex + 1
		}
		charMap[ch] = R
		maxLen = max(maxLen, R-L+1)
	}

	return maxLen
}

func mainlengthOfLongestSubstring() {
	s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}
