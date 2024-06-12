package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/koko-eating-bananas/

func check(A []int, h, K int) bool {
	var result int

	for _, v := range A {
		result += (v + K - 1) / K
	}

	return result <= h
}

// TC - O(NLogN), SC - O(N)
func minEatingSpeed(A []int, h int) int {
	L, R, res := 1, maxArr(A), 1

	for L <= R {
		mid := L + (R-L)/2
		if check(A, h, mid) {
			res = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return res
}

func maxArr(A []int) int {
	max := math.MinInt
	for _, v := range A {
		if v > max {
			max = v
		}
	}
	return max
}

func mainKoko() {
	piles := []int{312884470}
	h := 312884469
	fmt.Println(minEatingSpeed(piles, h))
}
