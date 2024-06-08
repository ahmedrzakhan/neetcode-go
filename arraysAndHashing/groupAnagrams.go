package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/group-anagrams/

// TC - O(N*K), SC - O(N), K is length of string
func groupAnagrams(A []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, str := range A {
		var count [26]int

		for _, ch := range str {
			lCh := strings.ToLower(string(ch))[0]
			count[lCh-'a']++
		}

		anagramMap[count] = append(anagramMap[count], str)
	}

	result := make([][]string, 0)

	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func mainGroupAnag() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
}
