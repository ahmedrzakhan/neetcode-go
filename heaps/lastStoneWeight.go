package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/last-stone-weight/

// TC  - O(NLogN), SC - O(N)
type IntHeapLSW []int

func (h *IntHeap) LenLSW() int {
	return len(*h)
}

func (h *IntHeap) LessLSW(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) SwapLSW(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) PushLSW(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) PopLSW() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(A []int) int {
	h := IntHeap(A)

	heap.Init(&h)

	for h.Len() > 1 {
		first := heap.Pop(&h).(int)
		second := heap.Pop(&h).(int)

		diff := first - second

		if diff > 0 {
			heap.Push(&h, diff)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return heap.Pop(&h).(int)
}

func mainLSW() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Printf("The weight of the last stone is: %d\n", lastStoneWeight(stones))
}
