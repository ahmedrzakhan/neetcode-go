package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/

// TC - O(N), SC - O(N)
func isValid(S string) bool {
	stack := make([]string, 0)

	for _, ch := range S {
		brack := string(ch)
		if brack == "{" || brack == "(" || brack == "[" {
			stack = append(stack, brack)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				top := stack[len(stack)-1]
				if brack == "}" && top != "{" ||
					brack == ")" && top != "(" ||
					brack == "]" && top != "[" {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

func mainisValid() {
	s := "()"
	fmt.Println(isValid(s))
}
