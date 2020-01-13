package backtrack

import (
	"fmt"
	"testing"
)

const ROW = 8

var index int

func TestEightQueue(t *testing.T) {
	EightQueue()
}

func EightQueue() {
	result := make([]int, ROW)
	for key := range result {
		result[key] = -1
	}
	calEightQueue(0, result)
}

func calEightQueue(i int, result []int) {
	if i == ROW {
		index++
		fmt.Println(index)
		printQueues(result)
		return
	}
	for column := 0; column < ROW; column++ {
		if isOk(i, column, result) {
			result[i] = column
			calEightQueue(i+1, result)
		}
	}
}

func isOk(row int, column int, result []int) bool {
	left, right := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		if result[i] == column {
			return false
		}
		if left >= 0 && result[i] == left {
			return false
		}
		if right < ROW && result[i] == right {
			return false
		}
		left--
		right++
	}
	return true
}

func printQueues(result []int) {
	for i := 0; i < ROW; i++ {
		for j := 0; j < ROW; j++ {
			if result[i] == j {
				fmt.Printf(" Q")
			} else {
				fmt.Printf(" *")
			}
		}
		fmt.Println()
	}
}
