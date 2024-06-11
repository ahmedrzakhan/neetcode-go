package main

import (
	"fmt"
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/description/

// TC - O(N), SC - O(N)
func isPalindrome(S string) bool {
	var cleanedStr strings.Builder

	for _, ch := range S {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			cleanedStr.WriteRune(unicode.ToLower(ch))
		}
	}

	cleanedString := cleanedStr.String()
	lenOfCleanStr := len(cleanedString)

	for i := 0; i < lenOfCleanStr/2; i++ {
		if cleanedString[i] != cleanedString[lenOfCleanStr-1-i] {
			return false
		}
	}

	return true
}

func mainisPalindrome() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
