package search

import "math"

// Jump implementation of jump search algorithm.
// Algorithm: Given an sorted array arr[] of n elements, write a function that works by jumping multiple steps ahead in sorted list until it find an item larger than target
func Jump(arr []int, target int) int {
	n := len(arr)
	step := math.Floor(math.Sqrt(float64(n)))
	prev := 0

	for arr[int(math.Min(step, float64(n))-1)] < target {
		prev = int(step)
		step += math.Floor(math.Sqrt(math.Floor(float64(n))))
		if prev >= n {
			return -1
		}
	}

	for arr[prev] < target {
		prev++
		if prev == int(math.Min(step, float64(n))) {
			return -1
		}
	}

	if arr[prev] == target {
		return prev
	}

	return -1
}
