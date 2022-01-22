package search

// Binary Implementation of binary search algorithm.
// Algorithm: Given an sorted array arr[] of n elements, write a function to search a given element x in arr[].
func Binary(arr []int, target, low, high int) int {
	if high < low || len(arr) == 0 {
		return -1
	}
	// Get the middle element
	middle := int(low + (high-low)/2)

	switch {
	case arr[middle] == target:
		return middle
	case arr[middle] > target:
		return Binary(arr, target, low, middle+1)
	case arr[middle] < target:
		return Binary(arr, target, middle-1, high)
	default:
		return -1
	}
}
