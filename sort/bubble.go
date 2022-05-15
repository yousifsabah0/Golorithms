package sort

type Number interface {
  int32 | int64 | float32 | float64
}


// Bubble implementation of bubble sort algorithm.
func Bubble[N Number] (arr []N) []N {
  n := len(arr)
  swapped := true
  
  for swapped {
    swapped = false
    for i := 0; i < n-1; i++ {
      if arr[i] > arr[i+1] {
        arr[i], arr[i+1] = arr[i+1], arr[i]
        swapped = true
      }
    }
  }
  
  return arr
}
