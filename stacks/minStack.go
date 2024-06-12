package main

import (
	"math"
)

// https://leetcode.com/problems/min-stack/

type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{
		data:    []int{},
		minData: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)

	if val <= this.minData[len(this.minData)-1] {
		this.minData = append(this.minData, val)
	}
}

func (this *MinStack) Pop() {
	var top int
	if len(this.data) > 0 {
		top = this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]

		if top == this.minData[len(this.minData)-1] {
			this.minData = this.minData[:len(this.minData)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.minData) > 0 {
		return this.minData[len(this.minData)-1]
	}
	return -1
}
