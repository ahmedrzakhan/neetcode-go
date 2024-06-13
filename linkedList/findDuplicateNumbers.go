package main

import "fmt"

// TC - O(N), SC - O(1)
func findDuplicate(A []int) int {
	slow, fast := A[0], A[A[0]]

	// Phase 1: Finding the intersection point of the two runners.
	for slow != fast {
		slow = A[slow]
		fast = A[A[fast]]
	}

	// Phase 2: Finding the entrance to the cycle.
	slow = 0
	for slow != fast {
		slow = A[slow]
		fast = A[fast]
	}
	return slow
}

func mainDup() {
	A := []int{1, 3, 4, 2, 2}
	fmt.Println("The duplicate number is:", findDuplicate(A))
}

/**
Why This Works
Cycle Guarantee: The duplicate number creates a loop because multiple indices will point to the same value.
Intersection Inevitability: In a cycle, a slow and fast runner will always eventually meet.
Cycle Entrance Property: The distance from the start of the array to the cycle entrance is the same as the distance from the meeting point within the cycle to the entrance. This is the key to finding the duplicate in the second phase.
Example
Let's say nums = [1, 3, 4, 2, 2].
The array forms a linked list: 0 -> 1 -> 3 -> 2 -> 4 -> 2 (cycle).
The pointers will meet at index 2 (value 4).
After resetting slow, both pointers move one step at a time and meet again at index 4 (value 2), the duplicate.
*/
