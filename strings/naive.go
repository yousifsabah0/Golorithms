package strings

// Naive implementation of Naive string search algorithm.
func Naive(str, pattern string) int {
	var strLength int = len(str)
	var patternLength int = len(pattern)

	for i := 0; i <= strLength-patternLength; i++ {
		var j int
		for j = 0; j < patternLength; j++ {
			if str[i+j] != pattern[j] {
				break
			}
		}

		if j == patternLength {
			return i
		}
	}
	return -1
}
