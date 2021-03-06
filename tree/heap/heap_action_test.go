package heap

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestHeap_InsertData2(t *testing.T) {
	a := []int{78, 14, 50, 20, 13, 9, 25, 8, 13, 37, 29, 33, 55, 52, 6, 17, 65, 23, 74, 43, 5, 29, 29, 72, 7, 13, 56, 21, 31, 66, 69, 69, 74, 12, 77, 23, 10, 6, 27, 63, 77, 21, 40, 10, 19, 59, 35, 40, 44, 4, 15, 29}
	h := NewHeap(len(a) + 1)

	for _, value := range a {
		h.InsertData(value, MaxTopSort)
	}

	for i := 1; i < len(a); i++ {
		h.RemoveTopData(MaxTopSort)
	}
	h.PrintHeap()
	//t.Log(h.a)
	//quickSort.QuickSort(a)
	//t.Log(a)
	//length := len(a)
	//if length % 2 == 0 {
	//	t.Log((float64(a[length/2 - 1]) + float64(a[length/2])) / float64(2))
	//}else {
	//	t.Log(a[length/2+1])
	//}
}

func TestHeap_InsertData(t *testing.T) {
	//堆排序简单版本
	h := NewHeap(51)
	//a := []int{23, 27, 30, 63, 73, 72, 67, 79, 95, 97}
	////建最大堆
	for i := 1; i <= 50; i++ {
		h.InsertData(rand.Intn(100), MaxTopSort)
	}
	//for _, value := range a {
	//	h.InsertData(value,MaxTopSort)
	//}
	h.PrintHeap()
	//每次删除顶元素 循环n-1次 从而形成从有序数组
	for i := 1; i < 50; i++ {
		h.RemoveTopData(MaxTopSort)
	}
	h.PrintHeap()
}

func TestMergedSortedData(t *testing.T) {
	//合并有序数组
	const MaxSize = 4
	a := [MaxSize][]int{
		{9, 10, 12, 30, 32, 32, 39, 61, 77, 87},
		{23, 27, 30, 63, 67, 72, 73, 79, 95, 97},
		{1, 5, 6, 18, 24, 56, 62, 65, 88, 92},
		{7, 15, 30, 31, 41, 43, 56, 71, 72, 97},
	}
	h := NewHeap(MaxSize + 1)
	b := make([]int, MaxSize)        //存放每个数组每次推进的位置
	countRow := make([]int, MaxSize) //存放每个数组的长度
	sumCountRow := 0
	var c []int //存放最终有序数组
	i := 0
	row := 0
	col := 0
	//首先用每个有序数组的首元素建最小堆
	for i < MaxSize {
		countRow[i] = len(a[i])
		sumCountRow += countRow[i]
		h.InsertData(a[i][0], MinTopSort)
		i++
	}
	//每次取堆顶元素放置最终排序数组
	for j := 0; j < sumCountRow; j++ {
		h.RemoveTopData(MinTopSort)   //取堆顶元素
		c = append(c, h.a[h.count+1]) //堆顶元素追加到最终数组
		for i = 0; i < MaxSize; i++ {
			//确定推进哪个有序数组
			if b[i] < countRow[i] && a[i][b[i]] == h.a[h.count+1] {
				b[i]++
				row = i
				col = b[i]
				break
			}
		}
		//某个数组遍历结束 则继续取堆顶元素
		if col >= countRow[i] {
			continue
		}
		//将数组添加到最小堆
		h.InsertData(a[row][col], MinTopSort)
	}
	t.Log(c)
	t.Log(len(c))
	t.Log(cap(c))
}

//topN 以有第k个元素值
func TestStaticDataTop(t *testing.T) {
	k := rand.Intn(10)
	if k == 0 {
		k = 1
	}
	var a []int
	for i := 1; i <= 10; i++ {
		a = append(a, rand.Intn(100))
	}
	t.Log(a)
	t.Log(k)

	//a := []int{38, 21, 55, 2, 60, 1, 55, 6, 93, 77}
	//k := 5

	h := NewHeap(k + 1)

	for index, value := range a {
		if index < k {
			h.InsertData(value, MinTopSort)
		} else {
			if value > h.a[1] {
				h.RemoveTopData(MinTopSort)
				h.InsertData(value, MinTopSort)
			}
		}
	}

	t.Logf("第%d大元素值为%d", k, h.a[1])
	for i := 1; i <= h.count; i++ {
		t.Log(h.a[i])
	}
}
