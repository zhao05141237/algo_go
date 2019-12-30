package string

import "math"

func rkSearch(main string, pattern string) int {
	if len(main) == 0 || len(pattern) >= len(main) || len(pattern) == 0 {
		return -1
	}
	patternHashValue := hashValue(pattern)
	var mainHashValue float64
	for i := 0; i <= len(main)-len(pattern); i++ {
		subStr := main[i : len(pattern)+i]
		if i == 0 {
			mainHashValue = hashValue(subStr)
		} else {
			mainHashValue = afterHashValue(mainHashValue, string(main[i-1]), string(main[len(pattern)+i-1]), len(pattern))
		}
		if patternHashValue == mainHashValue {
			return i + 1
		}
	}
	return -1

}

func hashValue(str string) float64 {
	s := float64(0)
	length := len(str)
	for index, value := range str {
		b := float64(value - 'a')
		s += b * math.Pow(float64(26), float64(length-index-1))
	}
	return s
}

func afterHashValue(beforeHashValue float64, beforeStr string, afterStr string, m int) float64 {
	//h[i] = (h[i-1] - 26 ^(m-1) * s[i-1] - 'a') * 26 + s[i+m-1] - 'a'
	return (beforeHashValue-math.Pow(26, float64(m-1))*float64(beforeStr[0]-'a'))*26 + float64(afterStr[0]-'a')
}
