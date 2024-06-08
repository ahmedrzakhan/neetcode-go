package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/top-k-frequent-elements/

type Pair struct {
	num       int
	frequency int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].frequency > h[j].frequency
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	N := len(old)
	X := old[N-1]
	*h = old[0 : N-1]
	return X
}

// TC - O(N), SC - O(N)
func topKFrequent(A []int, K int) []int {
	freqMap := make(map[int]int)

	for _, ele := range A {
		freqMap[ele]++
	}

	h := &PairHeap{}
	heap.Init(h)
	for num, freq := range freqMap {
		heap.Push(h, Pair{num, freq})
	}

	result := []int{}
	for i := 0; i < K; i++ {
		top := heap.Pop(h).(Pair)
		result = append(result, top.num)
	}
	return result
}

func mainTopKFreq() {
	nums := []int{-1, -1}
	k := 1
	fmt.Println("Top K Frequent Elements:", topKFrequent(nums, k)) // Output: [1 2]
}
