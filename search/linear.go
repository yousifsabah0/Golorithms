package search

// Linear Implementation of linear search algorithm.
// Algorithm: Given an array arr[] of n elements, write a function to search a given element x in arr[].
// Example:
// Input : arr[] = {10, 20, 80, 30, 60, 50,
//                     110, 100, 130, 170}
//          x = 110;
// Output : 6
// Element x is present at index 6
func Linear(arr []int, target int) int {
	for i := range arr {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
