package main

import "fmt"

// https://leetcode.com/problems/generate-parentheses/

// TC - O(2^N) SC - O(N)
func generateParantheses(N int) []string {
	var subsets []string
	helperGen(0, 0, N, "", &subsets)
	return subsets
}

func helperGen(open, close, N int, curSet string, subsets *[]string) {
	if len(curSet) == 2*N {
		*subsets = append(*subsets, curSet)
		return
	}

	if open < N {
		helperGen(open+1, close, N, curSet+"(", subsets)
	}
	if close < open {
		helperGen(open+1, close, N, curSet+")", subsets)
	}
}

func maiGen() {
	n := 3
	fmt.Println(generateParantheses(n))
}
