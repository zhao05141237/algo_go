package string

func getNextGs(pattern string) []int {
	patternLength := len(pattern)
	nexts := make([]int, patternLength)

	for key := range nexts {
		nexts[key] = -1
	}

	for i := 1; i < patternLength-1; i++ {
		j := nexts[i-1]
		for pattern[j+1] != pattern[i] && j >= 0 {
			j = nexts[j]
		}

		if pattern[j+1] == pattern[i] {
			j += 1
		}

		nexts[i] = j
	}

	return nexts

}

func findByKmp(main string, pattern string) int {
	mainLength := len(main)
	patternLength := len(pattern)

	if mainLength == 0 || patternLength >= mainLength || patternLength == 0 {
		return -1
	}
	nextGs := getNextGs(pattern)

	j := 0
	for i := 0; i < mainLength; i++ {
		for j > 0 && main[i] != pattern[j] {
			j = nextGs[j-1] + 1
		}
		if main[i] == pattern[j] {
			if j == patternLength-1 {
				return i - patternLength + 1
			}
			j++
		}
	}
	return -1
}
