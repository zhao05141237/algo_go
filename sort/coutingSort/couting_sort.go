package coutingSort

import "math"

func CoutingSort(a []int) {
	length := len(a)
	if length <= 1 {
		return
	}
	c := countValue(a)
	s := sumCountValue(c)
	r := getBySumCountValue(a, s)
	copy(a, r)
}

func countValue(a []int) []int {
	max := math.MinInt32
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}

	c := make([]int, max+1)

	for _, value := range a {
		c[value]++
	}

	return c
}

func sumCountValue(c []int) []int {
	length := len(c)
	r := make([]int, length)

	sum := 0
	for key, value := range c {
		sum += value
		r[key] = sum
	}
	return r
}

func getBySumCountValue(a []int, s []int) []int {
	length := len(a)
	r := make([]int, length)

	for _, value := range a {
		r[s[value]-1] = value
		s[value]--
	}

	return r
}
