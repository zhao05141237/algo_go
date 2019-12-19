package radixSort

import (
	"math"
)

func RadixStrSort(a []string) {
	maxStrBits := maxStrBits(a)
	c := make([][]string, 128)
	d := make([]int, 128)

	for i := 0; i < maxStrBits; i++ {
		for i := 0; i < 128; i++ {
			c[i] = nil
			d[i] = 0
		}
		for _, strValue := range a {
			var data int
			if len(strValue) < maxStrBits-i {
				data = 48
			} else {
				data = int(strValue[maxStrBits-i-1])
			}
			c[data] = append(c[data], strValue)
			d[data]++
		}

		t := 0
		for i := 0; i < 128; i++ {
			for j := 0; j < d[i]; j++ {
				a[t] = c[i][j]
				t++
			}
		}
	}
}

/**
获取数组中最大位数
*/
func maxStrBits(a []string) int {
	max := math.MinInt32
	for _, value := range a {
		r := strBits(value)
		if r >= max {
			max = r
		}
	}
	return max
}

func strBits(a string) int {
	return len(a)
}
