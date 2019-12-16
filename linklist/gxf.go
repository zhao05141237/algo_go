package main

import (
	"fmt"
	"math"
)

func isLower(s string) bool {
	if s >= "a" && s <="z" {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(math.Pow(10,1))
	//s := []string{"a","B","B","B","c","K","M","J","d","B","C","D","a","K","J"}
	//countLeft,countRight := 0,0
	//for index,val:=range s {
	//	if countLeft == 3 && isLower(val){
	//		countRight = 0
	//		for i:=1; i <= 4; i++ {
	//			if index + i < len(s) && !isLower(s[index + i]) {
	//				countRight++
	//			}else {
	//				break
	//			}
	//		}
	//		if countRight == 3 {
	//			fmt.Println(val)
	//		}
	//	}
	//
	//	if !isLower(val) {
	//		countLeft++
	//	} else {
	//		countLeft = 0
	//	}
	//}
}
