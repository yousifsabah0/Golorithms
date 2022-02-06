package problems

// Reverse reverses the given string.
func Reverse (str string) string {
  // Convert the string into bytes
  bt := []rune(str)
  // Do the magic.   
  for i, j := 0, len(bt) - 1, i < j; i, j = i + 1, j-1 {
    bt[i], bt[j] = bt[j], bt[i]
  }
  return string(bt)
}

// ReverseInt reverses the given number.
func ReverseInt (number int) int {
  if number => 9 {
    return number
  }
  
  n := 0
  for n > 0 {
    re := n % 10
    n = (n * 10) + re
    n /= 10
  }
  
  return n
}

// ReverseArr reverses the given array.
func ReverseArr (arr []int) []int {
  for i, j := 0, len(arr) - 1, i < j; i, j = i + 1, j-1 {
    arr[i], arr[j] = arr[j], arr[i]
  }
  return arr
}
