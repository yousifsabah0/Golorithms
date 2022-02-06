package sort

// Bubble implementation of bubble sort algorithm.
func Bubble (arr []int) []int {
  for i := 0; i < len(arr) - 1; i++ {
    for k := 0; k < len(arr) - 1 - i; k++ {
      if arr[i] < arr[k + 1] {
        arr[i], arr[k + 1] = arr[k + 1], arr[i]
      }
    }
  }
  return arr
}
