package backtrack

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	MaxPrice  = 300
	GoodPrice = 200
)

var (
	MinPrice int
	Trace    [][]int
	TraceRow int
)

func init() {
	MinPrice = 300
	Trace = make([][]int, NUMBERS)
	TraceRow = 0
}

func TestBagShopCat(t *testing.T) {
	BagShopCar()
}

func BagShopCar() {
	rand.Seed(time.Now().UnixNano())
	items := make([]int, NUMBERS)
	before := make([]int, 0)
	for key := range items {
		items[key] = rand.Intn(CW)
	}
	fmt.Println(items)
	bagShopCarBackTrack(0, items, 0, before)
	fmt.Println(MinPrice)
	//fmt.Println(Trace)
	printTrace()
	ShopCarDynamicProgram(items)
}

func bagShopCarBackTrack(i int, items []int, cPrice int, before []int) {
	if i == NUMBERS || cPrice > MaxPrice {
		return
	}
	if cPrice < MinPrice && cPrice >= GoodPrice {
		Trace[TraceRow] = make([]int, len(before))
		copy(Trace[TraceRow], before)
		TraceRow++
		MinPrice = cPrice
	}
	bagShopCarBackTrack(i+1, items, cPrice, before)
	if cPrice+items[i] <= MinPrice {
		before = append(before, items[i])
		bagShopCarBackTrack(i+1, items, cPrice+items[i], before)
	}
}

func printTrace() {
	for i := 0; i < TraceRow; i++ {
		lengthTrace := len(Trace[i])
		sum := 0
		for j := 0; j < lengthTrace; j++ {
			sum += Trace[i][j]
			if sum > MinPrice {
				break
			}
		}
		if sum == MinPrice {
			fmt.Println(Trace[i])
		}
	}
}
