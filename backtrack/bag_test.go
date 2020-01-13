package backtrack

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

const (
	NUMBERS = 10
	CW      = 100
)

var maxCw int

func init() {
	maxCw = math.MinInt8
}

func TestBag(t *testing.T) {
	Bag()
}

func TestSum(t *testing.T) {
	Sum()
}

func Sum() {
	a := []int{33, 95, 28, 12, 74, 4, 78, 14, 50, 21}
	s(0, 0, a, NUMBERS, CW)
}

func Bag() {
	rand.Seed(time.Now().UnixNano())
	items := make([]int, NUMBERS)
	for key := range items {
		items[key] = rand.Intn(CW)
	}
	//items := []int{2,2,4,6,3}
	fmt.Println(items)
	f(0, 0, items, NUMBERS, CW)
	fmt.Println(maxCw)
	//s(0,0,items,NUMBERS,CW)

	dynamicProgramming(items)
}

func dynamicProgramming(items []int) {
	c := make([][CW + 1]bool, NUMBERS)
	for i := 0; i < NUMBERS; i++ {
		for j := 0; j <= CW; j++ {
			c[i][j] = false
		}
	}

	c[0][0] = true
	if items[0] < CW {
		c[0][items[0]] = true
	}

	for i := 1; i < NUMBERS; i++ {
		for j := 0; j <= CW; j++ {
			//选择不放入i
			if c[i-1][j] {
				c[i][j] = true
			}
		}
		//选择放入i
		for j := 0; j <= CW-items[i]; j++ {
			if c[i-1][j] {
				c[i][items[i]+j] = true
			}
		}
	}

	//fmt.Println(c)

	for i := CW; i >= 0; i-- {
		if c[NUMBERS-1][i] {
			fmt.Printf("Max sum is %d", i)
			break
		}
	}
}

func dynamicProgrammingAnother(items []int) {
	c := make([]bool, CW+1)
	for key, _ := range c {
		c[key] = false
	}
	c[0] = true
	c[items[0]] = true

	for i := 1; i < NUMBERS; i++ {
		for j := CW - items[i]; j >= 0; j-- {
			if c[j] {
				c[j+items[i]] = true
			}
		}
	}

	for i := CW; i >= 0; i-- {
		if c[i] {
			fmt.Printf("Max sum is %d", i)
			break
		}
	}
}

func f(i int, cw int, items []int, n int, w int) {
	if i >= n || cw >= w {
		if cw > maxCw {
			maxCw = cw
		}
		return
	}
	f(i+1, cw, items, n, w)
	if cw+items[i] <= w {
		f(i+1, cw+items[i], items, n, w)
	}
}

func s(i int, cw int, items []int, n int, w int) {
	if i >= n || cw == maxCw {
		if cw == maxCw {
			fmt.Println(maxCw)
		}
		return
	}
	s(i+1, cw, items, n, w)
	if cw+items[i] <= w {
		//fmt.Printf("i:%d,weight:%d,cw:%d\n",i,items[i],cw)
		s(i+1, cw+items[i], items, n, w)
	} else {
		//fmt.Printf("out i:%d,weight:%d,cw:%d\n",i,items[i],cw)
	}
}
