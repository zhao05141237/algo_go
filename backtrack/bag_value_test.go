package backtrack

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

var MaxCValue int

func init() {
	MaxCValue = math.MinInt8
}

func TestBagValue(t *testing.T) {
	BagValue()
}

func BagValue() {
	rand.Seed(time.Now().UnixNano())
	items := make([]int, NUMBERS)
	value := make([]int, NUMBERS)
	for key := range items {
		items[key], value[key] = rand.Intn(CW), rand.Intn(CW)
	}
	//items := []int{2,2,4,6,3}
	//value := []int{3,4,8,9,6}
	fmt.Println(items)
	fmt.Println(value)
	bagValueBackTrack(0, items, value, 0, 0)
	fmt.Println(MaxCValue)
	bagValueDynamicProgram(items, value)
	bagValueDynamicProgramAnother(items, value)
}

func bagValueBackTrack(i int, items []int, value []int, w int, cValue int) {
	if i == NUMBERS || w == CW {
		if cValue > MaxCValue {
			MaxCValue = cValue
		}
		return
	}

	bagValueBackTrack(i+1, items, value, w, cValue)

	if w+items[i] <= CW {
		bagValueBackTrack(i+1, items, value, w+items[i], cValue+value[i])
	}
}

func bagValueDynamicProgram(items []int, value []int) {
	c := make([][CW + 1]int, NUMBERS)
	c[0][0] = 0
	c[0][items[0]] = value[0]

	for i := 1; i < NUMBERS; i++ {
		for j := 0; j <= CW; j++ {
			if c[i-1][j] >= 0 {
				c[i][j] = c[i-1][j]
			}
		}
		for j := 0; j <= CW-items[i]; j++ {
			if c[i-1][j] >= 0 {
				v := c[i-1][j] + value[i]
				if v > c[i][j+items[i]] {
					c[i][j+items[i]] = v
				}
			}
		}
	}

	//fmt.Println(c)

	MaxCValue = math.MinInt8

	for j := CW; j >= 0; j-- {
		if c[NUMBERS-1][j] > MaxCValue {
			MaxCValue = c[NUMBERS-1][j]
		}
	}

	fmt.Println(MaxCValue)
}

func bagValueDynamicProgramAnother(items []int, value []int) {
	c := make([]int, CW+1)
	c[items[0]] = value[0]

	for i := 1; i < NUMBERS; i++ {
		for j := CW - items[i]; j >= 0; j-- {
			if c[j] >= 0 {
				v := c[j] + value[i]
				if v > c[j+items[i]] {
					c[j+items[i]] = v
				}
			}
		}
	}

	//fmt.Println(c)

	MaxCValue = math.MinInt8

	for j := CW; j >= 0; j-- {
		if c[j] > MaxCValue {
			MaxCValue = c[j]
		}
	}

	fmt.Println(MaxCValue)
}
