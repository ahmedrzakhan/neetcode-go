package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

type IntHeap []int

type Kthlargest struct {
	k    int
	heap *IntHeap
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	N := len(old)
	x := old[N-1]
	*h = old[0 : N-1]
	return x
}

// TC - O(N * LogK), SC - O(K)
func Constructor(k int, nums []int) Kthlargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		// If the heap size is greater than k, remove the smallest element
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return Kthlargest{k: k, heap: h}
}

// TC - O(LogK), SC - O(K)
func (kl *Kthlargest) Add(val int) int {
	heap.Push(kl.heap, val)
	if kl.heap.Len() > kl.k {
		heap.Pop(kl.heap) // Pop the smallest element to maintain the size of the heap as k
	}
	return (*kl.heap)[0] // The root of the min-heap is the kth largest element
}

func mainKL() {
	k := 3
	nums := []int{4, 5, 8, 2}
	Kthlargest := Constructor(k, nums)

	fmt.Printf("Add 3; kth largest = %d\n", Kthlargest.Add(3))   // Returns 4
	fmt.Printf("Add 5; kth largest = %d\n", Kthlargest.Add(5))   // Returns 5
	fmt.Printf("Add 10; kth largest = %d\n", Kthlargest.Add(10)) // Returns 5
	fmt.Printf("Add 9; kth largest = %d\n", Kthlargest.Add(9))   // Returns 8
	fmt.Printf("Add 4; kth largest = %d\n", Kthlargest.Add(4))   // Returns 8
}
