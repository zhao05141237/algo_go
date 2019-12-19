package radixSort

import (
	"math"
)

/**
获取数组中最大位数
*/
func maxBits(a []int) int {
	max := math.MinInt32
	for _, value := range a {
		r := bits(value)
		if r >= max {
			max = r
		}
	}
	return max
}

func bits(a int) int {
	i := 0
	for {
		a = a / 10
		i++
		if a == 0 {
			break
		}
	}
	return i
}

func RadixSort(a []int) {
	maxBit := maxBits(a)

	number := 1
	c := make([][]int, 10)
	d := make([]int, 10)

	for i := 0; i < maxBit; i++ {
		for i := 0; i < 10; i++ {
			c[i] = nil
			d[i] = 0
		}
		for _, value := range a {
			data := value / number
			data = data % 10
			c[data] = append(c[data], value)
			d[data]++
		}
		t := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < d[i]; j++ {
				a[t] = c[i][j]
				t++
			}
		}
		number *= 10
	}
}
