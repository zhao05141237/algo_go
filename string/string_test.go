package string

import (
	"testing"
)

func TestBfSearch(t *testing.T) {
	t.Log(bfSearch("abcd227fac", "abcd"))
}

func TestRkHashValue(t *testing.T) {
	t.Log(hashValue("cba"))
}

func TestRkSearch(t *testing.T) {
	t.Log(rkSearch("abcdfac", "c"))
}

func TestBmSearch(t *testing.T) {
	mainStr := "abcacabcbcbacabc"
	patternStr := "cbacabc"
	mainLength := len(mainStr)
	patternLength := len(patternStr)
	t.Log(bmSearch(mainStr, mainLength, patternStr, patternLength))
	//t.Log(bmSearch("aaaaaaaaaaaaaaaa","baaa"))   //最坏前缀匹配为负数的情况
}

func TestBmFindGoodChar(t *testing.T) {
	t.Log(findGoodChar("abcacabcbcbacabc", "cbacabc"))
}

func TestBmGenerateGs(t *testing.T) {
	pattern := "aakdd"
	suffix := make([]int, len(pattern))
	prefix := make([]bool, len(pattern))
	bmGenerateGs(pattern, suffix, prefix)
	t.Log(suffix, prefix)
}

func TestGetNextGs(t *testing.T) {
	pattern := "abxabwabxad"
	t.Log(getNextGs(pattern))
}

func TestFindByKmp(t *testing.T) {
	s := "abc abcdab abcdabcdabde"
	pattern := "bcdabd"
	t.Log(findByKmp(s, pattern)) // 16

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "abab"
	t.Log(findByKmp(s, pattern)) //11

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "ababacd"
	t.Log(findByKmp(s, pattern)) //-1

	s = "hello"
	pattern = "ll"
	t.Log(findByKmp(s, pattern)) //2
}
