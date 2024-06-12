package main

import "fmt"

// https://leetcode.com/problems/daily-temperatures/

// TC -O(N), SC - O(N)
func dailyTemperatures(A []int) []int {
	N := len(A)
	result := make([]int, N)
	var stack []int

	for i := N - 1; i >= 0; i-- {
		for len(stack) > 0 && A[i] >= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}

	return result
}

func maindailyTemperatures() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}

	fmt.Println(dailyTemperatures(temperatures))
}
