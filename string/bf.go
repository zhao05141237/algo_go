package string

func bfSearch(main string, pattern string) int {
	if len(main) == 0 || len(pattern) >= len(main) || len(pattern) == 0 {
		return -1
	}
	for i := 0; i <= len(main)-len(pattern); i++ {
		subStr := main[i : len(pattern)+i]
		if subStr == pattern {
			return i + 1
		}
	}
	return -1
}
