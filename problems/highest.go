package problems

// FindHighestElement program that efficiently finds the largest element present in the array.
func FindHighestElement(arr []int) int {
	n := len(arr)
	max := arr[0]
	for i := 0; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

// FindSecondHighestElement program that efficiently finds the second largest element present in the array.
func FindSecondHighestElement(arr []int) int {
	max := FindHighestElement(arr)
	arr = append(arr[:max], arr[max+1:]...)

	return FindHighestElement(arr)
}
