package main

import "fmt"

// https://leetcode.com/problems/search-a-2d-matrix/d

// TC - O(Log(R*C)), SC - O(1)
func searchMatrix(M [][]int, target int) bool {
	ROWS := len(M)
	COLS := len(M[0])

	TOP := 0
	BOT := ROWS - 1

	for TOP <= BOT {
		ROW := (TOP + BOT) / 2
		if target > M[ROW][len(M[ROW])-1] {
			TOP = ROW + 1
		} else if target < M[ROW][0] {
			BOT = ROW - 1
		} else {
			break
		}
	}

	if !(TOP <= BOT) {
		return false
	}

	ROW := (TOP + BOT) / 2
	L, R := 0, COLS-1

	for L <= R {
		mid := (L + R) / 2
		if target == M[ROW][mid] {
			return true
		} else if target > M[ROW][mid] {
			L = mid + 1
		} else if target < M[ROW][mid] {
			R = mid - 1
		}
	}

	return false
}

func mainsearchMatrix() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 23

	fmt.Println(searchMatrix(matrix, target))
}
