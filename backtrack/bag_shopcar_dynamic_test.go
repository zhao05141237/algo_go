package backtrack

import "fmt"

func ShopCarDynamicProgram(items []int) {
	c := make([][MaxPrice + 1]bool, NUMBERS)
	v := make([]int, 0)
	c[0][0] = true
	c[0][items[0]] = true
	for i := 1; i < NUMBERS; i++ {
		for j := 0; j <= MaxPrice; j++ {
			if c[i-1][j] {
				c[i][j] = true
			}
		}
		for j := 0; j <= MaxPrice-items[i]; j++ {
			if c[i-1][j] {
				c[i][j+items[i]] = true
			}
		}
	}
	var j int

	for j = GoodPrice; j <= MaxPrice; j++ {
		if c[NUMBERS-1][j] {
			break
		}
	}

	fmt.Printf("the value near %d is %d\n", GoodPrice, j)

	if j == MaxPrice {
		fmt.Println("have no")
		return
	}

	for i := NUMBERS - 1; i >= 1; i-- {
		if j-items[i] >= 0 && c[i-1][j-items[i]] {
			v = append(v, items[i])
			j = j - items[i]
		}
	}

	if j != 0 {
		v = append(v, items[0])
	}
	fmt.Println(v)
}
