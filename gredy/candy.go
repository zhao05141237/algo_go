package gredy

import "fmt"

/**
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
你需要按照以下要求，帮助老师给这些孩子分发糖果：
每个孩子至少分配到1个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？
示例1:
输入:[1,0,2]
输出:5
解释:你可以分别给这三个孩子分发 2、1、2 颗糖果。
示例2:
输入:[1,2,2]
输出:4
解释:你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
*/
func Candy(ratings []int) int {
	ratingsLength := len(ratings)
	result := make([]int, ratingsLength)
	for i := 0; i < ratingsLength; i++ {
		result[i] = 1
	}

	for i := 1; i < ratingsLength; i++ {
		if ratings[i-1] > ratings[i] {
			result[i-1] = result[i-1] + 1
			for j := i - 2; j >= 0; j-- {
				if ratings[j] > ratings[j+1] && result[j] < result[j+1]+1 {
					result[j] = result[j+1] + 1
				} else {
					break
				}
			}
		} else if ratings[i-1] < ratings[i] {
			result[i] = result[i-1] + 1
		}

		if i < ratingsLength-1 && ratings[i-1] < ratings[i] && ratings[i+1] < ratings[i] {
			result[i]--
		}
	}

	fmt.Println(result)

	sum := 0
	for i := 0; i < ratingsLength; i++ {
		sum += result[i]
	}
	return sum
}
