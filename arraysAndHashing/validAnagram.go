package main

import "fmt"

// https://leetcode.com/problems/valid-anagram/description/

// TC - O(N), SC - O(N)
func isAnagram(S, T string) bool {

	if len(S) != len(T) {
		return false
	}

	charMap := map[string]int{}

	for i := range S {
		charMap[string(S[i])]++
		charMap[string(T[i])]--
	}

	for i := range S {
		if charMap[string(T[i])] != 0 {
			return false
		}
	}

	return true
}

func mainAnagram() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
