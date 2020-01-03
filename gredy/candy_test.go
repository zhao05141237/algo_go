package gredy

import "testing"

func TestCandy(t *testing.T) {
	//输出:8 预期:7
	//[1,2,87,87,87,2,1]
	//ratings := []int{1,2,87,87,87,2,1}
	//ratings := []int{1,2,87,87,87,2,1}
	//ratings := []int{1,3,2,3,1,1}
	//输出9 预期11
	//ratings := []int{1,3,4,5,2}
	//1+2+2+1+2+2+1
	//ratings := []int{1,3,2}
	//输出12 预期13
	//ratings := []int{0,1,2,3,2,1}
	//10 11
	ratings := []int{3, 2, 1, 1, 4, 3, 3}
	t.Log(Candy(ratings))
}
