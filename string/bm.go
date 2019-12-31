package string

import "math"

func bmGenerateGs(pattern string, suffix []int, prefix []bool) {
	patternLength := len(pattern)
	for key, _ := range suffix {
		suffix[key] = -1
	}
	var j, k int
	for i := 0; i < patternLength-1; i++ {
		j = i
		k = 0
		for j >= 0 && pattern[j] == pattern[patternLength-1-k] {
			k++
			j--
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[k] = true
		}
	}
}

func bmGenerateBc(pattern string) []int {
	bc := make([]int, 256)
	for key := range bc {
		bc[key] = -1
	}

	for index, value := range pattern {
		bc[value] = index
	}

	return bc
}

/**
坏字符匹配 当发生不匹配的时候，我们把坏字符对应的模式串中的字符下标记作 si。如果坏字符在模式串中存在，我们把这个坏字符在模式串中的下标记作 xi。如果不存在，我们把 xi 记作 -1。那模式串往后移动的位数就等于 si-xi。（注意，我这里说的下标，都是字符在模式串的下标）
不过，单纯使用坏字符规则还是不够的。因为根据 si-xi 计算出来的移动位数，有可能是负数，比如主串是 aaaaaaaaaaaaaaaa，模式串是 baaa。不但不会向后滑动模式串，还有可能倒退。
*/
func findBadChar(main string, pattern string) int {
	mainLength := len(main)
	patternLength := len(pattern)
	bc := bmGenerateBc(pattern)
	if mainLength == 0 || patternLength >= mainLength || patternLength == 0 {
		return -1
	}

	var j int
	for i := 0; i <= mainLength-patternLength; {
		for j = patternLength - 1; j >= 0; j-- {
			if main[i+j] != pattern[j] {
				break
			}
		}
		if j < 0 {
			return i + 1 //i 即为要找的位置
		}
		i += j - bc[main[i+j]] //此处是关键点 也可能出现负数的情况
	}
	return -1
}

func findGoodChar(main string, pattern string) int {
	mainLength := len(main)
	patternLength := len(pattern)
	suffix := make([]int, patternLength)
	prefix := make([]bool, patternLength)
	bmGenerateGs(pattern, suffix, prefix)
	if mainLength == 0 || patternLength >= mainLength || patternLength == 0 {
		return -1
	}
	var j int
	for i := 0; i <= mainLength-patternLength; {
		for j = patternLength - 1; j >= 0; j-- {
			if main[i+j] != pattern[j] {
				break
			}
		}
		if j < 0 {
			return i + 1 //i 即为要找的位置
		}
		y := moveByGS(j, patternLength, suffix, prefix)
		i += y
	}
	return -1

}

func bmSearch(main string, mainLength int, pattern string, patternLength int) int {
	bc := bmGenerateBc(pattern)
	suffix := make([]int, patternLength)
	prefix := make([]bool, patternLength)
	bmGenerateGs(pattern, suffix, prefix)
	if mainLength == 0 || patternLength >= mainLength || patternLength == 0 {
		return -1
	}
	var j int
	for i := 0; i <= mainLength-patternLength; {
		for j = patternLength - 1; j >= 0; j-- {
			if main[i+j] != pattern[j] {
				break
			}
		}
		if j < 0 {
			return i + 1 //i 即为要找的位置
		}
		x := j - bc[main[i+j]] //此处是关键点 也可能出现负数的情况
		y := 0
		if j < patternLength-1 {
			y = moveByGS(j, patternLength, suffix, prefix)
		}
		i += int(math.Max(float64(x), float64(y)))
	}
	return -1
}

func moveByGS(j int, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if prefix[r] {
			return r
		}
	}
	return m
}
